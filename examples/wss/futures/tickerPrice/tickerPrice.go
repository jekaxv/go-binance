package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/wss"
)

func main() {
	client := binance.NewWsFuturesClient(wss.Options{
		Endpoint: wss.FuturesTestnetBaseURL,
	})
	resp, err := client.NewTickerPrice().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
