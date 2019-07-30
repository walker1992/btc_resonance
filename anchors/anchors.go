package anchors

import (
	"net/rpc"
)

type NewResonanceTask struct {
	Address   string
	Amount    string
	Signature string
}
type SigResonance struct {
	Address   string
	Amount    string
	Signature string
}

type Client struct {
	client *rpc.Client
	path   string
}

func NewClient(url string) (*Client, error) {
	client, err := rpc.DialHTTP("tcp", url)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		path:   url}, nil
}

func (c *Client) ResonanceTrade(body *ResonanceMsg) error {
	return c.client.Call("TradeServer.ResonanceTrade", body, nil)
}
