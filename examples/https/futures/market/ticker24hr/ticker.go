package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
)

func main() {
	client := binance.NewFuturesClient()
	resp, err := client.NewTicker24hr().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
