package wss

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/jekaxv/go-binance/types"
	"github.com/jekaxv/go-binance/utils"
	"log/slog"
	"math/rand"
	"time"
)

type Client struct {
	Opt    *Options
	Logger *slog.Logger
	conn   *websocket.Conn
	req    *request
}

// connect initializes the WebSocket connection.
func (c *Client) connect(ctx context.Context) error {
	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, c.Opt.Endpoint, nil)
	if err != nil {
		c.Logger.Debug("websocket dial failed", "endpoint", c.Opt.Endpoint, "error", err)
		return err
	}
	c.Logger.Debug("websocket connection established", "endpoint", c.Opt.Endpoint, "status", resp.Status)
	c.conn = conn
	return nil
}
func (c *Client) SetReq(method string, aType ...types.AuthType) {
	reqType := types.AuthNone
	if len(aType) > 0 {
		reqType = aType[0]
	}
	c.req = &request{Method: method, AuthType: reqType, Params: make(map[string]any)}
}
func (c *Client) Close() error {
	return c.close()
}

func (c *Client) SetParams(key string, value any) {
	c.req.Params[key] = value
}

func (c *Client) close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *Client) Combined(combine bool) {
	c.combined(combine)
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
		c.Logger.Debug("received pong", "time", lastResponse.Format(time.RFC3339))
		return nil
	})

	go func() {
		defer ticker.Stop()
		c.Logger.Debug("websocket keepalive started", "timeout", WebsocketStreamsTimeout.String())
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.conn.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				c.Logger.Debug("failed to send ping", "error", err)
				return
			}
			c.Logger.Debug("ping sent", "deadline", deadline.Format(time.RFC3339))
			<-ticker.C
			if time.Since(lastResponse) > WebsocketStreamsTimeout {
				return
			}
		}
	}()
}

func (c *Client) WsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.wsServe(ctx)
}

func (c *Client) wsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	onMessage := make(chan []byte, 8)
	onError := make(chan error)

	go func() {
		defer func() {
			close(onMessage)
			close(onError)
			c.Logger.Debug("websocket serve goroutine exited")
		}()

		if err := c.connect(ctx); err != nil {
			c.Logger.Debug("websocket connect failed", "error", err)
			onError <- err
			return
		}
		defer c.conn.Close()
		c.keepAlive()
		for {
			select {
			case <-ctx.Done():
				c.Logger.Debug("context done, websocket serve stopping")
				return
			default:
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					c.Logger.Debug("failed to read message from websocket", "error", err)
					onError <- err
					return
				}
				c.Logger.Debug("websocket message received", "length", len(message))
				onMessage <- message
			}
		}
	}()
	return onMessage, onError
}

func (c *Client) WsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.wsApiServe(ctx)
}

func (c *Client) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	onMessage := make(chan []byte, 8)
	onError := make(chan error)
	err := c.connect(ctx)
	c.Logger.Debug("attempting websocket connection", "endpoint", c.Opt.Endpoint)
	go func() {
		defer func() {
			close(onMessage)
			close(onError)
			c.Logger.Debug("wsApiServe goroutine exited")
		}()
		if err != nil {
			c.Logger.Debug("websocket connect failed", "error", err)
			onError <- err
			return
		}
		defer c.conn.Close()
		c.keepAlive()
		for {
			select {
			case <-ctx.Done():
				c.Logger.Debug("context done, exiting wsApiServe loop")
				return
			default:
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					c.Logger.Debug("error reading websocket message", "error", err)
					onError <- err
					return
				}
				c.Logger.Debug("received websocket message", "length", len(message))
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

func (c *Client) Send() error {
	return c.send()
}

func (c *Client) send() error {
	if c.conn == nil {
		c.Logger.Debug("cannot send: connection is nil")
		return errors.New("websocket connection is nil")
	}
	c.req.Id = uuid4()
	c.Logger.Debug("generating request ID", "id", c.req.Id)
	if c.req.AuthType == types.AuthSigned {
		c.req.Params["timestamp"] = time.Now().UnixMilli()
	}
	if c.req.AuthType == types.AuthApiKey || c.req.AuthType == types.AuthSigned {
		c.req.Params["apiKey"] = c.Opt.ApiKey
	}

	if c.req.AuthType == types.AuthSigned {
		if c.Opt.SignType == types.SignTypeRsa {
			c.req.SignFunc = utils.RsaSign
		} else if c.Opt.SignType == types.SignTypeEd25519 {
			c.req.SignFunc = utils.Ed25519Sign
		} else {
			c.req.SignFunc = utils.HmacSign
		}
		sortedData := utils.SortMap(c.req.Params)
		c.Logger.Debug("sorted params for signature", "params", sortedData)
		sign, err := c.req.SignFunc(c.Opt.ApiSecret, sortedData)
		if err != nil {
			c.Logger.Debug("signature generation failed", "error", err)
			return err
		}
		c.req.Params["signature"] = sign
		c.Logger.Debug("signature added to request", "signature", sign)
	}
	return c.conn.WriteJSON(c.req)
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
	c.req = &request{Method: "order.place", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateOrder{c: c}
}

// NewCreateTestOrder Test new order (TRADE)
func (c *Client) NewCreateTestOrder() *CreateTestOrder {
	c.req = &request{Method: "order.test", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateTestOrder{c: c}
}

// NewQueryOrder Query order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.req = &request{Method: "order.status", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &QueryOrder{c: c}
}

// NewCancelOrder Cancel order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.req = &request{Method: "order.cancel", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CancelOrder{c: c}
}

// NewCancelReplaceOrder Cancel and replace order (TRADE)
func (c *Client) NewCancelReplaceOrder() *CancelReplaceOrder {
	c.req = &request{Method: "order.cancelReplace", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CancelReplaceOrder{c: c}
}

// NewOpenOrdersStatus Current open orders (USER_DATA)
func (c *Client) NewOpenOrdersStatus() *OpenOrdersStatus {
	c.req = &request{Method: "openOrders.status", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &OpenOrdersStatus{c: c}
}

// NewCancelOpenOrder Cancel open orders (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	c.req = &request{Method: "openOrders.cancelAll", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CancelOpenOrder{c: c}
}

// NewCreateOCOOrder Place new Order list - OCO (TRADE)
func (c *Client) NewCreateOCOOrder() *CreateOCOOrder {
	c.req = &request{Method: "orderList.place.oco", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateOCOOrder{c: c}
}

// NewCreateOTOOrder Place new Order list - OTO (TRADE)
func (c *Client) NewCreateOTOOrder() *CreateOTOOrder {
	c.req = &request{Method: "orderList.place.oto", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateOTOOrder{c: c}
}

// NewCreateOTOCOOrder Place new Order list - OTOCO (TRADE)
func (c *Client) NewCreateOTOCOOrder() *CreateOTOCOOrder {
	c.req = &request{Method: "orderList.place.otoco", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateOTOCOOrder{c: c}
}

// NewQueryOrderList Query Order list (USER_DATA)
func (c *Client) NewQueryOrderList() *QueryOrderList {
	c.req = &request{Method: "orderList.status", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &QueryOrderList{c: c}
}

// NewCancelOrderList Cancel Order list (TRADE)
func (c *Client) NewCancelOrderList() *CancelOrderList {
	c.req = &request{Method: "orderList.cancel", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CancelOrderList{c: c}
}

// NewQueryOpenOrder Current open Order lists (USER_DATA)
func (c *Client) NewQueryOpenOrder() *QueryOpenOrder {
	c.req = &request{Method: "openOrderLists.status", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &QueryOpenOrder{c: c}
}

// NewCreateSOROrder Place new order using SOR (TRADE)
func (c *Client) NewCreateSOROrder() *CreateSOROrder {
	c.req = &request{Method: "sor.order.place", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateSOROrder{c: c}
}

// NewCreateTestSOROrder Test new order using SOR (TRADE)
func (c *Client) NewCreateTestSOROrder() *CreateTestSOROrder {
	c.req = &request{Method: "sor.order.test", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &CreateTestSOROrder{c: c}
}

// NewAccountInformation Account information (USER_DATA)
func (c *Client) NewAccountInformation() *AccountInformation {
	c.req = &request{Method: "account.status", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AccountInformation{c: c}
}

// NewUnfilledOrder Unfilled Order Count (USER_DATA)
func (c *Client) NewUnfilledOrder() *UnfilledOrder {
	c.req = &request{Method: "account.rateLimits.orders", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &UnfilledOrder{c: c}
}

// NewAccountOrderHistory Account order history (USER_DATA)
func (c *Client) NewAccountOrderHistory() *AccountOrderHistory {
	c.req = &request{Method: "allOrders", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AccountOrderHistory{c: c}
}

// NewAllOrderList Account Order list history (USER_DATA)
func (c *Client) NewAllOrderList() *AllOrderList {
	c.req = &request{Method: "allOrderLists", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AllOrderList{c: c}
}

// NewAccountTradeHistory Account trade history (USER_DATA)
func (c *Client) NewAccountTradeHistory() *AccountTradeHistory {
	c.req = &request{Method: "myTrades", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AccountTradeHistory{c: c}
}

// NewAccountPreventedMatches Account prevented matches (USER_DATA)
func (c *Client) NewAccountPreventedMatches() *AccountPreventedMatches {
	c.req = &request{Method: "myPreventedMatches", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AccountPreventedMatches{c: c}
}

// NewAccountAllocations Account allocations (USER_DATA)
func (c *Client) NewAccountAllocations() *AccountAllocations {
	c.req = &request{Method: "myAllocations", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AccountAllocations{c: c}
}

// NewAccountCommission Account Commission Rates (USER_DATA)
func (c *Client) NewAccountCommission() *AccountCommission {
	c.req = &request{Method: "account.commission", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &AccountCommission{c: c}
}

// NewSessionLogon Log in with API key (SIGNED)
func (c *Client) NewSessionLogon() *SessionLogon {
	c.req = &request{Method: "session.logon", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &SessionLogon{c: c}
}

// NewSessionStatus Query session status (SIGNED)
func (c *Client) NewSessionStatus() *SessionStatus {
	c.req = &request{Method: "session.status", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &SessionStatus{c: c}
}

// NewSessionLogout Log out of the session
func (c *Client) NewSessionLogout() *SessionLogout {
	c.req = &request{Method: "session.logout", Params: make(map[string]any), AuthType: types.AuthSigned}
	return &SessionLogout{c: c}
}
