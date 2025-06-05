package wfutures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/wss"
	"github.com/shopspring/decimal"
)

// AccountBalance Query account balance info
type AccountBalance struct {
	c *Client
}

type AccountBalanceResult struct {
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

type AccountBalanceResponse struct {
	wss.ApiResponse
	Result []*AccountBalanceResult `json:"result"`
}

func (s *AccountBalance) RecvWindow(recvWindow uint) *AccountBalance {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *AccountBalance) Do(ctx context.Context) (*AccountBalanceResponse, error) {
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
			var resp *AccountBalanceResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// AccountInfo Get current account information. User in single-asset/ multi-assets mode will see different value, see comments in response section for detail.
type AccountInfo struct {
	c *Client
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

type AccountInfoResult struct {
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

type AccountInfoResponse struct {
	wss.ApiResponse
	Result *AccountInfoResult `json:"result"`
}

func (s *AccountInfo) RecvWindow(recvWindow uint) *AccountInfo {
	s.c.setParams("recvWindow", recvWindow)
	return s
}

func (s *AccountInfo) Do(ctx context.Context) (*AccountInfoResponse, error) {
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
			var resp *AccountInfoResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
