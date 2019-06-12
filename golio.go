package golio

import (
	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
)

// Client is a client for both the Riot API and the Data Dragon service
type Client struct {
	*api.RiotAPIClient
	*api.DataDragonClient
}

// NewClient returns a new client for both the Riot API and the Data Dragon service
func NewClient(region api.Region, apiKey string, client api.Doer, logger log.FieldLogger) *Client {
	return &Client{
		RiotAPIClient:    api.NewRiotAPIClient(region, apiKey, client, logger),
		DataDragonClient: api.NewDataDragonClient(client, region, logger),
	}
}
