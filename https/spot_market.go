package https

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
	"github.com/shopspring/decimal"
)

// Depth Get depth of a market
type Depth struct {
	c      *Client
	symbol string
	limit  *uint
}

type DepthResponse struct {
	LastUpdateId int                 `json:"lastUpdateId"`
	Bids         [][]decimal.Decimal `json:"bids"` // first is PRICE, second is QTY
	Asks         [][]decimal.Decimal `json:"asks"`
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
	Id           uint64          `json:"id,omitempty"`
	Price        decimal.Decimal `json:"price,omitempty"`
	Qty          decimal.Decimal `json:"qty,omitempty"`
	Time         uint64          `json:"time,omitempty"`
	QuoteQty     decimal.Decimal `json:"quoteQty,omitempty"`
	IsBuyerMaker bool            `json:"isBuyerMaker"`
	IsBestMatch  bool            `json:"isBestMatch"`
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

func (s *HistoricalTrades) Do(ctx context.Context) ([]*TradesResponse, error) {
	var resp []*TradesResponse
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

type AggTrades struct {
	c         *Client
	symbol    string
	fromId    *uint64
	startTime *uint64
	endTime   *uint64
	limit     *uint
}

type AggTradesResponse struct {
	TradeId     int             `json:"a"`
	Price       decimal.Decimal `json:"p"`
	Quantity    decimal.Decimal `json:"q"`
	FirstId     int             `json:"f"`
	LastId      int             `json:"l"`
	Timestamp   int64           `json:"T"`
	IsMaker     bool            `json:"m"`
	IsBestPrice bool            `json:"M"`
}

func (s *AggTrades) Symbol(symbol string) *AggTrades {
	s.symbol = symbol
	return s
}

// FromId ID to get aggregate trades from INCLUSIVE.
func (s *AggTrades) FromId(fromId uint64) *AggTrades {
	s.fromId = &fromId
	return s
}

// StartTime Timestamp in ms to get aggregate trades from INCLUSIVE.
func (s *AggTrades) StartTime(startTime uint64) *AggTrades {
	s.startTime = &startTime
	return s
}

// EndTime Timestamp in ms to get aggregate trades until INCLUSIVE.
func (s *AggTrades) EndTime(endTime uint64) *AggTrades {
	s.endTime = &endTime
	return s
}

// Limit Default 500; max 1000.
func (s *AggTrades) Limit(limit uint) *AggTrades {
	s.limit = &limit
	return s
}

func (s *AggTrades) Do(ctx context.Context) ([]*AggTradesResponse, error) {
	var resp []*AggTradesResponse
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
	OpenTime                 uint64          `json:"openTime"`
	OpenPrice                decimal.Decimal `json:"openPrice"`
	HighPrice                decimal.Decimal `json:"highPrice"`
	LowPrice                 decimal.Decimal `json:"lowPrice"`
	ClosePrice               decimal.Decimal `json:"closePrice"`
	Volume                   decimal.Decimal `json:"volume"`
	CloseTime                uint64          `json:"closeTime"`
	QuoteAssetVolume         decimal.Decimal `json:"quoteAssetVolume"`
	NumberOfTrades           int             `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  decimal.Decimal `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume decimal.Decimal `json:"takerBuyQuoteAssetVolume"`
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
		openPrice, _ := decimal.NewFromString(v[1].(string))
		highPrice, _ := decimal.NewFromString(v[2].(string))
		lowPrice, _ := decimal.NewFromString(v[3].(string))
		closePrice, _ := decimal.NewFromString(v[4].(string))
		volumePrice, _ := decimal.NewFromString(v[5].(string))
		quoteAssetVolume, _ := decimal.NewFromString(v[7].(string))
		takerBuyBaseAssetVolume, _ := decimal.NewFromString(v[9].(string))
		takerBuyQuoteAssetVolume, _ := decimal.NewFromString(v[10].(string))
		resp = append(resp, &KlineDataResponse{
			OpenTime:                 uint64(v[0].(float64)),
			OpenPrice:                openPrice,
			HighPrice:                highPrice,
			LowPrice:                 lowPrice,
			ClosePrice:               closePrice,
			Volume:                   volumePrice,
			CloseTime:                uint64(v[6].(float64)),
			QuoteAssetVolume:         quoteAssetVolume,
			NumberOfTrades:           int(v[8].(float64)),
			TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
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
	Mins      int             `json:"mins"`      // Average price interval (in minutes)
	Price     decimal.Decimal `json:"price"`     // Average price
	CloseTime int64           `json:"closeTime"` // Last trade time
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
	Symbol             string          `json:"symbol"` // Symbol Name
	PriceChange        decimal.Decimal `json:"priceChange"`
	PriceChangePercent decimal.Decimal `json:"priceChangePercent"`
	WeightedAvgPrice   decimal.Decimal `json:"weightedAvgPrice"`
	PrevClosePrice     decimal.Decimal `json:"prevClosePrice"`
	LastPrice          decimal.Decimal `json:"lastPrice"` // Closing price of the interval
	LastQty            decimal.Decimal `json:"lastQty"`
	BidPrice           decimal.Decimal `json:"bidPrice"`
	BidQty             decimal.Decimal `json:"bidQty"`
	AskPrice           decimal.Decimal `json:"askPrice"`
	AskQty             decimal.Decimal `json:"askQty"`
	OpenPrice          decimal.Decimal `json:"openPrice"`   // Opening price of the Interval
	HighPrice          decimal.Decimal `json:"highPrice"`   // Highest price in the interval
	LowPrice           decimal.Decimal `json:"lowPrice"`    // Lowest  price in the interval
	Volume             decimal.Decimal `json:"volume"`      // Total trade volume (in base asset)
	QuoteVolume        decimal.Decimal `json:"quoteVolume"` // Total trade volume (in quote asset)
	OpenTime           int64           `json:"openTime"`    // Start of the ticker interval
	CloseTime          int64           `json:"closeTime"`   // End of the ticker interval
	FirstId            int             `json:"firstId"`     // First tradeId considered
	LastId             int             `json:"lastId"`      // Last tradeId considered
	Count              int             `json:"count"`
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
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
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
	Symbol   string          `json:"symbol"`
	BidPrice decimal.Decimal `json:"bidPrice"`
	BidQty   decimal.Decimal `json:"bidQty"`
	AskPrice decimal.Decimal `json:"askPrice"`
	AskQty   decimal.Decimal `json:"askQty"`
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
	Symbol             string          `json:"symbol"`
	PriceChange        decimal.Decimal `json:"priceChange"`
	PriceChangePercent decimal.Decimal `json:"priceChangePercent"`
	WeightedAvgPrice   decimal.Decimal `json:"weightedAvgPrice"`
	OpenPrice          decimal.Decimal `json:"openPrice"`
	HighPrice          decimal.Decimal `json:"highPrice"`
	LowPrice           decimal.Decimal `json:"lowPrice"`
	LastPrice          decimal.Decimal `json:"lastPrice"`
	Volume             decimal.Decimal `json:"volume"`
	QuoteVolume        decimal.Decimal `json:"quoteVolume"`
	OpenTime           int64           `json:"openTime"`
	CloseTime          int64           `json:"closeTime"`
	FirstId            int             `json:"firstId"`
	LastId             int             `json:"lastId"`
	Count              int             `json:"count"`
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
