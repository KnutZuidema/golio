package val

import (
	"fmt"

	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

// MatchClient provides methods for the match endpoints of the VALORANT API.
type MatchClient struct {
	c *internal.Client
}

// GetMatchByID returns information about a match using match id
func (cc *MatchClient) GetMatchByID(matchID string) (*Match, error) {
	logger := cc.logger().WithField("method", "GetMatchByID")
	url := endpointMatchByID
	var match *Match
	if err := cc.c.GetInto(fmt.Sprintf(url, matchID), &match); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return match, nil
}

// GetMatchListByPUUID returns match history as a list using player UUID
func (cc *MatchClient) GetMatchListByPUUID(puuid string) (*MatchList, error) {
	logger := cc.logger().WithField("method", "GetMatchListByPUUID")
	url := endpointMatchListByPUUID
	var matchList *MatchList
	if err := cc.c.GetInto(fmt.Sprintf(url, puuid), &matchList); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return matchList, nil
}

// GetRecentMatchesByQueue returns last match IDs for live regions and e-sports routing
func (cc *MatchClient) GetRecentMatchesByQueue(queue string) (*RecentMatches, error) {
	logger := cc.logger().WithField("method", "GetRecentMatchesByQueue")
	url := endpointRecentMatchesByQueue
	var recentMatches *RecentMatches
	if err := cc.c.GetInto(fmt.Sprintf(url, queue), &recentMatches); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return recentMatches, nil
}

func (cc *MatchClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "match")
}
