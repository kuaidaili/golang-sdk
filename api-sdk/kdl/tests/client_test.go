package tests

import (
	"fmt"
	"testing"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/auth"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/client"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
)

var dpsAuth = auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
var dpsClient = client.Client{Auth: dpsAuth}

var dpsCountAuth = auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
var dpsCountClient = client.Client{Auth: dpsCountAuth}

var kpsAuth = auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
var kpsClient = client.Client{Auth: kpsAuth}

var opsAuth = auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
var opsClient = client.Client{Auth: opsAuth}

var tpsAuth = auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
var tpsClient = client.Client{Auth: tpsAuth}

func TestGetOrderExpireTime(t *testing.T) {
	expireTime, err := dpsClient.GetOrderExpireTime(signtype.HmacSha1)
	fmt.Println("expireTime: ", expireTime)
	if err != nil {
		t.Error(err)
	}
}

func TestGetIPWhitelist(t *testing.T) {
	ips, err := dpsClient.GetIPWhitelist(signtype.SIMPLE)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
}

func TestSetIPWhitelist(t *testing.T) {
	_, err := dpsClient.SetIPWhitelist([]string{}, signtype.HmacSha1)
	if err != nil {
		t.Error(err)
	}
	_, err = dpsClient.SetIPWhitelist([]string{"102.173.59.220"}, signtype.HmacSha1)
	if err != nil {
		t.Error(err)
	}
	_, err = dpsClient.SetIPWhitelist([]string{"102.173.59.220", "103.173.59.220"}, signtype.HmacSha1)
	if err != nil {
		t.Error(err)
	}
	_, err = dpsClient.SetIPWhitelist([]string{"102.173.59.220", "103.173.59.220", "104.173.59.220"}, signtype.HmacSha1)
	if err == nil {
		t.Error("fail to assert over limit error")
	}
}

func TestGetIPBalance(t *testing.T) {
	balance, err := dpsCountClient.GetIPBalance(signtype.HmacSha1)
	fmt.Println("balance: ", balance)
	if err != nil {
		t.Error(err)
	}
}
