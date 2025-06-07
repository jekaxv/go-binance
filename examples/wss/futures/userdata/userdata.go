package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/types"
	"github.com/jekaxv/go-binance/wss"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelFunc()
	client := binance.NewFuturesWsClient(wss.Options{
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
		SignType:  types.SignTypeEd25519,
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
