package binance

import (
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
	"github.com/jekaxv/go-binance/futures"
	"github.com/jekaxv/go-binance/spot"
	"net/http"
)

func NewClient(opt ...core.Options) *spot.Client {
	return &spot.Client{
		Client: &core.Client{
			Opt:        core.NewOptions(opt...),
			HttpClient: http.DefaultClient,
		},
	}
}

func NewWsClient(opt ...core.Options) *spot.WsClient {
	return &spot.WsClient{
		WsClient: &core.WsClient{
			Opt: core.NewOptions(opt...),
		},
	}
}

func NewWsApiClient(opt ...core.Options) *spot.WsClient {
	return &spot.WsClient{
		WsClient: &core.WsClient{
			Opt: core.NewWsApiOptions(opt...),
		},
	}
}

func NewFuturesClient(opt ...core.Options) *futures.Client {
	return &futures.Client{
		Client: &core.Client{
			Opt:        core.NewFuturesOptions(opt...),
			HttpClient: http.DefaultClient,
		},
	}
}
func NewFuturesWsApiClient(opt ...core.Options) *futures.WsClient {
	return &futures.WsClient{
		WsClient: &core.WsClient{
			Opt: core.NewFuturesApiOptions(opt...),
		},
	}
}
func NewFuturesWsClient(opt ...core.Options) *futures.WsClient {
	return &futures.WsClient{
		WsClient: &core.WsClient{
			Opt: core.NewFuturesWsOptions(opt...),
		},
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
