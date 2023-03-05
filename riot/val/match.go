package val

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/yigithanbalci/golio/internal"
)

type MatchClient struct {
	c *internal.Client
}

func (cc *MatchClient) GetContent(locale LocalizedNamesDto) (*ContentInfoDto, error) {
	logger := cc.logger().WithField("method", "GetContent")
	url := endPointGetContent
	if locale != "" {
		url = fmt.Sprintf(endPointGetContent, locale)
	}
	var contents *ContentInfoDto
	if err := cc.c.GetInto(url, &contents); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return contents, nil
}

func (cc *MatchClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "match")
}
