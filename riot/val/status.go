package val

import (
	log "github.com/sirupsen/logrus"
	"github.com/yigithanbalci/golio/internal"
)

type StatusClient struct {
	c *internal.Client
}

func (cc *StatusClient) GetPlatformData() (*PlatformDataDto, error) {
	logger := cc.logger().WithField("method", "GetPlatformData")
	var platformData *PlatformDataDto
	if err := cc.c.GetInto(endpointGetPlatformData, &platformData); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return platformData, nil
}

func (cc *StatusClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "status")
}
