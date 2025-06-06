package hfutures

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type userdataTestSuite struct {
	baseTestSuite
}

func TestUserdata(t *testing.T) {
	suite.Run(t, new(userdataTestSuite))
}

func (s *userdataTestSuite) TestNewGetListenKey() {
	msg := []byte(`{
	  "listenKey": "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewGetListenKey().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ListenKeyResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.ListenKey, testResp.ListenKey, "ListenKey")
}

func (s *userdataTestSuite) TestNewKeepaliveListenKey() {
	msg := []byte(`{
	  "listenKey": "pqia91ma19a5s61cv6a81va65sdf19v8a65a1a5s61cv6a81va65sdf19v8a65a1"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewKeepaliveListenKey().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ListenKeyResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.ListenKey, testResp.ListenKey, "ListenKey")
}

func (s *userdataTestSuite) TestNewCloseListenKey() {
	msg := []byte(``)
	server := s.setup(msg)
	defer server.Close()
	err := s.client.NewCloseListenKey().Do(context.Background())
	r := s.r()
	r.Empty(err)
}
