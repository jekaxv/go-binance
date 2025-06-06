package wss

import (
	"context"
	"encoding/json"
)

// StartUserDataStream Start a new user data stream.
type StartUserDataStream struct {
	c *Client
}

type StartUserDataStreamResponse struct {
	ApiResponse
	Result struct {
		ListenKey string `json:"listenKey"`
	} `json:"result"`
}

func (s *StartUserDataStream) Do(ctx context.Context) (*StartUserDataStreamResponse, error) {
	s.c.combined(false)
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *StartUserDataStreamResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// PingUserDataStream Ping a user data stream to keep it alive.
type PingUserDataStream struct {
	c *Client
}

func (s *PingUserDataStream) ListenKey(listenKey string) *PingUserDataStream {
	s.c.req.Params["listenKey"] = listenKey
	return s
}

func (s *PingUserDataStream) Do(ctx context.Context) (*ApiResponse, error) {
	s.c.combined(false)
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *ApiResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// StopUserDataStream Explicitly stop and close the user data stream.
type StopUserDataStream struct {
	c *Client
}

func (s *StopUserDataStream) ListenKey(listenKey string) *StopUserDataStream {
	s.c.req.Params["listenKey"] = listenKey
	return s
}

func (s *StopUserDataStream) Do(ctx context.Context) (*ApiResponse, error) {
	s.c.combined(false)
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(); err != nil {
		return nil, err
	}
	defer s.c.close()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *ApiResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
