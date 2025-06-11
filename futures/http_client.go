package futures

import (
	"context"
	"github.com/jekaxv/go-binance/core"
	"net/http"
)

type Client struct {
	C *core.Client
}

func (c *Client) set(key string, value any) {
	c.C.Set(key, value)
}

func (c *Client) invoke(ctx context.Context) error {
	return c.C.Invoke(ctx)
}

func (c *Client) rawBody() []byte {
	return c.C.RawBody()
}

// NewPing Test connectivity
func (c *Client) NewPing() *Ping {
	c.C.SetReq("/fapi/v1/ping", http.MethodGet)
	return &Ping{c: c}
}

// NewServerTime Check server time
func (c *Client) NewServerTime() *ServerTime {
	c.C.SetReq("/fapi/v1/time", http.MethodGet)
	return &ServerTime{c: c}
}

// NewExchangeInfo Exchange information
func (c *Client) NewExchangeInfo() *ExchangeInfo {
	c.C.SetReq("/fapi/v1/exchangeInfo", http.MethodGet)
	return &ExchangeInfo{c: c}
}

// NewDepth order book
func (c *Client) NewDepth() *Depth {
	c.C.SetReq("/fapi/v1/depth", http.MethodGet)
	return &Depth{c: c}
}

// NewTrades Recent trades list
func (c *Client) NewTrades() *Trades {
	c.C.SetReq("/fapi/v1/trades", http.MethodGet)
	return &Trades{c: c}
}

// NewHistoricalTrades Old trade lookup
func (c *Client) NewHistoricalTrades() *HistoricalTrades {
	c.C.SetReq("/fapi/v1/historicalTrades", http.MethodGet, core.AuthApiKey)
	return &HistoricalTrades{c: c}
}

// NewAggTrades Compressed/Aggregate trades list
func (c *Client) NewAggTrades() *AggTrades {
	c.C.SetReq("/fapi/v1/aggTrades", http.MethodGet)
	return &AggTrades{c: c}
}

// NewKline Kline/Candlestick data
func (c *Client) NewKline() *KlineData {
	c.C.SetReq("/fapi/v1/klines", http.MethodGet)
	return &KlineData{c: c}
}

// NewContractKline Continuous Contract Kline/Candlestick Data
func (c *Client) NewContractKline() *ContractKline {
	c.C.SetReq("/fapi/v1/continuousKlines", http.MethodGet)
	return &ContractKline{c: c}
}

// NewIndexKline Index Price Kline/Candlestick Data
func (c *Client) NewIndexKline() *IndexKline {
	c.C.SetReq("/fapi/v1/indexPriceKlines", http.MethodGet)
	return &IndexKline{c: c}
}

// NewMarkKline Mark Price Kline/Candlestick Data
func (c *Client) NewMarkKline() *MarkKline {
	c.C.SetReq("/fapi/v1/markPriceKlines", http.MethodGet)
	return &MarkKline{c: c}
}

// NewPremiumKline Premium index Kline Data
func (c *Client) NewPremiumKline() *PremiumKline {
	c.C.SetReq("/fapi/v1/premiumIndexKlines", http.MethodGet)
	return &PremiumKline{c: c}
}

// NewMarkPrice Mark Price
func (c *Client) NewMarkPrice() *MarkPrice {
	c.C.SetReq("/fapi/v1/premiumIndex", http.MethodGet)
	return &MarkPrice{c: c}
}

// NewFundingRate Get Funding Rate History
func (c *Client) NewFundingRate() *FundingRate {
	c.C.SetReq("/fapi/v1/fundingRate", http.MethodGet)
	return &FundingRate{c: c}
}

// NewFundingInfo Get Funding Rate Info
func (c *Client) NewFundingInfo() *FundingInfo {
	c.C.SetReq("/fapi/v1/fundingInfo", http.MethodGet)
	return &FundingInfo{c: c}
}

// NewTicker24hr 24hr Ticker Price Change Statistics
func (c *Client) NewTicker24hr() *Ticker24hr {
	c.C.SetReq("/fapi/v1/ticker/24hr", http.MethodGet)
	return &Ticker24hr{c: c}
}

// NewTickerPrice Symbol Price Ticker
func (c *Client) NewTickerPrice() *TickerPrice {
	c.C.SetReq("/fapi/v2/ticker/price", http.MethodGet)
	return &TickerPrice{c: c}
}

// NewBookTicker Symbol Order Book Ticker
func (c *Client) NewBookTicker() *BookTicker {
	c.C.SetReq("/fapi/v1/ticker/bookTicker", http.MethodGet)
	return &BookTicker{c: c}
}

// NewDeliveryPrice Quarterly Contract Settlement Price
func (c *Client) NewDeliveryPrice() *DeliveryPrice {
	c.C.SetReq("/futures/data/delivery-price", http.MethodGet)
	return &DeliveryPrice{c: c}
}

// NewOpenInterest Open Interest
func (c *Client) NewOpenInterest() *OpenInterest {
	c.C.SetReq("/fapi/v1/openInterest", http.MethodGet)
	return &OpenInterest{c: c}
}

// NewOpenInterestHist Open Interest Statistics
func (c *Client) NewOpenInterestHist() *OpenInterestHist {
	c.C.SetReq("/futures/data/openInterestHist", http.MethodGet)
	return &OpenInterestHist{c: c}
}

// NewTopTraderPositionsRatio Top Trader Long/Short Ratio (Positions)
func (c *Client) NewTopTraderPositionsRatio() *TopTraderPositionsRatio {
	c.C.SetReq("/futures/data/topLongShortPositionRatio", http.MethodGet)
	return &TopTraderPositionsRatio{c: c}
}

// NewTopTraderAccountsRatio Top Trader Long/Short Ratio (Accounts)
func (c *Client) NewTopTraderAccountsRatio() *TopTraderAccountsRatio {
	c.C.SetReq("/futures/data/topLongShortAccountRatio", http.MethodGet)
	return &TopTraderAccountsRatio{c: c}
}

// NewSymbolRatio Long/Short Ratio
func (c *Client) NewSymbolRatio() *SymbolRatio {
	c.C.SetReq("/futures/data/globalLongShortAccountRatio", http.MethodGet)
	return &SymbolRatio{c: c}
}

// NewTakerVolume Taker Buy/Sell Volume
func (c *Client) NewTakerVolume() *TakerVolume {
	c.C.SetReq("/futures/data/takerlongshortRatio", http.MethodGet)
	return &TakerVolume{c: c}
}

// NewFutureBasis Basis
func (c *Client) NewFutureBasis() *FutureBasis {
	c.C.SetReq("/futures/data/basis", http.MethodGet)
	return &FutureBasis{c: c}
}

// NewIndexInfo Composite Index Symbol Information
func (c *Client) NewIndexInfo() *IndexInfo {
	c.C.SetReq("/fapi/v1/indexInfo", http.MethodGet)
	return &IndexInfo{c: c}
}

// NewAssetIndex Multi-Assets Mode Asset Index
func (c *Client) NewAssetIndex() *AssetIndex {
	c.C.SetReq("/fapi/v1/assetIndex", http.MethodGet)
	return &AssetIndex{c: c}
}

// NewConstituentsPrice Query Index Price Constituents
func (c *Client) NewConstituentsPrice() *ConstituentsPrice {
	c.C.SetReq("/fapi/v1/constituents", http.MethodGet)
	return &ConstituentsPrice{c: c}
}

// NewInsuranceBalance Query Insurance Fund Balance Snapshot
func (c *Client) NewInsuranceBalance() *InsuranceBalance {
	c.C.SetReq("/fapi/v1/insuranceBalance", http.MethodGet)
	return &InsuranceBalance{c: c}
}

// NewCreateOrder New Order(TRADE)
func (c *Client) NewCreateOrder() *CreateOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodPost, core.AuthSigned)
	return &CreateOrder{c: c}
}

// NewPlaceBatchOrder Place Multiple Orders(TRADE)
func (c *Client) NewPlaceBatchOrder() *PlaceBatchOrder {
	c.C.SetReq("/fapi/v1/batchOrders", http.MethodPost, core.AuthSigned)
	return &PlaceBatchOrder{c: c}
}

// NewModifyOrder Modify Order (TRADE)
func (c *Client) NewModifyOrder() *ModifyOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodPut, core.AuthSigned)
	return &ModifyOrder{c: c}
}

// NewModifyMultipleOrder Modify Multiple Orders(TRADE)
func (c *Client) NewModifyMultipleOrder() *ModifyMultipleOrder {
	c.C.SetReq("/fapi/v1/batchOrders", http.MethodPut, core.AuthSigned)
	return &ModifyMultipleOrder{c: c}
}

// NewOrderAmendment Get Order Modify History (USER_DATA)
func (c *Client) NewOrderAmendment() *OrderAmendment {
	c.C.SetReq("/fapi/v1/orderAmendment", http.MethodGet, core.AuthSigned)
	return &OrderAmendment{c: c}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodDelete, core.AuthSigned)
	return &CancelOrder{c: c}
}

// NewCancelMultipleOrder Cancel Multiple Orders (TRADE)
func (c *Client) NewCancelMultipleOrder() *CancelMultipleOrder {
	c.C.SetReq("/fapi/v1/batchOrders", http.MethodDelete, core.AuthSigned)
	return &CancelMultipleOrder{c: c}
}

// NewCancelOpenOrder Cancel All Open Orders (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	c.C.SetReq("/fapi/v1/allOpenOrders", http.MethodDelete, core.AuthSigned)
	return &CancelOpenOrder{c: c}
}

// NewCountdownCancelAll Auto-Cancel All Open Orders (TRADE)
func (c *Client) NewCountdownCancelAll() *CountdownCancelAll {
	c.C.SetReq("/fapi/v1/countdownCancelAll", http.MethodPost, core.AuthSigned)
	return &CountdownCancelAll{c: c}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodGet, core.AuthSigned)
	return &QueryOrder{c: c}
}

// NewQueryAllOrder All Orders (USER_DATA)
func (c *Client) NewQueryAllOrder() *QueryAllOrder {
	c.C.SetReq("/fapi/v1/allOrders", http.MethodGet, core.AuthSigned)
	return &QueryAllOrder{c: c}
}

// NewAllOpenOrder Current All Open Orders (USER_DATA)
func (c *Client) NewAllOpenOrder() *AllOpenOrder {
	c.C.SetReq("/fapi/v1/openOrders", http.MethodGet, core.AuthSigned)
	return &AllOpenOrder{c: c}
}

// NewQueryOpenOrder Query Current Open Order (USER_DATA)
func (c *Client) NewQueryOpenOrder() *QueryOpenOrder {
	c.C.SetReq("/fapi/v1/openOrder", http.MethodGet, core.AuthSigned)
	return &QueryOpenOrder{c: c}
}

// NewForceOrder User's Force Orders (USER_DATA)
func (c *Client) NewForceOrder() *ForceOrder {
	c.C.SetReq("/fapi/v1/forceOrders", http.MethodGet, core.AuthSigned)
	return &ForceOrder{c: c}
}

// NewUserTrades Account Trade List (USER_DATA)
func (c *Client) NewUserTrades() *UserTrades {
	c.C.SetReq("/fapi/v1/userTrades", http.MethodGet, core.AuthSigned)
	return &UserTrades{c: c}
}

// NewChangeMarginType Change Margin Type(TRADE)
func (c *Client) NewChangeMarginType() *ChangeMarginType {
	c.C.SetReq("/fapi/v1/marginType", http.MethodPost, core.AuthSigned)
	return &ChangeMarginType{c: c}
}

// NewChangePositionSide Change Position Mode(TRADE)
func (c *Client) NewChangePositionSide() *ChangePositionSide {
	c.C.SetReq("/fapi/v1/positionSide/dual", http.MethodPost, core.AuthSigned)
	return &ChangePositionSide{c: c}
}

// NewChangeLeverage Change Initial Leverage(TRADE)
func (c *Client) NewChangeLeverage() *ChangeLeverage {
	c.C.SetReq("/fapi/v1/leverage", http.MethodPost, core.AuthSigned)
	return &ChangeLeverage{c: c}
}

// NewChangeMultiAssetsMargin Change Multi-Assets Mode (TRADE)
func (c *Client) NewChangeMultiAssetsMargin() *ChangeMultiAssetsMargin {
	c.C.SetReq("/fapi/v1/multiAssetsMargin", http.MethodPost, core.AuthSigned)
	return &ChangeMultiAssetsMargin{c: c}
}

// NewChangePositionMargin Modify Isolated Position Margin(TRADE)
func (c *Client) NewChangePositionMargin() *ChangePositionMargin {
	c.C.SetReq("/fapi/v1/positionMargin", http.MethodPost, core.AuthSigned)
	return &ChangePositionMargin{c: c}
}

// NewPositionRisk Position Information V3 (USER_DATA)
func (c *Client) NewPositionRisk() *PositionRisk {
	c.C.SetReq("/fapi/v3/positionRisk", http.MethodGet, core.AuthSigned)
	return &PositionRisk{c: c}
}

// NewAdlQuantile Position ADL Quantile Estimation(USER_DATA)
func (c *Client) NewAdlQuantile() *AdlQuantile {
	c.C.SetReq("/fapi/v1/adlQuantile", http.MethodGet, core.AuthSigned)
	return &AdlQuantile{c: c}
}

// NewPositionMarginHistory Get Position Margin Change History (TRADE)
func (c *Client) NewPositionMarginHistory() *PositionMarginHistory {
	c.C.SetReq("/fapi/v1/positionMargin/history", http.MethodGet, core.AuthSigned)
	return &PositionMarginHistory{c: c}
}

// NewCreateTestOrder Test Order(TRADE)
func (c *Client) NewCreateTestOrder() *CreateTestOrder {
	c.C.SetReq("/fapi/v1/order/test", http.MethodPost, core.AuthSigned)
	return &CreateTestOrder{c: c}
}

// NewConvertExchangeInfo List All Convert Pairs
func (c *Client) NewConvertExchangeInfo() *ConvertExchangeInfo {
	c.C.SetReq("/fapi/v1/convert/exchangeInfo", http.MethodGet, core.AuthSigned)
	return &ConvertExchangeInfo{c: c}
}

// NewGetQuote Send Quote Request(USER_DATA)
func (c *Client) NewGetQuote() *GetQuote {
	c.C.SetReq("/fapi/v1/convert/getQuote", http.MethodPost, core.AuthSigned)
	return &GetQuote{c: c}
}

// NewAcceptQuote Accept the offered quote (USER_DATA)
func (c *Client) NewAcceptQuote() *AcceptQuote {
	c.C.SetReq("/fapi/v1/convert/acceptQuote", http.MethodPost, core.AuthSigned)
	return &AcceptQuote{c: c}
}

// NewConvertOrderStatus Order status(USER_DATA)
func (c *Client) NewConvertOrderStatus() *ConvertOrderStatus {
	c.C.SetReq("/fapi/v1/convert/orderStatus", http.MethodGet, core.AuthSigned)
	return &ConvertOrderStatus{c: c}
}

// NewQueryBalance Futures Account Balance V3 (USER_DATA)
func (c *Client) NewQueryBalance() *QueryBalance {
	c.C.SetReq("/fapi/v3/balance", http.MethodGet, core.AuthSigned)
	return &QueryBalance{c: c}
}

// NewAccountInfo Account Information V3(USER_DATA)
func (c *Client) NewAccountInfo() *AccountInfo {
	c.C.SetReq("/fapi/v3/account", http.MethodGet, core.AuthSigned)
	return &AccountInfo{c: c}
}

// NewCommissionRate User Commission Rate (USER_DATA)
func (c *Client) NewCommissionRate() *CommissionRate {
	c.C.SetReq("/fapi/v1/commissionRate", http.MethodGet, core.AuthSigned)
	return &CommissionRate{c: c}
}

// NewAccountConfig Futures Account Configuration(USER_DATA)
func (c *Client) NewAccountConfig() *AccountConfig {
	c.C.SetReq("/fapi/v1/accountConfig", http.MethodGet, core.AuthSigned)
	return &AccountConfig{c: c}
}

// NewSymbolConfig Symbol Configuration(USER_DATA)
func (c *Client) NewSymbolConfig() *SymbolConfig {
	c.C.SetReq("/fapi/v1/symbolConfig", http.MethodGet, core.AuthSigned)
	return &SymbolConfig{c: c}
}

// NewQueryRateLimit Query User Rate Limit (USER_DATA)
func (c *Client) NewQueryRateLimit() *QueryRateLimit {
	c.C.SetReq("/fapi/v1/rateLimit/order", http.MethodGet, core.AuthSigned)
	return &QueryRateLimit{c: c}
}

// NewLeverageBracket Notional and Leverage Brackets (USER_DATA)
func (c *Client) NewLeverageBracket() *LeverageBracket {
	c.C.SetReq("/fapi/v1/leverageBracket", http.MethodGet, core.AuthSigned)
	return &LeverageBracket{c: c}
}

// NewMultiAssetsMargin Get Current Multi-Assets Mode (USER_DATA)
func (c *Client) NewMultiAssetsMargin() *MultiAssetsMargin {
	c.C.SetReq("/fapi/v1/multiAssetsMargin", http.MethodGet, core.AuthSigned)
	return &MultiAssetsMargin{c: c}
}

// NewGetPositionSide Get Current Position Mode(USER_DATA)
func (c *Client) NewGetPositionSide() *GetPositionSide {
	c.C.SetReq("/fapi/v1/positionSide/dual", http.MethodGet, core.AuthSigned)
	return &GetPositionSide{c: c}
}

// NewQueryIncome Get Income History (USER_DATA)
func (c *Client) NewQueryIncome() *QueryIncome {
	c.C.SetReq("/fapi/v1/income", http.MethodGet, core.AuthSigned)
	return &QueryIncome{c: c}
}

// NewTradingStatus Futures Trading Quantitative Rules Indicators (USER_DATA)
func (c *Client) NewTradingStatus() *TradingStatus {
	c.C.SetReq("/fapi/v1/apiTradingStatus", http.MethodGet, core.AuthSigned)
	return &TradingStatus{c: c}
}

// NewTransactionHistory Get Download Id For Futures Transaction History(USER_DATA)
func (c *Client) NewTransactionHistory() *TransactionHistory {
	c.C.SetReq("/fapi/v1/income/asyn", http.MethodGet, core.AuthSigned)
	return &TransactionHistory{c: c}
}

// NewTransactionHistoryLink Get Futures Transaction History Download Link by Id (USER_DATA)
func (c *Client) NewTransactionHistoryLink() *TransactionHistoryLink {
	c.C.SetReq("/fapi/v1/income/asyn/id", http.MethodGet, core.AuthSigned)
	return &TransactionHistoryLink{c: c}
}

// NewOrderHistory Get Download Id For Futures Order History (USER_DATA)
func (c *Client) NewOrderHistory() *OrderHistory {
	c.C.SetReq("/fapi/v1/order/asyn", http.MethodGet, core.AuthSigned)
	return &OrderHistory{c: c}
}

// NewOrderHistoryLink Get Futures Order History Download Link by Id (USER_DATA)
func (c *Client) NewOrderHistoryLink() *OrderHistoryLink {
	c.C.SetReq("/fapi/v1/order/asyn/id", http.MethodGet, core.AuthSigned)
	return &OrderHistoryLink{c: c}
}

// NewTradeHistory Get Download Id For Futures Trade History (USER_DATA)
func (c *Client) NewTradeHistory() *TradeHistory {
	c.C.SetReq("/fapi/v1/trade/asyn", http.MethodGet, core.AuthSigned)
	return &TradeHistory{c: c}
}

// NewTradeHistoryLink Get Futures Trade History Download Link by Id (USER_DATA)
func (c *Client) NewTradeHistoryLink() *TradeHistoryLink {
	c.C.SetReq("/fapi/v1/trade/asyn/id", http.MethodGet, core.AuthSigned)
	return &TradeHistoryLink{c: c}
}

// NewChangeFeeBurn Toggle BNB Burn On Futures Trade (TRADE)
func (c *Client) NewChangeFeeBurn() *ChangeFeeBurn {
	c.C.SetReq("/fapi/v1/feeBurn", http.MethodPost, core.AuthSigned)
	return &ChangeFeeBurn{c: c}
}

// NewQueryFeeBurn Get BNB Burn Status (USER_DATA)
func (c *Client) NewQueryFeeBurn() *QueryFeeBurn {
	c.C.SetReq("/fapi/v1/feeBurn", http.MethodGet, core.AuthSigned)
	return &QueryFeeBurn{c: c}
}

// NewGetListenKey Start User Data Stream (USER_STREAM)
func (c *Client) NewGetListenKey() *GetListenKey {
	c.C.SetReq("/fapi/v1/listenKey", http.MethodPost, core.AuthApiKey)
	return &GetListenKey{c: c}
}

// NewKeepaliveListenKey Keepalive User Data Stream (USER_STREAM)
func (c *Client) NewKeepaliveListenKey() *KeepaliveListenKey {
	c.C.SetReq("/fapi/v1/listenKey", http.MethodPut, core.AuthApiKey)
	return &KeepaliveListenKey{c: c}
}

// NewCloseListenKey Close User Data Stream (USER_STREAM)
func (c *Client) NewCloseListenKey() *CloseListenKey {
	c.C.SetReq("/fapi/v1/listenKey", http.MethodDelete, core.AuthApiKey)
	return &CloseListenKey{c: c}
}
