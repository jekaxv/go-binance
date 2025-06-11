package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewFuturesClient(core.Options{
		ApiKey: "YOUR_API_KEY",
	})
	resp, err := client.NewHistoricalTrades().Symbol("BTCUSDT").Limit(10).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
