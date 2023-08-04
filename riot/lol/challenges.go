package lol

import (
	"fmt"
	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

type ChallengesClient struct {
	c *internal.Client
}

func (cc *ChallengesClient) GetConfig() ([]*ChallengeConfigInfo, error) {
	logger := cc.logger().WithField("method", "GetConfig")
	var challengeConfigs []*ChallengeConfigInfo
	if err := cc.c.GetInto(endpointChallengesConfig, &challengeConfigs); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return challengeConfigs, nil
}

type Percentiles map[string]float64

type PercentilesByChallenges map[string]Percentiles

func (cc *ChallengesClient) GetPercentiles() (PercentilesByChallenges, error) {
	logger := cc.logger().WithField("method", "GetPercentiles")
	var percentiles PercentilesByChallenges
	if err := cc.c.GetInto(endpointChallengesPercentiles, &percentiles); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return percentiles, nil
}

func (cc *ChallengesClient) GetConfigByChallengeID(challengeID int64) (*ChallengeConfigInfo, error) {
	logger := cc.logger().WithField("method", "GetConfigByChallengeID")
	var challengeConfig *ChallengeConfigInfo
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallengesConfigByChallengeID, challengeID), &challengeConfig); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return challengeConfig, nil
}

func (cc *ChallengesClient) GetLeaderBoardByChallengeIDAndLevel(challengeID int64, tier tier, limit int32) ([]*ApexPlayerInfo, error) {
	logger := cc.logger().WithField("method", "GetLeaderBoardByChallengeIDAndLevel")
	var apexPlayerInfo []*ApexPlayerInfo
	if tier == "" {
		tier = TierChallenger
	}
	if limit <= 0 {
		limit = 50
	}
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallengesLeaderboardsByChallengeIDAndLevel, challengeID, tier, limit), &apexPlayerInfo); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return apexPlayerInfo, nil
}

func (cc *ChallengesClient) GetPercentilesByChallengeID(challengeID int64) (Percentiles, error) {
	logger := cc.logger().WithField("method", "GetPercentilesByChallengeID")
	var percentiles Percentiles
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallengesPercentilesByChallengeID, challengeID), &percentiles); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return percentiles, nil
}

func (cc *ChallengesClient) GetPlayerDataByPUUID(uuid string) (*PlayerInfo, error) {
	logger := cc.logger().WithField("method", "GetPlayerDataByPUUID")
	var playerData *PlayerInfo
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallengesPlayerDataByPUUID, uuid), &playerData); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return playerData, nil
}

func (cc *ChallengesClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "challenges")
}
