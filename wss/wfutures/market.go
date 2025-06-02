package wfutures

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/wss"
	"github.com/shopspring/decimal"
)

// Depth Get current order book. Note that this request returns limited market depth
type Depth struct {
	c *Client
}

type DepthResult struct {
	LastUpdateId    int64               `json:"lastUpdateId"`
	OutputTime      uint64              `json:"E"`
	TransactionTime uint64              `json:"T"`
	Bids            [][]decimal.Decimal `json:"bids"` // [0]Price [1] Quantity
	Asks            [][]decimal.Decimal `json:"asks"` // [0]Price [1] Quantity
}

type DepthResponse struct {
	wss.ApiResponse
	Result *DepthResult `json:"result"`
}

func (s *Depth) Symbol(symbol string) *Depth {
	s.c.setParams("symbol", symbol)
	return s
}

// Limit Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000]
func (s *Depth) Limit(limit uint) *Depth {
	s.c.setParams("limit", limit)
	return s
}

func (s *Depth) Do(ctx context.Context) (*DepthResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *DepthResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// TickerPrice Latest price for a symbol or symbols.
type TickerPrice struct {
	c *Client
}

type TickerPriceResult struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
	Time   int64           `json:"time"`
}

type TickerPriceResponse struct {
	wss.ApiResponse
	Result *TickerPriceResult `json:"result"`
}

func (s *TickerPrice) Symbol(symbol string) *TickerPrice {
	s.c.setParams("symbol", symbol)
	return s
}

func (s *TickerPrice) Do(ctx context.Context) (*TickerPriceResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *TickerPriceResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// TickerBook Best price/qty on the order book for a symbol or symbols.
type TickerBook struct {
	c *Client
}

type TickerBookResult struct {
	LastUpdateId int             `json:"lastUpdateId"`
	Symbol       string          `json:"symbol"`
	BidPrice     decimal.Decimal `json:"bidPrice"`
	BidQty       decimal.Decimal `json:"bidQty"`
	AskPrice     decimal.Decimal `json:"askPrice"`
	AskQty       decimal.Decimal `json:"askQty"`
	Time         int64           `json:"time"`
}

type TickerBookResponse struct {
	wss.ApiResponse
	Result *TickerBookResult `json:"result"`
}

func (s *TickerBook) Symbol(symbol string) *TickerBook {
	s.c.setParams("symbol", symbol)
	return s
}

func (s *TickerBook) Do(ctx context.Context) (*TickerBookResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *TickerBookResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
