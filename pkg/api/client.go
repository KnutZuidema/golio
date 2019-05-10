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
		RiotAPIClient: &RiotAPIClient{
			region: region,
			apiKey: apiKey,
			client: client,
		},
		DataDragonClient: &DataDragonClient{
			DataDragonVersion:  "9.9.1",
			DataDragonLanguage: LanguageCodeUnitedStates,
			client:             client,
		},
	}
}
