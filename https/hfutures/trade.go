package hfutures

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jekaxv/go-binance/types"
	"github.com/shopspring/decimal"
	"strconv"
	"strings"
)

type OrderReq struct {
	Symbol                  string                      `json:"symbol,omitempty"`
	Side                    types.OrderSideEnum         `json:"side,omitempty"`
	PositionSide            types.PositionSideEnum      `json:"positionSide,omitempty"`
	OrderType               types.OrderTypeEnum         `json:"orderType,omitempty"`
	TimeInForce             types.TimeInForceEnum       `json:"timeInForce,omitempty"`
	Quantity                string                      `json:"quantity,omitempty"`
	ReduceOnly              string                      `json:"reduceOnly,omitempty"`
	Price                   string                      `json:"price,omitempty"`
	NewClientOrderId        string                      `json:"newClientOrderId,omitempty"`
	StopPrice               string                      `json:"stopPrice,omitempty"`
	ClosePosition           string                      `json:"closePosition,omitempty"`
	ActivationPrice         float64                     `json:"activationPrice,omitempty"`
	CallbackRate            float64                     `json:"callbackRate,omitempty"`
	WorkingType             types.WorkingType           `json:"workingType,omitempty"`
	PriceProtect            string                      `json:"priceProtect,omitempty"`
	NewOrderRespType        types.OrderResponseTypeEnum `json:"newOrderRespType,omitempty"`
	PriceMatch              string                      `json:"priceMatch,omitempty"`
	SelfTradePreventionMode types.STPModeEnum           `json:"selfTradePreventionMode,omitempty"`
	GoodTillDate            int64                       `json:"goodTillDate,omitempty"`
}

// CreateOrder Send in a new order.
// https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api
type CreateOrder struct {
	c                       *Client
	symbol                  string
	side                    types.OrderSideEnum
	positionSide            *types.PositionSideEnum
	orderType               types.OrderTypeEnum
	timeInForce             *types.TimeInForceEnum
	quantity                *string
	reduceOnly              *string
	price                   *string
	newClientOrderId        *string
	stopPrice               *string
	closePosition           *string // true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
	activationPrice         *float64
	callbackRate            *float64
	workingType             *types.WorkingType
	priceProtect            *string // "TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
	newOrderRespType        *types.OrderResponseTypeEnum
	priceMatch              *string
	selfTradePreventionMode *types.STPModeEnum
	goodTillDate            *int64
	recvWindow              *int64
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
	s.symbol = symbol
	return s
}

// Side BUY or SELL
func (s *CreateOrder) Side(side types.OrderSideEnum) *CreateOrder {
	s.side = side
	return s
}

func (s *CreateOrder) PositionSide(positionSide types.PositionSideEnum) *CreateOrder {
	s.positionSide = &positionSide
	return s
}

func (s *CreateOrder) Type(orderType types.OrderTypeEnum) *CreateOrder {
	s.orderType = orderType
	return s
}

func (s *CreateOrder) TimeInForce(timeInForce types.TimeInForceEnum) *CreateOrder {
	s.timeInForce = &timeInForce
	return s
}

func (s *CreateOrder) Quantity(quantity string) *CreateOrder {
	s.quantity = &quantity
	return s
}

// ReduceOnly "true" or "false". default "false". Cannot be sent in Hedge Mode; cannot be sent with closePosition=true
func (s *CreateOrder) ReduceOnly(reduceOnly string) *CreateOrder {
	s.reduceOnly = &reduceOnly
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

// StopPrice Used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, and TAKE_PROFIT_LIMIT orders.
func (s *CreateOrder) StopPrice(stopPrice string) *CreateOrder {
	s.stopPrice = &stopPrice
	return s
}

// ClosePosition true, false；Close-All，used with STOP_MARKET or TAKE_PROFIT_MARKET.
func (s *CreateOrder) ClosePosition(closePosition string) *CreateOrder {
	s.closePosition = &closePosition
	return s
}

func (s *CreateOrder) ActivationPrice(activationPrice float64) *CreateOrder {
	s.activationPrice = &activationPrice
	return s
}
func (s *CreateOrder) CallbackRate(callbackRate float64) *CreateOrder {
	s.callbackRate = &callbackRate
	return s
}

// NewOrderRespType set the response JSON.
// ACK, RESULT, or FULL; MARKET and LIMIT order types default to FULL, all other orders default to ACK.
func (s *CreateOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateOrder {
	s.newOrderRespType = &newOrderRespType
	return s
}
func (s *CreateOrder) WorkingType(workingType types.WorkingType) *CreateOrder {
	s.workingType = &workingType
	return s
}

// PriceProtect "TRUE" or "FALSE", default "FALSE". Used with STOP/STOP_MARKET or TAKE_PROFIT/TAKE_PROFIT_MARKET orders.
func (s *CreateOrder) PriceProtect(priceProtect string) *CreateOrder {
	s.priceProtect = &priceProtect
	return s
}
func (s *CreateOrder) PriceMatch(priceMatch string) *CreateOrder {
	s.priceMatch = &priceMatch
	return s
}

// SelfTradePreventionMode The allowed enums is dependent on what is configured on the symbol. The possible supported values are: STP Modes.
func (s *CreateOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateOrder {
	s.selfTradePreventionMode = &selfTradePreventionMode
	return s
}

func (s *CreateOrder) GoodTillDate(goodTillDate int64) *CreateOrder {
	s.goodTillDate = &goodTillDate
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *CreateOrder) RecvWindow(recvWindow int64) *CreateOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *CreateOrder) Do(ctx context.Context) (*OrderResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	if s.positionSide != nil {
		s.c.set("positionSide", *s.positionSide)
	}
	s.c.set("type", s.orderType)
	if s.timeInForce != nil {
		s.c.set("timeInForce", *s.timeInForce)
	}
	if s.quantity != nil {
		s.c.set("quantity", *s.quantity)
	}
	if s.reduceOnly != nil {
		s.c.set("reduceOnly", *s.reduceOnly)
	}
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.newClientOrderId != nil {
		s.c.set("newClientOrderId", *s.newClientOrderId)
	}
	if s.stopPrice != nil {
		s.c.set("stopPrice", *s.stopPrice)
	}
	if s.closePosition != nil {
		s.c.set("closePosition", *s.closePosition)
	}
	if s.stopPrice != nil {
		s.c.set("stopPrice", *s.stopPrice)
	}
	if s.activationPrice != nil {
		s.c.set("activationPrice", *s.activationPrice)
	}
	if s.callbackRate != nil {
		s.c.set("callbackRate", *s.callbackRate)
	}
	if s.workingType != nil {
		s.c.set("workingType", *s.workingType)
	}
	if s.priceProtect != nil {
		s.c.set("priceProtect", *s.priceProtect)
	}
	if s.newOrderRespType != nil {
		s.c.set("newOrderRespType", *s.newOrderRespType)
	}
	if s.priceMatch != nil {
		s.c.set("priceMatch", *s.priceMatch)
	}
	if s.selfTradePreventionMode != nil {
		s.c.set("selfTradePreventionMode", *s.selfTradePreventionMode)
	}
	if s.goodTillDate != nil {
		s.c.set("goodTillDate", *s.goodTillDate)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// PlaceBatchOrder Place Multiple Orders
// https://developers.binance.com/docs/derivatives/usds-margined-futures/trade/rest-api/Place-Multiple-Orders
type PlaceBatchOrder struct {
	c           *Client
	batchOrders []OrderReq
	recvWindow  *int64
}

type PlaceBatchOrderResponse struct {
	OrderResponse
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *PlaceBatchOrder) BatchOrders(batchOrders []OrderReq) *PlaceBatchOrder {
	s.batchOrders = batchOrders
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *PlaceBatchOrder) RecvWindow(recvWindow int64) *PlaceBatchOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *PlaceBatchOrder) Do(ctx context.Context) ([]*PlaceBatchOrderResponse, error) {
	orderJson, err := json.Marshal(s.batchOrders)
	if err != nil {
		return nil, err
	}
	s.c.set("batchOrders", string(orderJson))
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*PlaceBatchOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

type ModifyOrderReq struct {
	OrderId           int64               `json:"orderId"`
	OrigClientOrderId string              `json:"origClientOrderId"`
	Symbol            string              `json:"symbol"`
	Side              types.OrderSideEnum `json:"side"`
	Quantity          string              `json:"quantity"`
	Price             string              `json:"price"`
	PriceMatch        string              `json:"priceMatch"`
}

// ModifyOrder Order modify function, currently only LIMIT order modification is supported, modified orders will be reordered in the match queue
type ModifyOrder struct {
	c                 *Client
	orderId           *int64
	origClientOrderId *string
	symbol            string
	side              types.OrderSideEnum
	quantity          *string
	price             *string
	priceMatch        *string
	recvWindow        *int64
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
	s.orderId = &orderId
	return s
}

func (s *ModifyOrder) OrigClientOrderId(origClientOrderId string) *ModifyOrder {
	s.origClientOrderId = &origClientOrderId
	return s
}

func (s *ModifyOrder) Symbol(symbol string) *ModifyOrder {
	s.symbol = symbol
	return s
}

// Side BUY or SELL
func (s *ModifyOrder) Side(side types.OrderSideEnum) *ModifyOrder {
	s.side = side
	return s
}

func (s *ModifyOrder) Quantity(quantity string) *ModifyOrder {
	s.quantity = &quantity
	return s
}

func (s *ModifyOrder) Price(price string) *ModifyOrder {
	s.price = &price
	return s
}

func (s *ModifyOrder) PriceMatch(priceMatch string) *ModifyOrder {
	s.priceMatch = &priceMatch
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *ModifyOrder) RecvWindow(recvWindow int64) *ModifyOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *ModifyOrder) Do(ctx context.Context) (*ModifyOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("side", s.side)
	if s.orderId != nil {
		s.c.set("orderId", *s.orderId)
	}
	if s.origClientOrderId != nil {
		s.c.set("origClientOrderId", *s.origClientOrderId)
	}
	if s.quantity != nil {
		s.c.set("quantity", *s.quantity)
	}
	if s.price != nil {
		s.c.set("price", *s.price)
	}
	if s.priceMatch != nil {
		s.c.set("priceMatch", *s.priceMatch)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *ModifyOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ModifyMultipleOrder Modify Multiple Orders (TRADE)
type ModifyMultipleOrder struct {
	c           *Client
	batchOrders []ModifyOrderReq
	recvWindow  *int64
}

type ModifyMultipleOrderResponse struct {
	ModifyOrderResponse
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *ModifyMultipleOrder) BatchOrders(batchOrders []ModifyOrderReq) *ModifyMultipleOrder {
	s.batchOrders = batchOrders
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *ModifyMultipleOrder) RecvWindow(recvWindow int64) *ModifyMultipleOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *ModifyMultipleOrder) Do(ctx context.Context) ([]*ModifyMultipleOrderResponse, error) {
	orderJson, err := json.Marshal(s.batchOrders)
	if err != nil {
		return nil, err
	}
	s.c.set("batchOrders", string(orderJson))
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*ModifyMultipleOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OrderAmendment Get order modification history
type OrderAmendment struct {
	c                 *Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
	startTime         *int64
	endTime           *int64
	limit             *int64
	recvWindow        *int64
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
	s.symbol = symbol
	return s
}
func (s *OrderAmendment) OrderId(orderId int64) *OrderAmendment {
	s.orderId = &orderId
	return s
}
func (s *OrderAmendment) OrigClientOrderId(origClientOrderId string) *OrderAmendment {
	s.origClientOrderId = &origClientOrderId
	return s
}
func (s *OrderAmendment) StartTime(startTime int64) *OrderAmendment {
	s.startTime = &startTime
	return s
}
func (s *OrderAmendment) EndTime(endTime int64) *OrderAmendment {
	s.endTime = &endTime
	return s
}

// Limit Default 50; max 100
func (s *OrderAmendment) Limit(limit int64) *OrderAmendment {
	s.limit = &limit
	return s
}
func (s *OrderAmendment) RecvWindow(recvWindow int64) *OrderAmendment {
	s.recvWindow = &recvWindow
	return s
}

func (s *OrderAmendment) Do(ctx context.Context) ([]*OrderAmendmentResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.orderId != nil {
		s.c.set("orderId", *s.orderId)
	}
	if s.origClientOrderId != nil {
		s.c.set("origClientOrderId", *s.origClientOrderId)
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
	var resp []*OrderAmendmentResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOrder Cancel an active order.
type CancelOrder struct {
	c                 *Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (s *CancelOrder) Symbol(symbol string) *CancelOrder {
	s.symbol = symbol
	return s
}
func (s *CancelOrder) OrderId(orderId int64) *CancelOrder {
	s.orderId = &orderId
	return s
}
func (s *CancelOrder) OrigClientOrderId(origClientOrderId string) *CancelOrder {
	s.origClientOrderId = &origClientOrderId
	return s
}
func (s *CancelOrder) RecvWindow(recvWindow int64) *CancelOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *CancelOrder) Do(ctx context.Context) (*OrderResponse, error) {
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
	var resp *OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelMultipleOrder Cancel Multiple Orders
type CancelMultipleOrder struct {
	c                     *Client
	symbol                string
	orderIdList           []int64
	origClientOrderIdList []string
	recvWindow            *int64
}

func (s *CancelMultipleOrder) Symbol(symbol string) *CancelMultipleOrder {
	s.symbol = symbol
	return s
}
func (s *CancelMultipleOrder) OrderIdList(orderIdList []int64) *CancelMultipleOrder {
	s.orderIdList = orderIdList
	return s
}
func (s *CancelMultipleOrder) OrigClientOrderIdList(origClientOrderIdList []string) *CancelMultipleOrder {
	s.origClientOrderIdList = origClientOrderIdList
	return s
}
func (s *CancelMultipleOrder) RecvWindow(recvWindow int64) *CancelMultipleOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *CancelMultipleOrder) Do(ctx context.Context) ([]*OrderResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.orderIdList != nil {
		orderList := "["
		for _, orderId := range s.orderIdList {
			orderList += strconv.FormatInt(orderId, 10) + ","
		}
		orderList = strings.TrimRight(orderList, ",")
		orderList += "]"
		s.c.set("orderIdList", orderList)
	}
	if s.origClientOrderIdList != nil {
		clientOrderIdList := "["
		for _, clientOrderId := range s.origClientOrderIdList {
			clientOrderIdList += fmt.Sprintf(`"%s",`, clientOrderId)
		}
		clientOrderIdList = strings.TrimRight(clientOrderIdList, ",")
		clientOrderIdList += "]"
		s.c.set("origClientOrderIdList", clientOrderIdList)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CancelOpenOrder Cancel All Open Orders
type CancelOpenOrder struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

type CancelOpenOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *CancelOpenOrder) Symbol(symbol string) *CancelOpenOrder {
	s.symbol = symbol
	return s
}
func (s *CancelOpenOrder) RecvWindow(recvWindow int64) *CancelOpenOrder {
	s.recvWindow = &recvWindow
	return s
}
func (s *CancelOpenOrder) Do(ctx context.Context) (*CancelOpenOrderResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *CancelOpenOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CountdownCancelAll Cancel all open orders of the specified symbol at the end of the specified countdown. The endpoint should be called repeatedly as heartbeats so that the existing countdown time can be canceled and replaced by a new one.
type CountdownCancelAll struct {
	c             *Client
	symbol        string
	countdownTime *int64
	recvWindow    *int64
}

type CountdownCancelAllResponse struct {
	Symbol        string `json:"symbol"`
	CountdownTime string `json:"countdownTime"`
}

func (s *CountdownCancelAll) Symbol(symbol string) *CountdownCancelAll {
	s.symbol = symbol
	return s
}
func (s *CountdownCancelAll) CountdownTime(countdownTime int64) *CountdownCancelAll {
	s.countdownTime = &countdownTime
	return s
}
func (s *CountdownCancelAll) RecvWindow(recvWindow int64) *CountdownCancelAll {
	s.recvWindow = &recvWindow
	return s
}
func (s *CountdownCancelAll) Do(ctx context.Context) (*CountdownCancelAllResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.countdownTime != nil {
		s.c.set("countdownTime", *s.countdownTime)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *CountdownCancelAllResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryOrder Check an order's status.
type QueryOrder struct {
	c                 *Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (s *QueryOrder) Symbol(symbol string) *QueryOrder {
	s.symbol = symbol
	return s
}
func (s *QueryOrder) OrderId(orderId int64) *QueryOrder {
	s.orderId = &orderId
	return s
}
func (s *QueryOrder) OrigClientOrderId(origClientOrderId string) *QueryOrder {
	s.origClientOrderId = &origClientOrderId
	return s
}
func (s *QueryOrder) RecvWindow(recvWindow int64) *QueryOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryOrder) Do(ctx context.Context) (*OrderResponse, error) {
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
	var resp *OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryAllOrder Get all account orders; active, canceled, or filled.
type QueryAllOrder struct {
	c          *Client
	symbol     string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	limit      *int64
	recvWindow *int64
}

func (s *QueryAllOrder) Symbol(symbol string) *QueryAllOrder {
	s.symbol = symbol
	return s
}
func (s *QueryAllOrder) OrderId(orderId int64) *QueryAllOrder {
	s.orderId = &orderId
	return s
}
func (s *QueryAllOrder) StartTime(startTime int64) *QueryAllOrder {
	s.startTime = &startTime
	return s
}
func (s *QueryAllOrder) EndTime(endTime int64) *QueryAllOrder {
	s.endTime = &endTime
	return s
}
func (s *QueryAllOrder) Limit(limit int64) *QueryAllOrder {
	s.limit = &limit
	return s
}
func (s *QueryAllOrder) RecvWindow(recvWindow int64) *QueryAllOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryAllOrder) Do(ctx context.Context) ([]*OrderResponse, error) {
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
	var resp []*OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AllOpenOrder Get all open orders on a symbol.
type AllOpenOrder struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

func (s *AllOpenOrder) Symbol(symbol string) *AllOpenOrder {
	s.symbol = &symbol
	return s
}
func (s *AllOpenOrder) RecvWindow(recvWindow int64) *AllOpenOrder {
	s.recvWindow = &recvWindow
	return s
}
func (s *AllOpenOrder) Do(ctx context.Context) ([]*OrderResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryOpenOrder Query open order
type QueryOpenOrder struct {
	c                 *Client
	symbol            string
	orderId           *int64
	origClientOrderId *string
	recvWindow        *int64
}

func (s *QueryOpenOrder) Symbol(symbol string) *QueryOpenOrder {
	s.symbol = symbol
	return s
}
func (s *QueryOpenOrder) OrderId(orderId int64) *QueryOpenOrder {
	s.orderId = &orderId
	return s
}
func (s *QueryOpenOrder) OrigClientOrderId(origClientOrderId string) *QueryOpenOrder {
	s.origClientOrderId = &origClientOrderId
	return s
}
func (s *QueryOpenOrder) RecvWindow(recvWindow int64) *QueryOpenOrder {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryOpenOrder) Do(ctx context.Context) (*OrderResponse, error) {
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
	var resp *OrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ForceOrder Query user's Force Orders
type ForceOrder struct {
	c             *Client
	symbol        *string
	autoCloseType *types.AutoCloseType
	startTime     *int64
	endTime       *int64
	limit         *int64
	recvWindow    *int64
}

func (s *ForceOrder) Symbol(symbol string) *ForceOrder {
	s.symbol = &symbol
	return s
}
func (s *ForceOrder) AutoCloseType(autoCloseType types.AutoCloseType) *ForceOrder {
	s.autoCloseType = &autoCloseType
	return s
}
func (s *ForceOrder) StartTime(startTime int64) *ForceOrder {
	s.startTime = &startTime
	return s
}
func (s *ForceOrder) EndTime(endTime int64) *ForceOrder {
	s.endTime = &endTime
	return s
}

// Limit Default 50; max 100
func (s *ForceOrder) Limit(limit int64) *ForceOrder {
	s.limit = &limit
	return s
}
func (s *ForceOrder) RecvWindow(recvWindow int64) *ForceOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *ForceOrder) Do(ctx context.Context) ([]*ModifyOrderResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.autoCloseType != nil {
		s.c.set("autoCloseType", *s.autoCloseType)
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
	var resp []*ModifyOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// UserTrades Get trades for a specific account and symbol.
type UserTrades struct {
	c          *Client
	symbol     string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
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
	s.symbol = symbol
	return s
}
func (s *UserTrades) OrderId(orderId int64) *UserTrades {
	s.orderId = &orderId
	return s
}
func (s *UserTrades) StartTime(startTime int64) *UserTrades {
	s.startTime = &startTime
	return s
}
func (s *UserTrades) EndTime(endTime int64) *UserTrades {
	s.endTime = &endTime
	return s
}
func (s *UserTrades) FromId(fromId int64) *UserTrades {
	s.fromId = &fromId
	return s
}

// Limit Default 50; max 100
func (s *UserTrades) Limit(limit int64) *UserTrades {
	s.limit = &limit
	return s
}
func (s *UserTrades) RecvWindow(recvWindow int64) *UserTrades {
	s.recvWindow = &recvWindow
	return s
}

func (s *UserTrades) Do(ctx context.Context) ([]*UserTradesResponse, error) {
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
	if s.fromId != nil {
		s.c.set("fromId", *s.fromId)
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
	var resp []*UserTradesResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangeMarginType Change symbol level margin type
type ChangeMarginType struct {
	c          *Client
	symbol     string
	marginType types.MarginType
	recvWindow *int64
}
type ChangeMarginTypeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func (s *ChangeMarginType) Symbol(symbol string) *ChangeMarginType {
	s.symbol = symbol
	return s
}
func (s *ChangeMarginType) MarginType(marginType types.MarginType) *ChangeMarginType {
	s.marginType = marginType
	return s
}
func (s *ChangeMarginType) RecvWindow(recvWindow int64) *ChangeMarginType {
	s.recvWindow = &recvWindow
	return s
}
func (s *ChangeMarginType) Do(ctx context.Context) (*ChangeMarginTypeResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("marginType", s.marginType)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *ChangeMarginTypeResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangePositionSide Change user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol
type ChangePositionSide struct {
	c                *Client
	dualSidePosition string
	recvWindow       *int64
}

// DualSidePosition "true": Hedge Mode; "false": One-way Mode
func (s *ChangePositionSide) DualSidePosition(dualSidePosition string) *ChangePositionSide {
	s.dualSidePosition = dualSidePosition
	return s
}
func (s *ChangePositionSide) RecvWindow(recvWindow int64) *ChangePositionSide {
	s.recvWindow = &recvWindow
	return s
}
func (s *ChangePositionSide) Do(ctx context.Context) (*ChangeMarginTypeResponse, error) {
	s.c.set("dualSidePosition", s.dualSidePosition)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *ChangeMarginTypeResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangeLeverage Change user's initial leverage of specific symbol market.
type ChangeLeverage struct {
	c          *Client
	symbol     string
	leverage   int
	recvWindow *int64
}
type ChangeLeverageResponse struct {
	Leverage         int             `json:"leverage"`
	MaxNotionalValue decimal.Decimal `json:"maxNotionalValue"`
	Symbol           string          `json:"symbol"`
}

func (s *ChangeLeverage) Symbol(symbol string) *ChangeLeverage {
	s.symbol = symbol
	return s
}

// Leverage target initial leverage: int from 1 to 125
func (s *ChangeLeverage) Leverage(leverage int) *ChangeLeverage {
	s.leverage = leverage
	return s
}
func (s *ChangeLeverage) RecvWindow(recvWindow int64) *ChangeLeverage {
	s.recvWindow = &recvWindow
	return s
}
func (s *ChangeLeverage) Do(ctx context.Context) (*ChangeLeverageResponse, error) {
	s.c.set("symbol", s.symbol)
	s.c.set("leverage", s.leverage)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *ChangeLeverageResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangeMultiAssetsMargin Change user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
type ChangeMultiAssetsMargin struct {
	c                 *Client
	multiAssetsMargin string
	recvWindow        *int64
}

// MultiAssetsMargin "true": Multi-Assets Mode; "false": Single-Asset Mode
func (s *ChangeMultiAssetsMargin) MultiAssetsMargin(multiAssetsMargin string) *ChangeMultiAssetsMargin {
	s.multiAssetsMargin = multiAssetsMargin
	return s
}
func (s *ChangeMultiAssetsMargin) RecvWindow(recvWindow int64) *ChangeMultiAssetsMargin {
	s.recvWindow = &recvWindow
	return s
}
func (s *ChangeMultiAssetsMargin) Do(ctx context.Context) (*ChangeMarginTypeResponse, error) {
	s.c.set("multiAssetsMargin", s.multiAssetsMargin)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *ChangeMarginTypeResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangePositionMargin Modify Isolated Position Margin
type ChangePositionMargin struct {
	c            *Client
	symbol       string
	positionSide *types.PositionSideEnum
	amount       string
	type_        int
	recvWindow   *int64
}
type ChangePositionMarginResponse struct {
	Amount float64 `json:"amount"`
	Code   int     `json:"code"`
	Msg    string  `json:"msg"`
	Type   int     `json:"type"`
}

func (s *ChangePositionMargin) Symbol(symbol string) *ChangePositionMargin {
	s.symbol = symbol
	return s
}
func (s *ChangePositionMargin) PositionSide(positionSide types.PositionSideEnum) *ChangePositionMargin {
	s.positionSide = &positionSide
	return s
}
func (s *ChangePositionMargin) Amount(amount string) *ChangePositionMargin {
	s.amount = amount
	return s
}

// Type 1: Add position margin，2: Reduce position margin
func (s *ChangePositionMargin) Type(type_ int) *ChangePositionMargin {
	s.type_ = type_
	return s
}
func (s *ChangePositionMargin) RecvWindow(recvWindow int64) *ChangePositionMargin {
	s.recvWindow = &recvWindow
	return s
}
func (s *ChangePositionMargin) Do(ctx context.Context) (*ChangePositionMarginResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.positionSide != nil {
		s.c.set("positionSide", *s.positionSide)
	}
	s.c.set("amount", s.amount)
	s.c.set("type", s.type_)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *ChangePositionMarginResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}
