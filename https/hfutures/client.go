package hfutures

import (
	"context"
	"github.com/jekaxv/go-binance/https"
	"net/http"
)

type Client struct {
	C *https.Client
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
	c.C.SetReq("/fapi/v1/historicalTrades", http.MethodGet, https.AuthApiKey)
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
	c.C.SetReq("/fapi/v1/order", http.MethodPost, https.AuthSigned)
	return &CreateOrder{c: c}
}

// NewPlaceBatchOrder Place Multiple Orders(TRADE)
func (c *Client) NewPlaceBatchOrder() *PlaceBatchOrder {
	c.C.SetReq("/fapi/v1/batchOrders", http.MethodPost, https.AuthSigned)
	return &PlaceBatchOrder{c: c}
}

// NewModifyOrder Modify Order (TRADE)
func (c *Client) NewModifyOrder() *ModifyOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodPut, https.AuthSigned)
	return &ModifyOrder{c: c}
}

// NewModifyMultipleOrder Modify Multiple Orders(TRADE)
func (c *Client) NewModifyMultipleOrder() *ModifyMultipleOrder {
	c.C.SetReq("/fapi/v1/batchOrders", http.MethodPut, https.AuthSigned)
	return &ModifyMultipleOrder{c: c}
}

// NewOrderAmendment Get Order Modify History (USER_DATA)
func (c *Client) NewOrderAmendment() *OrderAmendment {
	c.C.SetReq("/fapi/v1/orderAmendment", http.MethodGet, https.AuthSigned)
	return &OrderAmendment{c: c}
}

// NewCancelOrder Cancel Order (TRADE)
func (c *Client) NewCancelOrder() *CancelOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodDelete, https.AuthSigned)
	return &CancelOrder{c: c}
}

// NewCancelMultipleOrder Cancel Multiple Orders (TRADE)
func (c *Client) NewCancelMultipleOrder() *CancelMultipleOrder {
	c.C.SetReq("/fapi/v1/batchOrders", http.MethodDelete, https.AuthSigned)
	return &CancelMultipleOrder{c: c}
}

// NewCancelOpenOrder Cancel All Open Orders (TRADE)
func (c *Client) NewCancelOpenOrder() *CancelOpenOrder {
	c.C.SetReq("/fapi/v1/allOpenOrders", http.MethodDelete, https.AuthSigned)
	return &CancelOpenOrder{c: c}
}

// NewCountdownCancelAll Auto-Cancel All Open Orders (TRADE)
func (c *Client) NewCountdownCancelAll() *CountdownCancelAll {
	c.C.SetReq("/fapi/v1/countdownCancelAll", http.MethodPost, https.AuthSigned)
	return &CountdownCancelAll{c: c}
}

// NewQueryOrder Query Order (USER_DATA)
func (c *Client) NewQueryOrder() *QueryOrder {
	c.C.SetReq("/fapi/v1/order", http.MethodGet, https.AuthSigned)
	return &QueryOrder{c: c}
}

// NewQueryAllOrder All Orders (USER_DATA)
func (c *Client) NewQueryAllOrder() *QueryAllOrder {
	c.C.SetReq("/fapi/v1/allOrders", http.MethodGet, https.AuthSigned)
	return &QueryAllOrder{c: c}
}

// NewAllOpenOrder Current All Open Orders (USER_DATA)
func (c *Client) NewAllOpenOrder() *AllOpenOrder {
	c.C.SetReq("/fapi/v1/openOrders", http.MethodGet, https.AuthSigned)
	return &AllOpenOrder{c: c}
}

// NewQueryOpenOrder Query Current Open Order (USER_DATA)
func (c *Client) NewQueryOpenOrder() *QueryOpenOrder {
	c.C.SetReq("/fapi/v1/openOrder", http.MethodGet, https.AuthSigned)
	return &QueryOpenOrder{c: c}
}

// NewForceOrder User's Force Orders (USER_DATA)
func (c *Client) NewForceOrder() *ForceOrder {
	c.C.SetReq("/fapi/v1/forceOrders", http.MethodGet, https.AuthSigned)
	return &ForceOrder{c: c}
}

// NewUserTrades Account Trade List (USER_DATA)
func (c *Client) NewUserTrades() *UserTrades {
	c.C.SetReq("/fapi/v1/userTrades", http.MethodGet, https.AuthSigned)
	return &UserTrades{c: c}
}

// NewChangeMarginType Change Margin Type(TRADE)
func (c *Client) NewChangeMarginType() *ChangeMarginType {
	c.C.SetReq("/fapi/v1/marginType", http.MethodPost, https.AuthSigned)
	return &ChangeMarginType{c: c}
}

// NewChangePositionSide Change Position Mode(TRADE)
func (c *Client) NewChangePositionSide() *ChangePositionSide {
	c.C.SetReq("/fapi/v1/positionSide/dual", http.MethodPost, https.AuthSigned)
	return &ChangePositionSide{c: c}
}

// NewChangeLeverage Change Initial Leverage(TRADE)
func (c *Client) NewChangeLeverage() *ChangeLeverage {
	c.C.SetReq("/fapi/v1/leverage", http.MethodPost, https.AuthSigned)
	return &ChangeLeverage{c: c}
}

// NewChangeMultiAssetsMargin Change Multi-Assets Mode (TRADE)
func (c *Client) NewChangeMultiAssetsMargin() *ChangeMultiAssetsMargin {
	c.C.SetReq("/fapi/v1/leverage", http.MethodPost, https.AuthSigned)
	return &ChangeMultiAssetsMargin{c: c}
}
