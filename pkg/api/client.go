package api

import (
	"net/http"
)

type Client struct {
	*RiotAPIClient
	*DataDragonClient
}

func NewClient(region region, apiKey string, client *http.Client) *Client {
	return &Client{
		RiotAPIClient:    NewRiotAPIClient(region, apiKey, client),
		DataDragonClient: NewDataDragonClient(client, region),
	}
}
