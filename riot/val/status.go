package val

import (
	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// StatusClient provides methods for the status endpoints of the VALORANT API.
type StatusClient struct {
	c *internal.Client
}

// GetPlatformData returns information about platform including maintenances and incidents
func (cc *StatusClient) GetPlatformData() (*PlatformData, error) {
	logger := cc.logger().WithField("method", "GetPlatformData")
	var platformData *PlatformData
	if err := cc.c.GetInto(endpointGetPlatformData, &platformData); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return platformData, nil
}

func (cc *StatusClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "status")
}
