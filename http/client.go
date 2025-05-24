package http

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Client struct {
	Opt        *Options
	HttpClient *http.Client
	fullUrl    string
	req        *request
	resp       *response
}

func (c *Client) parseRequest() error {
	if c.req.authType == authSigned {
		c.req.set("timestamp", time.Now().UnixMilli())
	}
	fullUrl := fmt.Sprintf("%s%s", c.Opt.Endpoint, c.req.path)
	query := c.req.query.Encode()
	form := c.req.form.Encode()
	header := http.Header{}
	if c.req.header != nil {
		header = c.req.header.Clone()
	}
	if c.req.authType == authApiKey || c.req.authType == authSigned {
		header.Set("X-MBX-APIKEY", c.Opt.ApiKey)
	}
	if form != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		c.req.body = bytes.NewBufferString(form)
	}
	if c.req.authType == authSigned {
		mac := hmac.New(sha256.New, []byte(c.Opt.ApiSecret))
		if _, err := mac.Write([]byte(fmt.Sprintf("%s%s", query, form))); err != nil {
			return err
		}
		sign := fmt.Sprintf("signature=%x", mac.Sum(nil))
		if query == "" {
			query = sign
		} else {
			query = fmt.Sprintf("%s&%s", query, sign)
		}
	}
	if query != "" {
		fullUrl = fmt.Sprintf("%s?%s", fullUrl, query)
	}
	c.fullUrl = fullUrl
	c.req.header = header
	return nil
}

func (c *Client) invoke(ctx context.Context) error {
	if err := c.parseRequest(); err != nil {
		return err
	}
	req, err := http.NewRequest(c.req.method, c.fullUrl, c.req.body)
	if err != nil {
		c.resp = &response{err: err}
		return err
	}
	req = req.WithContext(ctx)
	if c.req.header != nil {
		req.Header = c.req.header
	}
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.resp = &response{err: err}
		return err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		c.resp = &response{err: err}
		return err
	}
	defer res.Body.Close()
	c.resp = &response{rawBody: data, status: res.StatusCode}
	if res.StatusCode != 200 {
		c.resp.err = errors.New(string(data))
		return c.resp.err
	}
	return nil
}

// NewPing Test connectivity
func (c *Client) NewPing() *Ping {
	c.req = &request{path: "/api/v3/ping", method: http.MethodGet}
	return &Ping{c: c}
}

// NewServerTime Check server time
func (c *Client) NewServerTime() *ServerTime {
	c.req = &request{path: "/api/v3/time", method: http.MethodGet}
	return &ServerTime{c: c}
}

// NewExchangeInfo Exchange information
func (c *Client) NewExchangeInfo() *ExchangeInfo {
	c.req = &request{path: "/api/v3/exchangeInfo", method: http.MethodGet}
	return &ExchangeInfo{c: c}
}

// NewDepth order book
func (c *Client) NewDepth() *Depth {
	c.req = &request{path: "/api/v3/depth", method: http.MethodGet}
	return &Depth{c: c}
}

// NewTrades Recent trades list
func (c *Client) NewTrades() *Trades {
	c.req = &request{path: "/api/v3/trades", method: http.MethodGet}
	return &Trades{c: c}
}

// NewHistoricalTrades Old trade lookup
func (c *Client) NewHistoricalTrades() *HistoricalTrades {
	c.req = &request{path: "/api/v3/historicalTrades", method: http.MethodGet}
	return &HistoricalTrades{c: c}
}

// NewAggregateTrades Compressed/Aggregate trades list
func (c *Client) NewAggregateTrades() *AggregateTrades {
	c.req = &request{path: "/api/v3/aggTrades", method: http.MethodGet}
	return &AggregateTrades{c: c}
}

// NewKlineData Kline/Candlestick data
func (c *Client) NewKlineData() *KlineData {
	c.req = &request{path: "/api/v3/klines", method: http.MethodGet}
	return &KlineData{c: c}
}

// NewUIKlines UI Klines
func (c *Client) NewUIKlines() *UIKlines {
	c.req = &request{path: "/api/v3/uiKlines", method: http.MethodGet}
	return &UIKlines{c: c}
}

// NewAveragePrice Current average price
func (c *Client) NewAveragePrice() *AveragePrice {
	c.req = &request{path: "/api/v3/avgPrice", method: http.MethodGet}
	return &AveragePrice{c: c}
}

// NewTickerPrice24h 24hr ticker price change statistics
func (c *Client) NewTickerPrice24h() *TickerPrice24h {
	c.req = &request{path: "/api/v3/ticker/24hr", method: http.MethodGet}
	return &TickerPrice24h{c: c}
}

// NewTradingDayTicker Trading Day Ticker
func (c *Client) NewTradingDayTicker() *TradingDayTicker {
	c.req = &request{path: "/api/v3/ticker/tradingDay", method: http.MethodGet}
	return &TradingDayTicker{c: c}
}

// NewPriceTicker Symbol price ticker
func (c *Client) NewPriceTicker() *PriceTicker {
	c.req = &request{path: "/api/v3/ticker/price", method: http.MethodGet}
	return &PriceTicker{c: c}
}

// NewOrderBookTicker Symbol Order Book Ticker
func (c *Client) NewOrderBookTicker() *OrderBookTicker {
	c.req = &request{path: "/api/v3/ticker/bookTicker", method: http.MethodGet}
	return &OrderBookTicker{c: c}
}

// NewTicker Rolling window price change statistics
func (c *Client) NewTicker() *Ticker {
	c.req = &request{path: "/api/v3/ticker", method: http.MethodGet}
	return &Ticker{c: c}
}

// NewCreateOrder New order (TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	c.req = &request{path: "/api/v3/order", method: http.MethodPost, authType: authSigned}
	return &CreateOrder{c: c}
}

// NewTestCreateOrder Test new order (TRADE)
func (c *Client) NewTestCreateOrder() *TestCreateOrder {
	c.req = &request{path: "/api/v3/order/test", method: http.MethodPost, authType: authSigned}
	return &TestCreateOrder{c: c}
}

// NewQueryOrder Query order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.req = &request{path: "/api/v3/order", method: http.MethodGet, authType: authSigned}
	return &QueryOrder{c: c}
}

// NewCancelOrder Cancel order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.req = &request{path: "/api/v3/order", method: http.MethodDelete, authType: authSigned}
	return &CancelOrder{c: c}
}

// NewCancelOpenOrder Cancel All Open Orders on a Symbol (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	c.req = &request{path: "/api/v3/openOrders", method: http.MethodDelete, authType: authSigned}
	return &CancelOpenOrder{c: c}
}

// NewCancelReplace Cancel an Existing Order and Send a New Order (TRADE)
func (c *Client) NewCancelReplace() *CancelReplace {
	c.req = &request{path: "/api/v3/order/cancelReplace", method: http.MethodPost, authType: authSigned}
	return &CancelReplace{c: c}
}

// NewOpenOrders Current open orders (USER_DATA)
func (c *Client) NewOpenOrders() *OpenOrders {
	c.req = &request{path: "/api/v3/openOrders", method: http.MethodGet, authType: authSigned}
	return &OpenOrders{c: c}
}

// NewAllOrders All orders (USER_DATA)
func (c *Client) NewAllOrders() *AllOrders {
	c.req = &request{path: "/api/v3/allOrders", method: http.MethodGet, authType: authSigned}
	return &AllOrders{c: c}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *Client) NewCancelOrderList() *CancelOrderList {
	c.req = &request{path: "/api/v3/orderList", method: http.MethodDelete, authType: authSigned}
	return &CancelOrderList{c: c}
}

// NewQueryOrderList Query Order List (USER_DATA)
func (c *Client) NewQueryOrderList() *QueryOrderList {
	c.req = &request{path: "/api/v3/orderList", method: http.MethodGet, authType: authSigned}
	return &QueryOrderList{c: c}
}

// NewQueryAllOrderLists Query All Order Lists (USER_DATA)
func (c *Client) NewQueryAllOrderLists() *QueryAllOrderLists {
	c.req = &request{path: "/api/v3/allOrderList", method: http.MethodGet, authType: authSigned}
	return &QueryAllOrderLists{c: c}
}

// NewQueryOpenOrderList Query Open Order Lists (USER_DATA)
func (c *Client) NewQueryOpenOrderList() *QueryOpenOrderList {
	c.req = &request{path: "/api/v3/openOrderList", method: http.MethodGet, authType: authSigned}
	return &QueryOpenOrderList{c: c}
}

// NewCreateSOROrder Create SOR Order (TRADE)
func (c *Client) NewCreateSOROrder() *CreateSOROrder {
	c.req = &request{path: "/api/v3/sor/order", method: http.MethodPost, authType: authSigned}
	return &CreateSOROrder{c: c}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *Client) NewCreateTestSOROrder() *CreateTestSOROrder {
	c.req = &request{path: "/api/v3/sor/order/test", method: http.MethodPost, authType: authSigned}
	return &CreateTestSOROrder{c: c}
}

// NewAccountInfo Account information (USER_DATA)
func (c *Client) NewAccountInfo() *AccountInfo {
	c.req = &request{path: "/api/v3/account", method: http.MethodGet, authType: authSigned}
	return &AccountInfo{c: c}
}

// NewAccountTrade Account trade list (USER_DATA)
func (c *Client) NewAccountTrade() *AccountTrade {
	c.req = &request{path: "/api/v3/myTrades", method: http.MethodGet, authType: authSigned}
	return &AccountTrade{c: c}
}

// NewQueryUnfilledOrder Query Unfilled Order Count (USER_DATA)
func (c *Client) NewQueryUnfilledOrder() *QueryUnfilledOrder {
	c.req = &request{path: "/api/v3/rateLimit/order", method: http.MethodGet, authType: authSigned}
	return &QueryUnfilledOrder{c: c}
}

// NewQueryPreventedMatches Query Prevented Matches (USER_DATA)
func (c *Client) NewQueryPreventedMatches() *QueryPreventedMatches {
	c.req = &request{path: "/api/v3/myPreventedMatches", method: http.MethodGet, authType: authSigned}
	return &QueryPreventedMatches{c: c}
}

// NewQueryAllocations Query Allocations (USER_DATA)
func (c *Client) NewQueryAllocations() *QueryAllocations {
	c.req = &request{path: "/api/v3/myAllocations", method: http.MethodGet, authType: authSigned}
	return &QueryAllocations{c: c}
}

// NewQueryCommission Query Commission Rates (USER_DATA)
func (c *Client) NewQueryCommission() *QueryCommission {
	c.req = &request{path: "/api/v3/account/commission", method: http.MethodGet, authType: authSigned}
	return &QueryCommission{c: c}
}

// NewStartUserDataStream Start user data stream (USER_STREAM)
func (c *Client) NewStartUserDataStream() *StartUserDataStream {
	c.req = &request{path: "/api/v3/userDataStream", method: http.MethodPost, authType: authApiKey}
	return &StartUserDataStream{c: c}
}

// NewCloseUserDataStream Close user data stream (USER_STREAM)
func (c *Client) NewCloseUserDataStream() *CloseUserDataStream {
	c.req = &request{path: "/api/v3/userDataStream", method: http.MethodDelete, authType: authApiKey}
	return &CloseUserDataStream{c: c}
}

// NewPingUserDataStream Keepalive user data stream (USER_STREAM)
func (c *Client) NewPingUserDataStream() *PingUserDataStream {
	c.req = &request{path: "/api/v3/userDataStream", method: http.MethodPut, authType: authApiKey}
	return &PingUserDataStream{c: c}
}
