package golio

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Client is a client for both the Riot API and the Data Dragon service
type Client struct {
	*RiotAPIClient
	*DataDragonClient
}

// NewClient returns a new client for both the Riot API and the Data Dragon service
func NewClient(region region, apiKey string, client Doer, logger log.FieldLogger) *Client {
	return &Client{
		RiotAPIClient:    NewRiotAPIClient(region, apiKey, client, logger),
		DataDragonClient: NewDataDragonClient(client, region, logger),
	}
}

// Doer is an interface for any client that can process an HTTP request and return a response.
// This will most commonly be a simple HTTP client.
type Doer interface {
	// Do processes an HTTP request and returns the response
	Do(r *http.Request) (*http.Response, error)
}
