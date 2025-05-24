package ws

type authType int

const (
	authNone authType = iota
	authApiKey
	authSigned
)

type request struct {
	Id       string         `json:"id"`
	Method   string         `json:"method"`
	AuthType authType       `json:"-"`
	Params   map[string]any `json:"params,omitempty"`
}
