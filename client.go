package golio

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	*RiotAPIClient
	*DataDragonClient
}

func NewClient(region region, apiKey string, client *http.Client, logger log.FieldLogger) *Client {
	return &Client{
		RiotAPIClient:    NewRiotAPIClient(region, apiKey, client, logger),
		DataDragonClient: NewDataDragonClient(client, region, logger),
	}
}
