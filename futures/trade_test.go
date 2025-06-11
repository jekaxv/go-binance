package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiTradeTestSuite struct {
	baseHttpTestSuite
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
		Side(core.OrderSideBUY).
		Type(core.OrderTypeMARKET).
		Quantity("10").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertCreateOrderResponse(resp, testResp)
}

func (s *apiTradeTestSuite) assertCreateOrderResponse(r1, r2 *OrderResponse) {
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
		s.assertCreateOrderResponse(&resp[i].OrderResponse, &testResp[i].OrderResponse)
		r.Equal(testResp[i].Code, resp[i].Code, "code")
		r.Equal(testResp[i].Msg, resp[i].Msg, "msg")
	}
}

func (s *apiTradeTestSuite) TestNewModifyOrder() {
	msg := []byte(`{
		"orderId": 20072994037,
		"symbol": "BTCUSDT",
		"pair": "BTCUSDT",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumBase": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"closePosition": false,
		"side": "BUY",
		"positionSide": "LONG",
		"stopPrice": "0",
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false,
		"origType": "LIMIT",
		"priceMatch": "NONE",              
		"selfTradePreventionMode": "NONE", 
		"goodTillDate": 0,                 
		"updateTime": 1629182711600
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewModifyOrder().
		Symbol("BTCUSDT").
		OrderId(20072994037).
		Side(core.OrderSideBUY).
		Price("30005").
		Quantity("10").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ModifyOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertModifyOrderResponse(resp, testResp)
}

func (s *apiTradeTestSuite) assertModifyOrderResponse(r1, r2 *ModifyOrderResponse) {
	r := s.r()
	r.Equal(r2.OrderId, r1.OrderId, "OrderId")
	r.Equal(r2.Symbol, r1.Symbol, "Symbol")
	r.Equal(r2.Pair, r1.Pair, "Pair")
	r.Equal(r2.Status, r1.Status, "Status")
	r.Equal(r2.ClientOrderId, r1.ClientOrderId, "ClientOrderId")
	r.Equal(r2.Price, r1.Price, "Price")
	r.Equal(r2.AvgPrice, r1.AvgPrice, "AvgPrice")
	r.Equal(r2.OrigQty, r1.OrigQty, "OrigQty")
	r.Equal(r2.ExecutedQty, r1.ExecutedQty, "ExecutedQty")
	r.Equal(r2.CumQty, r1.CumQty, "CumQty")
	r.Equal(r2.CumBase, r1.CumBase, "CumBase")
	r.Equal(r2.TimeInForce, r1.TimeInForce, "TimeInForce")
	r.Equal(r2.Type, r1.Type, "Type")
	r.Equal(r2.ReduceOnly, r1.ReduceOnly, "ReduceOnly")
	r.Equal(r2.ClosePosition, r1.ClosePosition, "ClosePosition")
	r.Equal(r2.Side, r1.Side, "Side")
	r.Equal(r2.PositionSide, r1.PositionSide, "PositionSide")
	r.Equal(r2.StopPrice, r1.StopPrice, "StopPrice")
	r.Equal(r2.WorkingType, r1.WorkingType, "WorkingType")
	r.Equal(r2.PriceProtect, r1.PriceProtect, "PriceProtect")
	r.Equal(r2.OrigType, r1.OrigType, "OrigType")
	r.Equal(r2.PriceMatch, r1.PriceMatch, "PriceMatch")
	r.Equal(r2.SelfTradePreventionMode, r1.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r2.GoodTillDate, r1.GoodTillDate, "GoodTillDate")
	r.Equal(r2.UpdateTime, r1.UpdateTime, "UpdateTime")
}

func (s *apiTradeTestSuite) TestNewModifyMultipleOrder() {
	msg := []byte(`[
		{
			"orderId": 20072994037,
			"symbol": "BTCUSDT",
			"pair": "BTCUSDT",
			"status": "NEW",
			"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
			"price": "30005",
			"avgPrice": "0.0",
			"origQty": "1",
			"executedQty": "0",
			"cumQty": "0",
			"cumBase": "0",
			"timeInForce": "GTC",
			"type": "LIMIT",
			"reduceOnly": false,
			"closePosition": false,
			"side": "BUY",
			"positionSide": "LONG",
			"stopPrice": "0",
			"workingType": "CONTRACT_PRICE",
			"priceProtect": false,
			"origType": "LIMIT",
			"priceMatch": "NONE",              
			"selfTradePreventionMode": "NONE",
			"goodTillDate": 0,                
			"updateTime": 1629182711600
		},
		{
			"code": -2022, 
			"msg": "ReduceOnly Order is rejected."
		}
	]`)
	server := s.setup(msg)
	defer server.Close()
	orders := []ModifyOrderReq{{
		Symbol:   "BTCUSDT",
		OrderId:  20072994037,
		Side:     "BUY",
		Price:    "30005",
		Quantity: "10",
	}}

	resp, err := s.client.NewModifyMultipleOrder().BatchOrders(orders).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*ModifyMultipleOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertModifyOrderResponse(&resp[i].ModifyOrderResponse, &testResp[i].ModifyOrderResponse)
		r.Equal(testResp[i].Code, resp[i].Code, "code")
		r.Equal(testResp[i].Msg, resp[i].Msg, "msg")
	}
}

func (s *apiTradeTestSuite) TestNewOrderAmendment() {
	msg := []byte(`[
	{
		"amendmentId": 5325,
		"symbol": "BTCUSDT",
		"pair": "BTCUSDT",
		"orderId": 20072994037,
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"time": 1629182711787,
		"amendment": {
			"price": {
				"before": "30002",
				"after": "30005"
			},
			"origQty": {
				"before": "1",
				"after": "1"
			},
			"count": 1
		}
	}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOrderAmendment().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrderAmendmentResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertOrderAmendmentResponse(resp[i], testResp[i])
	}
}

func (s *apiTradeTestSuite) assertOrderAmendmentResponse(r1, r2 *OrderAmendmentResponse) {
	r := s.r()
	r.Equal(r2.AmendmentId, r1.AmendmentId, "AmendmentId")
	r.Equal(r2.Symbol, r1.Symbol, "Symbol")
	r.Equal(r2.Pair, r1.Pair, "Pair")
	r.Equal(r2.OrderId, r1.OrderId, "OrderId")
	r.Equal(r2.ClientOrderId, r1.ClientOrderId, "ClientOrderId")
	r.Equal(r2.Time, r1.Time, "Time")
	r.Equal(r1.Amendment.Price.After, r2.Amendment.Price.After, "Price.After")
	r.Equal(r1.Amendment.Price.Before, r2.Amendment.Price.Before, "Price.Before")
	r.Equal(r1.Amendment.OrigQty.After, r2.Amendment.OrigQty.After, "OrigQty.After")
	r.Equal(r1.Amendment.OrigQty.Before, r2.Amendment.OrigQty.Before, "OrigQty.Before")
	r.Equal(r1.Amendment.Count, r2.Amendment.Count, "Count")
}

func (s *apiTradeTestSuite) TestCancelOrder() {
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
	resp, err := s.client.NewCancelOrder().Symbol("BTCUSDT").
		OrderId(22542179).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertCreateOrderResponse(resp, testResp)
}
func (s *apiTradeTestSuite) TestNewCancelMultipleOrder() {
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
	resp, err := s.client.NewCancelMultipleOrder().Symbol("BTCUSDT").
		OrderIdList([]int64{22542179, 2344}).
		OrigClientOrderIdList([]string{"testOrder", "testOrder2"}).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertCreateOrderResponse(resp[i], testResp[i])
	}
}

func (s *apiTradeTestSuite) TestNewCancelOpenOrder() {
	msg := []byte(`{
		"code": 200, 
		"msg": "The operation of cancel all open order is done."
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOpenOrder().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CancelOpenOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Code, testResp.Code, "code")
	r.Equal(resp.Msg, testResp.Msg, "msg")
}

func (s *apiTradeTestSuite) TestNewCountdownCancelAll() {
	msg := []byte(`{
		"symbol": "BTCUSDT", 
		"countdownTime": "100000"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCountdownCancelAll().Symbol("BTCUSDT").
		CountdownTime(1000).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CountdownCancelAllResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Symbol, testResp.Symbol, "Symbol")
	r.Equal(resp.CountdownTime, testResp.CountdownTime, "CountdownTime")
}

func (s *apiTradeTestSuite) TestNewQueryOrder() {
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
	resp, err := s.client.NewQueryOrder().Symbol("BTCUSDT").
		OrderId(22542179).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertCreateOrderResponse(resp, testResp)
}
func (s *apiTradeTestSuite) TestNewQueryAllOrder() {
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
	resp, err := s.client.NewQueryAllOrder().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertCreateOrderResponse(resp[i], testResp[i])
	}
}
func (s *apiTradeTestSuite) TestNewAllOpenOrder() {
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
	resp, err := s.client.NewAllOpenOrder().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertCreateOrderResponse(resp[i], testResp[i])
	}
}

func (s *apiTradeTestSuite) TestNewQueryOpenOrder() {
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
	resp, err := s.client.NewQueryOpenOrder().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertCreateOrderResponse(resp, testResp)
}

func (s *apiTradeTestSuite) TestNewForceOrder() {
	msg := []byte(`[{
		"orderId": 20072994037,
		"symbol": "BTCUSDT",
		"pair": "BTCUSDT",
		"status": "NEW",
		"clientOrderId": "LJ9R4QZDihCaS8UAOOLpgW",
		"price": "30005",
		"avgPrice": "0.0",
		"origQty": "1",
		"executedQty": "0",
		"cumQty": "0",
		"cumBase": "0",
		"timeInForce": "GTC",
		"type": "LIMIT",
		"reduceOnly": false,
		"closePosition": false,
		"side": "BUY",
		"positionSide": "LONG",
		"stopPrice": "0",
		"workingType": "CONTRACT_PRICE",
		"priceProtect": false,
		"origType": "LIMIT",
		"priceMatch": "NONE",              
		"selfTradePreventionMode": "NONE", 
		"goodTillDate": 0,                 
		"updateTime": 1629182711600
	}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewForceOrder().
		Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*ModifyOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertModifyOrderResponse(resp[i], testResp[i])
	}
}
func (s *apiTradeTestSuite) TestNewUserTrades() {
	msg := []byte(`[
	  {
		"buyer": false,
		"commission": "-0.07819010",
		"commissionAsset": "USDT",
		"id": 698759,
		"maker": false,
		"orderId": 25851813,
		"price": "7819.01",
		"qty": "0.002",
		"quoteQty": "15.63802",
		"realizedPnl": "-0.91539999",
		"side": "SELL",
		"positionSide": "SHORT",
		"symbol": "BTCUSDT",
		"time": 1569514978020
	  }
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewUserTrades().
		Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*UserTradesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertUserTradesResponse(resp[i], testResp[i])
	}
}

func (s *apiTradeTestSuite) assertUserTradesResponse(r1, r2 *UserTradesResponse) {
	r := s.r()
	r.Equal(r2.Buyer, r1.Buyer, "Buyer")
	r.Equal(r2.Commission, r1.Commission, "Commission")
	r.Equal(r2.CommissionAsset, r1.CommissionAsset, "CommissionAsset")
	r.Equal(r2.Id, r1.Id, "Id")
	r.Equal(r2.Maker, r1.Maker, "Maker")
	r.Equal(r2.OrderId, r1.OrderId, "OrderId")
	r.Equal(r2.Price, r1.Price, "Price")
	r.Equal(r2.Qty, r1.Qty, "Qty")
	r.Equal(r2.QuoteQty, r1.QuoteQty, "QuoteQty")
	r.Equal(r2.RealizedPnl, r1.RealizedPnl, "RealizedPnl")
	r.Equal(r2.Side, r1.Side, "Side")
	r.Equal(r2.PositionSide, r1.PositionSide, "PositionSide")
	r.Equal(r2.Symbol, r1.Symbol, "Symbol")
	r.Equal(r2.Time, r1.Time, "Time")
}
func (s *apiTradeTestSuite) TestNewChangeMarginType() {
	msg := []byte(`{
		"code": 200, 
		"msg": "The operation of cancel all open order is done."
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewChangeMarginType().Symbol("BTCUSDT").
		MarginType(core.MarginTypeISOLATED).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ChangeMarginTypeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Code, testResp.Code, "code")
	r.Equal(resp.Msg, testResp.Msg, "msg")
}

func (s *apiTradeTestSuite) TestNewChangePositionSide() {
	msg := []byte(`{
		"code": 200, 
		"msg": "success"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewChangePositionSide().DualSidePosition("true").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ChangeMarginTypeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Code, testResp.Code, "code")
	r.Equal(resp.Msg, testResp.Msg, "msg")
}
func (s *apiTradeTestSuite) TestNewChangeLeverage() {
	msg := []byte(`{
		"leverage": 21,
		"maxNotionalValue": "1000000",
		"symbol": "BTCUSDT"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewChangeLeverage().Symbol("BTCUSDT").
		Leverage(20).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ChangeLeverageResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Leverage, testResp.Leverage, "leverage")
	r.Equal(resp.MaxNotionalValue, testResp.MaxNotionalValue, "maxNotionalValue")
	r.Equal(resp.Symbol, testResp.Symbol, "symbol")
}

func (s *apiTradeTestSuite) TestNewChangeMultiAssetsMargin() {
	msg := []byte(`{
		"code": 200, 
		"msg": "success"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewChangeMultiAssetsMargin().MultiAssetsMargin("true").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ChangeMarginTypeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Code, testResp.Code, "code")
	r.Equal(resp.Msg, testResp.Msg, "msg")
}

func (s *apiTradeTestSuite) TestNewChangePositionMargin() {
	msg := []byte(`{
		"amount": 100.0,
		"code": 200,
		"msg": "Successfully modify position margin.",
		"type": 1
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewChangePositionMargin().Symbol("BTCUSDT").
		Amount("100").
		Type(1).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ChangeMarginTypeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Code, testResp.Code, "code")
	r.Equal(resp.Msg, testResp.Msg, "msg")
}
func (s *apiTradeTestSuite) TestNewPositionRisk() {
	msg := []byte(`[
	  {
			"symbol": "ADAUSDT",
			"positionSide": "BOTH",               
			"positionAmt": "30",
			"entryPrice": "0.385",
			"breakEvenPrice": "0.385077",
			"markPrice": "0.41047590",
			"unRealizedProfit": "0.76427700",      
			"liquidationPrice": "0",
			"isolatedMargin": "0",
			"notional": "12.31427700",
			"marginAsset": "USDT",
			"isolatedWallet": "0",
			"initialMargin": "0.61571385",         
			"maintMargin": "0.08004280",          
			"positionInitialMargin": "0.61571385",
			"openOrderInitialMargin": "0",         
			"adl": 2,
			"bidNotional": "0",                   
			"askNotional": "0",                   
			"updateTime": 1720736417660
	  }
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewPositionRisk().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*PositionRiskResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertPositionRiskResponse(resp[i], testResp[i])
	}
}
func (s *apiTradeTestSuite) assertPositionRiskResponse(r1, r2 *PositionRiskResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PositionSide, r2.PositionSide, "PositionSide")
	r.Equal(r1.PositionAmt, r2.PositionAmt, "PositionAmt")
	r.Equal(r1.EntryPrice, r2.EntryPrice, "EntryPrice")
	r.Equal(r1.BreakEvenPrice, r2.BreakEvenPrice, "BreakEvenPrice")
	r.Equal(r1.MarkPrice, r2.MarkPrice, "MarkPrice")
	r.Equal(r1.UnRealizedProfit, r2.UnRealizedProfit, "UnRealizedProfit")
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

func (s *apiTradeTestSuite) TestNewAdlQuantile() {
	msg := []byte(`[
	  {
		"symbol": "ETHUSDT",
		"adlQuantile": {
		  "LONG": 3,
		  "SHORT": 3,
		  "HEDGE": 0
		}
	  },
	  {
		"symbol": "BTCUSDT",
		"adlQuantile": {
		  "LONG": 1,
		  "SHORT": 2,
		  "BOTH": 0
		}
	  }
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAdlQuantile().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*AdlQuantileResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		r.Equal(resp[i].Symbol, testResp[i].Symbol, "Symbol")
		r.Equal(resp[i].AdlQuantile.HEDGE, testResp[i].AdlQuantile.HEDGE, "AdlQuantile.HEDGE")
		r.Equal(resp[i].AdlQuantile.LONG, testResp[i].AdlQuantile.LONG, "AdlQuantile.LONG")
		r.Equal(resp[i].AdlQuantile.SHORT, testResp[i].AdlQuantile.SHORT, "AdlQuantile.SHORT")
		r.Equal(resp[i].AdlQuantile.BOTH, testResp[i].AdlQuantile.BOTH, "AdlQuantile.BOTH")
	}
}

func (s *apiTradeTestSuite) TestNewPositionMarginHistory() {
	msg := []byte(`[
		{
			"symbol": "BTCUSDT",
			"type": 1,
			"deltaType": "USER_ADJUST",
			"amount": "23.36332311",
			"asset": "USDT",
			"time": 1578047897183,
			"positionSide": "BOTH"
		},
		{
			"symbol": "BTCUSDT",
			"type": 1, 
			"deltaType": "USER_ADJUST",
			"amount": "100",
			"asset": "USDT",
			"time": 1578047900425,
			"positionSide": "LONG" 
		}
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewPositionMarginHistory().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*PositionMarginHistoryResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertPositionMarginHistoryResponse(resp[i], testResp[i])
	}
}

func (s *apiTradeTestSuite) assertPositionMarginHistoryResponse(r1, r2 *PositionMarginHistoryResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Type, r2.Type, "Type")
	r.Equal(r1.DeltaType, r2.DeltaType, "DeltaType")
	r.Equal(r1.Amount, r2.Amount, "Amount")
	r.Equal(r1.Asset, r2.Asset, "Asset")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.PositionSide, r2.PositionSide, "PositionSide")
}

func (s *apiTradeTestSuite) TestCreatTestOrder() {
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
	resp, err := s.client.NewCreateTestOrder().Symbol("BTCUSDT").
		Side(core.OrderSideBUY).
		Type(core.OrderTypeMARKET).
		Quantity("10").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertCreateOrderResponse(resp, testResp)
}
