package hfutures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
	"github.com/shopspring/decimal"
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
	c                       *Client
	recvWindow              *int64
}

type CreateOrderResponse struct {
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

func (s *CreateOrder) Do(ctx context.Context) (*CreateOrderResponse, error) {
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
	var resp *CreateOrderResponse
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
	CreateOrderResponse
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
