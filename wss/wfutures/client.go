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

func (c *Client) close() error {
	return c.C.Close()
}
func (c *Client) send() error {
	return c.C.Send()
}

func (c *Client) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.C.WsApiServe(ctx)
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
	c.C.SetReq("order.place")
	return &CreateOrder{c: c}
}

// NewModifyOrder Modify Order (TRADE)
func (c *Client) NewModifyOrder() *ModifyOrder {
	c.C.SetReq("order.modify")
	return &ModifyOrder{c: c}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.C.SetReq("order.cancel")
	return &CancelOrder{c: c}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.C.SetReq("order.status")
	return &QueryOrder{c: c}
}

// NewPositionInfo Position Information V2 (USER_DATA)
func (c *Client) NewPositionInfo() *PositionInfo {
	c.C.SetReq("v2/account.position")
	return &PositionInfo{c: c}
}
