package https

import (
	"github.com/jekaxv/go-binance/types"
	"net/http"
)

const (
	BaseURL    = "https://api.binance.com"
	TestnetURL = "https://testnet.binance.vision"

	FuturesUrl        = "https://fapi.binance.com"
	FuturesTestnetUrl = "https://testnet.binancefuture.com"
)

var DefaultClient = http.DefaultClient

type Options struct {
	ApiKey    string
	ApiSecret string
	Endpoint  string
	SignType  types.SignType
}

func (o *Options) init() {
	if o.Endpoint == "" {
		o.Endpoint = BaseURL
	}
}

func (o *Options) initFutures() {
	if o.Endpoint == "" {
		o.Endpoint = FuturesUrl
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
