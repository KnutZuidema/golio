package account

import (
	"fmt"

	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

type AccountClient struct {
	c *internal.Client
}

func (ac *AccountClient) GetByPuuid(puuid string) (*Account, error) {
	logger := ac.logger().WithField("method", "GetByPuuid")
	var account Account

	if err := ac.c.GetInto(
		fmt.Sprintf(endpointGetByPuuid, puuid),
		&account,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &account, nil
}

func (ac *AccountClient) logger() log.FieldLogger {
	return ac.c.Logger().WithField("category", "account")
}
