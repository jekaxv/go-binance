package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/http"
	"github.com/jekaxv/go-binance/types"
)

func main() {
	client := binance.NewClient(http.Options{
		Endpoint:  http.TestnetURL,
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	resp, err := client.NewCreateOrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).Type(types.OrderTypeMARKET).Quantity("0.001").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
