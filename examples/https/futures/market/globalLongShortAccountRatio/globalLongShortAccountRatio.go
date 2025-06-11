package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewFuturesClient()
	resp, err := client.NewSymbolRatio().Symbol("BTCUSDT").Period(core.Interval5m).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
