package wss

import (
	"github.com/gorilla/websocket"
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

func (s *baseTestSuite) mockClient(url string) {
	s.client.Client.Opt.Endpoint = url
}

func (s *baseTestSuite) SetupTest() {
	s.client = new(mockedClient)
	client := Client{
		Opt: &Options{
			ApiKey:    "Api Key",
			ApiSecret: "Api Secret",
		},
	}
	s.client.Client = &client
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) assertWsResponse(r1, r2 ApiResponse) {
	r := s.r()
	r.Equal(r1.Id, r2.Id, "id")
	r.Equal(r1.Status, r2.Status, "status")
	if r1.Error != nil {
		s.assertWsApiError(r1.Error, r2.Error)
	}
	for i := range r1.RateLimits {
		s.assertWsRateLimits(r1.RateLimits[i], r2.RateLimits[i])
	}
}

func (s *baseTestSuite) assertWsApiError(r1, r2 *ApiError) {
	r := s.r()
	r.Equal(r1.Code, r2.Code, "code")
	r.Equal(r1.Msg, r2.Msg, "msg")
}

func (s *baseTestSuite) assertWsRateLimits(r1, r2 *ApiRateLimit) {
	r := s.r()
	r.Equal(r1.Count, r2.Count, "count")
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.IntervalNum, r2.IntervalNum, "intervalNum")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
}

func (s *baseTestSuite) assertWsApiSymbol(r1, r2 *ApiSymbol) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.Status, r2.Status, "status")
	r.Equal(r1.BaseAsset, r2.BaseAsset, "baseAsset")
	r.Equal(r1.BaseAssetPrecision, r2.BaseAssetPrecision, "baseAssetPrecision")
	r.Equal(r1.QuoteAsset, r2.QuoteAsset, "quoteAsset")
	r.Equal(r1.QuotePrecision, r2.QuotePrecision, "quotePrecision")
	r.Equal(r1.QuoteAssetPrecision, r2.QuoteAssetPrecision, "quoteAssetPrecision")
	r.Equal(r1.BaseCommissionPrecision, r2.BaseCommissionPrecision, "baseCommissionPrecision")
	r.Equal(r1.QuoteCommissionPrecision, r2.QuoteCommissionPrecision, "quoteCommissionPrecision")
	for i := range r1.OrderTypes {
		r.Equal(r1.OrderTypes[i], r2.OrderTypes[i], "orderTypes")
	}
	r.Equal(r1.IcebergAllowed, r2.IcebergAllowed, "icebergAllowed")
	r.Equal(r1.OcoAllowed, r2.OcoAllowed, "ocoAllowed")
	r.Equal(r1.OtoAllowed, r2.OtoAllowed, "otoAllowed")
	r.Equal(r1.QuoteOrderQtyMarketAllowed, r2.QuoteOrderQtyMarketAllowed, "quoteOrderQtyMarketAllowed")
	r.Equal(r1.AllowTrailingStop, r2.AllowTrailingStop, "allowTrailingStop")
	r.Equal(r1.CancelReplaceAllowed, r2.CancelReplaceAllowed, "cancelReplaceAllowed")
	r.Equal(r1.IsSpotTradingAllowed, r2.IsSpotTradingAllowed, "isSpotTradingAllowed")
	r.Equal(r1.IsMarginTradingAllowed, r2.IsMarginTradingAllowed, "isMarginTradingAllowed")
	for i := range r1.Filters {
		s.assertWsApiFilter(r1.Filters[i], r2.Filters[i])
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

func (s *baseTestSuite) assertWsApiSor(r1, r2 *ApiSort) {
	r := s.r()
	r.Equal(r1.BaseAsset, r2.BaseAsset, "baseAsset")
	for i := range r1.Symbols {
		r.Equal(r1.Symbols[i], r2.Symbols[i], "symbols")
	}
}
func (s *baseTestSuite) assertWsApiOrder(r1, r2 *ApiOrder) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.OrderId, r2.OrderId, "orderId")
	r.Equal(r1.ClientOrderId, r2.ClientOrderId, "clientOrderId")
}

func (s *baseTestSuite) assertWsApiFilter(r1, r2 *ApiFilter) {
	r := s.r()
	r.Equal(r1.FilterType, r2.FilterType, "filterType")
	r.Equal(r1.MinPrice, r2.MinPrice, "minPrice")
	r.Equal(r1.MaxPrice, r2.MaxPrice, "maxPrice")
	r.Equal(r1.TickSize, r2.TickSize, "tickSize")
	r.Equal(r1.MinQty, r2.MinQty, "minQty")
	r.Equal(r1.MaxQty, r2.MaxQty, "maxQty")
	r.Equal(r1.StepSize, r2.StepSize, "stepSize")
}

func (s *baseTestSuite) assertWsApiFill(r1, r2 *ApiFill) {
	r := s.r()
	r.Equal(r1.Price, r2.Price, "price")
	r.Equal(r1.Qty, r2.Qty, "qty")
	r.Equal(r1.Commission, r2.Commission, "commission")
	r.Equal(r1.CommissionAsset, r2.CommissionAsset, "commissionAsset")
	r.Equal(r1.TradeId, r2.TradeId, "tradeId")
	r.Equal(r1.MatchType, r2.MatchType, "matchType")
	r.Equal(r1.AllocId, r2.AllocId, "allocId")
}
func (s *baseTestSuite) assertWsApiOrderReport(r1, r2 *ApiOrderReport) {
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

func (s *baseTestSuite) assertWsCommissionRate(r1, r2 *ApiCommissionRate) {
	r := s.r()
	r.Equal(r1.Maker, r2.Maker, "maker")
	r.Equal(r1.Taker, r2.Taker, "taker")
	r.Equal(r1.Buyer, r2.Buyer, "buyer")
	r.Equal(r1.Seller, r2.Seller, "seller")
}

func (s *baseTestSuite) assertWsBalance(r1, r2 *ApiBalance) {
	r := s.r()
	r.Equal(r1.Asset, r2.Asset, "asset")
	r.Equal(r1.Free, r2.Free, "free")
	r.Equal(r1.Locked, r2.Locked, "locked")
}
func (s *baseTestSuite) mockServer(msg []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		}
	}))
}

func (s *baseTestSuite) setup(msg []byte) *httptest.Server {
	server := s.mockServer(msg)
	s.mockClient("ws" + server.URL[4:])
	return server
}
