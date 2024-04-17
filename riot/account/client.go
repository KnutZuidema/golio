package account

import "github.com/KnutZuidema/golio/internal"

// Client pools all methods for endpoints of the League of Legends API.
type Client struct {
	c *internal.Client
}

// NewClient returns a new instance of a League of Legends client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		base,
	}
}
