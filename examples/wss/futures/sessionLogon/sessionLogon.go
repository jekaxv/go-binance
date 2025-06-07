package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/types"
	"github.com/jekaxv/go-binance/wss"
)

func main() {
	client := binance.NewFuturesWsApiClient(wss.Options{
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
		SignType:  types.SignTypeEd25519,
	})
	resp, err := client.NewSessionLogon().Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
