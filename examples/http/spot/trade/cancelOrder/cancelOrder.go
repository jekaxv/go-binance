package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/http"
)

func main() {
	client := binance.NewClient(http.Options{
		Endpoint:  http.TestnetURL,
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	resp, err := client.NewCancelOrder().Symbol("BTCUSDT").OrderId(4).Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
