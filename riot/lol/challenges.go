package lol

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yigithanbalci/golio/internal"
)

type ChallengesClient struct {
	c *internal.Client
}

func (cc *ChallengesClient) GetConfig() ([]*ChallengeConfigInfoDto, error) {
	logger := cc.logger().WithField("method", "GetConfig")
	var challengeConfigs []*ChallengeConfigInfoDto
	if err := cc.c.GetInto(endpointChallangesConfig, &challengeConfigs); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return challengeConfigs, nil
}

func (cc *ChallengesClient) GetPercentiles() (map[string]map[string]float64, error) {
	logger := cc.logger().WithField("method", "GetPercentiles")
	var percentiles map[string]map[string]float64
	if err := cc.c.GetInto(endpointChallangesPercentiles, &percentiles); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return percentiles, nil
}

func (cc *ChallengesClient) GetConfigWithChallengeId(challengeId int64) (*ChallengeConfigInfoDto, error) {
	logger := cc.logger().WithField("method", "getConfigWithChallengeId")
	var challengeConfig *ChallengeConfigInfoDto
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallangesConfigByChallengeId, challengeId), &challengeConfig); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return challengeConfig, nil
}

func (cc *ChallengesClient) GetLeaderBoardByChallengeIdAndLevel(challengeId int64, tier tier, limit int32) ([]*ApexPlayerInfoDto, error) {
	logger := cc.logger().WithField("method", "GetLeaderBoardByChallengeIdAndLevel")
	var apexPlayerInfo []*ApexPlayerInfoDto
	if tier == "" {
		tier = TierChallenger
	}
	if limit <= 0 {
		limit = 50
	}
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallangesLeaderboardsByChallengeIdAndLevel, challengeId, tier, limit), &apexPlayerInfo); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return apexPlayerInfo, nil
}

func (cc *ChallengesClient) GetPercentilesWithChallengeId(challengeId int64) (map[string]float64, error) {
	logger := cc.logger().WithField("method", "GetPercentilesWithChallengeId")
	var percentiles map[string]float64
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallangesPercentilesByChallengeId, challengeId), &percentiles); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return percentiles, nil
}

func (cc *ChallengesClient) GetPlayerDataWithPUUID(puuid string) (*PlayerInfoDto, error) {
	logger := cc.logger().WithField("method", "GetPlayerDataWithPUUID")
	var playerData *PlayerInfoDto
	if err := cc.c.GetInto(fmt.Sprintf(endpointChallangesPlayerdataByPuuid, puuid), &playerData); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return playerData, nil
}

func (cc *ChallengesClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "challenges")
}
