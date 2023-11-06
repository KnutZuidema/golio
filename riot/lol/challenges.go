package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// ChallengesClient provides methods for the challenges endpoints of the League of Legends API.
type ChallengesClient struct {
	c *internal.Client
}

// GetConfig returns all basic challenge configuration information
func (cc *ChallengesClient) GetConfig() ([]*ChallengeConfigInfo, error) {
	logger := cc.logger().WithField("method", "GetConfig")
	var challengeConfigs []*ChallengeConfigInfo
	if err := cc.c.GetInto(endpointChallengesConfig, &challengeConfigs); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return challengeConfigs, nil
}

// Percentiles a map of level to percentile used by PercentilesByChallenges as value
type Percentiles map[string]float64

// PercentilesByChallenges a map of level to percentile of players who have achieved it
type PercentilesByChallenges map[string]Percentiles

// GetPercentiles returns a map of level to percentile of players who have achieved it
func (cc *ChallengesClient) GetPercentiles() (PercentilesByChallenges, error) {
	logger := cc.logger().WithField("method", "GetPercentiles")
	var percentiles PercentilesByChallenges
	if err := cc.c.GetInto(endpointChallengesPercentiles, &percentiles); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return percentiles, nil
}

// GetConfigByChallengeID returns challenge configuration by ID
func (cc *ChallengesClient) GetConfigByChallengeID(challengeID int64) (*ChallengeConfigInfo, error) {
	logger := cc.logger().WithField("method", "GetConfigByChallengeID")
	var challengeConfig *ChallengeConfigInfo
	if err := cc.c.GetInto(
		fmt.Sprintf(endpointChallengesConfigByChallengeID, challengeID), &challengeConfig,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return challengeConfig, nil
}

// GetLeaderBoardByChallengeIDAndLevel returns top players for each level
func (cc *ChallengesClient) GetLeaderBoardByChallengeIDAndLevel(
	challengeID int64, tier tier, limit int32,
) ([]*ApexPlayerInfo, error) {
	logger := cc.logger().WithField("method", "GetLeaderBoardByChallengeIDAndLevel")
	var apexPlayerInfo []*ApexPlayerInfo
	if tier == "" {
		tier = TierChallenger
	}
	if limit <= 0 {
		limit = 50
	}
	if err := cc.c.GetInto(
		fmt.Sprintf(endpointChallengesLeaderboards, challengeID, tier, limit), &apexPlayerInfo,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return apexPlayerInfo, nil
}

// GetPercentilesByChallengeID returns map of level to percentiles of players who have achieved it for a challenge
func (cc *ChallengesClient) GetPercentilesByChallengeID(challengeID int64) (Percentiles, error) {
	logger := cc.logger().WithField("method", "GetPercentilesByChallengeID")
	var percentiles Percentiles
	if err := cc.c.GetInto(
		fmt.Sprintf(endpointChallengesPercentilesByChallengeID, challengeID), &percentiles,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return percentiles, nil
}

// GetPlayerDataByPUUID returns player information with list of all progressed challenges
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
