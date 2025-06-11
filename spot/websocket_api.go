package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

type ApiRateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

type ApiFilter struct {
	FilterType string          `json:"filterType"`
	MinPrice   decimal.Decimal `json:"minPrice,omitempty"`
	MaxPrice   decimal.Decimal `json:"maxPrice,omitempty"`
	TickSize   decimal.Decimal `json:"tickSize,omitempty"`
	MinQty     decimal.Decimal `json:"minQty,omitempty"`
	MaxQty     decimal.Decimal `json:"maxQty,omitempty"`
	StepSize   decimal.Decimal `json:"stepSize,omitempty"`
}

type ApiSort struct {
	BaseAsset string   `json:"baseAsset"`
	Symbols   []string `json:"symbols"`
}

type ApiFill struct {
	Price           decimal.Decimal `json:"price"`
	Qty             decimal.Decimal `json:"qty"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	TradeId         int             `json:"tradeId"`
	MatchType       string          `json:"matchType"`
	AllocId         int             `json:"allocId"`
}

type ApiOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int64  `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type ApiCommissionRate struct {
	Maker  decimal.Decimal `json:"maker"`
	Taker  decimal.Decimal `json:"taker"`
	Buyer  decimal.Decimal `json:"buyer"`
	Seller decimal.Decimal `json:"seller"`
}

type ApiOrderReport struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int64           `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	OrigClientOrderId       string          `json:"origClientOrderId"`
	TransactTime            int64           `json:"transactTime"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice,omitempty"`
	IcebergQty              decimal.Decimal `json:"icebergQty,omitempty"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
}

type ApiSymbol struct {
	Symbol                          string       `json:"symbol"`
	Status                          string       `json:"status"`
	BaseAsset                       string       `json:"baseAsset"`
	BaseAssetPrecision              int          `json:"baseAssetPrecision"`
	QuoteAsset                      string       `json:"quoteAsset"`
	QuotePrecision                  int          `json:"quotePrecision"`
	QuoteAssetPrecision             int          `json:"quoteAssetPrecision"`
	BaseCommissionPrecision         int          `json:"baseCommissionPrecision"`
	QuoteCommissionPrecision        int          `json:"quoteCommissionPrecision"`
	OrderTypes                      []string     `json:"orderTypes"`
	IcebergAllowed                  bool         `json:"icebergAllowed"`
	OcoAllowed                      bool         `json:"ocoAllowed"`
	OtoAllowed                      bool         `json:"otoAllowed"`
	QuoteOrderQtyMarketAllowed      bool         `json:"quoteOrderQtyMarketAllowed"`
	AllowTrailingStop               bool         `json:"allowTrailingStop"`
	CancelReplaceAllowed            bool         `json:"cancelReplaceAllowed"`
	IsSpotTradingAllowed            bool         `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed          bool         `json:"isMarginTradingAllowed"`
	Filters                         []*ApiFilter `json:"filters"`
	Permissions                     []string     `json:"permissions"`
	PermissionSets                  [][]string   `json:"permissionSets"`
	DefaultSelfTradePreventionMode  string       `json:"defaultSelfTradePreventionMode"`
	AllowedSelfTradePreventionModes []string     `json:"allowedSelfTradePreventionModes"`
}

// WsPing Test connectivity to the WebSocket API.
type WsPing struct {
	c *WsClient
}

type ApiError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type ApiResponse struct {
	Id         string          `json:"id"`
	Status     int             `json:"status"`
	RateLimits []*ApiRateLimit `json:"rateLimits,omitempty"`
	Error      *ApiError       `json:"error,omitempty"`
}

type WsPingResponse struct {
	ApiResponse
}

func (s *WsPing) Do(ctx context.Context) (*WsPingResponse, error) {
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
			var resp *WsPingResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsServerTime Test connectivity to the WebSocket API and get the current server time.
type WsServerTime struct {
	c *WsClient
}

type WsServerTimeResponse struct {
	ApiResponse
	Result struct {
		ServerTime int64 `json:"serverTime"`
	} `json:"result"`
}

func (s *WsServerTime) Do(ctx context.Context) (*WsServerTimeResponse, error) {
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
			var resp *WsServerTimeResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsExchangeInfo Query current exchange trading rules, rate limits, and symbol information.
type WsExchangeInfo struct {
	c *WsClient
}

type ExchangeInfoResult struct {
	Timezone   string          `json:"timezone"`
	ServerTime int64           `json:"serverTime"`
	RateLimits []*ApiRateLimit `json:"rateLimits"`
	Symbols    []*ApiSymbol    `json:"symbols"`
	Sors       []*ApiSort      `json:"sors"`
}

type WsExchangeInfoResponse struct {
	ApiResponse
	Result *ExchangeInfoResult `json:"result"`
}

// Symbol Describe a single symbol
func (s *WsExchangeInfo) Symbol(symbol string) *WsExchangeInfo {
	s.c.setParams("symbol", symbol)
	return s
}

// Symbols Describe multiple symbols
func (s *WsExchangeInfo) Symbols(symbols []string) *WsExchangeInfo {
	s.c.setParams("symbols", symbols)
	return s
}

// Permissions Filter symbols by permissions
func (s *WsExchangeInfo) Permissions(permissions []string) *WsExchangeInfo {
	s.c.setParams("permissions", permissions)
	return s
}

// ShowPermissionSets Controls whether the content of the permissionSets field is populated or not. Defaults to true.
func (s *WsExchangeInfo) ShowPermissionSets(showPermissionSets bool) *WsExchangeInfo {
	s.c.setParams("showPermissionSets", showPermissionSets)
	return s
}

// SymbolStatus  Filters symbols that have this tradingStatus.
// Valid values: TRADING, HALT, BREAK
// Cannot be used in combination with symbol or symbols
func (s *WsExchangeInfo) SymbolStatus(symbolStatus core.SymbolStatusEnum) *WsExchangeInfo {
	s.c.setParams("symbolStatus", "symbolStatus")
	return s
}
func (s *WsExchangeInfo) Do(ctx context.Context) (*WsExchangeInfoResponse, error) {
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
			var resp *WsExchangeInfoResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
