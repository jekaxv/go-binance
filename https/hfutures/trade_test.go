package hfutures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiTradeTestSuite struct {
	baseTestSuite
}

func TestApiTradeAccount(t *testing.T) {
	suite.Run(t, new(apiTradeTestSuite))
}

func (s *apiTradeTestSuite) TestCreatOrder() {
	msg := []byte(`{
  "clientOrderId": "testOrder",
  "cumQty": "0",
  "cumQuote": "0",
  "executedQty": "0",
  "orderId": 22542179,
  "avgPrice": "0.00000",
  "origQty": "10",
  "price": "0",
  "reduceOnly": false,
  "side": "BUY",
  "positionSide": "SHORT",
  "status": "NEW",
  "stopPrice": "9300",
  "closePosition": false,
  "symbol": "BTCUSDT",
  "timeInForce": "GTD",
  "type": "TRAILING_STOP_MARKET",
  "origType": "TRAILING_STOP_MARKET",
  "activatePrice": "9020",
  "priceRate": "0.3",
  "updateTime": 1566818724722,
  "workingType": "CONTRACT_PRICE",
  "priceProtect": false,
  "priceMatch": "NONE",
  "selfTradePreventionMode": "NONE",
  "goodTillDate": 1693207680000
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).
		Type(types.OrderTypeMARKET).
		Quantity("10").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertCreateOrderResponse(resp, testResp)
}

func (s *baseTestSuite) assertCreateOrderResponse(r1, r2 *CreateOrderResponse) {
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
	r.Equal(r2.ActivatePrice, r1.ActivatePrice, "ActivatePrice")
	r.Equal(r2.PriceRate, r1.PriceRate, "PriceRate")
	r.Equal(r2.UpdateTime, r1.UpdateTime, "UpdateTime")
	r.Equal(r2.WorkingType, r1.WorkingType, "WorkingType")
	r.Equal(r2.PriceProtect, r1.PriceProtect, "PriceProtect")
	r.Equal(r2.PriceMatch, r1.PriceMatch, "PriceMatch")
	r.Equal(r2.SelfTradePreventionMode, r1.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r2.GoodTillDate, r1.GoodTillDate, "GoodTillDate")
}

func (s *apiTradeTestSuite) TestPlaceBatchOrder() {
	msg := []byte(`[{
  "clientOrderId": "testOrder",
  "cumQty": "0",
  "cumQuote": "0",
  "executedQty": "0",
  "orderId": 22542179,
  "avgPrice": "0.00000",
  "origQty": "10",
  "price": "0",
  "reduceOnly": false,
  "side": "BUY",
  "positionSide": "SHORT",
  "status": "NEW",
  "stopPrice": "9300",
  "closePosition": false,
  "symbol": "BTCUSDT",
  "timeInForce": "GTD",
  "type": "TRAILING_STOP_MARKET",
  "origType": "TRAILING_STOP_MARKET",
  "activatePrice": "9020",
  "priceRate": "0.3",
  "updateTime": 1566818724722,
  "workingType": "CONTRACT_PRICE",
  "priceProtect": false,
  "priceMatch": "NONE",
  "selfTradePreventionMode": "NONE",
  "goodTillDate": 1693207680000
}]`)
	server := s.setup(msg)
	defer server.Close()
	orders := []OrderReq{{
		Symbol:    "BTCUSDT",
		Side:      "BUY",
		OrderType: "MARKET",
		Quantity:  "10",
	}}
	resp, err := s.client.NewPlaceBatchOrder().BatchOrders(orders).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*PlaceBatchOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertCreateOrderResponse(&resp[i].CreateOrderResponse, &testResp[i].CreateOrderResponse)
		r.Equal(testResp[i].Code, resp[i].Code, "code")
		r.Equal(testResp[i].Msg, resp[i].Msg, "msg")
	}
}
