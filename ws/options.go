package ws

import "time"

const (
	BaseURL    = "wss://stream.binance.com:9443"
	TestnetURL = "wss://stream.testnet.binance.vision"

	ApiBaseURL    = "wss://ws-api.binance.com:443/ws-api/v3"
	ApiTestnetURL = "wss://testnet.binance.vision/ws-api/v3"
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

func NewOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		o := Options{}
		o.init()
		opt = append(opt, o)
	}
	return &opt[0]
}

func NewApiOptions(opt ...Options) *Options {
	if len(opt) == 0 {
		o := Options{}
		o.initApi()
		opt = append(opt, o)
	}
	return &opt[0]
}
