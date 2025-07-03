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
func (c *WsClient) send(r *core.WsRequest) error {
	return c.Send(r)
}

func (c *WsClient) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.WsApiServe(ctx)
}

func (c *WsClient) NewWebsocketStreams() *WebsocketStreams {
	return &WebsocketStreams{c: c}
}

// NewDepth Order Book
func (c *WsClient) NewDepth() *WsDepth {
	return &WsDepth{c: c, r: c.SetReq("depth")}
}

// NewTickerPrice Symbol Price Ticker
func (c *WsClient) NewTickerPrice() *WsTickerPrice {
	return &WsTickerPrice{c: c, r: c.SetReq("ticker.price")}
}

// NewTickerBook Symbol Order Book Ticker
func (c *WsClient) NewTickerBook() *WsTickerBook {
	return &WsTickerBook{c: c, r: c.SetReq("ticker.book")}
}

// NewCreateOrder New Order(TRADE)
func (c *WsClient) NewCreateOrder() *WsCreateOrder {
	return &WsCreateOrder{c: c, r: c.SetReq("order.place", core.AuthSigned)}
}

// NewModifyOrder Modify Order (TRADE)
func (c *WsClient) NewModifyOrder() *WsModifyOrder {
	return &WsModifyOrder{c: c, r: c.SetReq("order.modify", core.AuthSigned)}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *WsClient) NewCancelOrder() *WsCancelOrder {
	return &WsCancelOrder{c: c, r: c.SetReq("order.cancel", core.AuthSigned)}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *WsClient) NewQueryOrder() *WsQueryOrder {
	return &WsQueryOrder{c: c, r: c.SetReq("order.status", core.AuthSigned)}
}

// NewPositionInfo Position Information V2 (USER_DATA)
func (c *WsClient) NewPositionInfo() *WsPositionInfo {
	return &WsPositionInfo{c: c, r: c.SetReq("v2/account.position", core.AuthSigned)}
}

// NewAccountBalance Futures Account Balance V2(USER_DATA)
func (c *WsClient) NewAccountBalance() *WsAccountBalance {
	return &WsAccountBalance{c: c, r: c.SetReq("v2/account.balance", core.AuthSigned)}
}

// NewAccountInfo Account Information V2(USER_DATA)
func (c *WsClient) NewAccountInfo() *WsAccountInfo {
	return &WsAccountInfo{c: c, r: c.SetReq("v2/account.status", core.AuthSigned)}
}

// NewSessionLogon Log in with API key (SIGNED)
func (c *WsClient) NewSessionLogon() *SessionLogon {
	return &SessionLogon{c: c, r: c.SetReq("session.logon", core.AuthSigned)}
}

// NewSessionStatus Query session status (SIGNED)
func (c *WsClient) NewSessionStatus() *SessionStatus {
	return &SessionStatus{c: c, r: c.SetReq("session.status", core.AuthSigned)}
}

// NewSessionLogout Log out of the session
func (c *WsClient) NewSessionLogout() *SessionLogout {
	return &SessionLogout{c: c, r: c.SetReq("session.logout", core.AuthSigned)}
}
