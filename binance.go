package binance

import (
	"encoding/json"
	"github.com/jekaxv/go-binance/https"
	"github.com/jekaxv/go-binance/https/hfutures"
	"github.com/jekaxv/go-binance/wss"
	"github.com/jekaxv/go-binance/wss/wfutures"
	"log/slog"
)

func NewClient(opt ...https.Options) *https.Client {
	return &https.Client{
		Opt:        https.NewOptions(opt...),
		HttpClient: https.DefaultClient,
		Logger:     slog.Default(),
	}
}

func NewWsClient(opt ...wss.Options) *wss.Client {
	return &wss.Client{
		Opt:    wss.NewOptions(opt...),
		Logger: slog.Default(),
	}
}

func NewWsApiClient(opt ...wss.Options) *wss.Client {
	return &wss.Client{
		Opt:    wss.NewApiOptions(opt...),
		Logger: slog.Default(),
	}
}

func NewFuturesClient(opt ...https.Options) *hfutures.Client {
	return &hfutures.Client{
		C: &https.Client{
			Opt:        https.NewFuturesOptions(opt...),
			HttpClient: https.DefaultClient,
			Logger:     slog.Default(),
		},
	}
}
func NewFuturesWsApiClient(opt ...wss.Options) *wfutures.Client {
	return &wfutures.Client{
		C: &wss.Client{
			Opt:    wss.NewFuturesApiOptions(opt...),
			Logger: slog.Default(),
		},
	}
}
func NewFuturesWsClient(opt ...wss.Options) *wfutures.Client {
	return &wfutures.Client{
		C: &wss.Client{
			Opt:    wss.NewFuturesWsOptions(opt...),
			Logger: slog.Default(),
		},
	}
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
