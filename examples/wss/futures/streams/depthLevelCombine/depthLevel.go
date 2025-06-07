package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/wss"
	"time"
)

func main() {
	client := binance.NewFuturesWsClient(wss.Options{
		Endpoint: wss.FuturesStreamUrl,
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	symbols := make(map[string]int)
	symbols["btcusdt"] = 5
	symbols["ETHUSDT"] = 5
	onMessage, onError := client.NewWebsocketStreams().
		SubscribeCombinedDepthLevel(symbols, "100ms").
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
