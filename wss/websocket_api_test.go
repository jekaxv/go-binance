package wss

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type websocketApiTestSuite struct {
	baseTestSuite
}

func TestWebsocketApi(t *testing.T) {
	suite.Run(t, new(websocketApiTestSuite))
}

func (s *websocketApiTestSuite) TestWebSocketConnection() {
	msg := []byte(`{"id":"90feb8c2-3308-4f8e-956d-fef891859370","status":200,"result":{},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":3}]}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewPing().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *PingResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestConnectivity(resp, testResp)
}

func (s *websocketApiTestSuite) assertTestConnectivity(r1, r2 *PingResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
}

func (s *websocketApiTestSuite) TestWebSocketServerTime() {
	msg := []byte(`{"id":"4aac6953-8c8a-45c8-9066-f92d881e97c0","status":200,"result":{"serverTime":1737440469538},"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":3}]}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCheckServerTime().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CheckServerTimeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestServerTime(resp, testResp)
}

func (s *websocketApiTestSuite) assertTestServerTime(r1, r2 *CheckServerTimeResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r2.Result.ServerTime, r1.Result.ServerTime)
}

func (s *websocketApiTestSuite) TestWebSocketExchangeInformation() {
	msg := []byte(`{
  "id": "0feed23a-f49c-4fb3-b953-bea3cebf32b6",
  "status": 200,
  "result": {
    "timezone": "UTC",
    "serverTime": 1737441362536,
    "rateLimits": [
      {
        "rateLimitType": "CONNECTIONS",
        "interval": "MINUTE",
        "intervalNum": 5,
        "limit": 300
      }
    ],
    "exchangeFilters": [],
    "symbols": [
      {
        "symbol": "BTCUSDT",
        "status": "TRADING",
        "baseAsset": "BTC",
        "baseAssetPrecision": 8,
        "quoteAsset": "USDT",
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
        "allowTrailingStop": true,
        "cancelReplaceAllowed": true,
        "isSpotTradingAllowed": true,
        "isMarginTradingAllowed": false,
        "filters": [
          {
            "filterType": "MAX_NUM_ALGO_ORDERS",
            "maxNumAlgoOrders": 5
          }
        ],
        "permissions": [],
        "permissionSets": [
          [
            "SPOT"
          ]
        ],
        "defaultSelfTradePreventionMode": "EXPIRE_MAKER",
        "allowedSelfTradePreventionModes": [
          "NONE",
          "EXPIRE_TAKER",
          "EXPIRE_MAKER",
          "EXPIRE_BOTH"
        ]
      }
    ]
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
	resp, err := s.client.NewExchangeInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ExchangeInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestExchangeInformation(resp, testResp)
}

func (s *websocketApiTestSuite) assertTestExchangeInformation(r1, r2 *ExchangeInfoResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r2.Result.ServerTime, r1.Result.ServerTime)
	r.Equal(r2.Result.Timezone, r1.Result.Timezone)
	for i := range r1.Result.RateLimits {
		s.assertWsRateLimits(r1.Result.RateLimits[i], r2.Result.RateLimits[i])
	}
	for i := range r1.Result.Symbols {
		s.assertWsApiSymbol(r1.Result.Symbols[i], r2.Result.Symbols[i])
	}
	for i := range r1.Result.Sors {
		s.assertWsApiSor(r1.Result.Sors[i], r2.Result.Sors[i])
	}
}
