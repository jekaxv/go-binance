package wfutures

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jekaxv/go-binance/types"
	"github.com/shopspring/decimal"
	"strings"
)

type WebsocketStreams struct {
	c *Client
}

// AggTradeService The Aggregate Trade Streams push market trade information that is aggregated for fills with same price and taking side every 100 milliseconds.
// Only market trades will be aggregated, which means the insurance fund trades and ADL trades won't be aggregated.
type AggTradeService struct {
	*WebsocketStreams
}

type AggTradeEvent struct {
	Event        string          `json:"e"`
	Time         int64           `json:"E"`
	Symbol       string          `json:"s"`
	AggTradeID   int64           `json:"a"`
	Price        decimal.Decimal `json:"p"`
	Quantity     decimal.Decimal `json:"q"`
	FirstTradeID int64           `json:"f"`
	LastTradeID  int64           `json:"l"`
	TradeTime    int64           `json:"T"`
	IsBuyerMaker bool            `json:"m"`
	Placeholder  bool            `json:"M"` // Ignore
}

// SubscribeAggTrade Stream Name: <symbol>@aggTrade
func (s *WebsocketStreams) SubscribeAggTrade(symbol string) *AggTradeService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@aggTrade", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &AggTradeService{s}
}

func (e *AggTradeService) Do(ctx context.Context) (<-chan *AggTradeEvent, <-chan error) {
	messageCh := make(chan *AggTradeEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *AggTradeEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedAggTradeService struct {
	*WebsocketStreams
}

type CombinedAggTradeEvent struct {
	Stream string         `json:"stream"`
	Data   *AggTradeEvent `json:"data"`
}

func (s *WebsocketStreams) SubscribeCombinedAggTrade(symbols []string) *CombinedAggTradeService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedAggTradeService{s}
}

func (e *CombinedAggTradeService) Do(ctx context.Context) (<-chan *CombinedAggTradeEvent, <-chan error) {
	messageCh := make(chan *CombinedAggTradeEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedAggTradeEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// MarkPriceService Mark price and funding rate for a single symbol pushed every 3 seconds or every second.
type MarkPriceService struct {
	*WebsocketStreams
}

type MarkPriceEvent struct {
	EventType       string          `json:"e"`
	EventTime       int64           `json:"E"`
	Symbol          string          `json:"s"`
	MarkPrice       decimal.Decimal `json:"p"`
	IndexPrice      decimal.Decimal `json:"i"`
	SettlePrice     decimal.Decimal `json:"P"`
	FundingRate     decimal.Decimal `json:"r"`
	NextFundingTime int64           `json:"T"`
}

// SubscribeMarkPrice Stream Name: <symbol>@markPrice
func (s *WebsocketStreams) SubscribeMarkPrice(symbol string) *MarkPriceService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@markPrice@1s", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &MarkPriceService{s}
}

func (e *MarkPriceService) Do(ctx context.Context) (<-chan *MarkPriceEvent, <-chan error) {
	messageCh := make(chan *MarkPriceEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *MarkPriceEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedMarkPriceService struct {
	*WebsocketStreams
}

type CombinedMarkPriceEvent struct {
	Stream string          `json:"stream"`
	Data   *MarkPriceEvent `json:"data"`
}

func (s *WebsocketStreams) SubscribeCombinedMarkPrice(symbols []string) *CombinedMarkPriceService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@aggTrade", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedMarkPriceService{s}
}

func (e *CombinedMarkPriceService) Do(ctx context.Context) (<-chan *CombinedMarkPriceEvent, <-chan error) {
	messageCh := make(chan *CombinedMarkPriceEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedMarkPriceEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// MarkPriceArrService Mark price and funding rate for all symbols pushed every 3 seconds or every second.
type MarkPriceArrService struct {
	*WebsocketStreams
}

// SubscribeMarkPriceArr Stream Name: !markPrice@arr
func (s *WebsocketStreams) SubscribeMarkPriceArr() *MarkPriceArrService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!markPrice@arr@1s", s.c.getEndpoint()))
	return &MarkPriceArrService{s}
}

func (e *MarkPriceArrService) Do(ctx context.Context) (<-chan []*MarkPriceEvent, <-chan error) {
	messageCh := make(chan []*MarkPriceEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event []*MarkPriceEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// KlineService The Kline/Candlestick Stream push updates to the current klines/candlestick every 250 milliseconds (if existing).
type KlineService struct {
	*WebsocketStreams
}

type KlineResult struct {
	StartTime        int64           `json:"t"`
	CloseTime        int64           `json:"T"`
	Symbol           string          `json:"s"`
	Interval         string          `json:"i"`
	FirstTradeId     int             `json:"f"`
	LastTradeId      int             `json:"L"`
	OpenPrice        decimal.Decimal `json:"o"`
	ClosePrice       decimal.Decimal `json:"c"`
	HighPrice        decimal.Decimal `json:"h"`
	LowPrice         decimal.Decimal `json:"l"`
	BaseAssetVolume  decimal.Decimal `json:"v"`
	NumberOfTrades   int             `json:"n"`
	IsClosed         bool            `json:"x"`
	QuoteAssetVolume decimal.Decimal `json:"q"`
	TakerBaseVolume  decimal.Decimal `json:"V"`
	TakerQuoteVolume decimal.Decimal `json:"Q"`
	Placeholder      string          `json:"B"`
}

type KlineEvent struct {
	Event  string       `json:"e"`
	Time   int64        `json:"E"`
	Symbol string       `json:"s"`
	Kline  *KlineResult `json:"k"`
}

// SubscribeKline Stream Name: <symbol>@kline_<interval>
func (s *WebsocketStreams) SubscribeKline(symbol string, interval types.IntervalEnum) *KlineService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@kline_%s", s.c.getEndpoint(), strings.ToLower(symbol), interval))
	return &KlineService{s}
}

func (e *KlineService) Do(ctx context.Context) (<-chan *KlineEvent, <-chan error) {
	messageCh := make(chan *KlineEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *KlineEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedKlineService struct {
	*WebsocketStreams
}

type CombinedKlineEvent struct {
	Stream string      `json:"stream"`
	Data   *KlineEvent `json:"data"`
}

// SubscribeCombinedKline Stream Name: <symbol>@kline_<interval>
func (s *WebsocketStreams) SubscribeCombinedKline(symbols map[string]string) *CombinedKlineService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for symbol, interval := range symbols {
		endpoint += fmt.Sprintf("%s@kline_%s", strings.ToLower(symbol), interval) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedKlineService{s}
}

func (e *CombinedKlineService) Do(ctx context.Context) (<-chan *CombinedKlineEvent, <-chan error) {
	messageCh := make(chan *CombinedKlineEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedKlineEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// ContractKlineService Continuous Contract Kline/Candlestick Streams
type ContractKlineService struct {
	*WebsocketStreams
}

type ContractKlineEvent struct {
	Event        string       `json:"e"`
	Time         int64        `json:"E"`
	Symbol       string       `json:"ps"`
	ContractType string       `json:"ct"`
	Kline        *KlineResult `json:"k"`
}

// SubscribeContractKline Stream Name: <pair>_<contractType>@continuousKline_<interval>
func (s *WebsocketStreams) SubscribeContractKline(symbol string, contractType types.ContractType, interval types.IntervalEnum) *ContractKlineService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s_%s@continuousKline_%s", s.c.getEndpoint(), strings.ToLower(symbol), strings.ToLower(string(contractType)), interval))
	return &ContractKlineService{s}
}

func (e *ContractKlineService) Do(ctx context.Context) (<-chan *ContractKlineEvent, <-chan error) {
	messageCh := make(chan *ContractKlineEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *ContractKlineEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// MiniTickerService 24hr rolling window mini-ticker statistics for a single symbol.
// These are NOT the statistics of the UTC day, but a 24hr rolling window from requestTime to 24hrs before.
type MiniTickerService struct {
	*WebsocketStreams
}

type MiniTickerEvent struct {
	Event                  string          `json:"e"`
	Time                   int64           `json:"E"`
	Symbol                 string          `json:"s"`
	OpenPrice              decimal.Decimal `json:"o"`
	ClosePrice             decimal.Decimal `json:"c"`
	HighPrice              decimal.Decimal `json:"h"`
	LowPrice               decimal.Decimal `json:"l"`
	TotalTradedBaseVolume  decimal.Decimal `json:"v"`
	TotalTradedQuoteVolume decimal.Decimal `json:"q"`
}

// SubscribeMiniTicker Stream Name: <symbol>@miniTicker
func (s *WebsocketStreams) SubscribeMiniTicker(symbol string) *MiniTickerService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@miniTicker", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &MiniTickerService{s}
}

func (e *MiniTickerService) Do(ctx context.Context) (<-chan *MiniTickerEvent, <-chan error) {
	messageCh := make(chan *MiniTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *MiniTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedMiniTickerService struct {
	*WebsocketStreams
}

type CombinedMiniTickerEvent struct {
	Stream string           `json:"stream"`
	Data   *MiniTickerEvent `json:"data"`
}

// SubscribeCombinedMiniTicker Stream Name: <symbol>@miniTicker
func (s *WebsocketStreams) SubscribeCombinedMiniTicker(symbols []string) *CombinedMiniTickerService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@miniTicker", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedMiniTickerService{s}
}

func (e *CombinedMiniTickerService) Do(ctx context.Context) (<-chan *CombinedMiniTickerEvent, <-chan error) {
	messageCh := make(chan *CombinedMiniTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedMiniTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// TickerService 24hr rolling window ticker statistics for a single symbol.
// These are NOT the statistics of the UTC day, but a 24hr rolling window from requestTime to 24hrs before.
type TickerService struct {
	*WebsocketStreams
}

type TickerEvent struct {
	Event                  string          `json:"e"`
	Time                   int64           `json:"E"`
	Symbol                 string          `json:"s"`
	PriceChange            decimal.Decimal `json:"p"`
	PriceChangePercent     decimal.Decimal `json:"P"`
	WeightedAveragePrice   decimal.Decimal `json:"w"`
	LastTradePrice         decimal.Decimal `json:"c"`
	LastQuantity           decimal.Decimal `json:"Q"`
	OpenPrice              decimal.Decimal `json:"o"`
	HighPrice              decimal.Decimal `json:"h"`
	LowPrice               decimal.Decimal `json:"l"`
	TotalTradedBaseVolume  decimal.Decimal `json:"v"`
	TotalTradedQuoteVolume decimal.Decimal `json:"q"`
	OpenTime               int             `json:"O"`
	CloseTime              int             `json:"C"`
	FirstTradeID           int             `json:"F"`
	LastTradeID            int             `json:"L"`
	NumberOfTrades         int             `json:"n"`
}

// SubscribeTicker Stream Name: <symbol>@ticker
func (s *WebsocketStreams) SubscribeTicker(symbol string) *TickerService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@ticker", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &TickerService{s}
}

func (e *TickerService) Do(ctx context.Context) (<-chan *TickerEvent, <-chan error) {
	messageCh := make(chan *TickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *TickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedTickerService struct {
	*WebsocketStreams
}

type CombinedTickerEvent struct {
	Stream string       `json:"stream"`
	Data   *TickerEvent `json:"data"`
}

// SubscribeCombinedTicker Stream Name: <symbol>@ticker
func (s *WebsocketStreams) SubscribeCombinedTicker(symbols []string) *CombinedTickerService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@ticker", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedTickerService{s}
}

func (e *CombinedTickerService) Do(ctx context.Context) (<-chan *CombinedTickerEvent, <-chan error) {
	messageCh := make(chan *CombinedTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// TickerArrService 24hr rolling window ticker statistics for all symbols that changed in an array.
// These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
// Note that only tickers that have changed will be present in the array.
type TickerArrService struct {
	*WebsocketStreams
}

// SubscribeTickerArr Stream Name: !ticker@arr
func (s *WebsocketStreams) SubscribeTickerArr() *TickerArrService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!ticker@arr", s.c.getEndpoint()))
	return &TickerArrService{s}
}

func (e *TickerArrService) Do(ctx context.Context) (<-chan []*TickerEvent, <-chan error) {
	messageCh := make(chan []*TickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event []*TickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// MiniTickerArrService 24hr rolling window mini-ticker statistics for all symbols that changed in an array.
// These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
// Note that only tickers that have changed will be present in the array.
type MiniTickerArrService struct {
	*WebsocketStreams
}

// SubscribeMiniTickerArr Stream Name: !miniTicker@arr
func (s *WebsocketStreams) SubscribeMiniTickerArr() *MiniTickerArrService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!miniTicker@arr", s.c.getEndpoint()))
	return &MiniTickerArrService{s}
}

func (e *MiniTickerArrService) Do(ctx context.Context) (<-chan []*MiniTickerEvent, <-chan error) {
	messageCh := make(chan []*MiniTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event []*MiniTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// BookTickerService Pushes any update to the best bid or ask's price or quantity in real-time for a specified symbol.
// Multiple <symbol>@bookTicker streams can be subscribed to over one connection.
type BookTickerService struct {
	*WebsocketStreams
}

type BookTickerEvent struct {
	Event           string          `json:"e"`
	EventTime       int64           `json:"E"`
	TransactionTime int64           `json:"T"`
	UpdateId        int             `json:"u"`
	Symbol          string          `json:"s"`
	BestBidPrice    decimal.Decimal `json:"b"`
	BestBidQty      decimal.Decimal `json:"B"`
	BestAskPrice    decimal.Decimal `json:"a"`
	BestAskQty      decimal.Decimal `json:"A"`
}

// SubscribeBookTicker Stream Name: <symbol>@bookTicker
func (s *WebsocketStreams) SubscribeBookTicker(symbol string) *BookTickerService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@bookTicker", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &BookTickerService{s}
}

func (e *BookTickerService) Do(ctx context.Context) (<-chan *BookTickerEvent, <-chan error) {
	messageCh := make(chan *BookTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *BookTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedBookTickerService struct {
	*WebsocketStreams
}

type CombinedBookTickerEvent struct {
	Stream string           `json:"stream"`
	Data   *BookTickerEvent `json:"data"`
}

// SubscribeCombinedBookTicker Stream Name: <symbol>@bookTicker
func (s *WebsocketStreams) SubscribeCombinedBookTicker(symbols []string) *CombinedBookTickerService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@bookTicker", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedBookTickerService{s}
}

func (e *CombinedBookTickerService) Do(ctx context.Context) (<-chan *CombinedBookTickerEvent, <-chan error) {
	messageCh := make(chan *CombinedBookTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedBookTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// BestBookTickerService Pushes any update to the best bid or ask's price or quantity in real-time for all symbols.
type BestBookTickerService struct {
	*WebsocketStreams
}

// SubscribeBestBookTicker Stream Name: !bookTicker
func (s *WebsocketStreams) SubscribeBestBookTicker() *BestBookTickerService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!bookTicker", s.c.getEndpoint()))
	return &BestBookTickerService{s}
}

func (e *BestBookTickerService) Do(ctx context.Context) (<-chan *BookTickerEvent, <-chan error) {
	messageCh := make(chan *BookTickerEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *BookTickerEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// ForceOrderService The Liquidation Order Snapshot Streams push force liquidation order information for specific symbol.
// For each symbol，only the latest one liquidation order within 1000ms will be pushed as the snapshot.
// If no liquidation happens in the interval of 1000ms, no stream will be pushed.
type ForceOrderService struct {
	*WebsocketStreams
}

type ForceOrderEvent struct {
	Event string `json:"e"`
	Time  int64  `json:"E"`
	O     struct {
		Symbol              string          `json:"s"`
		Side                string          `json:"S"`
		OrderType           string          `json:"o"`
		TimeInForce         string          `json:"f"`
		Quantity            decimal.Decimal `json:"q"`
		Price               decimal.Decimal `json:"p"`
		AveragePrice        decimal.Decimal `json:"ap"`
		Status              string          `json:"X"`
		FilledQuantity      decimal.Decimal `json:"l"`
		AccumulatedQuantity decimal.Decimal `json:"z"`
		TradeTime           int64           `json:"T"`
	} `json:"o"`
}

// SubscribeForceOrder Stream Name: <symbol>@forceOrder
func (s *WebsocketStreams) SubscribeForceOrder(symbol string) *ForceOrderService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@forceOrder", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &ForceOrderService{s}
}

func (e *ForceOrderService) Do(ctx context.Context) (<-chan *ForceOrderEvent, <-chan error) {
	messageCh := make(chan *ForceOrderEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *ForceOrderEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedForceOrderService struct {
	*WebsocketStreams
}

type CombinedForceOrderEvent struct {
	Stream string           `json:"stream"`
	Data   *ForceOrderEvent `json:"data"`
}

// SubscribeCombinedForceOrder Stream Name: <symbol>@forceOrder
func (s *WebsocketStreams) SubscribeCombinedForceOrder(symbols []string) *CombinedForceOrderService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@forceOrder", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedForceOrderService{s}
}

func (e *CombinedForceOrderService) Do(ctx context.Context) (<-chan *CombinedForceOrderEvent, <-chan error) {
	messageCh := make(chan *CombinedForceOrderEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedForceOrderEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// AllForceOrderService The All Liquidation Order Snapshot Streams push force liquidation order information for all symbols in the market.
// For each symbol，only the latest one liquidation order within 1000ms will be pushed as the snapshot.
// If no liquidation happens in the interval of 1000ms, no stream will be pushed.
type AllForceOrderService struct {
	*WebsocketStreams
}

// SubscribeAllForceOrder Stream Name: !forceOrder@arr
func (s *WebsocketStreams) SubscribeAllForceOrder() *AllForceOrderService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!forceOrder@arr", s.c.getEndpoint()))
	return &AllForceOrderService{s}
}

func (e *AllForceOrderService) Do(ctx context.Context) (<-chan *ForceOrderEvent, <-chan error) {
	messageCh := make(chan *ForceOrderEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *ForceOrderEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// DepthService Diff. Book Depth Streams
type DepthService struct {
	*WebsocketStreams
}

type DepthEvent struct {
	Event           string              `json:"e"`
	Time            int64               `json:"E"`
	TransactionTime int64               `json:"T"`
	Symbol          string              `json:"s"`
	FirstId         int64               `json:"U"`
	FinalId         int64               `json:"u"`
	LastId          int64               `json:"pu"`
	Bids            [][]decimal.Decimal `json:"b"` // Bids to be updated. [0]Price level to be updated, [1]Quantity
	Asks            [][]decimal.Decimal `json:"a"` // Asks to be updated. [0]Price level to be updated, [1]Quantity
}

// SubscribeDepth Stream Names: <symbol>@depth OR <symbol>@depth@100ms
// Update Speed: 250ms, 500ms, 100ms
func (s *WebsocketStreams) SubscribeDepth(symbol string, interval ...string) *DepthService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@depth", s.c.getEndpoint(), strings.ToLower(symbol)))
	if len(interval) != 0 {
		s.c.setEndpoint(fmt.Sprintf("%s@%s", s.c.getEndpoint(), interval[0]))
	}
	return &DepthService{s}
}

func (e *DepthService) Do(ctx context.Context) (<-chan *DepthEvent, <-chan error) {
	messageCh := make(chan *DepthEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *DepthEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedDepthService struct {
	*WebsocketStreams
}

type CombinedDepthEvent struct {
	Stream string      `json:"stream"`
	Data   *DepthEvent `json:"data"`
}

// SubscribeCombinedDepth Stream Names: <symbol>@depth OR <symbol>@depth@100ms
// Update Speed: 1000ms or 100ms
func (s *WebsocketStreams) SubscribeCombinedDepth(symbols []string, interval ...string) *CombinedDepthService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@depth", strings.ToLower(symbol))
		if len(interval) != 0 {
			endpoint += "@" + interval[0]
		}
		endpoint += "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedDepthService{s}
}

func (e *CombinedDepthService) Do(ctx context.Context) (<-chan *CombinedDepthEvent, <-chan error) {
	messageCh := make(chan *CombinedDepthEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedDepthEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// DepthLevelService Top <levels> bids and asks, pushed every second. Valid <levels> are 5, 10, or 20.
type DepthLevelService struct {
	*WebsocketStreams
}

// SubscribeDepthLevel Stream Names: <symbol>@depth<levels> OR <symbol>@depth<levels>@100ms
// Top <levels> bids and asks, pushed every second. Valid <levels> are 5, 10, or 20.
// Update Speed: 1000ms or 100ms
func (s *WebsocketStreams) SubscribeDepthLevel(symbol string, level int, interval ...string) *DepthLevelService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@depth%d", s.c.getEndpoint(), strings.ToLower(symbol), level))
	if len(interval) != 0 {
		s.c.setEndpoint(fmt.Sprintf("%s@%s", s.c.getEndpoint(), interval[0]))
	}
	return &DepthLevelService{s}
}

func (e *DepthLevelService) Do(ctx context.Context) (<-chan *DepthEvent, <-chan error) {
	messageCh := make(chan *DepthEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *DepthEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedDepthLevelService struct {
	*WebsocketStreams
}

type CombinedDepthLevelEvent struct {
	Stream string      `json:"stream"`
	Data   *DepthEvent `json:"data"`
}

// SubscribeCombinedDepthLevel Stream Names: <symbol>@depth<levels> OR <symbol>@depth<levels>@100ms
// Top <levels> bids and asks, pushed every second. Valid <levels> are 5, 10, or 20.
// Update Speed: 1000ms or 100ms
func (s *WebsocketStreams) SubscribeCombinedDepthLevel(symbols map[string]int, interval ...string) *CombinedDepthLevelService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for symbol, level := range symbols {
		endpoint += fmt.Sprintf("%s@depth%d", strings.ToLower(symbol), level)
		if len(interval) != 0 {
			endpoint += "@" + interval[0]
		}
		endpoint += "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedDepthLevelService{s}
}

func (e *CombinedDepthLevelService) Do(ctx context.Context) (<-chan *CombinedDepthLevelEvent, <-chan error) {
	messageCh := make(chan *CombinedDepthLevelEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedDepthLevelEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// CompositeIndexService Composite index information for index symbols pushed every second.
type CompositeIndexService struct {
	*WebsocketStreams
}

type CompositeIndexEvent struct {
	Event       string          `json:"e"`
	Time        int64           `json:"E"`
	Symbol      string          `json:"s"`
	Price       decimal.Decimal `json:"p"`
	BaseAsset   string          `json:"C"`
	Composition []struct {
		BaseAsset  string          `json:"b"`
		QuoteAsset string          `json:"q"`
		Quantity   decimal.Decimal `json:"w"`
		Percentage decimal.Decimal `json:"W"`
		IndexPrice decimal.Decimal `json:"i"`
	} `json:"c"`
}

// SubscribeCompositeIndex Stream Name: <symbol>@compositeIndex
func (s *WebsocketStreams) SubscribeCompositeIndex(symbol string) *CompositeIndexService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@compositeIndex", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &CompositeIndexService{s}
}

func (e *CompositeIndexService) Do(ctx context.Context) (<-chan *CompositeIndexEvent, <-chan error) {
	messageCh := make(chan *CompositeIndexEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CompositeIndexEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedCompositeIndexService struct {
	*WebsocketStreams
}

type CombinedCompositeIndexEvent struct {
	Stream string           `json:"stream"`
	Data   *ForceOrderEvent `json:"data"`
}

// SubscribeCombinedCompositeIndex Stream Name: <symbol>@compositeIndex
func (s *WebsocketStreams) SubscribeCombinedCompositeIndex(symbols []string) *CombinedCompositeIndexService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@compositeIndex", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedCompositeIndexService{s}
}

func (e *CombinedCompositeIndexService) Do(ctx context.Context) (<-chan *CombinedCompositeIndexEvent, <-chan error) {
	messageCh := make(chan *CombinedCompositeIndexEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedCompositeIndexEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// ContractInfoService ContractInfo stream pushes when contract info updates(listing/settlement/contract bracket update).
// bks field only shows up when bracket gets updated.
type ContractInfoService struct {
	*WebsocketStreams
}
type ContractInfoEvent struct {
	Event          string `json:"e"`
	Time           int64  `json:"E"`
	Symbol         string `json:"s"`
	Pair           string `json:"ps"`
	ContractType   string `json:"ct"`
	DeliveryTime   int64  `json:"dt"`
	OnboardTime    int64  `json:"ot"`
	ContractStatus string `json:"cs"`
	Bks            []struct {
		Bracket            int     `json:"bs"`
		FloorBracket       int     `json:"bnf"`
		CapBracket         int     `json:"bnc"`
		MaintenanceBracket float64 `json:"mmr"`
		Auxiliary          int     `json:"cf"`
		MinLeverage        int     `json:"mi"`
		MaxLeverage        int     `json:"ma"`
	} `json:"bks"`
}

// SubscribeContractInfo Stream Name: !contractInfo
func (s *WebsocketStreams) SubscribeContractInfo() *ContractInfoService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!contractInfo", s.c.getEndpoint()))
	return &ContractInfoService{s}
}

func (e *ContractInfoService) Do(ctx context.Context) (<-chan *ContractInfoEvent, <-chan error) {
	messageCh := make(chan *ContractInfoEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *ContractInfoEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// AssetIndexArrService Asset index for multi-assets mode user
type AssetIndexArrService struct {
	*WebsocketStreams
}
type AssetIndexArr struct {
	Event             string          `json:"e"`
	Time              int64           `json:"E"`
	Symbol            string          `json:"s"`
	IndexPrice        decimal.Decimal `json:"i"`
	BidBuffer         decimal.Decimal `json:"b"`
	AskBuffer         decimal.Decimal `json:"a"`
	BidRate           decimal.Decimal `json:"B"`
	AskRate           decimal.Decimal `json:"A"`
	ExchangeBidBuffer decimal.Decimal `json:"q"`
	ExchangeAskBuffer decimal.Decimal `json:"g"`
	ExchangeBidRate   decimal.Decimal `json:"Q"`
	ExchangeAskRate   decimal.Decimal `json:"G"`
}

// SubscribeAssetIndexArr Stream Name: !assetIndex@arr
func (s *WebsocketStreams) SubscribeAssetIndexArr() *AssetIndexArrService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!assetIndex@arr", s.c.getEndpoint()))
	return &AssetIndexArrService{s}
}

func (e *AssetIndexArrService) Do(ctx context.Context) (<-chan []*AssetIndexArr, <-chan error) {
	messageCh := make(chan []*AssetIndexArr, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event []*AssetIndexArr
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

// AssetIndexService Asset index for multi-assets mode user
type AssetIndexService struct {
	*WebsocketStreams
}

// SubscribeAssetIndex Stream Name: <symbol>@assetIndex
func (s *WebsocketStreams) SubscribeAssetIndex(symbol string) *AssetIndexService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@assetIndex", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &AssetIndexService{s}
}

func (e *AssetIndexService) Do(ctx context.Context) (<-chan *AssetIndexArr, <-chan error) {
	messageCh := make(chan *AssetIndexArr, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *AssetIndexArr
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}

type CombinedAssetIndexService struct {
	*WebsocketStreams
}

type CombinedAssetIndexEvent struct {
	Stream string         `json:"stream"`
	Data   *AssetIndexArr `json:"data"`
}

func (s *WebsocketStreams) SubscribeCombinedAssetIndex(symbols []string) *CombinedAssetIndexService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@assetIndex", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedAssetIndexService{s}
}

func (e *CombinedAssetIndexService) Do(ctx context.Context) (<-chan *CombinedAssetIndexEvent, <-chan error) {
	messageCh := make(chan *CombinedAssetIndexEvent, 8)
	errorCh := make(chan error)

	go func() {
		defer close(messageCh)
		defer close(errorCh)
		onMessage, onError := e.c.wsServe(ctx)
		for {
			select {
			case <-ctx.Done():
				return
			case message := <-onMessage:
				var event *CombinedAssetIndexEvent
				if err := json.Unmarshal(message, &event); err != nil {
					errorCh <- err
					continue
				}
				messageCh <- event
			case err := <-onError:
				errorCh <- err
				return
			}
		}
	}()
	return messageCh, errorCh
}
