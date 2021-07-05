package tft

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// SummonerClient provides methods for the summoner endpoints of the League of Legends API.
type SummonerClient struct {
	c *internal.Client
}

// GetByName returns the summoner with the given summoner name
func (s *SummonerClient) GetByName(name string) (*Summoner, error) {
	return s.getBy(identificationName, name, s.logger().WithField("method", "GetByName"))
}

// GetByAccountID returns the summoner with the given account ID
func (s *SummonerClient) GetByAccountID(id string) (*Summoner, error) {
	return s.getBy(identificationAccountID, id, s.logger().WithField("method", "GetByAccountID"))
}

// GetByPUUID returns the summoner with the given PUUID
func (s *SummonerClient) GetByPUUID(puuid string) (*Summoner, error) {
	return s.getBy(identificationPUUID, puuid, s.logger().WithField("method", "GetByPUUID"))
}

// GetByID returns the summoner with the given ID
func (s *SummonerClient) GetByID(summonerID string) (*Summoner, error) {
	return s.getBy(identificationSummonerID, summonerID, s.logger().WithField("method", "GetByID"))
}

func (s *SummonerClient) getBy(by identification, value string, logger log.FieldLogger) (*Summoner, error) {
	var endpoint string
	switch by {
	case identificationSummonerID:
		endpoint = fmt.Sprintf(endpointGetSummonerBySummonerID, value)
	default:
		endpoint = fmt.Sprintf(endpointGetSummonerBy, by, value)
	}
	var summoner *Summoner
	if err := s.c.GetInto(endpoint, &summoner); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return summoner, nil
}

func (s *SummonerClient) logger() log.FieldLogger {
	return s.c.Logger().WithField("category", "summoner")
}
