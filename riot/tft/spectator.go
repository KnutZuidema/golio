package tft

import (
	"fmt"

	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

// SpectatorClient provides methods for spectator endpoints of the League of Legends TFT API.
type SpectatorClient struct {
	c *internal.Client
}

// GetActiveGamesByPUUID returns current game information for the given puuid.
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

// GetFeaturedGames returns a list of featured games
func (sc *SpectatorClient) GetFeaturedGames() (*FeaturedGames, error) {
	logger := sc.logger().WithField("method", "GetFeaturedGames")
	var featuredGames FeaturedGames
	if err := sc.c.GetInto(endpointSpectatorFeaturedGames, &featuredGames); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &featuredGames, nil
}

func (sc *SpectatorClient) logger() log.FieldLogger {
	return sc.c.Logger().WithField("category", "spectator")
}
