package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewFuturesClient()
	resp, err := client.NewFutureBasis().
		Symbol("BTCUSDT").
		ContractType(core.ContractTypePERPETUAL).
		Period(core.Interval5m).
		Limit(30).
		Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
