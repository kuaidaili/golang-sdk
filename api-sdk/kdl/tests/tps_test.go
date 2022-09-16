package tests

import (
	"fmt"
	"testing"

	"kdl/signtype"
)

func TestTpsCurrentIP(t *testing.T) {
	curIP, err := tpsClient.TpsCurrentIP(signtype.HmacSha1)
	fmt.Println("curIp: ", curIP)
	if err != nil {
		t.Error(err)
	}
}

func TestChangeTpsIP(t *testing.T) {
	newIP, err := tpsClient.ChangeTpsIP(signtype.HmacSha1)
	fmt.Println("newIp: ", newIP)
	if err != nil {
		t.Error(err)
	}
}
