package tests

import (
	"fmt"
	"testing"

	"kdl/signtype"
)

func TestGetDps(t *testing.T) {
	params := make(map[string]interface{})
	params["format"] = "json"
	params["area"] = "北京,上海"
	ips, err := dpsClient.GetDps(2, signtype.HmacSha1, params)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckDpsValid(t *testing.T) {
	params := make(map[string]interface{})
	params["format"] = "json"
	ips, err := dpsClient.GetDps(2, signtype.HmacSha1, params)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
	valids, err := dpsClient.CheckDpsValid(ips, signtype.SIMPLE)
	fmt.Println("valids: ", valids)
	if err != nil {
		t.Error(err)
	}
}

func TestGetDpsValidTime(t *testing.T) {
	params := make(map[string]interface{})
	params["format"] = "json"
	ips, err := dpsClient.GetDps(2, signtype.HmacSha1, params)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
	seconds, err := dpsClient.GetDpsValidTime(ips, signtype.HmacSha1)
	fmt.Println("seconds: ", seconds)
	if err != nil {
		t.Error(err)
	}
}
