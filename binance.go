package binance

import (
	"encoding/json"
	"github.com/jekaxv/go-binance/http"
	"github.com/jekaxv/go-binance/ws"
)

type Client struct {
	*http.Client
}

type WebsocketClient struct {
	*ws.Client
}

func NewClient(opt ...http.Options) *Client {
	c := Client{
		Client: &http.Client{
			Opt:        http.NewOptions(opt...),
			HttpClient: http.DefaultClient,
		},
	}
	return &c
}

func NewWsClient(opt ...ws.Options) *WebsocketClient {
	c := WebsocketClient{
		Client: &ws.Client{
			Opt: ws.NewOptions(opt...),
		},
	}
	return &c
}

func NewWsApiClient(opt ...ws.Options) *WebsocketClient {
	c := WebsocketClient{
		Client: &ws.Client{
			Opt: ws.NewApiOptions(opt...),
		},
	}
	return &c
}

func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
