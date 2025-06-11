package core

import (
	"log/slog"
	"time"
)

const (
	BaseURL    = "https://api.binance.com"
	TestnetURL = "https://testnet.binance.vision"

	FuturesUrl        = "https://fapi.binance.com"
	FuturesTestnetUrl = "https://testnet.binancefuture.com"

	WsBaseURL    = "wss://stream.binance.com:9443"
	WsTestnetURL = "wss://stream.testnet.binance.vision"

	ApiBaseURL    = "wss://ws-api.binance.com:443/ws-api/v3"
	ApiTestnetURL = "wss://ws-api.testnet.binance.vision/ws-api/v3"

	FuturesBaseURL        = "wss://ws-fapi.binance.com/ws-fapi/v1"
	FuturesTestnetBaseURL = "wss://testnet.binancefuture.com/ws-fapi/v1"

	FuturesStreamUrl = "wss://fstream.binance.com"
)

var WebsocketStreamsTimeout = time.Second * 60

type Options struct {
	Endpoint  string
	ApiKey    string
	ApiSecret string
	SignType  SignType

	Logger *slog.Logger
}

func (o *Options) init() {
	if o.Endpoint == "" {
		o.Endpoint = BaseURL
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
}

func (o *Options) initFutures() {
	if o.Endpoint == "" {
		o.Endpoint = FuturesUrl
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
}

func NewOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].init()
	return &opt[0]
}

func NewFuturesOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].initFutures()
	return &opt[0]
}

func (o *Options) wsInit() {
	if o.Endpoint == "" {
		o.Endpoint = WsBaseURL
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
}
func (o *Options) initApi() {
	if o.Endpoint == "" {
		o.Endpoint = ApiBaseURL
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
}

func (o *Options) initFuturesApi() {
	if o.Endpoint == "" {
		o.Endpoint = FuturesBaseURL
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
}
func (o *Options) initFutureStream() {
	if o.Endpoint == "" {
		o.Endpoint = FuturesStreamUrl
	}
	if o.Logger == nil {
		o.Logger = slog.Default()
	}
}

func NewWsOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].wsInit()
	return &opt[0]
}

func NewWsApiOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].initApi()
	return &opt[0]
}

func NewFuturesApiOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].initFuturesApi()
	return &opt[0]
}

func NewFuturesWsOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].initFutureStream()
	return &opt[0]
}
