package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiMarketTestSuite struct {
	baseWsTestSuite
}

func TestWebsocketApiMarket(t *testing.T) {
	suite.Run(t, new(apiMarketTestSuite))
}

func (s *apiMarketTestSuite) TestMarketDepth() {
	msg := []byte(`{
  "id": "6f61cab4-116e-42f7-8264-8acb99e9f95e",
  "status": 200,
  "result": {
    "lastUpdateId": 5277603,
    "bids": [
      [
        "103795.20000000",
        "0.28870000"
      ],
      [
        "103792.00000000",
        "0.00444000"
      ]
    ],
    "asks": [
      [
        "104009.50000000",
        "0.00387000"
      ],
      [
        "104014.35000000",
        "0.04022000"
      ]
    ]
  },
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 7
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewDepth().Symbol("BTCUSDT").Limit(2).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsDepthResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketDepth(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketDepth(r1, r2 *WsDepthResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertMarketDepthResult(r1.Result, r2.Result)
}

func (s *apiMarketTestSuite) assertMarketDepthResult(r1, r2 *DepthResult) {
	r := s.r()
	r.Equal(r1.LastUpdateId, r2.LastUpdateId, "LastUpdateId")
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

func (s *apiMarketTestSuite) TestMarketTradesRecent() {
	msg := []byte(`{
  "id": "3f75853e-5031-4221-a6e8-7732358a2b16",
  "status": 200,
  "result": [
    {
      "id": 1249997,
      "price": "104237.54000000",
      "qty": "0.00019000",
      "quoteQty": "19.80513260",
      "time": 1737555541551,
      "isBuyerMaker": false,
      "isBestMatch": true
    },
    {
      "id": 1249998,
      "price": "104237.54000000",
      "qty": "0.00539000",
      "quoteQty": "561.84034060",
      "time": 1737555543079,
      "isBuyerMaker": false,
      "isBestMatch": true
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 27
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTradesRecent().Symbol("BTCUSDT").Limit(2).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTradesRecentResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTradesRecent(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketTradesRecent(r1, r2 *WsTradesRecentResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTradesResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarketTradesResult(r1, r2 *TradesResult) {
	r := s.r()
	r.Equal(r1.Id, r2.Id, "Id")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Qty, r2.Qty, "Qty")
	r.Equal(r1.QuoteQty, r2.QuoteQty, "QuoteQty")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.IsBuyerMaker, r2.IsBuyerMaker, "IsBuyerMaker")
	r.Equal(r1.IsBestMatch, r2.IsBestMatch, "IsBestMatch")
}

func (s *apiMarketTestSuite) TestMarketTradesHistorical() {
	msg := []byte(`{
  "id": "83cae65a-a017-417f-90a2-851bdab203e9",
  "status": 200,
  "result": [
    {
      "id": 1250233,
      "price": "104300.01000000",
      "qty": "0.00192000",
      "quoteQty": "200.25601920",
      "time": 1737555820082,
      "isBuyerMaker": false,
      "isBestMatch": true
    },
    {
      "id": 1250234,
      "price": "104300.01000000",
      "qty": "0.00076000",
      "quoteQty": "79.26800760",
      "time": 1737555821091,
      "isBuyerMaker": false,
      "isBestMatch": true
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 27
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTradesHistorical().Symbol("BTCUSDT").Limit(2).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTradesHistoricalResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTradesHistorical(resp, testResp)
}
func (s *apiMarketTestSuite) assertTestMarketTradesHistorical(r1, r2 *WsTradesHistoricalResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTradesResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) TestMarketTradesAggregate() {
	msg := []byte(`{
  "id": "68076b19-d7c5-4cb1-aa96-5ad65c7d1314",
  "status": 200,
  "result": [
    {
      "a": 1193213,
      "p": "104493.11000000",
      "q": "0.00130000",
      "f": 1250458,
      "l": 1250458,
      "T": 1737556039748,
      "m": false,
      "M": true
    },
    {
      "a": 1193214,
      "p": "104498.45000000",
      "q": "0.00220000",
      "f": 1250459,
      "l": 1250459,
      "T": 1737556039748,
      "m": false,
      "M": true
    }
  ],
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
	resp, err := s.client.NewTradesAggregate().Symbol("BTCUSDT").Limit(2).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTradesAggregateResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTradesAggregate(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketTradesAggregate(r1, r2 *WsTradesAggregateResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTradesAggregateResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarketTradesAggregateResult(r1, r2 *TradesAggregateResult) {
	r := s.r()
	r.Equal(r1.TradeId, r2.TradeId, "TradeId")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Quantity, r2.Quantity, "Quantity")
	r.Equal(r1.FirstId, r2.FirstId, "FirstId")
	r.Equal(r1.LastId, r2.LastId, "LastId")
	r.Equal(r1.Timestamp, r2.Timestamp, "Timestamp")
	r.Equal(r1.IsMaker, r2.IsMaker, "IsMaker")
	r.Equal(r1.IsBestPrice, r2.IsBestPrice, "IsBestPrice")
}
func toDecimal(v string) decimal.Decimal {
	float, _ := decimal.NewFromString(v)
	return float
}
func (s *apiMarketTestSuite) TestMarketKline() {
	msg := []byte(`{
  "id": "1b668cdc-72d7-43e4-952b-00bd50702f6b",
  "status": 200,
  "result": [
    [
      1737556140000,
      "104552.00000000",
      "104650.25000000",
      "104551.87000000",
      "104650.25000000",
      "0.34488000",
      1737556199999,
      "36077.69853200",
      122,
      "0.22480000",
      "23513.66461330",
      "0"
    ],
    [
      1737556200000,
      "104634.38000000",
      "104644.84000000",
      "104596.68000000",
      "104605.60000000",
      "0.20685000",
      1737556259999,
      "21637.65114840",
      96,
      "0.10046000",
      "10508.39407400",
      "0"
    ]
  ],
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
	resp, err := s.client.NewKline().Symbol("BTCUSDT").Interval(core.Interval1m).Limit(2).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := WsKlineResponse{
		ApiResponse: ApiResponse{
			Id:     "1b668cdc-72d7-43e4-952b-00bd50702f6b",
			Status: 200,
			RateLimits: []*ApiRateLimit{{
				RateLimitType: "REQUEST_WEIGHT",
				Interval:      "MINUTE",
				IntervalNum:   1,
				Limit:         6000,
				Count:         4,
			}},
		},
		Result: []*KlineResult{{
			OpenTime:                 1737556140000,
			OpenPrice:                toDecimal("104552.00000000"),
			HighPrice:                toDecimal("104650.25000000"),
			LowPrice:                 toDecimal("104551.87000000"),
			ClosePrice:               toDecimal("104650.25000000"),
			Volume:                   toDecimal("0.34488000"),
			CloseTime:                1737556199999,
			QuoteAssetVolume:         toDecimal("36077.69853200"),
			NumberOfTrades:           122,
			TakerBuyBaseAssetVolume:  toDecimal("0.22480000"),
			TakerBuyQuoteAssetVolume: toDecimal("23513.66461330"),
		}, {
			OpenTime:                 1737556200000,
			OpenPrice:                toDecimal("104634.38000000"),
			HighPrice:                toDecimal("104644.84000000"),
			LowPrice:                 toDecimal("104596.68000000"),
			ClosePrice:               toDecimal("104605.60000000"),
			Volume:                   toDecimal("0.20685000"),
			CloseTime:                1737556259999,
			QuoteAssetVolume:         toDecimal("21637.65114840"),
			NumberOfTrades:           96,
			TakerBuyBaseAssetVolume:  toDecimal("0.10046000"),
			TakerBuyQuoteAssetVolume: toDecimal("10508.39407400"),
		}},
	}
	s.assertTestMarketKline(resp, &testResp)
}

func (s *apiMarketTestSuite) assertTestMarketKline(r1, r2 *WsKlineResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketKlineResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarketKlineResult(r1, r2 *KlineResult) {
	r := s.r()
	r.Equal(r1.OpenTime, r2.OpenTime, "OpenTime")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.ClosePrice, r2.ClosePrice, "ClosePrice")
	r.Equal(r1.CloseTime, r2.CloseTime, "CloseTime")
	r.Equal(r1.Volume, r2.Volume, "Volume")
	r.Equal(r1.QuoteAssetVolume, r2.QuoteAssetVolume, "QuoteAssetVolume")
	r.Equal(r1.NumberOfTrades, r2.NumberOfTrades, "NumberOfTrades")
	r.Equal(r1.TakerBuyBaseAssetVolume, r2.TakerBuyBaseAssetVolume, "TakerBuyBaseAssetVolume")
	r.Equal(r1.TakerBuyQuoteAssetVolume, r2.TakerBuyQuoteAssetVolume, "TakerBuyQuoteAssetVolume")
}

func (s *apiMarketTestSuite) TestMarketUIKline() {
	msg := []byte(`{
  "id": "1b668cdc-72d7-43e4-952b-00bd50702f6b",
  "status": 200,
  "result": [
    [
      1737556140000,
      "104552.00000000",
      "104650.25000000",
      "104551.87000000",
      "104650.25000000",
      "0.34488000",
      1737556199999,
      "36077.69853200",
      122,
      "0.22480000",
      "23513.66461330",
      "0"
    ],
    [
      1737556200000,
      "104634.38000000",
      "104644.84000000",
      "104596.68000000",
      "104605.60000000",
      "0.20685000",
      1737556259999,
      "21637.65114840",
      96,
      "0.10046000",
      "10508.39407400",
      "0"
    ]
  ],
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
	resp, err := s.client.NewUiKlines().Symbol("BTCUSDT").Interval(core.Interval1m).Limit(2).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := WsKlineResponse{
		ApiResponse: ApiResponse{
			Id:     "1b668cdc-72d7-43e4-952b-00bd50702f6b",
			Status: 200,
			RateLimits: []*ApiRateLimit{{
				RateLimitType: "REQUEST_WEIGHT",
				Interval:      "MINUTE",
				IntervalNum:   1,
				Limit:         6000,
				Count:         4,
			}},
		},
		Result: []*KlineResult{{
			OpenTime:                 1737556140000,
			OpenPrice:                toDecimal("104552.00000000"),
			HighPrice:                toDecimal("104650.25000000"),
			LowPrice:                 toDecimal("104551.87000000"),
			ClosePrice:               toDecimal("104650.25000000"),
			Volume:                   toDecimal("0.34488000"),
			CloseTime:                1737556199999,
			QuoteAssetVolume:         toDecimal("36077.69853200"),
			NumberOfTrades:           122,
			TakerBuyBaseAssetVolume:  toDecimal("0.22480000"),
			TakerBuyQuoteAssetVolume: toDecimal("23513.66461330"),
		}, {
			OpenTime:                 1737556200000,
			OpenPrice:                toDecimal("104634.38000000"),
			HighPrice:                toDecimal("104644.84000000"),
			LowPrice:                 toDecimal("104596.68000000"),
			ClosePrice:               toDecimal("104605.60000000"),
			Volume:                   toDecimal("0.20685000"),
			CloseTime:                1737556259999,
			QuoteAssetVolume:         toDecimal("21637.65114840"),
			NumberOfTrades:           96,
			TakerBuyBaseAssetVolume:  toDecimal("0.10046000"),
			TakerBuyQuoteAssetVolume: toDecimal("10508.39407400"),
		}},
	}
	s.assertTestMarketKline(resp, &testResp)
}

func (s *apiMarketTestSuite) TestMarketAveragePrice() {
	msg := []byte(`{
	  "id": "ddbfb65f-9ebf-42ec-8240-8f0f91de0867",
	  "status": 200,
	  "result": {
		"mins": 5,                    
		"price": "9.35751834",        
		"closeTime": 1694061154503   
	  },
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 6000,
		  "count": 2
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAveragePrice().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsAveragePriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketAveragePrice(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketAveragePrice(r1, r2 *WsAveragePriceResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.CloseTime, r2.Result.CloseTime, "CloseTime")
	r.Equal(r1.Result.Mins, r2.Result.Mins, "Mins")
	r.Equal(r1.Result.Price, r2.Result.Price, "Price")
}

func (s *apiMarketTestSuite) TestMarketTicker24h() {
	msg := []byte(`{
  "id": "757f1d3c-9fce-41f9-92ae-de3069d7e4eb",
  "status": 200,
  "result": [{
    "symbol": "BTCUSDT",
    "priceChange": "449.67000000",
    "priceChangePercent": "0.433",
    "weightedAvgPrice": "102788.40769560",
    "prevClosePrice": "103804.46000000",
    "lastPrice": "104238.03000000",
    "lastQty": "0.00253000",
    "bidPrice": "104238.03000000",
    "bidQty": "0.00064000",
    "askPrice": "104247.96000000",
    "askQty": "0.01120000",
    "openPrice": "103788.36000000",
    "highPrice": "109027.00000000",
    "lowPrice": "25518.00000000",
    "volume": "334.66449000",
    "quoteVolume": "34399630.03935860",
    "openTime": 1737470638068,
    "closeTime": 1737557038068,
    "firstId": 1118635,
    "lastId": 1252322,
    "count": 133688
  }],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 8
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTicker24h().Symbols([]string{"BTCUSDT"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTicker24hResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTicker24h(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketTicker24h(r1, r2 *WsTicker24hResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTicker24hResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarketTicker24hResult(r1, r2 *Ticker24hResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PriceChange, r2.PriceChange, "PriceChange")
	r.Equal(r1.PriceChangePercent, r2.PriceChangePercent, "PriceChangePercent")
	r.Equal(r1.WeightedAvgPrice, r2.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(r1.PrevClosePrice, r2.PrevClosePrice, "PrevClosePrice")
	r.Equal(r1.LastPrice, r2.LastPrice, "LastPrice")
	r.Equal(r1.LastQty, r2.LastQty, "LastQty")
	r.Equal(r1.BidPrice, r2.BidPrice, "BidPrice")
	r.Equal(r1.BidQty, r2.BidQty, "BidQty")
	r.Equal(r1.AskPrice, r2.AskPrice, "AskPrice")
	r.Equal(r1.AskQty, r2.AskQty, "AskQty")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.Volume, r2.Volume, "Volume")
	r.Equal(r1.QuoteVolume, r2.QuoteVolume, "QuoteVolume")
	r.Equal(r1.OpenTime, r2.OpenTime, "OpenTime")
	r.Equal(r1.CloseTime, r2.CloseTime, "CloseTime")
	r.Equal(r1.FirstId, r2.FirstId, "FirstId")
	r.Equal(r1.LastId, r2.LastId, "LastId")
	r.Equal(r1.Count, r2.Count, "Count")
}

func (s *apiMarketTestSuite) TestMarketTickerTradingDay() {
	msg := []byte(`{
  "id": "fe204947-27a3-4aad-812d-07a7f056ba28",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "priceChange": "345.13000000",
      "priceChangePercent": "0.332",
      "weightedAvgPrice": "102541.04012050",
      "openPrice": "103861.28000000",
      "highPrice": "109027.00000000",
      "lowPrice": "25518.00000000",
      "lastPrice": "104206.41000000",
      "volume": "336.37872000",
      "quoteVolume": "34492623.82320220",
      "openTime": 1737470940000,
      "closeTime": 1737557359656,
      "firstId": 1119314,
      "lastId": 1252823,
      "count": 133510
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
	resp, err := s.client.NewTickerTradingDay().Symbols([]string{"BTCUSDT"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTickerTradingDayResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTickerTradingDay(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketTickerTradingDay(r1, r2 *WsTickerTradingDayResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTickerResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarketTickerResult(r1, r2 *TickerResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PriceChange, r2.PriceChange, "PriceChange")
	r.Equal(r1.PriceChangePercent, r2.PriceChangePercent, "PriceChangePercent")
	r.Equal(r1.WeightedAvgPrice, r2.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.LastPrice, r2.LastPrice, "LastPrice")
	r.Equal(r1.Volume, r2.Volume, "Volume")
	r.Equal(r1.QuoteVolume, r2.QuoteVolume, "QuoteVolume")
	r.Equal(r1.OpenTime, r2.OpenTime, "OpenTime")
	r.Equal(r1.CloseTime, r2.CloseTime, "CloseTime")
	r.Equal(r1.FirstId, r2.FirstId, "FirstId")
	r.Equal(r1.LastId, r2.LastId, "LastId")
	r.Equal(r1.Count, r2.Count, "Count")
}

func (s *apiMarketTestSuite) TestMarketTickerPrice() {
	msg := []byte(`{"id":"2ab37d21-ae20-4852-b8ef-a6e5a3d1e37e","status":200,"result":[{"symbol":"BTCUSDT","price":"104091.72000000"}],"rateLimits":[{"rateLimitType":"REQUEST_WEIGHT","interval":"MINUTE","intervalNum":1,"limit":6000,"count":4}]}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTickerPrice().Symbols([]string{"BTCUSDT"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTickerPriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTickerPrice(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketTickerPrice(r1, r2 *WsTickerPriceResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTickerPriceResult(r1.Result[i], r2.Result[i])
	}
}
func (s *apiMarketTestSuite) assertTestMarketTickerPriceResult(r1, r2 *TickerPriceResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Price, r2.Price, "Price")
}

func (s *apiMarketTestSuite) TestMarketTickerBook() {
	msg := []byte(`{
  "id": "3477cc2d-583b-42f4-817b-fc27622408c5",
  "status": 200,
  "result": [
    {
      "symbol": "BTCUSDT",
      "bidPrice": "104036.97000000",
      "bidQty": "0.01000000",
      "askPrice": "104050.08000000",
      "askQty": "0.01120000"
    },
    {
      "symbol": "ETHUSDT",
      "bidPrice": "3290.40000000",
      "bidQty": "4.37770000",
      "askPrice": "3290.41000000",
      "askQty": "6.03030000"
    }
  ],
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000,
      "count": 12
    }
  ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTickerBook().Symbols([]string{"BTCUSDT"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsTickerBookResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestMarketTickerBook(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestMarketTickerBook(r1, r2 *WsTickerBookResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertTestMarketTickerBookResult(r1.Result[i], r2.Result[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarketTickerBookResult(r1, r2 *TickerBookResult) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.BidPrice, r2.BidPrice, "BidPrice")
	r.Equal(r1.BidQty, r2.BidQty, "BidQty")
	r.Equal(r1.AskPrice, r2.AskPrice, "AskPrice")
	r.Equal(r1.AskQty, r2.AskQty, "AskQty")
}
