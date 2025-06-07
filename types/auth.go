package types

type AuthType int

const (
	AuthNone AuthType = iota
	AuthApiKey
	AuthSigned
)

type SignType int

const (
	SignTypeHmac SignType = iota
	SignTypeRsa
	SignTypeEd25519
)
