package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// RateLimit define rate limit
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         uint   `json:"limit"`
}

// ExchangeFilter define exchange filter
type ExchangeFilter struct {
	FilterType       string `json:"filterType"`
	MaxNumAlgoOrders int64  `json:"maxNumAlgoOrders"`
}

// SymbolInfo define symbol
type SymbolInfo struct {
	Symbol                          string          `json:"symbol"`
	Status                          string          `json:"status"`
	BaseAsset                       string          `json:"baseAsset"`
	BaseAssetPrecision              int64           `json:"baseAssetPrecision"`
	QuoteAsset                      string          `json:"quoteAsset"`
	QuotePrecision                  int64           `json:"quotePrecision"`
	QuoteAssetPrecision             int64           `json:"quoteAssetPrecision"`
	OrderTypes                      []string        `json:"orderTypes"`
	IcebergAllowed                  bool            `json:"icebergAllowed"`
	OcoAllowed                      bool            `json:"ocoAllowed"`
	QuoteOrderQtyMarketAllowed      bool            `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool            `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool            `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool            `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool            `json:"isMarginTradingAllowed"`
	Filters                         []*SymbolFilter `json:"filters"`
	Permissions                     []string        `json:"permissions"`
	PermissionSets                  [][]string      `json:"permissionSets"`
	DefaultSelfTradePreventionMode  string          `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string        `json:"allowedSelfTradePreventionModes"`
}

// SymbolFilter define symbol filter
type SymbolFilter struct {
	ApplyMinToMarket      bool             `json:"applyMinToMarket,omitempty"`
	ApplyMaxToMarket      bool             `json:"applyMaxToMarket,omitempty"`
	AskMultiplierDown     string           `json:"askMultiplierDown,omitempty"`
	AskMultiplierUp       string           `json:"askMultiplierUp,omitempty"`
	AvgPriceMins          int64            `json:"avgPriceMins,omitempty"`
	BidMultiplierDown     string           `json:"bidMultiplierDown,omitempty"`
	BidMultiplierUp       string           `json:"bidMultiplierUp,omitempty"`
	FilterType            string           `json:"filterType,omitempty"`
	Limit                 uint             `json:"limit,omitempty"`
	MaxNotional           string           `json:"maxNotional,omitempty"`
	MaxNumAlgoOrders      int64            `json:"maxNumAlgoOrders,omitempty"`
	MaxNumOrders          int64            `json:"maxNumOrders,omitempty"`
	MaxPrice              string           `json:"maxPrice,omitempty"`
	MaxQty                string           `json:"maxQty,omitempty"`
	MaxTrailingAboveDelta int64            `json:"maxTrailingAboveDelta,omitempty"`
	MaxTrailingBelowDelta int64            `json:"maxTrailingBelowDelta,omitempty"`
	MinNotional           string           `json:"minNotional,omitempty"`
	MinPrice              *decimal.Decimal `json:"minPrice,omitempty"`
	MinQty                *decimal.Decimal `json:"minQty,omitempty"`
	MinTrailingAboveDelta int64            `json:"minTrailingAboveDelta,omitempty"`
	MinTrailingBelowDelta int64            `json:"minTrailingBelowDelta,omitempty"`
	StepSize              string           `json:"stepSize,omitempty"`
	TickSize              *decimal.Decimal `json:"tickSize,omitempty"`
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

// ServerTime Test connectivity to the Rest API and get the current server time.
type ServerTime struct {
	c *Client
	r *core.Request
}

type ServerTimeResponse struct {
	ServerTime uint64 `json:"serverTime"`
}

func (s *ServerTime) Do(ctx context.Context) (*ServerTimeResponse, error) {
	var resp ServerTimeResponse
	if err := s.c.invoke(s.r, ctx); err != nil {
		return &resp, err
	}
	return &resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ExchangeInfo Current exchange trading rules and symbol information
type ExchangeInfo struct {
	c *Client
	r *core.Request
}

type ExchangeInfoResponse struct {
	Timezone        string            `json:"timezone"`
	ServerTime      uint64            `json:"serverTime"`
	RateLimits      []*RateLimit      `json:"rateLimits"`
	ExchangeFilters []*ExchangeFilter `json:"exchangeFilters"`
	Symbols         []*SymbolInfo     `json:"symbols"`
}

func (s *ExchangeInfo) Symbol(symbol string) *ExchangeInfo {
	s.r.Set("symbol", symbol)
	return s
}

func (s *ExchangeInfo) Symbols(symbols []string) *ExchangeInfo {
	s.r.Set("symbols", symbols)
	return s
}

// Permissions can support single or multiple values (e.g. SPOT, ["MARGIN","LEVERAGED"]). This cannot be used in combination with symbol or symbols.
func (s *ExchangeInfo) Permissions(permissions []core.PermissionEnum) *ExchangeInfo {
	s.r.Set("permissions", permissions)
	return s
}

// ShowPermissionSets Controls whether the content of the permissionSets field is populated or not. Defaults to true
func (s *ExchangeInfo) ShowPermissionSets(showPermissionSets bool) *ExchangeInfo {
	s.r.Set("showPermissionSets", showPermissionSets)
	return s
}

// SymbolStatus Filters symbols that have this tradingStatus. Valid values: TRADING, HALT, BREAK
// Cannot be used in combination with symbols or symbol.
func (s *ExchangeInfo) SymbolStatus(symbolStatus core.SymbolStatusEnum) *ExchangeInfo {
	s.r.Set("symbolStatus", symbolStatus)
	return s
}

func (s *ExchangeInfo) Do(ctx context.Context) (*ExchangeInfoResponse, error) {
	resp := new(ExchangeInfoResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}
