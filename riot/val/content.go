package val

import (
	"fmt"
	"github.com/KnutZuidema/golio/internal"
	log "github.com/sirupsen/logrus"
)

type ContentClient struct {
	c *internal.Client
}

func (cc *ContentClient) GetContent(locale Locale) (*ContentInfo, error) {
	logger := cc.logger().WithField("method", "GetContent")
	url := endPointGetContent
	if locale != "" {
		url = fmt.Sprintf(endPointGetContent, locale)
	}
	var contents *ContentInfo
	if err := cc.c.GetInto(url, &contents); err != nil {
		logger.Debug(err)
		fmt.Println(err)
		return nil, err
	}
	return contents, nil
}

func (cc *ContentClient) logger() log.FieldLogger {
	return cc.c.Logger().WithField("category", "content")
}
