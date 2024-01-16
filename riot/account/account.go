package account

import (
	"fmt"
	"github.com/KnutZuidema/golio/api"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// AccountClient provides methods for the account endpoints of the League of Legends API.
type AccountClient struct { //nolint:golint
	c *internal.Client
}

// GetByPUUID returns the account matching the PUUID
func (ac *AccountClient) GetByPUUID(puuid string) (*Account, error) {
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
func (ac *AccountClient) GetByRiotID(gameName, tagLine string) (*Account, error) {
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

func (ac *AccountClient) logger() log.FieldLogger {
	return ac.c.Logger().WithField("category", "account")
}
