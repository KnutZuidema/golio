// Package tft allows you to interact with the Teamfight Tactics API
package tft

import "github.com/KnutZuidema/golio/internal"

// Client pools all methods for endpoints of the League of Legends TFT API.
type Client struct {
	Spectator *SpectatorClient
	League    *LeagueClient
	Match     *MatchClient
	Status    *StatusClient
	Summoner  *SummonerClient
}

// NewClient returns a new instance of a League of Legends TFT client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		Spectator: &SpectatorClient{c: base},
		League:    &LeagueClient{c: base},
		Match:     &MatchClient{c: base},
		Status:    &StatusClient{c: base},
		Summoner:  &SummonerClient{c: base},
	}
}
