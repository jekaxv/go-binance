package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
)

func main() {
	client := binance.NewFuturesWsApiClient()
	resp, err := client.NewTickerBook().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
