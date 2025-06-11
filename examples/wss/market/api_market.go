package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewWsApiClient(core.Options{
		Endpoint:  core.ApiTestnetURL,
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	resp, err := client.NewDepth().Symbol("BTCUSDT").Limit(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
