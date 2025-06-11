package futures

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
)

type UserDataStream struct {
	*WebsocketStreams
}

type UserDataEventType string

const (
	listenKeyExpired                 = "listenKeyExpired"
	ACCOUNT_UPDATE                   = "ACCOUNT_UPDATE"
	MARGIN_CALL                      = "MARGIN_CALL"
	ORDER_TRADE_UPDATE               = "ORDER_TRADE_UPDATE"
	TRADE_LITE                       = "TRADE_LITE"
	ACCOUNT_CONFIG_UPDATE            = "ACCOUNT_CONFIG_UPDATE"
	STRATEGY_UPDATE                  = "STRATEGY_UPDATE"
	GRID_UPDATE                      = "GRID_UPDATE"
	CONDITIONAL_ORDER_TRIGGER_REJECT = "CONDITIONAL_ORDER_TRIGGER_REJECT"
)

type UserDataEvent struct {
	Event               UserDataEventType `json:"e"`
	Time                int64             `json:"E"`
	AccountUpdate       AccountUpdate
	ListenExpired       ListenExpired
	MarginCall          MarginCall
	OrderTradeUpdate    OrderTradeUpdate
	TradeLite           TradeLite
	AccountConfigUpdate AccountConfigUpdate
	UpdateStrategy      UpdateStrategy
	GridUpdate          GridUpdate
	OrderTriggerReject  OrderTriggerReject
}

type OrderTriggerReject struct {
	SendTime int64 `json:"T"`
	Or       struct {
		Symbol       string `json:"s"`
		OrderId      int64  `json:"i"`
		RejectReason string `json:"r"`
	} `json:"or"`
}

type GridUpdate struct {
	TransactionTime int64 `json:"T"`
	Gu              struct {
		StrategyID     int             `json:"si"`
		StrategyType   string          `json:"st"`
		StrategyStatus string          `json:"ss"`
		Symbol         string          `json:"s"`
		OpCode         int             `json:"c"`
		RealizedPNL    decimal.Decimal `json:"r"`
		AveragePrice   decimal.Decimal `json:"up"`
		Qty            decimal.Decimal `json:"uq"`
		Fee            decimal.Decimal `json:"uf"`
		MatchedPNL     decimal.Decimal `json:"mp"`
		UpdateT        int64           `json:"ut"`
	} `json:"gu"`
}

type UpdateStrategy struct {
	StrategyID     int    `json:"si"`
	StrategyType   string `json:"st"`
	StrategyStatus string `json:"ss"`
	Symbol         string `json:"s"`
	UpdateTime     int64  `json:"ut"`
	OpCode         int    `json:"c"`
}

type StrategyUpdate struct {
	TransactionTime int64          `json:"T"`
	UpdateStrategy  UpdateStrategy `json:"su"`
}

type AccountConfigUpdate struct {
	TransactionTime int64 `json:"T"`
	Ac              struct {
		Symbol   string `json:"s"`
		Leverage int    `json:"l"`
	} `json:"ac"`
	Ai struct {
		MultiAssetsMode bool `json:"j"`
	} `json:"ai"`
}

type TradeLite struct {
	TransactionTime int64           `json:"T"`
	Symbol          string          `json:"s"`
	Quantity        decimal.Decimal `json:"q"`
	Price           decimal.Decimal `json:"p"`
	MakerSide       bool            `json:"m"`
	ClientOrderId   string          `json:"c"`
	Side            string          `json:"S"`
	LastPrice       decimal.Decimal `json:"L"`
	LastQuantity    decimal.Decimal `json:"l"`
	TradeId         int             `json:"t"`
	OrderId         int             `json:"i"`
}

type UpdateOrder struct {
	Symbol              string          `json:"s"`
	ClientOrderId       string          `json:"c"`
	Side                string          `json:"S"`
	OrderType           string          `json:"o"`
	TimeInForce         string          `json:"f"`
	Quantity            decimal.Decimal `json:"q"`
	Price               decimal.Decimal `json:"p"`
	AveragePrice        decimal.Decimal `json:"ap"`
	StopPrice           decimal.Decimal `json:"sp"`
	ExecutionType       string          `json:"x"`
	OrderStatus         string          `json:"X"`
	OrderId             int             `json:"i"`
	LastQuantity        decimal.Decimal `json:"l"`
	AccumulatedQuantity decimal.Decimal `json:"z"`
	LastPrice           decimal.Decimal `json:"L"`
	CommissionAsset     string          `json:"N"`
	Commission          decimal.Decimal `json:"n"`
	TradeTime           int64           `json:"T"`
	TradeId             int             `json:"t"`
	BidsNotional        decimal.Decimal `json:"b"`
	AskNotional         decimal.Decimal `json:"a"`
	MakerSide           bool            `json:"m"`
	ReduceOnly          bool            `json:"R"`
	WorkingType         string          `json:"wt"`
	OriginalOrderType   string          `json:"ot"`
	PositionSide        string          `json:"ps"`
	ClosePushed         bool            `json:"cp"`
	ActivationPrice     decimal.Decimal `json:"AP"`
	CallbackRate        decimal.Decimal `json:"cr"`
	PriceProtection     bool            `json:"pP"`
	Si                  int             `json:"si"`
	Ss                  int             `json:"ss"`
	RealizedProfit      string          `json:"rp"`
	STPMode             string          `json:"V"`
	PriceMatch          string          `json:"pm"`
	Gtd                 int             `json:"gtd"`
}

// OrderTradeUpdate When new order created, order status changed will push such event. event type is ORDER_TRADE_UPDATE.
type OrderTradeUpdate struct {
	TransactionTime int64       `json:"T"`
	O               UpdateOrder `json:"o"`
}

type MarginCallPosition struct {
	Symbol            string          `json:"s"`
	PositionSide      string          `json:"ps"`
	PositionAmount    decimal.Decimal `json:"pa"`
	MarginType        string          `json:"mt"`
	IsolatedWallet    decimal.Decimal `json:"iw"`
	MarkPrice         decimal.Decimal `json:"mp"`
	UnrealizedPnL     decimal.Decimal `json:"up"`
	MaintenanceMargin decimal.Decimal `json:"mm"`
}

type MarginCall struct {
	CrossWallet string                `json:"cw"`
	Position    []*MarginCallPosition `json:"p"`
}

type ListenExpired struct {
	ListenKey string `json:"listenKey"`
}

type UpdateBalance struct {
	Asset         string          `json:"a"`
	WalletBalance decimal.Decimal `json:"wb"`
	CrossWallet   decimal.Decimal `json:"cw"`
	BalanceChange decimal.Decimal `json:"bc"`
}

type UpdatePosition struct {
	Symbol              string          `json:"s"`
	PositionAmount      decimal.Decimal `json:"pa"`
	EntryPrice          decimal.Decimal `json:"ep"`
	BreakevenPrice      decimal.Decimal `json:"bep"`
	AccumulatedRealized decimal.Decimal `json:"cr"`
	UnrealizedPnL       decimal.Decimal `json:"up"`
	MarginType          string          `json:"mt"`
	IsolatedWallet      decimal.Decimal `json:"iw"`
	PositionSide        string          `json:"ps"`
}

// AccountUpdate https://developers.binance.com/docs/derivatives/usds-margined-futures/user-data-streams/Event-Balance-and-Position-Update
type AccountUpdate struct {
	TransactionTime int64 `json:"T"`
	UpdateData      struct {
		ReasonType     string            `json:"m"`
		Balances       []*UpdateBalance  `json:"B"`
		UpdatePosition []*UpdatePosition `json:"P"`
	} `json:"a"`
}

func (s *WebsocketStreams) SubscribeUserData(listenKey string) *UserDataStream {
	s.c.combined(false)
	s.c.setEndpoint(fmt.Sprintf("%s/%s", s.c.getEndpoint(), listenKey))
	return &UserDataStream{s}
}

func (e *UserDataStream) Do(ctx context.Context) (<-chan *UserDataEvent, <-chan error) {
	messageCh := make(chan *UserDataEvent, 8)
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
				event, err := e.parseUserEvent(message)
				if err != nil {
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

func (e *UserDataStream) parseUserEvent(message []byte) (*UserDataEvent, error) {
	var event *UserDataEvent
	if err := json.Unmarshal(message, &event); err != nil {
		return event, err
	}
	switch event.Event {
	case listenKeyExpired:
		return event, json.Unmarshal(message, &event.ListenExpired)
	case ACCOUNT_UPDATE:
		return event, json.Unmarshal(message, &event.AccountUpdate)
	case MARGIN_CALL:
		return event, json.Unmarshal(message, &event.MarginCall)
	case ORDER_TRADE_UPDATE:
		return event, json.Unmarshal(message, &event.OrderTradeUpdate)
	case TRADE_LITE:
		return event, json.Unmarshal(message, &event.TradeLite)
	case ACCOUNT_CONFIG_UPDATE:
		return event, json.Unmarshal(message, &event.AccountConfigUpdate)
	case STRATEGY_UPDATE:
		return event, json.Unmarshal(message, &event.UpdateStrategy)
	case GRID_UPDATE:
		return event, json.Unmarshal(message, &event.GridUpdate)
	case CONDITIONAL_ORDER_TRIGGER_REJECT:
		return event, json.Unmarshal(message, &event.OrderTriggerReject)
	}
	return event, nil
}

// SessionLogon Authenticate WebSocket connection using the provided API key.
type SessionLogon struct {
	c *WsClient
}
type SessionResult struct {
	ApiKey           string `json:"apiKey"`
	AuthorizedSince  int64  `json:"authorizedSince"`
	ConnectedSince   int64  `json:"connectedSince"`
	ReturnRateLimits bool   `json:"returnRateLimits"`
	ServerTime       int64  `json:"serverTime"`
	UserDataStream   bool   `json:"userDataStream"`
}

type SessionResponse struct {
	ApiResponse
	Result *SessionResult `json:"result"`
}

func (s *SessionLogon) Do(ctx context.Context) (*SessionResponse, error) {
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
			var resp *SessionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// SessionStatus Query the status of the WebSocket connection, inspecting which API key (if any) is used to authorize requests.
type SessionStatus struct {
	c *WsClient
}

func (s *SessionStatus) Do(ctx context.Context) (*SessionResponse, error) {
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
			var resp *SessionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// SessionLogout Forget the API key previously authenticated.
// If the connection is not authenticated, this request does nothing.
type SessionLogout struct {
	c *WsClient
}

func (s *SessionLogout) Do(ctx context.Context) (*SessionResponse, error) {
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
			var resp *SessionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
