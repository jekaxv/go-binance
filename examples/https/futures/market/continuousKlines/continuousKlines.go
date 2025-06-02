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
		Endpoint: https.FuturesTestnetUrl,
	})
	resp, err := client.NewContractKline().
		Pair("BTCUSDT").
		ContractType(types.ContractTypePERPETUAL).
		Interval(types.Interval1m).
		Limit(5).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
