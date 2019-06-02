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
func NewClient(region region, apiKey string, client *http.Client, logger log.FieldLogger) *Client {
	return &Client{
		RiotAPIClient:    NewRiotAPIClient(region, apiKey, client, logger),
		DataDragonClient: NewDataDragonClient(client, region, logger),
	}
}
