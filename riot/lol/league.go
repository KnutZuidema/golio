package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yigithanbalci/golio/internal"
)

// LeagueClient provides methods for league endpoints of the League of Legends API.
type LeagueClient struct {
	c *internal.Client
}

// GetChallenger returns the current Challenger league for the Region
func (l *LeagueClient) GetChallenger(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetChallenger")
	var list *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetChallengerLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// GetGrandmaster returns the current Grandmaster league for the Region
func (l *LeagueClient) GetGrandmaster(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetGrandmaster")
	var list *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetGrandmasterLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// GetMaster returns the current Master league for the Region
func (l *LeagueClient) GetMaster(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetMaster")
	var list *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetMasterLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// ListBySummoner returns all leagues a summoner with the given ID is in
func (l *LeagueClient) ListBySummoner(summonerID string) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListBySummoner")
	var leagues []*LeagueItem
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeaguesBySummoner, summonerID), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// ListPlayers returns all players with a league specified by its queue, tier and division
func (l *LeagueClient) ListPlayers(queue queue, tier tier, division division) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListPlayers")
	var leagues []*LeagueItem
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeagues, queue, tier, division), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// Get returns a ranked league with the specified ID
func (l *LeagueClient) Get(leagueID string) (*LeagueList, error) {
	logger := l.logger().WithField("method", "Get")
	var leagues *LeagueList
	if err := l.c.GetInto(fmt.Sprintf(endpointGetLeague, leagueID), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

func (l *LeagueClient) logger() log.FieldLogger {
	return l.c.Logger().WithField("category", "league")
}
