package main

import (
	"context"
	"github.com/jekaxv/go-binance"
)

func main() {
	client := binance.NewFuturesClient()
	err := client.NewPing().Do(context.Background())
	if err != nil {
		panic(err)
	}
}
