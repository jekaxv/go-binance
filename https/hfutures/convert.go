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
