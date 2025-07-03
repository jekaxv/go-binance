package spot

import (
	"context"
	"github.com/jekaxv/go-binance/core"
)

type WsClient struct {
	*core.WsClient
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

// NewPing Test connectivity
func (c *WsClient) NewPing() *WsPing {
	return &WsPing{c: c, r: c.SetReq("ping")}
}

// NewCheckServerTime Check server time
func (c *WsClient) NewCheckServerTime() *WsServerTime {
	return &WsServerTime{c: c, r: c.SetReq("time")}
}

// NewExchangeInfo Exchange information
func (c *WsClient) NewExchangeInfo() *WsExchangeInfo {
	return &WsExchangeInfo{c: c, r: c.SetReq("exchangeInfo")}
}

// NewDepth Order book
func (c *WsClient) NewDepth() *WsDepth {
	return &WsDepth{c: c, r: c.SetReq("depth")}
}

// NewTradesRecent Recent trades
func (c *WsClient) NewTradesRecent() *WsTradesRecent {
	return &WsTradesRecent{c: c, r: c.SetReq("trades.recent")}
}

// NewTradesHistorical Historical trades
func (c *WsClient) NewTradesHistorical() *WsTradesHistorical {
	return &WsTradesHistorical{c: c, r: c.SetReq("trades.historical")}
}

// NewTradesAggregate Aggregate trades
func (c *WsClient) NewTradesAggregate() *WsTradesAggregate {
	return &WsTradesAggregate{c: c, r: c.SetReq("trades.aggregate")}
}

// NewKline Klines
func (c *WsClient) NewKline() *WsKline {
	return &WsKline{c: c, r: c.SetReq("klines")}
}

// NewUiKlines UI Klines
func (c *WsClient) NewUiKlines() *WsUiKlines {
	return &WsUiKlines{c: c, r: c.SetReq("uiKlines")}
}

// NewAveragePrice Current average price
func (c *WsClient) NewAveragePrice() *WsAveragePrice {
	return &WsAveragePrice{c: c, r: c.SetReq("avgPrice")}
}

// NewTicker24h 24hr ticker price change statistics
func (c *WsClient) NewTicker24h() *WsTicker24h {
	return &WsTicker24h{c: c, r: c.SetReq("ticker.24hr")}
}

// NewTickerTradingDay Trading Day Ticker
func (c *WsClient) NewTickerTradingDay() *WsTickerTradingDay {
	return &WsTickerTradingDay{c: c, r: c.SetReq("ticker.tradingDay")}
}

// NewTicker Rolling window price change statistics
func (c *WsClient) NewTicker() *WsTicker {
	return &WsTicker{c: c, r: c.SetReq("ticker")}
}

// NewTickerPrice Symbol price ticker
func (c *WsClient) NewTickerPrice() *WsTickerPrice {
	return &WsTickerPrice{c: c, r: c.SetReq("ticker.price")}
}

// NewTickerBook Symbol order book ticker
func (c *WsClient) NewTickerBook() *WsTickerBook {
	return &WsTickerBook{c: c, r: c.SetReq("ticker.book")}
}

// NewCreateOrder Place new order (TRADE)
func (c *WsClient) NewCreateOrder() *WsCreateOrder {
	return &WsCreateOrder{c: c, r: c.SetReq("order.place", core.AuthSigned)}
}

// NewCreateTestOrder Test new order (TRADE)
func (c *WsClient) NewCreateTestOrder() *WsCreateTestOrder {
	return &WsCreateTestOrder{c: c, r: c.SetReq("order.test", core.AuthSigned)}
}

// NewQueryOrder Query order (USER_DATA)
func (c *WsClient) NewQueryOrder() *WsQueryOrder {
	return &WsQueryOrder{c: c, r: c.SetReq("order.status", core.AuthSigned)}
}

// NewCancelOrder Cancel order (TRADE)
func (c *WsClient) NewCancelOrder() *WsCancelOrder {
	return &WsCancelOrder{c: c, r: c.SetReq("order.cancel", core.AuthSigned)}
}

// NewCancelReplaceOrder Cancel and replace order (TRADE)
func (c *WsClient) NewCancelReplaceOrder() *WsCancelReplaceOrder {
	return &WsCancelReplaceOrder{c: c, r: c.SetReq("order.cancelReplace", core.AuthSigned)}
}

// NewOpenOrdersStatus Current open orders (USER_DATA)
func (c *WsClient) NewOpenOrdersStatus() *WsOpenOrdersStatus {
	return &WsOpenOrdersStatus{c: c, r: c.SetReq("openOrders.status", core.AuthSigned)}
}

// NewCancelOpenOrder Cancel open orders (TRADE)
func (c *WsClient) NewCancelOpenOrder() *WsCancelOpenOrder {
	return &WsCancelOpenOrder{c: c, r: c.SetReq("openOrders.cancelAll", core.AuthSigned)}
}

// NewCreateOCOOrder Place new Order list - OCO (TRADE)
func (c *WsClient) NewCreateOCOOrder() *WsCreateOCOOrder {
	return &WsCreateOCOOrder{c: c, r: c.SetReq("orderList.place.oco", core.AuthSigned)}
}

// NewCreateOTOOrder Place new Order list - OTO (TRADE)
func (c *WsClient) NewCreateOTOOrder() *WsCreateOTOOrder {
	return &WsCreateOTOOrder{c: c, r: c.SetReq("orderList.place.oto", core.AuthSigned)}
}

// NewCreateOTOCOOrder Place new Order list - OTOCO (TRADE)
func (c *WsClient) NewCreateOTOCOOrder() *WsCreateOTOCOOrder {
	return &WsCreateOTOCOOrder{c: c, r: c.SetReq("orderList.place.otoco", core.AuthSigned)}
}

// NewQueryOrderList Query Order list (USER_DATA)
func (c *WsClient) NewQueryOrderList() *WsQueryOrderList {
	return &WsQueryOrderList{c: c, r: c.SetReq("orderList.status", core.AuthSigned)}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *WsClient) NewCancelOrderList() *WsCancelOrderList {
	return &WsCancelOrderList{c: c, r: c.SetReq("orderList.cancel", core.AuthSigned)}
}

// NewQueryOpenOrder Current open Order lists (USER_DATA)
func (c *WsClient) NewQueryOpenOrder() *WsQueryOpenOrder {
	return &WsQueryOpenOrder{c: c, r: c.SetReq("openOrderLists.status", core.AuthSigned)}
}

// NewCreateSOROrder Place new order using SOR (TRADE)
func (c *WsClient) NewCreateSOROrder() *WsCreateSOROrder {
	return &WsCreateSOROrder{c: c, r: c.SetReq("sor.order.place", core.AuthSigned)}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *WsClient) NewCreateTestSOROrder() *WsCreateTestSOROrder {
	return &WsCreateTestSOROrder{c: c, r: c.SetReq("sor.order.test", core.AuthSigned)}
}

// NewAccountInformation Account information (USER_DATA)
func (c *WsClient) NewAccountInformation() *AccountInformation {
	return &AccountInformation{c: c, r: c.SetReq("account.status", core.AuthSigned)}
}

// NewUnfilledOrder Unfilled Order Count (USER_DATA)
func (c *WsClient) NewUnfilledOrder() *UnfilledOrder {
	return &UnfilledOrder{c: c, r: c.SetReq("account.rateLimits.orders", core.AuthSigned)}
}

// NewAccountOrderHistory Account order history (USER_DATA)
func (c *WsClient) NewAccountOrderHistory() *AccountOrderHistory {
	return &AccountOrderHistory{c: c, r: c.SetReq("allOrders", core.AuthSigned)}
}

// NewAllOrderList Account Order list history (USER_DATA)
func (c *WsClient) NewAllOrderList() *AllOrderList {
	return &AllOrderList{c: c, r: c.SetReq("allOrderLists", core.AuthSigned)}
}

// NewAccountTradeHistory Account trade history (USER_DATA)
func (c *WsClient) NewAccountTradeHistory() *AccountTradeHistory {
	return &AccountTradeHistory{c: c, r: c.SetReq("myTrades", core.AuthSigned)}
}

// NewAccountPreventedMatches Account prevented matches (USER_DATA)
func (c *WsClient) NewAccountPreventedMatches() *AccountPreventedMatches {
	return &AccountPreventedMatches{c: c, r: c.SetReq("myPreventedMatches", core.AuthSigned)}
}

// NewAccountAllocations Account allocations (USER_DATA)
func (c *WsClient) NewAccountAllocations() *AccountAllocations {
	return &AccountAllocations{c: c, r: c.SetReq("myAllocations", core.AuthSigned)}
}

// NewAccountCommission Account Commission Rates (USER_DATA)
func (c *WsClient) NewAccountCommission() *AccountCommission {
	return &AccountCommission{c: c, r: c.SetReq("account.commission", core.AuthSigned)}
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
