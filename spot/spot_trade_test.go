package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type spotTradeTestSuite struct {
	baseHttpTestSuite
}

func TestApiTradeAccount(t *testing.T) {
	suite.Run(t, new(spotTradeTestSuite))
}

func (s *spotTradeTestSuite) TestCreateOrder() {
	msg := []byte(`{
  "symbol": "BTCUSDT",
  "orderId": 6566216,
  "orderListId": -1,
  "clientOrderId": "LjKnBt1VlCWvR9we9040De",
  "transactTime": 1748095899515,
  "price": "0.00000000",
  "origQty": "0.00100000",
  "executedQty": "0.00100000",
  "origQuoteOrderQty": "0.00000000",
  "cummulativeQuoteQty": "108.73145000",
  "status": "FILLED",
  "timeInForce": "GTC",
  "type": "MARKET",
  "side": "BUY",
  "workingTime": 1748095899515,
  "fills": [
    {
      "price": "108731.45000000",
      "qty": "0.00100000",
      "commission": "0.00000000",
      "commissionAsset": "BTC",
      "tradeId": 3001146
    }
  ],
  "selfTradePreventionMode": "EXPIRE_MAKER"
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOrder().Symbol("BTCUSDT").
		Side(core.OrderSideBUY).
		Type(core.OrderTypeMARKET).
		Quantity("0.001").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateOrder(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestCreateOrder(r1, r2 *CreateOrderResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.OrderId, r2.OrderId, "OrderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "OrderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "ClientOrderId")
	r.Equal(r1.TransactTime, r2.TransactTime, "TransactTime")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.OrigQty, r2.OrigQty, "OrigQty")
	r.Equal(r1.ExecutedQty, r2.ExecutedQty, "ExecutedQty")
	r.Equal(r1.OrigQuoteOrderQty, r2.OrigQuoteOrderQty, "OrigQuoteOrderQty")
	r.Equal(r1.CummulativeQuoteQty, r2.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(r1.Status, r2.Status, "Status")
	r.Equal(r1.TimeInForce, r2.TimeInForce, "TimeInForce")
	r.Equal(r1.Type, r2.Type, "Type")
	r.Equal(r1.Side, r2.Side, "Side")
	r.Equal(r1.WorkingTime, r2.WorkingTime, "WorkingTime")
	for i := range r1.Fills {
		s.assertApiFill(r1.Fills[i], r2.Fills[i])
	}
}

func (s *spotTradeTestSuite) TestTestCreateOrder() {
	msg := []byte(`{
  "standardCommissionForOrder": {
    "maker": "0.00000000",
    "taker": "0.00000000"
  },
  "taxCommissionForOrder": {
    "maker": "0.00000000",
    "taker": "0.00000000"
  },
  "discount": {
    "enabledForAccount": true,
    "enabledForSymbol": true,
    "discountAsset": null,
    "discount": "0.00000000"
  }
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTestCreateOrder().Symbol("BTCUSDT").
		Side(core.OrderSideBUY).
		Type(core.OrderTypeMARKET).
		Quantity("0.001").
		ComputeCommissionRates(true).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *TestCreateOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateTestOrder(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestCreateTestOrder(r1, r2 *TestCreateOrderResponse) {
	r := s.r()
	r.Equal(r1.StandardCommissionForOrder.Maker, r2.StandardCommissionForOrder.Maker, "StandardCommissionForOrder.maker")
	r.Equal(r1.StandardCommissionForOrder.Taker, r2.StandardCommissionForOrder.Taker, "StandardCommissionForOrder.taker")
	r.Equal(r1.TaxCommissionForOrder.Maker, r2.TaxCommissionForOrder.Maker, "TaxCommissionForOrder.maker")
	r.Equal(r1.TaxCommissionForOrder.Taker, r2.TaxCommissionForOrder.Taker, "TaxCommissionForOrder.taker")
	r.Equal(r1.Discount.EnabledForAccount, r2.Discount.EnabledForAccount, "Discount.enabledForAccount")
	r.Equal(r1.Discount.EnabledForSymbol, r2.Discount.EnabledForSymbol, "Discount.enabledForSymbol")
	r.Equal(r1.Discount.DiscountAsset, r2.Discount.DiscountAsset, "Discount.discountAsset")
	r.Equal(r1.Discount.Discount, r2.Discount.Discount, "Discount.discount")
}

func (s *spotTradeTestSuite) TestQueryOrder() {
	msg := []byte(`{
  "symbol": "LTCBTC",
  "orderId": 1,
  "orderListId": -1,
  "clientOrderId": "myOrder1",
  "price": "0.1",
  "origQty": "1.0",
  "executedQty": "0.0",
  "cummulativeQuoteQty": "0.0",
  "status": "NEW",
  "timeInForce": "GTC",
  "type": "LIMIT",
  "side": "BUY",
  "stopPrice": "0.0",
  "icebergQty": "0.0",
  "time": 1499827319559,
  "updateTime": 1499827319559,
  "isWorking": true,
  "workingTime": 1499827319559,
  "origQuoteOrderQty": "0.000000",
  "selfTradePreventionMode": "NONE"
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOrder().Symbol("LTCBTC").
		OrderId(6566216).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestQueryOrder(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestQueryOrder(r1, r2 *QueryOrderResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.OrderId, r2.OrderId, "OrderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "OrderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "ClientOrderId")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.OrigQty, r2.OrigQty, "OrigQty")
	r.Equal(r1.ExecutedQty, r2.ExecutedQty, "ExecutedQty")
	r.Equal(r1.CummulativeQuoteQty, r2.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(r1.Status, r2.Status, "Status")
	r.Equal(r1.TimeInForce, r2.TimeInForce, "TimeInForce")
	r.Equal(r1.Type, r2.Type, "Type")
	r.Equal(r1.Side, r2.Side, "Side")
	r.Equal(r1.StopPrice, r2.StopPrice, "StopPrice")
	r.Equal(r1.IcebergQty, r2.IcebergQty, "IcebergQty")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "UpdateTime")
	r.Equal(r1.IsWorking, r2.IsWorking, "IsWorking")
	r.Equal(r1.WorkingTime, r2.WorkingTime, "WorkingTime")
	r.Equal(r1.OrigQuoteOrderQty, r2.OrigQuoteOrderQty, "OrigQuoteOrderQty")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "SelfTradePreventionMode")
}

func (s *spotTradeTestSuite) TestNewCancelOrder() {
	msg := []byte(`{
  "symbol": "LTCBTC",
  "origClientOrderId": "myOrder1",
  "orderId": 4,
  "orderListId": -1,
  "clientOrderId": "cancelMyOrder1",
  "transactTime": 1684804350068,
  "price": "2.00000000",
  "origQty": "1.00000000",
  "executedQty": "0.00000000",
  "cummulativeQuoteQty": "0.00000000",
  "status": "CANCELED",
  "timeInForce": "GTC",
  "type": "LIMIT",
  "side": "BUY",
  "selfTradePreventionMode": "NONE"
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOrder().Symbol("LTCBTC").
		OrderId(6566216).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestQueryOrder(resp, testResp)
}

func (s *spotTradeTestSuite) TestNewCancelOpenOrder() {
	msg := []byte(`[
  {
    "symbol": "BTCUSDT",
    "origClientOrderId": "E6APeyTJvkMvLMYMqu1KQ4",
    "orderId": 11,
    "orderListId": -1,
    "clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
    "transactTime": 1684804350068,
    "price": "0.089853",
    "origQty": "0.178622",
    "executedQty": "0.000000",
    "cummulativeQuoteQty": "0.000000",
    "status": "CANCELED",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "BUY",
    "selfTradePreventionMode": "NONE"
  },
  {
    "symbol": "BTCUSDT",
    "origClientOrderId": "A3EF2HCwxgZPFMrfwbgrhv",
    "orderId": 13,
    "orderListId": -1,
    "clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
    "transactTime": 1684804350069,
    "price": "0.090430",
    "origQty": "0.178622",
    "executedQty": "0.000000",
    "cummulativeQuoteQty": "0.000000",
    "status": "CANCELED",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "BUY",
    "selfTradePreventionMode": "NONE"
  },
  {
    "orderListId": 1929,
    "contingencyType": "OCO",
    "listStatusType": "ALL_DONE",
    "listOrderStatus": "ALL_DONE",
    "listClientOrderId": "2inzWQdDvZLHbbAmAozX2N",
    "transactionTime": 1585230948299,
    "symbol": "BTCUSDT",
    "orders": [
      {
        "symbol": "BTCUSDT",
        "orderId": 20,
        "clientOrderId": "CwOOIPHSmYywx6jZX77TdL"
      },
      {
        "symbol": "BTCUSDT",
        "orderId": 21,
        "clientOrderId": "461cPg51vQjV3zIMOXNz39"
      }
    ],
    "orderReports": [
      {
        "symbol": "BTCUSDT",
        "origClientOrderId": "CwOOIPHSmYywx6jZX77TdL",
        "orderId": 20,
        "orderListId": 1929,
        "clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
        "transactTime": 1688005070874,
        "price": "0.668611",
        "origQty": "0.690354",
        "executedQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "CANCELED",
        "timeInForce": "GTC",
        "type": "STOP_LOSS_LIMIT",
        "side": "BUY",
        "stopPrice": "0.378131",
        "icebergQty": "0.017083",
        "selfTradePreventionMode": "NONE"
      },
      {
        "symbol": "BTCUSDT",
        "origClientOrderId": "461cPg51vQjV3zIMOXNz39",
        "orderId": 21,
        "orderListId": 1929,
        "clientOrderId": "pXLV6Hz6mprAcVYpVMTGgx",
        "transactTime": 1688005070874,
        "price": "0.008791",
        "origQty": "0.690354",
        "executedQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "CANCELED",
        "timeInForce": "GTC",
        "type": "LIMIT_MAKER",
        "side": "BUY",
        "icebergQty": "0.639962",
        "selfTradePreventionMode": "NONE"
      }
    ]
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOpenOrder().Symbol("BTCUSDT").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*CancelOpenOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCancelOpenOrder(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestCancelOpenOrder(r1, r2 []*CancelOpenOrderResponse) {
	for i := range r1 {
		s.assertTestOrderResult(r1[i], r2[i])
	}
}

func (s *spotTradeTestSuite) assertTestOrderResult(r1, r2 *CancelOpenOrderResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrigClientOrderId, r2.OrigClientOrderId, "origClientOrderId")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
	r.Equal(r1.TransactTime, r2.TransactTime, "transactTime")
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.OrigQty, r2.OrigQty, "origQty")
	r.Equal(r1.ExecutedQty, r2.ExecutedQty, "executedQty")
	r.Equal(r1.CummulativeQuoteQty, r2.CummulativeQuoteQty, "cummulativeQuoteQty")
	r.Equal(r1.Status, r2.Status, "status")
	r.Equal(r1.TimeInForce, r2.TimeInForce, "timeInForce")
	r.Equal(r1.Type, r2.Type, "type")
	r.Equal(r1.Side, r2.Side, "side")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "selfTradePreventionMode")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	for i := range r1.Orders {
		s.assertApiOrder(r1.Orders[i], r2.Orders[i])
	}
	for i := range r1.OrderReports {
		s.assertApiOrderReport(r1.OrderReports[i], r2.OrderReports[i])
	}
}

func (s *spotTradeTestSuite) TestNewCancelReplace() {
	msg := []byte(`{
  "cancelResult": "SUCCESS",
  "newOrderResult": "SUCCESS",
  "cancelResponse": {
    "symbol": "BTCUSDT",
    "origClientOrderId": "DnLo3vTAQcjha43lAZhZ0y",
    "orderId": 9,
    "orderListId": -1,
    "clientOrderId": "osxN3JXAtJvKvCqGeMWMVR",
    "transactTime": 1684804350068,
    "price": "0.01000000",
    "origQty": "0.000100",
    "executedQty": "0.00000000",
    "origQuoteOrderQty": "0.000000",
    "cummulativeQuoteQty": "0.00000000",
    "status": "CANCELED",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "SELL",
    "selfTradePreventionMode": "NONE"
  },
  "newOrderResponse": {
    "symbol": "BTCUSDT",
    "orderId": 10,
    "orderListId": -1,
    "clientOrderId": "wOceeeOzNORyLiQfw7jd8S",
    "transactTime": 1652928801803,
    "price": "0.02000000",
    "origQty": "0.040000",
    "executedQty": "0.00000000",
    "cummulativeQuoteQty": "0.00000000",
    "origQuoteOrderQty": "0.000000",
    "status": "NEW",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "BUY",
    "workingTime": 1669277163808,
    "fills": [],
    "selfTradePreventionMode": "NONE"
  }
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelReplace().Symbol("BTCUSDT").
		Side("BUY").
		Type("MARKET").
		CancelOrderId(15379221).
		CancelReplaceMode("ALLOW_FAILURE").
		Quantity("0.001").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CancelReplaceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCancelReplaceOrder(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestCancelReplaceOrder(r1, r2 *CancelReplaceResponse) {
	r := s.r()
	if r1.Code != 0 {
		r.Equal(r1.Code, r2.Code, "code")
		r.Equal(r1.Msg, r2.Msg, "msg")
		r.Equal(r1.Data.CancelResult, r2.Data.CancelResult, "result.cancelResult")
		r.Equal(r1.Data.NewOrderResult, r2.Data.NewOrderResult, "result.newOrderResult")
		s.assertApiOrderReport(r1.Data.CancelResponse, r2.Data.CancelResponse)
		s.assertApiOrderReport(r1.Data.NewOrderResponse, r2.Data.NewOrderResponse)
	}
	if r1.CancelResult != "" {
		r.Equal(r1.CancelResult, r2.CancelResult, "result.cancelResult")
		r.Equal(r1.NewOrderResult, r2.NewOrderResult, "result.newOrderResult")
		s.assertApiOrderReport(r1.CancelResponse, r2.CancelResponse)
		s.assertApiOrderReport(r1.NewOrderResponse, r2.NewOrderResponse)
	}
}

func (s *spotTradeTestSuite) TestNewAllOpenOrders() {
	msg := []byte(`[
  {
    "symbol": "LTCBTC",
    "orderId": 1,
    "orderListId": -1,
    "clientOrderId": "myOrder1",
    "price": "0.1",
    "origQty": "1.0",
    "executedQty": "0.0",
    "cummulativeQuoteQty": "0.0",
    "status": "NEW",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "BUY",
    "stopPrice": "0.0",
    "icebergQty": "0.0",
    "time": 1499827319559,
    "updateTime": 1499827319559,
    "isWorking": true,
    "origQuoteOrderQty": "0.000000",
    "workingTime": 1499827319559,
    "selfTradePreventionMode": "NONE"
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOpenOrders().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrdersResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestAllOpenOrders(resp[i], testResp[i])
	}
}

func (s *spotTradeTestSuite) assertTestAllOpenOrders(r1, r2 *OrdersResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.OrigQty, r2.OrigQty, "origQty")
	r.Equal(r1.ExecutedQty, r2.ExecutedQty, "executedQty")
	r.Equal(r1.OrigQuoteOrderQty, r2.OrigQuoteOrderQty, "origQuoteOrderQty")
	r.Equal(r1.CummulativeQuoteQty, r2.CummulativeQuoteQty, "cummulativeQuoteQty")
	r.Equal(r1.Status, r2.Status, "status")
	r.Equal(r1.TimeInForce, r2.TimeInForce, "timeInForce")
	r.Equal(r1.Type, r2.Type, "type")
	r.Equal(r1.Side, r2.Side, "side")
	r.Equal(r1.StopPrice, r2.StopPrice, "stopPrice")
	r.Equal(r1.IcebergQty, r2.IcebergQty, "icebergQty")
	r.Equal(r1.Time, r2.Time, "time")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "updateTime")
	r.Equal(r1.IsWorking, r2.IsWorking, "isWorking")
	r.Equal(r1.WorkingTime, r2.WorkingTime, "workingTime")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "selfTradePreventionMode")
}

func (s *spotTradeTestSuite) TestNewAllOrders() {
	msg := []byte(`[
  {
    "symbol": "LTCBTC",
    "orderId": 1,
    "orderListId": -1,
    "clientOrderId": "myOrder1",
    "price": "0.1",
    "origQty": "1.0",
    "executedQty": "0.0",
    "cummulativeQuoteQty": "0.0",
    "status": "NEW",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "BUY",
    "stopPrice": "0.0",
    "icebergQty": "0.0",
    "time": 1499827319559,
    "updateTime": 1499827319559,
    "isWorking": true,
    "origQuoteOrderQty": "0.000000",
    "workingTime": 1499827319559,
    "selfTradePreventionMode": "NONE"
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAllOrders().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrdersResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestAllOpenOrders(resp[i], testResp[i])
	}
}

func (s *spotTradeTestSuite) TestNewCancelOrderList() {
	msg := []byte(`{
  "orderListId": 0,
  "contingencyType": "OCO",
  "listStatusType": "ALL_DONE",
  "listOrderStatus": "ALL_DONE",
  "listClientOrderId": "C3wyj4WVEktd7u9aVBRXcN",
  "transactionTime": 1574040868128,
  "symbol": "LTCBTC",
  "orders": [
    {
      "symbol": "LTCBTC",
      "orderId": 2,
      "clientOrderId": "pO9ufTiFGg3nw2fOdgeOXa"
    },
    {
      "symbol": "LTCBTC",
      "orderId": 3,
      "clientOrderId": "TXOvglzXuaubXAaENpaRCB"
    }
  ],
  "orderReports": [
    {
      "symbol": "LTCBTC",
      "origClientOrderId": "pO9ufTiFGg3nw2fOdgeOXa",
      "orderId": 2,
      "orderListId": 0,
      "clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
      "transactTime": 1688005070874,
      "price": "1.00000000",
      "origQty": "10.00000000",
      "executedQty": "0.00000000",
      "origQuoteOrderQty": "0.000000",
      "cummulativeQuoteQty": "0.00000000",
      "status": "CANCELED",
      "timeInForce": "GTC",
      "type": "STOP_LOSS_LIMIT",
      "side": "SELL",
      "stopPrice": "1.00000000",
      "selfTradePreventionMode": "NONE"
    },
    {
      "symbol": "LTCBTC",
      "origClientOrderId": "TXOvglzXuaubXAaENpaRCB",
      "orderId": 3,
      "orderListId": 0,
      "clientOrderId": "unfWT8ig8i0uj6lPuYLez6",
      "transactTime": 1688005070874,
      "price": "3.00000000",
      "origQty": "10.00000000",
      "executedQty": "0.00000000",
      "origQuoteOrderQty": "0.000000",
      "cummulativeQuoteQty": "0.00000000",
      "status": "CANCELED",
      "timeInForce": "GTC",
      "type": "LIMIT_MAKER",
      "side": "SELL",
      "selfTradePreventionMode": "NONE"
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOrderList().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderListResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderListResponse(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestOrderListResponse(r1, r2 *OrderListResponse) {
	r := s.r()
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	for i := range r1.Orders {
		s.assertApiOrder(r1.Orders[i], r2.Orders[i])
	}
	for i := range r1.OrderReports {
		s.assertApiOrderReport(r1.OrderReports[i], r2.OrderReports[i])
	}
}

func (s *spotTradeTestSuite) TestNewQueryOrderList() {
	msg := []byte(`{
  "orderListId": 27,
  "contingencyType": "OCO",
  "listStatusType": "EXEC_STARTED",
  "listOrderStatus": "EXECUTING",
  "listClientOrderId": "h2USkA5YQpaXHPIrkd96xE",
  "transactionTime": 1565245656253,
  "symbol": "LTCBTC",
  "orders": [
    {
      "symbol": "LTCBTC",
      "orderId": 4,
      "clientOrderId": "qD1gy3kc3Gx0rihm9Y3xwS"
    },
    {
      "symbol": "LTCBTC",
      "orderId": 5,
      "clientOrderId": "ARzZ9I00CPM8i3NhmU9Ega"
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOrderList().OrigClientOrderId("TXOvglzXuaubXAaENpaRCB").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryOrderListResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestQueryOrderListResponse(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestQueryOrderListResponse(r1, r2 *QueryOrderListResponse) {
	r := s.r()
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	for i := range r1.Orders {
		s.assertApiOrder(r1.Orders[i], r2.Orders[i])
	}
}

func (s *spotTradeTestSuite) TestNewQueryAllOrderLists() {
	msg := []byte(`[
  {
    "orderListId": 29,
    "contingencyType": "OCO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "amEEAXryFzFwYF1FeRpUoZ",
    "transactionTime": 1565245913483,
    "symbol": "LTCBTC",
    "orders": [
      {
        "symbol": "LTCBTC",
        "orderId": 4,
        "clientOrderId": "oD7aesZqjEGlZrbtRpy5zB"
      },
      {
        "symbol": "LTCBTC",
        "orderId": 5,
        "clientOrderId": "Jr1h6xirOxgeJOUuYQS7V3"
      }
    ]
  },
  {
    "orderListId": 28,
    "contingencyType": "OCO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "hG7hFNxJV6cZy3Ze4AUT4d",
    "transactionTime": 1565245913407,
    "symbol": "LTCBTC",
    "orders": [
      {
        "symbol": "LTCBTC",
        "orderId": 2,
        "clientOrderId": "j6lFOfbmFMRjTYA7rRJ0LP"
      },
      {
        "symbol": "LTCBTC",
        "orderId": 3,
        "clientOrderId": "z0KCjOdditiLS5ekAFtK81"
      }
    ]
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryAllOrderLists().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryAllOrderListsResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestQueryAllOrderListsResponse(resp[i], testResp[i])
	}
}

func (s *spotTradeTestSuite) assertTestQueryAllOrderListsResponse(r1, r2 *QueryAllOrderListsResponse) {
	r := s.r()
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	for i := range r1.Orders {
		s.assertApiOrder(r1.Orders[i], r2.Orders[i])
	}
}

func (s *spotTradeTestSuite) TestNewQueryOpenOrderList() {
	msg := []byte(`[
  {
    "orderListId": 31,
    "contingencyType": "OCO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "wuB13fmulKj3YjdqWEcsnp",
    "transactionTime": 1565246080644,
    "symbol": "LTCBTC",
    "orders": [
      {
        "symbol": "LTCBTC",
        "orderId": 4,
        "clientOrderId": "r3EH2N76dHfLoSZWIUw1bT"
      },
      {
        "symbol": "LTCBTC",
        "orderId": 5,
        "clientOrderId": "Cv1SnyPD3qhqpbjpYEHbd2"
      }
    ]
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOpenOrderList().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryOpenOrderListResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestQueryOpenOrderListResponse(resp[i], testResp[i])
	}
}

func (s *spotTradeTestSuite) assertTestQueryOpenOrderListResponse(r1, r2 *QueryOpenOrderListResponse) {
	r := s.r()
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	for i := range r1.Orders {
		s.assertApiOrder(r1.Orders[i], r2.Orders[i])
	}
}

func (s *spotTradeTestSuite) TestNewCreateSOROrder() {
	msg := []byte(`{
  "symbol": "BTCUSDT",
  "orderId": 2,
  "orderListId": -1,
  "clientOrderId": "sBI1KM6nNtOfj5tccZSKly",
  "transactTime": 1689149087774,
  "price": "31000.00000000",
  "origQty": "0.50000000",
  "executedQty": "0.50000000",
  "origQuoteOrderQty": "0.000000",
  "cummulativeQuoteQty": "14000.00000000",
  "status": "FILLED",
  "timeInForce": "GTC",
  "type": "LIMIT",
  "side": "BUY",
  "workingTime": 1689149087774,
  "fills": [
    {
      "matchType": "ONE_PARTY_TRADE_REPORT",
      "price": "28000.00000000",
      "qty": "0.50000000",
      "commission": "0.00000000",
      "commissionAsset": "BTC",
      "tradeId": -1,
      "allocId": 0
    }
  ],
  "workingFloor": "SOR",
  "selfTradePreventionMode": "NONE",
  "usedSor": true
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateSOROrder().Symbol("BTCUSDT").
		Side("BUY").
		Type("MARKET").
		Quantity("0.001").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateSOROrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateSOROrderResponse(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestCreateSOROrderResponse(r1, r2 *CreateSOROrderResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.OrderId, r2.OrderId, "OrderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "OrderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "ClientOrderId")
	r.Equal(r1.TransactTime, r2.TransactTime, "TransactTime")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.OrigQty, r2.OrigQty, "OrigQty")
	r.Equal(r1.ExecutedQty, r2.ExecutedQty, "ExecutedQty")
	r.Equal(r1.OrigQuoteOrderQty, r2.OrigQuoteOrderQty, "OrigQuoteOrderQty")
	r.Equal(r1.CummulativeQuoteQty, r2.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(r1.Status, r2.Status, "Status")
	r.Equal(r1.TimeInForce, r2.TimeInForce, "TimeInForce")
	r.Equal(r1.Type, r2.Type, "Type")
	r.Equal(r1.Side, r2.Side, "Side")
	r.Equal(r1.WorkingTime, r2.WorkingTime, "WorkingTime")
	for i := range r1.Fills {
		s.assertApiFill(r1.Fills[i], r2.Fills[i])
	}
	r.Equal(r1.WorkingFloor, r2.WorkingFloor, "WorkingFloor")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r1.UsedSor, r2.UsedSor, "UsedSor")
}

func (s *spotTradeTestSuite) TestNewCreateTestSOROrder() {
	msg := []byte(`{
  "standardCommissionForOrder": {
    "maker": "0.00000112",
    "taker": "0.00000114"
  },
  "taxCommissionForOrder": {
    "maker": "0.00000112",
    "taker": "0.00000114"
  },
  "discount": {
    "enabledForAccount": true,
    "enabledForSymbol": true,
    "discountAsset": "BNB",
    "discount": "0.25000000"
  }
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateTestSOROrder().Symbol("BTCUSDT").
		Side("BUY").
		Type("MARKET").
		Quantity("0.001").
		ComputeCommissionRates(true).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *TestCreateOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateTestSOROrderResponse(resp, testResp)
}

func (s *spotTradeTestSuite) assertTestCreateTestSOROrderResponse(r1, r2 *TestCreateOrderResponse) {
	r := s.r()
	r.Equal(r1.StandardCommissionForOrder.Maker, r2.StandardCommissionForOrder.Maker, "StandardCommissionForOrder.maker")
	r.Equal(r1.StandardCommissionForOrder.Taker, r2.StandardCommissionForOrder.Taker, "StandardCommissionForOrder.taker")
	r.Equal(r1.TaxCommissionForOrder.Maker, r2.TaxCommissionForOrder.Maker, "TaxCommissionForOrder.maker")
	r.Equal(r1.TaxCommissionForOrder.Taker, r2.TaxCommissionForOrder.Taker, "TaxCommissionForOrder.taker")
	r.Equal(r1.Discount.EnabledForAccount, r2.Discount.EnabledForAccount, "Discount.enabledForAccount")
	r.Equal(r1.Discount.EnabledForSymbol, r2.Discount.EnabledForSymbol, "Discount.enabledForSymbol")
	r.Equal(r1.Discount.DiscountAsset, r2.Discount.DiscountAsset, "Discount.discountAsset")
	r.Equal(r1.Discount.Discount, r2.Discount.Discount, "Discount.discount")
}
