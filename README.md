[![Documentation](https://godoc.org/github.com/KnutZuidema/golio?status.svg)](https://godoc.org/github.com/KnutZuidema/golio)
[![CI](https://github.com/KnutZuidema/golio/actions/workflows/ci.yml/badge.svg)](https://github.com/KnutZuidema/golio/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/KnutZuidema/golio)](https://goreportcard.com/report/github.com/KnutZuidema/golio)
[![codecov](https://codecov.io/gh/KnutZuidema/golio/branch/master/graph/badge.svg)](https://codecov.io/gh/KnutZuidema/golio)

# Golio

Golio is a wrapper for the Riot API and the Data Dragon service.
It is written purely in Go and provides idiomatic access to all
API endpoints.

## Example

```go
package main

import (
	"fmt"

	"github.com/KnutZuidema/golio"
	"github.com/KnutZuidema/golio/api"
	"github.com/KnutZuidema/golio/riot/lol"
	"github.com/sirupsen/logrus"
)

func main() {
	client := golio.NewClient("API KEY",
		golio.WithRegion(api.RegionEuropeWest),
		golio.WithLogger(logrus.New().WithField("foo", "bar")))
	summoner, _ := client.Riot.LoL.Summoner.GetByName("SK Jenax")
	fmt.Printf("%s is a level %d summoner\n", summoner.Name, summoner.SummonerLevel)
	champion, _ := client.DataDragon.GetChampion("Ashe")
	mastery, err := client.Riot.LoL.ChampionMastery.Get(summoner.ID, champion.Key)
	if err != nil {
		fmt.Printf("%s has not played any games on %s\n", summoner.Name, champion.Name)
	} else {
		fmt.Printf("%s has mastery level %d with %d points on %s\n", summoner.Name, mastery.ChampionLevel,
			mastery.ChampionPoints, champion.Name)
	}
	challengers, _ := client.Riot.LoL.League.GetChallenger(lol.QueueRankedSolo)
	rank1 := challengers.GetRank(0)
	fmt.Printf("%s is the highest ranked player with %d league points\n", rank1.SummonerName, rank1.LeaguePoints)
}
```
