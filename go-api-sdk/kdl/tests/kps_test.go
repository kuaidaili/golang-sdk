package tests

import (
	"fmt"
	"kdl/signtype"
	"testing"
)

func TestGetKps(t *testing.T) {
	params := make(map[string]interface{})
	params["format"] = "json"
	ips, err := kpsClient.GetKps(2, signtype.HmacSha1, params)
	fmt.Println("ips: ", ips)
	if err != nil {
		t.Error(err)
	}
}
