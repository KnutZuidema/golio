[![Documentation](https://godoc.org/github.com/KnutZuidema/golio?status.svg)](https://godoc.org/github.com/KnutZuidema/golio)
[![Build Status](https://travis-ci.com/KnutZuidema/golio.svg?branch=master)](https://travis-ci.com/KnutZuidema/golio)
# Golio

Golio is a wrapper for the Riot API and the Data Dragon service.
It is written purely in Go and provides idiomatic access to all
API endpoints.

## Example

```go
package main

import (
	"fmt"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio"
)

func main() {
	client := golio.NewClient(golio.RegionEuropeWest, "API KEY",
		http.DefaultClient, log.StandardLogger())
	summoner, _ := client.GetSummonerByName("SK Jenax")
	fmt.Printf("%s is a level %d summoner\n", summoner.Name, summoner.SummonerLevel)
	champion, _ := client.GetChampion("Ashe")
	mastery, err := client.GetChampionMastery(summoner.ID, champion.Key)
	if err != nil {
		fmt.Printf("%s has not played any games on %s\n", summoner.Name, champion.Name)
	} else {
		fmt.Printf("%s has mastery level %d with %d points on %s\n", summoner.Name, mastery.ChampionLevel,
			mastery.ChampionPoints, champion.Name)
	}
	challengers, _ := client.GetChallengerLeague(api.QueueRankedSolo)
	rank1 := challengers.GetRank(0)
	fmt.Printf("%s is the highest ranked player with %d league points\n", rank1.SummonerName, rank1.LeaguePoints)
}
```
