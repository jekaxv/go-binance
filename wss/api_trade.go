package wss

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
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

// CreateOrder Send in a new order.
type CreateOrder struct {
	c *Client
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

type CreateOrderResponse struct {
	ApiResponse
	Result *CreateOrderResult `json:"result"`
}

func (s *CreateOrder) Symbol(symbol string) *CreateOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *CreateOrder) Side(side types.OrderSideEnum) *CreateOrder {
	s.c.req.Params["side"] = side
	return s
}
func (s *CreateOrder) Type(orderType types.OrderTypeEnum) *CreateOrder {
	s.c.req.Params["type"] = orderType
	return s
}
func (s *CreateOrder) TimeInForce(timeInForce types.TimeInForceEnum) *CreateOrder {
	s.c.req.Params["timeInForce"] = timeInForce
	return s
}
func (s *CreateOrder) Price(price string) *CreateOrder {
	s.c.req.Params["price"] = price
	return s
}
func (s *CreateOrder) Quantity(quantity string) *CreateOrder {
	s.c.req.Params["quantity"] = quantity
	return s
}
func (s *CreateOrder) QuoteOrderQty(quoteOrderQty string) *CreateOrder {
	s.c.req.Params["quoteOrderQty"] = quoteOrderQty
	return s
}
func (s *CreateOrder) NewClientOrderId(newClientOrderId string) *CreateOrder {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}
func (s *CreateOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateOrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CreateOrder) StopPrice(stopPrice string) *CreateOrder {
	s.c.req.Params["stopPrice"] = stopPrice
	return s
}
func (s *CreateOrder) TrailingDelta(trailingDelta int) *CreateOrder {
	s.c.req.Params["trailingDelta"] = trailingDelta
	return s
}
func (s *CreateOrder) IcebergQty(icebergQty string) *CreateOrder {
	s.c.req.Params["icebergQty"] = icebergQty
	return s
}
func (s *CreateOrder) StrategyId(strategyId int64) *CreateOrder {
	s.c.req.Params["strategyId"] = strategyId
	return s
}
func (s *CreateOrder) StrategyType(strategyType int) *CreateOrder {
	s.c.req.Params["strategyType"] = strategyType
	return s
}
func (s *CreateOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateOrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

func (s *CreateOrder) RecvWindow(recvWindow int) *CreateOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateOrder) Do(ctx context.Context) (*CreateOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CreateOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

type CreateTestOrder struct {
	c *Client
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

type CreateTestOrderResult struct {
	StandardCommissionForOrder *OrderCommission `json:"standardCommissionForOrder"`
	TaxCommissionForOrder      *OrderCommission `json:"taxCommissionForOrder"`
	Discount                   *Discount        `json:"discount"`
}

type CreateOrderTestResponse struct {
	ApiResponse
	Result *CreateTestOrderResult `json:"result"`
}

func (s *CreateTestOrder) Symbol(symbol string) *CreateTestOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *CreateTestOrder) Side(side types.OrderSideEnum) *CreateTestOrder {
	s.c.req.Params["side"] = side
	return s
}
func (s *CreateTestOrder) Type(orderType types.OrderTypeEnum) *CreateTestOrder {
	s.c.req.Params["type"] = orderType
	return s
}
func (s *CreateTestOrder) TimeInForce(timeInForce types.TimeInForceEnum) *CreateTestOrder {
	s.c.req.Params["timeInForce"] = timeInForce
	return s
}
func (s *CreateTestOrder) Price(price string) *CreateTestOrder {
	s.c.req.Params["price"] = price
	return s
}
func (s *CreateTestOrder) Quantity(quantity string) *CreateTestOrder {
	s.c.req.Params["quantity"] = quantity
	return s
}
func (s *CreateTestOrder) QuoteOrderQty(quoteOrderQty string) *CreateTestOrder {
	s.c.req.Params["quoteOrderQty"] = quoteOrderQty
	return s
}
func (s *CreateTestOrder) NewClientOrderId(newClientOrderId string) *CreateTestOrder {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}
func (s *CreateTestOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateTestOrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CreateTestOrder) StopPrice(stopPrice string) *CreateTestOrder {
	s.c.req.Params["stopPrice"] = stopPrice
	return s
}
func (s *CreateTestOrder) TrailingDelta(trailingDelta int) *CreateTestOrder {
	s.c.req.Params["trailingDelta"] = trailingDelta
	return s
}
func (s *CreateTestOrder) IcebergQty(icebergQty string) *CreateTestOrder {
	s.c.req.Params["icebergQty"] = icebergQty
	return s
}
func (s *CreateTestOrder) StrategyId(strategyId int64) *CreateTestOrder {
	s.c.req.Params["strategyId"] = strategyId
	return s
}
func (s *CreateTestOrder) StrategyType(strategyType int) *CreateTestOrder {
	s.c.req.Params["strategyType"] = strategyType
	return s
}
func (s *CreateTestOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateTestOrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

func (s *CreateTestOrder) RecvWindow(recvWindow int) *CreateTestOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateTestOrder) ComputeCommissionRates(computeCommissionRates bool) *CreateTestOrder {
	s.c.req.Params["computeCommissionRates"] = computeCommissionRates
	return s
}

func (s *CreateTestOrder) Do(ctx context.Context) (*CreateOrderTestResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CreateOrderTestResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// QueryOrder Check execution status of an order.
type QueryOrder struct {
	c *Client
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

type QueryOrderResponse struct {
	ApiResponse
	Result *QueryOrderResult `json:"result"`
}

func (s *QueryOrder) Symbol(symbol string) *QueryOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *QueryOrder) OrderId(orderId int64) *QueryOrder {
	s.c.req.Params["orderId"] = orderId
	return s
}
func (s *QueryOrder) OrigClientOrderId(origClientOrderId string) *QueryOrder {
	s.c.req.Params["origClientOrderId"] = origClientOrderId
	return s
}

func (s *QueryOrder) RecvWindow(recvWindow int) *QueryOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *QueryOrder) Do(ctx context.Context) (*QueryOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *QueryOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CancelOrder Cancel an active order.
type CancelOrder struct {
	c *Client
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

type CancelOrderResponse struct {
	ApiResponse
	Result *CancelOrderResult `json:"result"`
}

func (s *CancelOrder) Symbol(symbol string) *CancelOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *CancelOrder) OrderId(orderId int64) *CancelOrder {
	s.c.req.Params["orderId"] = orderId
	return s
}

func (s *CancelOrder) OrigClientOrderId(origClientOrderId string) *CancelOrder {
	s.c.req.Params["origClientOrderId"] = origClientOrderId
	return s
}

func (s *CancelOrder) NewClientOrderId(newClientOrderId string) *CancelOrder {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}

// CancelRestrictions Supported values:
// ONLY_NEW - Cancel will succeed if the order status is NEW.
// ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED.
func (s *CancelOrder) CancelRestrictions(cancelRestrictions types.CancelRestrictionEnum) *CancelOrder {
	s.c.req.Params["cancelRestrictions"] = cancelRestrictions
	return s
}

func (s *CancelOrder) RecvWindow(recvWindow int) *CancelOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CancelOrder) Do(ctx context.Context) (*CancelOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CancelOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CancelReplaceOrder Cancel an existing order and immediately place a new order instead of the canceled one.
type CancelReplaceOrder struct {
	c *Client
}
type CancelReplaceOrderResult struct {
	CancelResult     string          `json:"cancelResult"`
	NewOrderResult   string          `json:"newOrderResult"`
	CancelResponse   *ApiOrderReport `json:"cancelResponse"`
	NewOrderResponse *ApiOrderReport `json:"newOrderResponse"`
}
type CancelReplaceOrderResponse struct {
	Id         string                    `json:"id"`
	Status     int                       `json:"status"`
	RateLimits []*ApiRateLimit           `json:"rateLimits,omitempty"`
	Error      *ApiOrderError            `json:"error,omitempty"`
	Result     *CancelReplaceOrderResult `json:"result"`
}

func (s *CancelReplaceOrder) Symbol(symbol string) *CancelReplaceOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}

// CancelReplaceMode Available cancelReplaceMode options:
// STOP_ON_FAILURE – if cancellation request fails, new order placement will not be attempted.
// ALLOW_FAILURE – new order placement will be attempted even if the cancel request fails.
func (s *CancelReplaceOrder) CancelReplaceMode(cancelReplaceMode types.CancelReplaceModeEnum) *CancelReplaceOrder {
	s.c.req.Params["cancelReplaceMode"] = cancelReplaceMode
	return s
}

// CancelOrderId Cancel order by orderId
func (s *CancelReplaceOrder) CancelOrderId(cancelOrderId int64) *CancelReplaceOrder {
	s.c.req.Params["cancelOrderId"] = cancelOrderId
	return s
}

// CancelOrigClientOrderId Cancel order by clientOrderId
func (s *CancelReplaceOrder) CancelOrigClientOrderId(cancelOrigClientOrderId string) *CancelReplaceOrder {
	s.c.req.Params["cancelOrigClientOrderId"] = cancelOrigClientOrderId
	return s
}

// CancelNewClientOrderId New ID for the canceled order. Automatically generated if not sent
func (s *CancelReplaceOrder) CancelNewClientOrderId(cancelNewClientOrderId string) *CancelReplaceOrder {
	s.c.req.Params["cancelNewClientOrderId"] = cancelNewClientOrderId
	return s
}
func (s *CancelReplaceOrder) Side(side types.OrderSideEnum) *CancelReplaceOrder {
	s.c.req.Params["side"] = side
	return s
}
func (s *CancelReplaceOrder) Type(type_ types.OrderTypeEnum) *CancelReplaceOrder {
	s.c.req.Params["type"] = type_
	return s
}
func (s *CancelReplaceOrder) TimeInForce(timeInForce types.TimeInForceEnum) *CancelReplaceOrder {
	s.c.req.Params["timeInForce"] = timeInForce
	return s
}
func (s *CancelReplaceOrder) Price(price string) *CancelReplaceOrder {
	s.c.req.Params["price"] = price
	return s
}
func (s *CancelReplaceOrder) Quantity(quantity string) *CancelReplaceOrder {
	s.c.req.Params["quantity"] = quantity
	return s
}
func (s *CancelReplaceOrder) QuoteOrderQty(quoteOrderQty string) *CancelReplaceOrder {
	s.c.req.Params["quoteOrderQty"] = quoteOrderQty
	return s
}

// NewClientOrderId Arbitrary unique ID among open orders. Automatically generated if not sent
func (s *CancelReplaceOrder) NewClientOrderId(newClientOrderId string) *CancelReplaceOrder {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}

// NewOrderRespType Select response format: ACK, RESULT, FULL.
// MARKET and LIMIT orders produce FULL response by default, other order types default to ACK.
func (s *CancelReplaceOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CancelReplaceOrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CancelReplaceOrder) StopPrice(stopPrice string) *CancelReplaceOrder {
	s.c.req.Params["stopPrice"] = stopPrice
	return s
}
func (s *CancelReplaceOrder) TrailingDelta(trailingDelta string) *CancelReplaceOrder {
	s.c.req.Params["trailingDelta"] = trailingDelta
	return s
}
func (s *CancelReplaceOrder) IcebergQty(icebergQty string) *CancelReplaceOrder {
	s.c.req.Params["icebergQty"] = icebergQty
	return s
}

// StrategyId Arbitrary numeric value identifying the order within an order strategy.
func (s *CancelReplaceOrder) StrategyId(strategyId int64) *CancelReplaceOrder {
	s.c.req.Params["strategyId"] = strategyId
	return s
}

// StrategyType Arbitrary numeric value identifying the order strategy.
// Values smaller than 1000000 are reserved and cannot be used.
func (s *CancelReplaceOrder) StrategyType(strategyType int) *CancelReplaceOrder {
	s.c.req.Params["strategyType"] = strategyType
	return s
}
func (s *CancelReplaceOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CancelReplaceOrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

// CancelRestrictions Supported values:
// ONLY_NEW - Cancel will succeed if the order status is NEW.
// ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED. For more information please refer to Regarding cancelRestrictions.
func (s *CancelReplaceOrder) CancelRestrictions(cancelRestrictions types.CancelRestrictionEnum) *CancelReplaceOrder {
	s.c.req.Params["cancelRestrictions"] = cancelRestrictions
	return s
}

// OrderRateLimitExceededMode Supported values:
// DO_NOTHING (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit
// CANCEL_ONLY - will always cancel the order.
func (s *CancelReplaceOrder) OrderRateLimitExceededMode(orderRateLimitExceededMode types.OrderExceededModeEnum) *CancelReplaceOrder {
	s.c.req.Params["orderRateLimitExceededMode"] = orderRateLimitExceededMode
	return s
}

func (s *CancelReplaceOrder) RecvWindow(recvWindow int) *CancelReplaceOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CancelReplaceOrder) Do(ctx context.Context) (*CancelReplaceOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CancelReplaceOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// OpenOrdersStatus Query execution status of all open orders.
type OpenOrdersStatus struct {
	c *Client
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
type OpenOrdersStatusResponse struct {
	ApiResponse
	Result []*OpenOrdersStatusResult `json:"result"`
}

func (s *OpenOrdersStatus) Symbol(symbol string) *OpenOrdersStatus {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *OpenOrdersStatus) RecvWindow(recvWindow int) *OpenOrdersStatus {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}
func (s *OpenOrdersStatus) Do(ctx context.Context) (*OpenOrdersStatusResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *OpenOrdersStatusResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CancelOpenOrder Cancel all open orders on a symbol. This includes orders that are part of an order list.
type CancelOpenOrder struct {
	c *Client
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

type CancelOpenOrderResponse struct {
	ApiResponse
	Result []*OrderResult `json:"result"`
}

func (s *CancelOpenOrder) Symbol(symbol string) *CancelOpenOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}

func (s *CancelOpenOrder) RecvWindow(recvWindow int) *CancelOpenOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CancelOpenOrder) Do(ctx context.Context) (*CancelOpenOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CancelOpenOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CreateOCOOrder Send in an one-cancels the other (OCO) pair, where activation of one order immediately cancels the other.
type CreateOCOOrder struct {
	c *Client
}

func (s *CreateOCOOrder) Symbol(symbol string) *CreateOCOOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}

// ListClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent.
// A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.
// listClientOrderId is distinct from the aboveClientOrderId and the belowCLientOrderId
func (s *CreateOCOOrder) ListClientOrderId(listClientOrderId string) *CreateOCOOrder {
	s.c.req.Params["listClientOrderId"] = listClientOrderId
	return s
}
func (s *CreateOCOOrder) Side(side types.OrderSideEnum) *CreateOCOOrder {
	s.c.req.Params["side"] = side
	return s
}

// Quantity for both orders of the order list.
func (s *CreateOCOOrder) Quantity(quantity string) *CreateOCOOrder {
	s.c.req.Params["quantity"] = quantity
	return s
}

// AboveType Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
func (s *CreateOCOOrder) AboveType(aboveType types.OrderTypeEnum) *CreateOCOOrder {
	s.c.req.Params["aboveType"] = aboveType
	return s
}

// AboveClientOrderId Arbitrary unique ID among open orders for the above order. Automatically generated if not sent
func (s *CreateOCOOrder) AboveClientOrderId(aboveClientOrderId string) *CreateOCOOrder {
	s.c.req.Params["aboveClientOrderId"] = aboveClientOrderId
	return s
}

// AboveIcebergQty Note that this can only be used if aboveTimeInForce is GTC.
func (s *CreateOCOOrder) AboveIcebergQty(aboveIcebergQty int64) *CreateOCOOrder {
	s.c.req.Params["aboveIcebergQty"] = aboveIcebergQty
	return s
}

// AbovePrice Can be used if aboveType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
func (s *CreateOCOOrder) AbovePrice(abovePrice string) *CreateOCOOrder {
	s.c.req.Params["abovePrice"] = abovePrice
	return s
}

// AboveStopPrice Can be used if aboveType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, TAKE_PROFIT_LIMIT
// Either aboveStopPrice or aboveTrailingDelta or both, must be specified.
func (s *CreateOCOOrder) AboveStopPrice(aboveStopPrice string) *CreateOCOOrder {
	s.c.req.Params["aboveStopPrice"] = aboveStopPrice
	return s
}
func (s *CreateOCOOrder) AboveTrailingDelta(aboveTrailingDelta int64) *CreateOCOOrder {
	s.c.req.Params["aboveTrailingDelta"] = aboveTrailingDelta
	return s
}

// AboveTimeInForce Required if aboveType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT.
func (s *CreateOCOOrder) AboveTimeInForce(aboveTimeInForce types.TimeInForceEnum) *CreateOCOOrder {
	s.c.req.Params["aboveTimeInForce"] = aboveTimeInForce
	return s
}

// AboveStrategyId Arbitrary numeric value identifying the above order within an order strategy.
func (s *CreateOCOOrder) AboveStrategyId(aboveStrategyId int64) *CreateOCOOrder {
	s.c.req.Params["aboveStrategyId"] = aboveStrategyId
	return s
}

// AboveStrategyType Arbitrary numeric value identifying the above order strategy.
// Values smaller than 1000000 are reserved and cannot be used.
func (s *CreateOCOOrder) AboveStrategyType(aboveStrategyType int64) *CreateOCOOrder {
	s.c.req.Params["aboveStrategyType"] = aboveStrategyType
	return s
}

// BelowType Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
func (s *CreateOCOOrder) BelowType(belowType types.OrderTypeEnum) *CreateOCOOrder {
	s.c.req.Params["belowType"] = belowType
	return s
}

func (s *CreateOCOOrder) BelowClientOrderId(belowClientOrderId string) *CreateOCOOrder {
	s.c.req.Params["belowClientOrderId"] = belowClientOrderId
	return s
}

// BelowIcebergQty Note that this can only be used if belowTimeInForce is GTC.
func (s *CreateOCOOrder) BelowIcebergQty(belowIcebergQty int64) *CreateOCOOrder {
	s.c.req.Params["belowIcebergQty"] = belowIcebergQty
	return s
}

// BelowPrice Can be used if belowType is STOP_LOSS_LIMIT , LIMIT_MAKER, or TAKE_PROFIT_LIMIT to specify the limit price.
func (s *CreateOCOOrder) BelowPrice(belowPrice string) *CreateOCOOrder {
	s.c.req.Params["belowPrice"] = belowPrice
	return s
}

// BelowStopPrice Can be used if belowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT.
func (s *CreateOCOOrder) BelowStopPrice(belowStopPrice string) *CreateOCOOrder {
	s.c.req.Params["belowStopPrice"] = belowStopPrice
	return s
}
func (s *CreateOCOOrder) BelowTrailingDelta(belowTrailingDelta int64) *CreateOCOOrder {
	s.c.req.Params["belowTrailingDelta"] = belowTrailingDelta
	return s
}

// BelowTimeInForce Required if belowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT
func (s *CreateOCOOrder) BelowTimeInForce(belowTimeInForce string) *CreateOCOOrder {
	s.c.req.Params["belowTimeInForce"] = belowTimeInForce
	return s
}
func (s *CreateOCOOrder) BelowStrategyId(belowStrategyId int64) *CreateOCOOrder {
	s.c.req.Params["belowStrategyId"] = belowStrategyId
	return s
}
func (s *CreateOCOOrder) BelowStrategyType(belowStrategyType int64) *CreateOCOOrder {
	s.c.req.Params["belowStrategyType"] = belowStrategyType
	return s
}

// NewOrderRespType Select response format: ACK, RESULT, FULL
func (s *CreateOCOOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateOCOOrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CreateOCOOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateOCOOrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}
func (s *CreateOCOOrder) RecvWindow(recvWindow int) *CreateOCOOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateOCOOrder) Do(ctx context.Context) (*OrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
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

// CreateOTOOrder Places an OTO.
// An OTO (One-Triggers-the-Other) is an order list comprised of 2 orders.
// The first order is called the working order and must be LIMIT or LIMIT_MAKER. Initially, only the working order goes on the order book.
// The second order is called the pending order. It can be any order type except for MARKET orders using parameter quoteOrderQty. The pending order is only placed on the order book when the working order gets fully filled.
// If either the working order or the pending order is cancelled individually, the other order in the order list will also be canceled or expired.
// OTOs add 2 orders to the unfilled order count, EXCHANGE_MAX_NUM_ORDERS filter and MAX_NUM_ORDERS filter.
type CreateOTOOrder struct {
	c *Client
}

func (s *CreateOTOOrder) Symbol(symbol string) *CreateOTOOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}

// ListClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent.
// A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.
// listClientOrderId is distinct from the workingClientOrderId and the pendingClientOrderId.
func (s *CreateOTOOrder) ListClientOrderId(listClientOrderId string) *CreateOTOOrder {
	s.c.req.Params["listClientOrderId"] = listClientOrderId
	return s
}

func (s *CreateOTOOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateOTOOrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}

func (s *CreateOTOOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateOTOOrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

// WorkingType Supported values: LIMIT,LIMIT_MAKER
func (s *CreateOTOOrder) WorkingType(workingType types.OrderTypeEnum) *CreateOTOOrder {
	s.c.req.Params["workingType"] = workingType
	return s
}

func (s *CreateOTOOrder) WorkingSide(workingSide types.OrderSideEnum) *CreateOTOOrder {
	s.c.req.Params["workingSide"] = workingSide
	return s
}

// WorkingClientOrderId Arbitrary unique ID among open orders for the working order.
// Automatically generated if not sent.
func (s *CreateOTOOrder) WorkingClientOrderId(workingClientOrderId string) *CreateOTOOrder {
	s.c.req.Params["workingClientOrderId"] = workingClientOrderId
	return s
}
func (s *CreateOTOOrder) WorkingPrice(workingPrice string) *CreateOTOOrder {
	s.c.req.Params["workingPrice"] = workingPrice
	return s
}

// WorkingQuantity Sets the quantity for the working order.
func (s *CreateOTOOrder) WorkingQuantity(workingQuantity string) *CreateOTOOrder {
	s.c.req.Params["workingQuantity"] = workingQuantity
	return s
}

// WorkingIcebergQty This can only be used if workingTimeInForce is GTC, or if workingType is LIMIT_MAKER.
func (s *CreateOTOOrder) WorkingIcebergQty(workingIcebergQty string) *CreateOTOOrder {
	s.c.req.Params["workingIcebergQty"] = workingIcebergQty
	return s
}
func (s *CreateOTOOrder) WorkingTimeInForce(workingTimeInForce types.TimeInForceEnum) *CreateOTOOrder {
	s.c.req.Params["workingTimeInForce"] = workingTimeInForce
	return s
}

// WorkingStrategyId	Arbitrary numeric value identifying the working order within an order strategy.
func (s *CreateOTOOrder) WorkingStrategyId(workingStrategyId int64) *CreateOTOOrder {
	s.c.req.Params["workingStrategyId"] = workingStrategyId
	return s
}
func (s *CreateOTOOrder) WorkingStrategyType(workingStrategyType int) *CreateOTOOrder {
	s.c.req.Params["workingStrategyType"] = workingStrategyType
	return s
}

// PendingType Supported values: Order Types Note that MARKET orders using quoteOrderQty are not supported.
func (s *CreateOTOOrder) PendingType(pendingType types.OrderTypeEnum) *CreateOTOOrder {
	s.c.req.Params["pendingType"] = pendingType
	return s
}
func (s *CreateOTOOrder) PendingSide(pendingSide types.OrderSideEnum) *CreateOTOOrder {
	s.c.req.Params["pendingSide"] = pendingSide
	return s
}

// PendingClientOrderId Arbitrary unique ID among open orders for the pending order.
// Automatically generated if not sent.
func (s *CreateOTOOrder) PendingClientOrderId(pendingClientOrderId string) *CreateOTOOrder {
	s.c.req.Params["pendingClientOrderId"] = pendingClientOrderId
	return s
}
func (s *CreateOTOOrder) PendingPrice(pendingPrice string) *CreateOTOOrder {
	s.c.req.Params["pendingPrice"] = pendingPrice
	return s
}
func (s *CreateOTOOrder) PendingStopPrice(pendingStopPrice string) *CreateOTOOrder {
	s.c.req.Params["pendingStopPrice"] = pendingStopPrice
	return s
}
func (s *CreateOTOOrder) PendingTrailingDelta(pendingTrailingDelta string) *CreateOTOOrder {
	s.c.req.Params["pendingTrailingDelta"] = pendingTrailingDelta
	return s
}
func (s *CreateOTOOrder) PendingQuantity(pendingQuantity string) *CreateOTOOrder {
	s.c.req.Params["pendingQuantity"] = pendingQuantity
	return s
}
func (s *CreateOTOOrder) PendingIcebergQty(pendingIcebergQty string) *CreateOTOOrder {
	s.c.req.Params["pendingIcebergQty"] = pendingIcebergQty
	return s
}
func (s *CreateOTOOrder) PendingTimeInForce(pendingTimeInForce types.TimeInForceEnum) *CreateOTOOrder {
	s.c.req.Params["pendingTimeInForce"] = pendingTimeInForce
	return s
}
func (s *CreateOTOOrder) PendingStrategyId(pendingStrategyId int64) *CreateOTOOrder {
	s.c.req.Params["pendingStrategyId"] = pendingStrategyId
	return s
}
func (s *CreateOTOOrder) PendingStrategyType(pendingStrategyType int) *CreateOTOOrder {
	s.c.req.Params["pendingStrategyType"] = pendingStrategyType
	return s
}

func (s *CreateOTOOrder) RecvWindow(recvWindow int) *CreateOTOOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateOTOOrder) Do(ctx context.Context) (*OrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
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

// CreateOTOCOOrder Place an OTOCO.
// An OTOCO (One-Triggers-One-Cancels-the-Other) is an order list comprised of 3 orders.
// The first order is called the working order and must be LIMIT or LIMIT_MAKER. Initially, only the working order goes on the order book.
// The behavior of the working order is the same as the OTO.
// OTOCO has 2 pending orders (pending above and pending below), forming an OCO pair. The pending orders are only placed on the order book when the working order gets fully filled.
// OTOCOs add 3 orders to the unfilled order count, EXCHANGE_MAX_NUM_ORDERS filter, and MAX_NUM_ORDERS filter.
type CreateOTOCOOrder struct {
	c *Client
}

func (s *CreateOTOCOOrder) Symbol(symbol string) *CreateOTOCOOrder {
	s.c.req.Params["symbol"] = symbol
	return s
}

// ListClientOrderId Arbitrary unique ID among open order lists. Automatically generated if not sent.
// A new order list with the same listClientOrderId is accepted only when the previous one is filled or completely expired.
// listClientOrderId is distinct from the workingClientOrderId, pendingAboveClientOrderId, and the pendingBelowClientOrderId.
func (s *CreateOTOCOOrder) ListClientOrderId(listClientOrderId string) *CreateOTOCOOrder {
	s.c.req.Params["listClientOrderId"] = listClientOrderId
	return s
}
func (s *CreateOTOCOOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateOTOCOOrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CreateOTOCOOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateOTOCOOrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

// WorkingType Supported values: LIMIT, LIMIT_MAKER
func (s *CreateOTOCOOrder) WorkingType(workingType types.OrderTypeEnum) *CreateOTOCOOrder {
	s.c.req.Params["workingType"] = workingType
	return s
}
func (s *CreateOTOCOOrder) WorkingSide(workingSide types.OrderSideEnum) *CreateOTOCOOrder {
	s.c.req.Params["workingSide"] = workingSide
	return s
}

// WorkingClientOrderId Arbitrary unique ID among open orders for the working order.
// Automatically generated if not sent.
func (s *CreateOTOCOOrder) WorkingClientOrderId(workingClientOrderId string) *CreateOTOCOOrder {
	s.c.req.Params["workingClientOrderId"] = workingClientOrderId
	return s
}
func (s *CreateOTOCOOrder) WorkingPrice(workingPrice string) *CreateOTOCOOrder {
	s.c.req.Params["workingPrice"] = workingPrice
	return s
}
func (s *CreateOTOCOOrder) WorkingQuantity(workingQuantity string) *CreateOTOCOOrder {
	s.c.req.Params["workingQuantity"] = workingQuantity
	return s
}
func (s *CreateOTOCOOrder) WorkingIcebergQty(workingIcebergQty string) *CreateOTOCOOrder {
	s.c.req.Params["workingIcebergQty"] = workingIcebergQty
	return s
}
func (s *CreateOTOCOOrder) WorkingTimeInForce(workingTimeInForce types.TimeInForceEnum) *CreateOTOCOOrder {
	s.c.req.Params["workingTimeInForce"] = workingTimeInForce
	return s
}
func (s *CreateOTOCOOrder) WorkingStrategyId(workingStrategyId int64) *CreateOTOCOOrder {
	s.c.req.Params["workingStrategyId"] = workingStrategyId
	return s
}
func (s *CreateOTOCOOrder) WorkingStrategyType(workingStrategyType int) *CreateOTOCOOrder {
	s.c.req.Params["workingStrategyType"] = workingStrategyType
	return s
}
func (s *CreateOTOCOOrder) PendingSide(pendingSide types.OrderSideEnum) *CreateOTOCOOrder {
	s.c.req.Params["pendingSide"] = pendingSide
	return s
}
func (s *CreateOTOCOOrder) PendingQuantity(pendingQuantity string) *CreateOTOCOOrder {
	s.c.req.Params["pendingQuantity"] = pendingQuantity
	return s
}

// PendingAboveType Supported values: STOP_LOSS_LIMIT, STOP_LOSS, LIMIT_MAKER, TAKE_PROFIT, TAKE_PROFIT_LIMIT
func (s *CreateOTOCOOrder) PendingAboveType(pendingAboveType types.OrderTypeEnum) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveType"] = pendingAboveType
	return s
}
func (s *CreateOTOCOOrder) PendingAboveClientOrderId(pendingAboveClientOrderId string) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveClientOrderId"] = pendingAboveClientOrderId
	return s
}
func (s *CreateOTOCOOrder) PendingAbovePrice(pendingAbovePrice string) *CreateOTOCOOrder {
	s.c.req.Params["pendingAbovePrice"] = pendingAbovePrice
	return s
}
func (s *CreateOTOCOOrder) PendingAboveStopPrice(pendingAboveStopPrice string) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveStopPrice"] = pendingAboveStopPrice
	return s
}
func (s *CreateOTOCOOrder) PendingAboveTrailingDelta(pendingAboveTrailingDelta string) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveTrailingDelta"] = pendingAboveTrailingDelta
	return s
}
func (s *CreateOTOCOOrder) PendingAboveIcebergQty(pendingAboveIcebergQty string) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveIcebergQty"] = pendingAboveIcebergQty
	return s
}
func (s *CreateOTOCOOrder) PendingAboveTimeInForce(pendingAboveTimeInForce types.TimeInForceEnum) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveTimeInForce"] = pendingAboveTimeInForce
	return s
}
func (s *CreateOTOCOOrder) PendingAboveStrategyId(pendingAboveStrategyId int64) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveStrategyId"] = pendingAboveStrategyId
	return s
}
func (s *CreateOTOCOOrder) PendingAboveStrategyType(pendingAboveStrategyType int) *CreateOTOCOOrder {
	s.c.req.Params["pendingAboveStrategyType"] = pendingAboveStrategyType
	return s
}

// PendingBelowType Supported values: STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT,TAKE_PROFIT_LIMIT
func (s *CreateOTOCOOrder) PendingBelowType(pendingBelowType types.OrderTypeEnum) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowType"] = pendingBelowType
	return s
}
func (s *CreateOTOCOOrder) PendingBelowClientOrderId(pendingBelowClientOrderId string) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowClientOrderId"] = pendingBelowClientOrderId
	return s
}

// PendingBelowPrice Can be used if pendingBelowType is STOP_LOSS_LIMIT or TAKE_PROFIT_LIMIT to specify limit price
func (s *CreateOTOCOOrder) PendingBelowPrice(pendingBelowPrice string) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowPrice"] = pendingBelowPrice
	return s
}

// PendingBelowStopPrice Can be used if pendingBelowType is STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT or TAKE_PROFIT_LIMIT.
func (s *CreateOTOCOOrder) PendingBelowStopPrice(pendingBelowStopPrice string) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowStopPrice"] = pendingBelowStopPrice
	return s
}
func (s *CreateOTOCOOrder) PendingBelowTrailingDelta(pendingBelowTrailingDelta string) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowTrailingDelta"] = pendingBelowTrailingDelta
	return s
}

// PendingBelowIcebergQty This can only be used if pendingBelowTimeInForce is GTC, or if pendingBelowType is LIMIT_MAKER.
func (s *CreateOTOCOOrder) PendingBelowIcebergQty(pendingBelowIcebergQty string) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowIcebergQty"] = pendingBelowIcebergQty
	return s
}
func (s *CreateOTOCOOrder) PendingBelowTimeInForce(pendingBelowTimeInForce types.TimeInForceEnum) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowTimeInForce"] = pendingBelowTimeInForce
	return s
}
func (s *CreateOTOCOOrder) PendingBelowStrategyId(pendingBelowStrategyId int64) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowStrategyId"] = pendingBelowStrategyId
	return s
}
func (s *CreateOTOCOOrder) PendingBelowStrategyType(pendingBelowStrategyType int) *CreateOTOCOOrder {
	s.c.req.Params["pendingBelowStrategyType"] = pendingBelowStrategyType
	return s
}

func (s *CreateOTOCOOrder) RecvWindow(recvWindow int) *CreateOTOCOOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateOTOCOOrder) Do(ctx context.Context) (*OrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
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

// QueryOrderList Check execution status of an Order list.
type QueryOrderList struct {
	c *Client
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

type OrderListResponse struct {
	ApiResponse
	Result *OrderListResult `json:"result"`
}

func (s *QueryOrderList) OrigClientOrderId(origClientOrderId string) *QueryOrderList {
	s.c.req.Params["origClientOrderId"] = origClientOrderId
	return s
}
func (s *QueryOrderList) OrderListId(orderListId int64) *QueryOrderList {
	s.c.req.Params["orderListId"] = orderListId
	return s
}
func (s *QueryOrderList) RecvWindow(recvWindow int) *QueryOrderList {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}
func (s *QueryOrderList) Do(ctx context.Context) (*OrderListResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *OrderListResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CancelOrderList Cancel an active order list.
type CancelOrderList struct {
	c *Client
}

func (s *CancelOrderList) Symbol(symbol int) *CancelOrderList {
	s.c.req.Params["symbol"] = symbol
	return s
}

// OrderListId Cancel order list by orderListId
func (s *CancelOrderList) OrderListId(orderListId int64) *CancelOrderList {
	s.c.req.Params["orderListId"] = orderListId
	return s
}

// ListClientOrderId Cancel order list by listClientId
func (s *CancelOrderList) ListClientOrderId(listClientOrderId string) *CancelOrderList {
	s.c.req.Params["listClientOrderId"] = listClientOrderId
	return s
}

// NewClientOrderId New ID for the canceled order list. Automatically generated if not sent
func (s *CancelOrderList) NewClientOrderId(newClientOrderId string) *CancelOrderList {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}
func (s *CancelOrderList) RecvWindow(recvWindow int) *CancelOrderList {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}
func (s *CancelOrderList) Do(ctx context.Context) (*OrderListResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *OrderListResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// QueryOpenOrder Query execution status of all open order lists.
type QueryOpenOrder struct {
	c *Client
}

type QueryOpenOrderResponse struct {
	ApiResponse
	Result []*OrderListResult `json:"result"`
}

func (s *QueryOpenOrder) RecvWindow(recvWindow int) *QueryOpenOrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}
func (s *QueryOpenOrder) Do(ctx context.Context) (*QueryOpenOrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *QueryOpenOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CreateSOROrder Places an order using smart order routing (SOR).
type CreateSOROrder struct {
	c *Client
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

type CreateSOROrderResponse struct {
	ApiResponse
	Result []*CreateSOROrderResult `json:"result"`
}

func (s *CreateSOROrder) Symbol(symbol string) *CreateSOROrder {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *CreateSOROrder) Side(side types.OrderSideEnum) *CreateSOROrder {
	s.c.req.Params["side"] = side
	return s
}
func (s *CreateSOROrder) Type(type_ types.OrderTypeEnum) *CreateSOROrder {
	s.c.req.Params["type"] = type_
	return s
}

// TimeInForce Applicable only to LIMIT order type
func (s *CreateSOROrder) TimeInForce(timeInForce types.TimeInForceEnum) *CreateSOROrder {
	s.c.req.Params["timeInForce"] = timeInForce
	return s
}

// Price Applicable only to LIMIT order type
func (s *CreateSOROrder) Price(price string) *CreateSOROrder {
	s.c.req.Params["price"] = price
	return s
}
func (s *CreateSOROrder) Quantity(quantity string) *CreateSOROrder {
	s.c.req.Params["quantity"] = quantity
	return s
}

// NewClientOrderId Arbitrary unique ID among open orders. Automatically generated if not sent
func (s *CreateSOROrder) NewClientOrderId(newClientOrderId string) *CreateSOROrder {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}

// NewOrderRespType Select response format: ACK, RESULT, FULL.
// MARKET and LIMIT orders use FULL by default.
func (s *CreateSOROrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateSOROrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CreateSOROrder) IcebergQty(icebergQty string) *CreateSOROrder {
	s.c.req.Params["icebergQty"] = icebergQty
	return s
}
func (s *CreateSOROrder) StrategyId(strategyId int64) *CreateSOROrder {
	s.c.req.Params["strategyId"] = strategyId
	return s
}
func (s *CreateSOROrder) StrategyType(strategyType int) *CreateSOROrder {
	s.c.req.Params["strategyType"] = strategyType
	return s
}
func (s *CreateSOROrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateSOROrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

func (s *CreateSOROrder) RecvWindow(recvWindow int) *CreateSOROrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateSOROrder) Do(ctx context.Context) (*CreateSOROrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CreateSOROrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// CreateTestSOROrder Test new order creation and signature/recvWindow using smart order routing (SOR). Creates and validates a new order but does not send it into the matching engine.
type CreateTestSOROrder struct {
	c *Client
}

type CreateTestSOROrderResponse struct {
	ApiResponse
	Result *CreateTestOrderResult `json:"result"`
}

func (s *CreateTestSOROrder) Symbol(symbol string) *CreateTestSOROrder {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *CreateTestSOROrder) Side(side types.OrderSideEnum) *CreateTestSOROrder {
	s.c.req.Params["side"] = side
	return s
}
func (s *CreateTestSOROrder) Type(type_ types.OrderTypeEnum) *CreateTestSOROrder {
	s.c.req.Params["type"] = type_
	return s
}
func (s *CreateTestSOROrder) TimeInForce(timeInForce types.TimeInForceEnum) *CreateTestSOROrder {
	s.c.req.Params["timeInForce"] = timeInForce
	return s
}
func (s *CreateTestSOROrder) Price(price string) *CreateTestSOROrder {
	s.c.req.Params["price"] = price
	return s
}
func (s *CreateTestSOROrder) Quantity(quantity string) *CreateTestSOROrder {
	s.c.req.Params["quantity"] = quantity
	return s
}
func (s *CreateTestSOROrder) NewClientOrderId(newClientOrderId string) *CreateTestSOROrder {
	s.c.req.Params["newClientOrderId"] = newClientOrderId
	return s
}
func (s *CreateTestSOROrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateTestSOROrder {
	s.c.req.Params["newOrderRespType"] = newOrderRespType
	return s
}
func (s *CreateTestSOROrder) IcebergQty(icebergQty string) *CreateTestSOROrder {
	s.c.req.Params["icebergQty"] = icebergQty
	return s
}
func (s *CreateTestSOROrder) StrategyId(strategyId int64) *CreateTestSOROrder {
	s.c.req.Params["strategyId"] = strategyId
	return s
}
func (s *CreateTestSOROrder) StrategyType(strategyType int) *CreateTestSOROrder {
	s.c.req.Params["strategyType"] = strategyType
	return s
}
func (s *CreateTestSOROrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateTestSOROrder {
	s.c.req.Params["selfTradePreventionMode"] = selfTradePreventionMode
	return s
}

func (s *CreateTestSOROrder) RecvWindow(recvWindow int) *CreateTestSOROrder {
	s.c.req.Params["recvWindow"] = recvWindow
	return s
}

func (s *CreateTestSOROrder) ComputeCommissionRates(computeCommissionRates bool) *CreateTestSOROrder {
	s.c.req.Params["computeCommissionRates"] = computeCommissionRates
	return s
}

func (s *CreateTestSOROrder) Do(ctx context.Context) (*CreateTestSOROrderResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *CreateTestSOROrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
