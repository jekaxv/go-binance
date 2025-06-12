package futures

import (
	"context"
	"github.com/jekaxv/go-binance/core"
)

type WsClient struct {
	*core.WsClient
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
	c.SetParams(key, value)
}

func (c *WsClient) wsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.WsServe(ctx)
}

func (c *WsClient) combined(combine bool) {
	c.Combined(combine)
}

func (c *WsClient) getEndpoint() string {
	return c.Opt.Endpoint
}

func (c *WsClient) setEndpoint(endpoint string) {
	c.Opt.Endpoint = endpoint
}

func (c *WsClient) close() error {
	return c.Close()
}
func (c *WsClient) send() error {
	return c.Send()
}

func (c *WsClient) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.WsApiServe(ctx)
}

func (c *WsClient) NewWebsocketStreams() *WebsocketStreams {
	return &WebsocketStreams{c: c}
}

// NewDepth Order Book
func (c *WsClient) NewDepth() *WsDepth {
	c.SetReq("depth")
	return &WsDepth{c: c}
}

// NewTickerPrice Symbol Price Ticker
func (c *WsClient) NewTickerPrice() *WsTickerPrice {
	c.SetReq("ticker.price")
	return &WsTickerPrice{c: c}
}

// NewTickerBook Symbol Order Book Ticker
func (c *WsClient) NewTickerBook() *WsTickerBook {
	c.SetReq("ticker.book")
	return &WsTickerBook{c: c}
}

// NewCreateOrder New Order(TRADE)
func (c *WsClient) NewCreateOrder() *WsCreateOrder {
	c.SetReq("order.place", core.AuthSigned)
	return &WsCreateOrder{c: c}
}

// NewModifyOrder Modify Order (TRADE)
func (c *WsClient) NewModifyOrder() *WsModifyOrder {
	c.SetReq("order.modify", core.AuthSigned)
	return &WsModifyOrder{c: c}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *WsClient) NewCancelOrder() *WsCancelOrder {
	c.SetReq("order.cancel", core.AuthSigned)
	return &WsCancelOrder{c: c}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *WsClient) NewQueryOrder() *WsQueryOrder {
	c.SetReq("order.status", core.AuthSigned)
	return &WsQueryOrder{c: c}
}

// NewPositionInfo Position Information V2 (USER_DATA)
func (c *WsClient) NewPositionInfo() *WsPositionInfo {
	c.SetReq("v2/account.position", core.AuthSigned)
	return &WsPositionInfo{c: c}
}

// NewAccountBalance Futures Account Balance V2(USER_DATA)
func (c *WsClient) NewAccountBalance() *WsAccountBalance {
	c.SetReq("v2/account.balance", core.AuthSigned)
	return &WsAccountBalance{c: c}
}

// NewAccountInfo Account Information V2(USER_DATA)
func (c *WsClient) NewAccountInfo() *WsAccountInfo {
	c.SetReq("v2/account.status", core.AuthSigned)
	return &WsAccountInfo{c: c}
}

// NewSessionLogon Log in with API key (SIGNED)
func (c *WsClient) NewSessionLogon() *SessionLogon {
	c.SetReq("session.logon", core.AuthSigned)
	return &SessionLogon{c: c}
}

// NewSessionStatus Query session status (SIGNED)
func (c *WsClient) NewSessionStatus() *SessionStatus {
	c.SetReq("session.status", core.AuthSigned)
	return &SessionStatus{c: c}
}

// NewSessionLogout Log out of the session
func (c *WsClient) NewSessionLogout() *SessionLogout {
	c.SetReq("session.logout", core.AuthSigned)
	return &SessionLogout{c: c}
}
