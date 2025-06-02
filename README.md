## go-binance

A simple Go library for interacting with the Binance API.

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

### New Client

Initialize the client with your API key and secret. The endpoint is optional, default is "https://api.binance.com".

```go
client := binance.NewClient(http.Options{
    ApiKey:    "YourApiKey",
    ApiSecret: "YourApiSecret",
})
```
### Create Order

```go
package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/https"
	"github.com/jekaxv/go-binance/types"
)

func main() {
	client := binance.NewClient(https.Options{
		Endpoint:  https.TestnetURL,
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	resp, err := client.NewCreateOrder().Symbol("BTCUSDT").
		Side(types.OrderSideBUY).Type(types.OrderTypeMARKET).Quantity("0.001").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(binance.PrettyPrint(resp))
}
```

## Websocket
### New Client
Initialize the client with your API key and secret. The endpoint is optional, default is "wss://stream.binance.com:9443".

```go
package main

import (
	"context"
	"fmt"
	"github.com/jekaxv/go-binance"
	"github.com/jekaxv/go-binance/wss"
)

func main() {
	client := binance.NewWsApiClient(wss.Options{
		ApiKey:    "YourApiKey",
		ApiSecret: "YourApiSecret",
	})
	resp, err := client.NewDepth().Symbol("BTCUSDT").Do(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(binance.PrettyPrint(resp))
}
```