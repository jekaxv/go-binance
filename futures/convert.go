package futures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/shopspring/decimal"
)

// ConvertExchangeInfo Query for all convertible token pairs and the tokens’ respective upper/lower limits
type ConvertExchangeInfo struct {
	c *Client
	r *core.Request
}

func (s *ConvertExchangeInfo) FromAsset(fromAsset string) *ConvertExchangeInfo {
	s.r.Set("fromAsset", fromAsset)
	return s
}

func (s *ConvertExchangeInfo) ToAsset(toAsset string) *ConvertExchangeInfo {
	s.r.Set("toAsset", toAsset)
	return s
}

type ConvertExchangeInfoResponse struct {
	FromAsset          string          `json:"fromAsset"`
	ToAsset            string          `json:"toAsset"`
	FromAssetMinAmount decimal.Decimal `json:"fromAssetMinAmount"`
	FromAssetMaxAmount decimal.Decimal `json:"fromAssetMaxAmount"`
	ToAssetMinAmount   decimal.Decimal `json:"toAssetMinAmount"`
	ToAssetMaxAmount   decimal.Decimal `json:"toAssetMaxAmount"`
}

func (s *ConvertExchangeInfo) Do(ctx context.Context) ([]*ConvertExchangeInfoResponse, error) {
	resp := make([]*ConvertExchangeInfoResponse, 0)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// GetQuote Request a quote for the requested token pairs
type GetQuote struct {
	c *Client
	r *core.Request
}

func (s *GetQuote) FromAsset(fromAsset string) *GetQuote {
	s.r.Set("fromAsset", fromAsset)
	return s
}

func (s *GetQuote) ToAsset(toAsset string) *GetQuote {
	s.r.Set("toAsset", toAsset)
	return s
}
func (s *GetQuote) FromAmount(fromAmount float64) *GetQuote {
	s.r.Set("fromAmount", fromAmount)
	return s
}
func (s *GetQuote) ToAmount(toAmount float64) *GetQuote {
	s.r.Set("toAmount", toAmount)
	return s
}

// ValidTime 10s, default 10s
func (s *GetQuote) ValidTime(validTime string) *GetQuote {
	s.r.Set("validTime", validTime)
	return s
}
func (s *GetQuote) RecvWindow(recvWindow int64) *GetQuote {
	s.r.Set("recvWindow", recvWindow)
	return s
}

type GetQuoteResponse struct {
	QuoteId        string          `json:"quoteId"`
	Ratio          decimal.Decimal `json:"ratio"`
	InverseRatio   decimal.Decimal `json:"inverseRatio"`
	ValidTimestamp int64           `json:"validTimestamp"`
	ToAmount       decimal.Decimal `json:"toAmount"`
	FromAmount     decimal.Decimal `json:"fromAmount"`
}

func (s *GetQuote) Do(ctx context.Context) (*GetQuoteResponse, error) {
	resp := new(GetQuoteResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// AcceptQuote Accept the offered quote by quote ID.
type AcceptQuote struct {
	c *Client
	r *core.Request
}

func (s *AcceptQuote) QuoteId(quoteId string) *AcceptQuote {
	s.r.Set("quoteId", quoteId)
	return s
}
func (s *AcceptQuote) RecvWindow(recvWindow int64) *AcceptQuote {
	s.r.Set("recvWindow", recvWindow)
	return s
}

type AcceptQuoteResponse struct {
	OrderId     string `json:"orderId"`
	CreateTime  int64  `json:"createTime"`
	OrderStatus string `json:"orderStatus"`
}

func (s *AcceptQuote) Do(ctx context.Context) (*AcceptQuoteResponse, error) {
	resp := new(AcceptQuoteResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}

// ConvertOrderStatus Query order status by order ID.
type ConvertOrderStatus struct {
	c *Client
	r *core.Request
}

func (s *ConvertOrderStatus) QuoteId(quoteId string) *ConvertOrderStatus {
	s.r.Set("quoteId", quoteId)
	return s
}
func (s *ConvertOrderStatus) OrderId(orderId string) *ConvertOrderStatus {
	s.r.Set("orderId", orderId)
	return s
}

type ConvertOrderStatusResponse struct {
	OrderId      int64  `json:"orderId"`
	OrderStatus  string `json:"orderStatus"`
	FromAsset    string `json:"fromAsset"`
	FromAmount   string `json:"fromAmount"`
	ToAsset      string `json:"toAsset"`
	ToAmount     string `json:"toAmount"`
	Ratio        string `json:"ratio"`
	InverseRatio string `json:"inverseRatio"`
	CreateTime   int64  `json:"createTime"`
}

func (s *ConvertOrderStatus) Do(ctx context.Context) (*ConvertOrderStatusResponse, error) {
	resp := new(ConvertOrderStatusResponse)
	if err := s.c.invoke(s.r, ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), resp)
}
