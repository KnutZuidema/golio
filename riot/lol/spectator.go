package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/yigithanbalci/golio/internal"
)

// SpectatorClient provides methods for the spectator endpoints of the League of Legends API.
type SpectatorClient struct {
	c *internal.Client
}

// GetCurrent returns a currently running game for a summoner
func (s *SpectatorClient) GetCurrent(summonerID string) (*GameInfo, error) {
	logger := s.logger().WithField("method", "GetCurrent")
	var games GameInfo
	if err := s.c.GetInto(fmt.Sprintf(endpointGetCurrentGame, summonerID), &games); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &games, nil
}

// ListFeatured returns the currently featured games
func (s *SpectatorClient) ListFeatured() (*FeaturedGames, error) {
	logger := s.logger().WithField("method", "ListFeatured")
	var games FeaturedGames
	if err := s.c.GetInto(endpointGetFeaturedGames, &games); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &games, nil
}

func (s *SpectatorClient) logger() log.FieldLogger {
	return s.c.Logger().WithField("category", "spectator")
}
