package riot

import log "github.com/sirupsen/logrus"

type statusClient struct {
	c *Client
}

// Get returns the current status of the services for the Region
func (s *statusClient) Get() (*Status, error) {
	logger := s.logger().WithField("method", "Get")
	var status *Status
	if err := s.c.getInto(endpointGetStatus, &status); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return status, nil
}
func (s *statusClient) logger() log.FieldLogger {
	return s.c.logger().WithField("category", "status")
}
