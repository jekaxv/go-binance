package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
)

func main() {
	client := binance.NewClient()
	resp, err := client.NewDepth().Symbol("BTCUSDT").Limit(10).Do(context.Background())
	if err != nil {
		panic(err)
	}
	mul := resp.Asks[0][0].Mul(resp.Asks[0][1])
	fmt.Println("Ask Price:", mul)
	fmt.Println(binance.PrettyPrint(resp))
}
