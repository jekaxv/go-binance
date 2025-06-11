package spot

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type spotAccountTestSuite struct {
	baseHttpTestSuite
}

func TestApiAccount(t *testing.T) {
	suite.Run(t, new(spotAccountTestSuite))
}

func (s *spotAccountTestSuite) TestAccountInformation() {
	msg := []byte(`{
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
  "updateTime": 1746623165116,
  "accountType": "SPOT",
  "balances": [
    {
      "asset": "ETH",
      "free": "1.00000000",
      "locked": "0.00000000"
    }
  ],
  "permissions": [
    "SPOT"
  ],
  "uid": 10000
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAccountInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountInfo(resp, testResp)
}

func (s *spotAccountTestSuite) assertTestAccountInfo(r1, r2 *AccountInfoResponse) {
	r := s.r()
	r.Equal(r1.MakerCommission, r2.MakerCommission, "makerCommission")
	r.Equal(r1.TakerCommission, r2.TakerCommission, "takerCommission")
	r.Equal(r1.BuyerCommission, r2.BuyerCommission, "buyerCommission")
	r.Equal(r1.SellerCommission, r2.SellerCommission, "sellerCommission")
	r.Equal(r1.CanTrade, r2.CanTrade, "canTrade")
	r.Equal(r1.CanWithdraw, r2.CanWithdraw, "canWithdraw")
	r.Equal(r1.CanDeposit, r2.CanDeposit, "canDeposit")
	s.assertCommissionRate(r1.CommissionRates, r2.CommissionRates)
	r.Equal(r1.Brokered, r2.Brokered, "brokered")
	r.Equal(r1.RequireSelfTradePrevention, r2.RequireSelfTradePrevention, "requireSelfTradePrevention")
	r.Equal(r1.PreventSor, r2.PreventSor, "preventSor")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "updateTime")
	r.Equal(r1.AccountType, r2.AccountType, "accountType")
	for i := range r1.Balances {
		s.assertBalance(r1.Balances[i], r2.Balances[i])
	}
	for i := range r1.Permissions {
		r.Equal(r1.Permissions[i], r2.Permissions[i], "permissions")
	}
	r.Equal(r1.Uid, r2.Uid, "uid")
}

func (s *spotAccountTestSuite) TestAccountTradeList() {
	msg := []byte(`[
  {
    "symbol": "BNBBTC",
    "id": 28457,
    "orderId": 100234,
    "orderListId": -1,
    "price": "4.00000100",
    "qty": "12.00000000",
    "quoteQty": "48.000012",
    "commission": "10.10000000",
    "commissionAsset": "BNB",
    "time": 1499865549590,
    "isBuyer": true,
    "isMaker": false,
    "isBestMatch": true
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAccountTrade().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*AccountTradeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestAccountOrderHistoryResult(resp[i], testResp[i])
	}
}

func (s *spotAccountTestSuite) assertTestAccountOrderHistoryResult(r1, r2 *AccountTradeResponse) {
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

func (s *spotAccountTestSuite) TestUnfilledOrder() {
	msg := []byte(`[{"rateLimitType":"ORDERS","interval":"SECOND","intervalNum":10,"limit":50,"count":0},{"rateLimitType":"ORDERS","interval":"DAY","intervalNum":1,"limit":160000,"count":0}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryUnfilledOrder().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryUnfilledOrderResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestUnfilledOrder(resp, testResp)

}

func (s *spotAccountTestSuite) assertTestUnfilledOrder(r1, r2 []*QueryUnfilledOrderResponse) {
	for i := range r1 {
		s.assertTestUnfilledOrderResult(r1[i], r2[i])
	}
}

func (s *spotAccountTestSuite) assertTestUnfilledOrderResult(r1, r2 *QueryUnfilledOrderResponse) {
	r := s.r()
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.IntervalNum, r2.IntervalNum, "intervalNum")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.Count, r2.Count, "count")
}

func (s *spotAccountTestSuite) TestQueryPreventedMatches() {
	msg := []byte(`[
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
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryPreventedMatches().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryPreventedMatchesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestQueryPreventedMatches(resp[i], testResp[i])
	}
}

func (s *spotAccountTestSuite) assertTestQueryPreventedMatches(r1, r2 *QueryPreventedMatchesResponse) {
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

func (s *spotAccountTestSuite) TestQueryAllocations() {
	msg := []byte(`[
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
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryAllocations().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryAllocationsResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestQueryAllocations(resp[i], testResp[i])
	}
}

func (s *spotAccountTestSuite) assertTestQueryAllocations(r1, r2 *QueryAllocationsResponse) {
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

func (s *spotAccountTestSuite) TestQueryCommissionRates() {
	msg := []byte(`{
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
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryCommission().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryCommissionResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountCommission(resp, testResp)
}

func (s *spotAccountTestSuite) assertTestAccountCommission(r1, r2 *QueryCommissionResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	s.assertCommissionRate(r1.StandardCommission, r2.StandardCommission)
	s.assertCommissionRate(r1.TaxCommission, r2.TaxCommission)
	r.Equal(r1.Discount.EnabledForAccount, r2.Discount.EnabledForAccount, "discount.EnabledForAccount")
	r.Equal(r1.Discount.EnabledForSymbol, r2.Discount.EnabledForSymbol, "discount.EnabledForSymbol")
	r.Equal(r1.Discount.DiscountAsset, r2.Discount.DiscountAsset, "discount.DiscountAsset")
	r.Equal(r1.Discount.Discount, r2.Discount.Discount, "discount.Discount")
}

func (s *spotAccountTestSuite) TestUserDataStreamEndpoints() {
	msg := []byte(`{"listenKey":"pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewStartUserDataStream().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *StartUserDataStreamResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(testResp.ListenKey, resp.ListenKey, "listenKey")
}

func (s *spotAccountTestSuite) TestCloseUserDataStream() {
	msg := []byte(`{}`)
	server := s.setup(msg)
	defer server.Close()
	err := s.client.NewCloseUserDataStream().ListenKey("pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1").Do(context.Background())
	r := s.r()
	r.Empty(err)
}

func (s *spotAccountTestSuite) TestPingUserDataStream() {
	msg := []byte(`{}`)
	server := s.setup(msg)
	defer server.Close()
	err := s.client.NewPingUserDataStream().ListenKey("pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1").Do(context.Background())
	r := s.r()
	r.Empty(err)
}
