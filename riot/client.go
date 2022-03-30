// Package riot provides methods for accessing the Riot API for League of Legends.
// This includes dynamic data like the current game a summoner is in or their ranked standing.
package riot

import (
	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/KnutZuidema/golio/riot/lor"
	"github.com/KnutZuidema/golio/riot/tft"
)

// Client provides access to all Riot API endpoints
type Client struct {
	// Deprecated: Use Client.LoL.ChampionMastery instead. Will be removed in a future release.
	ChampionMastery *lol.ChampionMasteryClient
	// Deprecated: Use Client.LoL.Champion instead. Will be removed in a future release.
	Champion *lol.ChampionClient
	// Deprecated: Use Client.LoL.League instead. Will be removed in a future release.
	League *lol.LeagueClient
	// Deprecated: Use Client.LoL.Status instead. Will be removed in a future release.
	Status *lol.StatusClient
	// Deprecated: Use Client.LoL.Match instead. Will be removed in a future release.
	Match *lol.MatchClient
	// Deprecated: Use Client.LoL.Spectator instead. Will be removed in a future release.
	Spectator *lol.SpectatorClient
	// Deprecated: Use Client.LoL.Summoner instead. Will be removed in a future release.
	Summoner *lol.SummonerClient
	// Deprecated: Use Client.LoL.ThirdPartyCode instead. Will be removed in a future release.
	ThirdPartyCode *lol.ThirdPartyCodeClient
	// Deprecated: Use Client.LoL.Tournament instead. Will be removed in a future release.
	Tournament *lol.TournamentClient

	LoL *lol.Client
	LoR *lor.Client
	TfT *tft.Client
}

// NewClient returns a new api client for the Riot API
func NewClient(region api.Region, apiKey string, client internal.Doer, logger log.FieldLogger) *Client {
	baseClient := internal.NewClient(region, apiKey, client, logger)
	c := &Client{
		LoL: lol.NewClient(baseClient),
		LoR: lor.NewClient(baseClient),
		TfT: tft.NewClient(baseClient),
	}

	// TODO: deprecated, remove in a future release
	c.ChampionMastery = c.LoL.ChampionMastery
	c.Summoner = c.LoL.Summoner
	c.Champion = c.LoL.Champion
	c.League = c.LoL.League
	c.Status = c.LoL.Status
	c.Match = c.LoL.Match
	c.Spectator = c.LoL.Spectator
	c.Tournament = c.LoL.Tournament
	c.ThirdPartyCode = c.LoL.ThirdPartyCode
	return c
}
