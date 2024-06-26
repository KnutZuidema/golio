package tft

import (
	"fmt"
	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

// SummonerClient provides methods for summoner endpoints of the League of Legends TFT API.
type SummonerClient struct {
	c *internal.Client
}

// GetSummonerByAccountID returns a summoner by account ID
func (sc *SummonerClient) GetSummonerByAccountID(encryptedAccountID string) (*Summoner, error) {
	logger := sc.logger().WithField("method", "GetSummonerByAccount")
	url := fmt.Sprintf(endpointSummonerByAccount, encryptedAccountID)
	var out *Summoner
	if err := sc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetSummonerByPUUID returns a summoner by PUUID
func (sc *SummonerClient) GetSummonerByPUUID(puuid string) (*Summoner, error) {
	logger := sc.logger().WithField("method", "GetSummonerByPUUID")
	url := fmt.Sprintf(endpointSummonerByPUUID, puuid)
	var out *Summoner
	if err := sc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetSummonerByMe returns a summoner by access token
func (sc *SummonerClient) GetSummonerByMe(authorization string) (*Summoner, error) {
	logger := sc.logger().WithField("method", "GetSummonerByMe")
	var out *Summoner
	if err := sc.c.GetInto(endpointSummonerByMe, &out, internal.WithHeader("Authorization", authorization)); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

// GetSummonerBySummonerID returns a summoner by summoner ID
func (sc *SummonerClient) GetSummonerBySummonerID(summonerID string) (*Summoner, error) {
	logger := sc.logger().WithField("method", "GetSummonerBySummonerID")
	url := fmt.Sprintf(endpointSummonerBySummonerID, summonerID)
	var out *Summoner
	if err := sc.c.GetInto(url, &out); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return out, nil
}

func (sc *SummonerClient) logger() log.FieldLogger {
	return sc.c.Logger().WithField("category", "summoner")
}
