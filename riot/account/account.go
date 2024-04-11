package account

import (
	"fmt"

	"github.com/KnutZuidema/golio/api"

	log "github.com/sirupsen/logrus"
)

// GetByPUUID returns the account matching the PUUID
func (ac *Client) GetByPUUID(puuid string) (*Account, error) {
	logger := ac.logger().WithField("method", "GetByPUUID")
	var account Account
	c := *ac.c
	c.Region = api.Region(api.RegionToRoute[c.Region])

	if err := c.GetInto(
		fmt.Sprintf(endpointGetByPUUID, puuid),
		&account,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &account, nil
}

// GetByRiotID returns the account matching the riot id
func (ac *Client) GetByRiotID(gameName, tagLine string) (*Account, error) {
	logger := ac.logger().WithField("method", "GetByRiotID")
	var account Account
	c := *ac.c
	c.Region = api.Region(api.RegionToRoute[c.Region])

	if err := c.GetInto(
		fmt.Sprintf(endpointGetByRiotID, gameName, tagLine),
		&account,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &account, nil
}

func (ac *Client) logger() log.FieldLogger {
	return ac.c.Logger().WithField("category", "account")
}
