package internal

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
)

const (
	apiURLFormat      = "%s://%s.%s%s"
	baseURL           = "api.riotgames.com"
	scheme            = "https"
	apiTokenHeaderKey = "X-Riot-Token"
)

// Client provides methods for communication with the Riot API.
type Client struct {
	L      log.FieldLogger
	Region api.Region
	APIKey string
	Client Doer
}

// NewClient returns a new client.
func NewClient(region api.Region, key string, client Doer, logger log.FieldLogger) *Client {
	return &Client{
		L:      logger,
		Region: region,
		APIKey: key,
		Client: client,
	}
}

// GetInto processes a GET request and saves the response body into the given target.
func (c *Client) GetInto(endpoint string, target interface{}) error {
	logger := c.Logger().WithFields(
		log.Fields{
			"method":   "GetInto",
			"endpoint": endpoint,
		},
	)
	response, err := c.Get(endpoint)
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

// PostInto processes a POST request and saves the response body into the given target.
func (c *Client) PostInto(endpoint string, body, target interface{}) error {
	logger := c.Logger().WithFields(
		log.Fields{
			"method":   "PostInto",
			"endpoint": endpoint,
		},
	)
	response, err := c.Post(endpoint, body)
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

// Put processes a PUT request.
func (c *Client) Put(endpoint string, body interface{}) error {
	logger := c.Logger().WithFields(
		log.Fields{
			"method":   "Put",
			"endpoint": endpoint,
		},
	)
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		logger.Debug(err)
		return err
	}
	_, err := c.DoRequest("PUT", endpoint, buf)
	return err
}

// Get processes a GET request.
func (c *Client) Get(endpoint string) (*http.Response, error) {
	return c.DoRequest("GET", endpoint, nil)
}

// Post processes a POST request.
func (c *Client) Post(endpoint string, body interface{}) (*http.Response, error) {
	logger := c.Logger().WithFields(
		log.Fields{
			"method":   "Post",
			"endpoint": endpoint,
		},
	)
	buf := &bytes.Buffer{}
	if err := json.NewEncoder(buf).Encode(body); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return c.DoRequest("POST", endpoint, buf)
}

// DoRequest processes a http.Request and returns the response.
// Rate-Limiting and retrying is handled via the corresponding response headers.
func (c *Client) DoRequest(method, endpoint string, body io.Reader) (*http.Response, error) {
	logger := c.Logger().WithFields(
		log.Fields{
			"method":   "DoRequest",
			"endpoint": endpoint,
		},
	)
	request, err := c.NewRequest(method, endpoint, body)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	response, err := c.Client.Do(request)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	if response.StatusCode == http.StatusServiceUnavailable {
		logger.Info("service unavailable, retrying")
		time.Sleep(time.Second)
		response, err = c.Client.Do(request)
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
		return c.DoRequest(method, endpoint, body)
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

// NewRequest returns a new http.Request with necessary headers et.
func (c *Client) NewRequest(method, endpoint string, body io.Reader) (*http.Request, error) {
	logger := c.Logger().WithFields(
		log.Fields{
			"method":   "NewRequest",
			"endpoint": endpoint,
		},
	)
	request, err := http.NewRequest(method, fmt.Sprintf(apiURLFormat, scheme, c.Region, baseURL, endpoint), body)
	if err != nil {
		logger.Debug(err)
		return nil, err
	}
	request.Header.Add(apiTokenHeaderKey, c.APIKey)
	request.Header.Add("Accept", "application/json")
	return request, nil
}

// Logger returns a logger with client specific fields set.
func (c *Client) Logger() log.FieldLogger {
	return c.L.WithField("region", c.Region)
}
