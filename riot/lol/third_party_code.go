package lol

import (
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// ThirdPartyCodeClient provides methods for the third party code endpoints of the
// League of Legends API.
type ThirdPartyCodeClient struct {
	c *internal.Client
}

// Get returns the third party code for the given summoner id
func (t *ThirdPartyCodeClient) Get(summonerID string) (string, error) {
	logger := t.logger().WithFields(
		log.Fields{
			"method": "Get",
		},
	)
	var code string
	if err := t.c.GetInto(fmt.Sprintf(endpointGetThirdPartyCode, summonerID), &code); err != nil {
		logger.Debug(err)
		return "", err
	}
	return code, nil
}

func (t *ThirdPartyCodeClient) logger() log.FieldLogger {
	return t.c.Logger().WithField("category", "third party code")
}
