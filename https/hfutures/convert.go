package hfutures

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
)

// ConvertExchangeInfo Query for all convertible token pairs and the tokens’ respective upper/lower limits
type ConvertExchangeInfo struct {
	c         *Client
	fromAsset *string
	toAsset   *string
}

func (s *ConvertExchangeInfo) FromAsset(fromAsset string) *ConvertExchangeInfo {
	s.fromAsset = &fromAsset
	return s
}

func (s *ConvertExchangeInfo) ToAsset(toAsset string) *ConvertExchangeInfo {
	s.toAsset = &toAsset
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
	if s.fromAsset != nil {
		s.c.set("fromAsset", *s.fromAsset)
	}
	if s.toAsset != nil {
		s.c.set("toAsset", *s.toAsset)
	}
	var resp []*ConvertExchangeInfoResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// GetQuote Request a quote for the requested token pairs
type GetQuote struct {
	c          *Client
	fromAsset  string
	toAsset    string
	fromAmount *float64
	toAmount   *float64
	validTime  *string
	recvWindow *int64
}

func (s *GetQuote) FromAsset(fromAsset string) *GetQuote {
	s.fromAsset = fromAsset
	return s
}

func (s *GetQuote) ToAsset(toAsset string) *GetQuote {
	s.toAsset = toAsset
	return s
}
func (s *GetQuote) FromAmount(fromAmount float64) *GetQuote {
	s.fromAmount = &fromAmount
	return s
}
func (s *GetQuote) ToAmount(toAmount float64) *GetQuote {
	s.toAmount = &toAmount
	return s
}

// ValidTime 10s, default 10s
func (s *GetQuote) ValidTime(validTime string) *GetQuote {
	s.validTime = &validTime
	return s
}
func (s *GetQuote) RecvWindow(recvWindow int64) *GetQuote {
	s.recvWindow = &recvWindow
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
	s.c.set("fromAsset", s.fromAsset)
	s.c.set("toAsset", s.toAsset)
	if s.fromAmount != nil {
		s.c.set("fromAmount", *s.fromAmount)
	}
	if s.toAmount != nil {
		s.c.set("toAmount", *s.toAmount)
	}
	if s.validTime != nil {
		s.c.set("validTime", *s.validTime)
	}
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *GetQuoteResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// AcceptQuote Accept the offered quote by quote ID.
type AcceptQuote struct {
	c          *Client
	quoteId    string
	recvWindow *int64
}

func (s *AcceptQuote) QuoteId(quoteId string) *AcceptQuote {
	s.quoteId = quoteId
	return s
}
func (s *AcceptQuote) RecvWindow(recvWindow int64) *AcceptQuote {
	s.recvWindow = &recvWindow
	return s
}

type AcceptQuoteResponse struct {
	OrderId     string `json:"orderId"`
	CreateTime  int64  `json:"createTime"`
	OrderStatus string `json:"orderStatus"`
}

func (s *AcceptQuote) Do(ctx context.Context) (*AcceptQuoteResponse, error) {
	s.c.set("quoteId", s.quoteId)
	if s.recvWindow != nil {
		s.c.set("recvWindow", *s.recvWindow)
	}
	var resp *AcceptQuoteResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// ConvertOrderStatus Query order status by order ID.
type ConvertOrderStatus struct {
	c       *Client
	quoteId *string
	orderId *string
}

func (s *ConvertOrderStatus) QuoteId(quoteId string) *ConvertOrderStatus {
	s.quoteId = &quoteId
	return s
}
func (s *ConvertOrderStatus) OrderId(orderId string) *ConvertOrderStatus {
	s.orderId = &orderId
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
	if s.quoteId != nil {
		s.c.set("quoteId", *s.quoteId)
	}
	if s.orderId != nil {
		s.c.set("orderId", *s.orderId)
	}
	var resp *ConvertOrderStatusResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}
