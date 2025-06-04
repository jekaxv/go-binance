package wfutures

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type marketTestSuite struct {
	baseTestSuite
}

func TestWebsocketMarket(t *testing.T) {
	suite.Run(t, new(marketTestSuite))
}

func (s *marketTestSuite) TestNewDepth() {
	msg := []byte(`{
  "id": "51e2affb-0aba-4821-ba75-f2625006eb43",
  "status": 200,
  "result": {
    "lastUpdateId": 1027024,
    "E": 1589436922972,  
    "T": 1589436922959,   
    "bids": [
      [
        "4.00000000",     
        "431.00000000"    
      ]
    ],
    "asks": [
      [
        "4.00000200",
        "12.00000000"
      ]
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 2400,
      "count": 5
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewDepth().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *DepthResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestDepthResponse(resp, testResp)
}

func (s *marketTestSuite) assertTestDepthResponse(r1, r2 *DepthResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertMarketDepthResult(r1.Result, r2.Result)
}

func (s *marketTestSuite) assertMarketDepthResult(r1, r2 *DepthResult) {
	r := s.r()
	r.Equal(r1.LastUpdateId, r2.LastUpdateId, "LastUpdateId")
	r.Equal(r1.OutputTime, r2.OutputTime, "OutputTime")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "TransactionTime")
	for i := range r1.Asks {
		for j := range r1.Asks[i] {
			r.Equal(r1.Asks[i][j], r2.Asks[i][j], "Asks")
		}
	}
	for i := range r1.Bids {
		for j := range r1.Bids[i] {
			r.Equal(r1.Bids[i][j], r2.Bids[i][j], "Bids")
		}
	}
}

func (s *marketTestSuite) TestNewTickerPrice() {
	msg := []byte(`{
	  "id": "9d32157c-a556-4d27-9866-66760a174b57",
	  "status": 200,
	  "result": [
		{
			"symbol": "BTCUSDT",
			"price": "6000.01",
			"time": 1589437530011
		}
	  ],
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 2400,
		  "count": 2
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTickerPrice().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *TickerPriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestTickerPriceResponse(resp, testResp)
}

func (s *marketTestSuite) assertTestTickerPriceResponse(r1, r2 *TickerPriceResponse) {
	r := s.r()
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		r.Equal(r1.Result[i].Symbol, r2.Result[i].Symbol, "Symbol")
		r.Equal(r1.Result[i].Price, r2.Result[i].Price, "Price")
		r.Equal(r1.Result[i].Time, r2.Result[i].Time, "Time")
	}
}

func (s *marketTestSuite) TestNewTickerBook() {
	msg := []byte(`{
	  "id": "9d32157c-a556-4d27-9866-66760a174b57",
	  "status": 200,
	  "result": [
		{
		  "lastUpdateId": 1027024,
		  "symbol": "BTCUSDT",
		  "bidPrice": "4.00000000",
		  "bidQty": "431.00000000",
		  "askPrice": "4.00000200",
		  "askQty": "9.00000000",
		  "time": 1589437530011
		}
	  ],
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 2400,
		  "count": 2
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTickerBook().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *TickerBookResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestTickerBookResponse(resp, testResp)
}

func (s *marketTestSuite) assertTestTickerBookResponse(r1, r2 *TickerBookResponse) {
	r := s.r()
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		r.Equal(r1.Result[i].Symbol, r2.Result[i].Symbol, "Symbol")
		r.Equal(r1.Result[i].BidPrice, r2.Result[i].BidPrice, "BidPrice")
		r.Equal(r1.Result[i].BidQty, r2.Result[i].BidQty, "BidQty")
		r.Equal(r1.Result[i].AskPrice, r2.Result[i].AskPrice, "AskPrice")
		r.Equal(r1.Result[i].AskQty, r2.Result[i].AskQty, "AskQty")
		r.Equal(r1.Result[i].Time, r2.Result[i].Time, "Time")
	}
}
