package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	client := binance.NewFuturesWsApiClient(core.Options{
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	resp, err := client.NewAccountInfo().Do(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
