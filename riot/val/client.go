package val

import "github.com/yigithanbalci/golio/internal"

// Client pools all methods for endpoints of the Valorant API.
type Client struct {
	Content *ContentClient
	Status  *StatusClient
	Ranked  *RankedClient
	Match   *MatchClient
}

// NewClient returns a new instance of a League of Legends client.
func NewClient(base *internal.Client) *Client {
	return &Client{
		Content: &ContentClient{c: base},
		Status:  &StatusClient{c: base},
		Ranked:  &RankedClient{c: base},
		Match:   &MatchClient{c: base},
	}
}
