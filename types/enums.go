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
	PositionSide_BOTH  PositionSideEnum = "BOTH"
	PositionSide_LONG                   = "LONG"
	PositionSide_SHORT                  = "SHORT"
)

type WorkingType string

const (
	WorkingTypeCONTRACT_PRICE WorkingType = "CONTRACT_PRICE"
	WorkingTypeMARK_PRICE     WorkingType = "MARK_PRICE"
)

type AutoCloseType string

const (
	AutoCloaseTypeLIQUIDATION AutoCloseType = "LIQUIDATION"
	AutoCloaseTypeADL                       = "ADL"
)

type MarginType string

const (
	MarginTypeISOLATED MarginType = "SOLATED"
	MarginTypeICROSSED            = "CROSSED"
)

type IncomeType string

const (
	IncomeType_TRANSFER                    IncomeType = "TRANSFER"
	IncomeType_WELCOME_BONUS                          = "WELCOME_BONUS"
	IncomeType_REALIZED_PNL                           = "REALIZED_PNL"
	IncomeType_FUNDING_FEE                            = "FUNDING_FEE"
	IncomeType_COMMISSION                             = "COMMISSION"
	IncomeType_INSURANCE_CLEAR                        = "INSURANCE_CLEAR"
	IncomeType_REFERRAL_KICKBACK                      = "REFERRAL_KICKBACK"
	IncomeType_COMMISSION_REBATE                      = "COMMISSION_REBATE"
	IncomeType_API_REBATE                             = "API_REBATE"
	IncomeType_CONTEST_REWARD                         = "CONTEST_REWARD"
	IncomeType_CROSS_COLLATERAL_TRANSFER              = "CROSS_COLLATERAL_TRANSFER"
	IncomeType_OPTIONS_PREMIUM_FEE                    = "OPTIONS_PREMIUM_FEE"
	IncomeType_OPTIONS_SETTLE_PROFIT                  = "OPTIONS_SETTLE_PROFIT"
	IncomeType_INTERNAL_TRANSFER                      = "INTERNAL_TRANSFER"
	IncomeType_AUTO_EXCHANGE                          = "AUTO_EXCHANGE"
	IncomeType_DELIVERED_SETTELMENT                   = "DELIVERED_SETTELMENT"
	IncomeType_COIN_SWAP_DEPOSIT                      = "COIN_SWAP_DEPOSIT"
	IncomeType_COIN_SWAP_WITHDRAW                     = "COIN_SWAP_WITHDRAW"
	IncomeType_POSITION_LIMIT_INCREASE_FEE            = "POSITION_LIMIT_INCREASE_FEE"
	IncomeType_STRATEGY_UMFUTURES_TRANSFER            = "STRATEGY_UMFUTURES_TRANSFER"
	IncomeType_FEE_RETURN                             = "FEE_RETURN"
	IncomeType_BFUSD_REWARD                           = "BFUSD_REWARD"
)
