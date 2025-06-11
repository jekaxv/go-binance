package futures

import (
	"context"
	"encoding/json"
)

// GetListenKey Start a new user data stream. The stream will close after 60 minutes unless a keepalive is sent.
// If the account has an active listenKey, that listenKey will be returned and its validity will be extended for 60 minutes.
type GetListenKey struct {
	c *Client
}

type ListenKeyResponse struct {
	ListenKey string `json:"listenKey"`
}

func (s *GetListenKey) Do(ctx context.Context) (*ListenKeyResponse, error) {
	var resp *ListenKeyResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

// KeepaliveListenKey Keepalive a user data stream to prevent a time out.
// User data streams will close after 60 minutes. It's recommended to send a ping about every 60 minutes.
type KeepaliveListenKey struct {
	c *Client
}

func (s *KeepaliveListenKey) Do(ctx context.Context) (*ListenKeyResponse, error) {
	var resp *ListenKeyResponse
	if err := s.c.invoke(ctx); err != nil {
		return resp, err
	}
	return resp, json.Unmarshal(s.c.rawBody(), &resp)
}

type CloseListenKey struct {
	c *Client
}

func (s *CloseListenKey) Do(ctx context.Context) error {
	if err := s.c.invoke(ctx); err != nil {
		return err
	}
	return nil
}
