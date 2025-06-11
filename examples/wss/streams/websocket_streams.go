package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
	"time"
)

func main() {
	client := binance.NewWsClient(core.Options{
		Endpoint:  core.TestnetURL,
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	onMessage, onError := client.NewWebsocketStreams().SubscribeDepth("btcusdt").Do(ctx)
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
