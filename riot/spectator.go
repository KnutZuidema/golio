package riot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type spectatorClient struct {
	c *Client
}

// GetCurrent returns a currently running game for a summoner
func (s *spectatorClient) GetCurrent(summonerID string) (*GameInfo, error) {
	logger := s.logger().WithField("method", "GetCurrent")
	var games GameInfo
	if err := s.c.getInto(fmt.Sprintf(endpointGetCurrentGame, summonerID), &games); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &games, nil
}

// ListFeatured returns the currently featured games
func (s *spectatorClient) ListFeatured() (*FeaturedGames, error) {
	logger := s.logger().WithField("method", "ListFeatured")
	var games FeaturedGames
	if err := s.c.getInto(endpointGetFeaturedGames, &games); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return &games, nil
}

func (s *spectatorClient) logger() log.FieldLogger {
	return s.c.logger().WithField("category", "spectator")
}
