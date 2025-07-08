package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// WsDepth Get current order book.
type WsDepth struct {
	c *WsClient
	r *core.WsRequest
}

type DepthResult struct {
	LastUpdateId int64               `json:"lastUpdateId"`
	Bids         [][]decimal.Decimal `json:"bids"` // [0]Price [1] Quantity
	Asks         [][]decimal.Decimal `json:"asks"` // [0]Price [1] Quantity
}

type WsDepthResponse struct {
	ApiResponse
	Result *DepthResult `json:"result"`
}

func (s *WsDepth) Symbol(symbol string) *WsDepth {
	s.r.Set("symbol", symbol)
	return s
}

// Limit Default 100; max 5000
func (s *WsDepth) Limit(limit int) *WsDepth {
	s.r.Set("limit", limit)
	return s
}

func (s *WsDepth) Do(ctx context.Context) (*WsDepthResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsDepthResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsTradesRecent Get recent trades
type WsTradesRecent struct {
	c *WsClient
	r *core.WsRequest
}

type TradesResult struct {
	Id           int             `json:"id"`
	Price        decimal.Decimal `json:"price"`
	Qty          decimal.Decimal `json:"qty"`
	QuoteQty     decimal.Decimal `json:"quoteQty"`
	Time         int64           `json:"time"`
	IsBuyerMaker bool            `json:"isBuyerMaker"`
	IsBestMatch  bool            `json:"isBestMatch"`
}

type WsTradesRecentResponse struct {
	ApiResponse
	Result []*TradesResult `json:"result"`
}

func (s *WsTradesRecent) Symbol(symbol string) *WsTradesRecent {
	s.r.Set("symbol", symbol)
	return s
}

// Limit Default 500; max 1000
func (s *WsTradesRecent) Limit(limit int) *WsTradesRecent {
	s.r.Set("limit", limit)
	return s
}
func (s *WsTradesRecent) Do(ctx context.Context) (*WsTradesRecentResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsTradesRecentResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsTradesHistorical Get historical trades.
type WsTradesHistorical struct {
	c *WsClient
	r *core.WsRequest
}

type WsTradesHistoricalResponse struct {
	ApiResponse
	Result []*TradesResult `json:"result"`
}

func (s *WsTradesHistorical) Symbol(symbol string) *WsTradesHistorical {
	s.r.Set("symbol", symbol)
	return s
}

// FromId Trade ID to begin at
func (s *WsTradesHistorical) FromId(fromId int64) *WsTradesHistorical {
	s.r.Set("fromId", fromId)
	return s
}

// Limit Default 500; max 1000
func (s *WsTradesHistorical) Limit(limit int) *WsTradesHistorical {
	s.r.Set("limit", limit)
	return s
}

func (s *WsTradesHistorical) Do(ctx context.Context) (*WsTradesHistoricalResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsTradesHistoricalResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsTradesAggregate Get aggregate trades.
type WsTradesAggregate struct {
	c *WsClient
	r *core.WsRequest
}

type TradesAggregateResult struct {
	TradeId     int             `json:"a"`
	Price       decimal.Decimal `json:"p"`
	Quantity    decimal.Decimal `json:"q"`
	FirstId     int             `json:"f"`
	LastId      int             `json:"l"`
	Timestamp   int64           `json:"T"`
	IsMaker     bool            `json:"m"`
	IsBestPrice bool            `json:"M"`
}

type WsTradesAggregateResponse struct {
	ApiResponse
	Result []*TradesAggregateResult `json:"result"`
}

func (s *WsTradesAggregate) Symbol(symbol string) *WsTradesAggregate {
	s.r.Set("symbol", symbol)
	return s
}

// FromId Aggregate trade ID to begin at
func (s *WsTradesAggregate) FromId(fromId int64) *WsTradesAggregate {
	s.r.Set("fromId", fromId)
	return s
}
func (s *WsTradesAggregate) StartTime(startTime int64) *WsTradesAggregate {
	s.r.Set("startTime", startTime)
	return s
}
func (s *WsTradesAggregate) EndTime(endTime int64) *WsTradesAggregate {
	s.r.Set("endTime", endTime)
	return s
}
func (s *WsTradesAggregate) Limit(limit int) *WsTradesAggregate {
	s.r.Set("limit", limit)
	return s
}

func (s *WsTradesAggregate) Do(ctx context.Context) (*WsTradesAggregateResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsTradesAggregateResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsKline Get klines (candlestick bars).
// Klines are uniquely identified by their open & close time.
type WsKline struct {
	c *WsClient
	r *core.WsRequest
}
type KlineResult struct {
	OpenTime                 int64           `json:"openTime"`
	OpenPrice                decimal.Decimal `json:"openPrice"`
	HighPrice                decimal.Decimal `json:"highPrice"`
	LowPrice                 decimal.Decimal `json:"lowPrice"`
	ClosePrice               decimal.Decimal `json:"closePrice"`
	Volume                   decimal.Decimal `json:"volume"`
	CloseTime                int64           `json:"closeTime"`
	QuoteAssetVolume         decimal.Decimal `json:"quoteAssetVolume"`
	NumberOfTrades           int             `json:"numberOfTrades"`
	TakerBuyBaseAssetVolume  decimal.Decimal `json:"takerBuyBaseAssetVolume"`
	TakerBuyQuoteAssetVolume decimal.Decimal `json:"takerBuyQuoteAssetVolume"`
}

type KlineRawResult struct {
	ApiResponse
	Result [][]any `json:"result"`
}

type WsKlineResponse struct {
	ApiResponse
	Result []*KlineResult `json:"result"`
}

func (s *WsKline) Symbol(symbol string) *WsKline {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsKline) Interval(interval core.IntervalEnum) *WsKline {
	s.r.Set("interval", interval)
	return s
}
func (s *WsKline) StartTime(startTime int64) *WsKline {
	s.r.Set("startTime", startTime)
	return s
}
func (s *WsKline) EndTime(endTime int64) *WsKline {
	s.r.Set("endTime", endTime)
	return s
}
func (s *WsKline) TimeZone(timeZone string) *WsKline {
	s.r.Set("timeZone", timeZone)
	return s
}
func (s *WsKline) Limit(limit int) *WsKline {
	s.r.Set("limit", limit)
	return s
}

func (s *WsKline) Do(ctx context.Context) (*WsKlineResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var raw *KlineRawResult
			if err := json.Unmarshal(message, &raw); err != nil {
				return nil, err
			}
			resp := new(WsKlineResponse)
			resp.ApiResponse = raw.ApiResponse
			resp.Result = parseKlineData(raw.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// WsUiKlines Get klines (candlestick bars) optimized for presentation.
type WsUiKlines struct {
	c *WsClient
	r *core.WsRequest
}

func (s *WsUiKlines) Symbol(symbol string) *WsUiKlines {
	s.r.Set("symbol", symbol)
	return s
}
func (s *WsUiKlines) Interval(interval string) *WsUiKlines {
	s.r.Set("interval", interval)
	return s
}
func (s *WsUiKlines) StartTime(startTime int64) *WsUiKlines {
	s.r.Set("startTime", startTime)
	return s
}
func (s *WsUiKlines) EndTime(endTime int64) *WsUiKlines {
	s.r.Set("endTime", endTime)
	return s
}
func (s *WsUiKlines) TimeZone(timeZone string) *WsUiKlines {
	s.r.Set("timeZone", timeZone)
	return s
}
func (s *WsUiKlines) Limit(limit int) *WsUiKlines {
	s.r.Set("limit", limit)
	return s
}

func (s *WsUiKlines) Do(ctx context.Context) (*WsKlineResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var raw *KlineRawResult
			if err := json.Unmarshal(message, &raw); err != nil {
				return nil, err
			}
			resp := new(WsKlineResponse)
			resp.ApiResponse = raw.ApiResponse
			resp.Result = parseKlineData(raw.Result)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// WsAveragePrice Get current average price for a symbol.
type WsAveragePrice struct {
	c *WsClient
	r *core.WsRequest
}

type AveragePriceResult struct {
	Mins      int             `json:"mins"`
	Price     decimal.Decimal `json:"price"`
	CloseTime int64           `json:"closeTime"`
}

type WsAveragePriceResponse struct {
	ApiResponse
	Result *AveragePriceResult `json:"result"`
}

func (s *WsAveragePrice) Symbol(symbol string) *WsAveragePrice {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsAveragePrice) Do(ctx context.Context) (*WsAveragePriceResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *WsAveragePriceResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsTicker24h Get 24-hour rolling window price change statistics.
type WsTicker24h struct {
	c *WsClient
	r *core.WsRequest
}

type Ticker24hResult struct {
	Symbol             string          `json:"symbol"`
	PriceChange        decimal.Decimal `json:"priceChange"`
	PriceChangePercent decimal.Decimal `json:"priceChangePercent"`
	WeightedAvgPrice   decimal.Decimal `json:"weightedAvgPrice"`
	PrevClosePrice     decimal.Decimal `json:"prevClosePrice"`
	LastPrice          decimal.Decimal `json:"lastPrice"`
	LastQty            decimal.Decimal `json:"lastQty"`
	BidPrice           decimal.Decimal `json:"bidPrice"`
	BidQty             decimal.Decimal `json:"bidQty"`
	AskPrice           decimal.Decimal `json:"askPrice"`
	AskQty             decimal.Decimal `json:"askQty"`
	OpenPrice          decimal.Decimal `json:"openPrice"`
	HighPrice          decimal.Decimal `json:"highPrice"`
	LowPrice           decimal.Decimal `json:"lowPrice"`
	Volume             decimal.Decimal `json:"volume"`
	QuoteVolume        decimal.Decimal `json:"quoteVolume"`
	OpenTime           int64           `json:"openTime"`
	CloseTime          int64           `json:"closeTime"`
	FirstId            int             `json:"firstId"`
	LastId             int             `json:"lastId"`
	Count              int             `json:"count"`
}

type WsTicker24hSingleResponse struct {
	ApiResponse
	Result *Ticker24hResult `json:"result"`
}

type WsTicker24hResponse struct {
	ApiResponse
	Result []*Ticker24hResult `json:"result"`
}

// Symbol Query ticker for a single symbol
func (s *WsTicker24h) Symbol(symbol string) *WsTicker24h {
	s.r.Set("symbol", symbol)
	return s
}

// Symbols Query ticker for multiple symbols
func (s *WsTicker24h) Symbols(symbols []string) *WsTicker24h {
	s.r.Set("symbols", symbols)
	return s
}

// Type ticker type: FULL (default) or MINI
func (s *WsTicker24h) Type(type_ core.TickerTypeEnum) *WsTicker24h {
	s.r.Set("type", type_)
	return s
}

func (s *WsTicker24h) Do(ctx context.Context) (*WsTicker24hResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(WsTicker24hResponse)
			if s.r.Get("symbols") != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(WsTicker24hSingleResponse)
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

// WsTickerTradingDay Price change statistics for a trading day.
type WsTickerTradingDay struct {
	c *WsClient
	r *core.WsRequest
}

type WsTickerTradingDaySingleResponse struct {
	ApiResponse
	Result *TickerResult `json:"result"`
}

type WsTickerTradingDayResponse struct {
	ApiResponse
	Result []*TickerResult `json:"result"`
}

func (s *WsTickerTradingDay) Symbol(symbol string) *WsTickerTradingDay {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsTickerTradingDay) Symbols(symbols []string) *WsTickerTradingDay {
	s.r.Set("symbols", symbols)
	return s
}

// TimeZone Default: 0 (UTC)
func (s *WsTickerTradingDay) TimeZone(timeZone string) *WsTickerTradingDay {
	s.r.Set("timeZone", timeZone)
	return s
}

// Type Supported values: FULL or MINI.
// If none provided, the default is FULL
func (s *WsTickerTradingDay) Type(type_ core.TickerTypeEnum) *WsTickerTradingDay {
	s.r.Set("type", type_)
	return s
}

func (s *WsTickerTradingDay) Do(ctx context.Context) (*WsTickerTradingDayResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(WsTickerTradingDayResponse)
			if s.r.Get("symbols") != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(WsTickerTradingDaySingleResponse)
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

// WsTicker Get rolling window price change statistics with a custom window.
// This request is similar to ticker.24hr, but statistics are computed on demand using the arbitrary window you specify.
type WsTicker struct {
	c *WsClient
	r *core.WsRequest
}
type TickerResult struct {
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
	FirstId            int64           `json:"firstId"`
	LastId             int64           `json:"lastId"`
	Count              int             `json:"count"`
}

type TickerSingleResponse struct {
	ApiResponse
	Result *TickerResult `json:"result"`
}

type WsTickerResponse struct {
	ApiResponse
	Result []*TickerResult `json:"result"`
}

func (s *WsTicker) Symbol(symbol string) *WsTicker {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsTicker) Symbols(symbols []string) *WsTicker {
	s.r.Set("symbols", symbols)
	return s
}

func (s *WsTicker) Type(type_ core.TickerTypeEnum) *WsTicker {
	s.r.Set("type", type_)
	return s
}

// WindowSize Supported window sizes:
// minutes	1m, 2m ... 59m
// hours	1h, 2h ... 23h
// days	1d, 2d ... 7d
func (s *WsTicker) WindowSize(windowSize string) *WsTicker {
	s.r.Set("windowSize", windowSize)
	return s
}

func (s *WsTicker) Do(ctx context.Context) (*WsTickerResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(WsTickerResponse)
			if s.r.Get("symbols") != nil {
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

// WsTickerPrice Get the latest market price for a symbol.
type WsTickerPrice struct {
	c *WsClient
	r *core.WsRequest
}

type TickerPriceResult struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
}

type WsTickerPriceSingleResponse struct {
	ApiResponse
	Result *TickerPriceResult `json:"result"`
}

type WsTickerPriceResponse struct {
	ApiResponse
	Result []*TickerPriceResult `json:"result"`
}

func (s *WsTickerPrice) Symbol(symbol string) *WsTickerPrice {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsTickerPrice) Symbols(symbols []string) *WsTickerPrice {
	s.r.Set("symbols", symbols)
	return s
}

func (s *WsTickerPrice) Do(ctx context.Context) (*WsTickerPriceResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(WsTickerPriceResponse)
			if s.r.Get("symbols") != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(WsTickerPriceSingleResponse)
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

// WsTickerBook Get the current best price and quantity on the order book.
type WsTickerBook struct {
	c *WsClient
	r *core.WsRequest
}

type TickerBookResult struct {
	Symbol   string          `json:"symbol"`
	BidPrice decimal.Decimal `json:"bidPrice"`
	BidQty   decimal.Decimal `json:"bidQty"`
	AskPrice decimal.Decimal `json:"askPrice"`
	AskQty   decimal.Decimal `json:"askQty"`
}

type WsTickerBookSingleResponse struct {
	ApiResponse
	Result *TickerBookResult `json:"result"`
}

type WsTickerBookResponse struct {
	ApiResponse
	Result []*TickerBookResult `json:"result"`
}

func (s *WsTickerBook) Symbol(symbol string) *WsTickerBook {
	s.r.Set("symbol", symbol)
	return s
}

func (s *WsTickerBook) Symbols(symbols []string) *WsTickerBook {
	s.r.Set("symbols", symbols)
	return s
}

func (s *WsTickerBook) Do(ctx context.Context) (*WsTickerBookResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			resp := new(WsTickerBookResponse)
			if s.r.Get("symbols") != nil {
				return resp, json.Unmarshal(message, &resp)
			}
			single := new(WsTickerBookSingleResponse)
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
