package wss

import (
	"github.com/jekaxv/go-binance/types"
	"github.com/jekaxv/go-binance/utils"
)

type request struct {
	Id       string         `json:"id"`
	Method   string         `json:"method"`
	Params   map[string]any `json:"params,omitempty"`
	AuthType types.AuthType `json:"-"`
	SignFunc utils.SignFunc `json:"-"`
}
