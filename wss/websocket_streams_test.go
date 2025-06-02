package wss

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type websocketStreamsTestSuite struct {
	baseTestSuite
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

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeTrade() {
	msg := []byte(`{"e":"trade","E":1737445349276,"s":"BTCUSDT","t":1075359,"p":"102315.35000000","q":"0.00009000","T":1737445349275,"m":false,"M":true}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeTrade("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *TradeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeTrade(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedTradee() {
	msg := []byte(`{"stream":"btcusdt@trade","data":{"e":"trade","E":1737445709952,"s":"BTCUSDT","t":1075762,"p":"102283.99000000","q":"0.00009000","T":1737445709952,"m":true,"M":true}}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedTrade([]string{"btcusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedTradeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedTrade(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedTrade(r1, r2 *CombinedTradeEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeTrade(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeTrade(r1, r2 *TradeEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.TradeId, r2.TradeId, "TradeId")
	r.Equal(r1.Price, r2.Price, "Price")
	r.Equal(r1.Quantity, r2.Quantity, "Quantity")
	r.Equal(r1.TradeTime, r2.TradeTime, "TradeTime")
	r.Equal(r1.IsBuyerMaker, r2.IsBuyerMaker, "IsBuyerMaker")
	r.Equal(r1.Placeholder, r2.Placeholder, "Placeholder")
}

func (s *websocketStreamsTestSuite) TestSubscribeKline() {
	msg := []byte(`{
  "e": "kline",
  "E": 1737447130004,
  "s": "BTCUSDT",
  "k": {
    "t": 1737447120000,
    "T": 1737447179999,
    "s": "BTCUSDT",
    "i": "1m",
    "f": 1078381,
    "L": 1078403,
    "o": "101852.71000000",
    "c": "101852.71000000",
    "h": "101853.37000000",
    "l": "101852.71000000",
    "v": "0.06244000",
    "n": 23,
    "x": false,
    "q": "6359.68560820",
    "V": "0.00363000",
    "Q": "369.72773310",
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
	r.Equal(r1.Kline.StartTime, r2.Kline.StartTime, "Kline.StartTime")
	r.Equal(r1.Kline.CloseTime, r2.Kline.CloseTime, "Kline.CloseTime")
	r.Equal(r1.Kline.Symbol, r2.Kline.Symbol, "Kline.Symbol")
	r.Equal(r1.Kline.Interval, r2.Kline.Interval, "Kline.Interval")
	r.Equal(r1.Kline.FirstTradeId, r2.Kline.FirstTradeId, "Kline.FirstTradeId")
	r.Equal(r1.Kline.LastTradeId, r2.Kline.LastTradeId, "Kline.LastTradeId")
	r.Equal(r1.Kline.OpenPrice, r2.Kline.OpenPrice, "Kline.OpenPrice")
	r.Equal(r1.Kline.ClosePrice, r2.Kline.ClosePrice, "Kline.ClosePrice")
	r.Equal(r1.Kline.HighPrice, r2.Kline.HighPrice, "Kline.HighPrice")
	r.Equal(r1.Kline.LowPrice, r2.Kline.LowPrice, "Kline.LowPrice")
	r.Equal(r1.Kline.BaseAssetVolume, r2.Kline.BaseAssetVolume, "Kline.BaseAssetVolume")
	r.Equal(r1.Kline.NumberOfTrades, r2.Kline.NumberOfTrades, "Kline.NumberOfTrades")
	r.Equal(r1.Kline.IsClosed, r2.Kline.IsClosed, "Kline.IsClosed")
	r.Equal(r1.Kline.QuoteAssetVolume, r2.Kline.QuoteAssetVolume, "Kline.QuoteAssetVolume")
	r.Equal(r1.Kline.TakerBaseVolume, r2.Kline.TakerBaseVolume, "Kline.TakerBaseVolume")
	r.Equal(r1.Kline.TakerQuoteVolume, r2.Kline.TakerQuoteVolume, "Kline.TakerQuoteVolume")
	r.Equal(r1.Kline.Placeholder, r2.Kline.Placeholder, "Kline.Placeholder")
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

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeTicker() {
	msg := []byte(`{"e":"24hrTicker","E":1737510248866,"s":"BTCUSDT","p":"4434.00000000","P":"4.376","w":"98412.56689980","x":"101372.11000000","c":"105766.00000000","Q":"0.00015000","b":"105498.24000000","B":"0.00172000","a":"105766.00000000","A":"0.00025000","o":"101332.00000000","h":"108228.00000000","l":"25262.00000000","v":"310.58183000","q":"30565155.12273710","O":1737423848866,"C":1737510248866,"F":1029731,"L":1174610,"n":144880}`)
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
    "E": 1737510374743,
    "s": "BTCUSDT",
    "p": "4274.48000000",
    "P": "4.222",
    "w": "98414.60451365",
    "x": "101279.06000000",
    "c": "105510.47000000",
    "Q": "0.00015000",
    "b": "105498.42000000",
    "B": "0.00284000",
    "a": "105510.47000000",
    "A": "0.00405000",
    "o": "101235.99000000",
    "h": "108228.00000000",
    "l": "25262.00000000",
    "v": "310.54254000",
    "q": "30561921.25876560",
    "O": 1737423974743,
    "C": 1737510374743,
    "F": 1029964,
    "L": 1174818,
    "n": 144855
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
	r.Equal(r1.FirstTradePrice, r2.FirstTradePrice, "FirstTradePrice")
	r.Equal(r1.LastTradePrice, r2.LastTradePrice, "LastTradePrice")
	r.Equal(r1.LastQuantity, r2.LastQuantity, "LastQuantity")
	r.Equal(r1.BestBidPrice, r2.BestBidPrice, "BestBidPrice")
	r.Equal(r1.BestBidQuantity, r2.BestBidQuantity, "BestBidQuantity")
	r.Equal(r1.BestAskPrice, r2.BestAskPrice, "BestAskPrice")
	r.Equal(r1.BestAskQuantity, r2.BestAskQuantity, "BestAskQuantity")
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
  "E": 1737510503064,
  "s": "ETHBTC",
  "p": "-0.00047000",
  "P": "-1.470",
  "w": "0.03160217",
  "x": "0.03199000",
  "c": "0.03151000",
  "Q": "0.07970000",
  "b": "0.03151000",
  "B": "0.20600000",
  "a": "0.03153000",
  "A": "12.54560000",
  "o": "0.03198000",
  "h": "0.03332000",
  "l": "0.03109000",
  "v": "2410.65430000",
  "q": "76.18191045",
  "O": 1737424103063,
  "C": 1737510503063,
  "F": 187970,
  "L": 217530,
  "n": 29561
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

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeTickerWindowSize() {
	msg := []byte(`{"e":"1hTicker","E":1737510739169,"s":"BTCUSDT","p":"-121.89000000","P":"-0.115","w":"105901.46914485","o":"105871.50000000","h":"107357.47000000","l":"105457.69000000","c":"105749.61000000","v":"10.82688000","q":"1146582.49825500","O":1737507120000,"C":1737510738330,"F":1169205,"L":1175588,"n":6384}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeTickerWindowSize("btcusdt", "1h").Do(context.Background())
	r := s.r()
	var testResp *TickerWindowSizeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeTickerWindowSize(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedTickerWindowSize() {
	msg := []byte(`{"stream":"btcusdt@ticker_1h","data":{"e":"1hTicker","E":1737510915179,"s":"BTCUSDT","p":"-104.85000000","P":"-0.099","w":"105896.30634406","o":"105899.99000000","h":"107357.47000000","l":"105457.69000000","c":"105795.14000000","v":"11.08641000","q":"1174009.86961580","O":1737507300000,"C":1737510915163,"F":1169469,"L":1175978,"n":6510}}`)
	server := s.setup(msg)
	defer server.Close()
	symbols := make(map[string]string)
	symbols["btcusdt"] = "1h"
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedTickerWindowSize(symbols).Do(context.Background())
	r := s.r()
	var testResp *CombinedTickerWindowSizeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedTickerWindowSize(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedTickerWindowSize(r1, r2 *CombinedTickerWindowSizeEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeTickerWindowSize(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeTickerWindowSize(r1, r2 *TickerWindowSizeEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PriceChange, r2.PriceChange, "PriceChange")
	r.Equal(r1.PriceChangePercent, r2.PriceChangePercent, "PriceChangePercent")
	r.Equal(r1.OpenPrice, r2.OpenPrice, "OpenPrice")
	r.Equal(r1.HighPrice, r2.HighPrice, "HighPrice")
	r.Equal(r1.LowPrice, r2.LowPrice, "LowPrice")
	r.Equal(r1.LastPrice, r2.LastPrice, "LastPrice")
	r.Equal(r1.WeightedAveragePrice, r2.WeightedAveragePrice, "WeightedAveragePrice")
	r.Equal(r1.TotalTradedBaseVolume, r2.TotalTradedBaseVolume, "TotalTradedBaseVolume")
	r.Equal(r1.TotalTradedQuoteVolume, r2.TotalTradedQuoteVolume, "TotalTradedQuoteVolume")
	r.Equal(r1.OpenTime, r2.OpenTime, "OpenTime")
	r.Equal(r1.CloseTime, r2.CloseTime, "CloseTime")
	r.Equal(r1.FirstTradeID, r2.FirstTradeID, "FirstTradeID")
	r.Equal(r1.LastTradeID, r2.LastTradeID, "LastTradeID")
	r.Equal(r1.NumberOfTrades, r2.NumberOfTrades, "NumberOfTrades")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeTickerWindowSizeArr() {
	msg := []byte(`[{
  "e": "1hTicker",
  "E": 1737538649880,
  "s": "ETHBTC",
  "p": "0.00008000",
  "P": "0.256",
  "w": "0.03134652",
  "o": "0.03131000",
  "h": "0.03157000",
  "l": "0.03127000",
  "c": "0.03139000",
  "v": "80.29600000",
  "q": "2.51700051",
  "O": 1737535020000,
  "C": 1737538649070,
  "F": 224022,
  "L": 225009,
  "n": 988
}]`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeTickerWindowSizeArr("1h").Do(context.Background())
	r := s.r()
	var testResp []*TickerWindowSizeEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			for i := range event {
				s.assertTestSubscribeTickerWindowSize(event[i], testResp[i])
			}
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeBookTicker() {
	msg := []byte(`{"u":5118292,"s":"BTCUSDT","b":"105024.19000000","B":"0.00296000","a":"105024.20000000","A":"0.00434000"}`)
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
	msg := []byte(`{"stream":"btcusdt@bookTicker","data":{"u":5119060,"s":"BTCUSDT","b":"104965.00000000","B":"0.00358000","a":"104965.01000000","A":"0.00393000"}}`)
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
	r.Equal(r1.UpdateId, r2.UpdateId, "UpdateId")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.BestBidPrice, r2.BestBidPrice, "BestBidPrice")
	r.Equal(r1.BestBidQty, r2.BestBidQty, "BestBidQty")
	r.Equal(r1.BestAskPrice, r2.BestAskPrice, "BestAskPrice")
	r.Equal(r1.BestAskQty, r2.BestAskQty, "BestAskQty")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeAvgPrice() {
	msg := []byte(`{"e":"avgPrice","E":1737539463819,"s":"BTCUSDT","i":"5m","w":"105024.24317124","T":1737539463818}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeAvgPrice("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *AvgPriceEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeAvgPrice(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedAvgPrice() {
	msg := []byte(`{"stream":"btcusdt@avgPrice","data":{"e":"avgPrice","E":1737539556088,"s":"BTCUSDT","i":"5m","w":"105022.81101341","T":1737539556088}}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedAvgPrice([]string{"btcusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedAvgPriceEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertSubscribeCombinedAvgPrice(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) assertSubscribeCombinedAvgPrice(r1, r2 *CombinedAvgPriceEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeAvgPrice(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeAvgPrice(r1, r2 *AvgPriceEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.Interval, r2.Interval, "Interval")
	r.Equal(r1.AveragePrice, r2.AveragePrice, "AveragePrice")
	r.Equal(r1.TradeTime, r2.TradeTime, "TradeTime")
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeDepthLevel() {
	msg := []byte(`{"lastUpdateId":5129092,"bids":[["104920.00000000","0.00263000"],["104919.66000000","0.00258000"],["104919.10000000","0.00444000"],["104918.63000000","0.00415000"],["104918.59000000","0.00263000"]],"asks":[["104920.01000000","0.00053000"],["104920.58000000","0.00296000"],["104920.76000000","0.00415000"],["104921.06000000","0.00248000"],["104921.64000000","0.00372000"]]}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeDepthLevel("btcusdt", 5).Do(context.Background())
	r := s.r()
	var testResp *DepthLevelEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeDepthLevel(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedDepthLevel() {
	msg := []byte(`{
  "stream": "btcusdt@depth5@100ms",
  "data": {
    "lastUpdateId": 5137787,
    "bids": [
      [
        "104791.99000000",
        "0.00268000"
      ],
      [
        "104791.88000000",
        "0.00397000"
      ],
      [
        "104791.49000000",
        "0.00012000"
      ],
      [
        "104791.36000000",
        "0.00425000"
      ],
      [
        "104791.25000000",
        "0.00354000"
      ]
    ],
    "asks": [
      [
        "104792.01000000",
        "0.00093000"
      ],
      [
        "104792.25000000",
        "0.00349000"
      ],
      [
        "104792.59000000",
        "0.00354000"
      ],
      [
        "104792.78000000",
        "0.00320000"
      ],
      [
        "104792.79000000",
        "0.00263000"
      ]
    ]
  }
}`)
	server := s.setup(msg)
	defer server.Close()
	symbols := make(map[string]int)
	symbols["btcusdt"] = 5
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedDepthLevel(symbols).Do(context.Background())
	r := s.r()
	var testResp *CombinedDepthLevelEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedDepthLevel(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedDepthLevel(r1, r2 *CombinedDepthLevelEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeDepthLevel(r1.Data, r2.Data)
}
func (s *websocketStreamsTestSuite) assertTestSubscribeDepthLevel(r1, r2 *DepthLevelEvent) {
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

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeDepth() {
	msg := []byte(`{"e":"depthUpdate","E":1737541559004,"s":"BTCUSDT","U":5142163,"u":5142171,"b":[["104778.32000000","0.00425000"],["104774.76000000","0.00000000"],["104774.71000000","0.00244000"],["104767.84000000","0.00000000"],["104710.20000000","0.00000000"]],"a":[["104789.35000000","0.00344000"],["104791.72000000","0.00463000"],["136211.81000000","0.00232000"]]}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeDepth("btcusdt").Do(context.Background())
	r := s.r()
	var testResp *DepthEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeDepth(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}

func (s *websocketStreamsTestSuite) TestWebSocketSubscribeCombinedDepth() {
	msg := []byte(`{"stream":"btcusdt@depth","data":{"e":"depthUpdate","E":1737554575004,"s":"BTCUSDT","U":5273187,"u":5273193,"b":[["104112.54000000","0.00000000"],["104102.13000000","0.65934000"],["104036.00000000","0.00308000"],["104022.80000000","0.00000000"]],"a":[["104122.95000000","0.01120000"],["104383.57000000","0.00000000"]]}}`)
	server := s.setup(msg)
	defer server.Close()
	onMessage, onError := s.client.NewWebsocketStreams().SubscribeCombinedDepth([]string{"btcusdt"}).Do(context.Background())
	r := s.r()
	var testResp *CombinedDepthEvent
	r.Empty(json.Unmarshal(msg, &testResp))
	for {
		select {
		case event := <-onMessage:
			s.assertTestSubscribeCombinedDepth(event, testResp)
			return
		case err := <-onError:
			s.Error(err)
			return
		}
	}
}
func (s *websocketStreamsTestSuite) assertTestSubscribeCombinedDepth(r1, r2 *CombinedDepthEvent) {
	r := s.r()
	r.Equal(r1.Stream, r2.Stream, "Stream")
	s.assertTestSubscribeDepth(r1.Data, r2.Data)
}

func (s *websocketStreamsTestSuite) assertTestSubscribeDepth(r1, r2 *DepthEvent) {
	r := s.r()
	r.Equal(r1.Event, r2.Event, "Event")
	r.Equal(r1.Time, r2.Time, "Time")
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.FirstId, r2.FirstId, "FirstId")
	r.Equal(r1.FinalId, r2.FinalId, "FinalId")
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
