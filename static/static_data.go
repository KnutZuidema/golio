// Package static provides methods to access static data and constant values used by the Riot API.
// These values will rarely be updated, only if e.g. a new season starts or a new game mode is added.
package static

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/yigithanbalci/golio/api"
	"github.com/yigithanbalci/golio/internal"
)

// Client provides access to static data provided by Riot
// data is fetched on the first call to each method and cached for further calls
type Client struct {
	logger  logrus.FieldLogger
	client  internal.Doer
	mutexes map[string]*sync.RWMutex
	cache   map[string]interface{}
}

// NewClient returns a new client
func NewClient(doer internal.Doer, logger logrus.FieldLogger) *Client {
	mutexes := map[string]*sync.RWMutex{
		"seasons":   {},
		"queues":    {},
		"maps":      {},
		"gameModes": {},
		"gameTypes": {},
	}
	return &Client{
		logger:  logger,
		client:  doer,
		mutexes: mutexes,
		cache:   map[string]interface{}{},
	}
}

// GetSeasons returns static data for seasons
func (c *Client) GetSeasons() ([]Season, error) {
	mu := c.mutexes["seasons"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	seasons, ok := c.cache["seasons"].([]Season)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointSeasons, &seasons); err != nil {
			return nil, err
		}
		c.cache["seasons"] = seasons
	}
	res := make([]Season, len(seasons))
	copy(res, seasons)
	return res, nil
}

// GetSeason returns the season for the specified id or an error if no season for the id exists
func (c *Client) GetSeason(id int) (Season, error) {
	seasons, err := c.GetSeasons()
	if err != nil {
		return Season{}, err
	}
	for _, season := range seasons {
		if season.ID == id {
			return season, nil
		}
	}
	return Season{}, api.ErrNotFound
}

// GetQueues returns static data for queues
func (c *Client) GetQueues() ([]Queue, error) {
	mu := c.mutexes["queues"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	queues, ok := c.cache["queues"].([]Queue)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointQueues, &queues); err != nil {
			return nil, err
		}
		c.cache["queues"] = queues
	}
	res := make([]Queue, len(queues))
	copy(res, queues)
	return res, nil
}

// GetQueue returns the queue for the specified id or an error if no queue for the id exists
func (c *Client) GetQueue(id int) (Queue, error) {
	queues, err := c.GetQueues()
	if err != nil {
		return Queue{}, err
	}
	for _, queue := range queues {
		if queue.ID == id {
			return queue, nil
		}
	}
	return Queue{}, api.ErrNotFound
}

// GetMaps returns static data for maps
func (c *Client) GetMaps() ([]Map, error) {
	mu := c.mutexes["maps"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	maps, ok := c.cache["maps"].([]Map)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointMaps, &maps); err != nil {
			return nil, err
		}
		c.cache["maps"] = maps
	}
	res := make([]Map, len(maps))
	copy(res, maps)
	return res, nil
}

// GetMap returns the map for the specified id or an error if no map for the id exists
func (c *Client) GetMap(id int) (Map, error) {
	mapps, err := c.GetMaps()
	if err != nil {
		return Map{}, err
	}
	for _, mapp := range mapps {
		if mapp.ID == id {
			return mapp, nil
		}
	}
	return Map{}, api.ErrNotFound
}

// GetGameModes returns static data for game modes
func (c *Client) GetGameModes() ([]GameMode, error) {
	mu := c.mutexes["gameModes"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	gameModes, ok := c.cache["gameModes"].([]GameMode)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointGameModes, &gameModes); err != nil {
			return nil, err
		}
		c.cache["gameModes"] = gameModes
	}
	res := make([]GameMode, len(gameModes))
	copy(res, gameModes)
	return res, nil
}

// GetGameMode returns the game mode for the specified id or an error if no mode for the id exists
func (c *Client) GetGameMode(mode string) (GameMode, error) {
	modes, err := c.GetGameModes()
	if err != nil {
		return GameMode{}, err
	}
	for _, mo := range modes {
		if mo.Mode == mode {
			return mo, nil
		}
	}
	return GameMode{}, api.ErrNotFound
}

// GetGameTypes returns static data for game types
func (c *Client) GetGameTypes() ([]GameType, error) {
	mu := c.mutexes["gameTypes"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	gameTypes, ok := c.cache["gameTypes"].([]GameType)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointGameTypes, &gameTypes); err != nil {
			return nil, err
		}
		c.cache["gameTypes"] = gameTypes
	}
	res := make([]GameType, len(gameTypes))
	copy(res, gameTypes)
	return res, nil
}

// GetGameType returns the game type for the specified id or an error if no type for the id exists
func (c *Client) GetGameType(typ string) (GameType, error) {
	types, err := c.GetGameTypes()
	if err != nil {
		return GameType{}, err
	}
	for _, mo := range types {
		if mo.Type == typ {
			return mo, nil
		}
	}
	return GameType{}, api.ErrNotFound
}

// ClearCaches clears caches for all methods
func (c *Client) ClearCaches() {
	c.cache = map[string]interface{}{}
}

func (c *Client) getInto(endpoint string, target interface{}) error {
	req, _ := http.NewRequest(http.MethodGet, endpoint, nil)
	resp, err := c.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err, ok := api.StatusToError[resp.StatusCode]
		if !ok {
			err = api.Error{
				Message:    "unknown error reason",
				StatusCode: resp.StatusCode,
			}
		}
		return err
	}
	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return err
	}
	return nil
}
