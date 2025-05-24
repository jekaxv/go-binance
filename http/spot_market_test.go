package http

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiMarketTestSuite struct {
	baseTestSuite
}

func TestApiMarketAccount(t *testing.T) {
	suite.Run(t, new(apiMarketTestSuite))
}

func (s *apiMarketTestSuite) TestNewDepth() {
	msg := []byte(`{
  "lastUpdateId": 1027024,
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
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewDepth().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *DepthResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestDepthResponse(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestDepthResponse(r1, r2 *DepthResponse) {
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

func (s *apiMarketTestSuite) TestNewTrades() {
	msg := []byte(`[
  {
    "id": 28457,
    "price": "4.00000100",
    "qty": "12.00000000",
    "quoteQty": "48.000012",
    "time": 1499865549590,
    "isBuyerMaker": true,
    "isBestMatch": true
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTrades().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TradesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTradesResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestTradesResponse(r1, r2 *TradesResponse) {
	r := s.r()
	r.Equal(r1.Id, r2.Id, "Id")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Qty, r2.Qty, "Qty")
	r.Equal(r1.QuoteQty, r2.QuoteQty, "QuoteQty")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.IsBuyerMaker, r2.IsBuyerMaker, "IsBuyerMaker")
	r.Equal(r1.IsBestMatch, r2.IsBestMatch, "IsBestMatch")
}

func (s *apiMarketTestSuite) TestNewHistoricalTrades() {
	msg := []byte(`[
  {
    "id": 28457,
    "price": "4.00000100",
    "qty": "12.00000000",
    "quoteQty": "48.000012",
    "time": 1499865549590,
    "isBuyerMaker": true,
    "isBestMatch": true
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewHistoricalTrades().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*HistoricalTradesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestHistoricalTradesResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestHistoricalTradesResponse(r1, r2 *HistoricalTradesResponse) {
	r := s.r()
	r.Equal(r1.Id, r2.Id, "Id")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Qty, r2.Qty, "Qty")
	r.Equal(r1.QuoteQty, r2.QuoteQty, "QuoteQty")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.IsBuyerMaker, r2.IsBuyerMaker, "IsBuyerMaker")
	r.Equal(r1.IsBestMatch, r2.IsBestMatch, "IsBestMatch")
}

func (s *apiMarketTestSuite) TestNewAggregateTrades() {
	msg := []byte(`[
  {
    "a": 26129,
    "p": "0.01633102",
    "q": "4.70443515",
    "f": 27781,
    "l": 27781,
    "T": 1498793709153,
    "m": true,
    "M": true
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAggregateTrades().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*AggregateTradesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestAggregateTradesResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestAggregateTradesResponse(r1, r2 *AggregateTradesResponse) {
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

func (s *apiMarketTestSuite) TestNewKlineData() {
	msg := []byte(`[
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
  ]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewKlineData().Symbol("BTCUSDT").Interval(types.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
		OpenTime:                 1737556140000,
		OpenPrice:                "104552.00000000",
		HighPrice:                "104650.25000000",
		LowPrice:                 "104551.87000000",
		ClosePrice:               "104650.25000000",
		Volume:                   "0.34488000",
		CloseTime:                1737556199999,
		QuoteAssetVolume:         "36077.69853200",
		NumberOfTrades:           122,
		TakerBuyBaseAssetVolume:  "0.22480000",
		TakerBuyQuoteAssetVolume: "23513.66461330",
	}, {
		OpenTime:                 1737556200000,
		OpenPrice:                "104634.38000000",
		HighPrice:                "104644.84000000",
		LowPrice:                 "104596.68000000",
		ClosePrice:               "104605.60000000",
		Volume:                   "0.20685000",
		CloseTime:                1737556259999,
		QuoteAssetVolume:         "21637.65114840",
		NumberOfTrades:           96,
		TakerBuyBaseAssetVolume:  "0.10046000",
		TakerBuyQuoteAssetVolume: "10508.39407400",
	}}
	for i := range resp {
		s.assertTestKlineDataResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestKlineDataResponse(r1, r2 *KlineDataResponse) {
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

func (s *apiMarketTestSuite) TestNewUIKlines() {
	msg := []byte(`[
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
  ]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewUIKlines().Symbol("BTCUSDT").Interval(types.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
		OpenTime:                 1737556140000,
		OpenPrice:                "104552.00000000",
		HighPrice:                "104650.25000000",
		LowPrice:                 "104551.87000000",
		ClosePrice:               "104650.25000000",
		Volume:                   "0.34488000",
		CloseTime:                1737556199999,
		QuoteAssetVolume:         "36077.69853200",
		NumberOfTrades:           122,
		TakerBuyBaseAssetVolume:  "0.22480000",
		TakerBuyQuoteAssetVolume: "23513.66461330",
	}, {
		OpenTime:                 1737556200000,
		OpenPrice:                "104634.38000000",
		HighPrice:                "104644.84000000",
		LowPrice:                 "104596.68000000",
		ClosePrice:               "104605.60000000",
		Volume:                   "0.20685000",
		CloseTime:                1737556259999,
		QuoteAssetVolume:         "21637.65114840",
		NumberOfTrades:           96,
		TakerBuyBaseAssetVolume:  "0.10046000",
		TakerBuyQuoteAssetVolume: "10508.39407400",
	}}
	for i := range resp {
		s.assertTestKlineDataResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewAveragePrice() {
	msg := []byte(`{
  "mins": 5,
  "price": "9.35751834",
  "closeTime": 1694061154503
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAveragePrice().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AveragePriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAveragePriceResponse(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestAveragePriceResponse(r1, r2 *AveragePriceResponse) {
	r := s.r()
	r.Equal(r1.CloseTime, r2.CloseTime, "CloseTime")
	r.Equal(r1.Mins, r2.Mins, "Mins")
	r.Equal(r1.Price, r2.Price, "Price")
}

func (s *apiMarketTestSuite) TestNewTickerPrice24h() {
	msg := []byte(`[{
  "symbol": "BNBBTC",
  "priceChange": "-94.99999800",
  "priceChangePercent": "-95.960",
  "weightedAvgPrice": "0.29628482",
  "prevClosePrice": "0.10002000",
  "lastPrice": "4.00000200",
  "lastQty": "200.00000000",
  "bidPrice": "4.00000000",
  "bidQty": "100.00000000",
  "askPrice": "4.00000200",
  "askQty": "100.00000000",
  "openPrice": "99.00000000",
  "highPrice": "100.00000000",
  "lowPrice": "0.10000000",
  "volume": "8913.30000000",
  "quoteVolume": "15.30000000",
  "openTime": 1499783499040,
  "closeTime": 1499869899040,
  "firstId": 28385,
  "lastId": 28460,
  "count": 76
}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTickerPrice24h().Symbols([]string{"BTCUSDT"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TickerPrice24hResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTickerPrice24hResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestTickerPrice24hResponse(r1, r2 *TickerPrice24hResponse) {
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

func (s *apiMarketTestSuite) TestNewTradingDayTicker() {
	msg := []byte(`[
  {
    "symbol": "BTCUSDT",
    "priceChange": "-83.13000000",
    "priceChangePercent": "-0.317",
    "weightedAvgPrice": "26234.58803036",
    "openPrice": "26304.80000000",
    "highPrice": "26397.46000000",
    "lowPrice": "26088.34000000",
    "lastPrice": "26221.67000000",
    "volume": "18495.35066000",
    "quoteVolume": "485217905.04210480",
    "openTime": 1695686400000,
    "closeTime": 1695772799999,
    "firstId": 3220151555,
    "lastId": 3220849281,
    "count": 697727
  },
  {
    "symbol": "BNBUSDT",
    "priceChange": "2.60000000",
    "priceChangePercent": "1.238",
    "weightedAvgPrice": "211.92276958",
    "openPrice": "210.00000000",
    "highPrice": "213.70000000",
    "lowPrice": "209.70000000",
    "lastPrice": "212.60000000",
    "volume": "280709.58900000",
    "quoteVolume": "59488753.54750000",
    "openTime": 1695686400000,
    "closeTime": 1695772799999,
    "firstId": 672397461,
    "lastId": 672496158,
    "count": 98698
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTradingDayTicker().Symbols([]string{"BTCUSDT", "BNBUSDT"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TickerResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTickerResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestTickerResponse(r1, r2 *TickerResponse) {
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

func (s *apiMarketTestSuite) TestNewPriceTicker() {
	msg := []byte(`[
  {
    "symbol": "LTCBTC",
    "price": "4.00000200"
  },
  {
    "symbol": "ETHBTC",
    "price": "0.07946600"
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewPriceTicker().Symbols([]string{"LTCBTC", "ETHBTC"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*PriceTickerResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestPriceTickerResponse(resp[i], testResp[i])
	}
}
func (s *apiMarketTestSuite) assertTestPriceTickerResponse(r1, r2 *PriceTickerResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Price, r2.Price, "Price")
}

func (s *apiMarketTestSuite) TestNewOrderBookTicker() {
	msg := []byte(`[
  {
    "symbol": "LTCBTC",
    "bidPrice": "4.00000000",
    "bidQty": "431.00000000",
    "askPrice": "4.00000200",
    "askQty": "9.00000000"
  },
  {
    "symbol": "ETHBTC",
    "bidPrice": "0.07946700",
    "bidQty": "9.00000000",
    "askPrice": "100000.00000000",
    "askQty": "1000.00000000"
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOrderBookTicker().Symbols([]string{"LTCBTC", "ETHBTC"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OrderBookTickerResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestOrderBookTickerResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestOrderBookTickerResponse(r1, r2 *OrderBookTickerResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.BidPrice, r2.BidPrice, "BidPrice")
	r.Equal(r1.BidQty, r2.BidQty, "BidQty")
	r.Equal(r1.AskPrice, r2.AskPrice, "AskPrice")
	r.Equal(r1.AskQty, r2.AskQty, "AskQty")
}

func (s *apiMarketTestSuite) TestNewTicker() {
	msg := []byte(`[
  {
    "symbol": "BTCUSDT",
    "priceChange": "-154.13000000",
    "priceChangePercent": "-0.740",
    "weightedAvgPrice": "20677.46305250",
    "openPrice": "20825.27000000",
    "highPrice": "20972.46000000",
    "lowPrice": "20327.92000000",
    "lastPrice": "20671.14000000",
    "volume": "72.65112300",
    "quoteVolume": "1502240.91155513",
    "openTime": 1655432400000,
    "closeTime": 1655446835460,
    "firstId": 11147809,
    "lastId": 11149775,
    "count": 1967
  },
  {
    "symbol": "BNBBTC",
    "priceChange": "0.00008530",
    "priceChangePercent": "0.823",
    "weightedAvgPrice": "0.01043129",
    "openPrice": "0.01036170",
    "highPrice": "0.01049850",
    "lowPrice": "0.01033870",
    "lastPrice": "0.01044700",
    "volume": "166.67000000",
    "quoteVolume": "1.73858301",
    "openTime": 1655432400000,
    "closeTime": 1655446835460,
    "firstId": 2351674,
    "lastId": 2352034,
    "count": 361
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTicker().Symbols([]string{"BTCUSDT", "BNBBTC"}).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TickerResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTickerResponse(resp[i], testResp[i])
	}
}
