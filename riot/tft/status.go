package tft

import (
	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

// StatusClient provides methods for status endpoints of the League of Legends TFT API.
type StatusClient struct {
	c *internal.Client
}

// GetPlatformData returns Teamfight Tactics status for the given platform
func (sc *StatusClient) GetPlatformData() (*PlatformData, error) {
	logger := sc.logger().WithField("method", "GetPlatformData")
	var out *PlatformData
	if err := sc.c.GetInto(endpointStatusPlatformData, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (sc *StatusClient) logger() log.FieldLogger {
	return sc.c.Logger().WithField("category", "status")
}
