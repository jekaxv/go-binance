## go-binance

### Overview

This Go package provides a simple, typed client for interacting with the Binance REST and WebSocket APIs. It supports:

- Spot trading: Market data, account information, and trade endpoints. 
- Futures trading (WebSocket): Real-time data streams via WebSocket for futures markets.

The package wraps the core HTTP and WebSocket clients and exposes domain-specific APIs under spot and futures namespaces.

## Installation

```shell
go get github.com/jekaxv/go-binance
```

## Import

```go
import (
    "github.com/jekaxv/go-binance"
)
```

## Quickstart

### Creating a Client

Initialize the client with your API key and secret. The endpoint is optional, default is "https://api.binance.com".

```go
client := binance.NewClient(core.Options{
    ApiKey:    "YOUR_API_KEY",
    ApiSecret: "YOUR_API_SECRET",
})
```
### Create Order

```go
package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewClient(core.Options{
		Endpoint:  core.TestnetURL,
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	resp, err := client.NewCreateOrder().Symbol("BTCUSDT").
		Side(core.OrderSideBUY).Type(core.OrderTypeMARKET).Quantity("0.001").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
```

## Websocket
### Creating a WebSocket Client
Initialize the client with your API key and secret. The endpoint is optional, default is "wss://stream.binance.com:9443".

```go
package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/core"
)

func main() {
	client := binance.NewWsApiClient(core.Options{
		ApiKey:    "YOUR_API_KEY",
		ApiSecret: "YOUR_API_SECRET",
	})
	resp, err := client.NewDepth().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
```

More examples can be found in [examples](https://github.com/jekaxv/go-binance/tree/main/examples)