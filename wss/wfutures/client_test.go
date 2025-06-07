package wfutures

import (
	"github.com/gorilla/websocket"
	"github.com/jekaxv/go-binance/wss"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"net/http"
	"net/http/httptest"
)

type mockedClient struct {
	mock.Mock
	*Client
}

type baseTestSuite struct {
	suite.Suite
	client *mockedClient
}

func (s *baseTestSuite) mockClient(url string) {
	s.client.Client.C.Opt.Endpoint = url
}

func (s *baseTestSuite) SetupTest() {
	s.client = new(mockedClient)
	client := Client{
		C: &wss.Client{
			Opt: &wss.Options{
				ApiKey:    "YOUR_API_KEY",
				ApiSecret: "YOUR_API_SECRET",
			},
			Logger: slog.Default(),
		},
	}
	s.client.Client = &client
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) mockServer(msg []byte) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		defer conn.Close()

		for {
			if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
				return
			}
		}
	}))
}

func (s *baseTestSuite) setup(msg []byte) *httptest.Server {
	server := s.mockServer(msg)
	s.mockClient("ws" + server.URL[4:])
	return server
}

func (s *baseTestSuite) assertWsResponse(r1, r2 wss.ApiResponse) {
	r := s.r()
	r.Equal(r1.Id, r2.Id, "id")
	r.Equal(r1.Status, r2.Status, "status")
	if r1.Error != nil {
		s.assertWsApiError(r1.Error, r2.Error)
	}
	for i := range r1.RateLimits {
		s.assertWsRateLimits(r1.RateLimits[i], r2.RateLimits[i])
	}
}

func (s *baseTestSuite) assertWsApiError(r1, r2 *wss.ApiError) {
	r := s.r()
	r.Equal(r1.Code, r2.Code, "code")
	r.Equal(r1.Msg, r2.Msg, "msg")
}

func (s *baseTestSuite) assertWsRateLimits(r1, r2 *wss.ApiRateLimit) {
	r := s.r()
	r.Equal(r1.Count, r2.Count, "count")
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.IntervalNum, r2.IntervalNum, "intervalNum")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
}
