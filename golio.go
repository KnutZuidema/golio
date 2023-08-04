// Package golio is a wrapper for the Riot API and the Data Dragon service. It is written purely in Go and provides
// idiomatic access to all API endpoints.
//
// Example:
//
//	client := golio.NewClient("API KEY",
//	              golio.WithRegion(api.RegionNorthAmerica),
//	              golio.WithLogger(logrus.New().WithField("foo", "bar")))
//	summoner, _ := client.Riot.Summoner.GetByName("SK Jenax")
//	fmt.Printf("%s is a level %d summoner\n", summoner.Name, summoner.SummonerLevel)
//	champion, _ := client.DataDragon.GetChampion("Ashe")
//	mastery, err := client.Riot.ChampionMastery.Get(summoner.ID, champion.Key)
//	if err != nil {
//	fmt.Printf("%s has not played any games on %s\n", summoner.Name, champion.Name)
//	} else {
//	fmt.Printf("%s has mastery level %d with %d points on %s\n", summoner.Name, mastery.ChampionLevel,
//	mastery.ChampionPoints, champion.Name)
//	}
//	challengers, _ := client.Riot.League.GetChallenger(api.QueueRankedSolo)
//	rank1 := challengers.GetRank(0)
//	fmt.Printf("%s is the highest ranked player with %d league points\n", rank1.SummonerName, rank1.LeaguePoints)
package golio

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/datadragon"
	"github.com/KnutZuidema/golio/internal"
	"github.com/KnutZuidema/golio/riot"
	"github.com/KnutZuidema/golio/static"
)

// Client is a client for both the Riot API and the Data Dragon service
type Client struct {
	client     internal.Doer
	logger     log.FieldLogger
	region     api.Region
	apiKey     string
	Riot       *riot.Client
	DataDragon *datadragon.Client
	Static     *static.Client
}

// Option is used to alter the attributes of a client
type Option func(*Client)

// WithClient sets the given http client for the golio client
func WithClient(c internal.Doer) Option {
	return func(client *Client) {
		client.client = c
	}
}

// WithLogger sets the given logger for the golio client
func WithLogger(l log.FieldLogger) Option {
	return func(client *Client) {
		client.logger = l
	}
}

// WithRegion sets the given region for the golio client
func WithRegion(r api.Region) Option {
	return func(client *Client) {
		client.region = r
	}
}

// NewClient returns a new client for both the Riot API and the Data Dragon service
func NewClient(apiKey string, options ...Option) *Client {
	c := &Client{
		client: http.DefaultClient,
		logger: log.StandardLogger(),
		region: api.RegionEuropeWest,
		apiKey: apiKey,
	}
	for _, opt := range options {
		opt(c)
	}
	c.Riot = riot.NewClient(c.region, c.apiKey, c.client, c.logger)
	c.DataDragon = datadragon.NewClient(c.client, c.region, c.logger)
	c.Static = static.NewClient(c.client, c.logger)
	return c
}
