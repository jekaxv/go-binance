package http

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/types"
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
	ApplyMinToMarket      bool   `json:"applyMinToMarket"`
	ApplyMaxToMarket      bool   `json:"applyMaxToMarket"`
	AskMultiplierDown     string `json:"askMultiplierDown"`
	AskMultiplierUp       string `json:"askMultiplierUp"`
	AvgPriceMins          int64  `json:"avgPriceMins"`
	BidMultiplierDown     string `json:"bidMultiplierDown"`
	BidMultiplierUp       string `json:"bidMultiplierUp"`
	FilterType            string `json:"filterType"`
	Limit                 uint   `json:"limit"`
	MaxNotional           string `json:"maxNotional"`
	MaxNumAlgoOrders      int64  `json:"maxNumAlgoOrders"`
	MaxNumOrders          int64  `json:"maxNumOrders"`
	MaxPrice              string `json:"maxPrice"`
	MaxQty                string `json:"maxQty"`
	MaxTrailingAboveDelta int64  `json:"maxTrailingAboveDelta"`
	MaxTrailingBelowDelta int64  `json:"maxTrailingBelowDelta"`
	MinNotional           string `json:"minNotional"`
	MinPrice              string `json:"minPrice"`
	MinQty                string `json:"minQty"`
	MinTrailingAboveDelta int64  `json:"minTrailingAboveDelta"`
	MinTrailingBelowDelta int64  `json:"minTrailingBelowDelta"`
	StepSize              string `json:"stepSize"`
	TickSize              string `json:"tickSize"`
}

// Ping Test connectivity to the Rest API.
type Ping struct {
	c *Client
}

func (s *Ping) Do(ctx context.Context) error {
	if err := s.c.invoke(ctx); err != nil {
		return err
	}
	return nil
}

// ServerTime Test connectivity to the Rest API and get the current server time.
type ServerTime struct {
	c *Client
}

type ServerTimeResponse struct {
	ServerTime uint64 `json:"serverTime"`
}

func (s *ServerTime) Do(ctx context.Context) (*ServerTimeResponse, error) {
	var resp ServerTimeResponse
	if err := s.c.invoke(ctx); err != nil {
		return &resp, err
	}
	return &resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// ExchangeInfo Current exchange trading rules and symbol information
type ExchangeInfo struct {
	c                  *Client
	symbol             *string
	symbols            []string
	permissions        []types.PermissionEnum
	showPermissionSets *bool
	symbolStatus       *types.SymbolStatusEnum
}

type ExchangeInfoResponse struct {
	Timezone        string            `json:"timezone"`
	ServerTime      uint64            `json:"serverTime"`
	RateLimits      []*RateLimit      `json:"rateLimits"`
	ExchangeFilters []*ExchangeFilter `json:"exchangeFilters"`
	Symbols         []*SymbolInfo     `json:"symbols"`
}

func (s *ExchangeInfo) Symbol(symbol string) *ExchangeInfo {
	s.symbol = &symbol
	return s
}

func (s *ExchangeInfo) Symbols(symbols []string) *ExchangeInfo {
	s.symbols = symbols
	return s
}

// Permissions can support single or multiple values (e.g. SPOT, ["MARGIN","LEVERAGED"]). This cannot be used in combination with symbol or symbols.
func (s *ExchangeInfo) Permissions(permissions []types.PermissionEnum) *ExchangeInfo {
	s.permissions = permissions
	return s
}

// ShowPermissionSets Controls whether the content of the permissionSets field is populated or not. Defaults to true
func (s *ExchangeInfo) ShowPermissionSets(showPermissionSets bool) *ExchangeInfo {
	s.showPermissionSets = &showPermissionSets
	return s
}

// SymbolStatus Filters symbols that have this tradingStatus. Valid values: TRADING, HALT, BREAK
// Cannot be used in combination with symbols or symbol.
func (s *ExchangeInfo) SymbolStatus(symbolStatus types.SymbolStatusEnum) *ExchangeInfo {
	s.symbolStatus = &symbolStatus
	return s
}

func (s *ExchangeInfo) Do(ctx context.Context) (*ExchangeInfoResponse, error) {
	var resp ExchangeInfoResponse
	if s.symbol != nil {
		s.c.req.set("symbol", *s.symbol)
	}
	if len(s.symbols) != 0 {
		s.c.req.set("symbols", s.symbols)
	}
	if len(s.permissions) == 1 {
		s.c.req.set("permissions", s.permissions[0])
	} else if len(s.permissions) >= 2 {
		s.c.req.set("permissions", s.permissions)
	}
	if s.showPermissionSets != nil {
		s.c.req.set("showPermissionSets", *s.showPermissionSets)
	}
	if s.symbolStatus != nil {
		s.c.req.set("symbolStatus", *s.symbolStatus)
	}
	if err := s.c.invoke(ctx); err != nil {
		return &resp, err
	}
	return &resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}
