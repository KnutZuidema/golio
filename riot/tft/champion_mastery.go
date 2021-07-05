package tft

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// ChampionMasteryClient provides methods for the champion mastery endpoints of the
// League of Legends API.
type ChampionMasteryClient struct {
	c *internal.Client
}

// List returns information about masteries for the summoner with the given ID
func (c *ChampionMasteryClient) List(summonerID string) ([]*ChampionMastery, error) {
	logger := c.logger().WithField("method", "List")
	var masteries []*ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMasteries, summonerID),
		&masteries,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return masteries, nil
}

// Get returns information about the mastery of the champion with the given ID the summoner with the
// given ID has
func (c *ChampionMasteryClient) Get(summonerID, championID string) (*ChampionMastery, error) {
	logger := c.logger().WithField("method", "Get")
	var mastery *ChampionMastery
	if err := c.c.GetInto(
		fmt.Sprintf(endpointGetChampionMastery, summonerID, championID),
		&mastery,
	); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return mastery, nil
}

// GetTotal returns the accumulated mastery score of all champions played by the summoner with the
// given ID
func (c *ChampionMasteryClient) GetTotal(summonerID string) (int, error) {
	logger := c.logger().WithField("method", "GetTotal")
	var score int
	if err := c.c.GetInto(fmt.Sprintf(endpointGetChampionMasteryTotalScore, summonerID), &score); err != nil {
		logger.Debug(err)
		return 0, err
	}
	return score, nil
}

func (c *ChampionMasteryClient) logger() log.FieldLogger {
	return c.c.Logger().WithField("category", "champion mastery")
}
