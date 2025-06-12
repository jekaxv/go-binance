package futures

import (
	"github.com/gorilla/websocket"
	"github.com/jekaxv/go-binance/core"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"log/slog"
	"net/http"
	"net/http/httptest"
)

type mockedWsClient struct {
	mock.Mock
	*WsClient
}

type baseWsTestSuite struct {
	suite.Suite
	client *mockedWsClient
}

func (s *baseWsTestSuite) mockClient(url string) {
	s.client.WsClient.Opt.Endpoint = url
}

func (s *baseWsTestSuite) SetupTest() {
	s.client = new(mockedWsClient)
	client := WsClient{
		&core.WsClient{
			Opt: &core.Options{
				ApiKey:    "YOUR_API_KEY",
				ApiSecret: "YOUR_API_SECRET",
				Logger:    slog.Default(),
			},
		},
	}
	s.client.WsClient = &client
}

func (s *baseWsTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseWsTestSuite) mockServer(msg []byte) *httptest.Server {
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

func (s *baseWsTestSuite) setup(msg []byte) *httptest.Server {
	server := s.mockServer(msg)
	s.mockClient("ws" + server.URL[4:])
	return server
}

func (s *baseWsTestSuite) assertWsResponse(r1, r2 ApiResponse) {
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

func (s *baseWsTestSuite) assertWsApiError(r1, r2 *ApiError) {
	r := s.r()
	r.Equal(r1.Code, r2.Code, "code")
	r.Equal(r1.Msg, r2.Msg, "msg")
}

func (s *baseWsTestSuite) assertWsRateLimits(r1, r2 *ApiRateLimit) {
	r := s.r()
	r.Equal(r1.Count, r2.Count, "count")
	r.Equal(r1.Interval, r2.Interval, "interval")
	r.Equal(r1.IntervalNum, r2.IntervalNum, "intervalNum")
	r.Equal(r1.Limit, r2.Limit, "limit")
	r.Equal(r1.RateLimitType, r2.RateLimitType, "rateLimitType")
}
