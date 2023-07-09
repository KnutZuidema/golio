// Package datadragon provides methods for retrieving data from the DataDragon API.
// This data is only updated for every new version of League of Legends.
package datadragon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	log "github.com/sirupsen/logrus"

	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
)

const (
	latestRuneAndMasteryVersion = "7.23.1"
	fallbackVersion             = "9.10.1"
	fallbackLanguage            = LanguageCodeUnitedStates
)

var (
	regionToRealmRegion = map[api.Region]string{
		api.RegionEuropeWest:        "euw",
		api.RegionEuropeNorthEast:   "eun",
		api.RegionJapan:             "jp",
		api.RegionKorea:             "kr",
		api.RegionLatinAmericaNorth: "lan",
		api.RegionLatinAmericaSouth: "las",
		api.RegionNorthAmerica:      "na",
		api.RegionOceania:           "oce",
		api.RegionPBE:               "pbe",
		api.RegionRussia:            "ru",
		api.RegionTurkey:            "tr",
		api.RegionBrasil:            "br",
	}
)

// Client provides access to all data provided by the Data Dragon service
type Client struct {
	logger             log.FieldLogger
	Version            string
	Language           languageCode
	client             internal.Doer
	championsMu        sync.RWMutex
	championsByName    map[string]ChampionDataExtended
	getChampionsToggle uint32
	profileIconsMu     sync.RWMutex
	profileIcons       []ProfileIcon
	itemsMu            sync.RWMutex
	items              []Item
	masteriesMu        sync.RWMutex
	masteries          []Mastery
	runesMu            sync.RWMutex
	runes              []Item
	summonersMu        sync.RWMutex
	summoners          []SummonerSpell
}

// NewClient returns a new client for the Data Dragon service.
func NewClient(client internal.Doer, region api.Region, logger log.FieldLogger) *Client {
	c := &Client{
		client:          client,
		logger:          logger.WithField("client", "data dragon"),
		championsByName: map[string]ChampionDataExtended{},
	}
	if err := c.init(regionToRealmRegion[region]); err != nil {
		c.Version = fallbackVersion
		c.Language = fallbackLanguage
	}
	return c
}

func (c *Client) init(region string) error {
	var res struct {
		Version  string `json:"v"`
		Language string `json:"l"`
	}
	response, err := c.doRequest(dataDragonBaseURL, fmt.Sprintf("/realms/%s.json", region))
	if err != nil {
		return err
	}
	if response.Body == nil {
		return fmt.Errorf("no response body")
	}
	if err := json.NewDecoder(response.Body).Decode(&res); err != nil {
		return err
	}
	c.Version = res.Version
	c.Language = languageCode(res.Language)
	return nil
}

// GetChampions returns all existing champions
func (c *Client) GetChampions() ([]ChampionData, error) {
	unlock, toggle := internal.RWLockToggle(&c.championsMu)
	defer unlock()
	if atomic.CompareAndSwapUint32(&c.getChampionsToggle, 0, 1) {
		toggle()
		var champions map[string]ChampionData
		if err := c.getInto("/champion.json", &champions); err != nil {
			return nil, err
		}
		for _, champion := range champions {
			data := ChampionDataExtended{ChampionData: champion}
			c.championsByName[champion.Name] = data
		}
	}
	res := make([]ChampionData, 0, len(c.championsByName))
	for _, champion := range c.championsByName {
		res = append(res, champion.ChampionData)
	}
	return res, nil
}

// GetChampionByID returns information about the champion with the given id
func (c *Client) GetChampionByID(id string) (ChampionDataExtended, error) {
	champions, err := c.GetChampions()
	if err != nil {
		return ChampionDataExtended{}, err
	}
	for _, champion := range champions {
		if champion.Key == id {
			return c.GetChampion(champion.Name)
		}
	}
	return ChampionDataExtended{}, api.ErrNotFound
}

// GetChampion returns information about the champion with the given name
func (c *Client) GetChampion(name string) (ChampionDataExtended, error) {
	unlock, toggle := internal.RWLockToggle(&c.championsMu)
	defer unlock()
	champion, ok := c.championsByName[name]
	if !ok || champion.Lore == "" {
		toggle()
		var data map[string]ChampionDataExtended
		if err := c.getInto(fmt.Sprintf("/champion/%s.json", name), &data); err != nil {
			return ChampionDataExtended{}, err
		}
		champion, ok = data[name]
		if !ok {
			return ChampionDataExtended{}, api.ErrNotFound
		}
		c.championsByName[name] = champion
	}
	return champion, nil
}

// GetProfileIcons returns all existing profile icons
func (c *Client) GetProfileIcons() ([]ProfileIcon, error) {
	unlock, toggle := internal.RWLockToggle(&c.profileIconsMu)
	defer unlock()
	if len(c.profileIcons) < 1 {
		toggle()
		var res map[string]ProfileIcon
		if err := c.getInto("/profileicon.json", &res); err != nil {
			return nil, err
		}
		c.profileIcons = make([]ProfileIcon, 0, len(res))
		for _, profileIcon := range res {
			c.profileIcons = append(c.profileIcons, profileIcon)
		}
	}
	res := make([]ProfileIcon, len(c.profileIcons))
	copy(res, c.profileIcons)
	return res, nil
}

// GetProfileIcon return information about the profile icon with the given id
func (c *Client) GetProfileIcon(id int) (ProfileIcon, error) {
	icons, err := c.GetProfileIcons()
	if err != nil {
		return ProfileIcon{}, err
	}
	for _, icon := range icons {
		if int(icon.ID) == id {
			return icon, nil
		}
	}
	return ProfileIcon{}, api.ErrNotFound
}

// GetItems returns all existing items
func (c *Client) GetItems() ([]Item, error) {
	unlock, toggle := internal.RWLockToggle(&c.itemsMu)
	defer unlock()
	if len(c.items) < 1 {
		toggle()
		var res map[string]Item
		if err := c.getInto("/item.json", &res); err != nil {
			return nil, err
		}
		c.items = make([]Item, 0, len(res))
		for id, item := range res {
			item.ID = id
			c.items = append(c.items, item)
		}
	}
	res := make([]Item, len(c.items))
	copy(res, c.items)
	return res, nil
}

// GetItem return information about the item with the given id
func (c *Client) GetItem(id string) (Item, error) {
	items, err := c.GetItems()
	if err != nil {
		return Item{}, err
	}
	for _, item := range items {
		if item.ID == id {
			return item, nil
		}
	}
	return Item{}, api.ErrNotFound
}

// GetMasteries returns all existing masteries. Masteries were removed in patch 7.23.1. If any version higher than that
// is specified the last available version will be used instead.
func (c *Client) GetMasteries() ([]Mastery, error) {
	unlock, toggle := internal.RWLockToggle(&c.masteriesMu)
	defer unlock()
	if len(c.masteries) < 1 {
		toggle()
		var res map[string]Mastery
		if err := c.getInto("/mastery.json", &res); err != nil {
			return nil, err
		}
		c.masteries = make([]Mastery, 0, len(res))
		for _, mastery := range res {
			c.masteries = append(c.masteries, mastery)
		}
	}
	res := make([]Mastery, len(c.masteries))
	copy(res, c.masteries)
	return res, nil
}

// GetMastery returns information about the mastery with the given id
func (c *Client) GetMastery(id int) (Mastery, error) {
	masteries, err := c.GetMasteries()
	if err != nil {
		return Mastery{}, err
	}
	for _, mastery := range masteries {
		if mastery.ID == id {
			return mastery, nil
		}
	}
	return Mastery{}, api.ErrNotFound
}

// GetRunes returns all existing runes. Runes were removed in patch 7.23.1. If any version higher than that
// is specified the last available version will be used instead.
func (c *Client) GetRunes() ([]Item, error) {
	unlock, toggle := internal.RWLockToggle(&c.runesMu)
	defer unlock()
	if len(c.runes) < 1 {
		toggle()
		var res map[string]Item
		if err := c.getInto("/rune.json", &res); err != nil {
			return nil, err
		}
		c.runes = make([]Item, 0, len(res))
		for id, runeItem := range res {
			runeItem.ID = id
			c.runes = append(c.runes, runeItem)
		}
	}
	res := make([]Item, len(c.runes))
	copy(res, c.runes)
	return res, nil
}

// GetRune returns information about the rune with the given id
func (c *Client) GetRune(id string) (Item, error) {
	runes, err := c.GetRunes()
	if err != nil {
		return Item{}, err
	}
	for _, r := range runes {
		if r.ID == id {
			return r, nil
		}
	}
	return Item{}, api.ErrNotFound
}

// GetSummonerSpells returns all existing summoner spells
func (c *Client) GetSummonerSpells() ([]SummonerSpell, error) {
	unlock, toggle := internal.RWLockToggle(&c.summonersMu)
	defer unlock()
	if len(c.summoners) < 1 {
		toggle()
		var res map[string]SummonerSpell
		if err := c.getInto("/summoner.json", &res); err != nil {
			return nil, err
		}
		c.summoners = make([]SummonerSpell, 0, len(res))
		for _, summoner := range res {
			c.summoners = append(c.summoners, summoner)
		}
	}
	res := make([]SummonerSpell, len(c.summoners))
	copy(res, c.summoners)
	return res, nil
}

// GetSummonerSpell returns information about the summoner spell with the given id
func (c *Client) GetSummonerSpell(id string) (SummonerSpell, error) {
	summonerSpells, err := c.GetSummonerSpells()
	if err != nil {
		return SummonerSpell{}, err
	}
	for _, summonerSpell := range summonerSpells {
		if summonerSpell.Key == id {
			return summonerSpell, nil
		}
	}
	return SummonerSpell{}, api.ErrNotFound
}

// ClearCaches resets all caches of the data dragon client
func (c *Client) ClearCaches() {
	c.championsMu.Lock()
	c.championsByName = map[string]ChampionDataExtended{}
	atomic.StoreUint32(&c.getChampionsToggle, 0)
	c.championsMu.Unlock()
	c.masteriesMu.Lock()
	c.masteries = []Mastery{}
	c.masteriesMu.Unlock()
	c.profileIconsMu.Lock()
	c.profileIcons = []ProfileIcon{}
	c.profileIconsMu.Unlock()
	c.itemsMu.Lock()
	c.items = []Item{}
	c.itemsMu.Unlock()
	c.summonersMu.Lock()
	c.summoners = []SummonerSpell{}
	c.summonersMu.Unlock()
	c.runesMu.Lock()
	c.runes = []Item{}
	c.runesMu.Unlock()
}

func (c *Client) getInto(endpoint string, target interface{}) error {
	response, err := c.doRequest(dataDragonDataURLFormat, endpoint)
	if err != nil {
		return err
	}
	var ddResponse dataDragonResponse
	if err = json.NewDecoder(response.Body).Decode(&ddResponse); err != nil {
		return err
	}
	// this can not return an error. the error would have been returned during the above decode already
	data, _ := json.Marshal(ddResponse.Data)
	return json.Unmarshal(data, &target)
}

func (c *Client) doRequest(format dataDragonURL, endpoint string) (*http.Response, error) {
	request, err := c.newRequest(format, endpoint)
	if err != nil {
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		var err error
		err, ok := api.StatusToError[response.StatusCode]
		if !ok {
			err = api.Error{
				Message:    "unknown error reason",
				StatusCode: response.StatusCode,
			}
		}
		return nil, err
	}
	return response, nil
}

func (c *Client) newRequest(format dataDragonURL, endpoint string) (*http.Request, error) {
	var version string
	if (strings.Contains(endpoint, "rune") || strings.Contains(endpoint, "mastery")) &&
		versionGreaterThan(c.Version, latestRuneAndMasteryVersion) {
		version = latestRuneAndMasteryVersion
	} else {
		version = c.Version
	}
	var url string
	switch format {
	case dataDragonDataURLFormat:
		url = fmt.Sprintf(string(format), version, c.Language)
	case dataDragonImageURLFormat:
		url = fmt.Sprintf(string(format), version)
	default:
		url = string(format)
	}
	url = "https://" + url + endpoint
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return request, nil
}

func versionGreaterThan(v1, v2 string) bool {
	v1Split := strings.Split(v1, ".")
	v2Split := strings.Split(v2, ".")
	for i := 0; i < len(v1Split) && i < len(v2Split); i++ {
		int1, err := strconv.Atoi(v1Split[i])
		if err != nil {
			return false
		}
		int2, err := strconv.Atoi(v2Split[i])
		if err != nil {
			return false
		}
		if int1 > int2 {
			return true
		}
	}
	return false
}

type dataDragonResponse struct {
	Type    string
	Format  string
	Version string
	Data    interface{}
}
