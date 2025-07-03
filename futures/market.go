package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int    `json:"limit"`
}

// ExchangeFilter define exchange filter
type ExchangeFilter struct {
	FilterType       string `json:"filterType"`
	MaxNumAlgoOrders int64  `json:"maxNumAlgoOrders"`
}

// Ping Test connectivity to the Rest API.
type Ping struct {
	c *Client
	r *core.Request
}

func (s *Ping) Do(ctx context.Context) error {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return err
	}
	return nil
}

type ServerTime struct {
	c *Client
	r *core.Request
}

type ServerTimeResponse struct {
	ServerTime int64 `json:"serverTime"`
}

func (s *ServerTime) Do(ctx context.Context) (*ServerTimeResponse, error) {
	resp := new(ServerTimeResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ExchangeInfo Current exchange trading rules and symbol information
type ExchangeInfo struct {
	c *Client
	r *core.Request
}

type SymbolFilter struct {
	FilterType        string          `json:"filterType"`
	MaxPrice          decimal.Decimal `json:"maxPrice,omitempty"`
	MinPrice          decimal.Decimal `json:"minPrice,omitempty"`
	TickSize          decimal.Decimal `json:"tickSize,omitempty"`
	MaxQty            decimal.Decimal `json:"maxQty,omitempty"`
	MinQty            decimal.Decimal `json:"minQty,omitempty"`
	StepSize          string          `json:"stepSize,omitempty"`
	Limit             int             `json:"limit,omitempty"`
	Notional          decimal.Decimal `json:"notional,omitempty"`
	MultiplierUp      decimal.Decimal `json:"multiplierUp,omitempty"`
	MultiplierDown    decimal.Decimal `json:"multiplierDown,omitempty"`
	MultiplierDecimal string          `json:"multiplierDecimal,omitempty"`
}

type SymbolInfo struct {
	Symbol                string          `json:"symbol"`
	Pair                  string          `json:"pair"`
	ContractType          string          `json:"contractType"`
	DeliveryDate          int64           `json:"deliveryDate"`
	OnboardDate           int64           `json:"onboardDate"`
	Status                string          `json:"status"`
	MaintMarginPercent    decimal.Decimal `json:"maintMarginPercent"`
	RequiredMarginPercent decimal.Decimal `json:"requiredMarginPercent"`
	BaseAsset             string          `json:"baseAsset"`
	QuoteAsset            string          `json:"quoteAsset"`
	MarginAsset           string          `json:"marginAsset"`
	PricePrecision        int             `json:"pricePrecision"`
	QuantityPrecision     int             `json:"quantityPrecision"`
	BaseAssetPrecision    int             `json:"baseAssetPrecision"`
	QuotePrecision        int             `json:"quotePrecision"`
	UnderlyingType        string          `json:"underlyingType"`
	UnderlyingSubType     []string        `json:"underlyingSubType"`
	SettlePlan            int             `json:"settlePlan"`
	TriggerProtect        decimal.Decimal `json:"triggerProtect"`
	Filters               []*SymbolFilter `json:"filters"`
	OrderType             []string        `json:"OrderType"`
	TimeInForce           []string        `json:"timeInForce"`
	LiquidationFee        string          `json:"liquidationFee"`
	MarketTakeBound       string          `json:"marketTakeBound"`
}

type ExchangeInfoResponse struct {
	Timezone        string            `json:"timezone"`
	ServerTime      int64             `json:"serverTime"`
	RateLimits      []*RateLimit      `json:"rateLimits"`
	ExchangeFilters []*ExchangeFilter `json:"exchangeFilters"`
	Symbols         []*SymbolInfo     `json:"symbols"`
}

func (s *ExchangeInfo) Do(ctx context.Context) (*ExchangeInfoResponse, error) {
	resp := new(ExchangeInfoResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// Depth Get depth of a market
type Depth struct {
	c *Client
	r *core.Request
}

type DepthResponse struct {
	LastUpdateId    int                 `json:"lastUpdateId"`
	OutputTime      int64               `json:"E"`
	TransactionTime int64               `json:"T"`
	Bids            [][]decimal.Decimal `json:"bids"` // first is PRICE, second is QTY
	Asks            [][]decimal.Decimal `json:"asks"`
}

func (s *Depth) Symbol(symbol string) *Depth {
	s.r.Set("symbol", symbol)
	return s
}

// Limit Default 500; max 1000. If limit > 1000, then the response will truncate to 1000.
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
	TradeId   int             `json:"a"`
	Price     decimal.Decimal `json:"p"`
	Quantity  decimal.Decimal `json:"q"`
	FirstId   int             `json:"f"`
	LastId    int             `json:"l"`
	Timestamp int64           `json:"T"`
	IsMaker   bool            `json:"m"`
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

type KlineDataResponse struct {
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

// Limit Default 500; max 1000.
func (s *KlineData) Limit(limit int) *KlineData {
	s.r.Set("limit", limit)
	return s
}

func (s *KlineData) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.rawBody())
}

func parseKlineData(rawBody []byte) ([]*KlineDataResponse, error) {
	resp := make([]*KlineDataResponse, 0)
	res := make([][]any, 0)
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
	return resp, nil
}

// ContractKline Kline/candlestick bars for a specific contract type. Klines are uniquely identified by their open time.
type ContractKline struct {
	c *Client
	r *core.Request
}

func (s *ContractKline) Pair(pair string) *ContractKline {
	s.r.Set("pair", pair)
	return s
}

func (s *ContractKline) ContractType(contractType core.ContractType) *ContractKline {
	s.r.Set("contractType", contractType)
	return s
}

func (s *ContractKline) Interval(interval core.IntervalEnum) *ContractKline {
	s.r.Set("interval", interval)
	return s
}

func (s *ContractKline) StartTime(startTime int64) *ContractKline {
	s.r.Set("startTime", startTime)
	return s
}

func (s *ContractKline) EndTime(endTime int64) *ContractKline {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500; max 1000.
func (s *ContractKline) Limit(limit int) *ContractKline {
	s.r.Set("limit", limit)
	return s
}

func (s *ContractKline) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.rawBody())
}

// IndexKline Kline/candlestick bars for the index price of a pair. Klines are uniquely identified by their open time.
type IndexKline struct {
	c *Client
	r *core.Request
}

func (s *IndexKline) Pair(pair string) *IndexKline {
	s.r.Set("pair", pair)
	return s
}

func (s *IndexKline) Interval(interval core.IntervalEnum) *IndexKline {
	s.r.Set("interval", interval)
	return s
}

func (s *IndexKline) StartTime(startTime int64) *IndexKline {
	s.r.Set("startTime", startTime)
	return s
}

func (s *IndexKline) EndTime(endTime int64) *IndexKline {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500; max 1000.
func (s *IndexKline) Limit(limit int) *IndexKline {
	s.r.Set("limit", limit)
	return s
}

func (s *IndexKline) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.rawBody())
}

// MarkKline Kline/candlestick bars for the mark price of a symbol. Klines are uniquely identified by their open time.
type MarkKline struct {
	c *Client
	r *core.Request
}

func (s *MarkKline) Symbol(symbol string) *MarkKline {
	s.r.Set("symbol", symbol)
	return s
}

func (s *MarkKline) Interval(interval core.IntervalEnum) *MarkKline {
	s.r.Set("interval", interval)
	return s
}

func (s *MarkKline) StartTime(startTime int64) *MarkKline {
	s.r.Set("startTime", startTime)
	return s
}

func (s *MarkKline) EndTime(endTime int64) *MarkKline {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500; max 1000.
func (s *MarkKline) Limit(limit int) *MarkKline {
	s.r.Set("limit", limit)
	return s
}

func (s *MarkKline) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.rawBody())
}

// PremiumKline Premium index kline bars of a symbol. Klines are uniquely identified by their open time.
type PremiumKline struct {
	c *Client
	r *core.Request
}

func (s *PremiumKline) Symbol(symbol string) *PremiumKline {
	s.r.Set("symbol", symbol)
	return s
}

func (s *PremiumKline) Interval(interval core.IntervalEnum) *PremiumKline {
	s.r.Set("interval", interval)
	return s
}

func (s *PremiumKline) StartTime(startTime int64) *PremiumKline {
	s.r.Set("startTime", startTime)
	return s
}

func (s *PremiumKline) EndTime(endTime int64) *PremiumKline {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 500; max 1000.
func (s *PremiumKline) Limit(limit int) *PremiumKline {
	s.r.Set("limit", limit)
	return s
}

func (s *PremiumKline) Do(ctx context.Context) ([]*KlineDataResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	return parseKlineData(s.c.rawBody())
}

// MarkPrice Mark Price and Funding Rate
type MarkPrice struct {
	c *Client
	r *core.Request
}

type MarkPriceResponse struct {
	Symbol               string `json:"symbol"`
	MarkPrice            string `json:"markPrice"`
	IndexPrice           string `json:"indexPrice"`
	EstimatedSettlePrice string `json:"estimatedSettlePrice"`
	LastFundingRate      string `json:"lastFundingRate"`
	InterestRate         string `json:"interestRate"`
	NextFundingTime      int64  `json:"nextFundingTime"`
	Time                 int64  `json:"time"`
}

func (s *MarkPrice) Symbol(symbol string) *MarkPrice {
	s.r.Set("symbol", symbol)
	return s
}

func (s *MarkPrice) Do(ctx context.Context) ([]*MarkPriceResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*MarkPriceResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(MarkPriceResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// FundingRate Get Funding Rate History
type FundingRate struct {
	c *Client
	r *core.Request
}

type FundingRateResponse struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
	FundingTime int64  `json:"fundingTime"`
	MarkPrice   string `json:"markPrice"`
}

func (s *FundingRate) Symbol(symbol string) *FundingRate {
	s.r.Set("symbol", symbol)
	return s
}

func (s *FundingRate) StartTime(startTime int64) *FundingRate {
	s.r.Set("startTime", startTime)
	return s
}

func (s *FundingRate) EndTime(endTime int64) *FundingRate {
	s.r.Set("endTime", endTime)
	return s
}

// Limit Default 100; max 1000.
func (s *FundingRate) Limit(limit int) *FundingRate {
	s.r.Set("limit", limit)
	return s
}

func (s *FundingRate) Do(ctx context.Context) ([]*FundingRateResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*FundingRateResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// FundingInfo Query funding rate info for symbols that had FundingRateCap/ FundingRateFloor / fundingIntervalHours adjustment
// 0 share 500/5min/IP rate limit with GET /fapi/v1/fundingInfo
type FundingInfo struct {
	c *Client
	r *core.Request
}

type FundingInfoResponse struct {
	Symbol                   string `json:"symbol"`
	AdjustedFundingRateCap   string `json:"adjustedFundingRateCap"`
	AdjustedFundingRateFloor string `json:"adjustedFundingRateFloor"`
	FundingIntervalHours     int    `json:"fundingIntervalHours"`
	Disclaimer               bool   `json:"disclaimer"`
}

func (s *FundingInfo) Do(ctx context.Context) ([]*FundingInfoResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*FundingInfoResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// Ticker24hr 24 hour rolling window price change statistics.
// Careful when accessing this with no symbol.
type Ticker24hr struct {
	c *Client
	r *core.Request
}

func (s *Ticker24hr) Symbol(symbol string) *Ticker24hr {
	s.r.Set("symbol", symbol)
	return s
}

type TickerStatisticsResponse struct {
	Symbol             string `json:"symbol"`
	PriceChange        string `json:"priceChange"`
	PriceChangePercent string `json:"priceChangePercent"`
	WeightedAvgPrice   string `json:"weightedAvgPrice"`
	LastPrice          string `json:"lastPrice"`
	LastQty            string `json:"lastQty"`
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

func (s *Ticker24hr) Do(ctx context.Context) ([]*TickerStatisticsResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*TickerStatisticsResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(TickerStatisticsResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// TickerPrice Latest price for a symbol or symbols.
type TickerPrice struct {
	c *Client
	r *core.Request
}

type TickerPriceResponse struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
	Time   int64  `json:"time"`
}

func (s *TickerPrice) Symbol(symbol string) *TickerPrice {
	s.r.Set("symbol", symbol)
	return s
}

func (s *TickerPrice) Do(ctx context.Context) ([]*TickerPriceResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*TickerPriceResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(TickerPriceResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// BookTicker Best price/qty on the order book for a symbol or symbols.
type BookTicker struct {
	c *Client
	r *core.Request
}

type BookTickerResponse struct {
	Symbol   string          `json:"symbol"`
	BidPrice decimal.Decimal `json:"bidPrice"`
	BidQty   decimal.Decimal `json:"bidQty"`
	AskPrice decimal.Decimal `json:"askPrice"`
	AskQty   decimal.Decimal `json:"askQty"`
	Time     int64           `json:"time"`
}

func (s *BookTicker) Symbol(symbol string) *BookTicker {
	s.r.Set("symbol", symbol)
	return s
}

func (s *BookTicker) Do(ctx context.Context) ([]*BookTickerResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*BookTickerResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(BookTickerResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// DeliveryPrice Quarterly Contract Settlement Price
type DeliveryPrice struct {
	c *Client
	r *core.Request
}

type DeliveryPriceResponse struct {
	DeliveryTime  int64   `json:"deliveryTime"`
	DeliveryPrice float64 `json:"deliveryPrice"`
}

func (s *DeliveryPrice) Symbol(symbol string) *DeliveryPrice {
	s.r.Set("pair", symbol)
	return s
}

func (s *DeliveryPrice) Do(ctx context.Context) ([]*DeliveryPriceResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*DeliveryPriceResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OpenInterest Get present open interest of a specific symbol.
type OpenInterest struct {
	c *Client
	r *core.Request
}

type OpenInterestResponse struct {
	OpenInterest decimal.Decimal `json:"openInterest"`
	Symbol       string          `json:"symbol"`
	Time         int64           `json:"time"`
}

func (s *OpenInterest) Symbol(symbol string) *OpenInterest {
	s.r.Set("symbol", symbol)
	return s
}

func (s *OpenInterest) Do(ctx context.Context) (*OpenInterestResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(OpenInterestResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// OpenInterestHist Open Interest Statistics
type OpenInterestHist struct {
	c *Client
	r *core.Request
}

type OpenInterestHistResponse struct {
	Symbol               string          `json:"symbol"`
	SumOpenInterest      decimal.Decimal `json:"sumOpenInterest"`
	SumOpenInterestValue decimal.Decimal `json:"sumOpenInterestValue"`
	Timestamp            int64           `json:"timestamp"`
}

func (s *OpenInterestHist) Symbol(symbol string) *OpenInterestHist {
	s.r.Set("symbol", symbol)
	return s
}

// Period "5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (s *OpenInterestHist) Period(period core.IntervalEnum) *OpenInterestHist {
	s.r.Set("period", period)
	return s
}

// Limit default 30, max 500
func (s *OpenInterestHist) Limit(limit int) *OpenInterestHist {
	s.r.Set("limit", limit)
	return s
}

func (s *OpenInterestHist) StartTime(startTime int64) *OpenInterestHist {
	s.r.Set("startTime", startTime)
	return s
}

func (s *OpenInterestHist) EndTime(endTime int64) *OpenInterestHist {
	s.r.Set("endTime", endTime)
	return s
}

func (s *OpenInterestHist) Do(ctx context.Context) ([]*OpenInterestHistResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*OpenInterestHistResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TopTraderPositionsRatio The proportion of net long and net short positions to total open positions of the top 20% users with the highest margin balance.
// Long Position % = Long positions of top traders / Total open positions of top traders Short Position % = Short positions of top traders / Total open positions of top traders Long/Short Ratio (Positions) = Long Position % / Short Position %
type TopTraderPositionsRatio struct {
	c *Client
	r *core.Request
}

type TopTraderRatioResponse struct {
	Symbol         string          `json:"symbol"`
	LongShortRatio decimal.Decimal `json:"longShortRatio"`
	LongAccount    decimal.Decimal `json:"longAccount"`
	ShortAccount   decimal.Decimal `json:"shortAccount"`
	Timestamp      int64           `json:"timestamp"`
}

func (s *TopTraderPositionsRatio) Symbol(symbol string) *TopTraderPositionsRatio {
	s.r.Set("symbol", symbol)
	return s
}

// Period "5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (s *TopTraderPositionsRatio) Period(period core.IntervalEnum) *TopTraderPositionsRatio {
	s.r.Set("period", period)
	return s
}

// Limit default 30, max 500
func (s *TopTraderPositionsRatio) Limit(limit int) *TopTraderPositionsRatio {
	s.r.Set("limit", limit)
	return s
}

func (s *TopTraderPositionsRatio) StartTime(startTime int64) *TopTraderPositionsRatio {
	s.r.Set("startTime", startTime)
	return s
}

func (s *TopTraderPositionsRatio) EndTime(endTime int64) *TopTraderPositionsRatio {
	s.r.Set("endTime", endTime)
	return s
}

func (s *TopTraderPositionsRatio) Do(ctx context.Context) ([]*TopTraderRatioResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*TopTraderRatioResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TopTraderAccountsRatio The proportion of net long and net short accounts to total accounts of the top 20% users with the highest margin balance.
type TopTraderAccountsRatio struct {
	c *Client
	r *core.Request
}

func (s *TopTraderAccountsRatio) Symbol(symbol string) *TopTraderAccountsRatio {
	s.r.Set("symbol", symbol)
	return s
}

// Period "5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (s *TopTraderAccountsRatio) Period(period core.IntervalEnum) *TopTraderAccountsRatio {
	s.r.Set("period", period)
	return s
}

// Limit default 30, max 500
func (s *TopTraderAccountsRatio) Limit(limit int) *TopTraderAccountsRatio {
	s.r.Set("limit", limit)
	return s
}

func (s *TopTraderAccountsRatio) StartTime(startTime int64) *TopTraderAccountsRatio {
	s.r.Set("startTime", startTime)
	return s
}

func (s *TopTraderAccountsRatio) EndTime(endTime int64) *TopTraderAccountsRatio {
	s.r.Set("endTime", endTime)
	return s
}

func (s *TopTraderAccountsRatio) Do(ctx context.Context) ([]*TopTraderRatioResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*TopTraderRatioResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// SymbolRatio Query symbol Long/Short Ratio
type SymbolRatio struct {
	c *Client
	r *core.Request
}

func (s *SymbolRatio) Symbol(symbol string) *SymbolRatio {
	s.r.Set("symbol", symbol)
	return s
}

// Period "5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (s *SymbolRatio) Period(period core.IntervalEnum) *SymbolRatio {
	s.r.Set("period", period)
	return s
}

// Limit default 30, max 500
func (s *SymbolRatio) Limit(limit int) *SymbolRatio {
	s.r.Set("limit", limit)
	return s
}

func (s *SymbolRatio) StartTime(startTime int64) *SymbolRatio {
	s.r.Set("startTime", startTime)
	return s
}

func (s *SymbolRatio) EndTime(endTime int64) *SymbolRatio {
	s.r.Set("endTime", endTime)
	return s
}

func (s *SymbolRatio) Do(ctx context.Context) ([]*TopTraderRatioResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*TopTraderRatioResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TakerVolume Taker Buy/Sell Volume
type TakerVolume struct {
	c *Client
	r *core.Request
}

type TakerVolumeResponse struct {
	BuySellRatio decimal.Decimal `json:"buySellRatio"`
	BuyVol       decimal.Decimal `json:"buyVol"`
	SellVol      decimal.Decimal `json:"sellVol"`
	Timestamp    int64           `json:"timestamp"`
}

func (s *TakerVolume) Symbol(symbol string) *TakerVolume {
	s.r.Set("symbol", symbol)
	return s
}

// Period "5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (s *TakerVolume) Period(period core.IntervalEnum) *TakerVolume {
	s.r.Set("period", period)
	return s
}

// Limit default 30, max 500
func (s *TakerVolume) Limit(limit int) *TakerVolume {
	s.r.Set("limit", limit)
	return s
}

func (s *TakerVolume) StartTime(startTime int64) *TakerVolume {
	s.r.Set("startTime", startTime)
	return s
}

func (s *TakerVolume) EndTime(endTime int64) *TakerVolume {
	s.r.Set("endTime", endTime)
	return s
}

func (s *TakerVolume) Do(ctx context.Context) ([]*TakerVolumeResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*TakerVolumeResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// FutureBasis Query future basis
type FutureBasis struct {
	c *Client
	r *core.Request
}

type FutureBasisResponse struct {
	IndexPrice          decimal.Decimal `json:"indexPrice"`
	ContractType        string          `json:"contractType"`
	BasisRate           decimal.Decimal `json:"basisRate"`
	FuturesPrice        decimal.Decimal `json:"futuresPrice"`
	AnnualizedBasisRate string          `json:"annualizedBasisRate"`
	Basis               decimal.Decimal `json:"basis"`
	Pair                string          `json:"pair"`
	Timestamp           int64           `json:"timestamp"`
}

func (s *FutureBasis) Symbol(symbol string) *FutureBasis {
	s.r.Set("symbol", symbol)
	return s
}

func (s *FutureBasis) ContractType(contractType core.ContractType) *FutureBasis {
	s.r.Set("contractType", contractType)
	return s
}

// Period "5m","15m","30m","1h","2h","4h","6h","12h","1d"
func (s *FutureBasis) Period(period core.IntervalEnum) *FutureBasis {
	s.r.Set("period", period)
	return s
}

// Limit default 30, max 500
func (s *FutureBasis) Limit(limit int) *FutureBasis {
	s.r.Set("limit", limit)
	return s
}

func (s *FutureBasis) StartTime(startTime int64) *FutureBasis {
	s.r.Set("startTime", startTime)
	return s
}

func (s *FutureBasis) EndTime(endTime int64) *FutureBasis {
	s.r.Set("endTime", endTime)
	return s
}

func (s *FutureBasis) Do(ctx context.Context) ([]*FutureBasisResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*FutureBasisResponse, 0)
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// IndexInfo Query composite index symbol information
type IndexInfo struct {
	c *Client
	r *core.Request
}

type IndexInfoResponse struct {
	Symbol        string       `json:"symbol"`
	Time          int64        `json:"time"`
	Component     string       `json:"component"`
	BaseAssetList []*BaseAsset `json:"baseAssetList"`
}

type BaseAsset struct {
	BaseAsset          string          `json:"baseAsset"`
	QuoteAsset         string          `json:"quoteAsset"`
	WeightInQuantity   decimal.Decimal `json:"weightInQuantity"`
	WeightInPercentage decimal.Decimal `json:"weightInPercentage"`
}

func (s *IndexInfo) Symbol(symbol string) *IndexInfo {
	s.r.Set("symbol", symbol)
	return s
}

func (s *IndexInfo) Do(ctx context.Context) ([]*IndexInfoResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*IndexInfoResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(IndexInfoResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// AssetIndex asset index for Multi-Assets mode
type AssetIndex struct {
	c *Client
	r *core.Request
}

type AssetIndexResponse struct {
	Symbol                string          `json:"symbol"`
	Time                  int64           `json:"time"`
	Index                 decimal.Decimal `json:"index"`
	BidBuffer             decimal.Decimal `json:"bidBuffer"`
	AskBuffer             decimal.Decimal `json:"askBuffer"`
	BidRate               decimal.Decimal `json:"bidRate"`
	AskRate               decimal.Decimal `json:"askRate"`
	AutoExchangeBidBuffer decimal.Decimal `json:"autoExchangeBidBuffer"`
	AutoExchangeAskBuffer decimal.Decimal `json:"autoExchangeAskBuffer"`
	AutoExchangeBidRate   decimal.Decimal `json:"autoExchangeBidRate"`
	AutoExchangeAskRate   decimal.Decimal `json:"autoExchangeAskRate"`
}

func (s *AssetIndex) Symbol(symbol string) *AssetIndex {
	s.r.Set("symbol", symbol)
	return s
}

func (s *AssetIndex) Do(ctx context.Context) ([]*AssetIndexResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*AssetIndexResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(AssetIndexResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// ConstituentsPrice Query index price constituents
type ConstituentsPrice struct {
	c *Client
	r *core.Request
}

type ConstituentResponse struct {
	Exchange string          `json:"exchange"`
	Symbol   string          `json:"symbol"`
	Price    decimal.Decimal `json:"price"`
	Weight   decimal.Decimal `json:"weight"`
}

type ConstituentsPriceResponse struct {
	Symbol       string                 `json:"symbol"`
	Time         int64                  `json:"time"`
	Constituents []*ConstituentResponse `json:"constituents"`
}

func (s *ConstituentsPrice) Symbol(symbol string) *ConstituentsPrice {
	s.r.Set("symbol", symbol)
	return s
}

func (s *ConstituentsPrice) Do(ctx context.Context) (*ConstituentsPriceResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := new(ConstituentsPriceResponse)
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// InsuranceBalance Query Insurance Fund Balance Snapshot
type InsuranceBalance struct {
	c *Client
	r *core.Request
}

type InsuranceBalanceAsset struct {
	Asset         string `json:"asset"`
	MarginBalance string `json:"marginBalance"`
	UpdateTime    int64  `json:"updateTime"`
}

type InsuranceBalanceResponse struct {
	Symbols []string                `json:"symbols"`
	Assets  []InsuranceBalanceAsset `json:"assets"`
}

func (s *InsuranceBalance) Symbol(symbol string) *InsuranceBalance {
	s.r.Set("symbol", symbol)
	return s
}

func (s *InsuranceBalance) Do(ctx context.Context) ([]*InsuranceBalanceResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	resp := make([]*InsuranceBalanceResponse, 0)
	if s.r.GetQuery("symbol") == "" {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	res := new(InsuranceBalanceResponse)
	if err := json.Unmarshal(s.c.rawBody(), res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}
