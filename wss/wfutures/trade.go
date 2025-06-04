package wfutures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
	"github.com/jekaxv/go-binance/wss"
	"github.com/shopspring/decimal"
)

// CreateOrder Send in a new order.
type CreateOrder struct {
	c *Client
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

type OrderResponse struct {
	wss.ApiResponse
	Result *OrderResult `json:"result"`
}

func (s *CreateOrder) Symbol(symbol string) *CreateOrder {
	s.c.setParams("symbol", symbol)
	return s
}

func (s *CreateOrder) Side(side types.OrderSideEnum) *CreateOrder {
	s.c.setParams("side", side)
	return s
}
func (s *CreateOrder) PositionSide(positionSide types.PositionSideEnum) *CreateOrder {
	s.c.setParams("positionSide", positionSide)
	return s
}

func (s *CreateOrder) Type(orderType types.OrderTypeEnum) *CreateOrder {
	s.c.setParams("type", orderType)
	return s
}

func (s *CreateOrder) TimeInForce(timeInForce types.TimeInForceEnum) *CreateOrder {
	s.c.setParams("timeInForce", timeInForce)
	return s
}
func (s *CreateOrder) Quantity(quantity float64) *CreateOrder {
	s.c.setParams("quantity", quantity)
	return s
}
func (s *CreateOrder) ReduceOnly(reduceOnly string) *CreateOrder {
	s.c.setParams("reduceOnly", reduceOnly)
	return s
}

func (s *CreateOrder) Price(price float64) *CreateOrder {
	s.c.setParams("price", price)
	return s
}
func (s *CreateOrder) NewClientOrderId(newClientOrderId string) *CreateOrder {
	s.c.setParams("newClientOrderId", newClientOrderId)
	return s
}
func (s *CreateOrder) StopPrice(stopPrice float64) *CreateOrder {
	s.c.setParams("stopPrice", stopPrice)
	return s
}
func (s *CreateOrder) ClosePosition(closePosition string) *CreateOrder {
	s.c.setParams("closePosition", closePosition)
	return s
}
func (s *CreateOrder) ActivationPrice(activationPrice float64) *CreateOrder {
	s.c.setParams("activationPrice", activationPrice)
	return s
}
func (s *CreateOrder) CallbackRate(callbackRate float64) *CreateOrder {
	s.c.setParams("callbackRate", callbackRate)
	return s
}
func (s *CreateOrder) WorkingType(workingType types.WorkingType) *CreateOrder {
	s.c.setParams("workingType", workingType)
	return s
}
func (s *CreateOrder) PriceProtect(priceProtect string) *CreateOrder {
	s.c.setParams("priceProtect", priceProtect)
	return s
}
func (s *CreateOrder) NewOrderRespType(newOrderRespType types.OrderResponseTypeEnum) *CreateOrder {
	s.c.setParams("newOrderRespType", newOrderRespType)
	return s
}

func (s *CreateOrder) PriceMatch(priceMatch string) *CreateOrder {
	s.c.setParams("priceMatch", priceMatch)
	return s
}
func (s *CreateOrder) SelfTradePreventionMode(selfTradePreventionMode types.STPModeEnum) *CreateOrder {
	s.c.setParams("selfTradePreventionMode", selfTradePreventionMode)
	return s
}
func (s *CreateOrder) GoodTillDate(goodTillDate int64) *CreateOrder {
	s.c.setParams("goodTillDate", goodTillDate)
	return s
}
func (s *CreateOrder) RecvWindow(recvWindow int64) *CreateOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *CreateOrder) Do(ctx context.Context) (*OrderResponse, error) {
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

// ModifyOrder Order modify function, currently only LIMIT order modification is supported, modified orders will be reordered in the match queue
type ModifyOrder struct {
	c *Client
}

func (s *ModifyOrder) OrderId(orderId int64) *ModifyOrder {
	s.c.setParams("orderId", orderId)
	return s
}

func (s *ModifyOrder) OrigClientOrderId(origClientOrderId string) *ModifyOrder {
	s.c.setParams("origClientOrderId", origClientOrderId)
	return s
}

func (s *ModifyOrder) Symbol(symbol string) *ModifyOrder {
	s.c.setParams("symbol", symbol)
	return s
}

// Side BUY or SELL
func (s *ModifyOrder) Side(side types.OrderSideEnum) *ModifyOrder {
	s.c.setParams("side", side)
	return s
}

func (s *ModifyOrder) Quantity(quantity float64) *ModifyOrder {
	s.c.setParams("quantity", quantity)
	return s
}

func (s *ModifyOrder) Price(price float64) *ModifyOrder {
	s.c.setParams("price", price)
	return s
}

func (s *ModifyOrder) PriceMatch(priceMatch string) *ModifyOrder {
	s.c.setParams("priceMatch", priceMatch)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *ModifyOrder) RecvWindow(recvWindow int64) *ModifyOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *ModifyOrder) Do(ctx context.Context) (*OrderResponse, error) {
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

// CancelOrder Cancel an active order.
type CancelOrder struct {
	c *Client
}

func (s *CancelOrder) Symbol(symbol string) *CancelOrder {
	s.c.setParams("symbol", symbol)
	return s
}
func (s *CancelOrder) OrderId(orderId int64) *CancelOrder {
	s.c.setParams("orderId", orderId)
	return s
}
func (s *CancelOrder) OrigClientOrderId(origClientOrderId string) *CancelOrder {
	s.c.setParams("origClientOrderId", origClientOrderId)
	return s
}
func (s *CancelOrder) RecvWindow(recvWindow int64) *CancelOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *CancelOrder) Do(ctx context.Context) (*OrderResponse, error) {
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

// QueryOrder Check an order's status.
type QueryOrder struct {
	c *Client
}

func (s *QueryOrder) Symbol(symbol string) *QueryOrder {
	s.c.setParams("symbol", symbol)
	return s
}
func (s *QueryOrder) OrderId(orderId int64) *QueryOrder {
	s.c.setParams("orderId", orderId)
	return s
}
func (s *QueryOrder) OrigClientOrderId(origClientOrderId string) *QueryOrder {
	s.c.setParams("origClientOrderId", origClientOrderId)
	return s
}
func (s *QueryOrder) RecvWindow(recvWindow int64) *QueryOrder {
	s.c.setParams("recvWindow", recvWindow)
	return s
}
func (s *QueryOrder) Do(ctx context.Context) (*OrderResponse, error) {
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

// PositionInfo Get current position information(only symbol that has position or open orders will be returned).
type PositionInfo struct {
	c *Client
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
	wss.ApiResponse
	Result []*PositionInfoResult `json:"result"`
}

func (s *PositionInfo) Symbol(symbol string) *PositionInfo {
	s.c.setParams("symbol", symbol)
	return s
}
func (s *PositionInfo) RecvWindow(recvWindow int64) *PositionInfo {
	s.c.setParams("recvWindow", recvWindow)
	return s
}
func (s *PositionInfo) Do(ctx context.Context) (*PositionInfoResponse, error) {
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
