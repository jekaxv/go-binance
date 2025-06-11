package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// QueryBalance Query account balance info
type QueryBalance struct {
	c          *Client
	recvWindow *int64
}

type QueryBalanceResponse struct {
	AccountAlias       string          `json:"accountAlias"`
	Asset              string          `json:"asset"`
	Balance            decimal.Decimal `json:"balance"`
	CrossWalletBalance decimal.Decimal `json:"crossWalletBalance"`
	CrossUnPnl         decimal.Decimal `json:"crossUnPnl"`
	AvailableBalance   decimal.Decimal `json:"availableBalance"`
	MaxWithdrawAmount  decimal.Decimal `json:"maxWithdrawAmount"`
	MarginAvailable    bool            `json:"marginAvailable"`
	UpdateTime         int64           `json:"updateTime"`
}

func (s *QueryBalance) RecvWindow(recvWindow int64) *QueryBalance {
	s.recvWindow = &recvWindow
	return s
}

func (s *QueryBalance) Do(ctx context.Context) ([]*QueryBalanceResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp []*QueryBalanceResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AccountInfo Get current account information. User in single-asset/ multi-assets mode will see different value, see comments in response section for detail.
type AccountInfo struct {
	c          *Client
	recvWindow *int64
}

type AccountAsset struct {
	Asset                  string          `json:"asset"`
	WalletBalance          decimal.Decimal `json:"walletBalance"`
	UnrealizedProfit       decimal.Decimal `json:"unrealizedProfit"`
	MarginBalance          decimal.Decimal `json:"marginBalance"`
	MaintMargin            decimal.Decimal `json:"maintMargin"`
	InitialMargin          decimal.Decimal `json:"initialMargin"`
	PositionInitialMargin  decimal.Decimal `json:"positionInitialMargin"`
	OpenOrderInitialMargin decimal.Decimal `json:"openOrderInitialMargin"`
	CrossWalletBalance     decimal.Decimal `json:"crossWalletBalance"`
	CrossUnPnl             decimal.Decimal `json:"crossUnPnl"`
	AvailableBalance       decimal.Decimal `json:"availableBalance"`
	MaxWithdrawAmount      decimal.Decimal `json:"maxWithdrawAmount"`
	MarginAvailable        bool            `json:"marginAvailable"`
	UpdateTime             int64           `json:"updateTime"`
}

type AccountPosition struct {
	Symbol           string          `json:"symbol"`
	PositionSide     string          `json:"positionSide"`
	PositionAmt      decimal.Decimal `json:"positionAmt"`
	UnrealizedProfit decimal.Decimal `json:"unrealizedProfit"`
	IsolatedMargin   decimal.Decimal `json:"isolatedMargin"`
	Notional         decimal.Decimal `json:"notional"`
	IsolatedWallet   decimal.Decimal `json:"isolatedWallet"`
	InitialMargin    decimal.Decimal `json:"initialMargin"`
	MaintMargin      decimal.Decimal `json:"maintMargin"`
	UpdateTime       int             `json:"updateTime"`
}

type AccountInfoResponse struct {
	TotalInitialMargin          decimal.Decimal    `json:"totalInitialMargin"`
	TotalMaintMargin            decimal.Decimal    `json:"totalMaintMargin"`
	TotalWalletBalance          decimal.Decimal    `json:"totalWalletBalance"`
	TotalUnrealizedProfit       decimal.Decimal    `json:"totalUnrealizedProfit"`
	TotalMarginBalance          decimal.Decimal    `json:"totalMarginBalance"`
	TotalPositionInitialMargin  decimal.Decimal    `json:"totalPositionInitialMargin"`
	TotalOpenOrderInitialMargin decimal.Decimal    `json:"totalOpenOrderInitialMargin"`
	TotalCrossWalletBalance     decimal.Decimal    `json:"totalCrossWalletBalance"`
	TotalCrossUnPnl             decimal.Decimal    `json:"totalCrossUnPnl"`
	AvailableBalance            decimal.Decimal    `json:"availableBalance"`
	MaxWithdrawAmount           decimal.Decimal    `json:"maxWithdrawAmount"`
	Assets                      []*AccountAsset    `json:"assets"`
	Positions                   []*AccountPosition `json:"positions"`
}

func (s *AccountInfo) RecvWindow(recvWindow int64) *AccountInfo {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountInfo) Do(ctx context.Context) (*AccountInfoResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *AccountInfoResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CommissionRate Get User Commission Rate
type CommissionRate struct {
	c          *Client
	symbol     string
	recvWindow *int64
}

type CommissionRateResponse struct {
	Symbol              string          `json:"symbol"`
	MakerCommissionRate decimal.Decimal `json:"makerCommissionRate"`
	TakerCommissionRate decimal.Decimal `json:"takerCommissionRate"`
}

func (s *CommissionRate) Symbol(symbol string) *CommissionRate {
	s.symbol = symbol
	return s
}

func (s *CommissionRate) RecvWindow(recvWindow int64) *CommissionRate {
	s.recvWindow = &recvWindow
	return s
}

func (s *CommissionRate) Do(ctx context.Context) (*CommissionRateResponse, error) {
	s.c.set("symbol", s.symbol)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *CommissionRateResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AccountConfig Query account configuration
type AccountConfig struct {
	c          *Client
	recvWindow *int64
}

type AccountConfigResponse struct {
	FeeTier           int  `json:"feeTier"`
	CanTrade          bool `json:"canTrade"`
	CanDeposit        bool `json:"canDeposit"`
	CanWithdraw       bool `json:"canWithdraw"`
	DualSidePosition  bool `json:"dualSidePosition"`
	UpdateTime        int  `json:"updateTime"`
	MultiAssetsMargin bool `json:"multiAssetsMargin"`
	TradeGroupId      int  `json:"tradeGroupId"`
}

func (s *AccountConfig) RecvWindow(recvWindow int64) *AccountConfig {
	s.recvWindow = &recvWindow
	return s
}

func (s *AccountConfig) Do(ctx context.Context) (*AccountConfigResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *AccountConfigResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// SymbolConfig Get current account symbol configuration.
type SymbolConfig struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}
type SymbolConfigResponse struct {
	Symbol           string          `json:"symbol"`
	MarginType       string          `json:"marginType"`
	IsAutoAddMargin  string          `json:"isAutoAddMargin"`
	Leverage         int             `json:"leverage"`
	MaxNotionalValue decimal.Decimal `json:"maxNotionalValue"`
}

func (s *SymbolConfig) Symbol(symbol string) *SymbolConfig {
	s.symbol = &symbol
	return s
}
func (s *SymbolConfig) RecvWindow(recvWindow int64) *SymbolConfig {
	s.recvWindow = &recvWindow
	return s
}
func (s *SymbolConfig) Do(ctx context.Context) ([]*SymbolConfigResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp []*SymbolConfigResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryRateLimit Query User Rate Limit
type QueryRateLimit struct {
	c          *Client
	recvWindow *int64
}

type QueryRateLimitResponse struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	IntervalNum   int    `json:"intervalNum"`
	Limit         int    `json:"limit"`
}

func (s *QueryRateLimit) RecvWindow(recvWindow int64) *QueryRateLimit {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryRateLimit) Do(ctx context.Context) ([]*QueryRateLimitResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp []*QueryRateLimitResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// LeverageBracket Query user notional and leverage bracket on speicfic symbol
type LeverageBracket struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}

type Bracket struct {
	Bracket          int     `json:"bracket"`
	InitialLeverage  int     `json:"initialLeverage"`
	NotionalCap      int     `json:"notionalCap"`
	NotionalFloor    int     `json:"notionalFloor"`
	MaintMarginRatio float64 `json:"maintMarginRatio"`
	Cum              int     `json:"cum"`
}
type LeverageBracketResponse struct {
	Symbol       string     `json:"symbol"`
	NotionalCoef float64    `json:"notionalCoef"`
	Brackets     []*Bracket `json:"brackets"`
}

func (s *LeverageBracket) Symbol(symbol string) *LeverageBracket {
	s.symbol = &symbol
	return s
}
func (s *LeverageBracket) RecvWindow(recvWindow int64) *LeverageBracket {
	s.recvWindow = &recvWindow
	return s
}
func (s *LeverageBracket) Do(ctx context.Context) ([]*LeverageBracketResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp []*LeverageBracketResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	if s.symbol == nil {
		return resp, json.Unmarshal(s.c.rawBody(), &resp)
	}
	var res *LeverageBracketResponse
	if err := json.Unmarshal(s.c.rawBody(), &res); err != nil {
		return nil, err
	}
	resp = append(resp, res)
	return resp, nil
}

// MultiAssetsMargin Get user's Multi-Assets mode (Multi-Assets Mode or Single-Asset Mode) on Every symbol
type MultiAssetsMargin struct {
	c          *Client
	recvWindow *int64
}
type MultiAssetsMarginResponse struct {
	MultiAssetsMargin bool `json:"multiAssetsMargin"` // Multi-Assets Mode; "false": Single-Asset Mode
}

func (s *MultiAssetsMargin) RecvWindow(recvWindow int64) *MultiAssetsMargin {
	s.recvWindow = &recvWindow
	return s
}

func (s *MultiAssetsMargin) Do(ctx context.Context) (*MultiAssetsMarginResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *MultiAssetsMarginResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// GetPositionSide Get user's position mode (Hedge Mode or One-way Mode ) on EVERY symbol
type GetPositionSide struct {
	c          *Client
	recvWindow *int64
}

type GetPositionSideResponse struct {
	DualSidePosition bool `json:"dualSidePosition"` // "true": Hedge Mode; "false": One-way Mode
}

func (s *GetPositionSide) RecvWindow(recvWindow int64) *GetPositionSide {
	s.recvWindow = &recvWindow
	return s
}

func (s *GetPositionSide) Do(ctx context.Context) (*GetPositionSideResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *GetPositionSideResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryIncome Query income history
type QueryIncome struct {
	c          *Client
	symbol     *string
	incomeType *core.IncomeType
	startTime  *int64
	endTime    *int64
	page       *int
	limit      *int
	recvWindow *int64
}
type QueryIncomeResponse struct {
	Symbol     string          `json:"symbol"`
	IncomeType string          `json:"incomeType"`
	Income     decimal.Decimal `json:"income"`
	Asset      string          `json:"asset"`
	Info       string          `json:"info"`
	Time       int64           `json:"time"`
	TranId     int64           `json:"tranId"`
	TradeId    string          `json:"tradeId"`
}

func (s *QueryIncome) Symbol(symbol string) *QueryIncome {
	s.symbol = &symbol
	return s
}
func (s *QueryIncome) IncomeType(incomeType core.IncomeType) *QueryIncome {
	s.incomeType = &incomeType
	return s
}
func (s *QueryIncome) StartTime(startTime int64) *QueryIncome {
	s.startTime = &startTime
	return s
}
func (s *QueryIncome) EndTime(endTime int64) *QueryIncome {
	s.endTime = &endTime
	return s
}
func (s *QueryIncome) Page(page int) *QueryIncome {
	s.page = &page
	return s
}

// Limit Default 100; max 1000
func (s *QueryIncome) Limit(limit int) *QueryIncome {
	s.limit = &limit
	return s
}
func (s *QueryIncome) RecvWindow(recvWindow int64) *QueryIncome {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryIncome) Do(ctx context.Context) ([]*QueryIncomeResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.incomeType != nil {
		s.c.set("incomeType", *s.incomeType)
	}
	if s.startTime != nil {
		s.c.set("startTime", *s.startTime)
	}
	if s.endTime != nil {
		s.c.set("endTime", *s.endTime)
	}
	if s.page != nil {
		s.c.set("page", *s.page)
	}
	if s.limit != nil {
		s.c.set("limit", *s.limit)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp []*QueryIncomeResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TradingStatus Futures trading quantitative rules indicators
type TradingStatus struct {
	c          *Client
	symbol     *string
	recvWindow *int64
}
type IndicatorResult struct {
	IsLocked           bool    `json:"isLocked"`
	PlannedRecoverTime int64   `json:"plannedRecoverTime"`
	Indicator          string  `json:"indicator"`
	Value              float64 `json:"value"`
	TriggerValue       float64 `json:"triggerValue"`
}
type TradingStatusResponse struct {
	Indicators map[string][]*IndicatorResult `json:"indicators"`
	UpdateTime int64                         `json:"updateTime"`
}

func (s *TradingStatus) Symbol(symbol string) *TradingStatus {
	s.symbol = &symbol
	return s
}
func (s *TradingStatus) RecvWindow(recvWindow int64) *TradingStatus {
	s.recvWindow = &recvWindow
	return s
}
func (s *TradingStatus) Do(ctx context.Context) (*TradingStatusResponse, error) {
	if s.symbol != nil {
		s.c.set("symbol", *s.symbol)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *TradingStatusResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TransactionHistory Get download id for futures transaction history
type TransactionHistory struct {
	c          *Client
	startTime  int64
	endTime    int64
	recvWindow *int64
}
type HistoryResponse struct {
	AvgCostTimestampOfLast30D int    `json:"avgCostTimestampOfLast30d"`
	DownloadId                string `json:"downloadId"`
}

func (s *TransactionHistory) StartTime(startTime int64) *TransactionHistory {
	s.startTime = startTime
	return s
}
func (s *TransactionHistory) EndTime(endTime int64) *TransactionHistory {
	s.endTime = endTime
	return s
}
func (s *TransactionHistory) RecvWindow(recvWindow int64) *TransactionHistory {
	s.recvWindow = &recvWindow
	return s
}
func (s *TransactionHistory) Do(ctx context.Context) (*HistoryResponse, error) {
	s.c.set("startTime", s.startTime)
	s.c.set("endTime", s.endTime)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *HistoryResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TransactionHistoryLink Get futures transaction history download link by Id
type TransactionHistoryLink struct {
	c          *Client
	downloadId string
	recvWindow *int64
}
type HistoryLinkResponse struct {
	DownloadId          string      `json:"downloadId"`
	Status              string      `json:"status"` // Enum：completed，processing
	Url                 string      `json:"url"`
	Notified            bool        `json:"notified"`
	ExpirationTimestamp int         `json:"expirationTimestamp"`
	IsExpired           interface{} `json:"isExpired"`
}

func (s *TransactionHistoryLink) DownloadId(downloadId string) *TransactionHistoryLink {
	s.downloadId = downloadId
	return s
}
func (s *TransactionHistoryLink) RecvWindow(recvWindow int64) *TransactionHistoryLink {
	s.recvWindow = &recvWindow
	return s
}
func (s *TransactionHistoryLink) Do(ctx context.Context) (*HistoryLinkResponse, error) {
	s.c.set("downloadId", s.downloadId)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *HistoryLinkResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OrderHistory Get Download Id For Futures Order History
type OrderHistory struct {
	c          *Client
	startTime  int64
	endTime    int64
	recvWindow *int64
}

func (s *OrderHistory) StartTime(startTime int64) *OrderHistory {
	s.startTime = startTime
	return s
}
func (s *OrderHistory) EndTime(endTime int64) *OrderHistory {
	s.endTime = endTime
	return s
}
func (s *OrderHistory) RecvWindow(recvWindow int64) *OrderHistory {
	s.recvWindow = &recvWindow
	return s
}
func (s *OrderHistory) Do(ctx context.Context) (*HistoryResponse, error) {
	s.c.set("startTime", s.startTime)
	s.c.set("endTime", s.endTime)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *HistoryResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// OrderHistoryLink Get futures order history download link by Id
type OrderHistoryLink struct {
	c          *Client
	downloadId string
	recvWindow *int64
}

func (s *OrderHistoryLink) DownloadId(downloadId string) *OrderHistoryLink {
	s.downloadId = downloadId
	return s
}
func (s *OrderHistoryLink) RecvWindow(recvWindow int64) *OrderHistoryLink {
	s.recvWindow = &recvWindow
	return s
}
func (s *OrderHistoryLink) Do(ctx context.Context) (*HistoryLinkResponse, error) {
	s.c.set("downloadId", s.downloadId)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *HistoryLinkResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TradeHistory Get download id for futures trade history
type TradeHistory struct {
	c          *Client
	startTime  int64
	endTime    int64
	recvWindow *int64
}

func (s *TradeHistory) StartTime(startTime int64) *TradeHistory {
	s.startTime = startTime
	return s
}
func (s *TradeHistory) EndTime(endTime int64) *TradeHistory {
	s.endTime = endTime
	return s
}
func (s *TradeHistory) RecvWindow(recvWindow int64) *TradeHistory {
	s.recvWindow = &recvWindow
	return s
}
func (s *TradeHistory) Do(ctx context.Context) (*HistoryResponse, error) {
	s.c.set("startTime", s.startTime)
	s.c.set("endTime", s.endTime)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *HistoryResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// TradeHistoryLink Get futures trade download link by Id
type TradeHistoryLink struct {
	c          *Client
	downloadId string
	recvWindow *int64
}

func (s *TradeHistoryLink) DownloadId(downloadId string) *TradeHistoryLink {
	s.downloadId = downloadId
	return s
}
func (s *TradeHistoryLink) RecvWindow(recvWindow int64) *TradeHistoryLink {
	s.recvWindow = &recvWindow
	return s
}
func (s *TradeHistoryLink) Do(ctx context.Context) (*HistoryLinkResponse, error) {
	s.c.set("downloadId", s.downloadId)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *HistoryLinkResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ChangeFeeBurn Change user's BNB Fee Discount (Fee Discount On or Fee Discount Off ) on EVERY symbol
type ChangeFeeBurn struct {
	c          *Client
	feeBurn    string
	recvWindow *int64
}
type ChangeFeeBurnResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// FeeBurn 	"true": Fee Discount On; "false": Fee Discount Off
func (s *ChangeFeeBurn) FeeBurn(feeBurn string) *ChangeFeeBurn {
	s.feeBurn = feeBurn
	return s
}
func (s *ChangeFeeBurn) RecvWindow(recvWindow int64) *ChangeFeeBurn {
	s.recvWindow = &recvWindow
	return s
}
func (s *ChangeFeeBurn) Do(ctx context.Context) (*ChangeFeeBurnResponse, error) {
	s.c.set("feeBurn", s.feeBurn)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *ChangeFeeBurnResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// QueryFeeBurn Get user's BNB Fee Discount (Fee Discount On or Fee Discount Off )
type QueryFeeBurn struct {
	c          *Client
	recvWindow *int64
}
type QueryFeeBurnResponse struct {
	FeeBurn bool `json:"feeBurn"`
}

func (s *QueryFeeBurn) RecvWindow(recvWindow int64) *QueryFeeBurn {
	s.recvWindow = &recvWindow
	return s
}
func (s *QueryFeeBurn) Do(ctx context.Context) (*QueryFeeBurnResponse, error) {
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *QueryFeeBurnResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}
