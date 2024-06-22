package tft

import (
	"fmt"

	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

type SpectatorClient struct {
	c *internal.Client
}

func (sc *SpectatorClient) GetActiveGamesByPUUID(puuid string) (*CurrentGameInfo, error) {
	logger := sc.logger().WithField("method", "GetActiveGamesByPUUID")
	url := fmt.Sprintf(endpointSpectatorActiveGamedByPUUID, puuid)
	var currentGameInfo CurrentGameInfo
	if err := sc.c.GetInto(url, &currentGameInfo); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &currentGameInfo, nil
}

func (sc *SpectatorClient) GetFeaturedGames() (*FeaturedGames, error) {
	logger := sc.logger().WithField("method", "GetFeaturedGames")
	var featuredGames FeaturedGames
	if err := sc.c.GetInto(endpointSpectatorFeaturedGames, &featuredGames); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &featuredGames, nil
}

func (l *SpectatorClient) logger() log.FieldLogger {
	return l.c.Logger().WithField("category", "spectator")
}
