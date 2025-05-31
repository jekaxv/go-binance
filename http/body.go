package http

import (
	"fmt"
	"github.com/shopspring/decimal"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type authType int

const (
	authNone authType = iota
	authApiKey
	authSigned
)

type request struct {
	method   string
	path     string
	authType authType
	query    url.Values
	form     url.Values
	header   http.Header
	body     io.Reader
}

type response struct {
	status  int
	err     error
	rawBody []byte
}

func (r *request) set(key string, value any) *request {
	if r.query == nil {
		r.query = url.Values{}
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		if elems, ok := value.([]string); ok {
			r.query.Set(key, `["`+strings.Join(elems, `","`)+`"]`)
		} else {
			r.query.Set(key, fmt.Sprintf("%v", value))
		}
	default:
		r.query.Set(key, fmt.Sprintf("%v", value))
	}
	return r
}

type SpotFill struct {
	Price           decimal.Decimal `json:"price"`
	Qty             decimal.Decimal `json:"qty"`
	Commission      decimal.Decimal `json:"commission"`
	CommissionAsset string          `json:"commissionAsset"`
	TradeId         int             `json:"tradeId"`
	MatchType       string          `json:"matchType"`
	AllocId         int             `json:"allocId"`
}

type OrderReport struct {
	Symbol                  string          `json:"symbol"`
	OrderId                 int64           `json:"orderId"`
	OrderListId             int             `json:"orderListId"`
	ClientOrderId           string          `json:"clientOrderId"`
	OrigClientOrderId       string          `json:"origClientOrderId"`
	TransactTime            int64           `json:"transactTime"`
	Price                   decimal.Decimal `json:"price"`
	OrigQty                 decimal.Decimal `json:"origQty"`
	ExecutedQty             decimal.Decimal `json:"executedQty"`
	OrigQuoteOrderQty       decimal.Decimal `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     decimal.Decimal `json:"cummulativeQuoteQty"`
	Status                  string          `json:"status"`
	TimeInForce             string          `json:"timeInForce"`
	Type                    string          `json:"type"`
	Side                    string          `json:"side"`
	StopPrice               decimal.Decimal `json:"stopPrice,omitempty"`
	IcebergQty              decimal.Decimal `json:"icebergQty,omitempty"`
	SelfTradePreventionMode string          `json:"selfTradePreventionMode"`
}

type SpotOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type SpotCommission struct {
	Maker  decimal.Decimal `json:"maker"`
	Taker  decimal.Decimal `json:"taker"`
	Buyer  decimal.Decimal `json:"buyer"`
	Seller decimal.Decimal `json:"seller"`
}

type ApiBalance struct {
	Asset  string          `json:"asset"`
	Free   decimal.Decimal `json:"free"`
	Locked decimal.Decimal `json:"locked"`
}
