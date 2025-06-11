package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/stretchr/testify/suite"
	"testing"
)

type websocketStreamsTestSuite struct {
	baseWsTestSuite
}

func TestWebsocketStreams(t *testing.T) {
	suite.Run(t, new(websocketStreamsTestSuite))
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeAggTrade() {
	msg := []byte(`{"e":"aggTrade","E":1737443769749,"s":"BTCUSDT","a":1019485,"p":"102342.24000000","q":"0.00254000","f":1071934,"l":1071934,"T":1737443769749,"m":false,"M":true}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeAggTrade("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *AggTradeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeAggTrade(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) assertTestSubscribeAggTrade(r1, r2 *AggTradeEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.AggTradeID, r2.AggTradeID, "AggTradeID")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Quantity, r2.Quantity, "Quantity")
	r.Equal(r1.FirstTradeID, r2.FirstTradeID, "FirstTradeID")
	r.Equal(r1.LastTradeID, r2.LastTradeID, "LastTradeID")
	r.Equal(r1.TradeTime, r2.TradeTime, "TradeTime")
	r.Equal(r1.IsBuyerMaker, r2.IsBuyerMaker, "IsBuyerMaker")
	r.Equal(r1.Placeholder, r2.Placeholder, "Placeholder")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedAggTrade() {
	msg := []byte(`{"stream":"btcusdt@aggTrade","data":{"e":"aggTrade","E":1737444959539,"s":"BTCUSDT","a":1022040,"p":"102433.97000000","q":"0.00016000","f":1074598,"l":1074598,"T":1737444959539,"m":true,"M":true}}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedAggTrade([]string{"btcusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedAggTradeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedAggTrade(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedAggTrade(r1, r2 *CombinedAggTradeEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeAggTrade(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeMarkPrice() {
	msg := []byte(`{
		"e": "markPriceUpdate",
		"E": 1749131782001,
		"s": "BTCUSDT",
		"p": "104352.4",
		"i": "104382.33456522",
		"P": "105274.63385347",
		"r": "0.00007303",
		"T": 1749139200000
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeMarkPrice("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *MarkPriceEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeMarkPrice(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) assertTestSubscribeMarkPrice(r1, r2 *MarkPriceEvent) {
	r := s.r()
	r.Equal(r1.EventType, r2.EventType, "EventType")
	r.Equal(r1.EventTime, r2.EventTime, "EventTime")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.MarkPrice, r2.MarkPrice, "MarkPrice")
	r.Equal(r1.IndexPrice, r2.IndexPrice, "IndexPrice")
	r.Equal(r1.SettlePrice, r2.SettlePrice, "SettlePrice")
	r.Equal(r1.FundingRate, r2.FundingRate, "FundingRate")
	r.Equal(r1.NextFundingTime, r2.NextFundingTime, "NextFundingTime")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedMarkPrice() {
	msg := []byte(`{
		"stream": "btcusdt@aggTrade",
		"data": {
			"e": "aggTrade",
			"E": 1749132017297,
			"s": "BTCUSDT",
			"p": "104415.9",
			"i": "0",
			"P": "0",
			"r": "0",
			"T": 1749132017289
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().
		SubscribeCombinedMarkPrice([]string{"btcusdt"}).
		Do(context.Background())
	r := s.r()
	var testResp *CombinedMarkPriceEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedMarkPrice(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedMarkPrice(r1, r2 *CombinedMarkPriceEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeMarkPrice(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeSubscribeMarkPriceArr() {
	msg := []byte(`[{
		"e": "markPriceUpdate",
		"E": 1749132112001,
		"s": "MATICUSDC",
		"p": "0.37910467",
		"i": "0.3791",
		"P": "0.3791",
		"r": "0.0001",
		"T": 1749139200000
	}]`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeMarkPriceArr().Do(context.Background())
	r := s.r()
	var testResp []*MarkPriceEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			for i := range testResp {
				s.assertTestSubscribeMarkPrice(event[i], testResp[i])
			}
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestSubscribeKline() {
	msg := []byte(`{
		"e": "kline",
		"E": 1749132381466,
		"s": "BTCUSDT",
		"k": {
			"t": 1749132360000,
			"T": 1749132419999,
			"s": "BTCUSDT",
			"i": "1m",
			"f": 6371666483,
			"L": 6371668775,
			"o": "104341.5",
			"c": "104293.8",
			"h": "104341.5",
			"l": "104284.4",
			"v": "129.729",
			"n": 2292,
			"x": false,
			"q": "13532464.6015",
			"V": "66.792",
			"Q": "6967306.7488",
			"B": "0"
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeKline("btcusdt", "1m").Do(context.Background())
	r := s.r()
	var testResp *KlineEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeKline(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestSubscribeCombinedKline() {
	msg := []byte(`{
	  "stream": "btcusdt@kline_1m",
	  "data": {
		"e": "kline",
		"E": 1737447340064,
		"s": "BTCUSDT",
		"k": {
		  "t": 1737447300000,
		  "T": 1737447359999,
		  "s": "BTCUSDT",
		  "i": "1m",
		  "f": 1078720,
		  "L": 1078780,
		  "o": "101796.98000000",
		  "c": "101817.78000000",
		  "h": "101818.00000000",
		  "l": "101796.97000000",
		  "v": "0.05522000",
		  "n": 61,
		  "x": false,
		  "q": "5621.98160280",
		  "V": "0.04785000",
		  "Q": "4871.61906530",
		  "B": "0"
		}
	  }
	}`)
	server := s.setup(msg)
	defer server.Close()
	symbols := make(map[string]string)
	symbols["btcusdt"] = "1m"
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedKline(symbols).Do(context.Background())
	r := s.r()
	var testResp *CombinedKlineEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedKline(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedKline(r1, r2 *CombinedKlineEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeKline(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeKline(r1, r2 *KlineEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	s.assertTestSubscribeKlineResult(r1.Kline, r2.Kline)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeKlineResult(r1, r2 *KlineResult) {
	r := s.r()
	r.Equal(r1.StartTime, r2.StartTime, "Kline.StartTime")
	r.Equal(r1.CloseTime, r2.CloseTime, "Kline.CloseTime")
	r.Equal(r1.Symbol, r2.Symbol, "Kline.Symbol")
	r.Equal(r1.Interval, r2.Interval, "Kline.Interval")
	r.Equal(r1.FirstTradeId, r2.FirstTradeId, "Kline.FirstTradeId")
	r.Equal(r1.LastTradeId, r2.LastTradeId, "Kline.LastTradeId")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "Kline.OpenPrice")
	r.Equal(r1.ClosePrice, r2.ClosePrice, "Kline.ClosePrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "Kline.HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "Kline.LowPrice")
	r.Equal(r1.BaseAssetVolume, r2.BaseAssetVolume, "Kline.BaseAssetVolume")
	r.Equal(r1.NumberOfTrades, r2.NumberOfTrades, "Kline.NumberOfTrades")
	r.Equal(r1.IsClosed, r2.IsClosed, "Kline.IsClosed")
	r.Equal(r1.QuoteAssetVolume, r2.QuoteAssetVolume, "Kline.QuoteAssetVolume")
	r.Equal(r1.TakerBaseVolume, r2.TakerBaseVolume, "Kline.TakerBaseVolume")
	r.Equal(r1.TakerQuoteVolume, r2.TakerQuoteVolume, "Kline.TakerQuoteVolume")
	r.Equal(r1.Placeholder, r2.Placeholder, "Kline.Placeholder")
}

func (s *websocketStreamsTestSuite) TestSubscribeContractKline() {
	msg := []byte(`{
		"e": "continuous_kline",
		"E": 1749133211883,
		"ps": "BTCUSDT",
		"ct": "PERPETUAL",
		"k": {
			"t": 1749133200000,
			"T": 1749133259999,
			"s": "",
			"i": "1m",
			"f": 7715339655404,
			"L": 7715343081442,
			"o": "103958",
			"c": "103933.8",
			"h": "103963.9",
			"l": "103890.7",
			"v": "110.028",
			"n": 2038,
			"x": false,
			"q": "11434915.6191",
			"V": "46.929",
			"Q": "4877534.8862",
			"B": "0"
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	symbols := make(map[string]string)
	symbols["btcusdt"] = "1m"
	onMessage, onError := s.client.NewWebsocketStreams().
		SubscribeContractKline("BTCUSDT", core.ContractTypePERPETUAL, core.Interval1m).
		Do(context.Background())
	r := s.r()
	var testResp *ContractKlineEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeContractKline(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeContractKline(r1, r2 *ContractKlineEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.ContractType, r2.ContractType, "ContractType")
	s.assertTestSubscribeKlineResult(r1.Kline, r2.Kline)
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeMiniTicker() {
	msg := []byte(`{"e":"24hrMiniTicker","E":1737450828073,"s":"BTCUSDT","c":"102681.77000000","o":"108240.00000000","h":"108751.78000000","l":"25368.00000000","v":"234.17280000","q":"22492957.53468620"}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeMiniTicker("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *MiniTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeMiniTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedMiniTicker() {
	msg := []byte(`{"stream":"ethusdt@miniTicker","data":{"e":"24hrMiniTicker","E":1737509415381,"s":"ETHUSDT","c":"3339.23000000","o":"3227.51000000","h":"4232.04000000","l":"1196.62000000","v":"5617.01730000","q":"18532488.10400400"}}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedMiniTicker([]string{"ethusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedMiniTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedMiniTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedMiniTicker(r1, r2 *CombinedMiniTickerEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeMiniTicker(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeMiniTicker(r1, r2 *MiniTickerEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.ClosePrice, r2.ClosePrice, "ClosePrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.TotalTradedBaseVolume, r2.TotalTradedBaseVolume, "TotalTradedBaseVolume")
	r.Equal(r1.TotalTradedQuoteVolume, r2.TotalTradedQuoteVolume, "TotalTradedQuoteVolume")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeTicker() {
	msg := []byte(`{
		"e": "24hrTicker",
		"E": 1749134076288,
		"s": "BTCUSDT",
		"p": "-560.1",
		"P": "-0.535",
		"w": "104882.55",
		"c": "104162",
		"Q": "0.014",
		"o": "104722.1",
		"h": "105878.3",
		"l": "103825",
		"v": "127047.822",
		"q": "13325099636.94",
		"O": 1749047640000,
		"C": 1749134076283,
		"F": 6369496440,
		"L": 6371843306,
		"n": 2346522
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeTicker("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *TickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedTicker() {
	msg := []byte(`{
		"stream": "btcusdt@ticker",
		"data": {
			"e": "24hrTicker",
			"E": 1749134076288,
			"s": "BTCUSDT",
			"p": "-560.1",
			"P": "-0.535",
			"w": "104882.55",
			"c": "104162",
			"Q": "0.014",
			"o": "104722.1",
			"h": "105878.3",
			"l": "103825",
			"v": "127047.822",
			"q": "13325099636.94",
			"O": 1749047640000,
			"C": 1749134076283,
			"F": 6369496440,
			"L": 6371843306,
			"n": 2346522
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedTicker([]string{"btcusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedTicker(r1, r2 *CombinedTickerEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeTicker(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeTicker(r1, r2 *TickerEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PriceChange, r2.PriceChange, "PriceChange")
	r.Equal(r1.PriceChangePercent, r2.PriceChangePercent, "PriceChangePercent")
	r.Equal(r1.WeightedAveragePrice, r2.WeightedAveragePrice, "WeightedAveragePrice")
	r.Equal(r1.LastTradePrice, r2.LastTradePrice, "LastTradePrice")
	r.Equal(r1.LastQuantity, r2.LastQuantity, "LastQuantity")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.TotalTradedBaseVolume, r2.TotalTradedBaseVolume, "TotalTradedBaseVolume")
	r.Equal(r1.TotalTradedQuoteVolume, r2.TotalTradedQuoteVolume, "TotalTradedQuoteVolume")
	r.Equal(r1.OpenTime, r2.OpenTime, "OpenTime")
	r.Equal(r1.CloseTime, r2.CloseTime, "CloseTime")
	r.Equal(r1.FirstTradeID, r2.FirstTradeID, "FirstTradeID")
	r.Equal(r1.LastTradeID, r2.LastTradeID, "LastTradeID")
	r.Equal(r1.NumberOfTrades, r2.NumberOfTrades, "NumberOfTrades")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeTickerArr() {
	msg := []byte(`[{
		"e": "24hrTicker",
		"E": 1749134211104,
		"s": "ENAUSDT",
		"p": "-0.0303",
		"P": "-9.039",
		"w": "0.3224141",
		"c": "0.3049",
		"Q": "120",
		"o": "0.3352",
		"h": "0.3556",
		"l": "0.3017",
		"v": "1415830889",
		"q": "456483880.2727",
		"O": 1749047760000,
		"C": 1749134211100,
		"F": 475682873,
		"L": 476533606,
		"n": 850655
	}]`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeTickerArr().Do(context.Background())
	r := s.r()
	var testResp []*TickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			for i := range event {
				s.assertTestSubscribeTicker(event[i], testResp[i])
			}
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeMiniTickerArr() {
	msg := []byte(`[{
  "e": "24hrMiniTicker",
  "E": 1737509629525,
  "s": "XRPUSDT",
  "c": "3.19030000",
  "o": "3.11880000",
  "h": "4.00000000",
  "l": "1.10220000",
  "v": "2724144.00000000",
  "q": "8518048.14740000"
},{
  "e": "24hrMiniTicker",
  "E": 1737509629729,
  "s": "EOSUSDT",
  "c": "0.84530000",
  "o": "0.80160000",
  "h": "0.86450000",
  "l": "0.79410000",
  "v": "3785797.40000000",
  "q": "3133703.13657000"
}]`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeMiniTickerArr().Do(context.Background())
	r := s.r()
	var testResp []*MiniTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			for i := range event {
				s.assertTestSubscribeMiniTickerArr(event[i], testResp[i])
			}
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeMiniTickerArr(r1, r2 *MiniTickerEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.ClosePrice, r2.ClosePrice, "ClosePrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.TotalTradedBaseVolume, r2.TotalTradedBaseVolume, "TotalTradedBaseVolume")
	r.Equal(r1.TotalTradedQuoteVolume, r2.TotalTradedQuoteVolume, "TotalTradedQuoteVolume")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeBookTicker() {
	msg := []byte(`{
		"e": "bookTicker",
		"E": 1749134621774,
		"T": 1749134621774,
		"u": 7715549794604,
		"s": "BTCUSDT",
		"b": "104516.3",
		"B": "25.066",
		"a": "104516.4",
		"A": "1.904"
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeBookTicker("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *BookTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeBookTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedBookTicker() {
	msg := []byte(`{
		"stream": "btcusdt@bookTicker",
		"data": {
			"e": "bookTicker",
			"E": 1749134675289,
			"T": 1749134675289,
			"u": 7715556457811,
			"s": "BTCUSDT",
			"b": "104540.4",
			"B": "6.538",
			"a": "104540.5",
			"A": "5.862"
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedBookTicker([]string{"btcusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedBookTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedBookTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedBookTicker(r1, r2 *CombinedBookTickerEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeBookTicker(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeBookTicker(r1, r2 *BookTickerEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.EventTime, r2.EventTime, "EventTime")
	r.Equal(r1.TransactionTime, r2.TransactionTime, "TransactionTime")
	r.Equal(r1.UpdateId, r2.UpdateId, "UpdateId")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.BestBidPrice, r2.BestBidPrice, "BestBidPrice")
	r.Equal(r1.BestBidQty, r2.BestBidQty, "BestBidQty")
	r.Equal(r1.BestAskPrice, r2.BestAskPrice, "BestAskPrice")
	r.Equal(r1.BestAskQty, r2.BestAskQty, "BestAskQty")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeBestBookTicker() {
	msg := []byte(`{
		"e": "bookTicker",
		"E": 1749134952180,
		"T": 1749134952180,
		"u": 7715604306318,
		"s": "SUIUSDC",
		"b": "3.16036",
		"B": "1.7",
		"a": "3.16037",
		"A": "183"
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeBestBookTicker().Do(context.Background())
	r := s.r()
	var testResp *BookTickerEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeBookTicker(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) TestWebSocketSubscribeForceOrder() {
	msg := []byte(`{
		"e": "forceOrder",
		"E": 1749135788134,
		"o": {
			"s": "LPTUSDT",
			"S": "BUY",
			"o": "LIMIT",
			"f": "IOC",
			"q": "89.9",
			"p": "8.567",
			"ap": "8.49",
			"X": "FILLED",
			"l": "89.9",
			"z": "89.9",
			"T": 1749135788130
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeForceOrder("LPTUSDT").Do(context.Background())
	r := s.r()
	var testResp *ForceOrderEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeForceOrder(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeForceOrder(r1, r2 *ForceOrderEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.O.Symbol, r2.O.Symbol, "Symbol")
	r.Equal(r1.O.Side, r2.O.Side, "Side")
	r.Equal(r1.O.OrderType, r2.O.OrderType, "OrderType")
	r.Equal(r1.O.TimeInForce, r2.O.TimeInForce, "TimeInForce")
	r.Equal(r1.O.Quantity, r2.O.Quantity, "Quantity")
	r.Equal(r1.O.Price, r2.O.Price, "Price")
	r.Equal(r1.O.AveragePrice, r2.O.AveragePrice, "AveragePrice")
	r.Equal(r1.O.Status, r2.O.Status, "Status")
	r.Equal(r1.O.FilledQuantity, r2.O.FilledQuantity, "FilledQuantity")
	r.Equal(r1.O.AccumulatedQuantity, r2.O.AccumulatedQuantity, "AccumulatedQuantity")
	r.Equal(r1.O.TradeTime, r2.O.TradeTime, "TradeTime")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedForceOrder() {
	msg := []byte(`{
		"stream": "lptusdt@forceOrder",
		"data": {
			"e": "forceOrder",
			"E": 1749136143079,
			"o": {
				"s": "LPTUSDT",
				"S": "SELL",
				"o": "LIMIT",
				"f": "IOC",
				"q": "3.5",
				"p": "8.387",
				"ap": "8.465",
				"X": "FILLED",
				"l": "3.5",
				"z": "3.5",
				"T": 1749136143074
			}
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().
		SubscribeCombinedForceOrder([]string{"LPTUSDT"}).
		Do(context.Background())
	r := s.r()
	var testResp *CombinedForceOrderEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedForce(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedForce(r1, r2 *CombinedForceOrderEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeForceOrder(r1.Data, r2.Data)
}
