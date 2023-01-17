package lor

import "github.com/yigithanbalci/golio/internal"

// Client pools methods for the Legends of Runeterra API.
type Client struct {
	Ranked *RankedClient
}

// NewClient returns a new instance of a Legends of Runeterra client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		Ranked: &RankedClient{c: base},
	}
}
