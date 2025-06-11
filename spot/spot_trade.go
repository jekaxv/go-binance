package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// CreateOrder Send in a new order.
type CreateOrder struct {
	c                       *Client
	symbol                  string
	side                    core.OrderSideEnum
	orderType               core.OrderTypeEnum
	timeInForce             *core.TimeInForceEnum
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int
	stopPrice               *string
	trailingDelta           *int64
	icebergQty              *string
	newOrderRespType        *core.OrderResponseTypeEnum
	selfTradePreventionMode *core.STPModeEnum
	recvWindow              *int64
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
	s.symbol = symbol
	return s
}

// Side BUY or SELL
func (s *CreateOrder) Side(side core.OrderSideEnum) *CreateOrder {
	s.side = side
	return s
}

func (s *CreateOrder) Type(orderType core.OrderTypeEnum) *CreateOrder {
	s.orderType = orderType
	return s
}

func (s *CreateOrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateOrder {
	s.timeInForce = &timeInForce
	return s
}

func (s *CreateOrder) Quantity(quantity string) *CreateOrder {
	s.quantity = &quantity
	return s
}
func (s *CreateOrder) QuoteOrderQty(quoteOrderQty string) *CreateOrder {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *CreateOrder) Price(price string) *CreateOrder {
	s.price = &price
	return s
}

// NewClientOrderId A unique id among open orders. Automatically generated if not sent.
// Orders with the same newClientOrderID can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (s *CreateOrder) NewClientOrderId(newClientOrderId string) *CreateOrder {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *CreateOrder) StrategyId(strategyId int64) *CreateOrder {
	s.strategyId = &strategyId
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *CreateOrder) StrategyType(strategyType int) *CreateOrder {
	s.strategyType = &strategyType
	return s
}

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateOrder) StopPrice(stopPrice string) *CreateOrder {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateOrder) TrailingDelta(trailingDelta int64) *CreateOrder {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQty Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
func (s *CreateOrder) IcebergQty(icebergQty string) *CreateOrder {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *CreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateOrder {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *CreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateOrder {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateOrder) RecvWindow(recvWindow int64) *CreateOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *CreateOrder) Do(ctx context.Context) (*CreateOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	s.c.set("type", s.orderType)
	if s.timeInForce != nil {
		s.c.set("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		s.c.set("quantity", *s.quantity)
	}
	if s.quoteOrderQty != nil {
		s.c.set("quoteOrderQty", *s.quoteOrderQty)
	}
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		s.c.set("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		s.c.set("strategyType", *s.strategyType)
	}
	if s.stopPrice != nil {
		s.c.set("stopPrice", *s.stopPrice)
	}
	if s.trailingDelta != nil {
		s.c.set("trailingDelta", *s.trailingDelta)
	}
	if s.icebergQty != nil {
		s.c.set("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		s.c.set("newOrderRespType", *s.newOrderRespType)
	}
	if s.selfTradePreventionMode != nil {
		s.c.set("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *CreateOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

type TestCreateOrder struct {
	c                       *Client
	symbol                  string
	side                    core.OrderSideEnum
	orderType               core.OrderTypeEnum
	timeInForce             *core.TimeInForceEnum
	quantity                *string
	quoteOrderQty           *string
	price                   *string
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int
	stopPrice               *string
	trailingDelta           *int64
	icebergQty              *string
	newOrderRespType        *core.OrderResponseTypeEnum
	selfTradePreventionMode *core.STPModeEnum
	recvWindow              *int64
	computeCommissionRates  *bool
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
	s.symbol = symbol
	return s
}

// Side BUY or SELL
func (s *TestCreateOrder) Side(side core.OrderSideEnum) *TestCreateOrder {
	s.side = side
	return s
}

// Type LIMIT,MARKET,STOP_LOSS,STOP_LOSS_LIMIT,TAKE_PROFIT,TAKE_PROFIT_LIMIT,LIMIT_MAKER
func (s *TestCreateOrder) Type(orderType core.OrderTypeEnum) *TestCreateOrder {
	s.orderType = orderType
	return s
}

func (s *TestCreateOrder) TimeInForce(timeInForce core.TimeInForceEnum) *TestCreateOrder {
	s.timeInForce = &timeInForce
	return s
}

func (s *TestCreateOrder) Quantity(quantity string) *TestCreateOrder {
	s.quantity = &quantity
	return s
}
func (s *TestCreateOrder) QuoteOrderQty(quoteOrderQty string) *TestCreateOrder {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *TestCreateOrder) Price(price string) *TestCreateOrder {
	s.price = &price
	return s
}

// NewClientOrderId A unique id among open orders. Automatically generated if not sent.
// Orders with the same newClientOrderID can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (s *TestCreateOrder) NewClientOrderId(newClientOrderId string) *TestCreateOrder {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *TestCreateOrder) StrategyId(strategyId int64) *TestCreateOrder {
	s.strategyId = &strategyId
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *TestCreateOrder) StrategyType(strategyType int) *TestCreateOrder {
	s.strategyType = &strategyType
	return s
}

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *TestCreateOrder) StopPrice(stopPrice string) *TestCreateOrder {
	s.stopPrice = &stopPrice
	return s
}

// TrailingDelta Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *TestCreateOrder) TrailingDelta(trailingDelta int64) *TestCreateOrder {
	s.trailingDelta = &trailingDelta
	return s
}

// IcebergQty Used with LIMIT, STOP_LOSS_LIMIT, and TAKE_PROFIT_LIMIT to create an iceberg order.
func (s *TestCreateOrder) IcebergQty(icebergQty string) *TestCreateOrder {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *TestCreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *TestCreateOrder {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *TestCreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *TestCreateOrder {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *TestCreateOrder) RecvWindow(recvWindow int64) *TestCreateOrder {
	s.recvWindow = &recvWindow
	return s
}

// ComputeCommissionRates Default: false
func (s *TestCreateOrder) ComputeCommissionRates(computeCommissionRates bool) *TestCreateOrder {
	s.computeCommissionRates = &computeCommissionRates
	return s
}

func (s *TestCreateOrder) Do(ctx context.Context) (*TestCreateOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	s.c.set("type", s.orderType)
	if s.timeInForce != nil {
		s.c.set("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		s.c.set("quantity", *s.quantity)
	}
	if s.quoteOrderQty != nil {
		s.c.set("quoteOrderQty", *s.quoteOrderQty)
	}
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		s.c.set("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		s.c.set("strategyType", *s.strategyType)
	}
	if s.stopPrice != nil {
		s.c.set("stopPrice", *s.stopPrice)
	}
	if s.trailingDelta != nil {
		s.c.set("trailingDelta", *s.trailingDelta)
	}
	if s.icebergQty != nil {
		s.c.set("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		s.c.set("newOrderRespType", *s.newOrderRespType)
	}
	if s.selfTradePreventionMode != nil {
		s.c.set("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if s.computeCommissionRates != nil {
		s.c.set("computeCommissionRates", *s.computeCommissionRates)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	if s.computeCommissionRates != nil && *s.computeCommissionRates {
		var resp *TestCreateOrderResponse
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	return nil, nil
}

// QueryOrder Check an order's status.
type QueryOrder struct {
	c                 *Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
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
	s.symbol = symbol
	return s
}

// OrderId For some historical orders cummulativeQuoteQty will be < 0, meaning the data is not available at this time.
func (s *QueryOrder) OrderId(orderId int64) *QueryOrder {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId For some historical orders cummulativeQuoteQty will be < 0, meaning the data is not available at this time.
func (s *QueryOrder) OrigClientOrderId(origClientOrderId string) *QueryOrder {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *QueryOrder) RecvWindow(recvWindow int64) *QueryOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryOrder) Do(ctx context.Context) (*QueryOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.orderId != nil {
		s.c.set("orderId", *s.orderId)
	}
	if s.origClientOrderId != nil {
		s.c.set("origClientOrderId", *s.origClientOrderId)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *QueryOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOrder Cancel an active order.
//
//	Name				Type	Mandatory	Description
//	cancelRestrictions	ENUM	NO			Supported values:
//												ONLY_NEW - Cancel will succeed if the order status is NEW.
//												ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED.
//	recvWindow			LONG	NO			The value cannot be greater than 60000.
//	timestamp			LONG	YES
type CancelOrder struct {
	c                  *Client
	symbol             string
	orderId            *int64
	origClientOrderId  *string
	newClientOrderId   *string
	cancelRestrictions *core.CancelRestrictionEnum
	recvWindow         *int64
}

func (s *CancelOrder) Symbol(symbol string) *CancelOrder {
	s.symbol = symbol
	return s
}

// OrderId Either orderId or origClientOrderId must be sent. If both parameters are sent, orderId takes precedence.
func (s *CancelOrder) OrderId(orderId int64) *CancelOrder {
	s.orderId = &orderId
	return s
}

// OrigClientOrderId Either orderId or origClientOrderId must be sent. If both parameters are sent, orderId takes precedence.
func (s *CancelOrder) OrigClientOrderId(origClientOrderId string) *CancelOrder {
	s.origClientOrderId = &origClientOrderId
	return s
}

// NewClientOrderId Used to uniquely identify this cancel. Automatically generated by default.
func (s *CancelOrder) NewClientOrderId(newClientOrderId string) *CancelOrder {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *CancelOrder) CancelRestrictions(cancelRestrictions core.CancelRestrictionEnum) *CancelOrder {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

func (s *CancelOrder) RecvWindow(recvWindow int64) *CancelOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *CancelOrder) Do(ctx context.Context) (*QueryOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.orderId != nil {
		s.c.set("orderId", *s.orderId)
	}
	if s.origClientOrderId != nil {
		s.c.set("origClientOrderId", *s.origClientOrderId)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.cancelRestrictions != nil {
		s.c.set("cancelRestrictions", *s.cancelRestrictions)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *QueryOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOpenOrder Cancels all active orders on a symbol. This includes orders that are part of an order list.
type CancelOpenOrder struct {
	c          *Client
	symbol     string
	recvWindow *int64
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
	s.symbol = symbol
	return s
}
func (s *CancelOpenOrder) RecvWindow(recvWindow int64) *CancelOpenOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *CancelOpenOrder) Do(ctx context.Context) ([]*CancelOpenOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*CancelOpenOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelReplace Cancel an Existing Order and Send a New Order (TRADE)
// Cancels an existing order and places a new order on the same symbol.
// Filters and Order Count are evaluated before the processing of the cancellation and order placement occurs.
// A new order that was not attempted (i.e. when newOrderResult: NOT_ATTEMPTED ), will still increase the order count by 1.
type CancelReplace struct {
	c                          *Client
	symbol                     string
	side                       core.OrderSideEnum
	orderType                  core.OrderTypeEnum
	cancelReplaceMode          core.CancelReplaceModeEnum
	timeInForce                *core.TimeInForceEnum
	quantity                   *string
	quoteOrderQty              *string
	price                      *string
	cancelNewClientOrderId     *string
	cancelOrigClientOrderId    *string
	cancelOrderId              *int64
	newClientOrderId           *string
	strategyId                 *int64
	strategyType               *int
	stopPrice                  *string
	trailingDelta              *int64
	icebergQty                 *string
	newOrderRespType           *core.OrderResponseTypeEnum
	selfTradePreventionMode    *core.STPModeEnum
	cancelRestrictions         *core.CancelRestrictionEnum
	orderRateLimitExceededMode *core.OrderExceededModeEnum
	recvWindow                 *int64
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
	s.symbol = symbol
	return s
}

// Side BUY or SELL
func (s *CancelReplace) Side(side core.OrderSideEnum) *CancelReplace {
	s.side = side
	return s
}

// Type LIMIT,MARKET,STOP_LOSS,STOP_LOSS_LIMIT,TAKE_PROFIT,TAKE_PROFIT_LIMIT,LIMIT_MAKER
func (s *CancelReplace) Type(orderType core.OrderTypeEnum) *CancelReplace {
	s.orderType = orderType
	return s
}

// CancelReplaceMode The allowed values are:
// STOP_ON_FAILURE - If the cancel request fails, the new order placement will not be attempted.
// ALLOW_FAILURE - new order placement will be attempted even if cancel request fails.
func (s *CancelReplace) CancelReplaceMode(cancelReplaceMode core.CancelReplaceModeEnum) *CancelReplace {
	s.cancelReplaceMode = cancelReplaceMode
	return s
}

func (s *CancelReplace) TimeInForce(timeInForce core.TimeInForceEnum) *CancelReplace {
	s.timeInForce = &timeInForce
	return s
}

func (s *CancelReplace) Quantity(quantity string) *CancelReplace {
	s.quantity = &quantity
	return s
}

func (s *CancelReplace) QuoteOrderQty(quoteOrderQty string) *CancelReplace {
	s.quoteOrderQty = &quoteOrderQty
	return s
}

func (s *CancelReplace) Price(price string) *CancelReplace {
	s.price = &price
	return s
}

func (s *CancelReplace) CancelNewClientOrderId(cancelNewClientOrderId string) *CancelReplace {
	s.cancelNewClientOrderId = &cancelNewClientOrderId
	return s
}

func (s *CancelReplace) CancelOrigClientOrderId(cancelOrigClientOrderId string) *CancelReplace {
	s.cancelOrigClientOrderId = &cancelOrigClientOrderId
	return s
}

func (s *CancelReplace) CancelOrderId(cancelOrderId int64) *CancelReplace {
	s.cancelOrderId = &cancelOrderId
	return s
}

func (s *CancelReplace) NewClientOrderId(newClientOrderId string) *CancelReplace {
	s.newClientOrderId = &newClientOrderId
	return s
}

func (s *CancelReplace) StrategyId(strategyId int64) *CancelReplace {
	s.strategyId = &strategyId
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *CancelReplace) StrategyType(strategyType int) *CancelReplace {
	s.strategyType = &strategyType
	return s
}

func (s *CancelReplace) StopPrice(stopPrice string) *CancelReplace {
	s.stopPrice = &stopPrice
	return s
}

func (s *CancelReplace) TrailingDelta(trailingDelta int64) *CancelReplace {
	s.trailingDelta = &trailingDelta
	return s
}

func (s *CancelReplace) IcebergQty(icebergQty string) *CancelReplace {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType Allowed values:
// ACK, RESULT, FULL
// MARKET and LIMIT orders types default to FULL; all other orders default to ACK
func (s *CancelReplace) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CancelReplace {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol.
func (s *CancelReplace) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CancelReplace {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CancelReplace) CancelRestrictions(cancelRestrictions core.CancelRestrictionEnum) *CancelReplace {
	s.cancelRestrictions = &cancelRestrictions
	return s
}

func (s *CancelReplace) OrderRateLimitExceededMode(orderRateLimitExceededMode core.OrderExceededModeEnum) *CancelReplace {
	s.orderRateLimitExceededMode = &orderRateLimitExceededMode
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CancelReplace) RecvWindow(recvWindow int64) *CancelReplace {
	s.recvWindow = &recvWindow
	return s
}

func (s *CancelReplace) Do(ctx context.Context) (*CancelReplaceResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	s.c.set("type", s.orderType)
	s.c.set("cancelReplaceMode", s.cancelReplaceMode)
	if s.timeInForce != nil {
		s.c.set("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		s.c.set("quantity", *s.quantity)
	}
	if s.quoteOrderQty != nil {
		s.c.set("quoteOrderQty", *s.quoteOrderQty)
	}
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.cancelNewClientOrderId != nil {
		s.c.set("cancelNewClientOrderId", *s.cancelNewClientOrderId)
	}
	if s.cancelOrigClientOrderId != nil {
		s.c.set("cancelOrigClientOrderId", *s.cancelOrigClientOrderId)
	}
	if s.cancelOrderId != nil {
		s.c.set("cancelOrderId", *s.cancelOrderId)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		s.c.set("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		s.c.set("strategyType", *s.strategyType)
	}
	if s.stopPrice != nil {
		s.c.set("stopPrice", *s.stopPrice)
	}
	if s.trailingDelta != nil {
		s.c.set("trailingDelta", *s.trailingDelta)
	}
	if s.icebergQty != nil {
		s.c.set("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		s.c.set("newOrderRespType", *s.newOrderRespType)
	}
	if s.selfTradePreventionMode != nil {
		s.c.set("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.cancelRestrictions != nil {
		s.c.set("cancelRestrictions", *s.cancelRestrictions)
	}
	if s.orderRateLimitExceededMode != nil {
		s.c.set("orderRateLimitExceededMode", *s.orderRateLimitExceededMode)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *CancelReplaceResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OpenOrders Get all open orders on a symbol. Careful when accessing this with no symbol.
type OpenOrders struct {
	c          *Client
	symbol     *string
	recvWindow *int64
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
	s.symbol = &symbol
	return s
}
func (s *OpenOrders) RecvWindow(recvWindow int64) *OpenOrders {
	s.recvWindow = &recvWindow
	return s
}

func (s *OpenOrders) Do(ctx context.Context) ([]*OrdersResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*OrdersResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AllOrders Get all account orders; active, canceled, or filled.
type AllOrders struct {
	c          *Client
	symbol     string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
}

func (s *AllOrders) Symbol(symbol string) *AllOrders {
	s.symbol = symbol
	return s
}

func (s *AllOrders) OrderId(orderId int64) *AllOrders {
	s.orderId = &orderId
	return s
}

func (s *AllOrders) StartTime(startTime int64) *AllOrders {
	s.startTime = &startTime
	return s
}

func (s *AllOrders) EndTime(endTime int64) *AllOrders {
	s.endTime = &endTime
	return s
}

// Limit Default 500; max 1000.
func (s *AllOrders) Limit(limit int) *AllOrders {
	s.limit = &limit
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *AllOrders) RecvWindow(recvWindow int64) *AllOrders {
	s.recvWindow = &recvWindow
	return s
}

func (s *AllOrders) Do(ctx context.Context) ([]*OrdersResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.orderId != nil {
		s.c.set("orderId", *s.orderId)
	}
	if s.startTime != nil {
		s.c.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.set("endTime", *s.endTime)
	}
	if s.limit != nil {
		s.c.set("limit", *s.limit)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*OrdersResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOrderList Cancel an entire Order list
type CancelOrderList struct {
	c                 *Client
	symbol            string
	orderListId       *int64
	listClientOrderId *string
	newClientOrderId  *string
	recvWindow        *int64
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
	s.symbol = symbol
	return s
}

// OrderListId Either orderListId or listClientOrderId must be provided
func (s *CancelOrderList) OrderListId(orderListId int64) *CancelOrderList {
	s.orderListId = &orderListId
	return s
}

// ListClientOrderId Either orderListId or listClientOrderId must be provided
func (s *CancelOrderList) ListClientOrderId(listClientOrderId string) *CancelOrderList {
	s.listClientOrderId = &listClientOrderId
	return s
}

// NewClientOrderId Used to uniquely identify this cancel. Automatically generated by default
func (s *CancelOrderList) NewClientOrderId(newClientOrderId string) *CancelOrderList {
	s.newClientOrderId = &newClientOrderId
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CancelOrderList) RecvWindow(recvWindow int64) *CancelOrderList {
	s.recvWindow = &recvWindow
	return s
}

func (s *CancelOrderList) Do(ctx context.Context) (*OrderListResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.orderListId != nil {
		s.c.set("orderListId", *s.orderListId)
	}
	if s.listClientOrderId != nil {
		s.c.set("listClientOrderId", *s.listClientOrderId)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *OrderListResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryOrderList Retrieves a specific order list based on provided optional parameters.
type QueryOrderList struct {
	c                 *Client
	orderListId       *int64
	origClientOrderId *string
	recvWindow        *int64
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
	s.orderListId = &orderListId
	return s
}

// OrigClientOrderId Either orderListId or listClientOrderId must be provided
func (s *QueryOrderList) OrigClientOrderId(origClientOrderId string) *QueryOrderList {
	s.origClientOrderId = &origClientOrderId
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *QueryOrderList) RecvWindow(recvWindow int64) *QueryOrderList {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryOrderList) Do(ctx context.Context) (*QueryOrderListResponse, error) {
	if s.orderListId != nil {
		s.c.set("orderListId", *s.orderListId)
	}
	if s.origClientOrderId != nil {
		s.c.set("origClientOrderId", *s.origClientOrderId)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *QueryOrderListResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryAllOrderLists Retrieves all order lists based on provided optional parameters.
// Note that the time between startTime and endTime can't be longer than 24 hours.
type QueryAllOrderLists struct {
	c          *Client
	fromId     *int64
	startTime  *int64
	endTime    *int64
	limit      *int
	recvWindow *int64
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
	s.fromId = &fromId
	return s
}
func (s *QueryAllOrderLists) StartTime(startTime int64) *QueryAllOrderLists {
	s.startTime = &startTime
	return s
}
func (s *QueryAllOrderLists) EndTime(endTime int64) *QueryAllOrderLists {
	s.endTime = &endTime
	return s
}

// Limit The default is 500; max is 1000
func (s *QueryAllOrderLists) Limit(limit int) *QueryAllOrderLists {
	s.limit = &limit
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *QueryAllOrderLists) RecvWindow(recvWindow int64) *QueryAllOrderLists {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryAllOrderLists) Do(ctx context.Context) ([]*QueryAllOrderListsResponse, error) {
	if s.fromId != nil {
		s.c.set("fromId", *s.fromId)
	}
	if s.startTime != nil {
		s.c.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.set("endTime", *s.endTime)
	}
	if s.limit != nil {
		s.c.set("limit", *s.limit)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*QueryAllOrderListsResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryOpenOrderList Query Open Order lists
type QueryOpenOrderList struct {
	c          *Client
	recvWindow *int64
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

func (s *QueryOpenOrderList) RecvWindow(recvWindow int64) *QueryOpenOrderList {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryOpenOrderList) Do(ctx context.Context) ([]*QueryOpenOrderListResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*QueryOpenOrderListResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CreateSOROrder Places an order using smart order routing (SOR).
type CreateSOROrder struct {
	c                       *Client
	symbol                  string
	side                    core.OrderSideEnum
	orderType               core.OrderTypeEnum
	timeInForce             *core.TimeInForceEnum
	quantity                string
	price                   *string
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int
	icebergQty              *string
	newOrderRespType        *core.OrderResponseTypeEnum
	selfTradePreventionMode *core.STPModeEnum
	recvWindow              *int64
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
	s.symbol = symbol
	return s
}
func (s *CreateSOROrder) Side(side core.OrderSideEnum) *CreateSOROrder {
	s.side = side
	return s
}
func (s *CreateSOROrder) Type(orderType core.OrderTypeEnum) *CreateSOROrder {
	s.orderType = orderType
	return s
}
func (s *CreateSOROrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateSOROrder {
	s.timeInForce = &timeInForce
	return s
}
func (s *CreateSOROrder) Quantity(quantity string) *CreateSOROrder {
	s.quantity = quantity
	return s
}
func (s *CreateSOROrder) Price(price string) *CreateSOROrder {
	s.price = &price
	return s
}
func (s *CreateSOROrder) NewClientOrderId(newClientOrderId string) *CreateSOROrder {
	s.newClientOrderId = &newClientOrderId
	return s
}
func (s *CreateSOROrder) StrategyId(strategyId int64) *CreateSOROrder {
	s.strategyId = &strategyId
	return s
}

// StrategyType The value cannot be less than 1000000.
func (s *CreateSOROrder) StrategyType(strategyType int) *CreateSOROrder {
	s.strategyType = &strategyType
	return s
}

// IcebergQty Used with LIMIT to create an iceberg order.
func (s *CreateSOROrder) IcebergQty(icebergQty string) *CreateSOROrder {
	s.icebergQty = &icebergQty
	return s
}

// NewOrderRespType set the response JSON. ACK, RESULT, or FULL. Default to FULL
func (s *CreateSOROrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateSOROrder {
	s.newOrderRespType = &newOrderRespType
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol.
func (s *CreateSOROrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateSOROrder {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateSOROrder) RecvWindow(recvWindow int64) *CreateSOROrder {
	s.recvWindow = &recvWindow
	return s
}
func (s *CreateSOROrder) Do(ctx context.Context) (*CreateSOROrderResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	s.c.set("type", s.orderType)
	if s.timeInForce != nil {
		s.c.set("timeInForce", *s.timeInForce)
	}
	s.c.set("quantity", s.quantity)
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		s.c.set("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		s.c.set("strategyType", *s.strategyType)
	}
	if s.icebergQty != nil {
		s.c.set("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		s.c.set("newOrderRespType", *s.newOrderRespType)
	}
	if s.selfTradePreventionMode != nil {
		s.c.set("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *CreateSOROrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CreateTestSOROrder Test new order creation and signature/recvWindow using smart order routing (SOR). Creates and validates a new order but does not send it into the matching engine.
type CreateTestSOROrder struct {
	c                       *Client
	symbol                  string
	side                    core.OrderSideEnum
	orderType               core.OrderTypeEnum
	timeInForce             *core.TimeInForceEnum
	quantity                string
	price                   *string
	newClientOrderId        *string
	strategyId              *int64
	strategyType            *int
	icebergQty              *string
	newOrderRespType        *core.OrderResponseTypeEnum
	selfTradePreventionMode *core.STPModeEnum
	recvWindow              *int64
	computeCommissionRates  *bool
}

func (s *CreateTestSOROrder) Symbol(symbol string) *CreateTestSOROrder {
	s.symbol = symbol
	return s
}
func (s *CreateTestSOROrder) Side(side core.OrderSideEnum) *CreateTestSOROrder {
	s.side = side
	return s
}
func (s *CreateTestSOROrder) Type(orderType core.OrderTypeEnum) *CreateTestSOROrder {
	s.orderType = orderType
	return s
}
func (s *CreateTestSOROrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateTestSOROrder {
	s.timeInForce = &timeInForce
	return s
}
func (s *CreateTestSOROrder) Quantity(quantity string) *CreateTestSOROrder {
	s.quantity = quantity
	return s
}
func (s *CreateTestSOROrder) Price(price string) *CreateTestSOROrder {
	s.price = &price
	return s
}
func (s *CreateTestSOROrder) NewClientOrderId(newClientOrderId string) *CreateTestSOROrder {
	s.newClientOrderId = &newClientOrderId
	return s
}
func (s *CreateTestSOROrder) StrategyId(strategyId int64) *CreateTestSOROrder {
	s.strategyId = &strategyId
	return s
}
func (s *CreateTestSOROrder) StrategyType(strategyType int) *CreateTestSOROrder {
	s.strategyType = &strategyType
	return s
}
func (s *CreateTestSOROrder) IcebergQty(icebergQty string) *CreateTestSOROrder {
	s.icebergQty = &icebergQty
	return s
}
func (s *CreateTestSOROrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateTestSOROrder {
	s.newOrderRespType = &newOrderRespType
	return s
}
func (s *CreateTestSOROrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateTestSOROrder {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}
func (s *CreateTestSOROrder) RecvWindow(recvWindow int64) *CreateTestSOROrder {
	s.recvWindow = &recvWindow
	return s
}
func (s *CreateTestSOROrder) ComputeCommissionRates(computeCommissionRates bool) *CreateTestSOROrder {
	s.computeCommissionRates = &computeCommissionRates
	return s
}
func (s *CreateTestSOROrder) Do(ctx context.Context) (*TestCreateOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	s.c.set("type", s.orderType)
	if s.timeInForce != nil {
		s.c.set("timeInForce", *s.timeInForce)
	}
	s.c.set("quantity", s.quantity)
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.strategyId != nil {
		s.c.set("strategyId", *s.strategyId)
	}
	if s.strategyType != nil {
		s.c.set("strategyType", *s.strategyType)
	}
	if s.icebergQty != nil {
		s.c.set("icebergQty", *s.icebergQty)
	}
	if s.newOrderRespType != nil {
		s.c.set("newOrderRespType", *s.newOrderRespType)
	}
	if s.selfTradePreventionMode != nil {
		s.c.set("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if s.computeCommissionRates != nil {
		s.c.set("computeCommissionRates", *s.computeCommissionRates)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	if s.computeCommissionRates != nil && *s.computeCommissionRates {
		var resp *TestCreateOrderResponse
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	return nil, nil
}
