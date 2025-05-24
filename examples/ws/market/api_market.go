package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/ws"
)

func main() {
	client := binance.NewWsApiClient(ws.Options{
		Endpoint:  ws.ApiTestnetURL,
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	resp, err := client.NewDepth().Symbol("BTCUSDT").Limit(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
