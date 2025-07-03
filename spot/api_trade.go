package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

type ApiOrderError struct {
	ApiError
	Data struct {
		CancelResult   string `json:"cancelResult"`
		NewOrderResult string `json:"newOrderResult"`
		CancelResponse struct {
			Code int    `json:"code"`
			Msg  string `json:"msg"`
		} `json:"cancelResponse"`
		NewOrderResponse *ApiOrderReport `json:"newOrderResponse"`
	} `json:"data"`
}

// WsCreateOrder Send in a new order.
type WsCreateOrder struct {
	c *WsClient
	r *core.WsRequest
}

type CreateOrderResult struct {
	Symbol              string          `json:"symbol"`
	OrderId             int64           `json:"orderId"`
	OrderListId         int             `json:"orderListId"`
	ClientOrderId       string          `json:"clientOrderId"`
	TransactTime        int64           `json:"transactTime"`
	Price               decimal.Decimal `json:"price"`
	OrigQty             decimal.Decimal `json:"origQty"`
	ExecutedQty         decimal.Decimal `json:"executedQty"`
	OrigQuoteOrderQty   decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty decimal.Decimal `json:"cummulativeQuoteQty"`
	Status              string          `json:"status"`
	TimeInForce         string          `json:"timeInForce"`
	Type                string          `json:"type"`
	Side                string          `json:"side"`
	WorkingTime         int64           `json:"workingTime"`
	Fills               []*ApiFill      `json:"fills"`
}

type WsCreateOrderResponse struct {
	ApiResponse
	Result *CreateOrderResult `json:"result"`
}

func (s *WsCreateOrder) Symbol(symbol string) *WsCreateOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsCreateOrder) Side(side core.OrderSideEnum) *WsCreateOrder {
	s.r.Set("side", side)
	return s
}
func (s *WsCreateOrder) Type(orderType core.OrderTypeEnum) *WsCreateOrder {
	s.r.Set("type", orderType)
	return s
}
func (s *WsCreateOrder) TimeInForce(timeInForce core.TimeInForceEnum) *WsCreateOrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}
func (s *WsCreateOrder) Price(price string) *WsCreateOrder {
	s.r.Set("price", price)
	return s
}
func (s *WsCreateOrder) Quantity(quantity string) *WsCreateOrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *WsCreateOrder) QuoteOrderQty(quoteOrderQty string) *WsCreateOrder {
	s.r.Set("quoteOrderQty", quoteOrderQty)
	return s
}
func (s *WsCreateOrder) NewClientOrderId(newClientOrderId string) *WsCreateOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *WsCreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCreateOrder) StopPrice(stopPrice string) *WsCreateOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}
func (s *WsCreateOrder) TrailingDelta(trailingDelta int) *WsCreateOrder {
	s.r.Set("trailingDelta", trailingDelta)
	return s
}
func (s *WsCreateOrder) IcebergQty(icebergQty string) *WsCreateOrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}
func (s *WsCreateOrder) StrategyId(strategyId int64) *WsCreateOrder {
	s.r.Set("strategyId", strategyId)
	return s
}
func (s *WsCreateOrder) StrategyType(strategyType int) *WsCreateOrder {
	s.r.Set("strategyType", strategyType)
	return s
}
func (s *WsCreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *WsCreateOrder) RecvWindow(recvWindow int) *WsCreateOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateOrder) Do(ctx context.Context) (*WsCreateOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCreateOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

type WsCreateTestOrder struct {
	c *WsClient
	r *core.WsRequest
}

type CreateTestOrderResult struct {
	StandardCommissionForOrder *OrderCommission `json:"standardCommissionForOrder"`
	TaxCommissionForOrder      *OrderCommission `json:"taxCommissionForOrder"`
	Discount                   *Discount        `json:"discount"`
}

type WsCreateOrderTestResponse struct {
	ApiResponse
	Result *CreateTestOrderResult `json:"result"`
}

func (s *WsCreateTestOrder) Symbol(symbol string) *WsCreateTestOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsCreateTestOrder) Side(side core.OrderSideEnum) *WsCreateTestOrder {
	s.r.Set("side", side)
	return s
}
func (s *WsCreateTestOrder) Type(orderType core.OrderTypeEnum) *WsCreateTestOrder {
	s.r.Set("type", orderType)
	return s
}
func (s *WsCreateTestOrder) TimeInForce(timeInForce core.TimeInForceEnum) *WsCreateTestOrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}
func (s *WsCreateTestOrder) Price(price string) *WsCreateTestOrder {
	s.r.Set("price", price)
	return s
}
func (s *WsCreateTestOrder) Quantity(quantity string) *WsCreateTestOrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *WsCreateTestOrder) QuoteOrderQty(quoteOrderQty string) *WsCreateTestOrder {
	s.r.Set("quoteOrderQty", quoteOrderQty)
	return s
}
func (s *WsCreateTestOrder) NewClientOrderId(newClientOrderId string) *WsCreateTestOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *WsCreateTestOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateTestOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCreateTestOrder) StopPrice(stopPrice string) *WsCreateTestOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}
func (s *WsCreateTestOrder) TrailingDelta(trailingDelta int) *WsCreateTestOrder {
	s.r.Set("trailingDelta", trailingDelta)
	return s
}
func (s *WsCreateTestOrder) IcebergQty(icebergQty string) *WsCreateTestOrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}
func (s *WsCreateTestOrder) StrategyId(strategyId int64) *WsCreateTestOrder {
	s.r.Set("strategyId", strategyId)
	return s
}
func (s *WsCreateTestOrder) StrategyType(strategyType int) *WsCreateTestOrder {
	s.r.Set("strategyType", strategyType)
	return s
}
func (s *WsCreateTestOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateTestOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *WsCreateTestOrder) RecvWindow(recvWindow int) *WsCreateTestOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateTestOrder) ComputeCommissionRates(computeCommissionRates bool) *WsCreateTestOrder {
	s.r.Set("computeCommissionRates", computeCommissionRates)
	return s
}

func (s *WsCreateTestOrder) Do(ctx context.Context) (*WsCreateOrderTestResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCreateOrderTestResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsQueryOrder Check execution status of an order.
type WsQueryOrder struct {
	c *WsClient
	r *core.WsRequest
}

type QueryOrderResult struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int64           `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	TrailingDelta           int             `json:"trailingDelta"`
	TrailingTime            int             `json:"trailingTime"`
	IcebergQty              decimal.Decimal `json:"icebergQty"`
	Time                    int64           `json:"time"`
	UpdateTime              int64           `json:"updateTime"`
	IsWorking               bool            `json:"isWorking"`
	WorkingTime             int64           `json:"workingTime"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	StrategyId              int             `json:"strategyId"`
	StrategyType            int             `json:"strategyType"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	PreventedMatchId        int             `json:"preventedMatchId"`
	PreventedQuantity       decimal.Decimal `json:"preventedQuantity"`
}

type WsQueryOrderResponse struct {
	ApiResponse
	Result *QueryOrderResult `json:"result"`
}

func (s *WsQueryOrder) Symbol(symbol string) *WsQueryOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsQueryOrder) OrderId(orderId int64) *WsQueryOrder {
	s.r.Set("orderId", orderId)
	return s
}
func (s *WsQueryOrder) OrigClientOrderId(origClientOrderId string) *WsQueryOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

func (s *WsQueryOrder) RecvWindow(recvWindow int) *WsQueryOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsQueryOrder) Do(ctx context.Context) (*WsQueryOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsQueryOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCancelOrder Cancel an active order.
type WsCancelOrder struct {
	c *WsClient
	r *core.WsRequest
}

type CancelOrderResult struct {
	Symbol                  string            `json:"symbol"`
	OrigClientOrderId       string            `json:"origClientOrderId"`
	OrderId                 int64             `json:"orderId"`
	OrderListId             int               `json:"orderListId"`
	ClientOrderId           string            `json:"clientOrderId"`
	TransactTime            int64             `json:"transactTime"`
	Price                   decimal.Decimal   `json:"price"`
	OrigQty                 decimal.Decimal   `json:"origQty"`
	ExecutedQty             decimal.Decimal   `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal   `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal   `json:"cummulativeQuoteQty"`
	Status                  string            `json:"status"`
	TimeInForce             string            `json:"timeInForce"`
	Type                    string            `json:"type"`
	Side                    string            `json:"side"`
	StopPrice               decimal.Decimal   `json:"stopPrice"`
	TrailingDelta           int               `json:"trailingDelta"`
	IcebergQty              decimal.Decimal   `json:"icebergQty"`
	StrategyId              int               `json:"strategyId"`
	StrategyType            int               `json:"strategyType"`
	SelfTradePreventionMode string            `json:"selfTradePreventionMode"`
	ContingencyType         string            `json:"contingencyType"`
	ListStatusType          string            `json:"listStatusType"`
	ListOrderStatus         string            `json:"listOrderStatus"`
	ListClientOrderId       string            `json:"listClientOrderId"`
	TransactionTime         int64             `json:"transactionTime"`
	Orders                  []*ApiOrder       `json:"orders"`
	OrderReports            []*ApiOrderReport `json:"orderReports"`
}

type WsCancelOrderResponse struct {
	ApiResponse
	Result *CancelOrderResult `json:"result"`
}

func (s *WsCancelOrder) Symbol(symbol string) *WsCancelOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsCancelOrder) OrderId(orderId int64) *WsCancelOrder {
	s.r.Set("orderId", orderId)
	return s
}

func (s *WsCancelOrder) OrigClientOrderId(origClientOrderId string) *WsCancelOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

func (s *WsCancelOrder) NewClientOrderId(newClientOrderId string) *WsCancelOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

// CancelRestrictions Supported values:
// ONLY_NEW - Cancel will succeed if the order status is NEW.
// ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED.
func (s *WsCancelOrder) CancelRestrictions(cancelRestrictions core.CancelRestrictionEnum) *WsCancelOrder {
	s.r.Set("cancelRestrictions", cancelRestrictions)
	return s
}

func (s *WsCancelOrder) RecvWindow(recvWindow int) *WsCancelOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCancelOrder) Do(ctx context.Context) (*WsCancelOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCancelOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCancelReplaceOrder Cancel an existing order and immediately place a new order instead of the canceled one.
type WsCancelReplaceOrder struct {
	c *WsClient
	r *core.WsRequest
}
type CancelReplaceOrderResult struct {
	CancelResult     string          `json:"cancelResult"`
	NewOrderResult   string          `json:"newOrderResult"`
	CancelResponse   *ApiOrderReport `json:"cancelResponse"`
	NewOrderResponse *ApiOrderReport `json:"newOrderResponse"`
}
type WsCancelReplaceOrderResponse struct {
	Id         string                    `json:"id"`
	Status     int                       `json:"status"`
	RateLimits []*ApiRateLimit           `json:"rateLimits,omitempty"`
	Error      *ApiOrderError            `json:"error,omitempty"`
	Result     *CancelReplaceOrderResult `json:"result"`
}

func (s *WsCancelReplaceOrder) Symbol(symbol string) *WsCancelReplaceOrder {
	s.r.Set("symbol", symbol)
	return s
}

// CancelReplaceMode Available cancelReplaceMode options:
// STOP_ON_FAILURE – if cancellation request fails, new order placement will not be attempted.
// ALLOW_FAILURE – new order placement will be attempted even if the cancel request fails.
func (s *WsCancelReplaceOrder) CancelReplaceMode(cancelReplaceMode core.CancelReplaceModeEnum) *WsCancelReplaceOrder {
	s.r.Set("cancelReplaceMode", cancelReplaceMode)
	return s
}

// CancelOrderId Cancel order by orderId
func (s *WsCancelReplaceOrder) CancelOrderId(cancelOrderId int64) *WsCancelReplaceOrder {
	s.r.Set("cancelOrderId", cancelOrderId)
	return s
}

// CancelOrigClientOrderId Cancel order by clientOrderId
func (s *WsCancelReplaceOrder) CancelOrigClientOrderId(cancelOrigClientOrderId string) *WsCancelReplaceOrder {
	s.r.Set("cancelOrigClientOrderId", cancelOrigClientOrderId)
	return s
}

// CancelNewClientOrderId New ID for the canceled order. Automatically generated if not sent
func (s *WsCancelReplaceOrder) CancelNewClientOrderId(cancelNewClientOrderId string) *WsCancelReplaceOrder {
	s.r.Set("cancelNewClientOrderId", cancelNewClientOrderId)
	return s
}
func (s *WsCancelReplaceOrder) Side(side core.OrderSideEnum) *WsCancelReplaceOrder {
	s.r.Set("side", side)
	return s
}
func (s *WsCancelReplaceOrder) Type(type_ core.OrderTypeEnum) *WsCancelReplaceOrder {
	s.r.Set("type", type_)
	return s
}
func (s *WsCancelReplaceOrder) TimeInForce(timeInForce core.TimeInForceEnum) *WsCancelReplaceOrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}
func (s *WsCancelReplaceOrder) Price(price string) *WsCancelReplaceOrder {
	s.r.Set("price", price)
	return s
}
func (s *WsCancelReplaceOrder) Quantity(quantity string) *WsCancelReplaceOrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *WsCancelReplaceOrder) QuoteOrderQty(quoteOrderQty string) *WsCancelReplaceOrder {
	s.r.Set("quoteOrderQty", quoteOrderQty)
	return s
}

// NewClientOrderId Arbitrary unique ID among open orders. Automatically generated if not sent
func (s *WsCancelReplaceOrder) NewClientOrderId(newClientOrderId string) *WsCancelReplaceOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

// NewOrderRespType Select response format: ACK, RESULT, FULL.
// MARKET and LIMIT orders produce FULL response by default, other order types default to ACK.
func (s *WsCancelReplaceOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCancelReplaceOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCancelReplaceOrder) StopPrice(stopPrice string) *WsCancelReplaceOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}
func (s *WsCancelReplaceOrder) TrailingDelta(trailingDelta string) *WsCancelReplaceOrder {
	s.r.Set("trailingDelta", trailingDelta)
	return s
}
func (s *WsCancelReplaceOrder) IcebergQty(icebergQty string) *WsCancelReplaceOrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}

// StrategyId Arbitrary numeric value identifying the order within an order strategy.
func (s *WsCancelReplaceOrder) StrategyId(strategyId int64) *WsCancelReplaceOrder {
	s.r.Set("strategyId", strategyId)
	return s
}

// StrategyType Arbitrary numeric value identifying the order strategy.
// Values smaller than 1000000 are reserved and cannot be used.
func (s *WsCancelReplaceOrder) StrategyType(strategyType int) *WsCancelReplaceOrder {
	s.r.Set("strategyType", strategyType)
	return s
}
func (s *WsCancelReplaceOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCancelReplaceOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

// CancelRestrictions Supported values:
// ONLY_NEW - Cancel will succeed if the order status is NEW.
// ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED. For more information please refer to Regarding cancelRestrictions.
func (s *WsCancelReplaceOrder) CancelRestrictions(cancelRestrictions core.CancelRestrictionEnum) *WsCancelReplaceOrder {
	s.r.Set("cancelRestrictions", cancelRestrictions)
	return s
}

// OrderRateLimitExceededMode Supported values:
// DO_NOTHING (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit
// CANCEL_ONLY - will always cancel the order.
func (s *WsCancelReplaceOrder) OrderRateLimitExceededMode(orderRateLimitExceededMode core.OrderExceededModeEnum) *WsCancelReplaceOrder {
	s.r.Set("orderRateLimitExceededMode", orderRateLimitExceededMode)
	return s
}

func (s *WsCancelReplaceOrder) RecvWindow(recvWindow int) *WsCancelReplaceOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCancelReplaceOrder) Do(ctx context.Context) (*WsCancelReplaceOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCancelReplaceOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsOpenOrdersStatus Query execution status of all open orders.
type WsOpenOrdersStatus struct {
	c *WsClient
	r *core.WsRequest
}

type OpenOrdersStatusResult struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int64           `json:"orderId"`
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
type WsOpenOrdersStatusResponse struct {
	ApiResponse
	Result []*OpenOrdersStatusResult `json:"result"`
}

func (s *WsOpenOrdersStatus) Symbol(symbol string) *WsOpenOrdersStatus {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsOpenOrdersStatus) RecvWindow(recvWindow int) *WsOpenOrdersStatus {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *WsOpenOrdersStatus) Do(ctx context.Context) (*WsOpenOrdersStatusResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsOpenOrdersStatusResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCancelOpenOrder Cancel all open orders on a symbol. This includes orders that are part of an order list.
type WsCancelOpenOrder struct {
	c *WsClient
	r *core.WsRequest
}

type OrderResult struct {
	Symbol                  string            `json:"symbol"`
	OrigClientOrderId       string            `json:"origClientOrderId"`
	OrderId                 int64             `json:"orderId"`
	OrderListId             int               `json:"orderListId"`
	ClientOrderId           string            `json:"clientOrderId"`
	TransactTime            int64             `json:"transactTime"`
	Price                   decimal.Decimal   `json:"price"`
	OrigQty                 decimal.Decimal   `json:"origQty"`
	ExecutedQty             decimal.Decimal   `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal   `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal   `json:"cummulativeQuoteQty"`
	Status                  string            `json:"status"`
	TimeInForce             string            `json:"timeInForce"`
	Type                    string            `json:"type"`
	Side                    string            `json:"side"`
	StopPrice               decimal.Decimal   `json:"stopPrice"`
	TrailingDelta           int               `json:"trailingDelta"`
	TrailingTime            int               `json:"trailingTime"`
	IcebergQty              decimal.Decimal   `json:"icebergQty"`
	StrategyId              int               `json:"strategyId"`
	StrategyType            int               `json:"strategyType"`
	SelfTradePreventionMode string            `json:"selfTradePreventionMode"`
	ContingencyType         string            `json:"contingencyType"`
	ListStatusType          string            `json:"listStatusType"`
	ListOrderStatus         string            `json:"listOrderStatus"`
	ListClientOrderId       string            `json:"listClientOrderId"`
	TransactionTime         int64             `json:"transactionTime"`
	Orders                  []*ApiOrder       `json:"orders"`
	OrderReports            []*ApiOrderReport `json:"orderReports"`
}

type OrderResponse struct {
	ApiResponse
	Result *OrderResult `json:"result"`
}

type WsCancelOpenOrderResponse struct {
	ApiResponse
	Result []*OrderResult `json:"result"`
}

func (s *WsCancelOpenOrder) Symbol(symbol string) *WsCancelOpenOrder {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsCancelOpenOrder) RecvWindow(recvWindow int) *WsCancelOpenOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCancelOpenOrder) Do(ctx context.Context) (*WsCancelOpenOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCancelOpenOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCreateOCOOrder Send in an one-cancels the other (OCO) pair, where activation of one order immediately cancels the other.
type WsCreateOCOOrder struct {
	c *WsClient
	r *core.WsRequest
}

func (s *WsCreateOCOOrder) Symbol(symbol string) *WsCreateOCOOrder {
	s.r.Set("symbol", symbol)
	return s
}

// ListClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent.
// A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.
// listClientOrderId is distinct from the aboveClientOrderId and the belowCLientOrderId
func (s *WsCreateOCOOrder) ListClientOrderId(listClientOrderId string) *WsCreateOCOOrder {
	s.r.Set("listClientOrderId", listClientOrderId)
	return s
}
func (s *WsCreateOCOOrder) Side(side core.OrderSideEnum) *WsCreateOCOOrder {
	s.r.Set("side", side)
	return s
}

// Quantity for both orders of the order list.
func (s *WsCreateOCOOrder) Quantity(quantity string) *WsCreateOCOOrder {
	s.r.Set("quantity", quantity)
	return s
}

// AboveType Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
func (s *WsCreateOCOOrder) AboveType(aboveType core.OrderTypeEnum) *WsCreateOCOOrder {
	s.r.Set("aboveType", aboveType)
	return s
}

// AboveClientOrderId Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
func (s *WsCreateOCOOrder) AboveClientOrderId(aboveClientOrderId string) *WsCreateOCOOrder {
	s.r.Set("aboveClientOrderId", aboveClientOrderId)
	return s
}

// AboveIcebergQty Note that this can only be used if aboveTimeInForce is GTC.
func (s *WsCreateOCOOrder) AboveIcebergQty(aboveIcebergQty int64) *WsCreateOCOOrder {
	s.r.Set("aboveIcebergQty", aboveIcebergQty)
	return s
}

// AbovePrice Can be used if aboveType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
func (s *WsCreateOCOOrder) AbovePrice(abovePrice string) *WsCreateOCOOrder {
	s.r.Set("abovePrice", abovePrice)
	return s
}

// AboveStopPrice Can be used if aboveType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT
// Either aboveStopPrice or aboveTrailingDelta or both, must be specified.
func (s *WsCreateOCOOrder) AboveStopPrice(aboveStopPrice string) *WsCreateOCOOrder {
	s.r.Set("aboveStopPrice", aboveStopPrice)
	return s
}
func (s *WsCreateOCOOrder) AboveTrailingDelta(aboveTrailingDelta int64) *WsCreateOCOOrder {
	s.r.Set("aboveTrailingDelta", aboveTrailingDelta)
	return s
}

// AboveTimeInForce Required if aboveType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT.
func (s *WsCreateOCOOrder) AboveTimeInForce(aboveTimeInForce core.TimeInForceEnum) *WsCreateOCOOrder {
	s.r.Set("aboveTimeInForce", aboveTimeInForce)
	return s
}

// AboveStrategyId Arbitrary numeric value identifying the above order within an order strategy.
func (s *WsCreateOCOOrder) AboveStrategyId(aboveStrategyId int64) *WsCreateOCOOrder {
	s.r.Set("aboveStrategyId", aboveStrategyId)
	return s
}

// AboveStrategyType Arbitrary numeric value identifying the above order strategy.
// Values smaller than 1000000 are reserved and cannot be used.
func (s *WsCreateOCOOrder) AboveStrategyType(aboveStrategyType int64) *WsCreateOCOOrder {
	s.r.Set("aboveStrategyType", aboveStrategyType)
	return s
}

// BelowType Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
func (s *WsCreateOCOOrder) BelowType(belowType core.OrderTypeEnum) *WsCreateOCOOrder {
	s.r.Set("belowType", belowType)
	return s
}

func (s *WsCreateOCOOrder) BelowClientOrderId(belowClientOrderId string) *WsCreateOCOOrder {
	s.r.Set("belowClientOrderId", belowClientOrderId)
	return s
}

// BelowIcebergQty Note that this can only be used if belowTimeInForce is GTC.
func (s *WsCreateOCOOrder) BelowIcebergQty(belowIcebergQty int64) *WsCreateOCOOrder {
	s.r.Set("belowIcebergQty", belowIcebergQty)
	return s
}

// BelowPrice Can be used if belowType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
func (s *WsCreateOCOOrder) BelowPrice(belowPrice string) *WsCreateOCOOrder {
	s.r.Set("belowPrice", belowPrice)
	return s
}

// BelowStopPrice Can be used if belowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT.
func (s *WsCreateOCOOrder) BelowStopPrice(belowStopPrice string) *WsCreateOCOOrder {
	s.r.Set("belowStopPrice", belowStopPrice)
	return s
}
func (s *WsCreateOCOOrder) BelowTrailingDelta(belowTrailingDelta int64) *WsCreateOCOOrder {
	s.r.Set("belowTrailingDelta", belowTrailingDelta)
	return s
}

// BelowTimeInForce Required if belowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT
func (s *WsCreateOCOOrder) BelowTimeInForce(belowTimeInForce string) *WsCreateOCOOrder {
	s.r.Set("belowTimeInForce", belowTimeInForce)
	return s
}
func (s *WsCreateOCOOrder) BelowStrategyId(belowStrategyId int64) *WsCreateOCOOrder {
	s.r.Set("belowStrategyId", belowStrategyId)
	return s
}
func (s *WsCreateOCOOrder) BelowStrategyType(belowStrategyType int64) *WsCreateOCOOrder {
	s.r.Set("belowStrategyType", belowStrategyType)
	return s
}

// NewOrderRespType Select response format: ACK, RESULT, FULL
func (s *WsCreateOCOOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateOCOOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCreateOCOOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateOCOOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}
func (s *WsCreateOCOOrder) RecvWindow(recvWindow int) *WsCreateOCOOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateOCOOrder) Do(ctx context.Context) (*OrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *OrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCreateOTOOrder Places an OTO.
// An OTO (One-Triggers-the-Other) is an order list comprised of 2 orders.
// The first order is called the working order and must be LIMIT or LIMIT_MAKER. Initially, only the working order goes on the order book.
// The second order is called the pending order. It can be any order type except for MARKET orders using parameter quoteOrderQty. The pending order is only placed on the order book when the working order gets fully filled.
// If either the working order or the pending order is cancelled individually, the other order in the order list will also be canceled or expired.
// OTOs add 2 orders to the unfilled order count, EXCHANGE_MAX_NUM_ORDERS filter and MAX_NUM_ORDERS filter.
type WsCreateOTOOrder struct {
	c *WsClient
	r *core.WsRequest
}

func (s *WsCreateOTOOrder) Symbol(symbol string) *WsCreateOTOOrder {
	s.r.Set("symbol", symbol)
	return s
}

// ListClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent.
// A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.
// listClientOrderId is distinct from the workingClientOrderId and the pendingClientOrderId.
func (s *WsCreateOTOOrder) ListClientOrderId(listClientOrderId string) *WsCreateOTOOrder {
	s.r.Set("listClientOrderId", listClientOrderId)
	return s
}

func (s *WsCreateOTOOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateOTOOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}

func (s *WsCreateOTOOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateOTOOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

// WorkingType Supported values: LIMIT,LIMIT_MAKER
func (s *WsCreateOTOOrder) WorkingType(workingType core.OrderTypeEnum) *WsCreateOTOOrder {
	s.r.Set("workingType", workingType)
	return s
}

func (s *WsCreateOTOOrder) WorkingSide(workingSide core.OrderSideEnum) *WsCreateOTOOrder {
	s.r.Set("workingSide", workingSide)
	return s
}

// WorkingClientOrderId Arbitrary unique ID among open orders for the working order.
// Automatically generated if not sent.
func (s *WsCreateOTOOrder) WorkingClientOrderId(workingClientOrderId string) *WsCreateOTOOrder {
	s.r.Set("workingClientOrderId", workingClientOrderId)
	return s
}
func (s *WsCreateOTOOrder) WorkingPrice(workingPrice string) *WsCreateOTOOrder {
	s.r.Set("workingPrice", workingPrice)
	return s
}

// WorkingQuantity Sets the quantity for the working order.
func (s *WsCreateOTOOrder) WorkingQuantity(workingQuantity string) *WsCreateOTOOrder {
	s.r.Set("workingQuantity", workingQuantity)
	return s
}

// WorkingIcebergQty This can only be used if workingTimeInForce is GTC, or if workingType is LIMIT_MAKER.
func (s *WsCreateOTOOrder) WorkingIcebergQty(workingIcebergQty string) *WsCreateOTOOrder {
	s.r.Set("workingIcebergQty", workingIcebergQty)
	return s
}
func (s *WsCreateOTOOrder) WorkingTimeInForce(workingTimeInForce core.TimeInForceEnum) *WsCreateOTOOrder {
	s.r.Set("workingTimeInForce", workingTimeInForce)
	return s
}

// WorkingStrategyId	Arbitrary numeric value identifying the working order within an order strategy.
func (s *WsCreateOTOOrder) WorkingStrategyId(workingStrategyId int64) *WsCreateOTOOrder {
	s.r.Set("workingStrategyId", workingStrategyId)
	return s
}
func (s *WsCreateOTOOrder) WorkingStrategyType(workingStrategyType int) *WsCreateOTOOrder {
	s.r.Set("workingStrategyType", workingStrategyType)
	return s
}

// PendingType Supported values: Order Types Note that MARKET orders using quoteOrderQty are not supported.
func (s *WsCreateOTOOrder) PendingType(pendingType core.OrderTypeEnum) *WsCreateOTOOrder {
	s.r.Set("pendingType", pendingType)
	return s
}
func (s *WsCreateOTOOrder) PendingSide(pendingSide core.OrderSideEnum) *WsCreateOTOOrder {
	s.r.Set("pendingSide", pendingSide)
	return s
}

// PendingClientOrderId Arbitrary unique ID among open orders for the pending order.
// Automatically generated if not sent.
func (s *WsCreateOTOOrder) PendingClientOrderId(pendingClientOrderId string) *WsCreateOTOOrder {
	s.r.Set("pendingClientOrderId", pendingClientOrderId)
	return s
}
func (s *WsCreateOTOOrder) PendingPrice(pendingPrice string) *WsCreateOTOOrder {
	s.r.Set("pendingPrice", pendingPrice)
	return s
}
func (s *WsCreateOTOOrder) PendingStopPrice(pendingStopPrice string) *WsCreateOTOOrder {
	s.r.Set("pendingStopPrice", pendingStopPrice)
	return s
}
func (s *WsCreateOTOOrder) PendingTrailingDelta(pendingTrailingDelta string) *WsCreateOTOOrder {
	s.r.Set("pendingTrailingDelta", pendingTrailingDelta)
	return s
}
func (s *WsCreateOTOOrder) PendingQuantity(pendingQuantity string) *WsCreateOTOOrder {
	s.r.Set("pendingQuantity", pendingQuantity)
	return s
}
func (s *WsCreateOTOOrder) PendingIcebergQty(pendingIcebergQty string) *WsCreateOTOOrder {
	s.r.Set("pendingIcebergQty", pendingIcebergQty)
	return s
}
func (s *WsCreateOTOOrder) PendingTimeInForce(pendingTimeInForce core.TimeInForceEnum) *WsCreateOTOOrder {
	s.r.Set("pendingTimeInForce", pendingTimeInForce)
	return s
}
func (s *WsCreateOTOOrder) PendingStrategyId(pendingStrategyId int64) *WsCreateOTOOrder {
	s.r.Set("pendingStrategyId", pendingStrategyId)
	return s
}
func (s *WsCreateOTOOrder) PendingStrategyType(pendingStrategyType int) *WsCreateOTOOrder {
	s.r.Set("pendingStrategyType", pendingStrategyType)
	return s
}

func (s *WsCreateOTOOrder) RecvWindow(recvWindow int) *WsCreateOTOOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateOTOOrder) Do(ctx context.Context) (*OrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *OrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCreateOTOCOOrder Place an OTOCO.
// An OTOCO (One-Triggers-One-Cancels-the-Other) is an order list comprised of 3 orders.
// The first order is called the working order and must be LIMIT or LIMIT_MAKER. Initially, only the working order goes on the order book.
// The behavior of the working order is the same as the OTO.
// OTOCO has 2 pending orders (pending above and pending below), forming an OCO pair. The pending orders are only placed on the order book when the working order gets fully filled.
// OTOCOs add 3 orders to the unfilled order count, EXCHANGE_MAX_NUM_ORDERS filter, and MAX_NUM_ORDERS filter.
type WsCreateOTOCOOrder struct {
	c *WsClient
	r *core.WsRequest
}

func (s *WsCreateOTOCOOrder) Symbol(symbol string) *WsCreateOTOCOOrder {
	s.r.Set("symbol", symbol)
	return s
}

// ListClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent.
// A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.
// listClientOrderId is distinct from the workingClientOrderId, pendingAboveClientOrderId, and the pendingBelowClientOrderId.
func (s *WsCreateOTOCOOrder) ListClientOrderId(listClientOrderId string) *WsCreateOTOCOOrder {
	s.r.Set("listClientOrderId", listClientOrderId)
	return s
}
func (s *WsCreateOTOCOOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateOTOCOOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCreateOTOCOOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateOTOCOOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

// WorkingType Supported values: LIMIT, LIMIT_MAKER
func (s *WsCreateOTOCOOrder) WorkingType(workingType core.OrderTypeEnum) *WsCreateOTOCOOrder {
	s.r.Set("workingType", workingType)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingSide(workingSide core.OrderSideEnum) *WsCreateOTOCOOrder {
	s.r.Set("workingSide", workingSide)
	return s
}

// WorkingClientOrderId Arbitrary unique ID among open orders for the working order.
// Automatically generated if not sent.
func (s *WsCreateOTOCOOrder) WorkingClientOrderId(workingClientOrderId string) *WsCreateOTOCOOrder {
	s.r.Set("workingClientOrderId", workingClientOrderId)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingPrice(workingPrice string) *WsCreateOTOCOOrder {
	s.r.Set("workingPrice", workingPrice)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingQuantity(workingQuantity string) *WsCreateOTOCOOrder {
	s.r.Set("workingQuantity", workingQuantity)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingIcebergQty(workingIcebergQty string) *WsCreateOTOCOOrder {
	s.r.Set("workingIcebergQty", workingIcebergQty)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingTimeInForce(workingTimeInForce core.TimeInForceEnum) *WsCreateOTOCOOrder {
	s.r.Set("workingTimeInForce", workingTimeInForce)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingStrategyId(workingStrategyId int64) *WsCreateOTOCOOrder {
	s.r.Set("workingStrategyId", workingStrategyId)
	return s
}
func (s *WsCreateOTOCOOrder) WorkingStrategyType(workingStrategyType int) *WsCreateOTOCOOrder {
	s.r.Set("workingStrategyType", workingStrategyType)
	return s
}
func (s *WsCreateOTOCOOrder) PendingSide(pendingSide core.OrderSideEnum) *WsCreateOTOCOOrder {
	s.r.Set("pendingSide", pendingSide)
	return s
}
func (s *WsCreateOTOCOOrder) PendingQuantity(pendingQuantity string) *WsCreateOTOCOOrder {
	s.r.Set("pendingQuantity", pendingQuantity)
	return s
}

// PendingAboveType Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
func (s *WsCreateOTOCOOrder) PendingAboveType(pendingAboveType core.OrderTypeEnum) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveType", pendingAboveType)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveClientOrderId(pendingAboveClientOrderId string) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveClientOrderId", pendingAboveClientOrderId)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAbovePrice(pendingAbovePrice string) *WsCreateOTOCOOrder {
	s.r.Set("pendingAbovePrice", pendingAbovePrice)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveStopPrice(pendingAboveStopPrice string) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveStopPrice", pendingAboveStopPrice)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveTrailingDelta(pendingAboveTrailingDelta string) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveTrailingDelta", pendingAboveTrailingDelta)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveIcebergQty(pendingAboveIcebergQty string) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveIcebergQty", pendingAboveIcebergQty)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveTimeInForce(pendingAboveTimeInForce core.TimeInForceEnum) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveTimeInForce", pendingAboveTimeInForce)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveStrategyId(pendingAboveStrategyId int64) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveStrategyId", pendingAboveStrategyId)
	return s
}
func (s *WsCreateOTOCOOrder) PendingAboveStrategyType(pendingAboveStrategyType int) *WsCreateOTOCOOrder {
	s.r.Set("pendingAboveStrategyType", pendingAboveStrategyType)
	return s
}

// PendingBelowType Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
func (s *WsCreateOTOCOOrder) PendingBelowType(pendingBelowType core.OrderTypeEnum) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowType", pendingBelowType)
	return s
}
func (s *WsCreateOTOCOOrder) PendingBelowClientOrderId(pendingBelowClientOrderId string) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowClientOrderId", pendingBelowClientOrderId)
	return s
}

// PendingBelowPrice Can be used if pendingBelowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT to specify limit price
func (s *WsCreateOTOCOOrder) PendingBelowPrice(pendingBelowPrice string) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowPrice", pendingBelowPrice)
	return s
}

// PendingBelowStopPrice Can be used if pendingBelowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT.
func (s *WsCreateOTOCOOrder) PendingBelowStopPrice(pendingBelowStopPrice string) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowStopPrice", pendingBelowStopPrice)
	return s
}
func (s *WsCreateOTOCOOrder) PendingBelowTrailingDelta(pendingBelowTrailingDelta string) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowTrailingDelta", pendingBelowTrailingDelta)
	return s
}

// PendingBelowIcebergQty This can only be used if pendingBelowTimeInForce is GTC, or if pendingBelowType is LIMIT_MAKER.
func (s *WsCreateOTOCOOrder) PendingBelowIcebergQty(pendingBelowIcebergQty string) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowIcebergQty", pendingBelowIcebergQty)
	return s
}
func (s *WsCreateOTOCOOrder) PendingBelowTimeInForce(pendingBelowTimeInForce core.TimeInForceEnum) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowTimeInForce", pendingBelowTimeInForce)
	return s
}
func (s *WsCreateOTOCOOrder) PendingBelowStrategyId(pendingBelowStrategyId int64) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowStrategyId", pendingBelowStrategyId)
	return s
}
func (s *WsCreateOTOCOOrder) PendingBelowStrategyType(pendingBelowStrategyType int) *WsCreateOTOCOOrder {
	s.r.Set("pendingBelowStrategyType", pendingBelowStrategyType)
	return s
}

func (s *WsCreateOTOCOOrder) RecvWindow(recvWindow int) *WsCreateOTOCOOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateOTOCOOrder) Do(ctx context.Context) (*OrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *OrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsQueryOrderList Check execution status of an Order list.
type WsQueryOrderList struct {
	c *WsClient
	r *core.WsRequest
}

type OrderListResult struct {
	OrderListId       int               `json:"orderListId"`
	ContingencyType   string            `json:"contingencyType"`
	ListStatusType    string            `json:"listStatusType"`
	ListOrderStatus   string            `json:"listOrderStatus"`
	ListClientOrderId string            `json:"listClientOrderId"`
	TransactionTime   int64             `json:"transactionTime"`
	Symbol            string            `json:"symbol"`
	Orders            []*ApiOrder       `json:"orders"`
	OrderReports      []*ApiOrderReport `json:"orderReports"`
}

type WsOrderListResponse struct {
	ApiResponse
	Result *OrderListResult `json:"result"`
}

func (s *WsQueryOrderList) OrigClientOrderId(origClientOrderId string) *WsQueryOrderList {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}
func (s *WsQueryOrderList) OrderListId(orderListId int64) *WsQueryOrderList {
	s.r.Set("orderListId", orderListId)
	return s
}
func (s *WsQueryOrderList) RecvWindow(recvWindow int) *WsQueryOrderList {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *WsQueryOrderList) Do(ctx context.Context) (*WsOrderListResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsOrderListResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCancelOrderList Cancel an active order list.
type WsCancelOrderList struct {
	c *WsClient
	r *core.WsRequest
}

func (s *WsCancelOrderList) Symbol(symbol int) *WsCancelOrderList {
	s.r.Set("symbol", symbol)
	return s
}

// OrderListId Cancel order list by orderListId
func (s *WsCancelOrderList) OrderListId(orderListId int64) *WsCancelOrderList {
	s.r.Set("orderListId", orderListId)
	return s
}

// ListClientOrderId Cancel order list by listClientId
func (s *WsCancelOrderList) ListClientOrderId(listClientOrderId string) *WsCancelOrderList {
	s.r.Set("listClientOrderId", listClientOrderId)
	return s
}

// NewClientOrderId New ID for the canceled order list. Automatically generated if not sent
func (s *WsCancelOrderList) NewClientOrderId(newClientOrderId string) *WsCancelOrderList {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *WsCancelOrderList) RecvWindow(recvWindow int) *WsCancelOrderList {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *WsCancelOrderList) Do(ctx context.Context) (*WsOrderListResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsOrderListResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsQueryOpenOrder Query execution status of all open order lists.
type WsQueryOpenOrder struct {
	c *WsClient
	r *core.WsRequest
}

type WsQueryOpenOrderResponse struct {
	ApiResponse
	Result []*OrderListResult `json:"result"`
}

func (s *WsQueryOpenOrder) RecvWindow(recvWindow int) *WsQueryOpenOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *WsQueryOpenOrder) Do(ctx context.Context) (*WsQueryOpenOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsQueryOpenOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCreateSOROrder Places an order using smart order routing (SOR).
type WsCreateSOROrder struct {
	c *WsClient
	r *core.WsRequest
}

type CreateSOROrderResult struct {
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
	Fills                   []*ApiFill      `json:"fills"`
	WorkingFloor            string          `json:"workingFloor"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	UsedSor                 bool            `json:"usedSor"`
}

type WsCreateSOROrderResponse struct {
	ApiResponse
	Result []*CreateSOROrderResult `json:"result"`
}

func (s *WsCreateSOROrder) Symbol(symbol string) *WsCreateSOROrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsCreateSOROrder) Side(side core.OrderSideEnum) *WsCreateSOROrder {
	s.r.Set("side", side)
	return s
}
func (s *WsCreateSOROrder) Type(type_ core.OrderTypeEnum) *WsCreateSOROrder {
	s.r.Set("type", type_)
	return s
}

// TimeInForce Applicable only to LIMIT order type
func (s *WsCreateSOROrder) TimeInForce(timeInForce core.TimeInForceEnum) *WsCreateSOROrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}

// Price Applicable only to LIMIT order type
func (s *WsCreateSOROrder) Price(price string) *WsCreateSOROrder {
	s.r.Set("price", price)
	return s
}
func (s *WsCreateSOROrder) Quantity(quantity string) *WsCreateSOROrder {
	s.r.Set("quantity", quantity)
	return s
}

// NewClientOrderId Arbitrary unique ID among open orders. Automatically generated if not sent
func (s *WsCreateSOROrder) NewClientOrderId(newClientOrderId string) *WsCreateSOROrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

// NewOrderRespType Select response format: ACK, RESULT, FULL.
// MARKET and LIMIT orders use FULL by default.
func (s *WsCreateSOROrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateSOROrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCreateSOROrder) IcebergQty(icebergQty string) *WsCreateSOROrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}
func (s *WsCreateSOROrder) StrategyId(strategyId int64) *WsCreateSOROrder {
	s.r.Set("strategyId", strategyId)
	return s
}
func (s *WsCreateSOROrder) StrategyType(strategyType int) *WsCreateSOROrder {
	s.r.Set("strategyType", strategyType)
	return s
}
func (s *WsCreateSOROrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateSOROrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *WsCreateSOROrder) RecvWindow(recvWindow int) *WsCreateSOROrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateSOROrder) Do(ctx context.Context) (*WsCreateSOROrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCreateSOROrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsCreateTestSOROrder Test new order creation and signature/recvWindow using smart order routing (SOR). Creates and validates a new order but does not send it into the matching engine.
type WsCreateTestSOROrder struct {
	c *WsClient
	r *core.WsRequest
}

type WsCreateTestSOROrderResponse struct {
	ApiResponse
	Result *CreateTestOrderResult `json:"result"`
}

func (s *WsCreateTestSOROrder) Symbol(symbol string) *WsCreateTestSOROrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsCreateTestSOROrder) Side(side core.OrderSideEnum) *WsCreateTestSOROrder {
	s.r.Set("side", side)
	return s
}
func (s *WsCreateTestSOROrder) Type(type_ core.OrderTypeEnum) *WsCreateTestSOROrder {
	s.r.Set("type", type_)
	return s
}
func (s *WsCreateTestSOROrder) TimeInForce(timeInForce core.TimeInForceEnum) *WsCreateTestSOROrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}
func (s *WsCreateTestSOROrder) Price(price string) *WsCreateTestSOROrder {
	s.r.Set("price", price)
	return s
}
func (s *WsCreateTestSOROrder) Quantity(quantity string) *WsCreateTestSOROrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *WsCreateTestSOROrder) NewClientOrderId(newClientOrderId string) *WsCreateTestSOROrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *WsCreateTestSOROrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateTestSOROrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *WsCreateTestSOROrder) IcebergQty(icebergQty string) *WsCreateTestSOROrder {
	s.r.Set("icebergQty", icebergQty)
	return s
}
func (s *WsCreateTestSOROrder) StrategyId(strategyId int64) *WsCreateTestSOROrder {
	s.r.Set("strategyId", strategyId)
	return s
}
func (s *WsCreateTestSOROrder) StrategyType(strategyType int) *WsCreateTestSOROrder {
	s.r.Set("strategyType", strategyType)
	return s
}
func (s *WsCreateTestSOROrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateTestSOROrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *WsCreateTestSOROrder) RecvWindow(recvWindow int) *WsCreateTestSOROrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateTestSOROrder) ComputeCommissionRates(computeCommissionRates bool) *WsCreateTestSOROrder {
	s.r.Set("computeCommissionRates", computeCommissionRates)
	return s
}

func (s *WsCreateTestSOROrder) Do(ctx context.Context) (*WsCreateTestSOROrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsCreateTestSOROrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
