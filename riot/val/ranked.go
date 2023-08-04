package val

import (
	"fmt"
	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

type RankedClient struct {
	c *internal.Client
}

func (cc *RankedClient) GetLeaderboardByActId(actId string, startIndex int32, size int32) (*Leaderboard, error) {
	logger := cc.logger().WithField("method", "GetLeaderboardByActId")
	var leaderboard *Leaderboard
	if startIndex < 0 {
		startIndex = 0
	}
	if size < 1 {
		size = 200
	}
	if err := cc.c.GetInto(fmt.Sprintf(endpointGetLeaderboardByActId+"?size=%d&startIndex=%d", actId, size, startIndex), &leaderboard); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return leaderboard, nil
}

func (cc *RankedClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "ranked")
}
