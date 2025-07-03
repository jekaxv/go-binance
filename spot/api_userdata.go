package spot

import (
	"context"
	"encoding/json"
	"github.com/jekaxv/go-binance/core"
)

// SessionLogon Authenticate WebSocket connection using the provided API key.
type SessionLogon struct {
	c *WsClient
	r *core.WsRequest
}
type SessionResult struct {
	ApiKey           string `json:"apiKey"`
	AuthorizedSince  int64  `json:"authorizedSince"`
	ConnectedSince   int64  `json:"connectedSince"`
	ReturnRateLimits bool   `json:"returnRateLimits"`
	ServerTime       int64  `json:"serverTime"`
	UserDataStream   bool   `json:"userDataStream"`
}

type SessionResponse struct {
	ApiResponse
	Result *SessionResult `json:"result"`
}

func (s *SessionLogon) Do(ctx context.Context) (*SessionResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *SessionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// SessionStatus Query the status of the WebSocket connection, inspecting which API key (if any) is used to authorize requests.
type SessionStatus struct {
	c *WsClient
	r *core.WsRequest
}

func (s *SessionStatus) Do(ctx context.Context) (*SessionResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *SessionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}

// SessionLogout Forget the API key previously authenticated.
// If the connection is not authenticated, this request does nothing.
type SessionLogout struct {
	c *WsClient
	r *core.WsRequest
}

func (s *SessionLogout) Do(ctx context.Context) (*SessionResponse, error) {
	onMessage, onError := s.c.wsApiServe(ctx)
	if err := s.c.send(s.r); err != nil {
		return nil, err
	}
	defer func(c *WsClient) {
		err := c.close()
		if err != nil {
			s.c.Opt.Logger.Debug("websocket close failed", "error", err)
		}
	}(s.c)
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case message := <-onMessage:
			var resp *SessionResponse
			return resp, json.Unmarshal(message, &resp)
		case err := <-onError:
			return nil, err
		}
	}
}
