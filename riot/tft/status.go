package tft

import (
	log "github.com/sirupsen/logrus"

	"github.com/KnutZuidema/golio/internal"
)

// StatusClient provides methods for the status endpoints of the League of Legends API.
type StatusClient struct {
	c *internal.Client
}

// Get returns the current status of the services for the Region
func (s *StatusClient) Get() (*Status, error) {
	logger := s.logger().WithField("method", "Get")
	var status *Status
	if err := s.c.GetInto(endpointGetStatus, &status); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return status, nil
}
func (s *StatusClient) logger() log.FieldLogger {
	return s.c.Logger().WithField("category", "status")
}
