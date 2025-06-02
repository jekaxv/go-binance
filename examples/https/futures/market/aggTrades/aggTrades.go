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
	resp, err := client.NewAggTrades().Symbol("BTCUSDT").Limit(5).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
