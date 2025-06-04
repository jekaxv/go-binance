package hfutures

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type convertTestSuite struct {
	baseTestSuite
}

func TestConvert(t *testing.T) {
	suite.Run(t, new(convertTestSuite))
}

func (s *convertTestSuite) TestNewConvertExchangeInfo() {
	msg := []byte(`[
	  {
		"fromAsset":"BTC",
		"toAsset":"USDT",
		"fromAssetMinAmount":"0.0004",
		"fromAssetMaxAmount":"50",
		"toAssetMinAmount":"20",
		"toAssetMaxAmount":"2500000"
	  }
	]`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewConvertExchangeInfo().Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp []*ConvertExchangeInfoResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	for i := range testResp {
		s.assertConvertExchangeInfoResponse(resp[i], testResp[i])
	}
}

func (s *convertTestSuite) assertConvertExchangeInfoResponse(r1, r2 *ConvertExchangeInfoResponse) {
	r := s.r()
	r.Equal(r1.FromAsset, r2.FromAsset, "FromAsset")
	r.Equal(r1.ToAsset, r2.ToAsset, "ToAsset")
	r.Equal(r1.FromAssetMaxAmount, r2.FromAssetMaxAmount, "FromAssetMaxAmount")
	r.Equal(r1.FromAssetMinAmount, r2.FromAssetMinAmount, "FromAssetMinAmount")
	r.Equal(r1.ToAsset, r2.ToAsset, "ToAsset")
	r.Equal(r1.ToAssetMaxAmount, r2.ToAssetMaxAmount, "ToAssetMaxAmount")
}
