package https

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiGeneralTestSuite struct {
	baseTestSuite
}

func TestApiGeneralAccount(t *testing.T) {
	suite.Run(t, new(apiGeneralTestSuite))
}

func (s *apiGeneralTestSuite) TestNewPing() {
	msg := []byte(`{}`)
	server := s.setup(msg)
	defer server.Close()
	err := s.client.NewPing().Do(context.Background())
	r := s.r()
	r.Empty(err)
}

func (s *apiGeneralTestSuite) TestNewServerTime() {
	msg := []byte(`{
  "serverTime": 1499827319559
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewServerTime().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ServerTimeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.ServerTime, testResp.ServerTime, "ServerTime")
}

func (s *apiGeneralTestSuite) TestNewExchangeInfo() {
	msg := []byte(`{
  "timezone": "UTC",
  "serverTime": 1565246363776,
  "rateLimits": [
    {}
  ],
  "exchangeFilters": [],
  "symbols": [
    {
      "symbol": "ETHBTC",
      "status": "TRADING",
      "baseAsset": "ETH",
      "baseAssetPrecision": 8,
      "quoteAsset": "BTC",
      "quotePrecision": 8,
      "quoteAssetPrecision": 8,
      "baseCommissionPrecision": 8,
      "quoteCommissionPrecision": 8,
      "orderTypes": [
        "LIMIT",
        "LIMIT_MAKER",
        "MARKET",
        "STOP_LOSS",
        "STOP_LOSS_LIMIT",
        "TAKE_PROFIT",
        "TAKE_PROFIT_LIMIT"
      ],
      "icebergAllowed": true,
      "ocoAllowed": true,
      "otoAllowed": true,
      "quoteOrderQtyMarketAllowed": true,
      "allowTrailingStop": false,
      "cancelReplaceAllowed": false,
      "allowAmend": false,
      "isSpotTradingAllowed": true,
      "isMarginTradingAllowed": true,
      "filters": [],
      "permissions": [],
      "permissionSets": [
        [
          "SPOT",
          "MARGIN"
        ]
      ],
      "defaultSelfTradePreventionMode": "NONE",
      "allowedSelfTradePreventionModes": [
        "NONE"
      ]
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewExchangeInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ExchangeInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestExchangeInfoResponse(resp, testResp)
}

func (s *apiGeneralTestSuite) assertTestExchangeInfoResponse(r1, r2 *ExchangeInfoResponse) {
	r := s.r()
	r.Equal(r1.Timezone, r2.Timezone, "Timezone")
	r.Equal(r1.ServerTime, r2.ServerTime, "ServerTime")
	for i := range r1.RateLimits {
		s.assertRateLimit(r1.RateLimits[i], r2.RateLimits[i])
	}
	for i := range r1.ExchangeFilters {
		s.assertExchangeFilter(r1.ExchangeFilters[i], r2.ExchangeFilters[i])
	}
	for i := range r1.Symbols {
		s.assertSymbolInfo(r1.Symbols[i], r2.Symbols[i])
	}
}
