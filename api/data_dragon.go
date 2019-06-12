package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/KnutZuidema/golio/model"

	log "github.com/sirupsen/logrus"
)

const (
	latestRuneAndMasteryVersion = "7.23.1"
	fallbackVersion             = "9.10.1"
	fallbackLanguage            = LanguageCodeUnitedStates
)

var (
	regionToRealmRegion = map[Region]string{
		RegionEuropeWest:        "euw",
		RegionEuropeNorthEast:   "eun",
		RegionJapan:             "jp",
		RegionKorea:             "kr",
		RegionLatinAmericaNorth: "lan",
		RegionLatinAmericaSouth: "las",
		RegionNorthAmerica:      "na",
		RegionOceania:           "oce",
		RegionPBE:               "pbe",
		RegionRussia:            "ru",
		RegionTurkey:            "tr",
		RegionBrasil:            "br",
	}
)

// DataDragonClient provides access to all data provided by the Data Dragon service
type DataDragonClient struct {
	logger   log.FieldLogger
	Version  string
	Language languageCode
	client   Doer
}

// NewDataDragonClient returns a new client for the Data Dragon service.
func NewDataDragonClient(client Doer, region Region, logger log.FieldLogger) *DataDragonClient {
	c := &DataDragonClient{
		client: client,
		logger: logger.WithField("client", "data dragon"),
	}
	if err := c.init(regionToRealmRegion[region]); err != nil {
		c.Version = fallbackVersion
		c.Language = fallbackLanguage
	}
	return c
}

func (c *DataDragonClient) init(region string) error {
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
func (c DataDragonClient) GetChampions() (map[string]model.ChampionData, error) {
	var champions map[string]model.ChampionData
	if err := c.getInto("/champion.json", &champions); err != nil {
		return nil, err
	}
	return champions, nil
}

// GetChampion returns information about the champion with the given name
func (c DataDragonClient) GetChampion(name string) (*model.ChampionDataExtended, error) {
	var data map[string]*model.ChampionDataExtended
	if err := c.getInto(fmt.Sprintf("/champion/%s.json", name), &data); err != nil {
		return nil, err
	}
	if data, ok := data[name]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("response does not contain requested champion data")
}

// GetProfileIcons returns all existing profile icons
func (c DataDragonClient) GetProfileIcons() (map[string]model.ProfileIcon, error) {
	var icons map[string]model.ProfileIcon
	if err := c.getInto("/profileicon.json", &icons); err != nil {
		return nil, err
	}
	return icons, nil
}

// GetItems returns all existing items
func (c DataDragonClient) GetItems() (map[string]model.Item, error) {
	var items map[string]model.Item
	if err := c.getInto("/item.json", &items); err != nil {
		return nil, err
	}
	return items, nil
}

// GetMasteries returns all existing masteries. Masteries were removed in patch 7.23.1. If any version higher than that
// is specified the last available version will be used instead.
func (c DataDragonClient) GetMasteries() (map[string]model.Mastery, error) {
	var masteries map[string]model.Mastery
	if err := c.getInto("/mastery.json", &masteries); err != nil {
		return nil, err
	}
	return masteries, nil
}

// GetSummonerSpells returns all existing summoner spells
func (c DataDragonClient) GetSummonerSpells() (map[string]model.SummonerSpell, error) {
	var spells map[string]model.SummonerSpell
	if err := c.getInto("/summoner.json", &spells); err != nil {
		return nil, err
	}
	return spells, nil
}

func (c DataDragonClient) getInto(endpoint string, target interface{}) error {
	response, err := c.doRequest(dataDragonDataURLFormat, endpoint)
	if err != nil {
		return err
	}
	var ddResponse model.DataDragonResponse
	if err = json.NewDecoder(response.Body).Decode(&ddResponse); err != nil {
		return err
	}
	data, err := json.Marshal(ddResponse.Data)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &target)
}

func (c DataDragonClient) doRequest(format dataDragonURL, endpoint string) (*http.Response, error) {
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
		err, ok := StatusToError[response.StatusCode]
		if !ok {
			err = Error{
				Message:    "unknown error reason",
				StatusCode: response.StatusCode,
			}
		}
		return nil, err
	}
	return response, nil
}

func (c DataDragonClient) newRequest(format dataDragonURL, endpoint string) (*http.Request, error) {
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
