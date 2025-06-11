package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
	"time"
)

func main() {
	client := binance.NewFuturesWsClient()
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	onMessage, onError := client.NewWebsocketStreams().
		SubscribeContractKline("BTCUSDT", core.ContractTypePERPETUAL, core.Interval1m).
		Do(ctx)
	for {
		select {
		case event := <-onMessage:
			fmt.Println(binance.PrettyPrint(event))
		case err := <-onError:
			fmt.Println(err)
			return
		case <-ctx.Done():
			fmt.Println("Timeout")
			return
		}
	}
}
