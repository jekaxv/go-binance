package spot

import (
	"context"
	"github.com/jekaxv/go-binance/core"
)

type WsClient struct {
	C *core.WsClient
}

func (c *WsClient) getParams(key string) any {
	return c.C.GetParams(key)
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

// NewPing Test connectivity
func (c *WsClient) NewPing() *WsPing {
	c.C.SetReq("ping")
	return &WsPing{c: c}
}

// NewCheckServerTime Check server time
func (c *WsClient) NewCheckServerTime() *WsServerTime {
	c.C.SetReq("time")
	return &WsServerTime{c: c}
}

// NewExchangeInfo Exchange information
func (c *WsClient) NewExchangeInfo() *WsExchangeInfo {
	c.C.SetReq("exchangeInfo")
	return &WsExchangeInfo{c: c}
}

// NewDepth Order book
func (c *WsClient) NewDepth() *WsDepth {
	c.C.SetReq("depth")
	return &WsDepth{c: c}
}

// NewTradesRecent Recent trades
func (c *WsClient) NewTradesRecent() *WsTradesRecent {
	c.C.SetReq("trades.recent")
	return &WsTradesRecent{c: c}
}

// NewTradesHistorical Historical trades
func (c *WsClient) NewTradesHistorical() *WsTradesHistorical {
	c.C.SetReq("trades.historical")
	return &WsTradesHistorical{c: c}
}

// NewTradesAggregate Aggregate trades
func (c *WsClient) NewTradesAggregate() *WsTradesAggregate {
	c.C.SetReq("trades.aggregate")
	return &WsTradesAggregate{c: c}
}

// NewKline Klines
func (c *WsClient) NewKline() *WsKline {
	c.C.SetReq("klines")
	return &WsKline{c: c}
}

// NewUiKlines UI Klines
func (c *WsClient) NewUiKlines() *WsUiKlines {
	c.C.SetReq("uiKlines")
	return &WsUiKlines{c: c}
}

// NewAveragePrice Current average price
func (c *WsClient) NewAveragePrice() *WsAveragePrice {
	c.C.SetReq("avgPrice")
	return &WsAveragePrice{c: c}
}

// NewTicker24h 24hr ticker price change statistics
func (c *WsClient) NewTicker24h() *WsTicker24h {
	c.C.SetReq("ticker.24hr")
	return &WsTicker24h{c: c}
}

// NewTickerTradingDay Trading Day Ticker
func (c *WsClient) NewTickerTradingDay() *WsTickerTradingDay {
	c.C.SetReq("ticker.tradingDay")
	return &WsTickerTradingDay{c: c}
}

// NewTicker Rolling window price change statistics
func (c *WsClient) NewTicker() *WsTicker {
	c.C.SetReq("ticker")
	return &WsTicker{c: c}
}

// NewTickerPrice Symbol price ticker
func (c *WsClient) NewTickerPrice() *WsTickerPrice {
	c.C.SetReq("ticker.price")
	return &WsTickerPrice{c: c}
}

// NewTickerBook Symbol order book ticker
func (c *WsClient) NewTickerBook() *WsTickerBook {
	c.C.SetReq("ticker.book")
	return &WsTickerBook{c: c}
}

// NewCreateOrder Place new order (TRADE)
func (c *WsClient) NewCreateOrder() *WsCreateOrder {
	c.C.SetReq("order.place", core.AuthSigned)
	return &WsCreateOrder{c: c}
}

// NewCreateTestOrder Test new order (TRADE)
func (c *WsClient) NewCreateTestOrder() *WsCreateTestOrder {
	c.C.SetReq("order.test", core.AuthSigned)
	return &WsCreateTestOrder{c: c}
}

// NewQueryOrder Query order (USER_DATA)
func (c *WsClient) NewQueryOrder() *WsQueryOrder {
	c.C.SetReq("order.status", core.AuthSigned)
	return &WsQueryOrder{c: c}
}

// NewCancelOrder Cancel order (TRADE)
func (c *WsClient) NewCancelOrder() *WsCancelOrder {
	c.C.SetReq("order.cancel", core.AuthSigned)
	return &WsCancelOrder{c: c}
}

// NewCancelReplaceOrder Cancel and replace order (TRADE)
func (c *WsClient) NewCancelReplaceOrder() *WsCancelReplaceOrder {
	c.C.SetReq("order.cancelReplace", core.AuthSigned)
	return &WsCancelReplaceOrder{c: c}
}

// NewOpenOrdersStatus Current open orders (USER_DATA)
func (c *WsClient) NewOpenOrdersStatus() *WsOpenOrdersStatus {
	c.C.SetReq("openOrders.status", core.AuthSigned)
	return &WsOpenOrdersStatus{c: c}
}

// NewCancelOpenOrder Cancel open orders (TRADE)
func (c *WsClient) NewCancelOpenOrder() *WsCancelOpenOrder {
	c.C.SetReq("openOrders.cancelAll", core.AuthSigned)
	return &WsCancelOpenOrder{c: c}
}

// NewCreateOCOOrder Place new Order list - OCO (TRADE)
func (c *WsClient) NewCreateOCOOrder() *WsCreateOCOOrder {
	c.C.SetReq("orderList.place.oco", core.AuthSigned)
	return &WsCreateOCOOrder{c: c}
}

// NewCreateOTOOrder Place new Order list - OTO (TRADE)
func (c *WsClient) NewCreateOTOOrder() *WsCreateOTOOrder {
	c.C.SetReq("orderList.place.oto", core.AuthSigned)
	return &WsCreateOTOOrder{c: c}
}

// NewCreateOTOCOOrder Place new Order list - OTOCO (TRADE)
func (c *WsClient) NewCreateOTOCOOrder() *WsCreateOTOCOOrder {
	c.C.SetReq("orderList.place.otoco", core.AuthSigned)
	return &WsCreateOTOCOOrder{c: c}
}

// NewQueryOrderList Query Order list (USER_DATA)
func (c *WsClient) NewQueryOrderList() *WsQueryOrderList {
	c.C.SetReq("orderList.status", core.AuthSigned)
	return &WsQueryOrderList{c: c}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *WsClient) NewCancelOrderList() *WsCancelOrderList {
	c.C.SetReq("orderList.cancel", core.AuthSigned)
	return &WsCancelOrderList{c: c}
}

// NewQueryOpenOrder Current open Order lists (USER_DATA)
func (c *WsClient) NewQueryOpenOrder() *WsQueryOpenOrder {
	c.C.SetReq("openOrderLists.status", core.AuthSigned)
	return &WsQueryOpenOrder{c: c}
}

// NewCreateSOROrder Place new order using SOR (TRADE)
func (c *WsClient) NewCreateSOROrder() *WsCreateSOROrder {
	c.C.SetReq("sor.order.place", core.AuthSigned)
	return &WsCreateSOROrder{c: c}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *WsClient) NewCreateTestSOROrder() *WsCreateTestSOROrder {
	c.C.SetReq("sor.order.test", core.AuthSigned)
	return &WsCreateTestSOROrder{c: c}
}

// NewAccountInformation Account information (USER_DATA)
func (c *WsClient) NewAccountInformation() *AccountInformation {
	c.C.SetReq("account.status", core.AuthSigned)
	return &AccountInformation{c: c}
}

// NewUnfilledOrder Unfilled Order Count (USER_DATA)
func (c *WsClient) NewUnfilledOrder() *UnfilledOrder {
	c.C.SetReq("account.rateLimits.orders", core.AuthSigned)
	return &UnfilledOrder{c: c}
}

// NewAccountOrderHistory Account order history (USER_DATA)
func (c *WsClient) NewAccountOrderHistory() *AccountOrderHistory {
	c.C.SetReq("allOrders", core.AuthSigned)
	return &AccountOrderHistory{c: c}
}

// NewAllOrderList Account Order list history (USER_DATA)
func (c *WsClient) NewAllOrderList() *AllOrderList {
	c.C.SetReq("allOrderLists", core.AuthSigned)
	return &AllOrderList{c: c}
}

// NewAccountTradeHistory Account trade history (USER_DATA)
func (c *WsClient) NewAccountTradeHistory() *AccountTradeHistory {
	c.C.SetReq("myTrades", core.AuthSigned)
	return &AccountTradeHistory{c: c}
}

// NewAccountPreventedMatches Account prevented matches (USER_DATA)
func (c *WsClient) NewAccountPreventedMatches() *AccountPreventedMatches {
	c.C.SetReq("myPreventedMatches", core.AuthSigned)
	return &AccountPreventedMatches{c: c}
}

// NewAccountAllocations Account allocations (USER_DATA)
func (c *WsClient) NewAccountAllocations() *AccountAllocations {
	c.C.SetReq("myAllocations", core.AuthSigned)
	return &AccountAllocations{c: c}
}

// NewAccountCommission Account Commission Rates (USER_DATA)
func (c *WsClient) NewAccountCommission() *AccountCommission {
	c.C.SetReq("account.commission", core.AuthSigned)
	return &AccountCommission{c: c}
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
