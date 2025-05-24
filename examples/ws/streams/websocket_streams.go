package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/ws"
	"time"
)

func main() {
	client := binance.NewWsClient(ws.Options{
		Endpoint:  ws.TestnetURL,
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	onMessage, onError := client.NewWebsocketStreams().SubscribeAggTrade("btcusdt").Do(ctx)
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
