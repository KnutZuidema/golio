package tft

import (
	"fmt"

	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

/*
https://developer.riotgames.com/apis#tft-league-v1/GET_getChallengerLeague
*/

type LeagueClient struct {
	c *internal.Client
}

func (lc *LeagueClient) GetChallenger() (interface{}, error) {
	logger := lc.logger().WithField("method", "GetChallengerLeaguesByQueue")
	var out interface{}
	if err := lc.c.GetInto(endpointLeagueChallenger, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) GetEntriesBySummoner(summonerID string) (interface{}, error) {
	logger := lc.logger().WithField("method", "GetEntriesBySummoner")
	url := fmt.Sprintf(endpointLeagueEntriesBySummoner, summonerID)
	var out interface{}
	if err := lc.c.GetInto(url, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) GetEntries(tier, division string) (interface{}, error) {
	logger := lc.logger().WithField("method", "GetEntries")
	url := fmt.Sprintf(endpointLeagueEntries, tier, division)
	var out interface{}
	if err := lc.c.GetInto(url, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) GetGrandMaster() (interface{}, error) {
	logger := lc.logger().WithField("method", "GetGrandMaster")
	var out interface{}
	if err := lc.c.GetInto(endpointLeagueGrandMaster, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) GetLeagues(leagueID string) (interface{}, error) {
	logger := lc.logger().WithField("method", "GetLeagues")
	url := fmt.Sprintf(endpointLeagueLeagues, leagueID)
	var out interface{}
	if err := lc.c.GetInto(url, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) GetMaster() (interface{}, error) {
	logger := lc.logger().WithField("method", "GetMaster")
	var out interface{}
	if err := lc.c.GetInto(endpointLeagueMaster, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) GetRatedLaddersByQueue(queue string) (interface{}, error) {
	logger := lc.logger().WithField("method", "GetRatedLaddersByQueue")
	url := fmt.Sprintf(endpointLeagueRatedLattersByQueue, queue)
	var out interface{}
	if err := lc.c.GetInto(url, out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (l *LeagueClient) logger() log.FieldLogger {
	return l.c.Logger().WithField("category", "league")
}
