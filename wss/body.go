package wss

type AuthType int

const (
	AuthNone AuthType = iota
	AuthApiKey
	AuthSigned
)

type request struct {
	Id       string         `json:"id"`
	Method   string         `json:"method"`
	AuthType AuthType       `json:"-"`
	Params   map[string]any `json:"params,omitempty"`
}
