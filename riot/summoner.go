package riot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type summonerClient struct {
	c *Client
}

// GetByName returns the summoner with the given summoner name
func (s *summonerClient) GetByName(name string) (*Summoner, error) {
	return s.getBy(identificationName, name, s.logger().WithField("method", "GetByName"))
}

// GetByAccountID returns the summoner with the given account ID
func (s *summonerClient) GetByAccountID(id string) (*Summoner, error) {
	return s.getBy(identificationAccountID, id, s.logger().WithField("method", "GetByAccountID"))
}

// GetByPUUID returns the summoner with the given PUUID
func (s *summonerClient) GetByPUUID(puuid string) (*Summoner, error) {
	return s.getBy(identificationPUUID, puuid, s.logger().WithField("method", "GetByPUUID"))
}

// GetByID returns the summoner with the given ID
func (s *summonerClient) GetByID(summonerID string) (*Summoner, error) {
	return s.getBy(identificationSummonerID, summonerID, s.logger().WithField("method", "GetByID"))
}

func (s *summonerClient) getBy(by identification, value string, logger log.FieldLogger) (*Summoner, error) {
	var endpoint string
	switch by {
	case identificationSummonerID:
		endpoint = fmt.Sprintf(endpointGetSummonerBySummonerID, value)
	default:
		endpoint = fmt.Sprintf(endpointGetSummonerBy, by, value)
	}
	var summoner *Summoner
	if err := s.c.getInto(endpoint, &summoner); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return summoner, nil
}

func (s *summonerClient) logger() log.FieldLogger {
	return s.c.logger().WithField("category", "summoner")
}
