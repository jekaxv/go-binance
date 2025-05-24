package http

import (
	"fmt"
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
	Price           string `json:"price"`
	Qty             string `json:"qty"`
	Commission      string `json:"commission"`
	CommissionAsset string `json:"commissionAsset"`
	TradeId         int    `json:"tradeId"`
	MatchType       string `json:"matchType"`
	AllocId         int    `json:"allocId"`
}

type SpotOrderReport struct {
	Symbol                  string `json:"symbol"`
	OrderId                 int64  `json:"orderId"`
	OrderListId             int    `json:"orderListId"`
	ClientOrderId           string `json:"clientOrderId"`
	OrigClientOrderId       string `json:"origClientOrderId"`
	TransactTime            int64  `json:"transactTime"`
	Price                   string `json:"price"`
	OrigQty                 string `json:"origQty"`
	ExecutedQty             string `json:"executedQty"`
	OrigQuoteOrderQty       string `json:"origQuoteOrderQty"`
	CummulativeQuoteQty     string `json:"cummulativeQuoteQty"`
	Status                  string `json:"status"`
	TimeInForce             string `json:"timeInForce"`
	Type                    string `json:"type"`
	Side                    string `json:"side"`
	StopPrice               string `json:"stopPrice,omitempty"`
	SelfTradePreventionMode string `json:"selfTradePreventionMode"`
}

type SpotOrder struct {
	Symbol        string `json:"symbol"`
	OrderId       int    `json:"orderId"`
	ClientOrderId string `json:"clientOrderId"`
}

type SpotCommission struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

type ApiBalance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
