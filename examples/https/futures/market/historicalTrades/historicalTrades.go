package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/https"
)

func main() {
	client := binance.NewFuturesClient(https.Options{
		ApiKey: "Api Key",
	})
	resp, err := client.NewHistoricalTrades().Symbol("BTCUSDT").Limit(10).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
