package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type tradeTestSuite struct {
	baseWsTestSuite
}

func TestWebsocketTrade(t *testing.T) {
	suite.Run(t, new(tradeTestSuite))
}

func (s *tradeTestSuite) TestNewCreateOrder() {
	msg := []byte(`{
		"id": "3f7df6e3-2df4-44b9-9919-d2f38f90a99a",
		"status": 200,
		"result": {
			"orderId": 325078477,
			"symbol": "BTCUSDT",
			"status": "NEW",
			"clientOrderId": "iCXL1BywlBaf2sesNUrVl3",
			"price": "43187.00",
			"avgPrice": "0.00",
			"origQty": "0.100",
			"executedQty": "0.000",
			"cumQty": "0.000",
			"cumQuote": "0.00000",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"reduceOnly": false,
			"closePosition": false,
			"side": "BUY",
			"positionSide": "BOTH",
			"stopPrice": "0.00",
			"workingType": "CONTRACT_PRICE",
			"priceProtect": false,
			"origType": "LIMIT",
			"priceMatch": "NONE",
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"updateTime": 1702555534435
		},
		"rateLimits": [
			{
				"rateLimitType": "ORDERS",
				"interval": "SECOND",
				"intervalNum": 10,
				"limit": 300,
				"count": 1
			},
			{
				"rateLimitType": "ORDERS",
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 1200,
				"count": 1
			},
			{
				"rateLimitType": "REQUEST_WEIGHT",
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 2400,
				"count": 1
			}
		]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOrder().
		Symbol("BTCUSDT").
		Type(core.OrderTypeLIMIT).
		PositionSide(core.PositionSide_BOTH).
		Price(11).
		Quantity(0.1).
		Side(core.OrderSideBUY).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderResponse(resp, testResp)
}

func (s *tradeTestSuite) assertTestOrderResponse(r1, r2 *WsOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertCreateOrderResult(r1.Result, r2.Result)
}

func (s *tradeTestSuite) assertCreateOrderResult(r1, r2 *OrderResult) {
	r := s.r()
	r.Equal(r2.ClientOrderId, r1.ClientOrderId, "ClientOrderId")
	r.Equal(r2.CumQty, r1.CumQty, "CumQty")
	r.Equal(r2.CumQuote, r1.CumQuote, "CumQuote")
	r.Equal(r2.ExecutedQty, r1.ExecutedQty, "ExecutedQty")
	r.Equal(r2.OrderId, r1.OrderId, "OrderId")
	r.Equal(r2.AvgPrice, r1.AvgPrice, "AvgPrice")
	r.Equal(r2.OrigQty, r1.OrigQty, "OrigQty")
	r.Equal(r2.Price, r1.Price, "Price")
	r.Equal(r2.ReduceOnly, r1.ReduceOnly, "ReduceOnly")
	r.Equal(r2.Side, r1.Side, "Side")
	r.Equal(r2.PositionSide, r1.PositionSide, "PositionSide")
	r.Equal(r2.Status, r1.Status, "Status")
	r.Equal(r2.StopPrice, r1.StopPrice, "StopPrice")
	r.Equal(r2.ClosePosition, r1.ClosePosition, "ClosePosition")
	r.Equal(r2.Symbol, r1.Symbol, "Symbol")
	r.Equal(r2.TimeInForce, r1.TimeInForce, "TimeInForce")
	r.Equal(r2.Type, r1.Type, "Type")
	r.Equal(r2.OrigType, r1.OrigType, "OrigType")
	r.Equal(r2.UpdateTime, r1.UpdateTime, "UpdateTime")
	r.Equal(r2.WorkingType, r1.WorkingType, "WorkingType")
	r.Equal(r2.PriceProtect, r1.PriceProtect, "PriceProtect")
	r.Equal(r2.PriceMatch, r1.PriceMatch, "PriceMatch")
	r.Equal(r2.SelfTradePreventionMode, r1.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r2.GoodTillDate, r1.GoodTillDate, "GoodTillDate")
}

func (s *tradeTestSuite) TestNewModifyOrder() {
	msg := []byte(`{
		"id": "3f7df6e3-2df4-44b9-9919-d2f38f90a99a",
		"status": 200,
		"result": {
			"orderId": 325078477,
			"symbol": "BTCUSDT",
			"status": "NEW",
			"clientOrderId": "iCXL1BywlBaf2sesNUrVl3",
			"price": "43187.00",
			"avgPrice": "0.00",
			"origQty": "0.100",
			"executedQty": "0.000",
			"cumQty": "0.000",
			"cumQuote": "0.00000",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"reduceOnly": false,
			"closePosition": false,
			"side": "BUY",
			"positionSide": "BOTH",
			"stopPrice": "0.00",
			"workingType": "CONTRACT_PRICE",
			"priceProtect": false,
			"origType": "LIMIT",
			"priceMatch": "NONE",
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,
			"updateTime": 1702555534435
		},
		"rateLimits": [
			{
				"rateLimitType": "ORDERS",
				"interval": "SECOND",
				"intervalNum": 10,
				"limit": 300,
				"count": 1
			},
			{
				"rateLimitType": "ORDERS",
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 1200,
				"count": 1
			},
			{
				"rateLimitType": "REQUEST_WEIGHT",
				"interval": "MINUTE",
				"intervalNum": 1,
				"limit": 2400,
				"count": 1
			}
		]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewModifyOrder().
		OrderId(1).
		Symbol("BTCUSDT").
		Price(11).
		Quantity(0.1).
		Side(core.OrderSideBUY).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderResponse(resp, testResp)
}

func (s *tradeTestSuite) TestNewCancelOrder() {
	msg := []byte(`{
	  "id": "5633b6a2-90a9-4192-83e7-925c90b6a2fd",
	  "status": 200,
	  "result": {
		"clientOrderId": "myOrder1",
		"cumQty": "0",
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 283194212,
		"origQty": "11",
		"origType": "TRAILING_STOP_MARKET",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"status": "CANCELED",
		"stopPrice": "9300",                
		"closePosition": false,  
		"symbol": "BTCUSDT",
		"timeInForce": "GTC",
		"type": "TRAILING_STOP_MARKET",
		"activatePrice": "9020",            
		"priceRate": "0.3",                
		"updateTime": 1571110484038,
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false,           
		"priceMatch": "NONE",              
		"selfTradePreventionMode": "NONE",
		"goodTillDate": 0                 
	  },
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 2400,
		  "count": 1
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOrder().
		OrderId(1).
		Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderResponse(resp, testResp)
}

func (s *tradeTestSuite) TestNewQueryOrder() {
	msg := []byte(`{
	  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",
	  "status": 200,
	  "result": {
		"avgPrice": "0.00000",
		"clientOrderId": "abc",
		"cumQuote": "0",
		"executedQty": "0",
		"orderId": 1917641,
		"origQty": "0.40",
		"origType": "TRAILING_STOP_MARKET",
		"price": "0",
		"reduceOnly": false,
		"side": "BUY",
		"positionSide": "SHORT",
		"status": "NEW",
		"stopPrice": "9300",
		"closePosition": false,
		"symbol": "BTCUSDT",
		"time": 1579276756075,
		"timeInForce": "GTC",
		"type": "TRAILING_STOP_MARKET",
		"activatePrice": "9020",
		"priceRate": "0.3",
		"updateTime": 1579276756075,
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false
	  }
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOrder().
		OrderId(1).
		Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderResponse(resp, testResp)
}

func (s *tradeTestSuite) TestNewPositionInfo() {
	msg := []byte(`{
	  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",
	  "status": 200,
	  "result": [
		{
		  "symbol": "BTCUSDT",
		  "positionSide": "BOTH",
		  "positionAmt": "1.000",
		  "entryPrice": "0.00000",
		  "breakEvenPrice": "0.0",
		  "markPrice": "6679.50671178",
		  "unrealizedProfit": "0.00000000",
		  "liquidationPrice": "0",
		  "isolatedMargin": "0.00000000",
		  "notional": "0",
		  "marginAsset": "USDT",
		  "isolatedWallet": "0",
		  "initialMargin": "0",
		  "maintMargin": "0",
		  "positionInitialMargin": "0",
		  "openOrderInitialMargin": "0",
		  "adl": 0,
		  "bidNotional": "0",
		  "askNotional": "0",
		  "updateTime": 0
		}
	  ],
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 2400,
		  "count": 20
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewPositionInfo().
		Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *PositionInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestPositionInfoResponse(resp, testResp)
}

func (s *tradeTestSuite) assertTestPositionInfoResponse(r1, r2 *PositionInfoResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertPositionInfoResult(r1.Result[i], r2.Result[i])
	}
}

func (s *tradeTestSuite) assertPositionInfoResult(r1, r2 *PositionInfoResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PositionSide, r2.PositionSide, "PositionSide")
	r.Equal(r1.PositionAmt, r2.PositionAmt, "PositionAmt")
	r.Equal(r1.EntryPrice, r2.EntryPrice, "EntryPrice")
	r.Equal(r1.BreakEvenPrice, r2.BreakEvenPrice, "BreakEvenPrice")
	r.Equal(r1.MarkPrice, r2.MarkPrice, "MarkPrice")
	r.Equal(r1.UnrealizedProfit, r2.UnrealizedProfit, "UnrealizedProfit")
	r.Equal(r1.LiquidationPrice, r2.LiquidationPrice, "LiquidationPrice")
	r.Equal(r1.IsolatedMargin, r2.IsolatedMargin, "IsolatedMargin")
	r.Equal(r1.Notional, r2.Notional, "Notional")
	r.Equal(r1.MarginAsset, r2.MarginAsset, "MarginAsset")
	r.Equal(r1.IsolatedWallet, r2.IsolatedWallet, "IsolatedWallet")
	r.Equal(r1.InitialMargin, r2.InitialMargin, "InitialMargin")
	r.Equal(r1.MaintMargin, r2.MaintMargin, "MaintMargin")
	r.Equal(r1.PositionInitialMargin, r2.PositionInitialMargin, "PositionInitialMargin")
	r.Equal(r1.OpenOrderInitialMargin, r2.OpenOrderInitialMargin, "OpenOrderInitialMargin")
	r.Equal(r1.Adl, r2.Adl, "Adl")
	r.Equal(r1.BidNotional, r2.BidNotional, "BidNotional")
	r.Equal(r1.AskNotional, r2.AskNotional, "AskNotional")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "UpdateTime")
}
