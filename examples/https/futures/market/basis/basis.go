package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/https"
	"github.com/jekaxv/go-binance/types"
)

func main() {
	client := binance.NewFuturesClient(https.Options{})
	resp, err := client.NewFutureBasis().
		Symbol("BTCUSDT").
		ContractType(types.ContractTypePERPETUAL).
		Period(types.Interval5m).
		Limit(30).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
