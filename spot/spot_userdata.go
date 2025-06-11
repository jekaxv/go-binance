package spot

import (
	"context"
	"encoding/json"
)

// StartUserDataStream Start a new user data stream. The stream will close after 60 minutes unless a keepalive is sent.
type StartUserDataStream struct {
	c *Client
}

type StartUserDataStreamResponse struct {
	ListenKey string `json:"listenKey"`
}

func (s *StartUserDataStream) Do(ctx context.Context) (*StartUserDataStreamResponse, error) {
	if err := s.c.invoke(ctx); err != nil {
		return nil, err
	}
	var resp *StartUserDataStreamResponse
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// CloseUserDataStream Close out a user data stream.
type CloseUserDataStream struct {
	c         *Client
	listenKey string
}

func (s *CloseUserDataStream) ListenKey(listenKey string) *CloseUserDataStream {
	s.listenKey = listenKey
	return s
}

func (s *CloseUserDataStream) Do(ctx context.Context) error {
	s.c.set("listenKey", s.listenKey)
	return s.c.invoke(ctx)
}

// PingUserDataStream Keepalive a user data stream to prevent a time out. User data streams will close after 60 minutes. It's recommended to send a ping about every 30 minutes.
type PingUserDataStream struct {
	c         *Client
	listenKey string
}

func (s *PingUserDataStream) ListenKey(listenKey string) *PingUserDataStream {
	s.listenKey = listenKey
	return s
}

func (s *PingUserDataStream) Do(ctx context.Context) error {
	s.c.set("listenKey", s.listenKey)
	return s.c.invoke(ctx)
}
