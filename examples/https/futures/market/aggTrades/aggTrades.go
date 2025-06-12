package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
	"log/slog"
	"os"
)

func main() {
	opts := &slog.HandlerOptions{Level: slog.LevelDebug}
	client := binance.NewFuturesClient(core.Options{
		Endpoint: core.FuturesTestnetUrl,
		Logger:   slog.New(slog.NewJSONHandler(os.Stdout, opts)),
	})
	resp, err := client.NewAggTrades().Symbol("BTCUSDT").Limit(5).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println(binance.PrettyPrint(resp))
}
