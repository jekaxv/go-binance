package http

import (
	"context"
	"encoding/json"
)

// AccountInfo Get current account information.
type AccountInfo struct {
	c                *Client
	omitZeroBalances *bool
	recvWindow       *int64
}

type AccountInfoResponse struct {
	MakerCommission            int             `json:"makerCommission"`
	TakerCommission            int             `json:"takerCommission"`
	BuyerCommission            int             `json:"buyerCommission"`
	SellerCommission           int             `json:"sellerCommission"`
	CommissionRates            *SpotCommission `json:"commissionRates"`
	CanTrade                   bool            `json:"canTrade"`
	CanWithdraw                bool            `json:"canWithdraw"`
	CanDeposit                 bool            `json:"canDeposit"`
	Brokered                   bool            `json:"brokered"`
	RequireSelfTradePrevention bool            `json:"requireSelfTradePrevention"`
	PreventSor                 bool            `json:"preventSor"`
	UpdateTime                 int             `json:"updateTime"`
	AccountType                string          `json:"accountType"`
	Balances                   []*ApiBalance   `json:"balances"`
	Permissions                []string        `json:"permissions"`
	Uid                        int             `json:"uid"`
}

// OmitZeroBalances When set to true, emits only the non-zero balances of an account.
// Default value: false
func (s *AccountInfo) OmitZeroBalances(omitZeroBalances bool) *AccountInfo {
	s.omitZeroBalances = &omitZeroBalances
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *AccountInfo) RecvWindow(recvWindow int64) *AccountInfo {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountInfo) Do(ctx context.Context) (*AccountInfoResponse, error) {
	if s.omitZeroBalances != nil {
		s.c.req.set("omitZeroBalances", *s.omitZeroBalances)
	}
	if s.recvWindow != nil {
		s.c.req.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *AccountInfoResponse
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// AccountTrade Get trades for a specific account and symbol.
type AccountTrade struct {
	c          *Client
	symbol     string
	orderId    *int64
	startTime  *int64
	endTime    *int64
	fromId     *int64
	limit      *int64
	recvWindow *int64
}

type AccountTradeResponse struct {
	Symbol          string `json:"symbol"`
	Id              int    `json:"id"`
	OrderId         int    `json:"orderId"`
	OrderListId     int    `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsBestMatch     bool   `json:"isBestMatch"`
}

func (s *AccountTrade) Symbol(symbol string) *AccountTrade {
	s.symbol = symbol
	return s
}

// OrderId This can only be used in combination with symbol.
func (s *AccountTrade) OrderId(orderId int64) *AccountTrade {
	s.orderId = &orderId
	return s
}
func (s *AccountTrade) StartTime(startTime int64) *AccountTrade {
	s.startTime = &startTime
	return s
}
func (s *AccountTrade) EndTime(endTime int64) *AccountTrade {
	s.endTime = &endTime
	return s
}

// FromId TradeId to fetch from. Default gets most recent trades.
func (s *AccountTrade) FromId(fromId int64) *AccountTrade {
	s.fromId = &fromId
	return s
}
func (s *AccountTrade) Limit(limit int64) *AccountTrade {
	s.limit = &limit
	return s
}
func (s *AccountTrade) RecvWindow(recvWindow int64) *AccountTrade {
	s.recvWindow = &recvWindow
	return s
}
func (s *AccountTrade) Do(ctx context.Context) ([]*AccountTradeResponse, error) {
	s.c.req.set("symbol", s.symbol)
	if s.orderId != nil {
		s.c.req.set("orderId", *s.orderId)
	}
	if s.startTime != nil {
		s.c.req.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.req.set("endTime", *s.endTime)
	}
	if s.fromId != nil {
		s.c.req.set("fromId", *s.fromId)
	}
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.recvWindow != nil {
		s.c.req.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*AccountTradeResponse
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// QueryUnfilledOrder Displays the user's unfilled order count for all intervals.
type QueryUnfilledOrder struct {
	c          *Client
	recvWindow *int64
}

type QueryUnfilledOrderResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}

func (s *QueryUnfilledOrder) RecvWindow(recvWindow int64) *QueryUnfilledOrder {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryUnfilledOrder) Do(ctx context.Context) ([]*QueryUnfilledOrderResponse, error) {
	if s.recvWindow != nil {
		s.c.req.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*QueryUnfilledOrderResponse
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// QueryPreventedMatches Displays the list of orders that were expired due to STP.
type QueryPreventedMatches struct {
	c                    *Client
	symbol               string
	preventedMatchId     *int64
	orderId              *int64
	fromPreventedMatchId *int64
	limit                *int64
	recvWindow           *int64
}
type QueryPreventedMatchesResponse struct {
	Symbol                  string `json:"symbol"`
	PreventedMatchId        int    `json:"preventedMatchId"`
	TakerOrderId            int    `json:"takerOrderId"`
	MakerSymbol             string `json:"makerSymbol"`
	MakerOrderId            int    `json:"makerOrderId"`
	TradeGroupId            int    `json:"tradeGroupId"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
	Price                   string `json:"price"`
	MakerPreventedQuantity  string `json:"makerPreventedQuantity"`
	TransactTime            int64  `json:"transactTime"`
}

func (s *QueryPreventedMatches) Symbol(symbol string) *QueryPreventedMatches {
	s.symbol = symbol
	return s
}
func (s *QueryPreventedMatches) PreventedMatchId(preventedMatchId int64) *QueryPreventedMatches {
	s.preventedMatchId = &preventedMatchId
	return s
}
func (s *QueryPreventedMatches) OrderId(orderId int64) *QueryPreventedMatches {
	s.orderId = &orderId
	return s
}
func (s *QueryPreventedMatches) FromPreventedMatchId(fromPreventedMatchId int64) *QueryPreventedMatches {
	s.fromPreventedMatchId = &fromPreventedMatchId
	return s
}
func (s *QueryPreventedMatches) Limit(limit int64) *QueryPreventedMatches {
	s.limit = &limit
	return s
}
func (s *QueryPreventedMatches) RecvWindow(recvWindow int64) *QueryPreventedMatches {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryPreventedMatches) Do(ctx context.Context) ([]*QueryPreventedMatchesResponse, error) {
	s.c.req.set("symbol", s.symbol)
	if s.preventedMatchId != nil {
		s.c.req.set("preventedMatchId", *s.preventedMatchId)
	}
	if s.orderId != nil {
		s.c.req.set("orderId", *s.orderId)
	}
	if s.fromPreventedMatchId != nil {
		s.c.req.set("fromPreventedMatchId", *s.fromPreventedMatchId)
	}
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.recvWindow != nil {
		s.c.req.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*QueryPreventedMatchesResponse
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// QueryAllocations Retrieves allocations resulting from SOR order placement.
type QueryAllocations struct {
	c                *Client
	symbol           string
	startTime        *int64
	endTime          *int64
	fromAllocationId *int64
	limit            *int64
	orderId          *int64
	recvWindow       *int64
}

type QueryAllocationsResponse struct {
	Symbol          string `json:"symbol"`
	AllocationId    int    `json:"allocationId"`
	AllocationType  string `json:"allocationType"`
	OrderId         int    `json:"orderId"`
	OrderListId     int    `json:"orderListId"`
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	QuoteQty        string `json:"quoteQty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	Time            int64  `json:"time"`
	IsBuyer         bool   `json:"isBuyer"`
	IsMaker         bool   `json:"isMaker"`
	IsAllocator     bool   `json:"isAllocator"`
}

func (s *QueryAllocations) Symbol(symbol string) *QueryAllocations {
	s.symbol = symbol
	return s
}
func (s *QueryAllocations) StartTime(startTime int64) *QueryAllocations {
	s.startTime = &startTime
	return s
}
func (s *QueryAllocations) EndTime(endTime int64) *QueryAllocations {
	s.endTime = &endTime
	return s
}
func (s *QueryAllocations) FromAllocationId(fromAllocationId int64) *QueryAllocations {
	s.fromAllocationId = &fromAllocationId
	return s
}
func (s *QueryAllocations) Limit(limit int64) *QueryAllocations {
	s.limit = &limit
	return s
}
func (s *QueryAllocations) OrderId(orderId int64) *QueryAllocations {
	s.orderId = &orderId
	return s
}
func (s *QueryAllocations) RecvWindow(recvWindow int64) *QueryAllocations {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryAllocations) Do(ctx context.Context) ([]*QueryAllocationsResponse, error) {
	s.c.req.set("symbol", s.symbol)
	if s.startTime != nil {
		s.c.req.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.req.set("endTime", *s.endTime)
	}
	if s.fromAllocationId != nil {
		s.c.req.set("fromAllocationId", *s.fromAllocationId)
	}
	if s.limit != nil {
		s.c.req.set("limit", *s.limit)
	}
	if s.orderId != nil {
		s.c.req.set("orderId", *s.orderId)
	}
	if s.recvWindow != nil {
		s.c.req.set("recvWindow", *s.recvWindow)
	}
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp []*QueryAllocationsResponse
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}

// QueryCommission Get current account commission rates.
type QueryCommission struct {
	c      *Client
	symbol string
}

type QueryCommissionResponse struct {
	Symbol             string          `json:"symbol"`
	StandardCommission *SpotCommission `json:"standardCommission"`
	TaxCommission      *SpotCommission `json:"taxCommission"`
	Discount           struct {
		EnabledForAccount bool   `json:"enabledForAccount"`
		EnabledForSymbol  bool   `json:"enabledForSymbol"`
		DiscountAsset     string `json:"discountAsset"`
		Discount          string `json:"discount"`
	} `json:"discount"`
}

func (s *QueryCommission) Symbol(symbol string) *QueryCommission {
	s.symbol = symbol
	return s
}
func (s *QueryCommission) Do(ctx context.Context) (*QueryCommissionResponse, error) {
	s.c.req.set("symbol", s.symbol)
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *QueryCommissionResponse
	return resp, json.Unmarshal(s.c.resp.rawBody, &resp)
}
