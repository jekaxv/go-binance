package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/https"
	"github.com/jekaxv/go-binance/types"
)

func main() {
	client := binance.NewFuturesClient(https.Options{
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: `YOUR_API_SECRET`,
		SignType:  types.SignTypeEd25519,
	})
	resp, err := client.NewGetListenKey().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
