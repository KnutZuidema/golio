package account

import "github.com/KnutZuidema/golio/internal"

type Client struct {
	AccountClient *AccountClient
}

func NewClient(base *internal.Client) *Client {
	return &Client{
		AccountClient: &AccountClient{base},
	}
}
