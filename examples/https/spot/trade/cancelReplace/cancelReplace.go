package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewClient(core.Options{
		Endpoint:  core.TestnetURL,
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	resp, err := client.NewCancelReplace().Symbol("BTCUSDT").
		CancelReplaceMode(core.ReplaceModeSTOP_ON_FAILURE).
		Side(core.OrderSideBUY).Type(core.OrderTypeMARKET).Quantity("0.001").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
