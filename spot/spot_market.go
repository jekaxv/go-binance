package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// Depth Get depth of a market
type Depth struct {
	c *Client
	r *core.Request
}

type DepthResponse struct {
	LastUpdateId int                 `json:"lastUpdateId"`
	Bids         [][]decimal.Decimal `json:"bids"` // first is PRICE, second is QTY
	Asks         [][]decimal.Decimal `json:"asks"`
}

func (s *Depth) Symbol(symbol string) *Depth {
	s.r.Set("symbol", symbol)
	return s
}

// Limit Default 100; max 5000. If limit > 5000, then the response will truncate to 5000.
func (s *Depth) Limit(limit int) *Depth {
	s.r.Set("limit", limit)
	return s
}

func (s *Depth) Do(ctx context.Context) (*DepthResponse, error) {
	resp := new(DepthResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// Trades Get recent trades.
type Trades struct {
	c *Client
	r *core.Request
}

type TradesResponse struct {
	Id           int64           `json:"id,omitempty"`
	Price        decimal.Decimal `json:"price,omitempty"`
	Qty          decimal.Decimal `json:"qty,omitempty"`
	Time         int64           `json:"time,omitempty"`
	QuoteQty     decimal.Decimal `json:"quoteQty,omitempty"`
	IsBuyerMaker bool            `json:"isBuyerMaker"`
	IsBestMatch  bool            `json:"isBestMatch"`
}

func (s *Trades) Symbol(symbol string) *Trades {
	s.r.Set("symbol", symbol)
	return s
}

// Limit Default 500; max 1000.
func (s *Trades) Limit(limit int) *Trades {
	s.r.Set("limit", limit)
	return s
}

func (s *Trades) Do(ctx context.Context) ([]*TradesResponse, error) {
	resp := make([]*TradesResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// HistoricalTrades Get older trades.
type HistoricalTrades struct {
	c *Client
	r *core.Request
}

func (s *HistoricalTrades) Symbol(symbol string) *HistoricalTrades {
	s.r.Set("symbol", symbol)
	return s
}

// Limit Default 500; max 1000.
func (s *HistoricalTrades) Limit(limit int) *HistoricalTrades {
	s.r.Set("limit", limit)
	return s
}

// FromId Trade id to fetch from. Default gets most recent trades.
func (s *HistoricalTrades) FromId(fromId int64) *HistoricalTrades {
	s.r.Set("fromId", fromId)
	return s
}

func (s *HistoricalTrades) Do(ctx context.Context) ([]*TradesResponse, error) {
	resp := make([]*TradesResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

type AggTrades struct {
	c *Client
	r *core.Request
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
	s.r.Set("symbol", symbol)
	return s
}

// FromId ID to get aggregate trades from INCLUSIVE.
func (s *AggTrades) FromId(fromId int64) *AggTrades {
	s.r.Set("fromId", fromId)
	return s
}

// StartTime Timestamp in ms to get aggregate trades from INCLUSIVE.
func (s *AggTrades) StartTime(startTime int64) *AggTrades {
	s.r.Set("startTime", startTime)
	return s
}

// EndTime Timestamp in ms to get aggregate trades until INCLUSIVE.
func (s *AggTrades) EndTime(endTime int64) *AggTrades {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500; max 1000.
func (s *AggTrades) Limit(limit int) *AggTrades {
	s.r.Set("limit", limit)
	return s
}

func (s *AggTrades) Do(ctx context.Context) ([]*AggTradesResponse, error) {
	resp := make([]*AggTradesResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// KlineData Kline/candlestick bars for a symbol. Klines are uniquely identified by their open time.
type KlineData struct {
	c *Client
	r *core.Request
}

func (s *KlineData) Symbol(symbol string) *KlineData {
	s.r.Set("symbol", symbol)
	return s
}

func (s *KlineData) Interval(interval core.IntervalEnum) *KlineData {
	s.r.Set("interval", interval)
	return s
}

func (s *KlineData) StartTime(startTime int64) *KlineData {
	s.r.Set("startTime", startTime)
	return s
}

func (s *KlineData) EndTime(endTime int64) *KlineData {
	s.r.Set("endTime", endTime)
	return s
}

// TimeZone Default: 0 (UTC)
func (s *KlineData) TimeZone(timeZone string) *KlineData {
	s.r.Set("timeZone", timeZone)
	return s
}

// Limit Default 500; max 1000.
func (s *KlineData) Limit(limit int) *KlineData {
	s.r.Set("limit", limit)
	return s
}

func (s *KlineData) Do(ctx context.Context) ([]*KlineResult, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	res := make([][]any, 0)
	if err := json.Unmarshal(s.c.rawBody(), &res); err != nil {
		return nil, err
	}
	return parseKlineData(res), nil
}

func parseKlineData(res [][]any) []*KlineResult {
	resp := make([]*KlineResult, 0)
	for _, v := range res {
		openPrice, _ := decimal.NewFromString(v[1].(string))
		highPrice, _ := decimal.NewFromString(v[2].(string))
		lowPrice, _ := decimal.NewFromString(v[3].(string))
		closePrice, _ := decimal.NewFromString(v[4].(string))
		volumePrice, _ := decimal.NewFromString(v[5].(string))
		quoteAssetVolume, _ := decimal.NewFromString(v[7].(string))
		takerBuyBaseAssetVolume, _ := decimal.NewFromString(v[9].(string))
		takerBuyQuoteAssetVolume, _ := decimal.NewFromString(v[10].(string))
		resp = append(resp, &KlineResult{
			OpenTime:                 int64(v[0].(float64)),
			OpenPrice:                openPrice,
			HighPrice:                highPrice,
			LowPrice:                 lowPrice,
			ClosePrice:               closePrice,
			Volume:                   volumePrice,
			CloseTime:                int64(v[6].(float64)),
			QuoteAssetVolume:         quoteAssetVolume,
			NumberOfTrades:           int(v[8].(float64)),
			TakerBuyBaseAssetVolume:  takerBuyBaseAssetVolume,
			TakerBuyQuoteAssetVolume: takerBuyQuoteAssetVolume,
		})
	}
	return resp
}

// UIKlines The request is similar to klines having the same parameters and response.
// uiKlines return modified kline data, optimized for presentation of candlestick charts.
type UIKlines struct {
	c *Client
	r *core.Request
}

func (s *UIKlines) Symbol(symbol string) *UIKlines {
	s.r.Set("symbol", symbol)
	return s
}

func (s *UIKlines) Interval(interval core.IntervalEnum) *UIKlines {
	s.r.Set("interval", interval)
	return s
}

func (s *UIKlines) StartTime(startTime int64) *UIKlines {
	s.r.Set("startTime", startTime)
	return s
}

func (s *UIKlines) EndTime(endTime int64) *UIKlines {
	s.r.Set("endTime", endTime)
	return s
}

// TimeZone Default: 0 (UTC)
func (s *UIKlines) TimeZone(timeZone string) *UIKlines {
	s.r.Set("timeZone", timeZone)
	return s
}

// Limit Default 500; max 1000.
func (s *UIKlines) Limit(limit int) *UIKlines {
	s.r.Set("limit", limit)
	return s
}

func (s *UIKlines) Do(ctx context.Context) ([]*KlineResult, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var res [][]any
	if err := json.Unmarshal(s.c.rawBody(), &res); err != nil {
		return nil, err
	}
	return parseKlineData(res), nil
}

// AveragePrice Current average price for a symbol.
type AveragePrice struct {
	c *Client
	r *core.Request
}

type AveragePriceResponse struct {
	Mins      int             `json:"mins"`      // Average price interval (in minutes)
	Price     decimal.Decimal `json:"price"`     // Average price
	CloseTime int64           `json:"closeTime"` // Last trade time
}

func (s *AveragePrice) Symbol(symbol string) *AveragePrice {
	s.r.Set("symbol", symbol)
	return s
}

func (s *AveragePrice) Do(ctx context.Context) (*AveragePriceResponse, error) {
	var resp *AveragePriceResponse
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TickerPrice24h 24 hour rolling window price change statistics. Careful when accessing this with no symbol.
type TickerPrice24h struct {
	c *Client
	r *core.Request
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
	s.r.Set("symbol", symbol)
	return s
}

func (s *TickerPrice24h) Symbols(symbols []string) *TickerPrice24h {
	s.r.Set("symbols", symbols)
	return s
}

// Type Supported values: FULL or MINI. If none provided, the default is FULL
func (s *TickerPrice24h) Type(tickerType core.TickerTypeEnum) *TickerPrice24h {
	s.r.Set("type", tickerType)
	return s
}

func (s *TickerPrice24h) Do(ctx context.Context) ([]*TickerPrice24hResponse, error) {
	resp := make([]*TickerPrice24hResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	var signalResp *TickerPrice24hResponse
	if err := json.Unmarshal(s.c.rawBody(), &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// TradingDayTicker Price change statistics for a trading day.
type TradingDayTicker struct {
	c *Client
	r *core.Request
}

func (s *TradingDayTicker) Symbol(symbol string) *TradingDayTicker {
	s.r.Set("symbol", symbol)
	return s
}

func (s *TradingDayTicker) Symbols(symbols []string) *TradingDayTicker {
	s.r.Set("symbols", symbols)
	return s
}

// TimeZone Default: 0 (UTC)
func (s *TradingDayTicker) TimeZone(timeZone string) *TradingDayTicker {
	s.r.Set("timeZone", timeZone)
	return s
}

// Type Supported values: FULL or MINI. If none provided, the default is FULL
func (s *TradingDayTicker) Type(tickerType core.TickerTypeEnum) *TradingDayTicker {
	s.r.Set("type", tickerType)
	return s
}

func (s *TradingDayTicker) Do(ctx context.Context) ([]*TickerResponse, error) {
	resp := make([]*TickerResponse, 0)

	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	var signalResp *TickerResponse
	if err := json.Unmarshal(s.c.rawBody(), &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// PriceTicker Latest price for a symbol or symbols.
type PriceTicker struct {
	c *Client
	r *core.Request
}

type PriceTickerResponse struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

func (s *PriceTicker) Symbol(symbol string) *PriceTicker {
	s.r.Set("symbol", symbol)
	return s
}

func (s *PriceTicker) Symbols(symbols []string) *PriceTicker {
	s.r.Set("symbols", symbols)
	return s
}

func (s *PriceTicker) Do(ctx context.Context) ([]*PriceTickerResponse, error) {
	resp := make([]*PriceTickerResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	var signalResp *PriceTickerResponse
	if err := json.Unmarshal(s.c.rawBody(), &signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// OrderBookTicker Best price/qty on the order book for a symbol or symbols.
type OrderBookTicker struct {
	c *Client
	r *core.Request
}

type OrderBookTickerResponse struct {
	Symbol   string          `json:"symbol"`
	BidPrice decimal.Decimal `json:"bidPrice"`
	BidQty   decimal.Decimal `json:"bidQty"`
	AskPrice decimal.Decimal `json:"askPrice"`
	AskQty   decimal.Decimal `json:"askQty"`
}

func (s *OrderBookTicker) Symbol(symbol string) *OrderBookTicker {
	s.r.Set("symbol", symbol)
	return s
}

func (s *OrderBookTicker) Symbols(symbols []string) *OrderBookTicker {
	s.r.Set("symbols", symbols)
	return s
}

func (s *OrderBookTicker) Do(ctx context.Context) ([]*OrderBookTickerResponse, error) {
	resp := make([]*OrderBookTickerResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	signalResp := new(OrderBookTickerResponse)
	if err := json.Unmarshal(s.c.rawBody(), signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}

// Ticker Rolling window price change statistics
type Ticker struct {
	c *Client
	r *core.Request
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
	s.r.Set("symbol", symbol)
	return s
}

func (s *Ticker) Symbols(symbols []string) *Ticker {
	s.r.Set("symbols", symbols)
	return s
}

// WindowSize Defaults to 1d if no parameter provided
// Supported windowSize values:
// 1m,2m....59m for minutes
// 1h, 2h....23h - for hours
// 1d...7d - for days
func (s *Ticker) WindowSize(windowSize string) *Ticker {
	s.r.Set("windowSize", windowSize)
	return s
}

// Type Supported values: FULL or MINI. If none provided, the default is FULL
func (s *Ticker) Type(tickerType core.TickerTypeEnum) *Ticker {
	s.r.Set("type", tickerType)
	return s
}

func (s *Ticker) Do(ctx context.Context) ([]*TickerResponse, error) {
	resp := make([]*TickerResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	signalResp := new(TickerResponse)
	if err := json.Unmarshal(s.c.rawBody(), signalResp); err != nil {
		return nil, err
	}
	resp = append(resp, signalResp)
	return resp, nil
}
