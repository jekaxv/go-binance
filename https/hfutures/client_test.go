package hfutures

import (
	"github.com/jekaxv/go-binance/https"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
)

type mockedClient struct {
	mock.Mock
	*Client
}

type baseTestSuite struct {
	suite.Suite
	client *mockedClient
}

func (s *baseTestSuite) SetupTest() {
	s.client = new(mockedClient)
	client := Client{
		C: &https.Client{
			Opt: &https.Options{
				ApiKey:    "Api Key",
				ApiSecret: "Api Secret",
			},
			HttpClient: http.DefaultClient,
		},
	}
	s.client.Client = &client
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) mockServer(msg []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(msg)
	}))
}

func (s *baseTestSuite) setup(msg []byte) *httptest.Server {
	server := s.mockServer(msg)
	s.client.C.Opt.Endpoint = server.URL
	return server
}

func (s *baseTestSuite) assertCommissionRate(r1, r2 *https.SpotCommission) {
	r := s.r()
	r.Equal(r1.Maker, r2.Maker, "maker")
	r.Equal(r1.Taker, r2.Taker, "taker")
	r.Equal(r1.Buyer, r2.Buyer, "buyer")
	r.Equal(r1.Seller, r2.Seller, "seller")
}

func (s *baseTestSuite) assertBalance(r1, r2 *https.ApiBalance) {
	r := s.r()
	r.Equal(r1.Asset, r2.Asset, "asset")
	r.Equal(r1.Free, r2.Free, "free")
	r.Equal(r1.Locked, r2.Locked, "locked")
}

func (s *baseTestSuite) assertApiFill(r1, r2 *https.SpotFill) {
	r := s.r()
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.Qty, r2.Qty, "qty")
	r.Equal(r1.Commission, r2.Commission, "commission")
	r.Equal(r1.CommissionAsset, r2.CommissionAsset, "commissionAsset")
	r.Equal(r1.TradeId, r2.TradeId, "tradeId")
	r.Equal(r1.MatchType, r2.MatchType, "matchType")
	r.Equal(r1.AllocId, r2.AllocId, "allocId")
}

func (s *baseTestSuite) assertApiOrder(r1, r2 *https.SpotOrder) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
}

func (s *baseTestSuite) assertApiOrderReport(r1, r2 *https.OrderReport) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.OrderListId, r2.OrderListId, "orderListId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
	r.Equal(r1.OrigClientOrderId, r2.OrigClientOrderId, "origClientOrderId")
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
	r.Equal(r1.SelfTradePreventionMode, r2.SelfTradePreventionMode, "selfTradePreventionMode")
}

func (s *baseTestSuite) assertRateLimit(r1, r2 *https.RateLimit) {
	r := s.r()
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
}

func (s *baseTestSuite) assertExchangeFilter(r1, r2 *https.ExchangeFilter) {
	r := s.r()
	r.Equal(r1.FilterType, r2.FilterType, "filterType")
	r.Equal(r1.MaxNumAlgoOrders, r2.MaxNumAlgoOrders, "maxNumAlgoOrders")
}

func (s *baseTestSuite) assertFilters(r1, r2 *SymbolFilter) {
	r := s.r()
	r.Equal(r1.FilterType, r2.FilterType, "filterType")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.MaxPrice, r2.MaxPrice, "maxPrice")
	r.Equal(r1.MaxQty, r2.MaxQty, "maxQty")
	r.Equal(r1.MinPrice, r2.MinPrice, "minPrice")
	r.Equal(r1.MinQty, r2.MinQty, "minQty")
	r.Equal(r1.StepSize, r2.StepSize, "stepSize")
	r.Equal(r1.TickSize, r2.TickSize, "tickSize")
}

func (s *baseTestSuite) assertSymbolInfo(r1, r2 *SymbolInfo) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.Pair, r2.Pair, "pair")
	r.Equal(r1.ContractType, r2.ContractType, "contractType")
	r.Equal(r1.DeliveryDate, r2.DeliveryDate, "deliveryDate")
	r.Equal(r1.OnboardDate, r2.OnboardDate, "onboardDate")
	r.Equal(r1.Status, r2.Status, "status")
	r.Equal(r1.MaintMarginPercent, r2.MaintMarginPercent, "maintMarginPercent")
	r.Equal(r1.RequiredMarginPercent, r2.RequiredMarginPercent, "requiredMarginPercent")
	r.Equal(r1.BaseAsset, r2.BaseAsset, "baseAsset")
	r.Equal(r1.QuoteAsset, r2.QuoteAsset, "quoteAsset")
	r.Equal(r1.MarginAsset, r2.MarginAsset, "marginAsset")
	r.Equal(r1.PricePrecision, r2.PricePrecision, "pricePrecision")
	r.Equal(r1.QuantityPrecision, r2.QuantityPrecision, "quantityPrecision")
	r.Equal(r1.BaseAssetPrecision, r2.BaseAssetPrecision, "baseAssetPrecision")
	r.Equal(r1.QuotePrecision, r2.QuotePrecision, "quotePrecision")
	r.Equal(r1.UnderlyingType, r2.UnderlyingType, "underlyingType")
	r.Equal(r1.UnderlyingSubType, r2.UnderlyingSubType, "underlyingSubType")
	r.Equal(r1.SettlePlan, r2.SettlePlan, "settlePlan")
	r.Equal(r1.TriggerProtect, r2.TriggerProtect, "triggerProtect")
	for i := range r1.OrderType {
		r.Equal(r1.OrderType[i], r2.OrderType[i], "orderTypes")
	}
	for i := range r1.Filters {
		s.assertFilters(r1.Filters[i], r2.Filters[i])
	}
	r.Equal(r1.LiquidationFee, r2.LiquidationFee, "liquidationFee")
	r.Equal(r1.MarketTakeBound, r2.MarketTakeBound, "marketTakeBound")
}
