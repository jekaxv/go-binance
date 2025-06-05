package wfutures

import (
	"context"
	"github.com/jekaxv/go-binance/wss"
)

type Client struct {
	C *wss.Client
}

func (c *Client) setParams(key string, value any) {
	c.C.SetParams(key, value)
}

func (c *Client) wsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.C.WsServe(ctx)
}

func (c *Client) combined(combine bool) {
	c.C.Combined(combine)
}

func (c *Client) getEndpoint() string {
	return c.C.Opt.Endpoint
}

func (c *Client) setEndpoint(endpoint string) {
	c.C.Opt.Endpoint = endpoint
}

func (c *Client) close() error {
	return c.C.Close()
}
func (c *Client) send() error {
	return c.C.Send()
}

func (c *Client) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.C.WsApiServe(ctx)
}

func (c *Client) NewWebsocketStreams() *WebsocketStreams {
	return &WebsocketStreams{c: c}
}

// NewDepth Order Book
func (c *Client) NewDepth() *Depth {
	c.C.SetReq("depth")
	return &Depth{c: c}
}

// NewTickerPrice Symbol Price Ticker
func (c *Client) NewTickerPrice() *TickerPrice {
	c.C.SetReq("ticker.price")
	return &TickerPrice{c: c}
}

// NewTickerBook Symbol Order Book Ticker
func (c *Client) NewTickerBook() *TickerBook {
	c.C.SetReq("ticker.book")
	return &TickerBook{c: c}
}

// NewCreateOrder New Order(TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	c.C.SetReq("order.place", wss.AuthSigned)
	return &CreateOrder{c: c}
}

// NewModifyOrder Modify Order (TRADE)
func (c *Client) NewModifyOrder() *ModifyOrder {
	c.C.SetReq("order.modify", wss.AuthSigned)
	return &ModifyOrder{c: c}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.C.SetReq("order.cancel", wss.AuthSigned)
	return &CancelOrder{c: c}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.C.SetReq("order.status", wss.AuthSigned)
	return &QueryOrder{c: c}
}

// NewPositionInfo Position Information V2 (USER_DATA)
func (c *Client) NewPositionInfo() *PositionInfo {
	c.C.SetReq("v2/account.position", wss.AuthSigned)
	return &PositionInfo{c: c}
}

// NewAccountBalance Futures Account Balance V2(USER_DATA)
func (c *Client) NewAccountBalance() *AccountBalance {
	c.C.SetReq("v2/account.balance", wss.AuthSigned)
	return &AccountBalance{c: c}
}

// NewAccountInfo Account Information V2(USER_DATA)
func (c *Client) NewAccountInfo() *AccountInfo {
	c.C.SetReq("v2/account.status", wss.AuthSigned)
	return &AccountInfo{c: c}
}
