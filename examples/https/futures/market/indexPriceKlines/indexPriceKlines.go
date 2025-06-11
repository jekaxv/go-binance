package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewFuturesClient()
	resp, err := client.NewIndexKline().
		Pair("BTCUSDT").
		Interval(core.Interval1m).
		Limit(5).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
