package riot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type leagueClient struct {
	c *Client
}

// GetChallenger returns the current Challenger league for the Region
func (l *leagueClient) GetChallenger(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetChallenger")
	var list *LeagueList
	if err := l.c.getInto(fmt.Sprintf(endpointGetChallengerLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// GetGrandmaster returns the current Grandmaster league for the Region
func (l *leagueClient) GetGrandmaster(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetGrandmaster")
	var list *LeagueList
	if err := l.c.getInto(fmt.Sprintf(endpointGetGrandmasterLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// GetMaster returns the current Master league for the Region
func (l *leagueClient) GetMaster(queue queue) (*LeagueList, error) {
	logger := l.logger().WithField("method", "GetMaster")
	var list *LeagueList
	if err := l.c.getInto(fmt.Sprintf(endpointGetMasterLeague, queue), &list); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return list, nil
}

// ListBySummoner returns all leagues a summoner with the given ID is in
func (l *leagueClient) ListBySummoner(summonerID string) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListBySummoner")
	var leagues []*LeagueItem
	if err := l.c.getInto(fmt.Sprintf(endpointGetLeaguesBySummoner, summonerID), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// ListPlayers returns all players with a league specified by its queue, tier and division
func (l *leagueClient) ListPlayers(queue queue, tier tier, division division) ([]*LeagueItem, error) {
	logger := l.logger().WithField("method", "ListPlayers")
	var leagues []*LeagueItem
	if err := l.c.getInto(fmt.Sprintf(endpointGetLeagues, queue, tier, division), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

// Get returns a ranked league with the specified ID
func (l *leagueClient) Get(leagueID string) (*LeagueList, error) {
	logger := l.logger().WithField("method", "Get")
	var leagues *LeagueList
	if err := l.c.getInto(fmt.Sprintf(endpointGetLeague, leagueID), &leagues); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return leagues, nil
}

func (l *leagueClient) logger() log.FieldLogger {
	return l.c.logger().WithField("category", "league")
}
