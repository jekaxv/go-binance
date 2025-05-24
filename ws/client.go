package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
)

type Client struct {
	Opt  *Options
	conn *websocket.Conn
	req  *request
}

// connect initializes the WebSocket connection.
func (c *Client) connect(ctx context.Context) error {
	conn, _, err := websocket.DefaultDialer.DialContext(ctx, c.Opt.Endpoint, nil)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Client) close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) combined(combine bool) {
	if combine {
		c.Opt.Endpoint = c.Opt.Endpoint + "/stream?streams="
	} else {
		c.Opt.Endpoint = c.Opt.Endpoint + "/ws"
	}
}

func (c *Client) keepAlive() {
	ticker := time.NewTicker(WebsocketStreamsTimeout)

	lastResponse := time.Now()
	c.conn.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		return nil
	})

	go func() {
		defer ticker.Stop()
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.conn.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				return
			}
			<-ticker.C
			if time.Since(lastResponse) > WebsocketStreamsTimeout {
				return
			}
		}
	}()
}

func (c *Client) wsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	onMessage := make(chan []byte, 8)
	onError := make(chan error)

	go func() {
		defer close(onMessage)
		defer close(onError)

		if err := c.connect(ctx); err != nil {
			onError <- err
			return
		}
		defer c.conn.Close()
		c.keepAlive()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					onError <- err
					return
				}
				onMessage <- message
			}
		}
	}()
	return onMessage, onError
}

func (c *Client) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	onMessage := make(chan []byte, 8)
	onError := make(chan error)
	err := c.connect(ctx)
	go func() {
		defer close(onMessage)
		defer close(onError)
		if err != nil {
			onError <- err
			return
		}
		defer c.conn.Close()
		c.keepAlive()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					onError <- err
					return
				}
				onMessage <- message
			}
		}
	}()
	return onMessage, onError
}

func uuid4() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]byte, 16)
	if _, err := r.Read(data); err != nil {
		return ""
	}
	data[6] = (data[6] & 0x0f) | 0x40
	data[8] = (data[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		data[0:4], data[4:6], data[6:8], data[8:10], data[10:])
}

func (c *Client) send() error {
	c.req.Id = uuid4()
	if c.conn == nil {
		return errors.New("websocket connection is nil")
	}
	if c.req.AuthType == authSigned {
		c.req.Params["timestamp"] = time.Now().UnixMilli()
	}
	if c.req.AuthType == authApiKey || c.req.AuthType == authSigned {
		c.req.Params["apiKey"] = c.Opt.ApiKey
	}
	if c.req.AuthType == authSigned {
		mac := hmac.New(sha256.New, []byte(c.Opt.ApiSecret))
		if _, err := mac.Write([]byte(SortMap(c.req.Params))); err != nil {
			return err
		}
		c.req.Params["signature"] = hex.EncodeToString(mac.Sum(nil))
	}
	return c.conn.WriteJSON(c.req)
}

func SortMap(params map[string]any) string {
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sortedParams []string
	for _, key := range keys {
		finalVar := ""
		value := params[key]
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			if elems, ok := value.([]string); ok {
				finalVar = `["` + strings.Join(elems, `","`) + `"]`
			} else {
				finalVar = fmt.Sprintf("%v", value)
			}
		default:
			finalVar = fmt.Sprintf("%v", value)
		}
		sortedParams = append(sortedParams, key+"="+finalVar)
	}
	return strings.Join(sortedParams, "&")
}

func (c *Client) NewWebsocketStreams() *WebsocketStreams {
	return &WebsocketStreams{c: c}
}

// NewPing Test connectivity
func (c *Client) NewPing() *Ping {
	c.req = &request{Method: "ping"}
	return &Ping{c: c}
}

// NewCheckServerTime Check server time
func (c *Client) NewCheckServerTime() *CheckServerTime {
	c.req = &request{Method: "time"}
	return &CheckServerTime{c: c}
}

// NewExchangeInfo Exchange information
func (c *Client) NewExchangeInfo() *ExchangeInfo {
	c.req = &request{Method: "exchangeInfo", Params: make(map[string]any)}
	return &ExchangeInfo{c: c}
}

// NewDepth Order book
func (c *Client) NewDepth() *Depth {
	c.req = &request{Method: "depth", Params: make(map[string]any)}
	return &Depth{c: c}
}

// NewTradesRecent Recent trades
func (c *Client) NewTradesRecent() *TradesRecent {
	c.req = &request{Method: "trades.recent", Params: make(map[string]any)}
	return &TradesRecent{c: c}
}

// NewTradesHistorical Historical trades
func (c *Client) NewTradesHistorical() *TradesHistorical {
	c.req = &request{Method: "trades.historical", Params: make(map[string]any)}
	return &TradesHistorical{c: c}
}

// NewTradesAggregate Aggregate trades
func (c *Client) NewTradesAggregate() *TradesAggregate {
	c.req = &request{Method: "trades.aggregate", Params: make(map[string]any)}
	return &TradesAggregate{c: c}
}

// NewKline Klines
func (c *Client) NewKline() *Kline {
	c.req = &request{Method: "klines", Params: make(map[string]any)}
	return &Kline{c: c}
}

// NewUiKlines UI Klines
func (c *Client) NewUiKlines() *UiKlines {
	c.req = &request{Method: "uiKlines", Params: make(map[string]any)}
	return &UiKlines{c: c}
}

// NewAveragePrice Current average price
func (c *Client) NewAveragePrice() *AveragePrice {
	c.req = &request{Method: "avgPrice", Params: make(map[string]any)}
	return &AveragePrice{c: c}
}

// NewTicker24h 24hr ticker price change statistics
func (c *Client) NewTicker24h() *Ticker24h {
	c.req = &request{Method: "ticker.24hr", Params: make(map[string]any)}
	return &Ticker24h{c: c}
}

// NewTickerTradingDay Trading Day Ticker
func (c *Client) NewTickerTradingDay() *TickerTradingDay {
	c.req = &request{Method: "ticker.tradingDay", Params: make(map[string]any)}
	return &TickerTradingDay{c: c}
}

// NewTicker Rolling window price change statistics
func (c *Client) NewTicker() *Ticker {
	c.req = &request{Method: "ticker", Params: make(map[string]any)}
	return &Ticker{c: c}
}

// NewTickerPrice Symbol price ticker
func (c *Client) NewTickerPrice() *TickerPrice {
	c.req = &request{Method: "ticker.price", Params: make(map[string]any)}
	return &TickerPrice{c: c}
}

// NewTickerBook Symbol order book ticker
func (c *Client) NewTickerBook() *TickerBook {
	c.req = &request{Method: "ticker.book", Params: make(map[string]any)}
	return &TickerBook{c: c}
}

// NewCreateOrder Place new order (TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	c.req = &request{Method: "order.place", Params: make(map[string]any), AuthType: authSigned}
	return &CreateOrder{c: c}
}

// NewCreateTestOrder Test new order (TRADE)
func (c *Client) NewCreateTestOrder() *CreateTestOrder {
	c.req = &request{Method: "order.test", Params: make(map[string]any), AuthType: authSigned}
	return &CreateTestOrder{c: c}
}

// NewQueryOrder Query order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.req = &request{Method: "order.status", Params: make(map[string]any), AuthType: authSigned}
	return &QueryOrder{c: c}
}

// NewCancelOrder Cancel order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.req = &request{Method: "order.cancel", Params: make(map[string]any), AuthType: authSigned}
	return &CancelOrder{c: c}
}

// NewCancelReplaceOrder Cancel and replace order (TRADE)
func (c *Client) NewCancelReplaceOrder() *CancelReplaceOrder {
	c.req = &request{Method: "order.cancelReplace", Params: make(map[string]any), AuthType: authSigned}
	return &CancelReplaceOrder{c: c}
}

// NewOpenOrdersStatus Current open orders (USER_DATA)
func (c *Client) NewOpenOrdersStatus() *OpenOrdersStatus {
	c.req = &request{Method: "openOrders.status", Params: make(map[string]any), AuthType: authSigned}
	return &OpenOrdersStatus{c: c}
}

// NewCancelOpenOrder Cancel open orders (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	c.req = &request{Method: "openOrders.cancelAll", Params: make(map[string]any), AuthType: authSigned}
	return &CancelOpenOrder{c: c}
}

// NewCreateOCOOrder Place new Order list - OCO (TRADE)
func (c *Client) NewCreateOCOOrder() *CreateOCOOrder {
	c.req = &request{Method: "orderList.place.oco", Params: make(map[string]any), AuthType: authSigned}
	return &CreateOCOOrder{c: c}
}

// NewCreateOTOOrder Place new Order list - OTO (TRADE)
func (c *Client) NewCreateOTOOrder() *CreateOTOOrder {
	c.req = &request{Method: "orderList.place.oto", Params: make(map[string]any), AuthType: authSigned}
	return &CreateOTOOrder{c: c}
}

// NewCreateOTOCOOrder Place new Order list - OTOCO (TRADE)
func (c *Client) NewCreateOTOCOOrder() *CreateOTOCOOrder {
	c.req = &request{Method: "orderList.place.otoco", Params: make(map[string]any), AuthType: authSigned}
	return &CreateOTOCOOrder{c: c}
}

// NewQueryOrderList Query Order list (USER_DATA)
func (c *Client) NewQueryOrderList() *QueryOrderList {
	c.req = &request{Method: "orderList.status", Params: make(map[string]any), AuthType: authSigned}
	return &QueryOrderList{c: c}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *Client) NewCancelOrderList() *CancelOrderList {
	c.req = &request{Method: "orderList.cancel", Params: make(map[string]any), AuthType: authSigned}
	return &CancelOrderList{c: c}
}

// NewQueryOpenOrder Current open Order lists (USER_DATA)
func (c *Client) NewQueryOpenOrder() *QueryOpenOrder {
	c.req = &request{Method: "openOrderLists.status", Params: make(map[string]any), AuthType: authSigned}
	return &QueryOpenOrder{c: c}
}

// NewCreateSOROrder Place new order using SOR (TRADE)
func (c *Client) NewCreateSOROrder() *CreateSOROrder {
	c.req = &request{Method: "sor.order.place", Params: make(map[string]any), AuthType: authSigned}
	return &CreateSOROrder{c: c}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *Client) NewCreateTestSOROrder() *CreateTestSOROrder {
	c.req = &request{Method: "sor.order.test", Params: make(map[string]any), AuthType: authSigned}
	return &CreateTestSOROrder{c: c}
}

// NewAccountInformation Account information (USER_DATA)
func (c *Client) NewAccountInformation() *AccountInformation {
	c.req = &request{Method: "account.status", Params: make(map[string]any), AuthType: authSigned}
	return &AccountInformation{c: c}
}

// NewUnfilledOrder Unfilled Order Count (USER_DATA)
func (c *Client) NewUnfilledOrder() *UnfilledOrder {
	c.req = &request{Method: "account.rateLimits.orders", Params: make(map[string]any), AuthType: authSigned}
	return &UnfilledOrder{c: c}
}

// NewAccountOrderHistory Account order history (USER_DATA)
func (c *Client) NewAccountOrderHistory() *AccountOrderHistory {
	c.req = &request{Method: "allOrders", Params: make(map[string]any), AuthType: authSigned}
	return &AccountOrderHistory{c: c}
}

// NewAllOrderList Account Order list history (USER_DATA)
func (c *Client) NewAllOrderList() *AllOrderList {
	c.req = &request{Method: "allOrderLists", Params: make(map[string]any), AuthType: authSigned}
	return &AllOrderList{c: c}
}

// NewAccountTradeHistory Account trade history (USER_DATA)
func (c *Client) NewAccountTradeHistory() *AccountTradeHistory {
	c.req = &request{Method: "myTrades", Params: make(map[string]any), AuthType: authSigned}
	return &AccountTradeHistory{c: c}
}

// NewAccountPreventedMatches Account prevented matches (USER_DATA)
func (c *Client) NewAccountPreventedMatches() *AccountPreventedMatches {
	c.req = &request{Method: "myPreventedMatches", Params: make(map[string]any), AuthType: authSigned}
	return &AccountPreventedMatches{c: c}
}

// NewAccountAllocations Account allocations (USER_DATA)
func (c *Client) NewAccountAllocations() *AccountAllocations {
	c.req = &request{Method: "myAllocations", Params: make(map[string]any), AuthType: authSigned}
	return &AccountAllocations{c: c}
}

// NewAccountCommission Account Commission Rates (USER_DATA)
func (c *Client) NewAccountCommission() *AccountCommission {
	c.req = &request{Method: "account.commission", Params: make(map[string]any), AuthType: authSigned}
	return &AccountCommission{c: c}
}

// NewStartUserDataStream Start user data stream (USER_STREAM)
func (c *Client) NewStartUserDataStream() *StartUserDataStream {
	c.req = &request{Method: "userDataStream.start", Params: make(map[string]any), AuthType: authApiKey}
	return &StartUserDataStream{c: c}
}

// NewPingUserDataStream Ping user data stream (USER_STREAM)
func (c *Client) NewPingUserDataStream() *PingUserDataStream {
	c.req = &request{Method: "userDataStream.ping", Params: make(map[string]any), AuthType: authApiKey}
	return &PingUserDataStream{c: c}
}

// NewStopUserDataStream Stop user data stream (USER_STREAM)
func (c *Client) NewStopUserDataStream() *StopUserDataStream {
	c.req = &request{Method: "userDataStream.stop", Params: make(map[string]any), AuthType: authApiKey}
	return &StopUserDataStream{c: c}
}
