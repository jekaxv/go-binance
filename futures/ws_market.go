package futures

import (
	"context"
	"encoding/json"
	"github.com/shopspring/decimal"
)

// WsDepth Get current order book. Note that this request returns limited market depth
type WsDepth struct {
	c *WsClient
}

type DepthResult struct {
	LastUpdateId    int64               `json:"lastUpdateId"`
	OutputTime      uint64              `json:"E"`
	TransactionTime uint64              `json:"T"`
	Bids            [][]decimal.Decimal `json:"bids"` // [0]Price [1] Quantity
	Asks            [][]decimal.Decimal `json:"asks"` // [0]Price [1] Quantity
}

type WsDepthResponse struct {
	ApiResponse
	Result *DepthResult `json:"result"`
}

func (s *WsDepth) Symbol(symbol string) *WsDepth {
	s.c.setParams("symbol", symbol)
	return s
}

// Limit Default 500; Valid limits:[5, 10, 20, 50, 100, 500, 1000]
func (s *WsDepth) Limit(limit uint) *WsDepth {
	s.c.setParams("limit", limit)
	return s
}

func (s *WsDepth) Do(ctx context.Context) (*WsDepthResponse, error) {
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
			var resp *WsDepthResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// WsTickerPrice Latest price for a symbol or symbols.
type WsTickerPrice struct {
	c      *WsClient
	symbol *string
}

type TickerPriceResult struct {
	Symbol string          `json:"symbol"`
	Price  decimal.Decimal `json:"price"`
	Time   int64           `json:"time"`
}

type WsTickerPriceResponse struct {
	ApiResponse
	Result []*TickerPriceResult `json:"result"`
}

func (s *WsTickerPrice) Symbol(symbol string) *WsTickerPrice {
	s.c.setParams("symbol", symbol)
	s.symbol = &symbol
	return s
}

func (s *WsTickerPrice) Do(ctx context.Context) (*WsTickerPriceResponse, error) {
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
			resp := new(WsTickerPriceResponse)
			if s.symbol == nil {
				return resp, json.Unmarshal(message, &resp)
			}
			var apiResp ApiResponse
			if err := json.Unmarshal(message, &apiResp); err != nil {
				return nil, err
			}
			resp.ApiResponse = apiResp
			var single *TickerPriceResult
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.Result = append(resp.Result, single)
			return resp, nil
		case err := <-onError:
			return nil, err
		}
	}
}

// WsTickerBook Best price/qty on the order book for a symbol or symbols.
type WsTickerBook struct {
	c      *WsClient
	symbol *string
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

type WsTickerBookResponse struct {
	ApiResponse
	Result []*TickerBookResult `json:"result"`
}

func (s *WsTickerBook) Symbol(symbol string) *WsTickerBook {
	s.c.setParams("symbol", symbol)
	s.symbol = &symbol
	return s
}

func (s *WsTickerBook) Do(ctx context.Context) (*WsTickerBookResponse, error) {
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
			resp := new(WsTickerBookResponse)
			if s.symbol == nil {
				return resp, json.Unmarshal(message, &resp)
			}
			var apiResp ApiResponse
			if err := json.Unmarshal(message, &apiResp); err != nil {
				return nil, err
			}
			resp.ApiResponse = apiResp
			var single *TickerBookResult
			if err := json.Unmarshal(message, &single); err != nil {
				return nil, err
			}
			resp.Result = append(resp.Result, single)
			return resp, nil

		case err := <-onError:
			return nil, err
		}
	}
}
