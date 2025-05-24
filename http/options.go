package http

import "net/http"

const (
	BaseURL    = "https://api.binance.com"
	TestnetURL = "https://testnet.binance.vision"
)

var DefaultClient = http.DefaultClient

type Options struct {
	ApiKey    string
	ApiSecret string
	Endpoint  string
}

func (o *Options) init() {
	if o.Endpoint == "" {
		o.Endpoint = BaseURL
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
