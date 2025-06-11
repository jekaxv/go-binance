package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
	"github.com/jekaxv/go-binance/futures"
)

func main() {
	client := binance.NewFuturesClient(core.Options{
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	orders := []futures.OrderReq{{
		Symbol:    "BTCUSDT",
		Side:      "BUY",
		OrderType: "MARKET",
		Quantity:  "10",
	}}
	resp, err := client.NewPlaceBatchOrder().BatchOrders(orders).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
