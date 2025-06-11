package futures

import (
	"context"
	"github.com/jekaxv/go-binance/core"
)

type WsClient struct {
	C *core.WsClient
}
type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
type ApiRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type ApiResponse struct {
	Id         string          `json:"id"`
	Status     int             `json:"status"`
	RateLimits []*ApiRateLimit `json:"rateLimits,omitempty"`
	Error      *ApiError       `json:"error,omitempty"`
}

func (c *WsClient) setParams(key string, value any) {
	c.C.SetParams(key, value)
}

func (c *WsClient) wsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.C.WsServe(ctx)
}

func (c *WsClient) combined(combine bool) {
	c.C.Combined(combine)
}

func (c *WsClient) getEndpoint() string {
	return c.C.Opt.Endpoint
}

func (c *WsClient) setEndpoint(endpoint string) {
	c.C.Opt.Endpoint = endpoint
}

func (c *WsClient) close() error {
	return c.C.Close()
}
func (c *WsClient) send() error {
	return c.C.Send()
}

func (c *WsClient) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.C.WsApiServe(ctx)
}

func (c *WsClient) NewWebsocketStreams() *WebsocketStreams {
	return &WebsocketStreams{c: c}
}

// NewDepth Order Book
func (c *WsClient) NewDepth() *WsDepth {
	c.C.SetReq("depth")
	return &WsDepth{c: c}
}

// NewTickerPrice Symbol Price Ticker
func (c *WsClient) NewTickerPrice() *WsTickerPrice {
	c.C.SetReq("ticker.price")
	return &WsTickerPrice{c: c}
}

// NewTickerBook Symbol Order Book Ticker
func (c *WsClient) NewTickerBook() *WsTickerBook {
	c.C.SetReq("ticker.book")
	return &WsTickerBook{c: c}
}

// NewCreateOrder New Order(TRADE)
func (c *WsClient) NewCreateOrder() *WsCreateOrder {
	c.C.SetReq("order.place", core.AuthSigned)
	return &WsCreateOrder{c: c}
}

// NewModifyOrder Modify Order (TRADE)
func (c *WsClient) NewModifyOrder() *WsModifyOrder {
	c.C.SetReq("order.modify", core.AuthSigned)
	return &WsModifyOrder{c: c}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *WsClient) NewCancelOrder() *WsCancelOrder {
	c.C.SetReq("order.cancel", core.AuthSigned)
	return &WsCancelOrder{c: c}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *WsClient) NewQueryOrder() *WsQueryOrder {
	c.C.SetReq("order.status", core.AuthSigned)
	return &WsQueryOrder{c: c}
}

// NewPositionInfo Position Information V2 (USER_DATA)
func (c *WsClient) NewPositionInfo() *WsPositionInfo {
	c.C.SetReq("v2/account.position", core.AuthSigned)
	return &WsPositionInfo{c: c}
}

// NewAccountBalance Futures Account Balance V2(USER_DATA)
func (c *WsClient) NewAccountBalance() *WsAccountBalance {
	c.C.SetReq("v2/account.balance", core.AuthSigned)
	return &WsAccountBalance{c: c}
}

// NewAccountInfo Account Information V2(USER_DATA)
func (c *WsClient) NewAccountInfo() *WsAccountInfo {
	c.C.SetReq("v2/account.status", core.AuthSigned)
	return &WsAccountInfo{c: c}
}

// NewSessionLogon Log in with API key (SIGNED)
func (c *WsClient) NewSessionLogon() *SessionLogon {
	c.C.SetReq("session.logon", core.AuthSigned)
	return &SessionLogon{c: c}
}

// NewSessionStatus Query session status (SIGNED)
func (c *WsClient) NewSessionStatus() *SessionStatus {
	c.C.SetReq("session.status", core.AuthSigned)
	return &SessionStatus{c: c}
}

// NewSessionLogout Log out of the session
func (c *WsClient) NewSessionLogout() *SessionLogout {
	c.C.SetReq("session.logout", core.AuthSigned)
	return &SessionLogout{c: c}
}
