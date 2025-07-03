package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiMarketTestSuite struct {
	baseHttpTestSuite
}

func TestApiGeneralAccount(t *testing.T) {
	suite.Run(t, new(apiMarketTestSuite))
}

func (s *apiMarketTestSuite) TestNewPing() {
	msg := []byte(`{}`)
	server := s.setup(msg)
	defer server.Close()
	err := s.client.NewPing().Do(context.Background())
	r := s.r()
	r.Empty(err)
}

func (s *apiMarketTestSuite) TestNewServerTime() {
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

func (s *apiMarketTestSuite) TestNewExchangeInfo() {
	msg := []byte(`{
  "timezone": "UTC",
  "serverTime": 1748597014323,
  "futuresType": "U_MARGINED",
  "rateLimits": [
    {
      "rateLimitType": "REQUEST_WEIGHT",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 6000
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "MINUTE",
      "intervalNum": 1,
      "limit": 1200
    },
    {
      "rateLimitType": "ORDERS",
      "interval": "SECOND",
      "intervalNum": 10,
      "limit": 300
    }
  ],
  "exchangeFilters": [],
  "assets": [
    {
      "asset": "USDT",
      "marginAvailable": true,
      "autoAssetExchange": "-100"
    },
    {
      "asset": "BTC",
      "marginAvailable": true,
      "autoAssetExchange": "-0.00100000"
    },
    {
      "asset": "BNB",
      "marginAvailable": true,
      "autoAssetExchange": "-0.00100000"
    },
    {
      "asset": "ETH",
      "marginAvailable": true,
      "autoAssetExchange": "-0.00100000"
    },
    {
      "asset": "USDC",
      "marginAvailable": true,
      "autoAssetExchange": "-0.10000000"
    },
    {
      "asset": "FDUSD",
      "marginAvailable": true,
      "autoAssetExchange": "0"
    },
    {
      "asset": "BNFCR",
      "marginAvailable": true,
      "autoAssetExchange": "-5000"
    },
    {
      "asset": "BFUSD",
      "marginAvailable": true,
      "autoAssetExchange": "0"
    }
  ],
  "symbols": [
    {
      "symbol": "BTCUSDT",
      "pair": "BTCUSDT",
      "contractType": "PERPETUAL",
      "deliveryDate": 4133404802000,
      "onboardDate": 1569398400000,
      "status": "TRADING",
      "maintMarginPercent": "2.5000",
      "requiredMarginPercent": "5.0000",
      "baseAsset": "BTC",
      "quoteAsset": "USDT",
      "marginAsset": "USDT",
      "pricePrecision": 2,
      "quantityPrecision": 3,
      "baseAssetPrecision": 8,
      "quotePrecision": 8,
      "underlyingType": "COIN",
      "underlyingSubType": [],
      "triggerProtect": "0.0500",
      "liquidationFee": "0.020000",
      "marketTakeBound": "0.30",
      "maxMoveOrderLimit": 1000,
      "filters": [
        {
          "filterType": "PRICE_FILTER",
          "tickSize": "0.10",
          "minPrice": "261.10",
          "maxPrice": "809484"
        },
        {
          "maxQty": "1000",
          "stepSize": "0.001",
          "filterType": "LOT_SIZE",
          "minQty": "0.001"
        },
        {
          "filterType": "MARKET_LOT_SIZE",
          "minQty": "0.001",
          "stepSize": "0.001",
          "maxQty": "1000"
        },
        {
          "filterType": "MAX_NUM_ORDERS",
          "limit": 200
        },
        {
          "limit": 10,
          "filterType": "MAX_NUM_ALGO_ORDERS"
        },
        {
          "notional": "100",
          "filterType": "MIN_NOTIONAL"
        },
        {
          "multiplierDecimal": "4",
          "multiplierUp": "1.5000",
          "multiplierDown": "0.5000",
          "filterType": "PERCENT_PRICE"
        },
        {
          "positionControlSide": "NONE",
          "filterType": "POSITION_RISK_CONTROL"
        }
      ],
      "orderTypes": [
        "LIMIT",
        "MARKET",
        "STOP",
        "STOP_MARKET",
        "TAKE_PROFIT",
        "TAKE_PROFIT_MARKET",
        "TRAILING_STOP_MARKET"
      ],
      "timeInForce": [
        "GTC",
        "IOC",
        "FOK",
        "GTX",
        "GTD"
      ],
      "permissionSets": [
        "GRID",
        "COPY"
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

func (s *apiMarketTestSuite) assertTestExchangeInfoResponse(r1, r2 *ExchangeInfoResponse) {
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

func (s *apiMarketTestSuite) TestNewDepth() {
	msg := []byte(`{
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
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewDepth().Symbol("BTCUSDT").Limit(1).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *DepthResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestDepthResponse(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestDepthResponse(r1, r2 *DepthResponse) {
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

func (s *apiMarketTestSuite) TestNewTrades() {
	msg := []byte(`[
  {
    "id": 28457,
    "price": "4.00000100",
    "qty": "12.00000000",
    "quoteQty": "48.000012",
    "time": 1499865549590,
    "isBuyerMaker": true
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
}

func (s *apiMarketTestSuite) TestNewHistoricalTrades() {
	msg := []byte(`[
  {
    "id": 28457,
    "price": "4.00000100",
    "qty": "12.00000000",
    "quoteQty": "48.000012",
    "time": 1499865549590,
    "isBuyerMaker": true
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewHistoricalTrades().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TradesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTradesResponse(resp[i], testResp[i])
	}
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
    "m": true
  }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAggTrades().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*AggTradesResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestAggregateTradesResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestAggregateTradesResponse(r1, r2 *AggTradesResponse) {
	r := s.r()
	r.Equal(r1.TradeId, r2.TradeId, "TradeId")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Quantity, r2.Quantity, "Quantity")
	r.Equal(r1.FirstId, r2.FirstId, "FirstId")
	r.Equal(r1.LastId, r2.LastId, "LastId")
	r.Equal(r1.Timestamp, r2.Timestamp, "Timestamp")
	r.Equal(r1.IsMaker, r2.IsMaker, "IsMaker")
}

func toDecimal(v string) decimal.Decimal {
	float, _ := decimal.NewFromString(v)
	return float
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
	resp, err := s.client.NewKline().Symbol("BTCUSDT").Interval(core.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
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

func (s *apiMarketTestSuite) TestNewContractKline() {
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
	resp, err := s.client.NewContractKline().Pair("BTCUSDT").ContractType(core.ContractTypePERPETUAL).Interval(core.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
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
	}}
	for i := range resp {
		s.assertTestKlineDataResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewIndexKline() {
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
	resp, err := s.client.NewIndexKline().Pair("BTCUSDT").Interval(core.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
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
	}}
	for i := range resp {
		s.assertTestKlineDataResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewMarkKline() {
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
	resp, err := s.client.NewMarkKline().Symbol("BTCUSDT").Interval(core.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
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
	}}
	for i := range resp {
		s.assertTestKlineDataResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewPremiumKline() {
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
	resp, err := s.client.NewPremiumKline().Symbol("BTCUSDT").Interval(core.Interval1m).Do(context.Background())
	r := s.r()
	r.Empty(err)
	testResp := []*KlineDataResponse{{
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
	}}
	for i := range resp {
		s.assertTestKlineDataResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewMarkPrice() {
	msg := []byte(`[{
		"symbol": "GALAUSDT",
		"markPrice": "0.01647150",
		"indexPrice": "0.01648441",
		"estimatedSettlePrice": "0.01654109",
		"lastFundingRate": "0.00010000",
		"interestRate": "0.00010000",
		"nextFundingTime": 1748851200000,
		"time": 1748840543000
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewMarkPrice().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*MarkPriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestMarkPriceResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestMarkPriceResponse(r1, r2 *MarkPriceResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.MarkPrice, r2.MarkPrice, "markPrice")
	r.Equal(r1.IndexPrice, r2.IndexPrice, "indexPrice")
	r.Equal(r1.EstimatedSettlePrice, r2.EstimatedSettlePrice, "estimatedSettlePrice")
	r.Equal(r1.LastFundingRate, r2.LastFundingRate, "lastFundingRate")
	r.Equal(r1.InterestRate, r2.InterestRate, "interestRate")
	r.Equal(r1.NextFundingTime, r2.NextFundingTime, "nextFundingTime")
	r.Equal(r1.Time, r2.Time, "time")
}

func (s *apiMarketTestSuite) TestNewFundingRate() {
	msg := []byte(`[
	{
		"symbol": "BTCUSDT",
		"fundingRate": "0.00010000",
		"fundingTime": 1748822400000,
		"markPrice": "105588.30108696"
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewFundingRate().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*FundingRateResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestFundingRateResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestFundingRateResponse(r1, r2 *FundingRateResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.FundingRate, r2.FundingRate, "fundingRate")
	r.Equal(r1.FundingTime, r2.FundingTime, "fundingTime")
	r.Equal(r1.MarkPrice, r2.MarkPrice, "markPrice")
}

func (s *apiMarketTestSuite) TestNewFundingInfo() {
	msg := []byte(`[
    {
        "symbol": "BLZUSDT",
        "adjustedFundingRateCap": "0.02500000",
        "adjustedFundingRateFloor": "-0.02500000",
        "fundingIntervalHours": 8,
        "disclaimer": false 
    }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewFundingInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*FundingInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestFundingInfoResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestFundingInfoResponse(r1, r2 *FundingInfoResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "symbol")
	r.Equal(r1.AdjustedFundingRateCap, r2.AdjustedFundingRateCap, "adjustedFundingRateCap")
	r.Equal(r1.AdjustedFundingRateFloor, r2.AdjustedFundingRateFloor, "adjustedFundingRateFloor")
	r.Equal(r1.FundingIntervalHours, r2.FundingIntervalHours, "fundingIntervalHours")
	r.Equal(r1.Disclaimer, r2.Disclaimer, "disclaimer")
}

func (s *apiMarketTestSuite) TestNewTickerPrice24h() {
	msg := []byte(`[
	{
  		"symbol": "BTCUSDT",
  		"priceChange": "-94.99999800",
  		"priceChangePercent": "-95.960",
  		"weightedAvgPrice": "0.29628482",
  		"lastPrice": "4.00000200",
  		"lastQty": "200.00000000",
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
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTicker24hr().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TickerStatisticsResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTickerPrice24hResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestTickerPrice24hResponse(r1, r2 *TickerStatisticsResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PriceChange, r2.PriceChange, "PriceChange")
	r.Equal(r1.PriceChangePercent, r2.PriceChangePercent, "PriceChangePercent")
	r.Equal(r1.WeightedAvgPrice, r2.WeightedAvgPrice, "WeightedAvgPrice")
	r.Equal(r1.LastPrice, r2.LastPrice, "LastPrice")
	r.Equal(r1.LastQty, r2.LastQty, "LastQty")
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

func (s *apiMarketTestSuite) TestNewPriceTicker() {
	msg := []byte(`[
	{
  		"symbol": "BTCUSDT",
  		"price": "6000.01",
  		"time": 1589437530011
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTickerPrice().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TickerPriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestPriceTickerResponse(resp[i], testResp[i])
	}
}
func (s *apiMarketTestSuite) assertTestPriceTickerResponse(r1, r2 *TickerPriceResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Time, r2.Time, "Time")
}

func (s *apiMarketTestSuite) TestNewBookTicker() {
	msg := []byte(`[
	{
		"symbol": "BTCUSDT",
		"bidPrice": "105464.2",
		"bidQty": "0.146",
		"askPrice": "105464.4",
		"askQty": "0.047",
		"time": 1748851798422
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewBookTicker().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*BookTickerResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestBookTickerResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestBookTickerResponse(r1, r2 *BookTickerResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.AskPrice, r2.AskPrice, "AskPrice")
	r.Equal(r1.AskQty, r2.AskQty, "AskQty")
	r.Equal(r1.BidPrice, r2.BidPrice, "BidPrice")
	r.Equal(r1.BidQty, r2.BidQty, "BidQty")
	r.Equal(r1.Time, r2.Time, "Time")
}

func (s *apiMarketTestSuite) TestNewDeliveryPrice() {
	msg := []byte(`[
    {
        "deliveryTime": 1695945600000,
        "deliveryPrice": 27103.00000000
    }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewDeliveryPrice().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*DeliveryPriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestDeliveryPriceResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestDeliveryPriceResponse(r1, r2 *DeliveryPriceResponse) {
	r := s.r()
	r.Equal(r1.DeliveryTime, r2.DeliveryTime, "deliveryTime")
	r.Equal(r1.DeliveryPrice, r2.DeliveryPrice, "deliveryPrice")
}

func (s *apiMarketTestSuite) TestNewOpenInterest() {
	msg := []byte(`{
	"openInterest": "10659.509", 
	"symbol": "BTCUSDT",
	"time": 1589437530011 
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOpenInterest().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *OpenInterestResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.OpenInterest, testResp.OpenInterest, "openInterest")
	r.Equal(resp.Symbol, testResp.Symbol, "symbol")
	r.Equal(resp.Time, testResp.Time, "time")
}

func (s *apiMarketTestSuite) TestNewOpenInterestHist() {
	msg := []byte(`[{
		"symbol": "BTCUSDT",
		"sumOpenInterest": "83495.746",
		"sumOpenInterestValue": "8832354647.467663",
		"timestamp": 1748853000000
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOpenInterestHist().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*OpenInterestHistResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestOpenInterestHistResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestOpenInterestHistResponse(r1, r2 *OpenInterestHistResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.SumOpenInterest, r2.SumOpenInterest, "SumOpenInterest")
	r.Equal(r1.SumOpenInterestValue, r2.SumOpenInterestValue, "SumOpenInterestValue")
	r.Equal(r1.Timestamp, r2.Timestamp, "Timestamp")
}

func (s *apiMarketTestSuite) TestNewTopTraderPositionsRatio() {
	msg := []byte(`[{
		"symbol": "BTCUSDT",
		"longShortRatio": "1.5278",
		"longAccount": "0.6044",
		"shortAccount": "0.3956",
		"timestamp": 1748853300000
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTopTraderPositionsRatio().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TopTraderRatioResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTopTraderRatioResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestTopTraderRatioResponse(r1, r2 *TopTraderRatioResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.LongShortRatio, r2.LongShortRatio, "LongShortRatio")
	r.Equal(r1.LongAccount, r2.LongAccount, "LongAccount")
	r.Equal(r1.ShortAccount, r2.ShortAccount, "ShortAccount")
	r.Equal(r1.Timestamp, r2.Timestamp, "Timestamp")
}

func (s *apiMarketTestSuite) TestNewTopTraderAccountsRatio() {
	msg := []byte(`[{
		"symbol": "BTCUSDT",
		"longShortRatio": "1.5278",
		"longAccount": "0.6044",
		"shortAccount": "0.3956",
		"timestamp": 1748853300000
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTopTraderAccountsRatio().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TopTraderRatioResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTopTraderRatioResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewSymbolRatio() {
	msg := []byte(`[{
		"symbol": "BTCUSDT",
		"longShortRatio": "1.5278",
		"longAccount": "0.6044",
		"shortAccount": "0.3956",
		"timestamp": 1748853300000
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewSymbolRatio().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TopTraderRatioResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTopTraderRatioResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) TestNewTakerVolume() {
	msg := []byte(`[{
		"buySellRatio": "0.4287",
		"buyVol": "213.545",
		"sellVol": "498.08",
		"timestamp": 1748853600000
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTakerVolume().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*TakerVolumeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestTakerVolumeResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestTakerVolumeResponse(r1, r2 *TakerVolumeResponse) {
	r := s.r()
	r.Equal(r1.BuySellRatio, r2.BuySellRatio, "BuySellRatio")
	r.Equal(r1.BuyVol, r2.BuyVol, "BuyVol")
	r.Equal(r1.SellVol, r2.SellVol, "SellVol")
	r.Equal(r1.Timestamp, r2.Timestamp, "Timestamp")
}
func (s *apiMarketTestSuite) TestNewFutureBasis() {
	msg := []byte(`[  
    {
        "indexPrice": "34400.15945055",
        "contractType": "PERPETUAL",
        "basisRate": "0.0004",
        "futuresPrice": "34414.10",
        "annualizedBasisRate": "",
        "basis": "13.94054945",
        "pair": "BTCUSDT",
        "timestamp": 1698742800000
    }
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewFutureBasis().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*FutureBasisResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestFutureBasisResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestFutureBasisResponse(r1, r2 *FutureBasisResponse) {
	r := s.r()
	r.Equal(r1.IndexPrice, r2.IndexPrice, "IndexPrice")
	r.Equal(r1.ContractType, r2.ContractType, "ContractType")
	r.Equal(r1.BasisRate, r2.BasisRate, "BasisRate")
	r.Equal(r1.FuturesPrice, r2.FuturesPrice, "FuturesPrice")
	r.Equal(r1.AnnualizedBasisRate, r2.AnnualizedBasisRate, "AnnualizedBasisRate")
	r.Equal(r1.Basis, r2.Basis, "Basis")
	r.Equal(r1.Pair, r2.Pair, "Pair")
	r.Equal(r1.Timestamp, r2.Timestamp, "Timestamp")
}

func (s *apiMarketTestSuite) TestNewIndexInfo() {
	msg := []byte(`[
		{ 
			"symbol": "DEFIUSDT",
			"time": 1589437530011,    
			"component": "baseAsset",
			"baseAssetList":[
				{
					"baseAsset":"BAL",
					"quoteAsset": "USDT",
					"weightInQuantity":"1.04406228",
					"weightInPercentage":"0.02783900"
				},
				{
					"baseAsset":"BAND",
					"quoteAsset": "USDT",
					"weightInQuantity":"3.53782729",
					"weightInPercentage":"0.03935200"
				}
			]
		}
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewIndexInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*IndexInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestIndexInfoResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestBaseAsset(r1, r2 *BaseAsset) {
	r := s.r()
	r.Equal(r1.BaseAsset, r2.BaseAsset, "BaseAsset")
	r.Equal(r1.QuoteAsset, r2.QuoteAsset, "QuoteAsset")
	r.Equal(r1.WeightInQuantity, r2.WeightInQuantity, "WeightInQuantity")
	r.Equal(r1.WeightInPercentage, r2.WeightInPercentage, "WeightInPercentage")
}

func (s *apiMarketTestSuite) assertTestIndexInfoResponse(r1, r2 *IndexInfoResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Component, r2.Component, "Component")
	for i := range r1.BaseAssetList {
		s.assertTestBaseAsset(r1.BaseAssetList[i], r2.BaseAssetList[i])
	}
}

func (s *apiMarketTestSuite) TestNewAssetInfo() {
	msg := []byte(`[
	{
		"symbol": "BTCUSD",
		"time": 1748855537004,
		"index": "104956.05365614",
		"bidBuffer": "0.05",
		"askBuffer": "0.05",
		"bidRate": "99708.25097333",
		"askRate": "110203.85633894",
		"autoExchangeBidBuffer": "0.025",
		"autoExchangeAskBuffer": "0.025",
		"autoExchangeBidRate": "102332.15231473",
		"autoExchangeAskRate": "107579.95499754"
	}
]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAssetIndex().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*AssetIndexResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestAssetIndexResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestAssetIndexResponse(r1, r2 *AssetIndexResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Index, r2.Index, "Index")
	r.Equal(r1.BidBuffer, r2.BidBuffer, "BidBuffer")
	r.Equal(r1.AskBuffer, r2.AskBuffer, "AskBuffer")
	r.Equal(r1.BidRate, r2.BidRate, "BidRate")
	r.Equal(r1.AskRate, r2.AskRate, "AskRate")
	r.Equal(r1.AutoExchangeBidBuffer, r2.AutoExchangeBidBuffer, "AutoExchangeBidBuffer")
	r.Equal(r1.AutoExchangeAskBuffer, r2.AutoExchangeAskBuffer, "AutoExchangeAskBuffer")
	r.Equal(r1.AutoExchangeBidRate, r2.AutoExchangeBidRate, "AutoExchangeBidRate")
	r.Equal(r1.AutoExchangeAskRate, r2.AutoExchangeAskRate, "AutoExchangeAskRate")
}

func (s *apiMarketTestSuite) TestNewConstituentsPrice() {
	msg := []byte(`{
    "symbol": "BTCUSDT",
    "time": 1745401553408,
    "constituents": [
        {
            "exchange": "binance",
            "symbol": "BTCUSDT",
            "price": "94057.03000000",
            "weight": "0.51282051"
        },
        {
            "exchange": "coinbase",
            "symbol": "BTC-USDT",
            "price": "94140.58000000",
            "weight": "0.15384615"
        },
        {
            "exchange": "gateio",
            "symbol": "BTC_USDT",
            "price": "94060.10000000",
            "weight": "0.02564103"
        },
        {
            "exchange": "kucoin",
            "symbol": "BTC-USDT",
            "price": "94096.70000000",
            "weight": "0.07692308"
        },
        {
            "exchange": "mxc",
            "symbol": "BTCUSDT",
            "price": "94057.02000000",
            "weight": "0.07692308"
        },
        {
            "exchange": "bitget",
            "symbol": "BTCUSDT",
            "price": "94064.03000000",
            "weight": "0.07692308"
        },
        {
            "exchange": "bybit",
            "symbol": "BTCUSDT",
            "price": "94067.90000000",
            "weight": "0.07692308"
        }
    ]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewConstituentsPrice().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ConstituentsPriceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestConstituentsPriceResponse(resp, testResp)
}

func (s *apiMarketTestSuite) assertTestConstituent(r1, r2 *ConstituentResponse) {
	r := s.r()
	r.Equal(r1.Exchange, r2.Exchange, "Exchange")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Weight, r2.Weight, "Weight")
}

func (s *apiMarketTestSuite) assertTestConstituentsPriceResponse(r1, r2 *ConstituentsPriceResponse) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Time, r2.Time, "Time")
	for i := range r1.Constituents {
		s.assertTestConstituent(r1.Constituents[i], r2.Constituents[i])
	}
}

func (s *apiMarketTestSuite) TestNewInsuranceBalance() {
	msg := []byte(`[{
   "symbols":[
      "BNBUSDT",
      "BTCUSDT",
      "BTCUSDT_250627",
      "BTCUSDT_250926",
      "ETHBTC",
      "ETHUSDT",
      "ETHUSDT_250627",
      "ETHUSDT_250926"
   ],
   "assets":[
      {
         "asset":"USDC",
         "marginBalance":"299999998.6497832",
         "updateTime":1745366402000
      },
      {
         "asset":"USDT",
         "marginBalance":"793930579.315848",
         "updateTime":1745366402000
      },
      {
         "asset":"BTC",
         "marginBalance":"61.73143554",
         "updateTime":1745366402000
      },
      {
         "asset":"BNFCR",
         "marginBalance":"633223.99396922",
         "updateTime":1745366402000
      }
   ]
}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewInsuranceBalance().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*InsuranceBalanceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		s.assertTestInsuranceBalanceResponse(resp[i], testResp[i])
	}
}

func (s *apiMarketTestSuite) assertTestInsuranceBalanceResponse(r1, r2 *InsuranceBalanceResponse) {
	r := s.r()
	for i := range r1.Symbols {
		r.Equal(r1.Symbols[i], r2.Symbols[i], "Symbols")
	}
	for i := range r1.Assets {
		s.Equal(r1.Assets[i], r2.Assets[i], "asset")
		s.Equal(r1.Assets[i].MarginBalance, r2.Assets[i].MarginBalance, "marginBalance")
		s.Equal(r1.Assets[i].UpdateTime, r2.Assets[i].UpdateTime, "updateTime")
	}
}
