package wss

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiUserdataTestSuite struct {
	baseTestSuite
}

func TestWebsocketApiUserdata(t *testing.T) {
	suite.Run(t, new(apiUserdataTestSuite))
}

func (s *apiUserdataTestSuite) TestStartUserDataStream() {
	msg := []byte(`{
	"id": "5ddbd23c-f137-4c63-bf95-82569cddc655",
	"status": 200,
	"rateLimits": [
		{
			"rateLimitType": "REQUEST_WEIGHT",
			"interval": "MINUTE",
			"intervalNum": 1,
			"limit": 6000,
			"count": 4
		}
	],
	"result": {
		"listenKey": "xxx"
	}
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewStartUserDataStream().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *StartUserDataStreamResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestStartUserDataStream(resp, testResp)
}

func (s *apiUserdataTestSuite) assertTestStartUserDataStream(r1, r2 *StartUserDataStreamResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	r := s.r()
	r.Equal(r1.Result.ListenKey, r2.Result.ListenKey, "listenKey")
}

func (s *apiUserdataTestSuite) TestPingUserDataStream() {
	msg := []byte(`{
	"id": "5ddbd23c-f137-4c63-bf95-82569cddc655",
	"status": 200,
	"rateLimits": [
		{
			"rateLimitType": "REQUEST_WEIGHT",
			"interval": "MINUTE",
			"intervalNum": 1,
			"limit": 6000,
			"count": 4
		}
	]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewPingUserDataStream().ListenKey("listenKey").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp ApiResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertWsResponse(*resp, testResp)
}

func (s *apiUserdataTestSuite) TestStopUserDataStream() {
	msg := []byte(`{
	"id": "5ddbd23c-f137-4c63-bf95-82569cddc655",
	"status": 200,
	"rateLimits": [
		{
			"rateLimitType": "REQUEST_WEIGHT",
			"interval": "MINUTE",
			"intervalNum": 1,
			"limit": 6000,
			"count": 4
		}
	]
}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewStopUserDataStream().ListenKey("listenKey").Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp ApiResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertWsResponse(*resp, testResp)
}
