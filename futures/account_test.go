package futures

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type accountTestSuite struct {
	baseHttpTestSuite
}

func TestAccount(t *testing.T) {
	suite.Run(t, new(accountTestSuite))
}

func (s *accountTestSuite) TestNewConvertExchangeInfo() {
	msg := []byte(`[
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
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryBalance().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryBalanceResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range testResp {
		s.assertMarketAccountBalanceResult(resp[i], testResp[i])
	}
}

func (s *accountTestSuite) assertMarketAccountBalanceResult(r1, r2 *QueryBalanceResponse) {
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

func (s *accountTestSuite) TestNewAccountInfo() {
	msg := []byte(`{
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
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAccountInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestAccountInfoResponse(resp, testResp)
}

func (s *accountTestSuite) assertTestAccountInfoResponse(r1, r2 *AccountInfoResponse) {
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

func (s *accountTestSuite) assertAccountAsset(r1, r2 *AccountAsset) {
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

func (s *accountTestSuite) assertAccountPosition(r1, r2 *AccountPosition) {
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

func (s *accountTestSuite) TestNewCommissionRate() {
	msg := []byte(`{
		"symbol": "BTCUSDT",
		"makerCommissionRate": "0.0002",
		"takerCommissionRate": "0.0004"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewCommissionRate().Symbol("BTCUSDT").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *CommissionRateResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Symbol, testResp.Symbol, "Symbol")
	r.Equal(resp.MakerCommissionRate, testResp.MakerCommissionRate, "MakerCommissionRate")
	r.Equal(resp.TakerCommissionRate, testResp.TakerCommissionRate, "TakerCommissionRate")
}

func (s *accountTestSuite) TestNewAccountConfig() {
	msg := []byte(`{   
		"feeTier": 0,             
		"canTrade": true,           
		"canDeposit": true,        
		"canWithdraw": true,       
		"dualSidePosition": true,
		"updateTime": 0,           
		"multiAssetsMargin": false,
		"tradeGroupId": -1
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAccountConfig().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AccountConfigResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.FeeTier, testResp.FeeTier, "FeeTier")
	r.Equal(resp.CanTrade, testResp.CanTrade, "CanTrade")
	r.Equal(resp.CanDeposit, testResp.CanDeposit, "CanDeposit")
	r.Equal(resp.CanWithdraw, testResp.CanWithdraw, "CanWithdraw")
	r.Equal(resp.DualSidePosition, testResp.DualSidePosition, "DualSidePosition")
	r.Equal(resp.UpdateTime, testResp.UpdateTime, "UpdateTime")
	r.Equal(resp.MultiAssetsMargin, testResp.MultiAssetsMargin, "MultiAssetsMargin")
	r.Equal(resp.TradeGroupId, testResp.TradeGroupId, "TradeGroupId")
}

func (s *accountTestSuite) TestNewSymbolConfig() {
	msg := []byte(`[
	  {
	  "symbol": "BTCUSDT", 
	  "marginType": "CROSSED",
	  "isAutoAddMargin": "false",
	  "leverage": 21,
	  "maxNotionalValue": "1000000"
	  }
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewSymbolConfig().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*SymbolConfigResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		r.Equal(resp[i].Symbol, testResp[i].Symbol, "Symbol")
		r.Equal(resp[i].MarginType, testResp[i].MarginType, "MarginType")
		r.Equal(resp[i].IsAutoAddMargin, testResp[i].IsAutoAddMargin, "IsAutoAddMargin")
		r.Equal(resp[i].Leverage, testResp[i].Leverage, "Leverage")
		r.Equal(resp[i].MaxNotionalValue, testResp[i].MaxNotionalValue, "MaxNotionalValue")
	}
}
func (s *accountTestSuite) TestNewQueryRateLimit() {
	msg := []byte(`[
	  {
		"rateLimitType": "ORDERS",
		"interval": "SECOND",
		"intervalNum": 10,
		"limit": 10000
	  },
	  {
		"rateLimitType": "ORDERS",
		"interval": "MINUTE",
		"intervalNum": 1,
		"limit": 20000
	  }
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryRateLimit().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryRateLimitResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		r.Equal(resp[i].RateLimitType, testResp[i].RateLimitType, "RateLimitType")
		r.Equal(resp[i].Interval, testResp[i].Interval, "Interval")
		r.Equal(resp[i].IntervalNum, testResp[i].IntervalNum, "IntervalNum")
		r.Equal(resp[i].Limit, testResp[i].Limit, "Limit")
	}
}
func (s *accountTestSuite) TestNewLeverageBracket() {
	msg := []byte(`[{
		"symbol": "ETHUSDT",
		"notionalCoef": 1.50,
		"brackets": [{
			"bracket": 1,
			"initialLeverage": 75,
			"notionalCap": 10000,
			"notionalFloor": 0,
			"maintMarginRatio": 0.0065,
			"cum":0
		}]
	}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewLeverageBracket().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*LeverageBracketResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		r.Equal(resp[i].Symbol, testResp[i].Symbol, "Symbol")
		r.Equal(resp[i].NotionalCoef, testResp[i].NotionalCoef, "NotionalCoef")
		for j := range resp[i].Brackets {
			r.Equal(resp[i].Brackets[j].Bracket, testResp[i].Brackets[j].Bracket, "Bracket")
			r.Equal(resp[i].Brackets[j].InitialLeverage, testResp[i].Brackets[j].InitialLeverage, "InitialLeverage")
			r.Equal(resp[i].Brackets[j].NotionalCap, testResp[i].Brackets[j].NotionalCap, "NotionalCap")
			r.Equal(resp[i].Brackets[j].NotionalFloor, testResp[i].Brackets[j].NotionalFloor, "NotionalFloor")
			r.Equal(resp[i].Brackets[j].MaintMarginRatio, testResp[i].Brackets[j].MaintMarginRatio, "MaintMarginRatio")
			r.Equal(resp[i].Brackets[j].Cum, testResp[i].Brackets[j].Cum, "Cum")
		}
	}
}

func (s *accountTestSuite) TestNewMultiAssetsMargin() {
	msg := []byte(`{
		"multiAssetsMargin": true 
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewMultiAssetsMargin().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *MultiAssetsMarginResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.MultiAssetsMargin, testResp.MultiAssetsMargin, "MultiAssetsMargin")
}

func (s *accountTestSuite) TestNewGetPositionSide() {
	msg := []byte(`{
		"dualSidePosition": true
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewGetPositionSide().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *GetPositionSideResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.DualSidePosition, testResp.DualSidePosition, "DualSidePosition")
}

func (s *accountTestSuite) TestNewQueryIncome() {
	msg := []byte(`[{
   		"symbol": "BTCUSDT",
    	"incomeType": "COMMISSION", 
    	"income": "-0.01000000",
    	"asset": "USDT",
    	"info":"COMMISSION",
    	"time": 1570636800000,
    	"tranId":9689322392,
    	"tradeId":"2059192"
	}]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryIncome().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*QueryIncomeResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range resp {
		r.Equal(resp[i].Symbol, testResp[i].Symbol, "Symbol")
		r.Equal(resp[i].IncomeType, testResp[i].IncomeType, "IncomeType")
		r.Equal(resp[i].Income, testResp[i].Income, "Income")
		r.Equal(resp[i].Asset, testResp[i].Asset, "Asset")
		r.Equal(resp[i].Info, testResp[i].Info, "Info")
		r.Equal(resp[i].Time, testResp[i].Time, "Time")
		r.Equal(resp[i].TranId, testResp[i].TranId, "TranId")
		r.Equal(resp[i].TradeId, testResp[i].TradeId, "TradeId")
	}
}

func (s *accountTestSuite) TestNewTradingStatus() {
	msg := []byte(`{
		"indicators":{
			"ACCOUNT":[
				{
					"indicator":"TMV", 
					"value":10,
					"triggerValue":1,
					"plannedRecoverTime":1644919865000,
					"isLocked":true
				}
			]
		},
		"updateTime":1644913304748
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTradingStatus().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *TradingStatusResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.UpdateTime, testResp.UpdateTime, "UpdateTime")
	for k := range resp.Indicators {
		r.Equal(resp.Indicators[k], testResp.Indicators[k], "Indicators")
		for j := range resp.Indicators[k] {
			r.Equal(resp.Indicators[k][j].Indicator, testResp.Indicators[k][j].Indicator, "Indicator")
			r.Equal(resp.Indicators[k][j].Value, testResp.Indicators[k][j].Value, "Value")
			r.Equal(resp.Indicators[k][j].TriggerValue, testResp.Indicators[k][j].TriggerValue, "TriggerValue")
			r.Equal(resp.Indicators[k][j].PlannedRecoverTime, testResp.Indicators[k][j].PlannedRecoverTime, "PlannedRecoverTime")
			r.Equal(resp.Indicators[k][j].IsLocked, testResp.Indicators[k][j].IsLocked, "IsLocked")
		}
	}
}

func (s *accountTestSuite) TestNewTransactionHistory() {
	msg := []byte(`{
		"avgCostTimestampOfLast30d":7241837,
		"downloadId":"546975389218332672"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTransactionHistory().StartTime(0).EndTime(0).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *HistoryResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.AvgCostTimestampOfLast30D, testResp.AvgCostTimestampOfLast30D, "AvgCostTimestampOfLast30D")
	r.Equal(resp.DownloadId, testResp.DownloadId, "DownloadId")
}

func (s *accountTestSuite) TestNewTransactionHistoryLink() {
	msg := []byte(`{
		"downloadId":"545923594199212032",
		"status":"processing",
		"url":"", 
		"notified":false,
		"expirationTimestamp":-1,
		"isExpired":null
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTransactionHistoryLink().DownloadId("545923594199212032").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *HistoryLinkResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.DownloadId, testResp.DownloadId, "DownloadId")
	r.Equal(resp.Status, testResp.Status, "Status")
	r.Equal(resp.Url, testResp.Url, "Url")
	r.Equal(resp.Notified, testResp.Notified, "Notified")
	r.Equal(resp.ExpirationTimestamp, testResp.ExpirationTimestamp, "ExpirationTimestamp")
	r.Equal(resp.IsExpired, testResp.IsExpired, "IsExpired")
}

func (s *accountTestSuite) TestNewOrderHistory() {
	msg := []byte(`{
		"avgCostTimestampOfLast30d":7241837,
		"downloadId":"546975389218332672"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOrderHistory().StartTime(0).EndTime(0).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *HistoryResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.AvgCostTimestampOfLast30D, testResp.AvgCostTimestampOfLast30D, "AvgCostTimestampOfLast30D")
	r.Equal(resp.DownloadId, testResp.DownloadId, "DownloadId")
}

func (s *accountTestSuite) TestNewOrderHistoryLink() {
	msg := []byte(`{
		"downloadId":"545923594199212032",
		"status":"processing",
		"url":"", 
		"notified":false,
		"expirationTimestamp":-1,
		"isExpired":null
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewOrderHistoryLink().DownloadId("545923594199212032").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *HistoryLinkResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.DownloadId, testResp.DownloadId, "DownloadId")
	r.Equal(resp.Status, testResp.Status, "Status")
	r.Equal(resp.Url, testResp.Url, "Url")
	r.Equal(resp.Notified, testResp.Notified, "Notified")
	r.Equal(resp.ExpirationTimestamp, testResp.ExpirationTimestamp, "ExpirationTimestamp")
	r.Equal(resp.IsExpired, testResp.IsExpired, "IsExpired")
}
func (s *accountTestSuite) TestNewTradeHistory() {
	msg := []byte(`{
		"avgCostTimestampOfLast30d":7241837,
		"downloadId":"546975389218332672"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTradeHistory().StartTime(0).EndTime(0).Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *HistoryResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.AvgCostTimestampOfLast30D, testResp.AvgCostTimestampOfLast30D, "AvgCostTimestampOfLast30D")
	r.Equal(resp.DownloadId, testResp.DownloadId, "DownloadId")
}

func (s *accountTestSuite) TestNewTradeHistoryLink() {
	msg := []byte(`{
		"downloadId":"545923594199212032",
		"status":"processing",
		"url":"", 
		"notified":false,
		"expirationTimestamp":-1,
		"isExpired":null
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewTradeHistoryLink().DownloadId("545923594199212032").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *HistoryLinkResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.DownloadId, testResp.DownloadId, "DownloadId")
	r.Equal(resp.Status, testResp.Status, "Status")
	r.Equal(resp.Url, testResp.Url, "Url")
	r.Equal(resp.Notified, testResp.Notified, "Notified")
	r.Equal(resp.ExpirationTimestamp, testResp.ExpirationTimestamp, "ExpirationTimestamp")
	r.Equal(resp.IsExpired, testResp.IsExpired, "IsExpired")
}

func (s *accountTestSuite) TestNewChangeFeeBurn() {
	msg := []byte(`{
		"code": 200,
		"msg": "success"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewChangeFeeBurn().FeeBurn("true").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ChangeFeeBurnResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.Code, testResp.Code, "Code")
	r.Equal(resp.Msg, testResp.Msg, "Msg")
}

func (s *accountTestSuite) TestNewQueryFeeBurn() {
	msg := []byte(`{
		"feeBurn": true
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewQueryFeeBurn().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *QueryFeeBurnResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.FeeBurn, testResp.FeeBurn, "FeeBurn")
}
