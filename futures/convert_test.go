package futures

import (
	"context"
	"encoding/json"
	"github.com/stretchr/testify/suite"
	"testing"
)

type convertTestSuite struct {
	baseHttpTestSuite
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

func (s *convertTestSuite) TestNewGetQuote() {
	msg := []byte(`{
	   "quoteId":"12415572564",
	   "ratio":"38163.7",
	   "inverseRatio":"0.0000262",
	   "validTimestamp":1623319461670,
	   "toAmount":"3816.37",
	   "fromAmount":"0.1"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewGetQuote().FromAsset("BTC").
		ToAsset("USDT").
		FromAmount(0.0004).
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *GetQuoteResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.QuoteId, testResp.QuoteId, "QuoteId")
	r.Equal(resp.Ratio, testResp.Ratio, "Ratio")
	r.Equal(resp.InverseRatio, testResp.InverseRatio, "InverseRatio")
	r.Equal(resp.ValidTimestamp, testResp.ValidTimestamp, "ValidTimestamp")
	r.Equal(resp.ToAmount, testResp.ToAmount, "ToAmount")
	r.Equal(resp.FromAmount, testResp.FromAmount, "FromAmount")
}

func (s *convertTestSuite) TestNewAcceptQuote() {
	msg := []byte(`{
	  "orderId":"933256278426274426",
	  "createTime":1623381330472,
	  "orderStatus":"PROCESS"
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewAcceptQuote().QuoteId("12415572564").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *AcceptQuoteResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.OrderId, testResp.OrderId, "OrderId")
	r.Equal(resp.CreateTime, testResp.CreateTime, "CreateTime")
	r.Equal(resp.OrderStatus, testResp.OrderStatus, "OrderStatus")
}

func (s *convertTestSuite) TestNewConvertOrderStatus() {
	msg := []byte(`{
	  "orderId":933256278426274426,
	  "orderStatus":"SUCCESS",
	  "fromAsset":"BTC",
	  "fromAmount":"0.00054414",
	  "toAsset":"USDT",
	  "toAmount":"20",
	  "ratio":"36755",
	  "inverseRatio":"0.00002721",
	  "createTime":1623381330472
	}`)
	server := s.setup(msg)
	defer server.Close()
	resp, err := s.client.NewConvertOrderStatus().QuoteId("12415572564").
		Do(context.Background())
	r := s.r()
	r.Empty(err)
	var testResp *ConvertOrderStatusResponse
	r.Empty(json.Unmarshal(msg, &testResp))
	r.Equal(resp.OrderId, testResp.OrderId, "OrderId")
	r.Equal(resp.OrderStatus, testResp.OrderStatus, "OrderStatus")
	r.Equal(resp.FromAsset, testResp.FromAsset, "FromAsset")
	r.Equal(resp.FromAmount, testResp.FromAmount, "FromAmount")
	r.Equal(resp.ToAsset, testResp.ToAsset, "ToAsset")
	r.Equal(resp.ToAmount, testResp.ToAmount, "ToAmount")
	r.Equal(resp.Ratio, testResp.Ratio, "Ratio")
	r.Equal(resp.InverseRatio, testResp.InverseRatio, "InverseRatio")
	r.Equal(resp.CreateTime, testResp.CreateTime, "CreateTime")
}
