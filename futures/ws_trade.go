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
	s.c.setParams("symbol", symbol)
	return s
}

func (s *WsCreateOrder) Side(side core.OrderSideEnum) *WsCreateOrder {
	s.c.setParams("side", side)
	return s
}
func (s *WsCreateOrder) PositionSide(positionSide core.PositionSideEnum) *WsCreateOrder {
	s.c.setParams("positionSide", positionSide)
	return s
}

func (s *WsCreateOrder) Type(orderType core.OrderTypeEnum) *WsCreateOrder {
	s.c.setParams("type", orderType)
	return s
}

func (s *WsCreateOrder) TimeInForce(timeInForce core.TimeInForceEnum) *WsCreateOrder {
	s.c.setParams("timeInForce", timeInForce)
	return s
}
func (s *WsCreateOrder) Quantity(quantity float64) *WsCreateOrder {
	s.c.setParams("quantity", quantity)
	return s
}
func (s *WsCreateOrder) ReduceOnly(reduceOnly string) *WsCreateOrder {
	s.c.setParams("reduceOnly", reduceOnly)
	return s
}

func (s *WsCreateOrder) Price(price float64) *WsCreateOrder {
	s.c.setParams("price", price)
	return s
}
func (s *WsCreateOrder) NewClientOrderId(newClientOrderId string) *WsCreateOrder {
	s.c.setParams("newClientOrderId", newClientOrderId)
	return s
}
func (s *WsCreateOrder) StopPrice(stopPrice float64) *WsCreateOrder {
	s.c.setParams("stopPrice", stopPrice)
	return s
}
func (s *WsCreateOrder) ClosePosition(closePosition string) *WsCreateOrder {
	s.c.setParams("closePosition", closePosition)
	return s
}
func (s *WsCreateOrder) ActivationPrice(activationPrice float64) *WsCreateOrder {
	s.c.setParams("activationPrice", activationPrice)
	return s
}
func (s *WsCreateOrder) CallbackRate(callbackRate float64) *WsCreateOrder {
	s.c.setParams("callbackRate", callbackRate)
	return s
}
func (s *WsCreateOrder) WorkingType(workingType core.WorkingType) *WsCreateOrder {
	s.c.setParams("workingType", workingType)
	return s
}
func (s *WsCreateOrder) PriceProtect(priceProtect string) *WsCreateOrder {
	s.c.setParams("priceProtect", priceProtect)
	return s
}
func (s *WsCreateOrder) NewOrderRespType(newOrderRespType core.OrderResponseTypeEnum) *WsCreateOrder {
	s.c.setParams("newOrderRespType", newOrderRespType)
	return s
}

func (s *WsCreateOrder) PriceMatch(priceMatch string) *WsCreateOrder {
	s.c.setParams("priceMatch", priceMatch)
	return s
}
func (s *WsCreateOrder) SelfTradePreventionMode(selfTradePreventionMode core.STPModeEnum) *WsCreateOrder {
	s.c.setParams("selfTradePreventionMode", selfTradePreventionMode)
	return s
}
func (s *WsCreateOrder) GoodTillDate(goodTillDate int64) *WsCreateOrder {
	s.c.setParams("goodTillDate", goodTillDate)
	return s
}
func (s *WsCreateOrder) RecvWindow(recvWindow int64) *WsCreateOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *WsCreateOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
}

func (s *WsModifyOrder) OrderId(orderId int64) *WsModifyOrder {
	s.c.setParams("orderId", orderId)
	return s
}

func (s *WsModifyOrder) OrigClientOrderId(origClientOrderId string) *WsModifyOrder {
	s.c.setParams("origClientOrderId", origClientOrderId)
	return s
}

func (s *WsModifyOrder) Symbol(symbol string) *WsModifyOrder {
	s.c.setParams("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *WsModifyOrder) Side(side core.OrderSideEnum) *WsModifyOrder {
	s.c.setParams("side", side)
	return s
}

func (s *WsModifyOrder) Quantity(quantity float64) *WsModifyOrder {
	s.c.setParams("quantity", quantity)
	return s
}

func (s *WsModifyOrder) Price(price float64) *WsModifyOrder {
	s.c.setParams("price", price)
	return s
}

func (s *WsModifyOrder) PriceMatch(priceMatch string) *WsModifyOrder {
	s.c.setParams("priceMatch", priceMatch)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *WsModifyOrder) RecvWindow(recvWindow int64) *WsModifyOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *WsModifyOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
}

func (s *WsCancelOrder) Symbol(symbol string) *WsCancelOrder {
	s.c.setParams("symbol", symbol)
	return s
}
func (s *WsCancelOrder) OrderId(orderId int64) *WsCancelOrder {
	s.c.setParams("orderId", orderId)
	return s
}
func (s *WsCancelOrder) OrigClientOrderId(origClientOrderId string) *WsCancelOrder {
	s.c.setParams("origClientOrderId", origClientOrderId)
	return s
}
func (s *WsCancelOrder) RecvWindow(recvWindow int64) *WsCancelOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *WsCancelOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
}

func (s *WsQueryOrder) Symbol(symbol string) *WsQueryOrder {
	s.c.setParams("symbol", symbol)
	return s
}
func (s *WsQueryOrder) OrderId(orderId int64) *WsQueryOrder {
	s.c.setParams("orderId", orderId)
	return s
}
func (s *WsQueryOrder) OrigClientOrderId(origClientOrderId string) *WsQueryOrder {
	s.c.setParams("origClientOrderId", origClientOrderId)
	return s
}
func (s *WsQueryOrder) RecvWindow(recvWindow int64) *WsQueryOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}
func (s *WsQueryOrder) Do(ctx context.Context) (*WsOrderResponse, error) {
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
	s.c.setParams("symbol", symbol)
	return s
}
func (s *WsPositionInfo) RecvWindow(recvWindow int64) *WsPositionInfo {
	s.c.setParams("recvWindow", recvWindow)
	return s
}
func (s *WsPositionInfo) Do(ctx context.Context) (*PositionInfoResponse, error) {
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
			var resp *PositionInfoResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
