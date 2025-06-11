package spot

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type apiUserdataTestSuite struct {
	baseWsTestSuite
}

func TestWebsocketApiUserdata(t *testing.T) {
	suite.Run(t, new(apiUserdataTestSuite))
}

func (s *apiUserdataTestSuite) TestSessionLogon() {
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
		"apiKey": "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"authorizedSince": 1649729878532,
		"connectedSince": 1649729873021,
		"returnRateLimits": false,
		"serverTime": 1649729878630,
		"userDataStream": false
		}
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewSessionLogon().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *SessionResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestSessionResponse(resp, testResp)
}

func (s *apiUserdataTestSuite) assertTestSessionResponse(r1, r2 *SessionResponse) {
	s.assertWsResponse(r1.ApiResponse, r2.ApiResponse)
	s.assertTestSessionResult(r1.Result, r2.Result)
}

func (s *apiUserdataTestSuite) assertTestSessionResult(r1, r2 *SessionResult) {
	r := s.r()
	r.Equal(r1.ApiKey, r2.ApiKey, "apiKey")
	r.Equal(r1.AuthorizedSince, r2.AuthorizedSince, "authorizedSince")
	r.Equal(r1.ConnectedSince, r2.ConnectedSince, "connectedSince")
	r.Equal(r1.ReturnRateLimits, r2.ReturnRateLimits, "returnRateLimits")
	r.Equal(r1.ServerTime, r2.ServerTime, "serverTime")
	r.Equal(r1.UserDataStream, r2.UserDataStream, "userDataStream")
}

func (s *apiUserdataTestSuite) TestSessionStatus() {
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
		"apiKey": "vmPUZE6mv9SD5VNHk4HlWFsOr6aKE2zvsw0MuIgwCIPy6utIco14y7Ju91duEh8A",
		"authorizedSince": 1649729878532,
		"connectedSince": 1649729873021,
		"returnRateLimits": false,
		"serverTime": 1649729878630,
		"userDataStream": false
	}
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewSessionStatus().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *SessionResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestSessionResponse(resp, testResp)
}

func (s *apiUserdataTestSuite) TestSessionLogout() {
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
		"apiKey": null,
    	"authorizedSince": null,
		"connectedSince": 1649729873021,
		"returnRateLimits": false,
		"serverTime": 1649729878630,
		"userDataStream": false
	}
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewSessionLogout().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *SessionResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	s.assertTestSessionResponse(resp, testResp)
}
