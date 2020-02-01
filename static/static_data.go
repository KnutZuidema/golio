// Package static provides methods to access static data and constant values used by the Riot API.
// These values will rarely be updated, only if e.g. a new season starts or a new game mode is added.
package static

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/model"
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
func (c *Client) GetSeasons() ([]model.Season, error) {
	mu := c.mutexes["seasons"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	seasons, ok := c.cache["seasons"].([]model.Season)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointSeasons, &seasons); err != nil {
			return nil, err
		}
		c.cache["seasons"] = seasons
	}
	res := make([]model.Season, len(seasons))
	copy(res, seasons)
	return res, nil
}

// GetQueues returns static data for queues
func (c *Client) GetQueues() ([]model.Queue, error) {
	mu := c.mutexes["queues"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	queues, ok := c.cache["queues"].([]model.Queue)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointQueues, &queues); err != nil {
			return nil, err
		}
		c.cache["queues"] = queues
	}
	res := make([]model.Queue, len(queues))
	copy(res, queues)
	return res, nil
}

// GetMaps returns static data for maps
func (c *Client) GetMaps() ([]model.Map, error) {
	mu := c.mutexes["maps"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	maps, ok := c.cache["maps"].([]model.Map)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointMaps, &maps); err != nil {
			return nil, err
		}
		c.cache["maps"] = maps
	}
	res := make([]model.Map, len(maps))
	copy(res, maps)
	return res, nil
}

// GetGameModes returns static data for game modes
func (c *Client) GetGameModes() ([]model.GameMode, error) {
	mu := c.mutexes["gameModes"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	gameModes, ok := c.cache["gameModes"].([]model.GameMode)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointGameModes, &gameModes); err != nil {
			return nil, err
		}
		c.cache["gameModes"] = gameModes
	}
	res := make([]model.GameMode, len(gameModes))
	copy(res, gameModes)
	return res, nil
}

// GetGameTypes returns static data for game types
func (c *Client) GetGameTypes() ([]model.GameType, error) {
	mu := c.mutexes["gameTypes"]
	unlock, toggle := internal.RWLockToggle(mu)
	defer unlock()
	gameTypes, ok := c.cache["gameTypes"].([]model.GameType)
	if !ok {
		toggle()
		if err := c.getInto(staticDataEndpointGameTypes, &gameTypes); err != nil {
			return nil, err
		}
		c.cache["gameTypes"] = gameTypes
	}
	res := make([]model.GameType, len(gameTypes))
	copy(res, gameTypes)
	return res, nil
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
