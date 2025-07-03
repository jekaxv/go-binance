package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// CreateOrder Send in a new order.
type CreateOrder struct {
	c *Client
	r *core.Request
}

type CreateOrderResponse struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int             `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
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
	WorkingTime             int64           `json:"workingTime"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	Fills                   []*Fill         `json:"fills"`
}

func (s *CreateOrder) Symbol(symbol string) *CreateOrder {
	s.r.Set("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *CreateOrder) Side(side core.OrderSideEnum) *CreateOrder {
	s.r.Set("side", side)
	return s
}

func (s *CreateOrder) Type(orderType core.OrderTypeEnum) *CreateOrder {
	s.r.Set("type", orderType)
	return s
}

func (s *CreateOrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateOrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}

func (s *CreateOrder) Quantity(quantity string) *CreateOrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *CreateOrder) QuoteOrderQty(quoteOrderQty string) *CreateOrder {
	s.r.Set("quoteOrderQty", quoteOrderQty)
	return s
}

func (s *CreateOrder) Price(price string) *CreateOrder {
	s.r.Set("price", price)
	return s
}

// NewClientOrderId A unique id among open orders. Automatically generated if not sent.
// Orders with the same newClientOrderID can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (s *CreateOrder) NewClientOrderId(newClientOrderId string) *CreateOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

func (s *CreateOrder) StrategyId(strategyId int64) *CreateOrder {
	s.r.Set("strategyId", strategyId)
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *CreateOrder) StrategyType(strategyType int) *CreateOrder {
	s.r.Set("strategyType", strategyType)
	return s
}

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateOrder) StopPrice(stopPrice string) *CreateOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}

// TrailingDelta Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateOrder) TrailingDelta(trailingDelta int64) *CreateOrder {
	s.r.Set("trailingDelta", trailingDelta)
	return s
}

// IcebergQty Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
func (s *CreateOrder) IcebergQty(icebergQty string) *CreateOrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *CreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *CreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateOrder) RecvWindow(recvWindow int) *CreateOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CreateOrder) Do(ctx context.Context) (*CreateOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(CreateOrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

type TestCreateOrder struct {
	c *Client
	r *core.Request
}

type OrderCommission struct {
	Maker decimal.Decimal `json:"maker"`
	Taker decimal.Decimal `json:"taker"`
}

type Discount struct {
	EnabledForAccount bool            `json:"enabledForAccount"`
	EnabledForSymbol  bool            `json:"enabledForSymbol"`
	DiscountAsset     string          `json:"discountAsset"`
	Discount          decimal.Decimal `json:"discount"`
}

type TestCreateOrderResponse struct {
	StandardCommissionForOrder OrderCommission `json:"standardCommissionForOrder"`
	TaxCommissionForOrder      OrderCommission `json:"taxCommissionForOrder"`
	Discount                   Discount        `json:"discount"`
}

func (s *TestCreateOrder) Symbol(symbol string) *TestCreateOrder {
	s.r.Set("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *TestCreateOrder) Side(side core.OrderSideEnum) *TestCreateOrder {
	s.r.Set("side", side)
	return s
}

// Type LIMIT,MARKET,STOP_LOSS,STOP_LOSS_LIMIT,TAKE_PROFIT,TAKE_PROFIT_LIMIT,LIMIT_MAKER
func (s *TestCreateOrder) Type(orderType core.OrderTypeEnum) *TestCreateOrder {
	s.r.Set("type", orderType)
	return s
}

func (s *TestCreateOrder) TimeInForce(timeInForce core.TimeInForceEnum) *TestCreateOrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}

func (s *TestCreateOrder) Quantity(quantity string) *TestCreateOrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *TestCreateOrder) QuoteOrderQty(quoteOrderQty string) *TestCreateOrder {
	s.r.Set("quoteOrderQty", quoteOrderQty)
	return s
}

func (s *TestCreateOrder) Price(price string) *TestCreateOrder {
	s.r.Set("price", price)
	return s
}

// NewClientOrderId A unique id among open orders. Automatically generated if not sent.
// Orders with the same newClientOrderID can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (s *TestCreateOrder) NewClientOrderId(newClientOrderId string) *TestCreateOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

func (s *TestCreateOrder) StrategyId(strategyId int64) *TestCreateOrder {
	s.r.Set("strategyId", strategyId)
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *TestCreateOrder) StrategyType(strategyType int) *TestCreateOrder {
	s.r.Set("strategyType", strategyType)
	return s
}

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *TestCreateOrder) StopPrice(stopPrice string) *TestCreateOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}

// TrailingDelta Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *TestCreateOrder) TrailingDelta(trailingDelta int64) *TestCreateOrder {
	s.r.Set("trailingDelta", trailingDelta)
	return s
}

// IcebergQty Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
func (s *TestCreateOrder) IcebergQty(icebergQty string) *TestCreateOrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *TestCreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *TestCreateOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *TestCreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *TestCreateOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *TestCreateOrder) RecvWindow(recvWindow int) *TestCreateOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

// ComputeCommissionRates Default: false
func (s *TestCreateOrder) ComputeCommissionRates(computeCommissionRates bool) *TestCreateOrder {
	s.r.Set("computeCommissionRates", computeCommissionRates)
	return s
}

func (s *TestCreateOrder) Do(ctx context.Context) (*TestCreateOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	if s.r.GetQuery("computeCommissionRates") == "true" {
		resp := new(TestCreateOrderResponse)
		return resp, json.Unmarshal(s.c.rawBody(), resp)
	}
	return nil, nil
}

// QueryOrder Check an order's status.
type QueryOrder struct {
	c *Client
	r *core.Request
}

type QueryOrderResponse struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int             `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	IcebergQty              decimal.Decimal `json:"icebergQty"`
	Time                    int64           `json:"time"`
	UpdateTime              int64           `json:"updateTime"`
	IsWorking               bool            `json:"isWorking"`
	WorkingTime             int64           `json:"workingTime"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
}

func (s *QueryOrder) Symbol(symbol string) *QueryOrder {
	s.r.Set("symbol", symbol)
	return s
}

// OrderId For some historical orders cummulativeQuoteQty will be < 0, meaning the data is not available at this time.
func (s *QueryOrder) OrderId(orderId int64) *QueryOrder {
	s.r.Set("orderId", orderId)
	return s
}

// OrigClientOrderId For some historical orders cummulativeQuoteQty will be < 0, meaning the data is not available at this time.
func (s *QueryOrder) OrigClientOrderId(origClientOrderId string) *QueryOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

func (s *QueryOrder) RecvWindow(recvWindow int) *QueryOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *QueryOrder) Do(ctx context.Context) (*QueryOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp *QueryOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOrder Cancel an active order.
type CancelOrder struct {
	c *Client
	r *core.Request
}

func (s *CancelOrder) Symbol(symbol string) *CancelOrder {
	s.r.Set("symbol", symbol)
	return s
}

// OrderId Either orderId or origClientOrderId must be sent. If both parameters are sent, orderId takes precedence.
func (s *CancelOrder) OrderId(orderId int64) *CancelOrder {
	s.r.Set("orderId", orderId)
	return s
}

// OrigClientOrderId Either orderId or origClientOrderId must be sent. If both parameters are sent, orderId takes precedence.
func (s *CancelOrder) OrigClientOrderId(origClientOrderId string) *CancelOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

// NewClientOrderId Used to uniquely identify this cancel. Automatically generated by default.
func (s *CancelOrder) NewClientOrderId(newClientOrderId string) *CancelOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

func (s *CancelOrder) CancelRestrictions(cancelRestrictions core.CancelRestrictionEnum) *CancelOrder {
	s.r.Set("cancelRestrictions", cancelRestrictions)
	return s
}

func (s *CancelOrder) RecvWindow(recvWindow int) *CancelOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CancelOrder) Do(ctx context.Context) (*QueryOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(QueryOrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// CancelOpenOrder Cancels all active orders on a symbol. This includes orders that are part of an order list.
type CancelOpenOrder struct {
	c *Client
	r *core.Request
}

type CancelOpenOrderResponse struct {
	Symbol                  string          `json:"symbol"`
	OrigClientOrderId       string          `json:"origClientOrderId,omitempty"`
	OrderId                 int             `json:"orderId,omitempty"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId,omitempty"`
	TransactTime            int64           `json:"transactTime,omitempty"`
	Price                   decimal.Decimal `json:"price,omitempty"`
	OrigQty                 decimal.Decimal `json:"origQty,omitempty"`
	ExecutedQty             decimal.Decimal `json:"executedQty,omitempty"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty,omitempty"`
	Status                  string          `json:"status,omitempty"`
	TimeInForce             string          `json:"timeInForce,omitempty"`
	Type                    string          `json:"type,omitempty"`
	Side                    string          `json:"side,omitempty"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode,omitempty"`
	ContingencyType         string          `json:"contingencyType,omitempty"`
	ListStatusType          string          `json:"listStatusType,omitempty"`
	ListOrderStatus         string          `json:"listOrderStatus,omitempty"`
	ListClientOrderId       string          `json:"listClientOrderId,omitempty"`
	TransactionTime         int64           `json:"transactionTime,omitempty"`
	Orders                  []*Order        `json:"orders,omitempty"`
	OrderReports            []*OrderReport  `json:"orderReports,omitempty"`
}

func (s *CancelOpenOrder) Symbol(symbol string) *CancelOpenOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *CancelOpenOrder) RecvWindow(recvWindow int) *CancelOpenOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CancelOpenOrder) Do(ctx context.Context) ([]*CancelOpenOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*CancelOpenOrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelReplace Cancel an Existing Order and Send a New Order (TRADE)
// Cancels an existing order and places a new order on the same symbol.
// Filters and Order Count are evaluated before the processing of the cancellation and order placement occurs.
// A new order that was not attempted (i.e. when newOrderResult: NOT_ATTEMPTED ), will still increase the order count by 1.
type CancelReplace struct {
	c *Client
	r *core.Request
}

type CancelReplaceResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		NewOrderResponse *OrderReport `json:"newOrderResponse"`
		CancelResult     string       `json:"cancelResult"`
		NewOrderResult   string       `json:"newOrderResult"`
		CancelResponse   *OrderReport `json:"cancelResponse"`
	} `json:"data"`
	CancelResult     string       `json:"cancelResult"`
	NewOrderResult   string       `json:"newOrderResult"`
	CancelResponse   *OrderReport `json:"cancelResponse"`
	NewOrderResponse *OrderReport `json:"newOrderResponse"`
}

func (s *CancelReplace) Symbol(symbol string) *CancelReplace {
	s.r.Set("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *CancelReplace) Side(side core.OrderSideEnum) *CancelReplace {
	s.r.Set("side", side)
	return s
}

// Type LIMIT,MARKET,STOP_LOSS,STOP_LOSS_LIMIT,TAKE_PROFIT,TAKE_PROFIT_LIMIT,LIMIT_MAKER
func (s *CancelReplace) Type(orderType core.OrderTypeEnum) *CancelReplace {
	s.r.Set("type", orderType)
	return s
}

// CancelReplaceMode The allowed values are:
// STOP_ON_FAILURE - If the cancel request fails, the new order placement will not be attempted.
// ALLOW_FAILURE - new order placement will be attempted even if cancel request fails.
func (s *CancelReplace) CancelReplaceMode(cancelReplaceMode core.CancelReplaceModeEnum) *CancelReplace {
	s.r.Set("cancelReplaceMode", cancelReplaceMode)
	return s
}

func (s *CancelReplace) TimeInForce(timeInForce core.TimeInForceEnum) *CancelReplace {
	s.r.Set("timeInForce", timeInForce)
	return s
}

func (s *CancelReplace) Quantity(quantity string) *CancelReplace {
	s.r.Set("quantity", quantity)
	return s
}

func (s *CancelReplace) QuoteOrderQty(quoteOrderQty string) *CancelReplace {
	s.r.Set("quoteOrderQty", quoteOrderQty)
	return s
}

func (s *CancelReplace) Price(price string) *CancelReplace {
	s.r.Set("price", price)
	return s
}

func (s *CancelReplace) CancelNewClientOrderId(cancelNewClientOrderId string) *CancelReplace {
	s.r.Set("cancelNewClientOrderId", cancelNewClientOrderId)
	return s
}

func (s *CancelReplace) CancelOrigClientOrderId(cancelOrigClientOrderId string) *CancelReplace {
	s.r.Set("cancelOrigClientOrderId", cancelOrigClientOrderId)
	return s
}

func (s *CancelReplace) CancelOrderId(cancelOrderId int64) *CancelReplace {
	s.r.Set("cancelOrderId", cancelOrderId)
	return s
}

func (s *CancelReplace) NewClientOrderId(newClientOrderId string) *CancelReplace {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

func (s *CancelReplace) StrategyId(strategyId int64) *CancelReplace {
	s.r.Set("strategyId", strategyId)
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *CancelReplace) StrategyType(strategyType int) *CancelReplace {
	s.r.Set("strategyType", strategyType)
	return s
}

func (s *CancelReplace) StopPrice(stopPrice string) *CancelReplace {
	s.r.Set("stopPrice", stopPrice)
	return s
}

func (s *CancelReplace) TrailingDelta(trailingDelta int64) *CancelReplace {
	s.r.Set("trailingDelta", trailingDelta)
	return s
}

func (s *CancelReplace) IcebergQty(icebergQty string) *CancelReplace {
	s.r.Set("icebergQty", icebergQty)
	return s
}

// NewOrderRespType Allowed values:
// ACK, RESULT, FULL
// MARKET and LIMIT orders types default to FULL; all other orders default to ACK
func (s *CancelReplace) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CancelReplace {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol.
func (s *CancelReplace) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CancelReplace {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *CancelReplace) CancelRestrictions(cancelRestrictions core.CancelRestrictionEnum) *CancelReplace {
	s.r.Set("cancelRestrictions", cancelRestrictions)
	return s
}

func (s *CancelReplace) OrderRateLimitExceededMode(orderRateLimitExceededMode core.OrderExceededModeEnum) *CancelReplace {
	s.r.Set("orderRateLimitExceededMode", orderRateLimitExceededMode)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CancelReplace) RecvWindow(recvWindow int) *CancelReplace {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CancelReplace) Do(ctx context.Context) (*CancelReplaceResponse, error) {
	resp := new(CancelReplaceResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OpenOrders Get all open orders on a symbol. Careful when accessing this with no symbol.
type OpenOrders struct {
	c *Client
	r *core.Request
}

type OrdersResponse struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int             `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	IcebergQty              decimal.Decimal `json:"icebergQty"`
	Time                    int64           `json:"time"`
	UpdateTime              int64           `json:"updateTime"`
	IsWorking               bool            `json:"isWorking"`
	WorkingTime             int64           `json:"workingTime"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
}

func (s *OpenOrders) Symbol(symbol string) *OpenOrders {
	s.r.Set("symbol", symbol)
	return s
}
func (s *OpenOrders) RecvWindow(recvWindow int) *OpenOrders {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *OpenOrders) Do(ctx context.Context) ([]*OrdersResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OrdersResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AllOrders Get all account orders; active, canceled, or filled.
type AllOrders struct {
	c *Client
	r *core.Request
}

func (s *AllOrders) Symbol(symbol string) *AllOrders {
	s.r.Set("symbol", symbol)
	return s
}

func (s *AllOrders) OrderId(orderId int64) *AllOrders {
	s.r.Set("orderId", orderId)
	return s
}

func (s *AllOrders) StartTime(startTime int64) *AllOrders {
	s.r.Set("startTime", startTime)
	return s
}

func (s *AllOrders) EndTime(endTime int64) *AllOrders {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500; max 1000.
func (s *AllOrders) Limit(limit int) *AllOrders {
	s.r.Set("limit", limit)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *AllOrders) RecvWindow(recvWindow int) *AllOrders {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *AllOrders) Do(ctx context.Context) ([]*OrdersResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OrdersResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOrderList Cancel an entire Order list
type CancelOrderList struct {
	c *Client
	r *core.Request
}

type OrderListResponse struct {
	OrderListId       int            `json:"orderListId"`
	ContingencyType   string         `json:"contingencyType"`
	ListStatusType    string         `json:"listStatusType"`
	ListOrderStatus   string         `json:"listOrderStatus"`
	ListClientOrderId string         `json:"listClientOrderId"`
	TransactionTime   int64          `json:"transactionTime"`
	Symbol            string         `json:"symbol"`
	Orders            []*Order       `json:"orders"`
	OrderReports      []*OrderReport `json:"orderReports"`
}

func (s *CancelOrderList) Symbol(symbol string) *CancelOrderList {
	s.r.Set("symbol", symbol)
	return s
}

// OrderListId Either orderListId or listClientOrderId must be provided
func (s *CancelOrderList) OrderListId(orderListId int64) *CancelOrderList {
	s.r.Set("orderListId", orderListId)
	return s
}

// ListClientOrderId Either orderListId or listClientOrderId must be provided
func (s *CancelOrderList) ListClientOrderId(listClientOrderId string) *CancelOrderList {
	s.r.Set("listClientOrderId", listClientOrderId)
	return s
}

// NewClientOrderId Used to uniquely identify this cancel. Automatically generated by default
func (s *CancelOrderList) NewClientOrderId(newClientOrderId string) *CancelOrderList {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CancelOrderList) RecvWindow(recvWindow int) *CancelOrderList {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CancelOrderList) Do(ctx context.Context) (*OrderListResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OrderListResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// QueryOrderList Retrieves a specific order list based on provided optional parameters.
type QueryOrderList struct {
	c *Client
	r *core.Request
}

type QueryOrderListResponse struct {
	OrderListId       int      `json:"orderListId"`
	ContingencyType   string   `json:"contingencyType"`
	ListStatusType    string   `json:"listStatusType"`
	ListOrderStatus   string   `json:"listOrderStatus"`
	ListClientOrderId string   `json:"listClientOrderId"`
	TransactionTime   int64    `json:"transactionTime"`
	Symbol            string   `json:"symbol"`
	Orders            []*Order `json:"orders"`
}

// OrderListId Either orderListId or listClientOrderId must be provided
func (s *QueryOrderList) OrderListId(orderListId int64) *QueryOrderList {
	s.r.Set("orderListId", orderListId)
	return s
}

// OrigClientOrderId Either orderListId or listClientOrderId must be provided
func (s *QueryOrderList) OrigClientOrderId(origClientOrderId string) *QueryOrderList {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *QueryOrderList) RecvWindow(recvWindow int) *QueryOrderList {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *QueryOrderList) Do(ctx context.Context) (*QueryOrderListResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(QueryOrderListResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// QueryAllOrderLists Retrieves all order lists based on provided optional parameters.
// Note that the time between startTime and endTime can't be longer than 24 hours.
type QueryAllOrderLists struct {
	c *Client
	r *core.Request
}
type QueryAllOrderListsResponse struct {
	OrderListId       int      `json:"orderListId"`
	ContingencyType   string   `json:"contingencyType"`
	ListStatusType    string   `json:"listStatusType"`
	ListOrderStatus   string   `json:"listOrderStatus"`
	ListClientOrderId string   `json:"listClientOrderId"`
	TransactionTime   int64    `json:"transactionTime"`
	Symbol            string   `json:"symbol"`
	Orders            []*Order `json:"orders"`
}

// FromId If supplied, neither startTime or endTime can be provided
func (s *QueryAllOrderLists) FromId(fromId int64) *QueryAllOrderLists {
	s.r.Set("fromId", fromId)
	return s
}
func (s *QueryAllOrderLists) StartTime(startTime int64) *QueryAllOrderLists {
	s.r.Set("startTime", startTime)
	return s
}
func (s *QueryAllOrderLists) EndTime(endTime int64) *QueryAllOrderLists {
	s.r.Set("endTime", endTime)
	return s
}

// Limit The default is 500; max is 1000
func (s *QueryAllOrderLists) Limit(limit int) *QueryAllOrderLists {
	s.r.Set("limit", limit)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *QueryAllOrderLists) RecvWindow(recvWindow int) *QueryAllOrderLists {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *QueryAllOrderLists) Do(ctx context.Context) ([]*QueryAllOrderListsResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*QueryAllOrderListsResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryOpenOrderList Query Open Order lists
type QueryOpenOrderList struct {
	c *Client
	r *core.Request
}

type QueryOpenOrderListResponse struct {
	OrderListId       int      `json:"orderListId"`
	ContingencyType   string   `json:"contingencyType"`
	ListStatusType    string   `json:"listStatusType"`
	ListOrderStatus   string   `json:"listOrderStatus"`
	ListClientOrderId string   `json:"listClientOrderId"`
	TransactionTime   int64    `json:"transactionTime"`
	Symbol            string   `json:"symbol"`
	Orders            []*Order `json:"orders"`
}

func (s *QueryOpenOrderList) RecvWindow(recvWindow int) *QueryOpenOrderList {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *QueryOpenOrderList) Do(ctx context.Context) ([]*QueryOpenOrderListResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*QueryOpenOrderListResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CreateSOROrder Places an order using smart order routing (SOR).
type CreateSOROrder struct {
	c *Client
	r *core.Request
}
type CreateSOROrderResponse struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int             `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
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
	WorkingTime             int64           `json:"workingTime"`
	Fills                   []*Fill         `json:"fills"`
	WorkingFloor            string          `json:"workingFloor"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	UsedSor                 bool            `json:"usedSor"`
}

func (s *CreateSOROrder) Symbol(symbol string) *CreateSOROrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *CreateSOROrder) Side(side core.OrderSideEnum) *CreateSOROrder {
	s.r.Set("side", side)
	return s
}
func (s *CreateSOROrder) Type(orderType core.OrderTypeEnum) *CreateSOROrder {
	s.r.Set("type", orderType)
	return s
}
func (s *CreateSOROrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateSOROrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}
func (s *CreateSOROrder) Quantity(quantity string) *CreateSOROrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *CreateSOROrder) Price(price string) *CreateSOROrder {
	s.r.Set("price", price)
	return s
}
func (s *CreateSOROrder) NewClientOrderId(newClientOrderId string) *CreateSOROrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *CreateSOROrder) StrategyId(strategyId int64) *CreateSOROrder {
	s.r.Set("strategyId", strategyId)
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *CreateSOROrder) StrategyType(strategyType int) *CreateSOROrder {
	s.r.Set("strategyType", strategyType)
	return s
}

// IcebergQty Used with LIMIT to create an iceberg order.
func (s *CreateSOROrder) IcebergQty(icebergQty string) *CreateSOROrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}

// NewOrderRespType set the response JSON. ACK, RESULT, or FULL. Default to FULL
func (s *CreateSOROrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateSOROrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol.
func (s *CreateSOROrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateSOROrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateSOROrder) RecvWindow(recvWindow int) *CreateSOROrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *CreateSOROrder) Do(ctx context.Context) (*CreateSOROrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(CreateSOROrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// CreateTestSOROrder Test new order creation and signature/recvWindow using smart order routing (SOR). Creates and validates a new order but does not send it into the matching engine.
type CreateTestSOROrder struct {
	c *Client
	r *core.Request
}

func (s *CreateTestSOROrder) Symbol(symbol string) *CreateTestSOROrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *CreateTestSOROrder) Side(side core.OrderSideEnum) *CreateTestSOROrder {
	s.r.Set("side", side)
	return s
}
func (s *CreateTestSOROrder) Type(orderType core.OrderTypeEnum) *CreateTestSOROrder {
	s.r.Set("type", orderType)
	return s
}
func (s *CreateTestSOROrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateTestSOROrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}
func (s *CreateTestSOROrder) Quantity(quantity string) *CreateTestSOROrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *CreateTestSOROrder) Price(price string) *CreateTestSOROrder {
	s.r.Set("price", price)
	return s
}
func (s *CreateTestSOROrder) NewClientOrderId(newClientOrderId string) *CreateTestSOROrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *CreateTestSOROrder) StrategyId(strategyId int64) *CreateTestSOROrder {
	s.r.Set("strategyId", strategyId)
	return s
}
func (s *CreateTestSOROrder) StrategyType(strategyType int) *CreateTestSOROrder {
	s.r.Set("strategyType", strategyType)
	return s
}
func (s *CreateTestSOROrder) IcebergQty(icebergQty string) *CreateTestSOROrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}
func (s *CreateTestSOROrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateTestSOROrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *CreateTestSOROrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateTestSOROrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}
func (s *CreateTestSOROrder) RecvWindow(recvWindow int) *CreateTestSOROrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *CreateTestSOROrder) ComputeCommissionRates(computeCommissionRates bool) *CreateTestSOROrder {
	s.r.Set("computeCommissionRates", computeCommissionRates)
	return s
}
func (s *CreateTestSOROrder) Do(ctx context.Context) (*TestCreateOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	if s.r.GetQuery("computeCommissionRates") == "true" {
		resp := new(TestCreateOrderResponse)
		return resp, json.Unmarshal(s.c.rawBody(), resp)
	}
	return nil, nil
}
