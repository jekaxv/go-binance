package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"time"
)

func main() {
	client := binance.NewFuturesWsClient()
	ctx, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	onMessage, onError := client.NewWebsocketStreams().
		SubscribeCompositeIndex("DEFIUSDT").
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
