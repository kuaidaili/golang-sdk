package tests

import (
	"fmt"
	"strings"
	"testing"

	"kdl/opslevel"
	"kdl/signtype"
)

func TestGetProxy(t *testing.T) {
	params := make(map[string]interface{})
	params["area"] = "北京,上海"
	ips, err := opsClient.GetProxy(2, opslevel.NORMAL, signtype.SIMPLE, params)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
}

func TestCheckOpsValid(t *testing.T) {
	params := make(map[string]interface{})
	params["area"] = "北京,上海"
	ips, err := opsClient.GetProxy(2, opslevel.NORMAL, signtype.SIMPLE, params)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
	ips = strings.Split(ips[0], "\n")
	valids, err := opsClient.CheckOpsValid(ips, signtype.HmacSha1)
	fmt.Println("valids: ", valids)
	if err != nil {
		fmt.Println(err)
	}
}
