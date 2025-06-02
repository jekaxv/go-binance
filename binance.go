package binance

import (
	"encoding/json"
	"github.com/jekaxv/go-binance/https"
	"github.com/jekaxv/go-binance/https/hfutures"
	"github.com/jekaxv/go-binance/wss"
	"github.com/jekaxv/go-binance/wss/wfutures"
)

func NewClient(opt ...https.Options) *https.Client {
	return &https.Client{
		Opt:        https.NewOptions(opt...),
		HttpClient: https.DefaultClient,
	}
}

func NewWsClient(opt ...wss.Options) *wss.Client {
	return &wss.Client{
		Opt: wss.NewOptions(opt...),
	}
}

func NewWsApiClient(opt ...wss.Options) *wss.Client {
	return &wss.Client{
		Opt: wss.NewApiOptions(opt...),
	}
}

func NewFuturesClient(opt ...https.Options) *hfutures.Client {
	return &hfutures.Client{
		C: &https.Client{
			Opt:        https.NewFuturesOptions(opt...),
			HttpClient: https.DefaultClient,
		},
	}
}
func NewWsFuturesClient(opt ...wss.Options) *wfutures.Client {
	return &wfutures.Client{
		C: &wss.Client{
			Opt: wss.NewFuturesOptions(opt...),
		},
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
