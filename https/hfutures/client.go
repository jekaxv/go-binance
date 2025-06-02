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
