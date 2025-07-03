package futures

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type OrderReq struct {
	Symbol                  string                     `json:"symbol,omitempty"`
	Side                    core.OrderSideEnum         `json:"side,omitempty"`
	PositionSide            core.PositionSideEnum      `json:"positionSide,omitempty"`
	OrderType               core.OrderTypeEnum         `json:"orderType,omitempty"`
	TimeInForce             core.TimeInForceEnum       `json:"timeInForce,omitempty"`
	Quantity                string                     `json:"quantity,omitempty"`
	ReduceOnly              string                     `json:"reduceOnly,omitempty"`
	Price                   string                     `json:"price,omitempty"`
	NewClientOrderId        string                     `json:"newClientOrderId,omitempty"`
	StopPrice               string                     `json:"stopPrice,omitempty"`
	ClosePosition           string                     `json:"closePosition,omitempty"`
	ActivationPrice         float64                    `json:"activationPrice,omitempty"`
	CallbackRate            float64                    `json:"callbackRate,omitempty"`
	WorkingType             core.WorkingType           `json:"workingType,omitempty"`
	PriceProtect            string                     `json:"priceProtect,omitempty"`
	NewOrderRespType        core.OrderResponseTypeEnum `json:"newOrderRespType,omitempty"`
	PriceMatch              string                     `json:"priceMatch,omitempty"`
	SelfTradePreventionMode core.STPModeEnum           `json:"selfTradePreventionMode,omitempty"`
	GoodTillDate            int64                      `json:"goodTillDate,omitempty"`
}

// CreateOrder Send in a new order.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api
type CreateOrder struct {
	c *Client
	r *core.Request
}

type OrderResponse struct {
	ClientOrderId           string          `json:"clientOrderId"`
	CumQty                  decimal.Decimal `json:"cumQty"`
	CumQuote                decimal.Decimal `json:"cumQuote"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	OrderId                 int             `json:"orderId"`
	AvgPrice                decimal.Decimal `json:"avgPrice"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	Price                   decimal.Decimal `json:"price"`
	ReduceOnly              bool            `json:"reduceOnly"`
	Side                    string          `json:"side"`
	PositionSide            string          `json:"positionSide"`
	Status                  string          `json:"status"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	ClosePosition           bool            `json:"closePosition"`
	Symbol                  string          `json:"symbol"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	OrigType                string          `json:"origType"`
	ActivatePrice           decimal.Decimal `json:"activatePrice"`
	PriceRate               decimal.Decimal `json:"priceRate"`
	UpdateTime              int64           `json:"updateTime"`
	WorkingType             string          `json:"workingType"`
	PriceProtect            bool            `json:"priceProtect"`
	PriceMatch              string          `json:"priceMatch"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	GoodTillDate            int64           `json:"goodTillDate"`
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

func (s *CreateOrder) PositionSide(positionSide core.PositionSideEnum) *CreateOrder {
	s.r.Set("positionSide", positionSide)
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

// ReduceOnly "true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
func (s *CreateOrder) ReduceOnly(reduceOnly string) *CreateOrder {
	s.r.Set("reduceOnly", reduceOnly)
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

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateOrder) StopPrice(stopPrice string) *CreateOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}

// ClosePosition true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
func (s *CreateOrder) ClosePosition(closePosition string) *CreateOrder {
	s.r.Set("closePosition", closePosition)
	return s
}

func (s *CreateOrder) ActivationPrice(activationPrice float64) *CreateOrder {
	s.r.Set("activationPrice", activationPrice)
	return s
}
func (s *CreateOrder) CallbackRate(callbackRate float64) *CreateOrder {
	s.r.Set("callbackRate", callbackRate)
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *CreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *CreateOrder) WorkingType(workingType core.WorkingType) *CreateOrder {
	s.r.Set("workingType", workingType)
	return s
}

// PriceProtect "TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
func (s *CreateOrder) PriceProtect(priceProtect string) *CreateOrder {
	s.r.Set("priceProtect", priceProtect)
	return s
}
func (s *CreateOrder) PriceMatch(priceMatch string) *CreateOrder {
	s.r.Set("priceMatch", priceMatch)
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *CreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *CreateOrder) GoodTillDate(goodTillDate int64) *CreateOrder {
	s.r.Set("goodTillDate", goodTillDate)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateOrder) RecvWindow(recvWindow int64) *CreateOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CreateOrder) Do(ctx context.Context) (*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// PlaceBatchOrder Place Multiple Orders
// https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Place-Multiple-Orders
type PlaceBatchOrder struct {
	c *Client
	r *core.Request
}

type PlaceBatchOrderResponse struct {
	OrderResponse
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *PlaceBatchOrder) BatchOrders(batchOrders []OrderReq) *PlaceBatchOrder {
	orderJson, err := json.Marshal(batchOrders)
	if err != nil {
		return s
	}
	s.r.Set("batchOrders", string(orderJson))
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *PlaceBatchOrder) RecvWindow(recvWindow int64) *PlaceBatchOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *PlaceBatchOrder) Do(ctx context.Context) ([]*PlaceBatchOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*PlaceBatchOrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

type ModifyOrderReq struct {
	OrderId           int64              `json:"orderId"`
	OrigClientOrderId string             `json:"origClientOrderId"`
	Symbol            string             `json:"symbol"`
	Side              core.OrderSideEnum `json:"side"`
	Quantity          string             `json:"quantity"`
	Price             string             `json:"price"`
	PriceMatch        string             `json:"priceMatch"`
}

// ModifyOrder Order modify function, currently only LIMIT order modification is supported, modified orders will be reordered in the match queue
type ModifyOrder struct {
	c *Client
	r *core.Request
}

type ModifyOrderResponse struct {
	OrderId                 int64           `json:"orderId"`
	Symbol                  string          `json:"symbol"`
	Pair                    string          `json:"pair"`
	Status                  string          `json:"status"`
	ClientOrderId           string          `json:"clientOrderId"`
	Price                   decimal.Decimal `json:"price"`
	AvgPrice                decimal.Decimal `json:"avgPrice"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	CumQty                  decimal.Decimal `json:"cumQty"`
	CumBase                 decimal.Decimal `json:"cumBase"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	ReduceOnly              bool            `json:"reduceOnly"`
	ClosePosition           bool            `json:"closePosition"`
	Side                    string          `json:"side"`
	PositionSide            string          `json:"positionSide"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	WorkingType             string          `json:"workingType"`
	PriceProtect            bool            `json:"priceProtect"`
	OrigType                string          `json:"origType"`
	PriceMatch              string          `json:"priceMatch"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	GoodTillDate            int             `json:"goodTillDate"`
	UpdateTime              int64           `json:"updateTime"`
}

func (s *ModifyOrder) OrderId(orderId int64) *ModifyOrder {
	s.r.Set("orderId", orderId)
	return s
}

func (s *ModifyOrder) OrigClientOrderId(origClientOrderId string) *ModifyOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

func (s *ModifyOrder) Symbol(symbol string) *ModifyOrder {
	s.r.Set("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *ModifyOrder) Side(side core.OrderSideEnum) *ModifyOrder {
	s.r.Set("side", side)
	return s
}

func (s *ModifyOrder) Quantity(quantity string) *ModifyOrder {
	s.r.Set("quantity", quantity)
	return s
}

func (s *ModifyOrder) Price(price string) *ModifyOrder {
	s.r.Set("price", price)
	return s
}

func (s *ModifyOrder) PriceMatch(priceMatch string) *ModifyOrder {
	s.r.Set("priceMatch", priceMatch)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *ModifyOrder) RecvWindow(recvWindow int64) *ModifyOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *ModifyOrder) Do(ctx context.Context) (*ModifyOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ModifyOrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ModifyMultipleOrder Modify Multiple Orders (TRADE)
type ModifyMultipleOrder struct {
	c *Client
	r *core.Request
}

type ModifyMultipleOrderResponse struct {
	ModifyOrderResponse
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *ModifyMultipleOrder) BatchOrders(batchOrders []ModifyOrderReq) *ModifyMultipleOrder {
	orderJson, err := json.Marshal(batchOrders)
	if err != nil {
		return s
	}
	s.r.Set("batchOrders", string(orderJson))
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *ModifyMultipleOrder) RecvWindow(recvWindow int64) *ModifyMultipleOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *ModifyMultipleOrder) Do(ctx context.Context) ([]*ModifyMultipleOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*ModifyMultipleOrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OrderAmendment Get order modification history
type OrderAmendment struct {
	c *Client
	r *core.Request
}

type AmendmentQuantity struct {
	Before decimal.Decimal `json:"before"`
	After  decimal.Decimal `json:"after"`
}

type Amendment struct {
	Price   *AmendmentQuantity `json:"price"`
	OrigQty *AmendmentQuantity `json:"origQty"`
	Count   int                `json:"count"`
}

type OrderAmendmentResponse struct {
	AmendmentId   int        `json:"amendmentId"`
	Symbol        string     `json:"symbol"`
	Pair          string     `json:"pair"`
	OrderId       int64      `json:"orderId"`
	ClientOrderId string     `json:"clientOrderId"`
	Time          int64      `json:"time"`
	Amendment     *Amendment `json:"amendment"`
}

func (s *OrderAmendment) Symbol(symbol string) *OrderAmendment {
	s.r.Set("symbol", symbol)
	return s
}
func (s *OrderAmendment) OrderId(orderId int64) *OrderAmendment {
	s.r.Set("orderId", orderId)
	return s
}
func (s *OrderAmendment) OrigClientOrderId(origClientOrderId string) *OrderAmendment {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}
func (s *OrderAmendment) StartTime(startTime int64) *OrderAmendment {
	s.r.Set("startTime", startTime)
	return s
}
func (s *OrderAmendment) EndTime(endTime int64) *OrderAmendment {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 50; max 100
func (s *OrderAmendment) Limit(limit int64) *OrderAmendment {
	s.r.Set("limit", limit)
	return s
}
func (s *OrderAmendment) RecvWindow(recvWindow int64) *OrderAmendment {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *OrderAmendment) Do(ctx context.Context) ([]*OrderAmendmentResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OrderAmendmentResponse, 0)
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
func (s *CancelOrder) OrderId(orderId int64) *CancelOrder {
	s.r.Set("orderId", orderId)
	return s
}
func (s *CancelOrder) OrigClientOrderId(origClientOrderId string) *CancelOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}
func (s *CancelOrder) RecvWindow(recvWindow int64) *CancelOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CancelOrder) Do(ctx context.Context) (*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// CancelMultipleOrder Cancel Multiple Orders
type CancelMultipleOrder struct {
	c *Client
	r *core.Request
}

func (s *CancelMultipleOrder) Symbol(symbol string) *CancelMultipleOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *CancelMultipleOrder) OrderIdList(orderIdList []int64) *CancelMultipleOrder {
	orderList := "["
	for _, orderId := range orderIdList {
		orderList += strconv.FormatInt(orderId, 10) + ","
	}
	orderList = strings.TrimRight(orderList, ",")
	orderList += "]"
	s.r.Set("orderIdList", orderIdList)
	return s
}
func (s *CancelMultipleOrder) OrigClientOrderIdList(origClientOrderIdList []string) *CancelMultipleOrder {
	clientOrderIdList := "["
	for _, clientOrderId := range origClientOrderIdList {
		clientOrderIdList += fmt.Sprintf(`"%s",`, clientOrderId)
	}
	clientOrderIdList = strings.TrimRight(clientOrderIdList, ",")
	clientOrderIdList += "]"
	s.r.Set("origClientOrderIdList", origClientOrderIdList)
	return s
}
func (s *CancelMultipleOrder) RecvWindow(recvWindow int64) *CancelMultipleOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CancelMultipleOrder) Do(ctx context.Context) ([]*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOpenOrder Cancel All Open Orders
type CancelOpenOrder struct {
	c *Client
	r *core.Request
}

type CancelOpenOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *CancelOpenOrder) Symbol(symbol string) *CancelOpenOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *CancelOpenOrder) RecvWindow(recvWindow int64) *CancelOpenOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *CancelOpenOrder) Do(ctx context.Context) (*CancelOpenOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(CancelOpenOrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// CountdownCancelAll Cancel all open orders of the specified symbol at the end of the specified countdown. The endpoint should be called repeatedly as heartbeats so that the existing countdown time can be canceled and replaced by a new one.
type CountdownCancelAll struct {
	c *Client
	r *core.Request
}

type CountdownCancelAllResponse struct {
	Symbol        string `json:"symbol"`
	CountdownTime string `json:"countdownTime"`
}

func (s *CountdownCancelAll) Symbol(symbol string) *CountdownCancelAll {
	s.r.Set("symbol", symbol)
	return s
}
func (s *CountdownCancelAll) CountdownTime(countdownTime int64) *CountdownCancelAll {
	s.r.Set("countdownTime", countdownTime)
	return s
}
func (s *CountdownCancelAll) RecvWindow(recvWindow int64) *CountdownCancelAll {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *CountdownCancelAll) Do(ctx context.Context) (*CountdownCancelAllResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(CountdownCancelAllResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// QueryOrder Check an order's status.
type QueryOrder struct {
	c *Client
	r *core.Request
}

func (s *QueryOrder) Symbol(symbol string) *QueryOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *QueryOrder) OrderId(orderId int64) *QueryOrder {
	s.r.Set("orderId", orderId)
	return s
}
func (s *QueryOrder) OrigClientOrderId(origClientOrderId string) *QueryOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}
func (s *QueryOrder) RecvWindow(recvWindow int64) *QueryOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *QueryOrder) Do(ctx context.Context) (*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// QueryAllOrder Get all account orders; active, canceled, or filled.
type QueryAllOrder struct {
	c *Client
	r *core.Request
}

func (s *QueryAllOrder) Symbol(symbol string) *QueryAllOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *QueryAllOrder) OrderId(orderId int64) *QueryAllOrder {
	s.r.Set("orderId", orderId)
	return s
}
func (s *QueryAllOrder) StartTime(startTime int64) *QueryAllOrder {
	s.r.Set("startTime", startTime)
	return s
}
func (s *QueryAllOrder) EndTime(endTime int64) *QueryAllOrder {
	s.r.Set("endTime", endTime)
	return s
}
func (s *QueryAllOrder) Limit(limit int64) *QueryAllOrder {
	s.r.Set("limit", limit)
	return s
}
func (s *QueryAllOrder) RecvWindow(recvWindow int64) *QueryAllOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *QueryAllOrder) Do(ctx context.Context) ([]*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AllOpenOrder Get all open orders on a symbol.
type AllOpenOrder struct {
	c *Client
	r *core.Request
}

func (s *AllOpenOrder) Symbol(symbol string) *AllOpenOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *AllOpenOrder) RecvWindow(recvWindow int64) *AllOpenOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AllOpenOrder) Do(ctx context.Context) ([]*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryOpenOrder Query open order
type QueryOpenOrder struct {
	c *Client
	r *core.Request
}

func (s *QueryOpenOrder) Symbol(symbol string) *QueryOpenOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *QueryOpenOrder) OrderId(orderId int64) *QueryOpenOrder {
	s.r.Set("orderId", orderId)
	return s
}
func (s *QueryOpenOrder) OrigClientOrderId(origClientOrderId string) *QueryOpenOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}
func (s *QueryOpenOrder) RecvWindow(recvWindow int64) *QueryOpenOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *QueryOpenOrder) Do(ctx context.Context) (*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ForceOrder Query user's Force Orders
type ForceOrder struct {
	c *Client
	r *core.Request
}

func (s *ForceOrder) Symbol(symbol string) *ForceOrder {
	s.r.Set("symbol", symbol)
	return s
}
func (s *ForceOrder) AutoCloseType(autoCloseType core.AutoCloseType) *ForceOrder {
	s.r.Set("autoCloseType", autoCloseType)
	return s
}
func (s *ForceOrder) StartTime(startTime int64) *ForceOrder {
	s.r.Set("startTime", startTime)
	return s
}
func (s *ForceOrder) EndTime(endTime int64) *ForceOrder {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 50; max 100
func (s *ForceOrder) Limit(limit int64) *ForceOrder {
	s.r.Set("limit", limit)
	return s
}
func (s *ForceOrder) RecvWindow(recvWindow int64) *ForceOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *ForceOrder) Do(ctx context.Context) ([]*ModifyOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*ModifyOrderResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// UserTrades Get trades for a specific account and symbol.
type UserTrades struct {
	c *Client
	r *core.Request
}

type UserTradesResponse struct {
	Buyer           bool            `json:"buyer"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	Id              int             `json:"id"`
	Maker           bool            `json:"maker"`
	OrderId         int             `json:"orderId"`
	Price           decimal.Decimal `json:"price"`
	Qty             decimal.Decimal `json:"qty"`
	QuoteQty        decimal.Decimal `json:"quoteQty"`
	RealizedPnl     decimal.Decimal `json:"realizedPnl"`
	Side            string          `json:"side"`
	PositionSide    string          `json:"positionSide"`
	Symbol          string          `json:"symbol"`
	Time            int64           `json:"time"`
}

func (s *UserTrades) Symbol(symbol string) *UserTrades {
	s.r.Set("symbol", symbol)
	return s
}
func (s *UserTrades) OrderId(orderId int64) *UserTrades {
	s.r.Set("orderId", orderId)
	return s
}
func (s *UserTrades) StartTime(startTime int64) *UserTrades {
	s.r.Set("startTime", startTime)
	return s
}
func (s *UserTrades) EndTime(endTime int64) *UserTrades {
	s.r.Set("endTime", endTime)
	return s
}
func (s *UserTrades) FromId(fromId int64) *UserTrades {
	s.r.Set("fromId", fromId)
	return s
}

// Limit Default 50; max 100
func (s *UserTrades) Limit(limit int64) *UserTrades {
	s.r.Set("limit", limit)
	return s
}
func (s *UserTrades) RecvWindow(recvWindow int64) *UserTrades {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *UserTrades) Do(ctx context.Context) ([]*UserTradesResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*UserTradesResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangeMarginType Change symbol level margin type
type ChangeMarginType struct {
	c *Client
	r *core.Request
}
type ChangeMarginTypeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *ChangeMarginType) Symbol(symbol string) *ChangeMarginType {
	s.r.Set("symbol", symbol)
	return s
}
func (s *ChangeMarginType) MarginType(marginType core.MarginType) *ChangeMarginType {
	s.r.Set("marginType", marginType)
	return s
}
func (s *ChangeMarginType) RecvWindow(recvWindow int64) *ChangeMarginType {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *ChangeMarginType) Do(ctx context.Context) (*ChangeMarginTypeResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ChangeMarginTypeResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ChangePositionSide Change user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol
type ChangePositionSide struct {
	c *Client
	r *core.Request
}

// DualSidePosition "true": Hedge Mode; "false": One-way Mode
func (s *ChangePositionSide) DualSidePosition(dualSidePosition string) *ChangePositionSide {
	s.r.Set("dualSidePosition", dualSidePosition)
	return s
}
func (s *ChangePositionSide) RecvWindow(recvWindow int64) *ChangePositionSide {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *ChangePositionSide) Do(ctx context.Context) (*ChangeMarginTypeResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ChangeMarginTypeResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ChangeLeverage Change user's initial leverage of specific symbol market.
type ChangeLeverage struct {
	c *Client
	r *core.Request
}
type ChangeLeverageResponse struct {
	Leverage         int             `json:"leverage"`
	MaxNotionalValue decimal.Decimal `json:"maxNotionalValue"`
	Symbol           string          `json:"symbol"`
}

func (s *ChangeLeverage) Symbol(symbol string) *ChangeLeverage {
	s.r.Set("symbol", symbol)
	return s
}

// Leverage target initial leverage: int from 1 to 125
func (s *ChangeLeverage) Leverage(leverage int) *ChangeLeverage {
	s.r.Set("leverage", leverage)
	return s
}
func (s *ChangeLeverage) RecvWindow(recvWindow int64) *ChangeLeverage {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *ChangeLeverage) Do(ctx context.Context) (*ChangeLeverageResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ChangeLeverageResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ChangeMultiAssetsMargin Change user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
type ChangeMultiAssetsMargin struct {
	c *Client
	r *core.Request
}

// MultiAssetsMargin "true": Multi-Assets Mode; "false": Single-Asset Mode
func (s *ChangeMultiAssetsMargin) MultiAssetsMargin(multiAssetsMargin string) *ChangeMultiAssetsMargin {
	s.r.Set("multiAssetsMargin", multiAssetsMargin)
	return s
}
func (s *ChangeMultiAssetsMargin) RecvWindow(recvWindow int64) *ChangeMultiAssetsMargin {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *ChangeMultiAssetsMargin) Do(ctx context.Context) (*ChangeMarginTypeResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ChangeMarginTypeResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ChangePositionMargin Modify Isolated Position Margin
type ChangePositionMargin struct {
	c *Client
	r *core.Request
}
type ChangePositionMarginResponse struct {
	Amount float64 `json:"amount"`
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Type   int     `json:"type"`
}

func (s *ChangePositionMargin) Symbol(symbol string) *ChangePositionMargin {
	s.r.Set("symbol", symbol)
	return s
}
func (s *ChangePositionMargin) PositionSide(positionSide core.PositionSideEnum) *ChangePositionMargin {
	s.r.Set("positionSide", positionSide)
	return s
}
func (s *ChangePositionMargin) Amount(amount string) *ChangePositionMargin {
	s.r.Set("amount", amount)
	return s
}

// Type 1: Add position margin，2: Reduce position margin
func (s *ChangePositionMargin) Type(type_ int) *ChangePositionMargin {
	s.r.Set("type", type_)
	return s
}
func (s *ChangePositionMargin) RecvWindow(recvWindow int64) *ChangePositionMargin {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *ChangePositionMargin) Do(ctx context.Context) (*ChangePositionMarginResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ChangePositionMarginResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// PositionRisk Get current position information.
type PositionRisk struct {
	c *Client
	r *core.Request
}

type PositionRiskResponse struct {
	Symbol                 string          `json:"symbol"`
	PositionSide           string          `json:"positionSide"`
	PositionAmt            decimal.Decimal `json:"positionAmt"`
	EntryPrice             decimal.Decimal `json:"entryPrice"`
	BreakEvenPrice         decimal.Decimal `json:"breakEvenPrice"`
	MarkPrice              decimal.Decimal `json:"markPrice"`
	UnRealizedProfit       decimal.Decimal `json:"unRealizedProfit"`
	LiquidationPrice       decimal.Decimal `json:"liquidationPrice"`
	IsolatedMargin         decimal.Decimal `json:"isolatedMargin"`
	Notional               decimal.Decimal `json:"notional"`
	MarginAsset            string          `json:"marginAsset"`
	IsolatedWallet         decimal.Decimal `json:"isolatedWallet"`
	InitialMargin          decimal.Decimal `json:"initialMargin"`
	MaintMargin            decimal.Decimal `json:"maintMargin"`
	PositionInitialMargin  decimal.Decimal `json:"positionInitialMargin"`
	OpenOrderInitialMargin decimal.Decimal `json:"openOrderInitialMargin"`
	Adl                    int             `json:"adl"`
	BidNotional            string          `json:"bidNotional"`
	AskNotional            string          `json:"askNotional"`
	UpdateTime             int64           `json:"updateTime"`
}

func (s *PositionRisk) Symbol(symbol string) *PositionRisk {
	s.r.Set("symbol", symbol)
	return s
}
func (s *PositionRisk) RecvWindow(recvWindow int64) *PositionRisk {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *PositionRisk) Do(ctx context.Context) ([]*PositionRiskResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*PositionRiskResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AdlQuantile Position ADL Quantile Estimation
type AdlQuantile struct {
	c *Client
	r *core.Request
}

type AdlQuantileResponse struct {
	Symbol      string `json:"symbol"`
	AdlQuantile struct {
		LONG  int `json:"LONG"`
		SHORT int `json:"SHORT"`
		HEDGE int `json:"HEDGE"`
		BOTH  int `json:"BOTH"`
	} `json:"adlQuantile"`
}

func (s *AdlQuantile) Symbol(symbol string) *AdlQuantile {
	s.r.Set("symbol", symbol)
	return s
}
func (s *AdlQuantile) RecvWindow(recvWindow int64) *AdlQuantile {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AdlQuantile) Do(ctx context.Context) ([]*AdlQuantileResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*AdlQuantileResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// PositionMarginHistory Get Position Margin Change History
type PositionMarginHistory struct {
	c *Client
	r *core.Request
}

type PositionMarginHistoryResponse struct {
	Symbol       string          `json:"symbol"`
	Type         int             `json:"type"`
	DeltaType    string          `json:"deltaType"`
	Amount       decimal.Decimal `json:"amount"`
	Asset        string          `json:"asset"`
	Time         int64           `json:"time"`
	PositionSide string          `json:"positionSide"`
}

func (s *PositionMarginHistory) Symbol(symbol string) *PositionMarginHistory {
	s.r.Set("symbol", symbol)
	return s
}
func (s *PositionMarginHistory) Type(type_ int) *PositionMarginHistory {
	s.r.Set("type", type_)
	return s
}
func (s *PositionMarginHistory) StartTime(startTime int64) *PositionMarginHistory {
	s.r.Set("startTime", startTime)
	return s
}
func (s *PositionMarginHistory) EndTime(endTime int64) *PositionMarginHistory {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500
func (s *PositionMarginHistory) Limit(limit int64) *PositionMarginHistory {
	s.r.Set("limit", limit)
	return s
}
func (s *PositionMarginHistory) RecvWindow(recvWindow int64) *PositionMarginHistory {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *PositionMarginHistory) Do(ctx context.Context) ([]*PositionMarginHistoryResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*PositionMarginHistoryResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CreateTestOrder Testing order request, this order will not be submitted to matching engine
type CreateTestOrder struct {
	c *Client
	r *core.Request
}

func (s *CreateTestOrder) Symbol(symbol string) *CreateTestOrder {
	s.r.Set("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *CreateTestOrder) Side(side core.OrderSideEnum) *CreateTestOrder {
	s.r.Set("side", side)
	return s
}

func (s *CreateTestOrder) PositionSide(positionSide core.PositionSideEnum) *CreateTestOrder {
	s.r.Set("positionSide", positionSide)
	return s
}

func (s *CreateTestOrder) Type(orderType core.OrderTypeEnum) *CreateTestOrder {
	s.r.Set("type", orderType)
	return s
}

func (s *CreateTestOrder) TimeInForce(timeInForce core.TimeInForceEnum) *CreateTestOrder {
	s.r.Set("timeInForce", timeInForce)
	return s
}

func (s *CreateTestOrder) Quantity(quantity string) *CreateTestOrder {
	s.r.Set("quantity", quantity)
	return s
}

// ReduceOnly "true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
func (s *CreateTestOrder) ReduceOnly(reduceOnly string) *CreateTestOrder {
	s.r.Set("reduceOnly", reduceOnly)
	return s
}
func (s *CreateTestOrder) Price(price string) *CreateTestOrder {
	s.r.Set("price", price)
	return s
}

// NewClientOrderId A unique id among open orders. Automatically generated if not sent.
// Orders with the same newClientOrderID can be accepted only when the previous one is filled, otherwise the order will be rejected.
func (s *CreateTestOrder) NewClientOrderId(newClientOrderId string) *CreateTestOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateTestOrder) StopPrice(stopPrice string) *CreateTestOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}

// ClosePosition true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
func (s *CreateTestOrder) ClosePosition(closePosition string) *CreateTestOrder {
	s.r.Set("closePosition", closePosition)
	return s
}

func (s *CreateTestOrder) ActivationPrice(activationPrice float64) *CreateTestOrder {
	s.r.Set("activationPrice", activationPrice)
	return s
}
func (s *CreateTestOrder) CallbackRate(callbackRate float64) *CreateTestOrder {
	s.r.Set("callbackRate", callbackRate)
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *CreateTestOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *CreateTestOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}
func (s *CreateTestOrder) WorkingType(workingType core.WorkingType) *CreateTestOrder {
	s.r.Set("workingType", workingType)
	return s
}

// PriceProtect "TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
func (s *CreateTestOrder) PriceProtect(priceProtect string) *CreateTestOrder {
	s.r.Set("priceProtect", priceProtect)
	return s
}
func (s *CreateTestOrder) PriceMatch(priceMatch string) *CreateTestOrder {
	s.r.Set("priceMatch", priceMatch)
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *CreateTestOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *CreateTestOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}

func (s *CreateTestOrder) GoodTillDate(goodTillDate int64) *CreateTestOrder {
	s.r.Set("goodTillDate", goodTillDate)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateTestOrder) RecvWindow(recvWindow int64) *CreateTestOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *CreateTestOrder) Do(ctx context.Context) (*OrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OrderResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}
