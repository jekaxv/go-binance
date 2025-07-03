package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// AccountInfo Get current account information.
type AccountInfo struct {
	c *Client
	r *core.Request
}

type AccountInfoResponse struct {
	MakerCommission            int           `json:"makerCommission"`
	TakerCommission            int           `json:"takerCommission"`
	BuyerCommission            int           `json:"buyerCommission"`
	SellerCommission           int           `json:"sellerCommission"`
	CommissionRates            *Commission   `json:"commissionRates"`
	CanTrade                   bool          `json:"canTrade"`
	CanWithdraw                bool          `json:"canWithdraw"`
	CanDeposit                 bool          `json:"canDeposit"`
	Brokered                   bool          `json:"brokered"`
	RequireSelfTradePrevention bool          `json:"requireSelfTradePrevention"`
	PreventSor                 bool          `json:"preventSor"`
	UpdateTime                 int           `json:"updateTime"`
	AccountType                string        `json:"accountType"`
	Balances                   []*ApiBalance `json:"balances"`
	Permissions                []string      `json:"permissions"`
	Uid                        int           `json:"uid"`
}

// OmitZeroBalances When set to true, emits only the non-zero balances of an account.
// Default value: false
func (s *AccountInfo) OmitZeroBalances(omitZeroBalances bool) *AccountInfo {
	s.r.Set("omitZeroBalances", omitZeroBalances)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *AccountInfo) RecvWindow(recvWindow int64) *AccountInfo {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *AccountInfo) Do(ctx context.Context) (*AccountInfoResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp *AccountInfoResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AccountTrade Get trades for a specific account and symbol.
type AccountTrade struct {
	c *Client
	r *core.Request
}

type AccountTradeResponse struct {
	Symbol          string          `json:"symbol"`
	Id              int             `json:"id"`
	OrderId         int             `json:"orderId"`
	OrderListId     int             `json:"orderListId"`
	Price           decimal.Decimal `json:"price"`
	Qty             decimal.Decimal `json:"qty"`
	QuoteQty        decimal.Decimal `json:"quoteQty"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	Time            int64           `json:"time"`
	IsBuyer         bool            `json:"isBuyer"`
	IsMaker         bool            `json:"isMaker"`
	IsBestMatch     bool            `json:"isBestMatch"`
}

func (s *AccountTrade) Symbol(symbol string) *AccountTrade {
	s.r.Set("symbol", symbol)
	return s
}

// OrderId This can only be used in combination with symbol.
func (s *AccountTrade) OrderId(orderId int64) *AccountTrade {
	s.r.Set("orderId", orderId)
	return s
}
func (s *AccountTrade) StartTime(startTime int64) *AccountTrade {
	s.r.Set("startTime", startTime)
	return s
}
func (s *AccountTrade) EndTime(endTime int64) *AccountTrade {
	s.r.Set("endTime", endTime)
	return s
}

// FromId TradeId to fetch from. Default gets most recent trades.
func (s *AccountTrade) FromId(fromId int64) *AccountTrade {
	s.r.Set("fromId", fromId)
	return s
}
func (s *AccountTrade) Limit(limit int64) *AccountTrade {
	s.r.Set("limit", limit)
	return s
}
func (s *AccountTrade) RecvWindow(recvWindow int64) *AccountTrade {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AccountTrade) Do(ctx context.Context) ([]*AccountTradeResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp []*AccountTradeResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryUnfilledOrder Displays the user's unfilled order count for all intervals.
type QueryUnfilledOrder struct {
	c *Client
	r *core.Request
}

type QueryUnfilledOrderResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

func (s *QueryUnfilledOrder) RecvWindow(recvWindow int64) *QueryUnfilledOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *QueryUnfilledOrder) Do(ctx context.Context) ([]*QueryUnfilledOrderResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp []*QueryUnfilledOrderResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryPreventedMatches Displays the list of orders that were expired due to STP.
type QueryPreventedMatches struct {
	c *Client
	r *core.Request
}
type QueryPreventedMatchesResponse struct {
	Symbol                  string          `json:"symbol"`
	PreventedMatchId        int             `json:"preventedMatchId"`
	TakerOrderId            int             `json:"takerOrderId"`
	MakerSymbol             string          `json:"makerSymbol"`
	MakerOrderId            int             `json:"makerOrderId"`
	TradeGroupId            int             `json:"tradeGroupId"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	Price                   decimal.Decimal `json:"price"`
	MakerPreventedQuantity  decimal.Decimal `json:"makerPreventedQuantity"`
	TransactTime            int64           `json:"transactTime"`
}

func (s *QueryPreventedMatches) Symbol(symbol string) *QueryPreventedMatches {
	s.r.Set("symbol", symbol)
	return s
}
func (s *QueryPreventedMatches) PreventedMatchId(preventedMatchId int64) *QueryPreventedMatches {
	s.r.Set("preventedMatchId", preventedMatchId)
	return s
}
func (s *QueryPreventedMatches) OrderId(orderId int64) *QueryPreventedMatches {
	s.r.Set("orderId", orderId)
	return s
}
func (s *QueryPreventedMatches) FromPreventedMatchId(fromPreventedMatchId int64) *QueryPreventedMatches {
	s.r.Set("fromPreventedMatchId", fromPreventedMatchId)
	return s
}
func (s *QueryPreventedMatches) Limit(limit int64) *QueryPreventedMatches {
	s.r.Set("limit", limit)
	return s
}
func (s *QueryPreventedMatches) RecvWindow(recvWindow int64) *QueryPreventedMatches {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *QueryPreventedMatches) Do(ctx context.Context) ([]*QueryPreventedMatchesResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp []*QueryPreventedMatchesResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryAllocations Retrieves allocations resulting from SOR order placement.
type QueryAllocations struct {
	c *Client
	r *core.Request
}

type QueryAllocationsResponse struct {
	Symbol          string          `json:"symbol"`
	AllocationId    int             `json:"allocationId"`
	AllocationType  string          `json:"allocationType"`
	OrderId         int             `json:"orderId"`
	OrderListId     int             `json:"orderListId"`
	Price           decimal.Decimal `json:"price"`
	Qty             decimal.Decimal `json:"qty"`
	QuoteQty        decimal.Decimal `json:"quoteQty"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	Time            int64           `json:"time"`
	IsBuyer         bool            `json:"isBuyer"`
	IsMaker         bool            `json:"isMaker"`
	IsAllocator     bool            `json:"isAllocator"`
}

func (s *QueryAllocations) Symbol(symbol string) *QueryAllocations {
	s.r.Set("symbol", symbol)
	return s
}
func (s *QueryAllocations) StartTime(startTime int64) *QueryAllocations {
	s.r.Set("startTime", startTime)
	return s
}
func (s *QueryAllocations) EndTime(endTime int64) *QueryAllocations {
	s.r.Set("endTime", endTime)
	return s
}
func (s *QueryAllocations) FromAllocationId(fromAllocationId int64) *QueryAllocations {
	s.r.Set("fromAllocationId", fromAllocationId)
	return s
}
func (s *QueryAllocations) Limit(limit int64) *QueryAllocations {
	s.r.Set("limit", limit)
	return s
}
func (s *QueryAllocations) OrderId(orderId int64) *QueryAllocations {
	s.r.Set("orderId", orderId)
	return s
}
func (s *QueryAllocations) RecvWindow(recvWindow int64) *QueryAllocations {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *QueryAllocations) Do(ctx context.Context) ([]*QueryAllocationsResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp []*QueryAllocationsResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryCommission Get current account commission rates.
type QueryCommission struct {
	c *Client
	r *core.Request
}

type QueryCommissionResponse struct {
	Symbol             string      `json:"symbol"`
	StandardCommission *Commission `json:"standardCommission"`
	TaxCommission      *Commission `json:"taxCommission"`
	Discount           *Discount   `json:"discount"`
}

func (s *QueryCommission) Symbol(symbol string) *QueryCommission {
	s.r.Set("symbol", symbol)
	return s
}
func (s *QueryCommission) Do(ctx context.Context) (*QueryCommissionResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp *QueryCommissionResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}
