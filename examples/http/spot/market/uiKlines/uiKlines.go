package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/types"
)

func main() {
	client := binance.NewClient()
	resp, err := client.NewUIKlines().Symbol("BTCUSDT").Interval(types.Interval1m).Limit(10).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
