package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
)

// StartUserDataStream Start a new user data stream. The stream will close after 60 minutes unless a keepalive is sent.
type StartUserDataStream struct {
	c *Client
	r *core.Request
}

type StartUserDataStreamResponse struct {
	ListenKey string `json:"listenKey"`
}

func (s *StartUserDataStream) Do(ctx context.Context) (*StartUserDataStreamResponse, error) {
	if err := s.c.invoke(s.r, ctx); err != nil {
		return nil, err
	}
	var resp *StartUserDataStreamResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CloseUserDataStream Close out a user data stream.
type CloseUserDataStream struct {
	c *Client
	r *core.Request
}

func (s *CloseUserDataStream) ListenKey(listenKey string) *CloseUserDataStream {
	s.r.Set("listenKey", listenKey)
	return s
}

func (s *CloseUserDataStream) Do(ctx context.Context) error {
	return s.c.invoke(s.r, ctx)
}

// PingUserDataStream Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
type PingUserDataStream struct {
	c *Client
	r *core.Request
}

func (s *PingUserDataStream) ListenKey(listenKey string) *PingUserDataStream {
	s.r.Set("listenKey", listenKey)
	return s
}

func (s *PingUserDataStream) Do(ctx context.Context) error {
	return s.c.invoke(s.r, ctx)
}
