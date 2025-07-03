package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// WsCreateOrder Send in a new order.
type WsCreateOrder struct {
	c *WsClient
	r *core.WsRequest
}

type OrderResult struct {
	OrderId                 int             `json:"orderId"`
	Symbol                  string          `json:"symbol"`
	Status                  string          `json:"status"`
	ClientOrderId           string          `json:"clientOrderId"`
	Price                   decimal.Decimal `json:"price"`
	AvgPrice                decimal.Decimal `json:"avgPrice"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	CumQty                  decimal.Decimal `json:"cumQty"`
	CumQuote                decimal.Decimal `json:"cumQuote"`
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

type WsOrderResponse struct {
	ApiResponse
	Result *OrderResult `json:"result"`
}

func (s *WsCreateOrder) Symbol(symbol string) *WsCreateOrder {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsCreateOrder) Side(side core.OrderSideEnum) *WsCreateOrder {
	s.r.Set("side", side)
	return s
}
func (s *WsCreateOrder) PositionSide(positionSide core.PositionSideEnum) *WsCreateOrder {
	s.r.Set("positionSide", positionSide)
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
func (s *WsCreateOrder) Quantity(quantity float64) *WsCreateOrder {
	s.r.Set("quantity", quantity)
	return s
}
func (s *WsCreateOrder) ReduceOnly(reduceOnly string) *WsCreateOrder {
	s.r.Set("reduceOnly", reduceOnly)
	return s
}

func (s *WsCreateOrder) Price(price float64) *WsCreateOrder {
	s.r.Set("price", price)
	return s
}
func (s *WsCreateOrder) NewClientOrderId(newClientOrderId string) *WsCreateOrder {
	s.r.Set("newClientOrderId", newClientOrderId)
	return s
}
func (s *WsCreateOrder) StopPrice(stopPrice float64) *WsCreateOrder {
	s.r.Set("stopPrice", stopPrice)
	return s
}
func (s *WsCreateOrder) ClosePosition(closePosition string) *WsCreateOrder {
	s.r.Set("closePosition", closePosition)
	return s
}
func (s *WsCreateOrder) ActivationPrice(activationPrice float64) *WsCreateOrder {
	s.r.Set("activationPrice", activationPrice)
	return s
}
func (s *WsCreateOrder) CallbackRate(callbackRate float64) *WsCreateOrder {
	s.r.Set("callbackRate", callbackRate)
	return s
}
func (s *WsCreateOrder) WorkingType(workingType core.WorkingType) *WsCreateOrder {
	s.r.Set("workingType", workingType)
	return s
}
func (s *WsCreateOrder) PriceProtect(priceProtect string) *WsCreateOrder {
	s.r.Set("priceProtect", priceProtect)
	return s
}
func (s *WsCreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateOrder {
	s.r.Set("newOrderRespType", newOrderRespType)
	return s
}

func (s *WsCreateOrder) PriceMatch(priceMatch string) *WsCreateOrder {
	s.r.Set("priceMatch", priceMatch)
	return s
}
func (s *WsCreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateOrder {
	s.r.Set("selfTradePreventionMode", selfTradePreventionMode)
	return s
}
func (s *WsCreateOrder) GoodTillDate(goodTillDate int64) *WsCreateOrder {
	s.r.Set("goodTillDate", goodTillDate)
	return s
}
func (s *WsCreateOrder) RecvWindow(recvWindow int64) *WsCreateOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCreateOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
			var resp *WsOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsModifyOrder Order modify function, currently only LIMIT order modification is supported, modified orders will be reordered in the match queue
type WsModifyOrder struct {
	c *WsClient
	r *core.WsRequest
}

func (s *WsModifyOrder) OrderId(orderId int64) *WsModifyOrder {
	s.r.Set("orderId", orderId)
	return s
}

func (s *WsModifyOrder) OrigClientOrderId(origClientOrderId string) *WsModifyOrder {
	s.r.Set("origClientOrderId", origClientOrderId)
	return s
}

func (s *WsModifyOrder) Symbol(symbol string) *WsModifyOrder {
	s.r.Set("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *WsModifyOrder) Side(side core.OrderSideEnum) *WsModifyOrder {
	s.r.Set("side", side)
	return s
}

func (s *WsModifyOrder) Quantity(quantity float64) *WsModifyOrder {
	s.r.Set("quantity", quantity)
	return s
}

func (s *WsModifyOrder) Price(price float64) *WsModifyOrder {
	s.r.Set("price", price)
	return s
}

func (s *WsModifyOrder) PriceMatch(priceMatch string) *WsModifyOrder {
	s.r.Set("priceMatch", priceMatch)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *WsModifyOrder) RecvWindow(recvWindow int64) *WsModifyOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsModifyOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
			var resp *WsOrderResponse
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
func (s *WsCancelOrder) RecvWindow(recvWindow int64) *WsCancelOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsCancelOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
			var resp *WsOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsQueryOrder Check an order's status.
type WsQueryOrder struct {
	c *WsClient
	r *core.WsRequest
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
func (s *WsQueryOrder) RecvWindow(recvWindow int64) *WsQueryOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *WsQueryOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
			var resp *WsOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsPositionInfo Get current position information(only symbol that has position or open orders will be returned).
type WsPositionInfo struct {
	c *WsClient
	r *core.WsRequest
}

type PositionInfoResult struct {
	Symbol                 string          `json:"symbol"`
	PositionSide           string          `json:"positionSide"`
	PositionAmt            decimal.Decimal `json:"positionAmt"`
	EntryPrice             decimal.Decimal `json:"entryPrice"`
	BreakEvenPrice         decimal.Decimal `json:"breakEvenPrice"`
	MarkPrice              decimal.Decimal `json:"markPrice"`
	UnrealizedProfit       decimal.Decimal `json:"unrealizedProfit"`
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
	UpdateTime             int             `json:"updateTime"`
}

type PositionInfoResponse struct {
	ApiResponse
	Result []*PositionInfoResult `json:"result"`
}

func (s *WsPositionInfo) Symbol(symbol string) *WsPositionInfo {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsPositionInfo) RecvWindow(recvWindow int64) *WsPositionInfo {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *WsPositionInfo) Do(ctx context.Context) (*PositionInfoResponse, error) {
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
			var resp *PositionInfoResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
