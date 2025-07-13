package core

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	Opt        *Options
	HttpClient *http.Client
	fullUrl    string
	resp       *response
}

func (c *Client) SetReq(path, method string, aType ...AuthType) *Request {
	reqType := AuthNone
	if len(aType) > 0 {
		reqType = aType[0]
	}
	return &Request{method: method, path: path, authType: reqType}
}

func (c *Client) parseRequest(r *Request) error {
	if r.authType == AuthSigned {
		r.Set("timestamp", time.Now().UnixMilli())
	}
	fullUrl := fmt.Sprintf("%s%s", c.Opt.Endpoint, r.path)
	query := r.query.Encode()
	form := r.form.Encode()
	header := http.Header{}
	if r.header != nil {
		header = r.header.Clone()
	}
	if r.authType == AuthApiKey || r.authType == AuthSigned {
		header.Set("X-MBX-APIKEY", c.Opt.ApiKey)
	}
	if form != "" {
		header.Set("Content-Type", "application/x-www-form-urlencoded")
		r.body = bytes.NewBufferString(form)
	}
	if r.authType == AuthSigned {
		params := fmt.Sprintf("%s%s", query, form)
		var sf SignFunc
		if c.Opt.SignType == SignTypeRsa {
			sf = RsaSign
		} else if c.Opt.SignType == SignTypeEd25519 {
			sf = Ed25519Sign
		} else {
			sf = HmacSign
		}
		sign, err := sf(c.Opt.ApiSecret, params)
		if err != nil {
			return err
		}
		sign = fmt.Sprintf("signature=%s", sign)
		if query == "" {
			query = sign
		} else {
			query = fmt.Sprintf("%s&%s", query, sign)
		}
	}
	if query != "" {
		fullUrl = fmt.Sprintf("%s?%s", fullUrl, query)
	}
	c.Opt.Logger.Debug("parsed request",
		"method", r.method,
		"path", r.path,
		"auth_type", r.authType,
		"full_url", fullUrl,
	)
	c.fullUrl = fullUrl
	r.header = header
	return nil
}

func (c *Client) RawBody() []byte {
	return c.resp.rawBody
}

func (c *Client) RawHeader() http.Header {
	return c.resp.rawHeader
}

func (c *Client) Invoke(r *Request, ctx context.Context) error {
	return c.invoke(r, ctx)
}

func (c *Client) invoke(r *Request, ctx context.Context) error {
	if err := c.parseRequest(r); err != nil {
		return err
	}
	req, err := http.NewRequest(r.method, c.fullUrl, r.body)
	if err != nil {
		c.Opt.Logger.Debug("failed to create new HTTP request", "error", err)
		c.resp = &response{err: err}
		return err
	}
	req = req.WithContext(ctx)
	if r.header != nil {
		req.Header = r.header
	}
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.resp = &response{err: err}
		return err
	}
	data, err := io.ReadAll(res.Body)
	if err != nil {
		c.resp = &response{err: err}
		return err
	}
	c.Opt.Logger.Debug("received HTTP response", "status", res.StatusCode)
	defer res.Body.Close()
	c.resp = &response{rawBody: data, status: res.StatusCode, rawHeader: res.Header}
	if res.StatusCode != 200 {
		c.Opt.Logger.Debug("HTTP response returned non-200", "status", res.StatusCode, "body", string(data))
		c.resp.err = errors.New(string(data))
		return c.resp.err
	}
	return nil
}

type WsClient struct {
	Opt     *Options
	conn    *websocket.Conn
	writeCh chan *WsRequest

	writerOnce   sync.Once
	cancelWriter context.CancelFunc
	writerWg     sync.WaitGroup
}

// connect initializes the WebSocket connection.
func (c *WsClient) connect(ctx context.Context) error {
	conn, resp, err := websocket.DefaultDialer.DialContext(ctx, c.Opt.Endpoint, nil)
	if err != nil {
		c.Opt.Logger.Debug("websocket dial failed", "endpoint", c.Opt.Endpoint, "error", err)
		return err
	}
	c.Opt.Logger.Debug("websocket connection established", "endpoint", c.Opt.Endpoint, "status", resp.Status)
	c.conn = conn
	return nil
}
func (c *WsClient) SetReq(method string, aType ...AuthType) *WsRequest {
	reqType := AuthNone
	if len(aType) > 0 {
		reqType = aType[0]
	}
	return &WsRequest{Method: method, AuthType: reqType, Params: make(map[string]any)}
}
func (c *WsClient) Close() error {
	if c.cancelWriter != nil {
		c.cancelWriter()
		c.writerWg.Wait()
	}
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *WsClient) Combined(combine bool) {
	c.combined(combine)
}

func (c *WsClient) combined(combine bool) {
	if combine {
		c.Opt.Endpoint = c.Opt.Endpoint + "/stream?streams="
	} else {
		c.Opt.Endpoint = c.Opt.Endpoint + "/ws"
	}
}

func (c *WsClient) keepAlive() {
	ticker := time.NewTicker(WebsocketStreamsTimeout)

	lastResponse := time.Now()
	c.conn.SetPongHandler(func(msg string) error {
		lastResponse = time.Now()
		c.Opt.Logger.Debug("received pong", "time", lastResponse.Format(time.RFC3339))
		return nil
	})

	go func() {
		defer ticker.Stop()
		c.Opt.Logger.Debug("websocket keepalive started", "timeout", WebsocketStreamsTimeout.String())
		for {
			deadline := time.Now().Add(10 * time.Second)
			err := c.conn.WriteControl(websocket.PingMessage, []byte{}, deadline)
			if err != nil {
				c.Opt.Logger.Debug("failed to send ping", "error", err)
				return
			}
			c.Opt.Logger.Debug("ping sent", "deadline", deadline.Format(time.RFC3339))
			<-ticker.C
			if time.Since(lastResponse) > WebsocketStreamsTimeout {
				return
			}
		}
	}()
}

func (c *WsClient) WsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.wsServe(ctx)
}

func (c *WsClient) wsServe(ctx context.Context) (<-chan []byte, <-chan error) {
	onMessage := make(chan []byte, 8)
	onError := make(chan error)

	go func() {
		defer func() {
			close(onMessage)
			close(onError)
			c.Opt.Logger.Debug("websocket serve goroutine exited")
		}()

		if err := c.connect(ctx); err != nil {
			c.Opt.Logger.Debug("websocket connect failed", "error", err)
			onError <- err
			return
		}
		defer c.conn.Close()
		c.keepAlive()
		for {
			select {
			case <-ctx.Done():
				c.Opt.Logger.Debug("context done, websocket serve stopping")
				return
			default:
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					c.Opt.Logger.Debug("failed to read message from websocket", "error", err)
					onError <- err
					return
				}
				c.Opt.Logger.Debug("websocket message received", "length", len(message))
				onMessage <- message
			}
		}
	}()
	return onMessage, onError
}

func (c *WsClient) WsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	return c.wsApiServe(ctx)
}

func (c *WsClient) wsApiServe(ctx context.Context) (<-chan []byte, <-chan error) {
	onMessage := make(chan []byte, 8)
	onError := make(chan error)
	err := c.connect(ctx)
	c.Opt.Logger.Debug("attempting websocket connection", "endpoint", c.Opt.Endpoint)
	go func() {
		defer func() {
			close(onMessage)
			close(onError)
			c.Opt.Logger.Debug("wsApiServe goroutine exited")
		}()
		if err != nil {
			c.Opt.Logger.Debug("websocket connect failed", "error", err)
			onError <- err
			return
		}
		defer c.conn.Close()
		c.keepAlive()
		// Initialize writer here after successful connection (optional, but ensures writer is ready)
		c.lazyInitWriter()
		for {
			select {
			case <-ctx.Done():
				c.Opt.Logger.Debug("context done, exiting wsApiServe loop")
				return
			default:
				_, message, err := c.conn.ReadMessage()
				if err != nil {
					c.Opt.Logger.Debug("error reading websocket message", "error", err)
					onError <- err
					return
				}
				c.Opt.Logger.Debug("received websocket message", "length", len(message))
				onMessage <- message
			}
		}
	}()
	return onMessage, onError
}

func uuid4() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]byte, 16)
	if _, err := r.Read(data); err != nil {
		return ""
	}
	data[6] = (data[6] & 0x0f) | 0x40
	data[8] = (data[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		data[0:4], data[4:6], data[6:8], data[8:10], data[10:])
}

func (c *WsClient) Send(r *WsRequest) error {
	return c.send(r)
}

func (c *WsClient) send(r *WsRequest) error {
	if c.conn == nil {
		c.Opt.Logger.Debug("cannot send: connection is nil")
		return errors.New("websocket connection is nil")
	}
	r.Id = uuid4()
	c.Opt.Logger.Debug("generating request ID", "id", r.Id)
	if r.AuthType == AuthSigned {
		r.Params["timestamp"] = time.Now().UnixMilli()
	}
	if r.AuthType == AuthApiKey || r.AuthType == AuthSigned {
		r.Params["apiKey"] = c.Opt.ApiKey
	}

	if r.AuthType == AuthSigned {
		var sf SignFunc
		if c.Opt.SignType == SignTypeRsa {
			sf = RsaSign
		} else if c.Opt.SignType == SignTypeEd25519 {
			sf = Ed25519Sign
		} else {
			sf = HmacSign
		}
		sortedData := SortMap(r.Params)
		c.Opt.Logger.Debug("sorted params for signature", "params", sortedData)
		sign, err := sf(c.Opt.ApiSecret, sortedData)
		if err != nil {
			c.Opt.Logger.Debug("signature generation failed", "error", err)
			return err
		}
		r.Params["signature"] = sign
		c.Opt.Logger.Debug("signature added to request", "signature", sign)
	}

	select {
	case c.writeCh <- r:
		return nil
	case <-time.After(5 * time.Second):
		c.Opt.Logger.Warn("Failed to send request to write channel: channel full or blocked.", "request_id", r.Id)
		return errors.New("failed to send request: write channel full or blocked")
	}
}

// lazyInitWriter initializes the write channel and starts the writer goroutine
// This method is called internally within the WsClient.
func (c *WsClient) lazyInitWriter() {
	c.writerOnce.Do(func() {
		c.writeCh = make(chan *WsRequest, 16) // Initialize with a suitable buffer size
		var ctx context.Context
		ctx, c.cancelWriter = context.WithCancel(context.Background())
		c.writerWg.Add(1)
		go c.writerGoroutine(ctx)
		c.Opt.Logger.Debug("writer goroutine started successfully")
	})
}

// writerGoroutine is a goroutine that continuously writes messages to the WebSocket connection.
func (c *WsClient) writerGoroutine(ctx context.Context) {
	defer c.writerWg.Done()
	for {
		select {
		case <-ctx.Done():
			c.Opt.Logger.Debug("Context cancelled, writer Goroutine exiting.")
			return
		case req, ok := <-c.writeCh:
			if !ok {
				c.Opt.Logger.Debug("Write channel closed, writer Goroutine exiting.")
				return
			}
			if c.conn == nil {
				c.Opt.Logger.Warn("WebSocket connection is nil, discarding request.", "id", req.Id)
				continue
			}
			if err := c.conn.WriteJSON(req); err != nil {
				c.Opt.Logger.Error("Failed to write JSON to WebSocket", "error", err, "request_id", req.Id)
			}
		}
	}
}
