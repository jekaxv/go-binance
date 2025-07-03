package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// WsAccountBalance Query account balance info
type WsAccountBalance struct {
	c *WsClient
	r *core.WsRequest
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

type WsAccountBalanceResponse struct {
	ApiResponse
	Result []*AccountBalanceResult `json:"result"`
}

func (s *WsAccountBalance) RecvWindow(recvWindow uint) *WsAccountBalance {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsAccountBalance) Do(ctx context.Context) (*WsAccountBalanceResponse, error) {
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
			var resp *WsAccountBalanceResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsAccountInfo Get current account information. User in single-asset/ multi-assets mode will see different value, see comments in response section for detail.
type WsAccountInfo struct {
	c *WsClient
	r *core.WsRequest
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

type WsAccountInfoResponse struct {
	ApiResponse
	Result *AccountInfoResult `json:"result"`
}

func (s *WsAccountInfo) RecvWindow(recvWindow uint) *WsAccountInfo {
	s.r.Set("recvWindow", recvWindow)
	return s
}

func (s *WsAccountInfo) Do(ctx context.Context) (*WsAccountInfoResponse, error) {
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
			var resp *WsAccountInfoResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
