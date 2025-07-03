package futures

import (
	"context"
	"github.com/jekaxv/go-binance/core"
	"net/http"
)

type Client struct {
	*core.Client
}

func (c *Client) invoke(r *core.Request, ctx context.Context) error {
	return c.Invoke(r, ctx)
}

func (c *Client) rawBody() []byte {
	return c.RawBody()
}

// NewPing Test connectivity
func (c *Client) NewPing() *Ping {
	return &Ping{c: c, r: c.SetReq("/fapi/v1/ping", http.MethodGet)}
}

// NewServerTime Check server time
func (c *Client) NewServerTime() *ServerTime {
	return &ServerTime{c: c, r: c.SetReq("/fapi/v1/time", http.MethodGet)}
}

// NewExchangeInfo Exchange information
func (c *Client) NewExchangeInfo() *ExchangeInfo {
	return &ExchangeInfo{c: c, r: c.SetReq("/fapi/v1/exchangeInfo", http.MethodGet)}
}

// NewDepth order book
func (c *Client) NewDepth() *Depth {
	return &Depth{c: c, r: c.SetReq("/fapi/v1/depth", http.MethodGet)}
}

// NewTrades Recent trades list
func (c *Client) NewTrades() *Trades {
	return &Trades{c: c, r: c.SetReq("/fapi/v1/trades", http.MethodGet)}
}

// NewHistoricalTrades Old trade lookup
func (c *Client) NewHistoricalTrades() *HistoricalTrades {
	return &HistoricalTrades{c: c, r: c.SetReq("/fapi/v1/historicalTrades", http.MethodGet, core.AuthApiKey)}
}

// NewAggTrades Compressed/Aggregate trades list
func (c *Client) NewAggTrades() *AggTrades {
	return &AggTrades{c: c, r: c.SetReq("/fapi/v1/aggTrades", http.MethodGet)}
}

// NewKline Kline/Candlestick data
func (c *Client) NewKline() *KlineData {
	return &KlineData{c: c, r: c.SetReq("/fapi/v1/klines", http.MethodGet)}
}

// NewContractKline Continuous Contract Kline/Candlestick Data
func (c *Client) NewContractKline() *ContractKline {
	return &ContractKline{c: c, r: c.SetReq("/fapi/v1/continuousKlines", http.MethodGet)}
}

// NewIndexKline Index Price Kline/Candlestick Data
func (c *Client) NewIndexKline() *IndexKline {
	return &IndexKline{c: c, r: c.SetReq("/fapi/v1/indexPriceKlines", http.MethodGet)}
}

// NewMarkKline Mark Price Kline/Candlestick Data
func (c *Client) NewMarkKline() *MarkKline {
	return &MarkKline{c: c, r: c.SetReq("/fapi/v1/markPriceKlines", http.MethodGet)}
}

// NewPremiumKline Premium index Kline Data
func (c *Client) NewPremiumKline() *PremiumKline {
	return &PremiumKline{c: c, r: c.SetReq("/fapi/v1/premiumIndexKlines", http.MethodGet)}
}

// NewMarkPrice Mark Price
func (c *Client) NewMarkPrice() *MarkPrice {
	return &MarkPrice{c: c, r: c.SetReq("/fapi/v1/premiumIndex", http.MethodGet)}
}

// NewFundingRate Get Funding Rate History
func (c *Client) NewFundingRate() *FundingRate {
	return &FundingRate{c: c, r: c.SetReq("/fapi/v1/fundingRate", http.MethodGet)}
}

// NewFundingInfo Get Funding Rate Info
func (c *Client) NewFundingInfo() *FundingInfo {
	return &FundingInfo{c: c, r: c.SetReq("/fapi/v1/fundingInfo", http.MethodGet)}
}

// NewTicker24hr 24hr Ticker Price Change Statistics
func (c *Client) NewTicker24hr() *Ticker24hr {
	return &Ticker24hr{c: c, r: c.SetReq("/fapi/v1/ticker/24hr", http.MethodGet)}
}

// NewTickerPrice Symbol Price Ticker
func (c *Client) NewTickerPrice() *TickerPrice {
	return &TickerPrice{c: c, r: c.SetReq("/fapi/v2/ticker/price", http.MethodGet)}
}

// NewBookTicker Symbol Order Book Ticker
func (c *Client) NewBookTicker() *BookTicker {
	return &BookTicker{c: c, r: c.SetReq("/fapi/v1/ticker/bookTicker", http.MethodGet)}
}

// NewDeliveryPrice Quarterly Contract Settlement Price
func (c *Client) NewDeliveryPrice() *DeliveryPrice {
	return &DeliveryPrice{c: c, r: c.SetReq("/futures/data/delivery-price", http.MethodGet)}
}

// NewOpenInterest Open Interest
func (c *Client) NewOpenInterest() *OpenInterest {
	return &OpenInterest{c: c, r: c.SetReq("/fapi/v1/openInterest", http.MethodGet)}
}

// NewOpenInterestHist Open Interest Statistics
func (c *Client) NewOpenInterestHist() *OpenInterestHist {
	return &OpenInterestHist{c: c, r: c.SetReq("/futures/data/openInterestHist", http.MethodGet)}
}

// NewTopTraderPositionsRatio Top Trader Long/Short Ratio (Positions)
func (c *Client) NewTopTraderPositionsRatio() *TopTraderPositionsRatio {
	return &TopTraderPositionsRatio{c: c, r: c.SetReq("/futures/data/topLongShortPositionRatio", http.MethodGet)}
}

// NewTopTraderAccountsRatio Top Trader Long/Short Ratio (Accounts)
func (c *Client) NewTopTraderAccountsRatio() *TopTraderAccountsRatio {
	return &TopTraderAccountsRatio{c: c, r: c.SetReq("/futures/data/topLongShortAccountRatio", http.MethodGet)}
}

// NewSymbolRatio Long/Short Ratio
func (c *Client) NewSymbolRatio() *SymbolRatio {
	return &SymbolRatio{c: c, r: c.SetReq("/futures/data/globalLongShortAccountRatio", http.MethodGet)}
}

// NewTakerVolume Taker Buy/Sell Volume
func (c *Client) NewTakerVolume() *TakerVolume {
	return &TakerVolume{c: c, r: c.SetReq("/futures/data/takerlongshortRatio", http.MethodGet)}
}

// NewFutureBasis Basis
func (c *Client) NewFutureBasis() *FutureBasis {
	return &FutureBasis{c: c, r: c.SetReq("/futures/data/basis", http.MethodGet)}
}

// NewIndexInfo Composite Index Symbol Information
func (c *Client) NewIndexInfo() *IndexInfo {
	return &IndexInfo{c: c, r: c.SetReq("/fapi/v1/indexInfo", http.MethodGet)}
}

// NewAssetIndex Multi-Assets Mode Asset Index
func (c *Client) NewAssetIndex() *AssetIndex {
	return &AssetIndex{c: c, r: c.SetReq("/fapi/v1/assetIndex", http.MethodGet)}
}

// NewConstituentsPrice Query Index Price Constituents
func (c *Client) NewConstituentsPrice() *ConstituentsPrice {
	return &ConstituentsPrice{c: c, r: c.SetReq("/fapi/v1/constituents", http.MethodGet)}
}

// NewInsuranceBalance Query Insurance Fund Balance Snapshot
func (c *Client) NewInsuranceBalance() *InsuranceBalance {
	return &InsuranceBalance{c: c, r: c.SetReq("/fapi/v1/insuranceBalance", http.MethodGet)}
}

// NewCreateOrder New Order(TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	return &CreateOrder{c: c, r: c.SetReq("/fapi/v1/order", http.MethodPost, core.AuthSigned)}
}

// NewPlaceBatchOrder Place Multiple Orders(TRADE)
func (c *Client) NewPlaceBatchOrder() *PlaceBatchOrder {
	return &PlaceBatchOrder{c: c, r: c.SetReq("/fapi/v1/batchOrders", http.MethodPost, core.AuthSigned)}
}

// NewModifyOrder Modify Order (TRADE)
func (c *Client) NewModifyOrder() *ModifyOrder {
	return &ModifyOrder{c: c, r: c.SetReq("/fapi/v1/order", http.MethodPut, core.AuthSigned)}
}

// NewModifyMultipleOrder Modify Multiple Orders(TRADE)
func (c *Client) NewModifyMultipleOrder() *ModifyMultipleOrder {
	return &ModifyMultipleOrder{c: c, r: c.SetReq("/fapi/v1/batchOrders", http.MethodPut, core.AuthSigned)}
}

// NewOrderAmendment Get Order Modify History (USER_DATA)
func (c *Client) NewOrderAmendment() *OrderAmendment {
	return &OrderAmendment{c: c, r: c.SetReq("/fapi/v1/orderAmendment", http.MethodGet, core.AuthSigned)}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	return &CancelOrder{c: c, r: c.SetReq("/fapi/v1/order", http.MethodDelete, core.AuthSigned)}
}

// NewCancelMultipleOrder Cancel Multiple Orders (TRADE)
func (c *Client) NewCancelMultipleOrder() *CancelMultipleOrder {
	return &CancelMultipleOrder{c: c, r: c.SetReq("/fapi/v1/batchOrders", http.MethodDelete, core.AuthSigned)}
}

// NewCancelOpenOrder Cancel All Open Orders (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	return &CancelOpenOrder{c: c, r: c.SetReq("/fapi/v1/allOpenOrders", http.MethodDelete, core.AuthSigned)}
}

// NewCountdownCancelAll Auto-Cancel All Open Orders (TRADE)
func (c *Client) NewCountdownCancelAll() *CountdownCancelAll {
	return &CountdownCancelAll{c: c, r: c.SetReq("/fapi/v1/countdownCancelAll", http.MethodPost, core.AuthSigned)}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	return &QueryOrder{c: c, r: c.SetReq("/fapi/v1/order", http.MethodGet, core.AuthSigned)}
}

// NewQueryAllOrder All Orders (USER_DATA)
func (c *Client) NewQueryAllOrder() *QueryAllOrder {
	return &QueryAllOrder{c: c, r: c.SetReq("/fapi/v1/allOrders", http.MethodGet, core.AuthSigned)}
}

// NewAllOpenOrder Current All Open Orders (USER_DATA)
func (c *Client) NewAllOpenOrder() *AllOpenOrder {
	return &AllOpenOrder{c: c, r: c.SetReq("/fapi/v1/openOrders", http.MethodGet, core.AuthSigned)}
}

// NewQueryOpenOrder Query Current Open Order (USER_DATA)
func (c *Client) NewQueryOpenOrder() *QueryOpenOrder {
	return &QueryOpenOrder{c: c, r: c.SetReq("/fapi/v1/openOrder", http.MethodGet, core.AuthSigned)}
}

// NewForceOrder User's Force Orders (USER_DATA)
func (c *Client) NewForceOrder() *ForceOrder {
	return &ForceOrder{c: c, r: c.SetReq("/fapi/v1/forceOrders", http.MethodGet, core.AuthSigned)}
}

// NewUserTrades Account Trade List (USER_DATA)
func (c *Client) NewUserTrades() *UserTrades {
	return &UserTrades{c: c, r: c.SetReq("/fapi/v1/userTrades", http.MethodGet, core.AuthSigned)}
}

// NewChangeMarginType Change Margin Type(TRADE)
func (c *Client) NewChangeMarginType() *ChangeMarginType {
	return &ChangeMarginType{c: c, r: c.SetReq("/fapi/v1/marginType", http.MethodPost, core.AuthSigned)}
}

// NewChangePositionSide Change Position Mode(TRADE)
func (c *Client) NewChangePositionSide() *ChangePositionSide {
	return &ChangePositionSide{c: c, r: c.SetReq("/fapi/v1/positionSide/dual", http.MethodPost, core.AuthSigned)}
}

// NewChangeLeverage Change Initial Leverage(TRADE)
func (c *Client) NewChangeLeverage() *ChangeLeverage {
	return &ChangeLeverage{c: c, r: c.SetReq("/fapi/v1/leverage", http.MethodPost, core.AuthSigned)}
}

// NewChangeMultiAssetsMargin Change Multi-Assets Mode (TRADE)
func (c *Client) NewChangeMultiAssetsMargin() *ChangeMultiAssetsMargin {
	return &ChangeMultiAssetsMargin{c: c, r: c.SetReq("/fapi/v1/multiAssetsMargin", http.MethodPost, core.AuthSigned)}
}

// NewChangePositionMargin Modify Isolated Position Margin(TRADE)
func (c *Client) NewChangePositionMargin() *ChangePositionMargin {
	return &ChangePositionMargin{c: c, r: c.SetReq("/fapi/v1/positionMargin", http.MethodPost, core.AuthSigned)}
}

// NewPositionRisk Position Information V3 (USER_DATA)
func (c *Client) NewPositionRisk() *PositionRisk {
	return &PositionRisk{c: c, r: c.SetReq("/fapi/v3/positionRisk", http.MethodGet, core.AuthSigned)}
}

// NewAdlQuantile Position ADL Quantile Estimation(USER_DATA)
func (c *Client) NewAdlQuantile() *AdlQuantile {
	return &AdlQuantile{c: c, r: c.SetReq("/fapi/v1/adlQuantile", http.MethodGet, core.AuthSigned)}
}

// NewPositionMarginHistory Get Position Margin Change History (TRADE)
func (c *Client) NewPositionMarginHistory() *PositionMarginHistory {
	return &PositionMarginHistory{c: c, r: c.SetReq("/fapi/v1/positionMargin/history", http.MethodGet, core.AuthSigned)}
}

// NewCreateTestOrder Test Order(TRADE)
func (c *Client) NewCreateTestOrder() *CreateTestOrder {
	return &CreateTestOrder{c: c, r: c.SetReq("/fapi/v1/order/test", http.MethodPost, core.AuthSigned)}
}

// NewConvertExchangeInfo List All Convert Pairs
func (c *Client) NewConvertExchangeInfo() *ConvertExchangeInfo {
	return &ConvertExchangeInfo{c: c, r: c.SetReq("/fapi/v1/convert/exchangeInfo", http.MethodGet, core.AuthSigned)}
}

// NewGetQuote Send Quote Request(USER_DATA)
func (c *Client) NewGetQuote() *GetQuote {
	return &GetQuote{c: c, r: c.SetReq("/fapi/v1/convert/getQuote", http.MethodPost, core.AuthSigned)}
}

// NewAcceptQuote Accept the offered quote (USER_DATA)
func (c *Client) NewAcceptQuote() *AcceptQuote {
	return &AcceptQuote{c: c, r: c.SetReq("/fapi/v1/convert/acceptQuote", http.MethodPost, core.AuthSigned)}
}

// NewConvertOrderStatus Order status(USER_DATA)
func (c *Client) NewConvertOrderStatus() *ConvertOrderStatus {
	return &ConvertOrderStatus{c: c, r: c.SetReq("/fapi/v1/convert/orderStatus", http.MethodGet, core.AuthSigned)}
}

// NewQueryBalance Futures Account Balance V3 (USER_DATA)
func (c *Client) NewQueryBalance() *QueryBalance {
	return &QueryBalance{c: c, r: c.SetReq("/fapi/v3/balance", http.MethodGet, core.AuthSigned)}
}

// NewAccountInfo Account Information V3(USER_DATA)
func (c *Client) NewAccountInfo() *AccountInfo {
	return &AccountInfo{c: c, r: c.SetReq("/fapi/v3/account", http.MethodGet, core.AuthSigned)}
}

// NewCommissionRate User Commission Rate (USER_DATA)
func (c *Client) NewCommissionRate() *CommissionRate {
	return &CommissionRate{c: c, r: c.SetReq("/fapi/v1/commissionRate", http.MethodGet, core.AuthSigned)}
}

// NewAccountConfig Futures Account Configuration(USER_DATA)
func (c *Client) NewAccountConfig() *AccountConfig {
	return &AccountConfig{c: c, r: c.SetReq("/fapi/v1/accountConfig", http.MethodGet, core.AuthSigned)}
}

// NewSymbolConfig Symbol Configuration(USER_DATA)
func (c *Client) NewSymbolConfig() *SymbolConfig {
	return &SymbolConfig{c: c, r: c.SetReq("/fapi/v1/symbolConfig", http.MethodGet, core.AuthSigned)}
}

// NewQueryRateLimit Query User Rate Limit (USER_DATA)
func (c *Client) NewQueryRateLimit() *QueryRateLimit {
	return &QueryRateLimit{c: c, r: c.SetReq("/fapi/v1/rateLimit/order", http.MethodGet, core.AuthSigned)}
}

// NewLeverageBracket Notional and Leverage Brackets (USER_DATA)
func (c *Client) NewLeverageBracket() *LeverageBracket {
	return &LeverageBracket{c: c, r: c.SetReq("/fapi/v1/leverageBracket", http.MethodGet, core.AuthSigned)}
}

// NewMultiAssetsMargin Get Current Multi-Assets Mode (USER_DATA)
func (c *Client) NewMultiAssetsMargin() *MultiAssetsMargin {
	return &MultiAssetsMargin{c: c, r: c.SetReq("/fapi/v1/multiAssetsMargin", http.MethodGet, core.AuthSigned)}
}

// NewGetPositionSide Get Current Position Mode(USER_DATA)
func (c *Client) NewGetPositionSide() *GetPositionSide {
	return &GetPositionSide{c: c, r: c.SetReq("/fapi/v1/positionSide/dual", http.MethodGet, core.AuthSigned)}
}

// NewQueryIncome Get Income History (USER_DATA)
func (c *Client) NewQueryIncome() *QueryIncome {
	return &QueryIncome{c: c, r: c.SetReq("/fapi/v1/income", http.MethodGet, core.AuthSigned)}
}

// NewTradingStatus Futures Trading Quantitative Rules Indicators (USER_DATA)
func (c *Client) NewTradingStatus() *TradingStatus {
	return &TradingStatus{c: c, r: c.SetReq("/fapi/v1/apiTradingStatus", http.MethodGet, core.AuthSigned)}
}

// NewTransactionHistory Get Download Id For Futures Transaction History(USER_DATA)
func (c *Client) NewTransactionHistory() *TransactionHistory {
	return &TransactionHistory{c: c, r: c.SetReq("/fapi/v1/income/asyn", http.MethodGet, core.AuthSigned)}
}

// NewTransactionHistoryLink Get Futures Transaction History Download Link by Id (USER_DATA)
func (c *Client) NewTransactionHistoryLink() *TransactionHistoryLink {
	return &TransactionHistoryLink{c: c, r: c.SetReq("/fapi/v1/income/asyn/id", http.MethodGet, core.AuthSigned)}
}

// NewOrderHistory Get Download Id For Futures Order History (USER_DATA)
func (c *Client) NewOrderHistory() *OrderHistory {
	return &OrderHistory{c: c, r: c.SetReq("/fapi/v1/order/asyn", http.MethodGet, core.AuthSigned)}
}

// NewOrderHistoryLink Get Futures Order History Download Link by Id (USER_DATA)
func (c *Client) NewOrderHistoryLink() *OrderHistoryLink {
	return &OrderHistoryLink{c: c, r: c.SetReq("/fapi/v1/order/asyn/id", http.MethodGet, core.AuthSigned)}
}

// NewTradeHistory Get Download Id For Futures Trade History (USER_DATA)
func (c *Client) NewTradeHistory() *TradeHistory {
	return &TradeHistory{c: c, r: c.SetReq("/fapi/v1/trade/asyn", http.MethodGet, core.AuthSigned)}
}

// NewTradeHistoryLink Get Futures Trade History Download Link by Id (USER_DATA)
func (c *Client) NewTradeHistoryLink() *TradeHistoryLink {
	return &TradeHistoryLink{c: c, r: c.SetReq("/fapi/v1/trade/asyn/id", http.MethodGet, core.AuthSigned)}
}

// NewChangeFeeBurn Toggle BNB Burn On Futures Trade (TRADE)
func (c *Client) NewChangeFeeBurn() *ChangeFeeBurn {
	return &ChangeFeeBurn{c: c, r: c.SetReq("/fapi/v1/feeBurn", http.MethodPost, core.AuthSigned)}
}

// NewQueryFeeBurn Get BNB Burn Status (USER_DATA)
func (c *Client) NewQueryFeeBurn() *QueryFeeBurn {
	return &QueryFeeBurn{c: c, r: c.SetReq("/fapi/v1/feeBurn", http.MethodGet, core.AuthSigned)}
}

// NewGetListenKey Start User Data Stream (USER_STREAM)
func (c *Client) NewGetListenKey() *GetListenKey {
	return &GetListenKey{c: c, r: c.SetReq("/fapi/v1/listenKey", http.MethodPost, core.AuthApiKey)}
}

// NewKeepaliveListenKey Keepalive User Data Stream (USER_STREAM)
func (c *Client) NewKeepaliveListenKey() *KeepaliveListenKey {
	return &KeepaliveListenKey{c: c, r: c.SetReq("/fapi/v1/listenKey", http.MethodPut, core.AuthApiKey)}
}

// NewCloseListenKey Close User Data Stream (USER_STREAM)
func (c *Client) NewCloseListenKey() *CloseListenKey {
	return &CloseListenKey{c: c, r: c.SetReq("/fapi/v1/listenKey", http.MethodDelete, core.AuthApiKey)}
}
