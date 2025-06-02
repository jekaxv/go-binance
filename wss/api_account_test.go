package wss

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiAccountTestSuite struct {
	baseTestSuite
}

func TestWebsocketApiAccount(t *testing.T) {
	suite.Run(t, new(apiAccountTestSuite))
}

func (s *apiAccountTestSuite) TestAccountInformation() {
	msg := []byte(`{
  "id": "d2a6c85d-ffb6-4acb-b630-323623a7cf1c",
  "status": 200,
  "result": {
    "makerCommission": 0,
    "takerCommission": 0,
    "buyerCommission": 0,
    "sellerCommission": 0,
    "commissionRates": {
      "maker": "0.00000000",
      "taker": "0.00000000",
      "buyer": "0.00000000",
      "seller": "0.00000000"
    },
    "canTrade": true,
    "canWithdraw": true,
    "canDeposit": true,
    "brokered": false,
    "requireSelfTradePrevention": false,
    "preventSor": false,
    "updateTime": 1737291872242,
    "accountType": "SPOT",
    "balances": [
      {
        "asset": "ETH",
        "free": "2.00000000",
        "locked": "0.00000000"
      }
    ],
    "permissions": [
      "SPOT"
    ],
    "uid": 10000
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
	resp, err := s.client.NewAccountInformation().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountInformationResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountInfo(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAccountInfo(r1, r2 *AccountInformationResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertTestAccountInfoResult(r1.Result, r2.Result)
}

func (s *apiAccountTestSuite) assertTestAccountInfoResult(r1, r2 *AccountInformationResult) {
	r := s.r()
	r.Equal(r1.MakerCommission, r2.MakerCommission, "makerCommission")
	r.Equal(r1.TakerCommission, r2.TakerCommission, "takerCommission")
	r.Equal(r1.BuyerCommission, r2.BuyerCommission, "buyerCommission")
	r.Equal(r1.SellerCommission, r2.SellerCommission, "sellerCommission")
	r.Equal(r1.CanTrade, r2.CanTrade, "canTrade")
	r.Equal(r1.CanWithdraw, r2.CanWithdraw, "canWithdraw")
	r.Equal(r1.CanDeposit, r2.CanDeposit, "canDeposit")
	s.assertWsCommissionRate(r1.CommissionRates, r2.CommissionRates)
	r.Equal(r1.Brokered, r2.Brokered, "brokered")
	r.Equal(r1.RequireSelfTradePrevention, r2.RequireSelfTradePrevention, "requireSelfTradePrevention")
	r.Equal(r1.PreventSor, r2.PreventSor, "preventSor")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "updateTime")
	r.Equal(r1.AccountType, r2.AccountType, "accountType")
	for i := range r1.Balances {
		s.assertWsBalance(r1.Balances[i], r2.Balances[i])
	}
	for i := range r1.Permissions {
		r.Equal(r1.Permissions[i], r2.Permissions[i], "permissions")
	}
	r.Equal(r1.Uid, r2.Uid, "uid")
}

func (s *apiAccountTestSuite) TestUnfilledOrder() {
	msg := []byte(`{"id":"5ac77bb0-2e78-4c11-a06b-b94e437e48d0","status":200,"result":[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":0}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":42}]}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewUnfilledOrder().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *UnfilledOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestUnfilledOrder(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestUnfilledOrder(r1, r2 *UnfilledOrderResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestUnfilledOrderResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiAccountTestSuite) assertTestUnfilledOrderResult(r1, r2 *UnfilledOrderResult) {
	r := s.r()
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.IntervalNum, r2.IntervalNum, "intervalNum")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.Count, r2.Count, "count")
}

func (s *apiAccountTestSuite) TestAccountOrderHistory() {
	msg := []byte(`{
  "id": "8409ac0b-7a9f-4bb8-b529-0ce9c909ed1e",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "orderId": 2805547,
      "orderListId": -1,
      "clientOrderId": "Qoy8DO6sIlFmfMtuSfUSRa",
      "price": "0.00000000",
      "origQty": "0.00010000",
      "executedQty": "0.00010000",
      "cummulativeQuoteQty": "9.43234600",
      "status": "FILLED",
      "timeInForce": "GTC",
      "type": "MARKET",
      "side": "BUY",
      "stopPrice": "0.00000000",
      "icebergQty": "0.00000000",
      "time": 1737291872242,
      "updateTime": 1737291872242,
      "isWorking": true,
      "workingTime": 1737291872242,
      "origQuoteOrderQty": "0.00000000",
      "selfTradePreventionMode": "EXPIRE_MAKER"
    }
  ],
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
	resp, err := s.client.NewAccountOrderHistory().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountOrderHistoryResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountOrderHistory(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAccountOrderHistory(r1, r2 *AccountOrderHistoryResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestAccountOrderHistoryResult(r1.Result[i], r2.Result[i])
	}
}
func (s *apiAccountTestSuite) assertTestAccountOrderHistoryResult(r1, r2 *AccountOrderHistoryResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.OrigQty, r2.OrigQty, "origQty")
	r.Equal(r1.ExecutedQty, r2.ExecutedQty, "executedQty")
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
	r.Equal(r1.OrigQuoteOrderQty, r2.OrigQuoteOrderQty, "origQuoteOrderQty")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "selfTradePreventionMode")
	r.Equal(r1.PreventedMatchId, r2.PreventedMatchId, "preventedMatchId")
	r.Equal(r1.PreventedQuantity, r2.PreventedQuantity, "preventedQuantity")
}

func (s *apiAccountTestSuite) TestAllOrderList() {
	msg := []byte(`{
  "id": "8617b7b3-1b3d-4dec-94cd-eefd929b8ceb",
  "status": 200,
  "result": [
    {
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
    }
  ],
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
	resp, err := s.client.NewAllOrderList().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AllOrderListResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAllOrderList(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAllOrderList(r1, r2 *AllOrderListResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestAllOrderListResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiAccountTestSuite) assertTestAllOrderListResult(r1, r2 *AllOrderListResult) {
	r := s.r()
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ContingencyType, r2.ContingencyType, "contingencyType")
	r.Equal(r1.ListStatusType, r2.ListStatusType, "listStatusType")
	r.Equal(r1.ListOrderStatus, r2.ListOrderStatus, "listOrderStatus")
	r.Equal(r1.ListClientOrderId, r2.ListClientOrderId, "listClientOrderId")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "transactionTime")
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	for i := range r1.Orders {
		s.assertWsApiOrder(r1.Orders[i], r2.Orders[i])
	}
}

func (s *apiAccountTestSuite) TestAccountTradeHistory() {
	msg := []byte(`{
  "id": "44bde973-abd5-47dc-ad6d-bacdea1c91d1",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "id": 1379446,
      "orderId": 4036985,
      "orderListId": -1,
      "price": "102300.00000000",
      "qty": "0.00010000",
      "quoteQty": "10.23000000",
      "commission": "0.00000000",
      "commissionAsset": "BTC",
      "time": 1737641826090,
      "isBuyer": true,
      "isMaker": false,
      "isBestMatch": true
    }
  ],
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
	resp, err := s.client.NewAccountTradeHistory().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountTradeHistoryResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountTradeHistory(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAccountTradeHistory(r1, r2 *AccountTradeHistoryResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestAccountTradeHistoryResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiAccountTestSuite) assertTestAccountTradeHistoryResult(r1, r2 *AccountTradeHistoryResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.Id, r2.Id, "id")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.Qty, r2.Qty, "qty")
	r.Equal(r1.QuoteQty, r2.QuoteQty, "quoteQty")
	r.Equal(r1.Commission, r2.Commission, "commission")
	r.Equal(r1.CommissionAsset, r2.CommissionAsset, "commissionAsset")
	r.Equal(r1.Time, r2.Time, "time")
	r.Equal(r1.IsBuyer, r2.IsBuyer, "isBuyer")
	r.Equal(r1.IsMaker, r2.IsMaker, "isMaker")
	r.Equal(r1.IsBestMatch, r2.IsBestMatch, "isBestMatch")
}

func (s *apiAccountTestSuite) TestAccountPreventedMatches() {
	msg := []byte(`{
  "id": "g4ce6a53-a39d-4f71-823b-4ab5r391d6y8",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "preventedMatchId": 1,
      "takerOrderId": 5,
      "makerSymbol": "BTCUSDT", 
      "makerOrderId": 3,
      "tradeGroupId": 1,
      "selfTradePreventionMode": "EXPIRE_MAKER",
      "price": "1.100000",
      "makerPreventedQuantity": "1.300000",
      "transactTime": 1669101687094
    }
  ],
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
	resp, err := s.client.NewAccountPreventedMatches().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountPreventedMatchesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountPreventedMatches(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAccountPreventedMatches(r1, r2 *AccountPreventedMatchesResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestAccountPreventedMatchesResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiAccountTestSuite) assertTestAccountPreventedMatchesResult(r1, r2 *AccountPreventedMatchesResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.PreventedMatchId, r2.PreventedMatchId, "preventedMatchId")
	r.Equal(r1.TakerOrderId, r2.TakerOrderId, "takerOrderId")
	r.Equal(r1.MakerSymbol, r2.MakerSymbol, "makerSymbol")
	r.Equal(r1.MakerOrderId, r2.MakerOrderId, "makerOrderId")
	r.Equal(r1.TradeGroupId, r2.TradeGroupId, "tradeGroupId")
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "selfTradePreventionMode")
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.MakerPreventedQuantity, r2.MakerPreventedQuantity, "makerPreventedQuantity")
	r.Equal(r1.TransactTime, r2.TransactTime, "transactTime")
}

func (s *apiAccountTestSuite) TestAccountAllocations() {
	msg := []byte(`{
  "id": "g4ce6a53-a39d-4f71-823b-4ab5r391d6y8",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "allocationId": 0,
      "allocationType": "SOR",
      "orderId": 500,
      "orderListId": -1,
      "price": "1.00000000",
      "qty": "0.10000000",
      "quoteQty": "0.10000000",
      "commission": "0.00000000",
      "commissionAsset": "BTC",
      "time": 1687319487614,
      "isBuyer": false,
      "isMaker": false,
      "isAllocator": false
    }
  ],
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
	resp, err := s.client.NewAccountAllocations().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountAllocationsResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountAllocations(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAccountAllocations(r1, r2 *AccountAllocationsResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestAccountAllocationsResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiAccountTestSuite) assertTestAccountAllocationsResult(r1, r2 *AccountAllocationsResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.AllocationId, r2.AllocationId, "allocationId")
	r.Equal(r1.AllocationType, r2.AllocationType, "allocationType")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.Qty, r2.Qty, "qty")
	r.Equal(r1.QuoteQty, r2.QuoteQty, "quoteQty")
	r.Equal(r1.Commission, r2.Commission, "commission")
	r.Equal(r1.CommissionAsset, r2.CommissionAsset, "commissionAsset")
	r.Equal(r1.Time, r2.Time, "time")
	r.Equal(r1.IsBuyer, r2.IsBuyer, "isBuyer")
	r.Equal(r1.IsMaker, r2.IsMaker, "isMaker")
	r.Equal(r1.IsAllocator, r2.IsAllocator, "isAllocator")
}

func (s *apiAccountTestSuite) TestAccountCommission() {
	msg := []byte(`{
  "id": "23c3f105-9778-429b-b264-10d848eab4c7",
  "status": 200,
  "result": {
    "symbol": "BTCUSDT",
    "standardCommission": {
      "maker": "0.00000000",
      "taker": "0.00000000",
      "buyer": "0.00000000",
      "seller": "0.00000000"
    },
    "taxCommission": {
      "maker": "0.00000000",
      "taker": "0.00000000",
      "buyer": "0.00000000",
      "seller": "0.00000000"
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
	resp, err := s.client.NewAccountCommission().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountCommissionResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountCommission(resp, testResp)
}

func (s *apiAccountTestSuite) assertTestAccountCommission(r1, r2 *AccountCommissionResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.Symbol, r2.Result.Symbol, "symbol")
	s.assertWsCommissionRate(r1.Result.StandardCommission, r2.Result.StandardCommission)
	s.assertWsCommissionRate(r1.Result.TaxCommission, r2.Result.TaxCommission)
	r.Equal(r1.Result.Discount.EnabledForAccount, r2.Result.Discount.EnabledForAccount, "discount.EnabledForAccount")
	r.Equal(r1.Result.Discount.EnabledForSymbol, r2.Result.Discount.EnabledForSymbol, "discount.EnabledForSymbol")
	r.Equal(r1.Result.Discount.DiscountAsset, r2.Result.Discount.DiscountAsset, "discount.DiscountAsset")
	r.Equal(r1.Result.Discount.Discount, r2.Result.Discount.Discount, "discount.Discount")
}
