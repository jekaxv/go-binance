package wss

import (
	"context"
	"encoding/json"
	"fmt"
)

type UserDataStream struct {
	*WebsocketStreams
}

type UserDataEventType string

const (
	outboundAccountPosition UserDataEventType = "outboundAccountPosition"
	balanceUpdate                             = "balanceUpdate"
	executionReport                           = "executionReport"
	listStatus                                = "listStatus"
	listenKeyExpired                          = "listenKeyExpired"
	eventStreamTerminated                     = "eventStreamTerminated"
	externalLockUpdate                        = "externalLockUpdate"
)

type UserDataEvent struct {
	Event              UserDataEventType `json:"e"`
	Time               int64             `json:"E"`
	AccountUpdate      AccountUpdate
	BalanceUpdate      BalanceUpdate
	OrderUpdate        OrderUpdate
	ListStatus         ListStatus
	ListenExpired      ListenExpired
	StreamTerminated   StreamTerminated
	ExternalLockUpdate ExternalLockUpdate
}

// AccountUpdate outboundAccountPosition is sent any time an account balance has changed and contains the assets that were possibly changed by the event that generated the balance change.
type AccountUpdate struct {
	AccountUpdateTime int64            `json:"u"`
	BalancesArray     []AccountBalance `json:"B"`
}

type AccountBalance struct {
	Asset  string `json:"a"`
	Free   string `json:"f"`
	Locked string `json:"l"`
}

// BalanceUpdate Balance Update occurs during the following:
// Deposits or withdrawals from the account
// Transfer of funds between accounts (e.g. Spot to Margin)
type BalanceUpdate struct {
	Asset        string `json:"a"`
	BalanceDelta string `json:"d"`
	ClearTime    int64  `json:"T"`
}

// OrderUpdate Orders are updated with the executionReport event.
// We recommend using the FIX API for better performance compared to using the User Data Streams.
type OrderUpdate struct {
	Symbol                  string `json:"s"` // Symbol
	ClientOrderId           string `json:"c"` // Client order ID
	Side                    string `json:"S"` // Side
	OrderType               string `json:"o"` // Order type
	TimeForce               string `json:"f"` // Time in force
	OrderQuantity           string `json:"q"` // Order quantity
	OrderPrice              string `json:"p"` // Order price
	StopPrice               string `json:"P"` // Stop price
	IcebergQuantity         string `json:"F"` // Iceberg quantity
	OrderListId             int    `json:"g"` // OrderListId
	OriginalOrderId         string `json:"C"` // Original client order ID; This is the ID of the order being canceled
	CurrentExecType         string `json:"x"` // Current execution type
	CurrentOrderStatus      string `json:"X"` // Current order status
	OrderRejectReason       string `json:"r"` // Order reject reason; will be an error code.
	OrderId                 int    `json:"i"` // Order ID
	LastExecQuantity        string `json:"l"` // Last executed quantity
	CumulativeQuantity      string `json:"z"` // Cumulative filled quantity
	LastExecPrice           string `json:"L"` // Last executed price
	CommissionAmount        string `json:"n"` // Commission amount
	CommissionAsset         string `json:"N"` // Commission asset
	TransactionTime         int64  `json:"T"` // Transaction time
	TradeId                 int    `json:"t"` // Trade ID
	PreventedMatchId        int    `json:"v"` // Prevented Match Id; This is only visible if the order expired due to STP
	ExecutionId             int    `json:"I"` // Execution Id
	IsInOrderBook           bool   `json:"w"` // Is the order on the book?
	IsMaker                 bool   `json:"m"` // Is this trade the maker side?
	Ignore                  bool   `json:"M"` // Ignore
	CreateTime              int64  `json:"O"` // Order creation time
	FilledQuoteVolume       string `json:"Z"` // Cumulative quote asset transacted quantity
	LatestQuoteVolume       string `json:"Y"` // Last quote asset transacted quantity (i.e. lastPrice * lastQty)
	QuoteVolume             string `json:"Q"` // Quote Order Quantity
	WorkingTime             int64  `json:"W"` // Working Time; This is only visible if the order has been placed on the book.
	SelfTradePreventionMode string `json:"V"` // SelfTradePreventionMode
}

type ListStatus struct {
	Symbol           string          `json:"s"` // Symbol
	OrderListId      int             `json:"g"` // OrderListId
	ContingencyType  string          `json:"c"` // Contingency Type
	ListStatusType   string          `json:"l"` // List Status Type
	ListOrderStatus  string          `json:"L"` // List Order Status
	ListRejectReason string          `json:"r"` // List Reject Reason
	ClientOrderId    string          `json:"C"` // List Client Order ID
	TransactionTime  int64           `json:"T"` // Transaction Time
	ListStatusObj    []ListStatusObj `json:"O"`
}

type ListStatusObj struct {
	Symbol        string `json:"s"` // Symbol
	OrderId       int    `json:"i"` // OrderId
	ClientOrderId string `json:"c"` // ClientOrderId
}

// ListenExpired Listen Key Expired
// This event is sent when the listen key expires.
// No more events will be sent after this until a new listenKey is created.
// This event will not be pushed when the stream is closed normally.
type ListenExpired struct {
	ListenKey string `json:"listenKey"`
}

// StreamTerminated This event appears only for WebSocket API.
// eventStreamTerminated is sent when the User Data Stream is stopped.
// For example, after you send a userDataStream.stop request, or a session.logout request.
type StreamTerminated struct {
	Event struct {
		E  string `json:"e"`
		E1 int64  `json:"E"`
	} `json:"event"`
}

// ExternalLockUpdate is sent when part of your spot wallet balance is locked/unlocked by an external system, for example when used as margin collateral.
type ExternalLockUpdate struct {
	Asset           string `json:"a"` // Asset
	Delta           string `json:"d"` // Delta
	TransactionTime int64  `json:"T"` // Transaction Time
}

func (s *WebsocketStreams) SubscribeUserData(listenKey string) *UserDataStream {
	s.c.combined(false)
	s.c.Opt.Endpoint = fmt.Sprintf("%s/%s", s.c.Opt.Endpoint, listenKey)
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
	case outboundAccountPosition:
		return event, json.Unmarshal(message, &event.AccountUpdate)
	case balanceUpdate:
		return event, json.Unmarshal(message, &event.BalanceUpdate)
	case executionReport:
		return event, json.Unmarshal(message, &event.OrderUpdate)
	case listStatus:
		return event, json.Unmarshal(message, &event.ListStatus)
	case listenKeyExpired:
		return event, json.Unmarshal(message, &event.ListenExpired)
	case eventStreamTerminated:
		return event, json.Unmarshal(message, &event.StreamTerminated)
	case externalLockUpdate:
		return event, json.Unmarshal(message, &event.ExternalLockUpdate)
	}
	return event, nil
}
