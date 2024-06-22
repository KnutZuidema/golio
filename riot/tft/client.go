package tft

import "github.com/KnutZuidema/golio/internal"

type Client struct {
	Spectator *SpectatorClient
}

func NewClient(base *internal.Client) *Client {
	return &Client{
		Spectator: &SpectatorClient{c: base},
	}
}
