package futures

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type accountWsTestSuite struct {
	baseWsTestSuite
}

func TestWebsocketAccount(t *testing.T) {
	suite.Run(t, new(accountWsTestSuite))
}

func (s *accountWsTestSuite) TestNewAccountBalance() {
	msg := []byte(`{
	  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",
	  "status": 200,
	  "result": [
		{
		  "accountAlias": "SgsR",
		  "asset": "USDT",
		  "balance": "122607.35137903",
		  "crossWalletBalance": "23.72469206",
		  "crossUnPnl": "0.00000000",
		  "availableBalance": "23.72469206",
		  "maxWithdrawAmount": "23.72469206",
		  "marginAvailable": true,
		  "updateTime": 1617939110373
		}
	  ],
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 2400,
		  "count": 20
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAccountBalance().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsAccountBalanceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountBalanceResponse(resp, testResp)
}

func (s *accountWsTestSuite) assertTestAccountBalanceResponse(r1, r2 *WsAccountBalanceResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	for i := range r1.Result {
		s.assertMarketAccountBalanceResult(r1.Result[i], r2.Result[i])
	}
}

func (s *accountWsTestSuite) assertMarketAccountBalanceResult(r1, r2 *AccountBalanceResult) {
	r := s.r()
	r.Equal(r1.AccountAlias, r2.AccountAlias, "AccountAlias")
	r.Equal(r1.Asset, r2.Asset, "Asset")
	r.Equal(r1.Balance, r2.Balance, "Balance")
	r.Equal(r1.CrossWalletBalance, r2.CrossWalletBalance, "CrossWalletBalance")
	r.Equal(r1.CrossUnPnl, r2.CrossUnPnl, "CrossUnPnl")
	r.Equal(r1.AvailableBalance, r2.AvailableBalance, "AvailableBalance")
	r.Equal(r1.MaxWithdrawAmount, r2.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal(r1.MarginAvailable, r2.MarginAvailable, "MarginAvailable")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "UpdateTime")
}

func (s *accountWsTestSuite) TestNewAccountInfo() {
	msg := []byte(`{
	  "id": "605a6d20-6588-4cb9-afa0-b0ab087507ba",
	  "status": 200,
	  "result": {
		"totalInitialMargin": "0.00000000",
		"totalMaintMargin": "0.00000000",
		"totalWalletBalance": "126.72469206",
		"totalUnrealizedProfit": "0.00000000",
		"totalMarginBalance": "126.72469206",
		"totalPositionInitialMargin": "0.00000000",
		"totalOpenOrderInitialMargin": "0.00000000",
		"totalCrossWalletBalance": "126.72469206",
		"totalCrossUnPnl": "0.00000000",
		"availableBalance": "126.72469206",
		"maxWithdrawAmount": "126.72469206",
		"assets": [
		  {
			"asset": "USDT",
			"walletBalance": "23.72469206",
			"unrealizedProfit": "0.00000000",
			"marginBalance": "23.72469206",
			"maintMargin": "0.00000000",
			"initialMargin": "0.00000000",
			"positionInitialMargin": "0.00000000",
			"openOrderInitialMargin": "0.00000000",
			"crossWalletBalance": "23.72469206",
			"crossUnPnl": "0.00000000",
			"availableBalance": "126.72469206",
			"maxWithdrawAmount": "23.72469206",
			"marginAvailable": true,
			"updateTime": 1625474304765
		  },
		  {
			"asset": "BUSD",
			"walletBalance": "103.12345678",
			"unrealizedProfit": "0.00000000",
			"marginBalance": "103.12345678",
			"maintMargin": "0.00000000",
			"initialMargin": "0.00000000",
			"positionInitialMargin": "0.00000000",
			"openOrderInitialMargin": "0.00000000",
			"crossWalletBalance": "103.12345678",
			"crossUnPnl": "0.00000000",
			"availableBalance": "126.72469206",
			"maxWithdrawAmount": "103.12345678",
			"marginAvailable": true,
			"updateTime": 1625474304765
		  }
		],
		"positions": [
		  {
			"symbol": "BTCUSDT",
			"positionSide": "BOTH",
			"positionAmt": "1.000",
			"unrealizedProfit": "0.00000000",
			"isolatedMargin": "0.00000000",
			"notional": "0",
			"isolatedWallet": "0",
			"initialMargin": "0",
			"maintMargin": "0",
			"updateTime": 0
		  }
		]
	  },
	  "rateLimits": [
		{
		  "rateLimitType": "REQUEST_WEIGHT",
		  "interval": "MINUTE",
		  "intervalNum": 1,
		  "limit": 2400,
		  "count": 20
		}
	  ]
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAccountInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *WsAccountInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountInfoResponse(resp, testResp)
}

func (s *accountWsTestSuite) assertTestAccountInfoResponse(r1, r2 *WsAccountInfoResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertAccountInfoResult(r1.Result, r2.Result)
}

func (s *accountWsTestSuite) assertAccountInfoResult(r1, r2 *AccountInfoResult) {
	r := s.r()
	r.Equal(r1.TotalInitialMargin, r2.TotalInitialMargin, "TotalInitialMargin")
	r.Equal(r1.TotalMaintMargin, r2.TotalMaintMargin, "TotalMaintMargin")
	r.Equal(r1.TotalWalletBalance, r2.TotalWalletBalance, "TotalWalletBalance")
	r.Equal(r1.TotalUnrealizedProfit, r2.TotalUnrealizedProfit, "TotalUnrealizedProfit")
	r.Equal(r1.TotalMarginBalance, r2.TotalMarginBalance, "TotalMarginBalance")
	r.Equal(r1.TotalPositionInitialMargin, r2.TotalPositionInitialMargin, "TotalPositionInitialMargin")
	r.Equal(r1.TotalOpenOrderInitialMargin, r2.TotalOpenOrderInitialMargin, "TotalOpenOrderInitialMargin")
	r.Equal(r1.TotalCrossWalletBalance, r2.TotalCrossWalletBalance, "TotalCrossWalletBalance")
	r.Equal(r1.TotalCrossUnPnl, r2.TotalCrossUnPnl, "TotalCrossUnPnl")
	r.Equal(r1.AvailableBalance, r2.AvailableBalance, "AvailableBalance")
	r.Equal(r1.MaxWithdrawAmount, r2.MaxWithdrawAmount, "MaxWithdrawAmount")
	for i := range r1.Assets {
		s.assertAccountAsset(r1.Assets[i], r2.Assets[i])
	}
	for i := range r1.Positions {
		s.assertAccountPosition(r1.Positions[i], r2.Positions[i])
	}
}

func (s *accountWsTestSuite) assertAccountAsset(r1, r2 *AccountAsset) {
	r := s.r()
	r.Equal(r1.Asset, r2.Asset, "Asset")
	r.Equal(r1.WalletBalance, r2.WalletBalance, "WalletBalance")
	r.Equal(r1.UnrealizedProfit, r2.UnrealizedProfit, "UnrealizedProfit")
	r.Equal(r1.MarginBalance, r2.MarginBalance, "MarginBalance")
	r.Equal(r1.MaintMargin, r2.MaintMargin, "MaintMargin")
	r.Equal(r1.InitialMargin, r2.InitialMargin, "InitialMargin")
	r.Equal(r1.PositionInitialMargin, r2.PositionInitialMargin, "PositionInitialMargin")
	r.Equal(r1.OpenOrderInitialMargin, r2.OpenOrderInitialMargin, "OpenOrderInitialMargin")
	r.Equal(r1.CrossWalletBalance, r2.CrossWalletBalance, "CrossWalletBalance")
	r.Equal(r1.CrossUnPnl, r2.CrossUnPnl, "CrossUnPnl")
	r.Equal(r1.AvailableBalance, r2.AvailableBalance, "AvailableBalance")
	r.Equal(r1.MaxWithdrawAmount, r2.MaxWithdrawAmount, "MaxWithdrawAmount")
	r.Equal(r1.MarginAvailable, r2.MarginAvailable, "MarginAvailable")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "UpdateTime")
}

func (s *accountWsTestSuite) assertAccountPosition(r1, r2 *AccountPosition) {
	r := s.r()
	r.Equal(r1.Symbol, r2.Symbol, "Symbol")
	r.Equal(r1.PositionSide, r2.PositionSide, "PositionSide")
	r.Equal(r1.PositionAmt, r2.PositionAmt, "PositionAmt")
	r.Equal(r1.UnrealizedProfit, r2.UnrealizedProfit, "UnrealizedProfit")
	r.Equal(r1.IsolatedMargin, r2.IsolatedMargin, "IsolatedMargin")
	r.Equal(r1.Notional, r2.Notional, "Notional")
	r.Equal(r1.IsolatedWallet, r2.IsolatedWallet, "IsolatedWallet")
	r.Equal(r1.InitialMargin, r2.InitialMargin, "InitialMargin")
	r.Equal(r1.MaintMargin, r2.MaintMargin, "MaintMargin")
	r.Equal(r1.UpdateTime, r2.UpdateTime, "UpdateTime")
}
