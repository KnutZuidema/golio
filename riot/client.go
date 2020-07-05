// Package riot provides methods for accessing the Riot API for League of Legends.
// This includes dynamic data like the current game a summoner is in or their ranked standing.
package riot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
)

// Client provides access to all Riot API endpoints
type Client struct {
	l      log.FieldLogger
	Region api.Region
	apiKey string
	client internal.Doer
	// Deprecated: Use Client.LoL.ChampionMastery instead. Will be removed in a future release.
	ChampionMastery *ChampionMasteryClient
	// Deprecated: Use Client.LoL.Champion instead. Will be removed in a future release.
	Champion *ChampionClient
	// Deprecated: Use Client.LoL.League instead. Will be removed in a future release.
	League *LeagueClient
	// Deprecated: Use Client.LoL.Status instead. Will be removed in a future release.
	Status *StatusClient
	// Deprecated: Use Client.LoL.Match instead. Will be removed in a future release.
	Match *MatchClient
	// Deprecated: Use Client.LoL.Spectator instead. Will be removed in a future release.
	Spectator *SpectatorClient
	// Deprecated: Use Client.LoL.Summoner instead. Will be removed in a future release.
	Summoner *SummonerClient
	// Deprecated: Use Client.LoL.ThirdPartyCode instead. Will be removed in a future release.
	ThirdPartyCode *ThirdPartyCodeClient
	// Deprecated: Use Client.LoL.Tournament instead. Will be removed in a future release.
	Tournament *TournamentClient
}

// NewClient returns a new api client for the Riot API
func NewClient(region api.Region, apiKey string, client internal.Doer, logger log.FieldLogger) *Client {
	c := &Client{
		Region: region,
		apiKey: apiKey,
		client: client,
		l:      logger.WithField("client", "riot api"),
	}
	common := &struct {
		c *Client
	}{
		c: c,
	}
	c.ChampionMastery = (*championMasteryClient)(common)
	c.Summoner = (*summonerClient)(common)
	c.Champion = (*championClient)(common)
	c.League = (*leagueClient)(common)
	c.Status = (*statusClient)(common)
	c.Match = (*matchClient)(common)
	c.Spectator = (*spectatorClient)(common)
	c.Tournament = (*tournamentClient)(common)
	c.ThirdPartyCode = (*thirdPartyCodeClient)(common)
	return c
}

func (c *Client) getInto(endpoint string, target interface{}) error {
	logger := c.logger().WithFields(log.Fields{
		"method":   "getInto",
		"endpoint": endpoint,
	})
	response, err := c.get(endpoint)
	if err != nil {
		logger.Debug(err)
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		logger.Debug(err)
		return err
	}
	return nil
}

func (c *Client) postInto(endpoint string, body, target interface{}) error {
	logger := c.logger().WithFields(log.Fields{
		"method":   "postInto",
		"endpoint": endpoint,
	})
	response, err := c.post(endpoint, body)
	if err != nil {
		logger.Debug(err)
		return err
	}
	if err := json.NewDecoder(response.Body).Decode(target); err != nil {
		logger.Debug(err)
		return err
	}
	return nil
}

func (c *Client) put(endpoint string, body interface{}) error {
	logger := c.logger().WithFields(log.Fields{
		"method":   "put",
		"endpoint": endpoint,
	})
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		logger.Debug(err)
		return err
	}
	_, err := c.doRequest("PUT", endpoint, buf)
	return err
}

func (c *Client) get(endpoint string) (*http.Response, error) {
	return c.doRequest("GET", endpoint, nil)
}

func (c *Client) post(endpoint string, body interface{}) (*http.Response, error) {
	logger := c.logger().WithFields(log.Fields{
		"method":   "post",
		"endpoint": endpoint,
	})
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return c.doRequest("POST", endpoint, buf)
}

func (c *Client) doRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	logger := c.logger().WithFields(log.Fields{
		"method":   "doRequest",
		"endpoint": endpoint,
	})
	request, err := c.newRequest(method, endpoint, body)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	response, err := c.client.Do(request)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	if response.StatusCode == http.StatusServiceUnavailable {
		logger.Info("service unavailable, retrying")
		time.Sleep(time.Second)
		response, err = c.client.Do(request)
		if err != nil {
			logger.Debug(err)
			return nil, err
		}
	}
	if response.StatusCode == http.StatusTooManyRequests {
		retry := response.Header.Get("Retry-After")
		seconds, err := strconv.Atoi(retry)
		if err != nil {
			logger.Debug(err)
			return nil, err
		}
		logger.Infof("rate limited, waiting %d seconds", seconds)
		time.Sleep(time.Duration(seconds) * time.Second)
		return c.doRequest(method, endpoint, body)
	}
	if response.StatusCode < 200 || response.StatusCode > 299 {
		logger.Debugf("error response: %v", response.Status)
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

func (c *Client) newRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	logger := c.logger().WithFields(log.Fields{
		"method":   "newRequest",
		"endpoint": endpoint,
	})
	request, err := http.NewRequest(method, fmt.Sprintf(apiURLFormat, scheme, c.Region, baseURL, endpoint), body)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	request.Header.Add(apiTokenHeaderKey, c.apiKey)
	request.Header.Add("Accept", "application/json")
	return request, nil
}

func (c *Client) logger() log.FieldLogger {
	return c.l.WithField("region", c.Region)
}
