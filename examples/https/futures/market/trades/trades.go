package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/https"
)

func main() {
	client := binance.NewFuturesClient(https.Options{
		Endpoint: https.FuturesTestnetUrl,
	})
	resp, err := client.NewTrades().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
