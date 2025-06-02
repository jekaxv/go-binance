package wss

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

func TestWebsocketApiTrade(t *testing.T) {
	suite.Run(t, new(apiTradeTestSuite))
}

func (s *apiTradeTestSuite) TestCreateOrder() {
	msg := []byte(`{
  "id": "56374a46-3061-486b-a311-99ee972eb648",
  "status": 200,
  "result": {
    "symbol": "BTCUSDT",
    "orderId": 12569099453,
    "orderListId": -1, 
    "clientOrderId": "4d96324ff9d44481926157ec08158a40",
    "transactTime": 1660801715639,
    "price": "23416.10000000",
    "origQty": "0.00847000",
    "executedQty": "0.00000000",
    "origQuoteOrderQty": "0.000000",
    "cummulativeQuoteQty": "0.00000000",
    "status": "NEW",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "SELL",
    "workingTime": 1660801715639,
    "selfTradePreventionMode": "NONE"
  },
  "rateLimits": [
    {
      "rateLimitType": "ORDERS",
      "interval": "SECOND",
      "intervalNum": 10,
      "limit": 50,
      "count": 1
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "DAY",
      "intervalNum": 1,
      "limit": 160000,
      "count": 1
    },
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).Type(types.OrderTypeMARKET).Quantity("0.001").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateOrder(resp, testResp)
}

func (s *apiTradeTestSuite) assertTestCreateOrder(r1, r2 *CreateOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.Symbol, r2.Result.Symbol, "Symbol")
	r.Equal(r1.Result.OrderId, r2.Result.OrderId, "OrderId")
	r.Equal(r1.Result.OrderListId, r2.Result.OrderListId, "OrderListId")
	r.Equal(r1.Result.ClientOrderId, r2.Result.ClientOrderId, "ClientOrderId")
	r.Equal(r1.Result.TransactTime, r2.Result.TransactTime, "TransactTime")
	r.Equal(r1.Result.Price, r2.Result.Price, "Price")
	r.Equal(r1.Result.OrigQty, r2.Result.OrigQty, "OrigQty")
	r.Equal(r1.Result.ExecutedQty, r2.Result.ExecutedQty, "ExecutedQty")
	r.Equal(r1.Result.OrigQuoteOrderQty, r2.Result.OrigQuoteOrderQty, "OrigQuoteOrderQty")
	r.Equal(r1.Result.CummulativeQuoteQty, r2.Result.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(r1.Result.Status, r2.Result.Status, "Status")
	r.Equal(r1.Result.TimeInForce, r2.Result.TimeInForce, "TimeInForce")
	r.Equal(r1.Result.Type, r2.Result.Type, "Type")
	r.Equal(r1.Result.Side, r2.Result.Side, "Side")
	r.Equal(r1.Result.WorkingTime, r2.Result.WorkingTime, "WorkingTime")
	for i := range r1.Result.Fills {
		s.assertWsApiFill(r1.Result.Fills[i], r2.Result.Fills[i])
	}
}

func (s *apiTradeTestSuite) TestCreateTestOrder() {
	msg := []byte(`{
  "id": "2794d94d-47f9-4ad9-8a33-5d662487477f",
  "status": 200,
  "result": {
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
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 22
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateTestOrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).Type(types.OrderTypeMARKET).Quantity("0.001").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateOrderTestResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateTestOrder(resp, testResp)
}
func (s *apiTradeTestSuite) assertTestCreateTestOrder(r1, r2 *CreateOrderTestResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.StandardCommissionForOrder.Maker, r2.Result.StandardCommissionForOrder.Maker, "StandardCommissionForOrder.maker")
	r.Equal(r1.Result.StandardCommissionForOrder.Taker, r2.Result.StandardCommissionForOrder.Taker, "StandardCommissionForOrder.taker")
	r.Equal(r1.Result.TaxCommissionForOrder.Maker, r2.Result.TaxCommissionForOrder.Maker, "TaxCommissionForOrder.maker")
	r.Equal(r1.Result.TaxCommissionForOrder.Taker, r2.Result.TaxCommissionForOrder.Taker, "TaxCommissionForOrder.taker")
	r.Equal(r1.Result.Discount.EnabledForAccount, r2.Result.Discount.EnabledForAccount, "Discount.enabledForAccount")
	r.Equal(r1.Result.Discount.EnabledForSymbol, r2.Result.Discount.EnabledForSymbol, "Discount.enabledForSymbol")
	r.Equal(r1.Result.Discount.DiscountAsset, r2.Result.Discount.DiscountAsset, "Discount.discountAsset")
	r.Equal(r1.Result.Discount.Discount, r2.Result.Discount.Discount, "Discount.discount")
}

func (s *apiTradeTestSuite) TestQueryOrder() {
	msg := []byte(`{
  "id": "aa62318a-5a97-4f3b-bdc7-640bbe33b291",
  "status": 200,
  "result": {
    "symbol": "BTCUSDT",
    "orderId": 12569099453,
    "orderListId": -1,
    "clientOrderId": "4d96324ff9d44481926157",
    "price": "23416.10000000",
    "origQty": "0.00847000",
    "executedQty": "0.00847000",
    "cummulativeQuoteQty": "198.33521500",
    "status": "FILLED",
    "timeInForce": "GTC",
    "type": "LIMIT",
    "side": "SELL",
    "stopPrice": "0.00000000",
    "trailingDelta": 10,
    "trailingTime": -1,
    "icebergQty": "0.00000000",
    "time": 1660801715639,
    "updateTime": 1660801717945,
    "isWorking": true,
    "workingTime": 1660801715639,
    "origQuoteOrderQty": "0.00000000",
    "strategyId": 37463720,
    "strategyType": 1000000,
    "selfTradePreventionMode": "NONE",
    "preventedMatchId": 0,
    "preventedQuantity": "1.200000"
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 4
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOrder().Symbol("BTCUSDT").
		OrderId(1).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestQueryOrder(resp, testResp)
}

func (s *apiTradeTestSuite) assertTestQueryOrder(r1, r2 *QueryOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.Symbol, r2.Result.Symbol, "Symbol")
	r.Equal(r1.Result.OrderId, r2.Result.OrderId, "OrderId")
	r.Equal(r1.Result.OrderListId, r2.Result.OrderListId, "OrderListId")
	r.Equal(r1.Result.ClientOrderId, r2.Result.ClientOrderId, "ClientOrderId")
	r.Equal(r1.Result.Price, r2.Result.Price, "Price")
	r.Equal(r1.Result.OrigQty, r2.Result.OrigQty, "OrigQty")
	r.Equal(r1.Result.ExecutedQty, r2.Result.ExecutedQty, "ExecutedQty")
	r.Equal(r1.Result.CummulativeQuoteQty, r2.Result.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(r1.Result.Status, r2.Result.Status, "Status")
	r.Equal(r1.Result.TimeInForce, r2.Result.TimeInForce, "TimeInForce")
	r.Equal(r1.Result.Type, r2.Result.Type, "Type")
	r.Equal(r1.Result.Side, r2.Result.Side, "Side")
	r.Equal(r1.Result.StopPrice, r2.Result.StopPrice, "StopPrice")
	r.Equal(r1.Result.TrailingDelta, r2.Result.TrailingDelta, "TrailingDelta")
	r.Equal(r1.Result.TrailingTime, r2.Result.TrailingTime, "TrailingTime")
	r.Equal(r1.Result.IcebergQty, r2.Result.IcebergQty, "IcebergQty")
	r.Equal(r1.Result.Time, r2.Result.Time, "Time")
	r.Equal(r1.Result.UpdateTime, r2.Result.UpdateTime, "UpdateTime")
	r.Equal(r1.Result.IsWorking, r2.Result.IsWorking, "IsWorking")
	r.Equal(r1.Result.WorkingTime, r2.Result.WorkingTime, "WorkingTime")
	r.Equal(r1.Result.OrigQuoteOrderQty, r2.Result.OrigQuoteOrderQty, "OrigQuoteOrderQty")
	r.Equal(r1.Result.StrategyId, r2.Result.StrategyId, "StrategyId")
	r.Equal(r1.Result.StrategyType, r2.Result.StrategyType, "StrategyType")
	r.Equal(r1.Result.SelfTradePreventionMode, r2.Result.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r1.Result.PreventedMatchId, r2.Result.PreventedMatchId, "PreventedMatchId")
	r.Equal(r1.Result.PreventedQuantity, r2.Result.PreventedQuantity, "PreventedQuantity")
}

func (s *apiTradeTestSuite) TestCancelOrder() {
	msg := []byte(`{
  "id": "16eaf097-bbec-44b9-96ff-e97e6e875870",
  "status": 200,
  "result": {
    "orderListId": 19431,
    "contingencyType": "OCO",
    "listStatusType": "ALL_DONE",
    "listOrderStatus": "ALL_DONE",
    "listClientOrderId": "iuVNVJYYrByz6C4yGOPPK0",
    "transactionTime": 1660803702431,
    "symbol": "BTCUSDT",
    "orders": [
      {
        "symbol": "BTCUSDT",
        "orderId": 12569099454,
        "clientOrderId": "Tnu2IP0J5Y4mxw3IATBfmW"
      }
    ],
    "orderReports": [
      {
        "symbol": "BTCUSDT",
        "origClientOrderId": "Tnu2IP0J5Y4mxw3IATBfmW",
        "orderId": 12569099454,
        "orderListId": 19431,
        "clientOrderId": "OFFXQtxVFZ6Nbcg4PgE2DA",
        "transactTime": 1684804350068,
        "price": "23400.00000000",
        "origQty": "0.00850000",
        "executedQty": "0.00000000",
        "cummulativeQuoteQty": "0.00000000",
        "status": "CANCELED",
        "timeInForce": "GTC",
        "type": "LIMIT_MAKER",
        "side": "BUY",
        "selfTradePreventionMode": "NONE"
      }
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOrder().Symbol("BTCUSDT").
		OrderId(1).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CancelOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCancelOrder(resp, testResp)
}

func (s *apiTradeTestSuite) assertTestCancelOrder(r1, r2 *CancelOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.Symbol, r2.Result.Symbol, "Symbol")
	r.Equal(r1.Result.OrigClientOrderId, r2.Result.OrigClientOrderId, "OrigClientOrderId")
	r.Equal(r1.Result.OrderId, r2.Result.OrderId, "OrderId")
	r.Equal(r1.Result.OrderListId, r2.Result.OrderListId, "OrderListId")
	r.Equal(r1.Result.ClientOrderId, r2.Result.ClientOrderId, "ClientOrderId")
	r.Equal(r1.Result.TransactTime, r2.Result.TransactTime, "TransactTime")
	r.Equal(r1.Result.Price, r2.Result.Price, "Price")
	r.Equal(r1.Result.OrigQty, r2.Result.OrigQty, "OrigQty")
	r.Equal(r1.Result.ExecutedQty, r2.Result.ExecutedQty, "ExecutedQty")
	r.Equal(r1.Result.OrigQuoteOrderQty, r2.Result.OrigQuoteOrderQty, "OrigQuoteOrderQty")
	r.Equal(r1.Result.CummulativeQuoteQty, r2.Result.CummulativeQuoteQty, "CummulativeQuoteQty")
	r.Equal(r1.Result.Status, r2.Result.Status, "Status")
	r.Equal(r1.Result.TimeInForce, r2.Result.TimeInForce, "TimeInForce")
	r.Equal(r1.Result.Type, r2.Result.Type, "Type")
	r.Equal(r1.Result.Side, r2.Result.Side, "Side")
	r.Equal(r1.Result.StopPrice, r2.Result.StopPrice, "StopPrice")
	r.Equal(r1.Result.TrailingDelta, r2.Result.TrailingDelta, "TrailingDelta")
	r.Equal(r1.Result.IcebergQty, r2.Result.IcebergQty, "IcebergQty")
	r.Equal(r1.Result.StrategyId, r2.Result.StrategyId, "StrategyId")
	r.Equal(r1.Result.StrategyType, r2.Result.StrategyType, "StrategyType")
	r.Equal(r1.Result.SelfTradePreventionMode, r2.Result.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r1.Result.ContingencyType, r2.Result.ContingencyType, "ContingencyType")
	r.Equal(r1.Result.ListStatusType, r2.Result.ListStatusType, "ListStatusType")
	r.Equal(r1.Result.ListOrderStatus, r2.Result.ListOrderStatus, "ListOrderStatus")
	r.Equal(r1.Result.ListClientOrderId, r2.Result.ListClientOrderId, "ListClientOrderId")
	r.Equal(r1.Result.TransactionTime, r2.Result.TransactionTime, "TransactionTime")
	for i := range r1.Result.Orders {
		s.assertWsApiOrder(r1.Result.Orders[i], r2.Result.Orders[i])
	}
	for i := range r1.Result.OrderReports {
		s.assertWsApiOrderReport(r1.Result.OrderReports[i], r2.Result.OrderReports[i])
	}
}

func (s *apiTradeTestSuite) TestCancelReplaceOrder() {
	msg := []byte(`{
  "id": "99de1036-b5e2-4e0f-9b5c-13d751c93a1a",
  "status": 200,
  "result": {
    "cancelResult": "SUCCESS",
    "newOrderResult": "SUCCESS",
    "cancelResponse": {
      "symbol": "BTCUSDT",
      "origClientOrderId": "4d96324ff9d44481926157",
      "orderId": 125690984230,
      "orderListId": -1,
      "clientOrderId": "91fe37ce9e69c90d6358c0",
      "transactTime": 1684804350068,
      "price": "23450.00000000",
      "origQty": "0.00847000",
      "executedQty": "0.00001000",
      "origQuoteOrderQty": "0.000000",
      "cummulativeQuoteQty": "0.23450000",
      "status": "CANCELED",
      "timeInForce": "GTC",
      "type": "LIMIT",
      "side": "SELL",
      "selfTradePreventionMode": "NONE"
    },
    "newOrderResponse": {
      "symbol": "BTCUSDT",
      "orderId": 12569099453,
      "orderListId": -1,
      "clientOrderId": "bX5wROblo6YeDwa9iTLeyY",
      "transactTime": 1660813156959,
      "price": "23416.10000000",
      "origQty": "0.00847000",
      "executedQty": "0.00000000",
      "origQuoteOrderQty": "0.000000",
      "cummulativeQuoteQty": "0.00000000",
      "status": "NEW",
      "timeInForce": "GTC",
      "type": "LIMIT",
      "side": "SELL",
      "selfTradePreventionMode": "NONE"
    }
  },
  "rateLimits": [
    {
      "rateLimitType": "ORDERS",
      "interval": "SECOND",
      "intervalNum": 10,
      "limit": 50,
      "count": 1
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "DAY",
      "intervalNum": 1,
      "limit": 160000,
      "count": 1
    },
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelReplaceOrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).
		CancelOrderId(1).
		CancelReplaceMode(types.ReplaceModeSTOP_ON_FAILURE).
		Type(types.OrderTypeLIMIT).
		Quantity("0.0001").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CancelReplaceOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCancelReplaceOrder(resp, testResp)
}
func (s *apiTradeTestSuite) assertTestCancelReplaceOrder(r1, r2 *CancelReplaceOrderResponse) {
	r := s.r()
	r.Equal(r1.Id, r2.Id, "Id")
	r.Equal(r1.Status, r2.Status, "Status")
	for i := range r1.RateLimits {
		s.assertWsRateLimits(r1.RateLimits[i], r2.RateLimits[i])
	}
	if r1.Error != nil {
		r.Equal(r1.Error.ApiError.Code, r2.Error.ApiError.Code, "error.apiError.code")
		r.Equal(r1.Error.ApiError.Msg, r2.Error.ApiError.Msg, "error.apiError.msg")
		r.Equal(r1.Error.Data.CancelResponse, r2.Error.Data.CancelResult, "error.data.cancelResponse")
		r.Equal(r1.Error.Data.NewOrderResult, r2.Error.Data.NewOrderResult, "error.data.newOrderResult")
		r.Equal(r1.Error.Data.CancelResponse.Code, r2.Error.Data.CancelResponse.Code, "error.data.cancelResponse.code")
		r.Equal(r1.Error.Data.CancelResponse.Msg, r2.Error.Data.CancelResponse.Msg, "error.data.cancelResponse.msg")
		s.assertWsApiOrderReport(r1.Error.Data.NewOrderResponse, r2.Error.Data.NewOrderResponse)
	}
	if r1.Result != nil {
		r.Equal(r1.Result.CancelResult, r2.Result.CancelResult, "result.cancelResult")
		r.Equal(r1.Result.NewOrderResult, r2.Result.NewOrderResult, "result.newOrderResult")
		s.assertWsApiOrderReport(r1.Result.CancelResponse, r2.Result.CancelResponse)
		s.assertWsApiOrderReport(r1.Result.NewOrderResponse, r2.Result.NewOrderResponse)
	}
}

func (s *apiTradeTestSuite) TestOpenOrdersStatus() {
	msg := []byte(`{
  "id": "55f07876-4f6f-4c47-87dc-43e5fff3f2e7",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "orderId": 12569099453,
      "orderListId": -1,
      "clientOrderId": "4d96324ff9d44481926157",
      "price": "23416.10000000",
      "origQty": "0.00847000",
      "executedQty": "0.00720000",
      "origQuoteOrderQty": "0.000000",
      "cummulativeQuoteQty": "172.43931000",
      "status": "PARTIALLY_FILLED",
      "timeInForce": "GTC",
      "type": "LIMIT",
      "side": "SELL",
      "stopPrice": "0.00000000",
      "icebergQty": "0.00000000",
      "time": 1660801715639,
      "updateTime": 1660801717945,
      "isWorking": true,
      "workingTime": 1660801715639,
      "origQuoteOrderQty": "0.00000000",
      "selfTradePreventionMode": "NONE"
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 6
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOpenOrdersStatus().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OpenOrdersStatusResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOpenOrdersStatus(resp, testResp)
}

func (s *apiTradeTestSuite) assertTestOpenOrdersStatus(r1, r2 *OpenOrdersStatusResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestOpenOrdersStatusResult(r1.Result[i], r2.Result[i])
	}
}
func (s *apiTradeTestSuite) assertTestOpenOrdersStatusResult(r1, r2 *OpenOrdersStatusResult) {
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

func (s *apiTradeTestSuite) TestCancelOpenOrder() {
	msg := []byte(`{
  "id": "778f938f-9041-4b88-9914-efbf64eeacc8",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "origClientOrderId": "4d96324ff9d44481926157",
      "orderId": 12569099453,
      "orderListId": -1,
      "clientOrderId": "91fe37ce9e69c90d6358c0",
      "transactTime": 1684804350068,
      "price": "23416.10000000",
      "origQty": "0.00847000",
      "executedQty": "0.00001000",
      "origQuoteOrderQty": "0.000000",
      "cummulativeQuoteQty": "0.23416100",
      "status": "CANCELED",
      "timeInForce": "GTC",
      "type": "LIMIT",
      "side": "SELL",
      "stopPrice": "0.00000000",
      "trailingDelta": 0,
      "trailingTime": -1,
      "icebergQty": "0.00000000",
      "strategyId": 37463720,
      "strategyType": 1000000,
      "selfTradePreventionMode": "NONE"
    },
    {
      "orderListId": 19431,
      "contingencyType": "OCO",
      "listStatusType": "ALL_DONE",
      "listOrderStatus": "ALL_DONE",
      "listClientOrderId": "iuVNVJYYrByz6C4yGOPPK0",
      "transactionTime": 1660803702431,
      "symbol": "BTCUSDT",
      "orders": [
        {
          "symbol": "BTCUSDT",
          "orderId": 12569099453,
          "clientOrderId": "bX5wROblo6YeDwa9iTLeyY"
        },
        {
          "symbol": "BTCUSDT",
          "orderId": 12569099454,
          "clientOrderId": "Tnu2IP0J5Y4mxw3IATBfmW"
        }
      ],
      "orderReports": [
        {
          "symbol": "BTCUSDT",
          "origClientOrderId": "bX5wROblo6YeDwa9iTLeyY",
          "orderId": 12569099453,
          "orderListId": 19431,
          "clientOrderId": "OFFXQtxVFZ6Nbcg4PgE2DA",
          "transactTime": 1684804350068,
          "price": "23450.50000000",
          "origQty": "0.00850000",
          "executedQty": "0.00000000",
          "origQuoteOrderQty": "0.000000",
          "cummulativeQuoteQty": "0.00000000",
          "status": "CANCELED",
          "timeInForce": "GTC",
          "type": "STOP_LOSS_LIMIT",
          "side": "BUY",
          "stopPrice": "23430.00000000",
          "selfTradePreventionMode": "NONE"
        },
        {
          "symbol": "BTCUSDT",
          "origClientOrderId": "Tnu2IP0J5Y4mxw3IATBfmW",
          "orderId": 12569099454,
          "orderListId": 19431,
          "clientOrderId": "OFFXQtxVFZ6Nbcg4PgE2DA",
          "transactTime": 1684804350068,
          "price": "23400.00000000",
          "origQty": "0.00850000",
          "executedQty": "0.00000000",
          "origQuoteOrderQty": "0.000000",
          "cummulativeQuoteQty": "0.00000000",
          "status": "CANCELED",
          "timeInForce": "GTC",
          "type": "LIMIT_MAKER",
          "side": "BUY",
          "selfTradePreventionMode": "NONE"
        }
      ]
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOpenOrder().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CancelOpenOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCancelOpenOrder(resp, testResp)
}

func (s *apiTradeTestSuite) assertTestCancelOpenOrder(r1, r2 *CancelOpenOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestOrderResult(r1.Result[i], r2.Result[i])
	}
}
func (s *apiTradeTestSuite) assertTestOrderResult(r1, r2 *OrderResult) {
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
	r.Equal(r1.OrigQuoteOrderQty, r2.OrigQuoteOrderQty, "origQuoteOrderQty")
	r.Equal(r1.CummulativeQuoteQty, r2.CummulativeQuoteQty, "cummulativeQuoteQty")
	r.Equal(r1.Status, r2.Status, "status")
	r.Equal(r1.TimeInForce, r2.TimeInForce, "timeInForce")
	r.Equal(r1.Type, r2.Type, "type")
	r.Equal(r1.Side, r2.Side, "side")
	r.Equal(r1.StopPrice, r2.StopPrice, "stopPrice")
	r.Equal(r1.TrailingDelta, r2.TrailingDelta, "trailingDelta")
	r.Equal(r1.TrailingTime, r2.TrailingTime, "trailingTime")
	r.Equal(r1.IcebergQty, r2.IcebergQty, "icebergQty")
	r.Equal(r1.StrategyId, r2.StrategyId, "strategyId")
	r.Equal(r1.StrategyType, r2.StrategyType, "strategyType")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "selfTradePreventionMode")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	for i := range r1.Orders {
		s.assertWsApiOrder(r1.Orders[i], r2.Orders[i])
	}
	for i := range r1.OrderReports {
		s.assertWsApiOrderReport(r1.OrderReports[i], r2.OrderReports[i])
	}
}

func (s *apiTradeTestSuite) TestCreateOCOOrder() {
	msg := []byte(`{
  "id": "56374a46-3261-486b-a211-99ed972eb648",
  "status": 200,
  "result":
  {
    "orderListId": 2,
    "contingencyType": "OCO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "cKPMnDCbcLQILtDYM4f4fX",
    "transactionTime": 1711062760648,
    "symbol": "LTCBNB",
    "orders":
    [
      {
        "symbol": "LTCBNB",
        "orderId": 2,
        "clientOrderId": "0m6I4wfxvTUrOBSMUl0OPU"
      },
      {
        "symbol": "LTCBNB",
        "orderId": 3,
        "clientOrderId": "Z2IMlR79XNY5LU0tOxrWyW"
      }
    ],
    "orderReports":
    [
      {
        "symbol": "LTCBNB",
        "orderId": 2,
        "orderListId": 2,
        "clientOrderId": "0m6I4wfxvTUrOBSMUl0OPU",
        "transactTime": 1711062760648,
        "price": "1.50000000",
        "origQty": "1.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.00000000",
        "status": "NEW",
        "timeInForce": "GTC",
        "type": "STOP_LOSS_LIMIT",
        "side": "BUY",
        "stopPrice": "1.50000001",
        "workingTime": -1,
        "selfTradePreventionMode": "NONE"
      },
      {
        "symbol": "LTCBNB",
        "orderId": 3,
        "orderListId": 2,
        "clientOrderId": "Z2IMlR79XNY5LU0tOxrWyW",
        "transactTime": 1711062760648,
        "price": "1.49999999",
        "origQty": "1.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.00000000",
        "status": "NEW",
        "timeInForce": "GTC",
        "type": "LIMIT_MAKER",
        "side": "BUY",
        "workingTime": 1711062760648,
        "selfTradePreventionMode": "NONE"
      }
    ]
  },
  "rateLimits":
  [
    {
      "rateLimitType": "ORDERS",
      "interval": "SECOND",
      "intervalNum": 10,
      "limit": 50,
      "count": 2
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "DAY",
      "intervalNum": 1,
      "limit": 160000,
      "count": 2
    },
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOCOOrder().Symbol("LTCBNB").
		Side(types.OrderSideBUY).
		Quantity("1").
		AboveType(types.OrderTypeSTOP_LOSS_LIMIT).
		AbovePrice("1.5").
		AboveStopPrice("1.50000001").
		AboveTimeInForce(types.TimeInForceGTC).
		BelowType(types.OrderTypeLIMIT_MAKER).
		BelowPrice("1.49999999").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrder(resp, testResp)
}

func (s *apiTradeTestSuite) assertTestOrder(r1, r2 *OrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertTestOrderResult(r1.Result, r2.Result)
}

func (s *apiTradeTestSuite) TestCreateOTOOrder() {
	msg := []byte(`{
  "id": "1712544395950",
  "status": 200,
  "result": {
    "orderListId": 626,
    "contingencyType": "OTO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "KA4EBjGnzvSwSCQsDdTrlf",
    "transactionTime": 1712544395981,
    "symbol": "1712544378871",
    "orders": [
      {
        "symbol": "LTCBNB",
        "orderId": 13,
        "clientOrderId": "YiAUtM9yJjl1a2jXHSp9Ny"
      },
      {
        "symbol": "LTCBNB",
        "orderId": 14,
        "clientOrderId": "9MxJSE1TYkmyx5lbGLve7R"
      }
    ],
    "orderReports": [
      {
        "symbol": "LTCBNB",
        "orderId": 13,
        "orderListId": 626,
        "clientOrderId": "YiAUtM9yJjl1a2jXHSp9Ny",
        "transactTime": 1712544395981,
        "price": "1.000000",
        "origQty": "1.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "NEW",
        "timeInForce": "GTC",
        "type": "LIMIT",
        "side": "SELL",
        "workingTime": 1712544395981,
        "selfTradePreventionMode": "NONE"
      },
      {
        "symbol": "LTCBNB",
        "orderId": 14,
        "orderListId": 626,
        "clientOrderId": "9MxJSE1TYkmyx5lbGLve7R",
        "transactTime": 1712544395981,
        "price": "0.000000",
        "origQty": "1.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "PENDING_NEW",
        "timeInForce": "GTC",
        "type": "MARKET",
        "side": "BUY",
        "workingTime": -1,
        "selfTradePreventionMode": "NONE"
      }
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "ORDERS",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 10000000,
      "count": 10
    },
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 1000,
      "count": 38
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOTOOrder().Symbol("LTCBNB").
		PendingSide(types.OrderSideBUY).
		PendingQuantity("1").
		PendingType(types.OrderTypeMARKET).
		WorkingPrice("1").
		WorkingQuantity("1").
		WorkingSide(types.OrderSideSELL).
		WorkingTimeInForce(types.TimeInForceGTC).
		WorkingType(types.OrderTypeLIMIT).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrder(resp, testResp)
}

func (s *apiTradeTestSuite) TestCreateOTOCOOrder() {
	msg := []byte(`{
  "id": "1712544408508",
  "status": 200,
  "result": {
    "orderListId": 629,
    "contingencyType": "OTO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "GaeJHjZPasPItFj4x7Mqm6",
    "transactionTime": 1712544408537,
    "symbol": "1712544378871",
    "orders": [
      {
        "symbol": "1712544378871",
        "orderId": 23,
        "clientOrderId": "OVQOpKwfmPCfaBTD0n7e7H"
      },
      {
        "symbol": "1712544378871",
        "orderId": 24,
        "clientOrderId": "YcCPKCDMQIjNvLtNswt82X"
      },
      {
        "symbol": "1712544378871",
        "orderId": 25,
        "clientOrderId": "ilpIoShcFZ1ZGgSASKxMPt"
      }
    ],
    "orderReports": [
      {
        "symbol": "LTCBNB",
        "orderId": 23,
        "orderListId": 629,
        "clientOrderId": "OVQOpKwfmPCfaBTD0n7e7H",
        "transactTime": 1712544408537,
        "price": "1.500000",
        "origQty": "1.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "NEW",
        "timeInForce": "GTC",
        "type": "LIMIT",
        "side": "BUY",
        "workingTime": 1712544408537,
        "selfTradePreventionMode": "NONE"
      },
      {
        "symbol": "LTCBNB",
        "orderId": 24,
        "orderListId": 629,
        "clientOrderId": "YcCPKCDMQIjNvLtNswt82X",
        "transactTime": 1712544408537,
        "price": "0.000000",
        "origQty": "5.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "PENDING_NEW",
        "timeInForce": "GTC",
        "type": "STOP_LOSS",
        "side": "SELL",
        "stopPrice": "0.500000",
        "workingTime": -1,
        "selfTradePreventionMode": "NONE"
      },
      {
        "symbol": "LTCBNB",
        "orderId": 25,
        "orderListId": 629,
        "clientOrderId": "ilpIoShcFZ1ZGgSASKxMPt",
        "transactTime": 1712544408537,
        "price": "5.000000",
        "origQty": "5.000000",
        "executedQty": "0.000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.000000",
        "status": "PENDING_NEW",
        "timeInForce": "GTC",
        "type": "LIMIT_MAKER",
        "side": "SELL",
        "workingTime": -1,
        "selfTradePreventionMode": "NONE"
      }
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "ORDERS",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 10000000,
      "count": 18
    },
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 1000,
      "count": 65
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateOTOCOOrder().Symbol("LTCBNB").
		PendingSide(types.OrderSideBUY).
		PendingQuantity("5").
		PendingSide(types.OrderSideSELL).
		PendingBelowPrice("5").
		PendingBelowType(types.OrderTypeLIMIT_MAKER).
		PendingAboveStopPrice("0.5").
		PendingAboveType(types.OrderTypeSTOP_LOSS).
		WorkingPrice("1.6").
		WorkingSide(types.OrderSideBUY).
		WorkingTimeInForce(types.TimeInForceGTC).
		WorkingType(types.OrderTypeLIMIT).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrder(resp, testResp)
}

func (s *apiTradeTestSuite) TestQueryOrderList() {
	msg := []byte(`{
  "id": "b53fd5ff-82c7-4a04-bd64-5f9dc42c2100",
  "status": 200,
  "result": {
    "orderListId": 1274512,
    "contingencyType": "OCO",
    "listStatusType": "EXEC_STARTED",
    "listOrderStatus": "EXECUTING",
    "listClientOrderId": "08985fedd9ea2cf6b28996",
    "transactionTime": 1660801713793,
    "symbol": "BTCUSDT",
    "orders": [
      {
        "symbol": "BTCUSDT",
        "orderId": 12569138901,
        "clientOrderId": "BqtFCj5odMoWtSqGk2X9tU"
      },
      {
        "symbol": "BTCUSDT",
        "orderId": 12569138902,
        "clientOrderId": "jLnZpj5enfMXTuhKB1d0us"
      }
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 4
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOrderList().OrigClientOrderId("1").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderListResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderList(resp, testResp)
}
func (s *apiTradeTestSuite) assertTestOrderList(r1, r2 *OrderListResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertTestOrderListResult(r1.Result, r2.Result)
}

func (s *apiTradeTestSuite) assertTestOrderListResult(r1, r2 *OrderListResult) {
	r := s.r()
	r.Equal(r1.OrderListId, r2.OrderListId, "OrderListId")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "ContingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "ListStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "ListOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "ListClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "TransactionTime")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	for i := range r1.Orders {
		s.assertWsApiOrder(r1.Orders[i], r2.Orders[i])
	}
	for i := range r1.OrderReports {
		s.assertWsApiOrderReport(r1.OrderReports[i], r2.OrderReports[i])
	}
}

func (s *apiTradeTestSuite) TestCancelOrderList() {
	msg := []byte(`{
  "id": "c5899911-d3f4-47ae-8835-97da553d27d0",
  "status": 200,
  "result": {
    "orderListId": 1274512,
    "contingencyType": "OCO",
    "listStatusType": "ALL_DONE",
    "listOrderStatus": "ALL_DONE",
    "listClientOrderId": "6023531d7edaad348f5aff",
    "transactionTime": 1660801720215,
    "symbol": "BTCUSDT",
    "orders": [
      {
        "symbol": "BTCUSDT",
        "orderId": 12569138901,
        "clientOrderId": "BqtFCj5odMoWtSqGk2X9tU"
      },
      {
        "symbol": "BTCUSDT",
        "orderId": 12569138902,
        "clientOrderId": "jLnZpj5enfMXTuhKB1d0us"
      }
    ],
    "orderReports": [
      {
        "symbol": "BTCUSDT",
        "orderId": 12569138901,
        "orderListId": 1274512,
        "clientOrderId": "BqtFCj5odMoWtSqGk2X9tU",
        "transactTime": 1660801720215,
        "price": "23410.00000000",
        "origQty": "0.00650000",
        "executedQty": "0.00000000",
        "origQuoteOrderQty": "0.000000",
        "cummulativeQuoteQty": "0.00000000",
        "status": "CANCELED",
        "timeInForce": "GTC",
        "type": "STOP_LOSS_LIMIT",
        "side": "SELL",
        "stopPrice": "23405.00000000",
        "selfTradePreventionMode": "NONE"
      },
      {
        "symbol": "BTCUSDT",
        "orderId": 12569138902,
        "orderListId": 1274512,
        "clientOrderId": "jLnZpj5enfMXTuhKB1d0us",
        "transactTime": 1660801720215,
        "price": "23420.00000000",
        "origQty": "0.00650000",
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
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCancelOrderList().OrderListId(1).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OrderListResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestOrderList(resp, testResp)
}

func (s *apiTradeTestSuite) TestQueryOpenOrder() {
	msg := []byte(`{
  "id": "3a4437e2-41a3-4c19-897c-9cadc5dce8b6",
  "status": 200,
  "result": [
    {
      "orderListId": 0,
      "contingencyType": "OCO",
      "listStatusType": "EXEC_STARTED",
      "listOrderStatus": "EXECUTING",
      "listClientOrderId": "08985fedd9ea2cf6b28996",
      "transactionTime": 1660801713793,
      "symbol": "BTCUSDT",
      "orders": [
        {
          "symbol": "BTCUSDT",
          "orderId": 4,
          "clientOrderId": "CUhLgTXnX5n2c0gWiLpV4d"
        },
        {
          "symbol": "BTCUSDT",
          "orderId": 5,
          "clientOrderId": "1ZqG7bBuYwaF4SU8CwnwHm"
        }
      ]
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 6
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryOpenOrder().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryOpenOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestQueryOpenOrder(resp, testResp)
}
func (s *apiTradeTestSuite) assertTestQueryOpenOrder(r1, r2 *QueryOpenOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestOrderListResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiTradeTestSuite) TestCreateSOROrder() {
	msg := []byte(`{
  "id": "3a4437e2-41a3-4c19-897c-9cadc5dce8b6",
  "status": 200,
  "result": [
    {
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
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 1
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateSOROrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).
		Type(types.OrderTypeLIMIT).
		Quantity("0.5").
		Price("100000").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateSOROrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateSOROrder(resp, testResp)
}
func (s *apiTradeTestSuite) assertTestCreateSOROrder(r1, r2 *CreateSOROrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestCreateSOROrderResult(r1.Result[i], r2.Result[i])
	}
}
func (s *apiTradeTestSuite) assertTestCreateSOROrderResult(r1, r2 *CreateSOROrderResult) {
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
		s.assertWsApiFill(r1.Fills[i], r2.Fills[i])
	}
	r.Equal(r1.WorkingFloor, r2.WorkingFloor, "WorkingFloor")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "SelfTradePreventionMode")
	r.Equal(r1.UsedSor, r2.UsedSor, "UsedSor")
}

func (s *apiTradeTestSuite) TestCreateTestSOROrder() {
	msg := []byte(`{
  "id": "3a4437e2-41a3-4c19-897c-9cadc5dce8b6",
  "status": 200,
  "result": {
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
      "discount": "0.25"                          
    }
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 20
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCreateTestSOROrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).
		Type(types.OrderTypeLIMIT).
		Quantity("0.5").
		Price("100000").ComputeCommissionRates(true).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CreateTestSOROrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestCreateTestSOROrder(resp, testResp)
}
func (s *apiTradeTestSuite) assertTestCreateTestSOROrder(r1, r2 *CreateTestSOROrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.StandardCommissionForOrder.Maker, r2.Result.StandardCommissionForOrder.Maker, "StandardCommissionForOrder.maker")
	r.Equal(r1.Result.StandardCommissionForOrder.Taker, r2.Result.StandardCommissionForOrder.Taker, "StandardCommissionForOrder.taker")
	r.Equal(r1.Result.TaxCommissionForOrder.Maker, r2.Result.TaxCommissionForOrder.Maker, "TaxCommissionForOrder.maker")
	r.Equal(r1.Result.TaxCommissionForOrder.Taker, r2.Result.TaxCommissionForOrder.Taker, "TaxCommissionForOrder.taker")
	r.Equal(r1.Result.Discount.EnabledForAccount, r2.Result.Discount.EnabledForAccount, "Discount.enabledForAccount")
	r.Equal(r1.Result.Discount.EnabledForSymbol, r2.Result.Discount.EnabledForSymbol, "Discount.enabledForSymbol")
	r.Equal(r1.Result.Discount.DiscountAsset, r2.Result.Discount.DiscountAsset, "Discount.discountAsset")
	r.Equal(r1.Result.Discount.Discount, r2.Result.Discount.Discount, "Discount.discount")
}
