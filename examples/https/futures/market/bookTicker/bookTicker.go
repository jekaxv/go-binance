package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewFuturesClient(core.Options{
		Endpoint: core.FuturesTestnetUrl,
	})
	resp, err := client.NewBookTicker().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
