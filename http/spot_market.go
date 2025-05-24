package http

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
)

// Depth Get depth of a market
type Depth struct {
	c      *Client
	symbol string
	limit  *uint
}

type DepthResponse struct {
	LastUpdateId int        `json:"lastUpdateId"`
	Bids         [][]string `json:"bids"` // first is PRICE, second is QTY
	Asks         [][]string `json:"asks"`
}

func (s *Depth) Symbol(symbol string) *Depth {
	s.symbol = symbol
	return s
}

// Limit Default 100; max 5000. If limit > 5000, then the response will truncate to 5000.
func (s *Depth) Limit(limit uint) *Depth {
	s.limit = &limit
	return s
}

func (s *Depth) Do(ctx context.Context) (*DepthResponse, error) {
	var resp DepthResponse
	s.c.req.set("symbol", s.symbol)
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if err := s.c.invoke(ctx); err != nil {
		return &resp, err
	}
	return &resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// Trades Get recent trades.
type Trades struct {
	c      *Client
	symbol string
	limit  *uint
}

type TradesResponse struct {
	Id           uint64 `json:"id,omitempty"`
	Price        string `json:"price,omitempty"`
	Qty          string `json:"qty,omitempty"`
	Time         uint64 `json:"time,omitempty"`
	QuoteQty     string `json:"quoteQty,omitempty"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

func (s *Trades) Symbol(symbol string) *Trades {
	s.symbol = symbol
	return s
}

// Limit Default 500; max 1000.
func (s *Trades) Limit(limit uint) *Trades {
	s.limit = &limit
	return s
}

func (s *Trades) Do(ctx context.Context) ([]*TradesResponse, error) {
	var resp []*TradesResponse
	s.c.req.set("symbol", s.symbol)
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// HistoricalTrades Get older trades.
type HistoricalTrades struct {
	c      *Client
	symbol string
	limit  *int
	fromId *uint64
}

type HistoricalTradesResponse struct {
	Id           int    `json:"id"`
	Price        string `json:"price"`
	Qty          string `json:"qty"`
	QuoteQty     string `json:"quoteQty"`
	Time         uint64 `json:"time"`
	IsBuyerMaker bool   `json:"isBuyerMaker"`
	IsBestMatch  bool   `json:"isBestMatch"`
}

func (s *HistoricalTrades) Symbol(symbol string) *HistoricalTrades {
	s.symbol = symbol
	return s
}

// Limit Default 500; max 1000.
func (s *HistoricalTrades) Limit(limit int) *HistoricalTrades {
	s.limit = &limit
	return s
}

// FromId Trade id to fetch from. Default gets most recent trades.
func (s *HistoricalTrades) FromId(fromId uint64) *HistoricalTrades {
	s.fromId = &fromId
	return s
}

func (s *HistoricalTrades) Do(ctx context.Context) ([]*HistoricalTradesResponse, error) {
	var resp []*HistoricalTradesResponse
	s.c.req.set("symbol", s.symbol)
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.fromId != nil {
		s.c.req.set("fromId", *s.fromId)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

type AggregateTrades struct {
	c         *Client
	symbol    string
	fromId    *uint64
	startTime *uint64
	endTime   *uint64
	limit     *uint
}

type AggregateTradesResponse struct {
	TradeId     int    `json:"a"`
	Price       string `json:"p"`
	Quantity    string `json:"q"`
	FirstId     int    `json:"f"`
	LastId      int    `json:"l"`
	Timestamp   int64  `json:"T"`
	IsMaker     bool   `json:"m"`
	IsBestPrice bool   `json:"M"`
}

func (s *AggregateTrades) Symbol(symbol string) *AggregateTrades {
	s.symbol = symbol
	return s
}

// FromId ID to get aggregate trades from INCLUSIVE.
func (s *AggregateTrades) FromId(fromId uint64) *AggregateTrades {
	s.fromId = &fromId
	return s
}

// StartTime Timestamp in ms to get aggregate trades from INCLUSIVE.
func (s *AggregateTrades) StartTime(startTime uint64) *AggregateTrades {
	s.startTime = &startTime
	return s
}

// EndTime Timestamp in ms to get aggregate trades until INCLUSIVE.
func (s *AggregateTrades) EndTime(endTime uint64) *AggregateTrades {
	s.endTime = &endTime
	return s
}

// Limit Default 500; max 1000.
func (s *AggregateTrades) Limit(limit uint) *AggregateTrades {
	s.limit = &limit
	return s
}

func (s *AggregateTrades) Do(ctx context.Context) ([]*AggregateTradesResponse, error) {
	var resp []*AggregateTradesResponse
	s.c.req.set("symbol", s.symbol)
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.fromId != nil {
		s.c.req.set("fromId", *s.fromId)
	}
	if s.startTime != nil {
		s.c.req.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.req.set("endTime", *s.endTime)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// KlineData Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.
type KlineData struct {
	c         *Client
	symbol    string
	interval  types.IntervalEnum
	startTime *uint64
	endTime   *uint64
	timeZone  *string
	limit     *uint
}

type KlineDataResponse struct {
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

func (s *KlineData) Symbol(symbol string) *KlineData {
	s.symbol = symbol
	return s
}

func (s *KlineData) Interval(interval types.IntervalEnum) *KlineData {
	s.interval = interval
	return s
}

func (s *KlineData) StartTime(startTime uint64) *KlineData {
	s.startTime = &startTime
	return s
}

func (s *KlineData) EndTime(endTime uint64) *KlineData {
	s.endTime = &endTime
	return s
}

// TimeZone Default: 0 (UTC)
func (s *KlineData) TimeZone(timeZone string) *KlineData {
	s.timeZone = &timeZone
	return s
}

// Limit Default 500; max 1000.
func (s *KlineData) Limit(limit uint) *KlineData {
	s.limit = &limit
	return s
}

func (s *KlineData) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	s.c.req.set("symbol", s.symbol)
	s.c.req.set("interval", s.interval)
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.startTime != nil {
		s.c.req.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.req.set("endTime", *s.endTime)
	}
	if s.timeZone != nil {
		s.c.req.set("timeZone", *s.timeZone)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.resp.rawBody)
}

func parseKlineData(rawBody []byte) ([]*KlineDataResponse, error) {
	var resp []*KlineDataResponse
	var res [][]any
	if err := json.Unmarshal(rawBody, &res); err != nil {
		return resp, err
	}
	for _, v := range res {
		resp = append(resp, &KlineDataResponse{
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
}

// UIKlines The request is similar to klines having the same parameters and response.
// uiKlines return modified kline data, optimized for presentation of candlestick charts.
type UIKlines struct {
	c         *Client
	symbol    string
	interval  types.IntervalEnum
	startTime *uint64
	endTime   *uint64
	timeZone  *string
	limit     *uint
}

func (s *UIKlines) Symbol(symbol string) *UIKlines {
	s.symbol = symbol
	return s
}

func (s *UIKlines) Interval(interval types.IntervalEnum) *UIKlines {
	s.interval = interval
	return s
}

func (s *UIKlines) StartTime(startTime uint64) *UIKlines {
	s.startTime = &startTime
	return s
}

func (s *UIKlines) EndTime(endTime uint64) *UIKlines {
	s.endTime = &endTime
	return s
}

// TimeZone Default: 0 (UTC)
func (s *UIKlines) TimeZone(timeZone string) *UIKlines {
	s.timeZone = &timeZone
	return s
}

// Limit Default 500; max 1000.
func (s *UIKlines) Limit(limit uint) *UIKlines {
	s.limit = &limit
	return s
}

func (s *UIKlines) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	s.c.req.set("symbol", s.symbol)
	s.c.req.set("interval", s.interval)
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.startTime != nil {
		s.c.req.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.req.set("endTime", *s.endTime)
	}
	if s.timeZone != nil {
		s.c.req.set("timeZone", *s.timeZone)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.resp.rawBody)
}

// AveragePrice Current average price for a symbol.
type AveragePrice struct {
	c      *Client
	symbol string
}

type AveragePriceResponse struct {
	Mins      int    `json:"mins"`      // Average price interval (in minutes)
	Price     string `json:"price"`     // Average price
	CloseTime int64  `json:"closeTime"` // Last trade time
}

func (s *AveragePrice) Symbol(symbol string) *AveragePrice {
	s.symbol = symbol
	return s
}

func (s *AveragePrice) Do(ctx context.Context) (*AveragePriceResponse, error) {
	var resp *AveragePriceResponse
	s.c.req.set("symbol", s.symbol)
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// TickerPrice24h 24 hour rolling window price change statistics. Careful when accessing this with no symbol.
type TickerPrice24h struct {
	c          *Client
	symbol     *string
	symbols    []string
	tickerType *types.TickerTypeEnum
}

type TickerPrice24hResponse struct {
	Symbol             string `json:"symbol"` // Symbol Name
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	PrevClosePrice     string `json:"prevClosePrice"`
	LastPrice          string `json:"lastPrice"` // Closing price of the interval
	LastQty            string `json:"lastQty"`
	BidPrice           string `json:"bidPrice"`
	BidQty             string `json:"bidQty"`
	AskPrice           string `json:"askPrice"`
	AskQty             string `json:"askQty"`
	OpenPrice          string `json:"openPrice"`   // Opening price of the Interval
	HighPrice          string `json:"highPrice"`   // Highest price in the interval
	LowPrice           string `json:"lowPrice"`    // Lowest  price in the interval
	Volume             string `json:"volume"`      // Total trade volume (in base asset)
	QuoteVolume        string `json:"quoteVolume"` // Total trade volume (in quote asset)
	OpenTime           int64  `json:"openTime"`    // Start of the ticker interval
	CloseTime          int64  `json:"closeTime"`   // End of the ticker interval
	FirstId            int    `json:"firstId"`     // First tradeId considered
	LastId             int    `json:"lastId"`      // Last tradeId considered
	Count              int    `json:"count"`
}

// Symbol Parameter symbol and symbols cannot be used in combination.
// If neither parameter is sent, tickers for all symbols will be returned in an array.
func (s *TickerPrice24h) Symbol(symbol string) *TickerPrice24h {
	s.symbol = &symbol
	return s
}

func (s *TickerPrice24h) Symbols(symbols []string) *TickerPrice24h {
	s.symbols = symbols
	return s
}

// Type Supported values: FULL or MINI. If none provided, the default is FULL
func (s *TickerPrice24h) Type(tickerType types.TickerTypeEnum) *TickerPrice24h {
	s.tickerType = &tickerType
	return s
}

func (s *TickerPrice24h) Do(ctx context.Context) ([]*TickerPrice24hResponse, error) {
	var resp []*TickerPrice24hResponse
	if s.symbol != nil {
		s.c.req.set("symbol", *s.symbol)
	}
	if len(s.symbols) != 0 {
		s.c.req.set("symbols", s.symbols)
	}
	if s.tickerType != nil {
		s.c.req.set("type", *s.tickerType)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	if s.symbol == nil {
		return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
	}
	var signalResp *TickerPrice24hResponse
	if err := json.Unmarshal(s.c.resp.rawBody, &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// TradingDayTicker Price change statistics for a trading day.
type TradingDayTicker struct {
	c          *Client
	symbol     *string
	symbols    []string
	timeZone   *string
	tickerType *types.TickerTypeEnum
}

func (s *TradingDayTicker) Symbol(symbol string) *TradingDayTicker {
	s.symbol = &symbol
	return s
}

func (s *TradingDayTicker) Symbols(symbols []string) *TradingDayTicker {
	s.symbols = symbols
	return s
}

// TimeZone Default: 0 (UTC)
func (s *TradingDayTicker) TimeZone(timeZone string) *TradingDayTicker {
	s.timeZone = &timeZone
	return s
}

// Type Supported values: FULL or MINI. If none provided, the default is FULL
func (s *TradingDayTicker) Type(tickerType types.TickerTypeEnum) *TradingDayTicker {
	s.tickerType = &tickerType
	return s
}

func (s *TradingDayTicker) Do(ctx context.Context) ([]*TickerResponse, error) {
	var resp []*TickerResponse
	if s.symbol != nil {
		s.c.req.set("symbol", *s.symbol)
	}
	if len(s.symbols) != 0 {
		s.c.req.set("symbols", s.symbols)
	}
	if s.tickerType != nil {
		s.c.req.set("type", *s.tickerType)
	}
	if s.timeZone != nil {
		s.c.req.set("timeZone", *s.timeZone)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	if s.symbol == nil {
		return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
	}
	var signalResp *TickerResponse
	if err := json.Unmarshal(s.c.resp.rawBody, &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// PriceTicker Latest price for a symbol or symbols.
type PriceTicker struct {
	c       *Client
	symbol  *string
	symbols []string
}

type PriceTickerResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func (s *PriceTicker) Symbol(symbol string) *PriceTicker {
	s.symbol = &symbol
	return s
}

func (s *PriceTicker) Symbols(symbols []string) *PriceTicker {
	s.symbols = symbols
	return s
}

func (s *PriceTicker) Do(ctx context.Context) ([]*PriceTickerResponse, error) {
	var resp []*PriceTickerResponse
	if s.symbol != nil {
		s.c.req.set("symbol", *s.symbol)
	}
	if len(s.symbols) != 0 {
		s.c.req.set("symbols", s.symbols)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	if s.symbol == nil {
		return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
	}
	var signalResp *PriceTickerResponse
	if err := json.Unmarshal(s.c.resp.rawBody, &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// OrderBookTicker Best price/qty on the order book for a symbol or symbols.
type OrderBookTicker struct {
	c       *Client
	symbol  *string
	symbols []string
}

type OrderBookTickerResponse struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

func (s *OrderBookTicker) Symbol(symbol string) *OrderBookTicker {
	s.symbol = &symbol
	return s
}

func (s *OrderBookTicker) Symbols(symbols []string) *OrderBookTicker {
	s.symbols = symbols
	return s
}

func (s *OrderBookTicker) Do(ctx context.Context) ([]*OrderBookTickerResponse, error) {
	var resp []*OrderBookTickerResponse
	if s.symbol != nil {
		s.c.req.set("symbol", *s.symbol)
	}
	if len(s.symbols) != 0 {
		s.c.req.set("symbols", s.symbols)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	if s.symbol == nil {
		return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
	}
	var signalResp *OrderBookTickerResponse
	if err := json.Unmarshal(s.c.resp.rawBody, &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// Ticker Rolling window price change statistics
type Ticker struct {
	c          *Client
	symbol     *string
	symbols    []string
	windowSize *string
	tickerType *types.TickerTypeEnum
}

type TickerResponse struct {
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
	FirstId            int    `json:"firstId"`
	LastId             int    `json:"lastId"`
	Count              int    `json:"count"`
}

func (s *Ticker) Symbol(symbol string) *Ticker {
	s.symbol = &symbol
	return s
}

func (s *Ticker) Symbols(symbols []string) *Ticker {
	s.symbols = symbols
	return s
}

// WindowSize Defaults to 1d if no parameter provided
// Supported windowSize values:
// 1m,2m....59m for minutes
// 1h, 2h....23h - for hours
// 1d...7d - for days
func (s *Ticker) WindowSize(windowSize string) *Ticker {
	s.windowSize = &windowSize
	return s
}

// Type Supported values: FULL or MINI. If none provided, the default is FULL
func (s *Ticker) Type(tickerType types.TickerTypeEnum) *Ticker {
	s.tickerType = &tickerType
	return s
}

func (s *Ticker) Do(ctx context.Context) ([]*TickerResponse, error) {
	var resp []*TickerResponse
	if s.symbol != nil {
		s.c.req.set("symbol", *s.symbol)
	}
	if len(s.symbols) != 0 {
		s.c.req.set("symbols", s.symbols)
	}
	if s.tickerType != nil {
		s.c.req.set("type", *s.tickerType)
	}
	if s.windowSize != nil {
		s.c.req.set("windowSize", *s.windowSize)
	}
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	if s.symbol == nil {
		return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
	}
	var signalResp *TickerResponse
	if err := json.Unmarshal(s.c.resp.rawBody, &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}
