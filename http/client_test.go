package http

import (
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
		Opt: &Options{
			ApiKey:    "Api Key",
			ApiSecret: "Api Secret",
		},
		HttpClient: http.DefaultClient,
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
	s.client.Opt.Endpoint = server.URL
	return server
}

func (s *baseTestSuite) assertCommissionRate(r1, r2 *SpotCommission) {
	r := s.r()
	r.Equal(r1.Maker, r2.Maker, "maker")
	r.Equal(r1.Taker, r2.Taker, "taker")
	r.Equal(r1.Buyer, r2.Buyer, "buyer")
	r.Equal(r1.Seller, r2.Seller, "seller")
}

func (s *baseTestSuite) assertBalance(r1, r2 *ApiBalance) {
	r := s.r()
	r.Equal(r1.Asset, r2.Asset, "asset")
	r.Equal(r1.Free, r2.Free, "free")
	r.Equal(r1.Locked, r2.Locked, "locked")
}

func (s *baseTestSuite) assertApiFill(r1, r2 *SpotFill) {
	r := s.r()
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.Qty, r2.Qty, "qty")
	r.Equal(r1.Commission, r2.Commission, "commission")
	r.Equal(r1.CommissionAsset, r2.CommissionAsset, "commissionAsset")
	r.Equal(r1.TradeId, r2.TradeId, "tradeId")
	r.Equal(r1.MatchType, r2.MatchType, "matchType")
	r.Equal(r1.AllocId, r2.AllocId, "allocId")
}

func (s *baseTestSuite) assertApiOrder(r1, r2 *SpotOrder) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
}

func (s *baseTestSuite) assertApiOrderReport(r1, r2 *OrderReport) {
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

func (s *baseTestSuite) assertRateLimit(r1, r2 *RateLimit) {
	r := s.r()
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
}

func (s *baseTestSuite) assertExchangeFilter(r1, r2 *ExchangeFilter) {
	r := s.r()
	r.Equal(r1.FilterType, r2.FilterType, "filterType")
	r.Equal(r1.MaxNumAlgoOrders, r2.MaxNumAlgoOrders, "maxNumAlgoOrders")
}

func (s *baseTestSuite) assertFilters(r1, r2 *SymbolFilter) {
	r := s.r()
	r.Equal(r1.ApplyMinToMarket, r2.ApplyMinToMarket, "applyMinToMarket")
	r.Equal(r1.ApplyMaxToMarket, r2.ApplyMaxToMarket, "applyMaxToMarket")
	r.Equal(r1.AskMultiplierDown, r2.AskMultiplierDown, "askMultiplierDown")
	r.Equal(r1.AskMultiplierUp, r2.AskMultiplierUp, "askMultiplierUp")
	r.Equal(r1.AvgPriceMins, r2.AvgPriceMins, "avgPriceMins")
	r.Equal(r1.BidMultiplierDown, r2.BidMultiplierDown, "bidMultiplierDown")
	r.Equal(r1.BidMultiplierUp, r2.BidMultiplierUp, "bidMultiplierUp")
	r.Equal(r1.FilterType, r2.FilterType, "filterType")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.MaxNotional, r2.MaxNotional, "maxNotional")
	r.Equal(r1.MaxNumAlgoOrders, r2.MaxNumAlgoOrders, "maxNumAlgoOrders")
	r.Equal(r1.MaxNumOrders, r2.MaxNumOrders, "maxNumOrders")
	r.Equal(r1.MaxPrice, r2.MaxPrice, "maxPrice")
	r.Equal(r1.MaxQty, r2.MaxQty, "maxQty")
	r.Equal(r1.MaxTrailingAboveDelta, r2.MaxTrailingAboveDelta, "maxTrailingAboveDelta")
	r.Equal(r1.MaxTrailingBelowDelta, r2.MaxTrailingBelowDelta, "maxTrailingBelowDelta")
	r.Equal(r1.MinNotional, r2.MinNotional, "minNotional")
	r.Equal(r1.MinPrice, r2.MinPrice, "minPrice")
	r.Equal(r1.MinQty, r2.MinQty, "minQty")
	r.Equal(r1.MinTrailingAboveDelta, r2.MinTrailingAboveDelta, "minTrailingAboveDelta")
	r.Equal(r1.MinTrailingBelowDelta, r2.MinTrailingBelowDelta, "minTrailingBelowDelta")
	r.Equal(r1.StepSize, r2.StepSize, "stepSize")
	r.Equal(r1.TickSize, r2.TickSize, "tickSize")
}

func (s *baseTestSuite) assertSymbolInfo(r1, r2 *SymbolInfo) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.Status, r2.Status, "status")
	r.Equal(r1.BaseAsset, r2.BaseAsset, "baseAsset")
	r.Equal(r1.BaseAssetPrecision, r2.BaseAssetPrecision, "baseAssetPrecision")
	r.Equal(r1.QuoteAsset, r2.QuoteAsset, "quoteAsset")
	r.Equal(r1.QuotePrecision, r2.QuotePrecision, "quotePrecision")
	r.Equal(r1.QuoteAssetPrecision, r2.QuoteAssetPrecision, "quoteAssetPrecision")
	for i := range r1.OrderTypes {
		r.Equal(r1.OrderTypes[i], r2.OrderTypes[i], "orderTypes")
	}
	r.Equal(r1.IcebergAllowed, r2.IcebergAllowed, "icebergAllowed")
	r.Equal(r1.OcoAllowed, r2.OcoAllowed, "ocoAllowed")
	r.Equal(r1.QuoteOrderQtyMarketAllowed, r2.QuoteOrderQtyMarketAllowed, "quoteOrderQtyMarketAllowed")
	r.Equal(r1.AllowTrailingStop, r2.AllowTrailingStop, "allowTrailingStop")
	r.Equal(r1.CancelReplaceAllowed, r2.CancelReplaceAllowed, "cancelReplaceAllowed")
	r.Equal(r1.IsSpotTradingAllowed, r2.IsSpotTradingAllowed, "isSpotTradingAllowed")
	r.Equal(r1.IsMarginTradingAllowed, r2.IsMarginTradingAllowed, "isMarginTradingAllowed")
	for i := range r1.Filters {
		s.assertFilters(r1.Filters[i], r2.Filters[i])
	}
	for i := range r1.Permissions {
		r.Equal(r1.Permissions[i], r2.Permissions[i], "permissions")
	}
	for i := range r1.PermissionSets {
		for j := range r1.PermissionSets[i] {
			r.Equal(r1.PermissionSets[i][j], r2.PermissionSets[i][j], "permissionSets")
		}
	}
	r.Equal(r1.DefaultSelfTradePreventionMode, r2.DefaultSelfTradePreventionMode, "defaultSelfTradePreventionMode")
	for i := range r1.AllowedSelfTradePreventionModes {
		r.Equal(r1.AllowedSelfTradePreventionModes[i], r2.AllowedSelfTradePreventionModes[i], "allowedSelfTradePreventionModes")
	}
}
