package futures

import (
	"github.com/jekaxv/go-binance/core"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"net/http"
	"net/http/httptest"
)

type mockedHttpClient struct {
	mock.Mock
	*Client
}

type baseHttpTestSuite struct {
	suite.Suite
	client *mockedHttpClient
}

func (s *baseHttpTestSuite) SetupTest() {
	s.client = new(mockedHttpClient)
	client := Client{
		C: &core.Client{
			Opt: &core.Options{
				ApiKey:    "YOUR_API_KEY",
				ApiSecret: "YOUR_API_SECRET",
			},
			HttpClient: http.DefaultClient,
			Logger:     slog.Default(),
		},
	}
	s.client.Client = &client
}

func (s *baseHttpTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseHttpTestSuite) mockServer(msg []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(msg)
	}))
}

func (s *baseHttpTestSuite) setup(msg []byte) *httptest.Server {
	server := s.mockServer(msg)
	s.client.C.Opt.Endpoint = server.URL
	return server
}

func (s *baseHttpTestSuite) assertFilters(r1, r2 *SymbolFilter) {
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

func (s *baseHttpTestSuite) assertRateLimit(r1, r2 *RateLimit) {
	r := s.r()
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
}

func (s *baseHttpTestSuite) assertExchangeFilter(r1, r2 *ExchangeFilter) {
	r := s.r()
	r.Equal(r1.FilterType, r2.FilterType, "filterType")
	r.Equal(r1.MaxNumAlgoOrders, r2.MaxNumAlgoOrders, "maxNumAlgoOrders")
}

func (s *baseHttpTestSuite) assertSymbolInfo(r1, r2 *SymbolInfo) {
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
