package riot

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type thirdPartyCodeClient struct {
	c *Client
}

// Get returns the third party code for the given summoner id
func (t *thirdPartyCodeClient) Get(summonerID string) (string, error) {
	logger := t.logger().WithFields(log.Fields{
		"method": "Get",
	})
	var code string
	if err := t.c.getInto(fmt.Sprintf(endpointGetThirdPartyCode, summonerID), &code); err != nil {
		logger.Debug(err)
		return "", err
	}
	return code, nil
}

func (t *thirdPartyCodeClient) logger() log.FieldLogger {
	return t.c.logger().WithField("category", "third party code")
}
