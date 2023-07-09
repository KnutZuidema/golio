package lor

import "github.com/yigithanbalci/golio/internal"

// RankedClient provides methods for the ranked endpoints of the Legends of Runeterra API.
type RankedClient struct {
	c *internal.Client
}

// GetMasters returns all players currently in the Master tier for the region.
func (c *RankedClient) GetMasters() ([]*Player, error) {
	var players []*Player
	if err := c.c.GetInto(endpointGetMaster, &players); err != nil {
		return nil, err
	}
	return players, nil
}
