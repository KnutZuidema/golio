package lol

import (
	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// ChampionClient provides methods for the champions endpoints of the League of Legends API.
type ChampionClient struct {
	Client *internal.Client
}

// GetFreeRotation returns information about the current free champion rotation
func (c *ChampionClient) GetFreeRotation() (*ChampionInfo, error) {
	logger := c.logger().WithField("method", "GetFreeRotation")
	var info *ChampionInfo
	if err := c.Client.GetInto(endpointGetFreeChampionRotation, &info); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return info, nil
}

func (c *ChampionClient) logger() log.FieldLogger {
	return c.Client.Logger().WithField("category", "champion")
}
