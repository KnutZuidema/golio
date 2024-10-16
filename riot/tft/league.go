package tft

import (
	"fmt"

	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

// LeagueClient provides methods for league endpoints of the League of Legends TFT API.
type LeagueClient struct {
	c *internal.Client
}

// GetChallenger returns the current Challenger league for the Region
func (lc *LeagueClient) GetChallenger(queue queue) (*LeagueList, error) {
	logger := lc.logger().WithField("method", "GetChallenger")
	if queue == "" {
		queue = QueueRankedTFT
	}
	url := fmt.Sprintf(endpointLeagueChallenger, queue)
	var out *LeagueList
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetEntriesBySummoner returns league entries for a given summoner ID
func (lc *LeagueClient) GetEntriesBySummoner(summonerID string) ([]*LeagueEntry, error) {
	logger := lc.logger().WithField("method", "GetEntriesBySummoner")
	url := fmt.Sprintf(endpointLeagueEntriesBySummoner, summonerID)
	var out []*LeagueEntry
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetEntries returns all the league entries
func (lc *LeagueClient) GetEntries(tier tier, division division) ([]*LeagueEntry, error) {
	logger := lc.logger().WithField("method", "GetEntries")
	url := fmt.Sprintf(endpointLeagueEntries, tier, division)
	var out []*LeagueEntry
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetGrandMaster returns the current GrandMaster league for the Region
func (lc *LeagueClient) GetGrandMaster(queue queue) (*LeagueList, error) {
	logger := lc.logger().WithField("method", "GetGrandMaster")
	if queue == "" {
		queue = QueueRankedTFT
	}
	url := fmt.Sprintf(endpointLeagueGrandMaster, queue)
	var out *LeagueList
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetLeagues returns league with given ID, including inactive entries
func (lc *LeagueClient) GetLeagues(leagueID string) (*LeagueList, error) {
	logger := lc.logger().WithField("method", "GetLeagues")
	url := fmt.Sprintf(endpointLeagueLeagues, leagueID)
	var out *LeagueList
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetMaster returns the current Master league for the Region
func (lc *LeagueClient) GetMaster(queue queue) (*LeagueList, error) {
	logger := lc.logger().WithField("method", "GetMaster")
	if queue == "" {
		queue = QueueRankedTFT
	}
	url := fmt.Sprintf(endpointLeagueMaster, queue)
	var out *LeagueList
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetRatedLaddersByQueue returns the top rated ladder for given queue
func (lc *LeagueClient) GetRatedLaddersByQueue(queue queue) ([]*TopRatedLadderEntry, error) {
	logger := lc.logger().WithField("method", "GetRatedLaddersByQueue")
	url := fmt.Sprintf(endpointLeagueRatedLattersByQueue, queue)
	var out []*TopRatedLadderEntry
	if err := lc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (lc *LeagueClient) logger() log.FieldLogger {
	return lc.c.Logger().WithField("category", "league")
}
