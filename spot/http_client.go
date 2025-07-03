package spot

import (
	"context"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
	"net/http"
)

type Client struct {
	*core.Client
}

func (c *Client) invoke(r *core.Request, ctx context.Context) error {
	return c.Invoke(r, ctx)
}

func (c *Client) rawBody() []byte {
	return c.RawBody()
}

type Fill struct {
	Price           decimal.Decimal `json:"price"`
	Qty             decimal.Decimal `json:"qty"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	TradeId         int             `json:"tradeId"`
	MatchType       string          `json:"matchType"`
	AllocId         int             `json:"allocId"`
}

type OrderReport struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int64           `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	OrigClientOrderId       string          `json:"origClientOrderId"`
	TransactTime            int64           `json:"transactTime"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice,omitempty"`
	IcebergQty              decimal.Decimal `json:"icebergQty,omitempty"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
}

type Order struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type Commission struct {
	Maker  decimal.Decimal `json:"maker"`
	Taker  decimal.Decimal `json:"taker"`
	Buyer  decimal.Decimal `json:"buyer"`
	Seller decimal.Decimal `json:"seller"`
}

type ApiBalance struct {
	Asset  string          `json:"asset"`
	Free   decimal.Decimal `json:"free"`
	Locked decimal.Decimal `json:"locked"`
}

// NewPing Test connectivity
func (c *Client) NewPing() *Ping {

	return &Ping{c: c, r: c.SetReq("/api/v3/ping", http.MethodGet)}
}

// NewServerTime Check server time
func (c *Client) NewServerTime() *ServerTime {
	return &ServerTime{c: c, r: c.SetReq("/api/v3/time", http.MethodGet)}
}

// NewExchangeInfo Exchange information
func (c *Client) NewExchangeInfo() *ExchangeInfo {
	return &ExchangeInfo{c: c, r: c.SetReq("/api/v3/exchangeInfo", http.MethodGet)}
}

// NewDepth order book
func (c *Client) NewDepth() *Depth {
	return &Depth{c: c, r: c.SetReq("/api/v3/depth", http.MethodGet)}
}

// NewTrades Recent trades list
func (c *Client) NewTrades() *Trades {
	return &Trades{c: c, r: c.SetReq("/api/v3/trades", http.MethodGet)}
}

// NewHistoricalTrades Old trade lookup
func (c *Client) NewHistoricalTrades() *HistoricalTrades {
	return &HistoricalTrades{c: c, r: c.SetReq("/api/v3/historicalTrades", http.MethodGet)}
}

// NewAggTrades Compressed/Aggregate trades list
func (c *Client) NewAggTrades() *AggTrades {
	return &AggTrades{c: c, r: c.SetReq("/api/v3/aggTrades", http.MethodGet)}
}

// NewKline Kline/Candlestick data
func (c *Client) NewKline() *KlineData {
	return &KlineData{c: c, r: c.SetReq("/api/v3/klines", http.MethodGet)}
}

// NewUIKlines UI Klines
func (c *Client) NewUIKlines() *UIKlines {
	return &UIKlines{c: c, r: c.SetReq("/api/v3/uiKlines", http.MethodGet)}
}

// NewAvgPrice Current average price
func (c *Client) NewAvgPrice() *AveragePrice {
	return &AveragePrice{c: c, r: c.SetReq("/api/v3/avgPrice", http.MethodGet)}
}

// NewTickerPrice24h 24hr ticker price change statistics
func (c *Client) NewTickerPrice24h() *TickerPrice24h {
	return &TickerPrice24h{c: c, r: c.SetReq("/api/v3/ticker/24hr", http.MethodGet)}
}

// NewTradingDayTicker Trading Day Ticker
func (c *Client) NewTradingDayTicker() *TradingDayTicker {
	return &TradingDayTicker{c: c, r: c.SetReq("/api/v3/ticker/tradingDay", http.MethodGet)}
}

// NewTickerPrice Symbol price ticker
func (c *Client) NewTickerPrice() *PriceTicker {
	return &PriceTicker{c: c, r: c.SetReq("/api/v3/ticker/price", http.MethodGet)}
}

// NewBookTicker Symbol Order Book Ticker
func (c *Client) NewBookTicker() *OrderBookTicker {
	return &OrderBookTicker{c: c, r: c.SetReq("/api/v3/ticker/bookTicker", http.MethodGet)}
}

// NewTicker Rolling window price change statistics
func (c *Client) NewTicker() *Ticker {
	return &Ticker{c: c, r: c.SetReq("/api/v3/ticker", http.MethodGet)}
}

// NewCreateOrder New order (TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	return &CreateOrder{c: c, r: c.SetReq("/api/v3/order", http.MethodPost, core.AuthSigned)}
}

// NewTestCreateOrder Test new order (TRADE)
func (c *Client) NewTestCreateOrder() *TestCreateOrder {
	return &TestCreateOrder{c: c, r: c.SetReq("/api/v3/order/test", http.MethodPost, core.AuthSigned)}
}

// NewQueryOrder Query order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	return &QueryOrder{c: c, r: c.SetReq("/api/v3/order", http.MethodGet, core.AuthSigned)}
}

// NewCancelOrder Cancel order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	return &CancelOrder{c: c, r: c.SetReq("/api/v3/order", http.MethodDelete, core.AuthSigned)}
}

// NewCancelOpenOrder Cancel All Open Orders on a Symbol (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	return &CancelOpenOrder{c: c, r: c.SetReq("/api/v3/openOrders", http.MethodDelete, core.AuthSigned)}
}

// NewCancelReplace Cancel an Existing Order and Send a New Order (TRADE)
func (c *Client) NewCancelReplace() *CancelReplace {
	return &CancelReplace{c: c, r: c.SetReq("/api/v3/order/cancelReplace", http.MethodPost, core.AuthSigned)}
}

// NewOpenOrders Current open orders (USER_DATA)
func (c *Client) NewOpenOrders() *OpenOrders {
	return &OpenOrders{c: c, r: c.SetReq("/api/v3/openOrders", http.MethodGet, core.AuthSigned)}
}

// NewAllOrders All orders (USER_DATA)
func (c *Client) NewAllOrders() *AllOrders {
	return &AllOrders{c: c, r: c.SetReq("/api/v3/allOrders", http.MethodGet, core.AuthSigned)}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *Client) NewCancelOrderList() *CancelOrderList {
	return &CancelOrderList{c: c, r: c.SetReq("/api/v3/orderList", http.MethodDelete, core.AuthSigned)}
}

// NewQueryOrderList Query Order List (USER_DATA)
func (c *Client) NewQueryOrderList() *QueryOrderList {
	return &QueryOrderList{c: c, r: c.SetReq("/api/v3/orderList", http.MethodGet, core.AuthSigned)}
}

// NewQueryAllOrderLists Query All Order Lists (USER_DATA)
func (c *Client) NewQueryAllOrderLists() *QueryAllOrderLists {
	return &QueryAllOrderLists{c: c, r: c.SetReq("/api/v3/allOrderList", http.MethodGet, core.AuthSigned)}
}

// NewQueryOpenOrderList Query Open Order Lists (USER_DATA)
func (c *Client) NewQueryOpenOrderList() *QueryOpenOrderList {
	return &QueryOpenOrderList{c: c, r: c.SetReq("/api/v3/openOrderList", http.MethodGet, core.AuthSigned)}
}

// NewCreateSOROrder Create SOR Order (TRADE)
func (c *Client) NewCreateSOROrder() *CreateSOROrder {
	return &CreateSOROrder{c: c, r: c.SetReq("/api/v3/sor/order", http.MethodPost, core.AuthSigned)}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *Client) NewCreateTestSOROrder() *CreateTestSOROrder {
	return &CreateTestSOROrder{c: c, r: c.SetReq("/api/v3/sor/order/test", http.MethodPost, core.AuthSigned)}
}

// NewAccountInfo Account information (USER_DATA)
func (c *Client) NewAccountInfo() *AccountInfo {
	return &AccountInfo{c: c, r: c.SetReq("/api/v3/account", http.MethodGet, core.AuthSigned)}
}

// NewAccountTrade Account trade list (USER_DATA)
func (c *Client) NewAccountTrade() *AccountTrade {
	return &AccountTrade{c: c, r: c.SetReq("/api/v3/myTrades", http.MethodGet, core.AuthSigned)}
}

// NewQueryUnfilledOrder Query Unfilled Order Count (USER_DATA)
func (c *Client) NewQueryUnfilledOrder() *QueryUnfilledOrder {
	return &QueryUnfilledOrder{c: c, r: c.SetReq("/api/v3/rateLimit/order", http.MethodGet, core.AuthSigned)}
}

// NewQueryPreventedMatches Query Prevented Matches (USER_DATA)
func (c *Client) NewQueryPreventedMatches() *QueryPreventedMatches {
	return &QueryPreventedMatches{c: c, r: c.SetReq("/api/v3/myPreventedMatches", http.MethodGet, core.AuthSigned)}
}

// NewQueryAllocations Query Allocations (USER_DATA)
func (c *Client) NewQueryAllocations() *QueryAllocations {
	return &QueryAllocations{c: c, r: c.SetReq("/api/v3/myAllocations", http.MethodGet, core.AuthSigned)}
}

// NewQueryCommission Query Commission Rates (USER_DATA)
func (c *Client) NewQueryCommission() *QueryCommission {
	return &QueryCommission{c: c, r: c.SetReq("/api/v3/account/commission", http.MethodGet, core.AuthSigned)}
}

// NewStartUserDataStream Start user data stream (USER_STREAM)
func (c *Client) NewStartUserDataStream() *StartUserDataStream {
	return &StartUserDataStream{c: c, r: c.SetReq("/api/v3/userDataStream", http.MethodPost, core.AuthApiKey)}
}

// NewCloseUserDataStream Close user data stream (USER_STREAM)
func (c *Client) NewCloseUserDataStream() *CloseUserDataStream {
	return &CloseUserDataStream{c: c, r: c.SetReq("/api/v3/userDataStream", http.MethodDelete, core.AuthApiKey)}
}

// NewPingUserDataStream Keepalive user data stream (USER_STREAM)
func (c *Client) NewPingUserDataStream() *PingUserDataStream {
	return &PingUserDataStream{c: c, r: c.SetReq("/api/v3/userDataStream", http.MethodPut, core.AuthApiKey)}
}
