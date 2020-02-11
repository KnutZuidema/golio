package riot

import log "github.com/sirupsen/logrus"

type championClient struct {
	c *Client
}

// GetFreeRotation returns information about the current free champion rotation
func (c *championClient) GetFreeRotation() (*ChampionInfo, error) {
	logger := c.logger().WithField("method", "GetFreeRotation")
	var info *ChampionInfo
	if err := c.c.getInto(endpointGetFreeChampionRotation, &info); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return info, nil
}

func (c *championClient) logger() log.FieldLogger {
	return c.c.logger().WithField("category", "champion")
}
