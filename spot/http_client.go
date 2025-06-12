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

func (c *Client) set(key string, value any) {
	c.Set(key, value)
}

func (c *Client) invoke(ctx context.Context) error {
	return c.Invoke(ctx)
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
	c.SetReq("/api/v3/ping", http.MethodGet)
	return &Ping{c: c}
}

// NewServerTime Check server time
func (c *Client) NewServerTime() *ServerTime {
	c.SetReq("/api/v3/time", http.MethodGet)
	return &ServerTime{c: c}
}

// NewExchangeInfo Exchange information
func (c *Client) NewExchangeInfo() *ExchangeInfo {
	c.SetReq("/api/v3/exchangeInfo", http.MethodGet)
	return &ExchangeInfo{c: c}
}

// NewDepth order book
func (c *Client) NewDepth() *Depth {
	c.SetReq("/api/v3/depth", http.MethodGet)
	return &Depth{c: c}
}

// NewTrades Recent trades list
func (c *Client) NewTrades() *Trades {
	c.SetReq("/api/v3/trades", http.MethodGet)
	return &Trades{c: c}
}

// NewHistoricalTrades Old trade lookup
func (c *Client) NewHistoricalTrades() *HistoricalTrades {
	c.SetReq("/api/v3/historicalTrades", http.MethodGet)
	return &HistoricalTrades{c: c}
}

// NewAggTrades Compressed/Aggregate trades list
func (c *Client) NewAggTrades() *AggTrades {
	c.SetReq("/api/v3/aggTrades", http.MethodGet)
	return &AggTrades{c: c}
}

// NewKline Kline/Candlestick data
func (c *Client) NewKline() *KlineData {
	c.SetReq("/api/v3/klines", http.MethodGet)
	return &KlineData{c: c}
}

// NewUIKlines UI Klines
func (c *Client) NewUIKlines() *UIKlines {
	c.SetReq("/api/v3/uiKlines", http.MethodGet)
	return &UIKlines{c: c}
}

// NewAvgPrice Current average price
func (c *Client) NewAvgPrice() *AveragePrice {
	c.SetReq("/api/v3/avgPrice", http.MethodGet)
	return &AveragePrice{c: c}
}

// NewTickerPrice24h 24hr ticker price change statistics
func (c *Client) NewTickerPrice24h() *TickerPrice24h {
	c.SetReq("/api/v3/ticker/24hr", http.MethodGet)
	return &TickerPrice24h{c: c}
}

// NewTradingDayTicker Trading Day Ticker
func (c *Client) NewTradingDayTicker() *TradingDayTicker {
	c.SetReq("/api/v3/ticker/tradingDay", http.MethodGet)
	return &TradingDayTicker{c: c}
}

// NewTickerPrice Symbol price ticker
func (c *Client) NewTickerPrice() *PriceTicker {
	c.SetReq("/api/v3/ticker/price", http.MethodGet)
	return &PriceTicker{c: c}
}

// NewBookTicker Symbol Order Book Ticker
func (c *Client) NewBookTicker() *OrderBookTicker {
	c.SetReq("/api/v3/ticker/bookTicker", http.MethodGet)
	return &OrderBookTicker{c: c}
}

// NewTicker Rolling window price change statistics
func (c *Client) NewTicker() *Ticker {
	c.SetReq("/api/v3/ticker", http.MethodGet)
	return &Ticker{c: c}
}

// NewCreateOrder New order (TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	c.SetReq("/api/v3/order", http.MethodPost, core.AuthSigned)
	return &CreateOrder{c: c}
}

// NewTestCreateOrder Test new order (TRADE)
func (c *Client) NewTestCreateOrder() *TestCreateOrder {
	c.SetReq("/api/v3/order/test", http.MethodPost, core.AuthSigned)
	return &TestCreateOrder{c: c}
}

// NewQueryOrder Query order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.SetReq("/api/v3/order", http.MethodGet, core.AuthSigned)
	return &QueryOrder{c: c}
}

// NewCancelOrder Cancel order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.SetReq("/api/v3/order", http.MethodDelete, core.AuthSigned)
	return &CancelOrder{c: c}
}

// NewCancelOpenOrder Cancel All Open Orders on a Symbol (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	c.SetReq("/api/v3/openOrders", http.MethodDelete, core.AuthSigned)
	return &CancelOpenOrder{c: c}
}

// NewCancelReplace Cancel an Existing Order and Send a New Order (TRADE)
func (c *Client) NewCancelReplace() *CancelReplace {
	c.SetReq("/api/v3/order/cancelReplace", http.MethodPost, core.AuthSigned)
	return &CancelReplace{c: c}
}

// NewOpenOrders Current open orders (USER_DATA)
func (c *Client) NewOpenOrders() *OpenOrders {
	c.SetReq("/api/v3/openOrders", http.MethodGet, core.AuthSigned)
	return &OpenOrders{c: c}
}

// NewAllOrders All orders (USER_DATA)
func (c *Client) NewAllOrders() *AllOrders {
	c.SetReq("/api/v3/allOrders", http.MethodGet, core.AuthSigned)
	return &AllOrders{c: c}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *Client) NewCancelOrderList() *CancelOrderList {
	c.SetReq("/api/v3/orderList", http.MethodDelete, core.AuthSigned)
	return &CancelOrderList{c: c}
}

// NewQueryOrderList Query Order List (USER_DATA)
func (c *Client) NewQueryOrderList() *QueryOrderList {
	c.SetReq("/api/v3/orderList", http.MethodGet, core.AuthSigned)
	return &QueryOrderList{c: c}
}

// NewQueryAllOrderLists Query All Order Lists (USER_DATA)
func (c *Client) NewQueryAllOrderLists() *QueryAllOrderLists {
	c.SetReq("/api/v3/allOrderList", http.MethodGet, core.AuthSigned)
	return &QueryAllOrderLists{c: c}
}

// NewQueryOpenOrderList Query Open Order Lists (USER_DATA)
func (c *Client) NewQueryOpenOrderList() *QueryOpenOrderList {
	c.SetReq("/api/v3/openOrderList", http.MethodGet, core.AuthSigned)
	return &QueryOpenOrderList{c: c}
}

// NewCreateSOROrder Create SOR Order (TRADE)
func (c *Client) NewCreateSOROrder() *CreateSOROrder {
	c.SetReq("/api/v3/sor/order", http.MethodPost, core.AuthSigned)
	return &CreateSOROrder{c: c}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *Client) NewCreateTestSOROrder() *CreateTestSOROrder {
	c.SetReq("/api/v3/sor/order/test", http.MethodPost, core.AuthSigned)
	return &CreateTestSOROrder{c: c}
}

// NewAccountInfo Account information (USER_DATA)
func (c *Client) NewAccountInfo() *AccountInfo {
	c.SetReq("/api/v3/account", http.MethodGet, core.AuthSigned)
	return &AccountInfo{c: c}
}

// NewAccountTrade Account trade list (USER_DATA)
func (c *Client) NewAccountTrade() *AccountTrade {
	c.SetReq("/api/v3/myTrades", http.MethodGet, core.AuthSigned)
	return &AccountTrade{c: c}
}

// NewQueryUnfilledOrder Query Unfilled Order Count (USER_DATA)
func (c *Client) NewQueryUnfilledOrder() *QueryUnfilledOrder {
	c.SetReq("/api/v3/rateLimit/order", http.MethodGet, core.AuthSigned)
	return &QueryUnfilledOrder{c: c}
}

// NewQueryPreventedMatches Query Prevented Matches (USER_DATA)
func (c *Client) NewQueryPreventedMatches() *QueryPreventedMatches {
	c.SetReq("/api/v3/myPreventedMatches", http.MethodGet, core.AuthSigned)
	return &QueryPreventedMatches{c: c}
}

// NewQueryAllocations Query Allocations (USER_DATA)
func (c *Client) NewQueryAllocations() *QueryAllocations {
	c.SetReq("/api/v3/myAllocations", http.MethodGet, core.AuthSigned)
	return &QueryAllocations{c: c}
}

// NewQueryCommission Query Commission Rates (USER_DATA)
func (c *Client) NewQueryCommission() *QueryCommission {
	c.SetReq("/api/v3/account/commission", http.MethodGet, core.AuthSigned)
	return &QueryCommission{c: c}
}

// NewStartUserDataStream Start user data stream (USER_STREAM)
func (c *Client) NewStartUserDataStream() *StartUserDataStream {
	c.SetReq("/api/v3/userDataStream", http.MethodPost, core.AuthApiKey)
	return &StartUserDataStream{c: c}
}

// NewCloseUserDataStream Close user data stream (USER_STREAM)
func (c *Client) NewCloseUserDataStream() *CloseUserDataStream {
	c.SetReq("/api/v3/userDataStream", http.MethodDelete, core.AuthApiKey)
	return &CloseUserDataStream{c: c}
}

// NewPingUserDataStream Keepalive user data stream (USER_STREAM)
func (c *Client) NewPingUserDataStream() *PingUserDataStream {
	c.SetReq("/api/v3/userDataStream", http.MethodPut, core.AuthApiKey)
	return &PingUserDataStream{c: c}
}
