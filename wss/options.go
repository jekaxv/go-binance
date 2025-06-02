package wss

import "time"

const (
	BaseURL    = "wss://stream.binance.com:9443"
	TestnetURL = "wss://stream.testnet.binance.vision"

	ApiBaseURL    = "wss://ws-api.binance.com:443/ws-api/v3"
	ApiTestnetURL = "wss://testnet.binance.vision/ws-api/v3"

	FuturesBaseURL        = "wss://ws-fapi.binance.com/ws-fapi/v1"
	FuturesTestnetBaseURL = "wss://testnet.binancefuture.com/ws-fapi/v1"
)

var WebsocketStreamsTimeout = time.Second * 60

type Options struct {
	Endpoint  string
	ApiKey    string
	ApiSecret string
}

func (o *Options) init() {
	if o.Endpoint == "" {
		o.Endpoint = BaseURL
	}
}
func (o *Options) initApi() {
	if o.Endpoint == "" {
		o.Endpoint = ApiBaseURL
	}
}

func (o *Options) initFuturesApi() {
	if o.Endpoint == "" {
		o.Endpoint = FuturesBaseURL
	}
}

func NewOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].init()
	return &opt[0]
}

func NewApiOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].initApi()
	return &opt[0]
}

func NewFuturesOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		opt = append(opt, Options{})
	}
	opt[0].initFuturesApi()
	return &opt[0]
}
