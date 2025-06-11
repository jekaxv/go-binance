package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
)

func main() {
	client := binance.NewFuturesClient()
	resp, err := client.NewFundingRate().Symbol("BTCUSDT").Limit(1).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
