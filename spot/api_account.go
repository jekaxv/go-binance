package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// AccountInformation Query information about your account.
type AccountInformation struct {
	c *WsClient
	r *core.WsRequest
}

type AccountInformationResult struct {
	MakerCommission            int                `json:"makerCommission"`
	TakerCommission            int                `json:"takerCommission"`
	BuyerCommission            int                `json:"buyerCommission"`
	SellerCommission           int                `json:"sellerCommission"`
	CanTrade                   bool               `json:"canTrade"`
	CanWithdraw                bool               `json:"canWithdraw"`
	CanDeposit                 bool               `json:"canDeposit"`
	CommissionRates            *ApiCommissionRate `json:"commissionRates"`
	Brokered                   bool               `json:"brokered"`
	RequireSelfTradePrevention bool               `json:"requireSelfTradePrevention"`
	PreventSor                 bool               `json:"preventSor"`
	UpdateTime                 int64              `json:"updateTime"`
	AccountType                string             `json:"accountType"`
	Balances                   []*ApiBalance      `json:"balances"`
	Permissions                []string           `json:"permissions"`
	Uid                        int                `json:"uid"`
}

type AccountInformationResponse struct {
	ApiResponse
	Result *AccountInformationResult `json:"result"`
}

// OmitZeroBalances When set to true, emits only the non-zero balances of an account.
// Default value: false
func (s *AccountInformation) OmitZeroBalances(omitZeroBalances bool) *AccountInformation {
	s.r.Set("omitZeroBalances", omitZeroBalances)
	return s
}

// RecvWindow The value cannot be greater than 60000
func (s *AccountInformation) RecvWindow(recvWindow int) *AccountInformation {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *AccountInformation) Do(ctx context.Context) (*AccountInformationResponse, error) {
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
			var resp *AccountInformationResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// UnfilledOrder Query your current unfilled order count for all intervals.
type UnfilledOrder struct {
	c *WsClient
	r *core.WsRequest
}

type UnfilledOrderResult struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
	Count         int    `json:"count"`
}
type UnfilledOrderResponse struct {
	ApiResponse
	Result []*UnfilledOrderResult `json:"result"`
}

func (s *UnfilledOrder) RecvWindow(recvWindow int) *UnfilledOrder {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *UnfilledOrder) Do(ctx context.Context) (*UnfilledOrderResponse, error) {
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
			var resp *UnfilledOrderResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// AccountOrderHistory Query information about all your orders – active, canceled, filled – filtered by time range.
type AccountOrderHistory struct {
	c *WsClient
	r *core.WsRequest
}

type AccountOrderHistoryResult struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int64           `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice"`
	IcebergQty              decimal.Decimal `json:"icebergQty"`
	Time                    int64           `json:"time"`
	UpdateTime              int64           `json:"updateTime"`
	IsWorking               bool            `json:"isWorking"`
	WorkingTime             int64           `json:"workingTime"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
	PreventedMatchId        int             `json:"preventedMatchId"`
	PreventedQuantity       string          `json:"preventedQuantity"`
}

type AccountOrderHistoryResponse struct {
	ApiResponse
	Result []*AccountOrderHistoryResult `json:"result"`
}

func (s *AccountOrderHistory) Symbol(symbol string) *AccountOrderHistory {
	s.r.Set("symbol", symbol)
	return s
}

// OrderId Order ID to begin at
func (s *AccountOrderHistory) OrderId(orderId int) *AccountOrderHistory {
	s.r.Set("orderId", orderId)
	return s
}
func (s *AccountOrderHistory) StartTime(startTime int) *AccountOrderHistory {
	s.r.Set("startTime", startTime)
	return s
}
func (s *AccountOrderHistory) EndTime(endTime int) *AccountOrderHistory {
	s.r.Set("endTime", endTime)
	return s
}
func (s *AccountOrderHistory) Limit(limit int) *AccountOrderHistory {
	s.r.Set("limit", limit)
	return s
}
func (s *AccountOrderHistory) RecvWindow(recvWindow int) *AccountOrderHistory {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AccountOrderHistory) Do(ctx context.Context) (*AccountOrderHistoryResponse, error) {
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
			var resp *AccountOrderHistoryResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// AllOrderList Query information about all your order lists, filtered by time range.
type AllOrderList struct {
	c *WsClient
	r *core.WsRequest
}

type AllOrderListResult struct {
	OrderListId       int         `json:"orderListId"`
	ContingencyType   string      `json:"contingencyType"`
	ListStatusType    string      `json:"listStatusType"`
	ListOrderStatus   string      `json:"listOrderStatus"`
	ListClientOrderId string      `json:"listClientOrderId"`
	TransactionTime   int64       `json:"transactionTime"`
	Symbol            string      `json:"symbol"`
	Orders            []*ApiOrder `json:"orders"`
}

type AllOrderListResponse struct {
	ApiResponse
	Result []*AllOrderListResult `json:"result"`
}

// FromId Order list ID to begin at
func (s *AllOrderList) FromId(fromId int) *AllOrderList {
	s.r.Set("fromId", fromId)
	return s
}
func (s *AllOrderList) StartTime(startTime int) *AllOrderList {
	s.r.Set("startTime", startTime)
	return s
}
func (s *AllOrderList) EndTime(endTime int) *AllOrderList {
	s.r.Set("endTime", endTime)
	return s
}
func (s *AllOrderList) Limit(limit int) *AllOrderList {
	s.r.Set("limit", limit)
	return s
}
func (s *AllOrderList) RecvWindow(recvWindow int) *AllOrderList {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AllOrderList) Do(ctx context.Context) (*AllOrderListResponse, error) {
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
			var resp *AllOrderListResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// AccountTradeHistory Query information about all your trades, filtered by time range.
type AccountTradeHistory struct {
	c *WsClient
	r *core.WsRequest
}
type AccountTradeHistoryResult struct {
	Symbol          string          `json:"symbol"`
	Id              int             `json:"id"`
	OrderId         int64           `json:"orderId"`
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
type AccountTradeHistoryResponse struct {
	ApiResponse
	Result []*AccountTradeHistoryResult `json:"result"`
}

func (s *AccountTradeHistory) Symbol(symbol string) *AccountTradeHistory {
	s.r.Set("symbol", symbol)
	return s
}
func (s *AccountTradeHistory) OrderId(orderId int) *AccountTradeHistory {
	s.r.Set("orderId", orderId)
	return s
}
func (s *AccountTradeHistory) StartTime(startTime int) *AccountTradeHistory {
	s.r.Set("startTime", startTime)
	return s
}
func (s *AccountTradeHistory) EndTime(endTime int) *AccountTradeHistory {
	s.r.Set("endTime", endTime)
	return s
}
func (s *AccountTradeHistory) FromId(fromId int) *AccountTradeHistory {
	s.r.Set("fromId", fromId)
	return s
}
func (s *AccountTradeHistory) Limit(limit int) *AccountTradeHistory {
	s.r.Set("limit", limit)
	return s
}
func (s *AccountTradeHistory) RecvWindow(recvWindow int) *AccountTradeHistory {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AccountTradeHistory) Do(ctx context.Context) (*AccountTradeHistoryResponse, error) {
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
			var resp *AccountTradeHistoryResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// AccountPreventedMatches Displays the list of orders that were expired due to STP.
type AccountPreventedMatches struct {
	c *WsClient
	r *core.WsRequest
}
type AccountPreventedMatchesResult struct {
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
type AccountPreventedMatchesResponse struct {
	ApiResponse
	Result []*AccountPreventedMatchesResult `json:"result"`
}

func (s *AccountPreventedMatches) Symbol(symbol string) *AccountPreventedMatches {
	s.r.Set("symbol", symbol)
	return s
}
func (s *AccountPreventedMatches) PreventedMatchId(preventedMatchId int) *AccountPreventedMatches {
	s.r.Set("preventedMatchId", preventedMatchId)
	return s
}
func (s *AccountPreventedMatches) OrderId(orderId int) *AccountPreventedMatches {
	s.r.Set("orderId", orderId)
	return s
}
func (s *AccountPreventedMatches) FromPreventedMatchId(fromPreventedMatchId int) *AccountPreventedMatches {
	s.r.Set("fromPreventedMatchId", fromPreventedMatchId)
	return s
}
func (s *AccountPreventedMatches) Limit(limit int) *AccountPreventedMatches {
	s.r.Set("limit", limit)
	return s
}
func (s *AccountPreventedMatches) RecvWindow(recvWindow int) *AccountPreventedMatches {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *AccountPreventedMatches) Do(ctx context.Context) (*AccountPreventedMatchesResponse, error) {
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
			var resp *AccountPreventedMatchesResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// AccountAllocations Retrieves allocations resulting from SOR order placement.
type AccountAllocations struct {
	c *WsClient
	r *core.WsRequest
}
type AccountAllocationsResult struct {
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
type AccountAllocationsResponse struct {
	ApiResponse
	Result []*AccountAllocationsResult `json:"result"`
}

func (s *AccountAllocations) Symbol(symbol string) *AccountAllocations {
	s.r.Set("symbol", symbol)
	return s
}
func (s *AccountAllocations) StartTime(startTime int) *AccountAllocations {
	s.r.Set("startTime", startTime)
	return s
}
func (s *AccountAllocations) EndTime(endTime int) *AccountAllocations {
	s.r.Set("endTime", endTime)
	return s
}
func (s *AccountAllocations) FromAllocationId(fromAllocationId int) *AccountAllocations {
	s.r.Set("fromAllocationId", fromAllocationId)
	return s
}
func (s *AccountAllocations) Limit(limit int) *AccountAllocations {
	s.r.Set("limit", limit)
	return s
}
func (s *AccountAllocations) OrderId(orderId int) *AccountAllocations {
	s.r.Set("orderId", orderId)
	return s
}
func (s *AccountAllocations) RecvWindow(recvWindow int) *AccountAllocations {
	s.r.Set("recvWindow", recvWindow)
	return s
}
func (s *AccountAllocations) Do(ctx context.Context) (*AccountAllocationsResponse, error) {
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
			var resp *AccountAllocationsResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

type AccountCommission struct {
	c *WsClient
	r *core.WsRequest
}
type AccountCommissionResult struct {
	Symbol             string             `json:"symbol"`
	StandardCommission *ApiCommissionRate `json:"standardCommission"`
	TaxCommission      *ApiCommissionRate `json:"taxCommission"`
	Discount           *Discount          `json:"discount"`
}

type AccountCommissionResponse struct {
	ApiResponse
	Result *AccountCommissionResult `json:"result"`
}

func (s *AccountCommission) Symbol(symbol string) *AccountCommission {
	s.r.Set("symbol", symbol)
	return s
}

func (s *AccountCommission) Do(ctx context.Context) (*AccountCommissionResponse, error) {
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
			var resp *AccountCommissionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
