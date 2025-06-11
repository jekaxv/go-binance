package spot

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
)

type WebsocketStreams struct {
	c *WsClient
}

// AggTradeService The Aggregate Trade Streams push trade information that is aggregated for a single taker order.
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

// TradeService The Trade Streams push raw trade information; each trade has a unique buyer and seller.
type TradeService struct {
	*WebsocketStreams
}

type TradeEvent struct {
	Event        string          `json:"e"`
	Time         int64           `json:"E"`
	Symbol       string          `json:"s"`
	TradeId      int64           `json:"t"`
	Price        decimal.Decimal `json:"p"`
	Quantity     decimal.Decimal `json:"q"`
	TradeTime    int64           `json:"T"`
	IsBuyerMaker bool            `json:"m"`
	Placeholder  bool            `json:"M"` // Ignore
}

// SubscribeTrade Stream Name: <symbol>@trade
func (s *WebsocketStreams) SubscribeTrade(symbol string) *TradeService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@trade", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &TradeService{s}
}

func (e *TradeService) Do(ctx context.Context) (<-chan *TradeEvent, <-chan error) {
	messageCh := make(chan *TradeEvent, 8)
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
				var event *TradeEvent
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

type CombinedTradeService struct {
	*WebsocketStreams
}

type CombinedTradeEvent struct {
	Stream string      `json:"stream"`
	Data   *TradeEvent `json:"data"`
}

func (s *WebsocketStreams) SubscribeCombinedTrade(symbols []string) *CombinedTradeService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@trade", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedTradeService{s}
}

func (e *CombinedTradeService) Do(ctx context.Context) (<-chan *CombinedTradeEvent, <-chan error) {
	messageCh := make(chan *CombinedTradeEvent, 8)
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
				var event *CombinedTradeEvent
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

// KlineService The Kline/Candlestick Stream push updates to the current klines/candlestick every second in UTC+0 timezone
type KlineService struct {
	*WebsocketStreams
}

type KlineEvent struct {
	Event  string `json:"e"`
	Time   int64  `json:"E"`
	Symbol string `json:"s"`
	Kline  struct {
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
	} `json:"k"`
}

// SubscribeKline Stream Name: <symbol>@kline_<interval>
func (s *WebsocketStreams) SubscribeKline(symbol, interval string) *KlineService {
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

// MiniTickerService 24hr rolling window mini-ticker statistics. These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
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

// TickerService 24hr rolling window ticker statistics for a single symbol.
// These are NOT the statistics of the UTC day, but a 24hr rolling window for the previous 24hrs.
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
	FirstTradePrice        decimal.Decimal `json:"x"`
	LastTradePrice         decimal.Decimal `json:"c"`
	LastQuantity           decimal.Decimal `json:"Q"`
	BestBidPrice           decimal.Decimal `json:"b"`
	BestBidQuantity        decimal.Decimal `json:"B"`
	BestAskPrice           decimal.Decimal `json:"a"`
	BestAskQuantity        decimal.Decimal `json:"A"`
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

// TickerWindowSizeService Rolling window ticker statistics for a single symbol, computed over multiple windows.
type TickerWindowSizeService struct {
	*WebsocketStreams
}

type TickerWindowSizeEvent struct {
	Event                  string          `json:"e"`
	Time                   int64           `json:"E"`
	Symbol                 string          `json:"s"`
	PriceChange            decimal.Decimal `json:"p"`
	PriceChangePercent     decimal.Decimal `json:"P"`
	OpenPrice              decimal.Decimal `json:"o"`
	HighPrice              decimal.Decimal `json:"h"`
	LowPrice               decimal.Decimal `json:"l"`
	LastPrice              decimal.Decimal `json:"c"`
	WeightedAveragePrice   decimal.Decimal `json:"w"`
	TotalTradedBaseVolume  decimal.Decimal `json:"v"`
	TotalTradedQuoteVolume decimal.Decimal `json:"q"`
	OpenTime               int             `json:"O"`
	CloseTime              int             `json:"C"`
	FirstTradeID           int             `json:"F"`
	LastTradeID            int             `json:"L"`
	NumberOfTrades         int             `json:"n"`
}

// SubscribeTickerWindowSize Stream Name: <symbol>@ticker_<window_size>
// windowSize: 1h,4h,1d
func (s *WebsocketStreams) SubscribeTickerWindowSize(symbol, windowSize string) *TickerWindowSizeService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@ticker_%s", s.c.getEndpoint(), strings.ToLower(symbol), windowSize))
	return &TickerWindowSizeService{s}
}

func (e *TickerWindowSizeService) Do(ctx context.Context) (<-chan *TickerWindowSizeEvent, <-chan error) {
	messageCh := make(chan *TickerWindowSizeEvent, 8)
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
				var event *TickerWindowSizeEvent
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

type CombinedTickerWindowSizeService struct {
	*WebsocketStreams
}

type CombinedTickerWindowSizeEvent struct {
	Stream string                 `json:"stream"`
	Data   *TickerWindowSizeEvent `json:"data"`
}

// SubscribeCombinedTickerWindowSize Stream Name: <symbol>@ticker_<window_size>
// windowSize: 1h,4h,1d
func (s *WebsocketStreams) SubscribeCombinedTickerWindowSize(symbols map[string]string) *CombinedTickerWindowSizeService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for symbol, windowSize := range symbols {
		endpoint += fmt.Sprintf("%s@ticker_%s", strings.ToLower(symbol), windowSize) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])

	return &CombinedTickerWindowSizeService{s}
}

func (e *CombinedTickerWindowSizeService) Do(ctx context.Context) (<-chan *CombinedTickerWindowSizeEvent, <-chan error) {
	messageCh := make(chan *CombinedTickerWindowSizeEvent, 8)
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
				var event *CombinedTickerWindowSizeEvent
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

// TickerWindowSizeArrService Rolling window ticker statistics for all market symbols, computed over multiple windows.
// Note that only tickers that have changed will be present in the array.
type TickerWindowSizeArrService struct {
	*WebsocketStreams
}

// SubscribeTickerWindowSizeArr Stream Name: !ticker_<window-size>@arr
func (s *WebsocketStreams) SubscribeTickerWindowSizeArr(windowSize string) *TickerWindowSizeArrService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/!ticker_%s@arr", s.c.getEndpoint(), windowSize))
	return &TickerWindowSizeArrService{s}
}

func (e *TickerWindowSizeArrService) Do(ctx context.Context) (<-chan []*TickerWindowSizeEvent, <-chan error) {
	messageCh := make(chan []*TickerWindowSizeEvent, 8)
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
				var event []*TickerWindowSizeEvent
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
	UpdateId     int             `json:"u"`
	Symbol       string          `json:"s"`
	BestBidPrice decimal.Decimal `json:"b"`
	BestBidQty   decimal.Decimal `json:"B"`
	BestAskPrice decimal.Decimal `json:"a"`
	BestAskQty   decimal.Decimal `json:"A"`
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

// AvgPriceService Average price streams push changes in the average price over a fixed time interval.
type AvgPriceService struct {
	*WebsocketStreams
}

type AvgPriceEvent struct {
	Event        string          `json:"e"`
	Time         int64           `json:"E"`
	Symbol       string          `json:"s"`
	Interval     string          `json:"i"`
	AveragePrice decimal.Decimal `json:"w"`
	TradeTime    int64           `json:"T"`
}

// SubscribeAvgPrice Stream Name: <symbol>@avgPrice
func (s *WebsocketStreams) SubscribeAvgPrice(symbol string) *AvgPriceService {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s@avgPrice", s.c.getEndpoint(), strings.ToLower(symbol)))
	return &AvgPriceService{s}
}

func (e *AvgPriceService) Do(ctx context.Context) (<-chan *AvgPriceEvent, <-chan error) {
	messageCh := make(chan *AvgPriceEvent, 8)
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
				var event *AvgPriceEvent
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

type CombinedAvgPriceService struct {
	*WebsocketStreams
}

type CombinedAvgPriceEvent struct {
	Stream string         `json:"stream"`
	Data   *AvgPriceEvent `json:"data"`
}

// SubscribeCombinedAvgPrice Stream Name: <symbol>@avgPrice
func (s *WebsocketStreams) SubscribeCombinedAvgPrice(symbols []string) *CombinedAvgPriceService {
	s.c.combined(true)
	endpoint := s.c.getEndpoint()
	for _, symbol := range symbols {
		endpoint += fmt.Sprintf("%s@avgPrice", strings.ToLower(symbol)) + "/"
	}
	s.c.setEndpoint(endpoint[:len(endpoint)-1])
	return &CombinedAvgPriceService{s}
}

func (e *CombinedAvgPriceService) Do(ctx context.Context) (<-chan *CombinedAvgPriceEvent, <-chan error) {
	messageCh := make(chan *CombinedAvgPriceEvent, 8)
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
				var event *CombinedAvgPriceEvent
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

type DepthLevelEvent struct {
	LastUpdateId int                 `json:"lastUpdateId"`
	Bids         [][]decimal.Decimal `json:"bids"` // Bids to be updated. [0]Price level to be updated, [1]Quantity
	Asks         [][]decimal.Decimal `json:"asks"` // Asks to be updated. [0]Price level to be updated, [1]Quantity
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

func (e *DepthLevelService) Do(ctx context.Context) (<-chan *DepthLevelEvent, <-chan error) {
	messageCh := make(chan *DepthLevelEvent, 8)
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
				var event *DepthLevelEvent
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
	Stream string           `json:"stream"`
	Data   *DepthLevelEvent `json:"data"`
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

// DepthService Order book price and quantity depth updates used to locally manage an order book.
type DepthService struct {
	*WebsocketStreams
}

type DepthEvent struct {
	Event   string              `json:"e"`
	Time    int64               `json:"E"`
	Symbol  string              `json:"s"`
	FirstId int64               `json:"U"`
	FinalId int64               `json:"u"`
	Bids    [][]decimal.Decimal `json:"b"` // Bids to be updated. [0]Price level to be updated, [1]Quantity
	Asks    [][]decimal.Decimal `json:"a"` // Asks to be updated. [0]Price level to be updated, [1]Quantity
}

// SubscribeDepth Stream Names: <symbol>@depth OR <symbol>@depth@100ms
// Update Speed: 1000ms or 100ms
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
