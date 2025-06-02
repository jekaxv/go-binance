package types

type PermissionEnum string

const (
	PermissionSPOT      PermissionEnum = "SPOT"
	PermissionMARGIN                   = "MARGIN"
	PermissionLEVERAGED                = "LEVERAGED"
)

type SymbolStatusEnum string

const (
	SymbolStatusPRE_TRADING  SymbolStatusEnum = "PRE_TRADING"
	SymbolStatusRADING                        = "TRADING"
	SymbolStatusOST_TRADING                   = "POST_TRADING"
	SymbolStatusND_OF_DAY                     = "END_OF_DAY"
	SymbolStatusALT                           = "HALT"
	SymbolStatusUCTION_MATCH                  = "AUCTION_MATCH"
	SymbolStatusREAK                          = "BREAK"
)

type IntervalEnum string

const (
	Interval1s  IntervalEnum = "1s"
	Interval1m               = "1m"
	Interval3m               = "3m"
	Interval5m               = "5m"
	Interval15m              = "15m"
	Interval30m              = "30m"
	Interval1h               = "1h"
	Interval2h               = "2h"
	Interval4h               = "4h"
	Interval6h               = "6h"
	Interval8h               = "8h"
	Interval12h              = "12h"
	Interval1d               = "1d"
	Interval3d               = "3d"
	Interval1w               = "1w"
	Interval1M               = "1M"
)

type TickerTypeEnum string

const (
	TickerTypeFULL TickerTypeEnum = "FULL"
	TickerTypeMINI                = "MINI"
)

type OrderSideEnum string

const (
	OrderSideBUY  OrderSideEnum = "BUY"
	OrderSideSELL               = "SELL"
)

type OrderTypeEnum string

const (
	OrderTypeLIMIT             OrderTypeEnum = "LIMIT"
	OrderTypeMARKET                          = "MARKET"
	OrderTypeSTOP_LOSS                       = "STOP_LOSS"
	OrderTypeSTOP_LOSS_LIMIT                 = "STOP_LOSS_LIMIT"
	OrderTypeTAKE_PROFIT                     = "TAKE_PROFIT"
	OrderTypeTAKE_PROFIT_LIMIT               = "TAKE_PROFIT_LIMIT"
	OrderTypeLIMIT_MAKER                     = "LIMIT_MAKER"
)

type TimeInForceEnum string

// GTC: Good Til Canceled An order will be on the book unless the order is canceled.
// IOC: Immediate Or Cancel An order will try to fill the order as much as it can before the order expires.
// FOK: Fill or Kill An order will expire if the full order cannot be filled upon execution.
const (
	TimeInForceGTC TimeInForceEnum = "GTC"
	TimeInForceIOC                 = "IOC"
	TimeInForceFOK                 = "FOK"
)

type OrderResponseTypeEnum string

const (
	OrderResponseTypeACK    OrderResponseTypeEnum = "ACK"
	OrderResponseTypeRESULT                       = "RESULT"
	OrderResponseTypeFULL                         = "FULL"
)

type STPModeEnum string

const (
	STPModeNONE         STPModeEnum = "NONE"
	STPModeEXPIRE_MAKER             = "EXPIRE_MAKER"
	STPModeEXPIRE_TAKER             = "EXPIRE_TAKER"
	STPModeEXPIRE_BOTH              = "EXPIRE_BOTH"
)

type CancelRestrictionEnum string

// ONLY_NEW - Cancel will succeed if the order status is NEW.
// ONLY_PARTIALLY_FILLED - Cancel will succeed if order status is PARTIALLY_FILLED
const (
	CancelRestrictionONLY_NEW              CancelRestrictionEnum = "ONLY_NEW"
	CancelRestrictionONLY_PARTIALLY_FILLED                       = "ONLY_PARTIALLY_FILLED"
)

type CancelReplaceModeEnum string

const (
	ReplaceModeSTOP_ON_FAILURE CancelReplaceModeEnum = "STOP_ON_FAILURE"
	ReplaceModeALLOW_FAILURE                         = "ALLOW_FAILURE"
)

type OrderExceededModeEnum string

// DO_NOTHING (default)- will only attempt to cancel the order if account has not exceeded the unfilled order rate limit
// CANCEL_ONLY - will always cancel the order
const (
	OrderExceededModeDO_NOTHING  OrderExceededModeEnum = "DO_NOTHING"
	OrderExceededModeCANCEL_ONLY                       = "CANCEL_ONLY"
)

type ContractType string

const (
	ContractTypePERPETUAL       ContractType = "PERPETUAL"
	ContractTypeCURRENT_QUARTER              = "CURRENT_QUARTER"
	ContractTypeNEXT_QUARTER                 = "NEXT_QUARTER"
)

type PositionSideEnum string

const (
	PositionSideBOTH  PositionSideEnum = "BOTH"
	PositionSideLONG                   = "LONG"
	PositionSideSHORT                  = "SHORT"
)

type WorkingType string

const (
	WorkingTypeCONTRACT_PRICE WorkingType = "CONTRACT_PRICE"
	WorkingTypeMARK_PRICE     WorkingType = "MARK_PRICE"
)
