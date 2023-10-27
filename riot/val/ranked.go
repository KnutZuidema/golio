package val

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yigithanbalci/golio/internal"
)

// RankedClient provides methods for the ranked endpoints of the VALORANT API.
type RankedClient struct {
	c *internal.Client
}

// GetLeaderboardByActId returns leaderboard for the competitive queue by act ID
func (cc *RankedClient) GetLeaderboardByActId(actId string, startIndex int32, size int32) (*Leaderboard, error) {
	logger := cc.logger().WithField("method", "GetLeaderboardByActId")
	var leaderboard *Leaderboard
	if startIndex < 0 {
		startIndex = 0
	}
	if size < 1 {
		size = 200
	}
	if err := cc.c.GetInto(
		fmt.Sprintf(endpointGetLeaderboardByActID+"?size=%d&startIndex=%d", actId, size, startIndex), &leaderboard,
	); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return leaderboard, nil
}

func (cc *RankedClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "ranked")
}
