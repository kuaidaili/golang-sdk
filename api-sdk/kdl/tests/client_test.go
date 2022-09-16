package tests

import (
	"fmt"
	"testing"

	"kdl/auth"
	"kdl/client"
	"kdl/signtype"
)

var secret_id = ""
var secret_key = ""

var dpsAuth = auth.Auth{SecretID: secret_id, SecretKey: secret_key}
var dpsClient = client.Client{Auth: dpsAuth}

var dpsCountAuth = auth.Auth{SecretID: secret_id, SecretKey: secret_key}
var dpsCountClient = client.Client{Auth: dpsCountAuth}

var kpsAuth = auth.Auth{SecretID: secret_id, SecretKey: secret_key}
var kpsClient = client.Client{Auth: kpsAuth}

var opsAuth = auth.Auth{SecretID: secret_id, SecretKey: secret_key}
var opsClient = client.Client{Auth: opsAuth}

var tpsAuth = auth.Auth{SecretID: secret_id, SecretKey: secret_key}
var tpsClient = client.Client{Auth: tpsAuth}

func TestGetOrderExpireTime(t *testing.T) {
	expireTime, err := dpsClient.GetOrderExpireTime(signtype.Token)
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
	balance, err := dpsCountClient.GetIPBalance(signtype.Token)
	fmt.Println("balance: ", balance)
	if err != nil {
		t.Error(err)
	}
}

func TestGetProxyAuthorization(t *testing.T) {
	balance, err := dpsCountClient.GetProxyAuthorization(0, signtype.Token)
	fmt.Println("proxyauthorization: ", balance)
	if err != nil {
		t.Error(err)
	}
}

func TestGetSecretToken(t *testing.T) {
	secret_token := dpsCountClient.GetSecretToken()
	fmt.Println("secret_token: ", secret_token)
}
