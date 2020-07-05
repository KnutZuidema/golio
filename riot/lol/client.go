package lol

import "github.com/KnutZuidema/golio/internal"

// Client pools all methods for endpoints of the League of Legends API.
type Client struct {
	ChampionMastery *ChampionMasteryClient
	Champion        *ChampionClient
	League          *LeagueClient
	Status          *StatusClient
	Match           *MatchClient
	Spectator       *SpectatorClient
	Summoner        *SummonerClient
	ThirdPartyCode  *ThirdPartyCodeClient
	Tournament      *TournamentClient
}

// NewClient returns a new instance of a League of Legends client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		ChampionMastery: &ChampionMasteryClient{Client: base},
		Summoner:        &SummonerClient{Client: base},
		Champion:        &ChampionClient{Client: base},
		League:          &LeagueClient{Client: base},
		Status:          &StatusClient{Client: base},
		Match:           &MatchClient{Client: base},
		Spectator:       &SpectatorClient{Client: base},
		Tournament:      &TournamentClient{Client: base},
		ThirdPartyCode:  &ThirdPartyCodeClient{Client: base},
	}
}
