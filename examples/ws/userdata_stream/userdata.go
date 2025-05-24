package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/ws"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	client := binance.NewWsClient(ws.Options{
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	onMessage, onError := client.NewWebsocketStreams().SubscribeUserData("ListenKey").Do(ctx)
	for {
		select {
		case event := <-onMessage:
			fmt.Printf("Received event: %+v\n", event)
		case err := <-onError:
			fmt.Printf("Error: %v\n", err)
			return
		case <-ctx.Done():
			fmt.Printf("Timeout")
			return
		}
	}
}
