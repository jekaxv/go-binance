package ws

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
)

// Depth Get current order book.
type Depth struct {
	c *Client
}

type DepthResult struct {
	LastUpdateId int64      `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"` // [0]Price [1] Quantity
	Asks         [][]string `json:"asks"` // [0]Price [1] Quantity
}

type DepthResponse struct {
	ApiResponse
	Result *DepthResult `json:"result"`
}

func (s *Depth) Symbol(symbol string) *Depth {
	s.c.req.Params["symbol"] = symbol
	return s
}

// Limit Default 100; max 5000
func (s *Depth) Limit(limit uint) *Depth {
	s.c.req.Params["limit"] = limit
	return s
}

func (s *Depth) Do(ctx context.Context) (*DepthResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *DepthResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// TradesRecent Get recent trades
type TradesRecent struct {
	c *Client
}

type TradesResult struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         int64  `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

type TradesRecentResponse struct {
	ApiResponse
	Result []*TradesResult `json:"result"`
}

func (s *TradesRecent) Symbol(symbol string) *TradesRecent {
	s.c.req.Params["symbol"] = symbol
	return s
}

// Limit Default 500; max 1000
func (s *TradesRecent) Limit(limit uint) *TradesRecent {
	s.c.req.Params["limit"] = limit
	return s
}
func (s *TradesRecent) Do(ctx context.Context) (*TradesRecentResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *TradesRecentResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// TradesHistorical Get historical trades.
type TradesHistorical struct {
	c *Client
}

type TradesHistoricalResponse struct {
	ApiResponse
	Result []*TradesResult `json:"result"`
}

func (s *TradesHistorical) Symbol(symbol string) *TradesHistorical {
	s.c.req.Params["symbol"] = symbol
	return s
}

// FromId Trade ID to begin at
func (s *TradesHistorical) FromId(fromId uint64) *TradesHistorical {
	s.c.req.Params["fromId"] = fromId
	return s
}

// Limit Default 500; max 1000
func (s *TradesHistorical) Limit(limit uint) *TradesHistorical {
	s.c.req.Params["limit"] = limit
	return s
}

func (s *TradesHistorical) Do(ctx context.Context) (*TradesHistoricalResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *TradesHistoricalResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// TradesAggregate Get aggregate trades.
type TradesAggregate struct {
	c *Client
}

type TradesAggregateResult struct {
	TradeId     int    `json:"a"`
	Price       string `json:"p"`
	Quantity    string `json:"q"`
	FirstId     int    `json:"f"`
	LastId      int    `json:"l"`
	Timestamp   int64  `json:"T"`
	IsMaker     bool   `json:"m"`
	IsBestPrice bool   `json:"M"`
}

type TradesAggregateResponse struct {
	ApiResponse
	Result []*TradesAggregateResult `json:"result"`
}

func (s *TradesAggregate) Symbol(symbol string) *TradesAggregate {
	s.c.req.Params["symbol"] = symbol
	return s
}

// FromId Aggregate trade ID to begin at
func (s *TradesAggregate) FromId(fromId uint64) *TradesAggregate {
	s.c.req.Params["fromId"] = fromId
	return s
}
func (s *TradesAggregate) StartTime(startTime uint64) *TradesAggregate {
	s.c.req.Params["startTime"] = startTime
	return s
}
func (s *TradesAggregate) EndTime(endTime uint64) *TradesAggregate {
	s.c.req.Params["endTime"] = endTime
	return s
}
func (s *TradesAggregate) Limit(limit uint) *TradesAggregate {
	s.c.req.Params["limit"] = limit
	return s
}

func (s *TradesAggregate) Do(ctx context.Context) (*TradesAggregateResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *TradesAggregateResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// Kline Get klines (candlestick bars).
// Klines are uniquely identified by their open & close time.
type Kline struct {
	c *Client
}
type KlineResult struct {
	OpenTime                 uint64 `json:"openTime"`
	OpenPrice                string `json:"openPrice"`
	HighPrice                string `json:"highPrice"`
	LowPrice                 string `json:"lowPrice"`
	ClosePrice               string `json:"closePrice"`
	Volume                   string `json:"volume"`
	CloseTime                uint64 `json:"closeTime"`
	QuoteAssetVolume         string `json:"quoteAssetVolume"`
	NumberOfTrades           int    `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  string `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume string `json:"takerBuyQuoteAssetVolume"`
}

type KlineRawResult struct {
	ApiResponse
	Result [][]any `json:"result"`
}

type KlineResponse struct {
	ApiResponse
	Result []*KlineResult `json:"result"`
}

func (s *Kline) Symbol(symbol string) *Kline {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *Kline) Interval(interval types.IntervalEnum) *Kline {
	s.c.req.Params["interval"] = interval
	return s
}
func (s *Kline) StartTime(startTime uint64) *Kline {
	s.c.req.Params["startTime"] = startTime
	return s
}
func (s *Kline) EndTime(endTime uint64) *Kline {
	s.c.req.Params["endTime"] = endTime
	return s
}
func (s *Kline) TimeZone(timeZone string) *Kline {
	s.c.req.Params["timeZone"] = timeZone
	return s
}
func (s *Kline) Limit(limit uint) *Kline {
	s.c.req.Params["limit"] = limit
	return s
}

func (s *Kline) Do(ctx context.Context) (*KlineResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var raw *KlineRawResult
			if err := json.Unmarshal(message, &raw); err != nil {
				return nil, err
			}
			resp := new(KlineResponse)
			resp.ApiResponse = raw.ApiResponse
			for _, v := range raw.Result {
				resp.Result = append(resp.Result, &KlineResult{
					OpenTime:                 uint64(v[0].(float64)),
					OpenPrice:                v[1].(string),
					HighPrice:                v[2].(string),
					LowPrice:                 v[3].(string),
					ClosePrice:               v[4].(string),
					Volume:                   v[5].(string),
					CloseTime:                uint64(v[6].(float64)),
					QuoteAssetVolume:         v[7].(string),
					NumberOfTrades:           int(v[8].(float64)),
					TakerBuyBaseAssetVolume:  v[9].(string),
					TakerBuyQuoteAssetVolume: v[10].(string),
				})
			}
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// UiKlines Get klines (candlestick bars) optimized for presentation.
type UiKlines struct {
	c *Client
}

func (s *UiKlines) Symbol(symbol string) *UiKlines {
	s.c.req.Params["symbol"] = symbol
	return s
}
func (s *UiKlines) Interval(interval string) *UiKlines {
	s.c.req.Params["interval"] = interval
	return s
}
func (s *UiKlines) StartTime(startTime uint64) *UiKlines {
	s.c.req.Params["startTime"] = startTime
	return s
}
func (s *UiKlines) EndTime(endTime uint64) *UiKlines {
	s.c.req.Params["endTime"] = endTime
	return s
}
func (s *UiKlines) TimeZone(timeZone string) *UiKlines {
	s.c.req.Params["timeZone"] = timeZone
	return s
}
func (s *UiKlines) Limit(limit uint) *UiKlines {
	s.c.req.Params["limit"] = limit
	return s
}

func (s *UiKlines) Do(ctx context.Context) (*KlineResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var raw *KlineRawResult
			if err := json.Unmarshal(message, &raw); err != nil {
				return nil, err
			}
			resp := new(KlineResponse)
			resp.ApiResponse = raw.ApiResponse
			for _, v := range raw.Result {
				resp.Result = append(resp.Result, &KlineResult{
					OpenTime:                 uint64(v[0].(float64)),
					OpenPrice:                v[1].(string),
					HighPrice:                v[2].(string),
					LowPrice:                 v[3].(string),
					ClosePrice:               v[4].(string),
					Volume:                   v[5].(string),
					CloseTime:                uint64(v[6].(float64)),
					QuoteAssetVolume:         v[7].(string),
					NumberOfTrades:           int(v[8].(float64)),
					TakerBuyBaseAssetVolume:  v[9].(string),
					TakerBuyQuoteAssetVolume: v[10].(string),
				})
			}
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// AveragePrice Get current average price for a symbol.
type AveragePrice struct {
	c *Client
}

type AveragePriceResult struct {
	Mins      int    `json:"mins"`
	Price     string `json:"price"`
	CloseTime int64  `json:"closeTime"`
}

type AveragePriceResponse struct {
	ApiResponse
	Result *AveragePriceResult `json:"result"`
}

func (s *AveragePrice) Symbol(symbol string) *AveragePrice {
	s.c.req.Params["symbol"] = symbol
	return s
}

func (s *AveragePrice) Do(ctx context.Context) (*AveragePriceResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *AveragePriceResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// Ticker24h Get 24-hour rolling window price change statistics.
type Ticker24h struct {
	c *Client
}

type Ticker24hResult struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

type Ticker24hSingleResponse struct {
	ApiResponse
	Result *Ticker24hResult `json:"result"`
}

type Ticker24hResponse struct {
	ApiResponse
	Result []*Ticker24hResult `json:"result"`
}

// Symbol Query ticker for a single symbol
func (s *Ticker24h) Symbol(symbol string) *Ticker24h {
	s.c.req.Params["symbol"] = symbol
	return s
}

// Symbols Query ticker for multiple symbols
func (s *Ticker24h) Symbols(symbols []string) *Ticker24h {
	s.c.req.Params["symbols"] = symbols
	return s
}

// Type ticker type: FULL (default) or MINI
func (s *Ticker24h) Type(type_ types.TickerTypeEnum) *Ticker24h {
	s.c.req.Params["type"] = type_
	return s
}

func (s *Ticker24h) Do(ctx context.Context) (*Ticker24hResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(Ticker24hResponse)
			if s.c.req.Params["symbols"] != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(Ticker24hSingleResponse)
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.ApiResponse = single.ApiResponse
			resp.Result = append(resp.Result, single.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// TickerTradingDay Price change statistics for a trading day.
type TickerTradingDay struct {
	c *Client
}

type TickerTradingDaySingleResponse struct {
	ApiResponse
	Result *TickerResult `json:"result"`
}

type TickerTradingDayResponse struct {
	ApiResponse
	Result []*TickerResult `json:"result"`
}

func (s *TickerTradingDay) Symbol(symbol string) *TickerTradingDay {
	s.c.req.Params["symbol"] = symbol
	return s
}

func (s *TickerTradingDay) Symbols(symbols []string) *TickerTradingDay {
	s.c.req.Params["symbols"] = symbols
	return s
}

// TimeZone Default: 0 (UTC)
func (s *TickerTradingDay) TimeZone(timeZone string) *TickerTradingDay {
	s.c.req.Params["timeZone"] = timeZone
	return s
}

// Type Supported values: FULL or MINI.
// If none provided, the default is FULL
func (s *TickerTradingDay) Type(type_ types.TickerTypeEnum) *TickerTradingDay {
	s.c.req.Params["type"] = type_
	return s
}

func (s *TickerTradingDay) Do(ctx context.Context) (*TickerTradingDayResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(TickerTradingDayResponse)
			if s.c.req.Params["symbols"] != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(TickerTradingDaySingleResponse)
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.ApiResponse = single.ApiResponse
			resp.Result = append(resp.Result, single.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// Ticker Get rolling window price change statistics with a custom window.
// This request is similar to ticker.24hr, but statistics are computed on demand using the arbitrary window you specify.
type Ticker struct {
	c *Client
}
type TickerResult struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	OpenPrice          string `json:"openPrice"`
	HighPrice          string `json:"highPrice"`
	LowPrice           string `json:"lowPrice"`
	LastPrice          string `json:"lastPrice"`
	Volume             string `json:"volume"`
	QuoteVolume        string `json:"quoteVolume"`
	OpenTime           int64  `json:"openTime"`
	CloseTime          int64  `json:"closeTime"`
	FirstId            int64  `json:"firstId"`
	LastId             int64  `json:"lastId"`
	Count              int    `json:"count"`
}

type TickerSingleResponse struct {
	ApiResponse
	Result *TickerResult `json:"result"`
}

type TickerResponse struct {
	ApiResponse
	Result []*TickerResult `json:"result"`
}

func (s *Ticker) Symbol(symbol string) *Ticker {
	s.c.req.Params["symbol"] = symbol
	return s
}

func (s *Ticker) Symbols(symbols []string) *Ticker {
	s.c.req.Params["symbols"] = symbols
	return s
}

func (s *Ticker) Type(type_ types.TickerTypeEnum) *Ticker {
	s.c.req.Params["type"] = type_
	return s
}

// WindowSize Supported window sizes:
// minutes	1m, 2m ... 59m
// hours	1h, 2h ... 23h
// days	1d, 2d ... 7d
func (s *Ticker) WindowSize(windowSize string) *Ticker {
	s.c.req.Params["windowSize"] = windowSize
	return s
}

func (s *Ticker) Do(ctx context.Context) (*TickerResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(TickerResponse)
			if s.c.req.Params["symbols"] != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(TickerSingleResponse)
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.ApiResponse = single.ApiResponse
			resp.Result = append(resp.Result, single.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// TickerPrice Get the latest market price for a symbol.
type TickerPrice struct {
	c *Client
}

type TickerPriceResult struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type TickerPriceSingleResponse struct {
	ApiResponse
	Result *TickerPriceResult `json:"result"`
}

type TickerPriceResponse struct {
	ApiResponse
	Result []*TickerPriceResult `json:"result"`
}

func (s *TickerPrice) Symbol(symbol string) *TickerPrice {
	s.c.req.Params["symbol"] = symbol
	return s
}

func (s *TickerPrice) Symbols(symbols []string) *TickerPrice {
	s.c.req.Params["symbols"] = symbols
	return s
}

func (s *TickerPrice) Do(ctx context.Context) (*TickerPriceResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(TickerPriceResponse)
			if s.c.req.Params["symbols"] != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(TickerPriceSingleResponse)
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.ApiResponse = single.ApiResponse
			resp.Result = append(resp.Result, single.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// TickerBook Get the current best price and quantity on the order book.
type TickerBook struct {
	c *Client
}

type TickerBookResult struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

type TickerBookSingleResponse struct {
	ApiResponse
	Result *TickerBookResult `json:"result"`
}

type TickerBookResponse struct {
	ApiResponse
	Result []*TickerBookResult `json:"result"`
}

func (s *TickerBook) Symbol(symbol string) *TickerBook {
	s.c.req.Params["symbol"] = symbol
	return s
}

func (s *TickerBook) Symbols(symbols []string) *TickerBook {
	s.c.req.Params["symbols"] = symbols
	return s
}

func (s *TickerBook) Do(ctx context.Context) (*TickerBookResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(TickerBookResponse)
			if s.c.req.Params["symbols"] != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(TickerBookSingleResponse)
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.ApiResponse = single.ApiResponse
			resp.Result = append(resp.Result, single.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}
