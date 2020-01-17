package client

import (
	"errors"
	"fmt"
	"kdl/endpoint"
	"kdl/signtype"
)

// TpsCurrentIP 获取隧道当前IP, 仅支持支持换IP周期>=1分钟的隧道代理订单
func (client Client) TpsCurrentIP(signType signtype.SignType) (string, error) {
	ep := endpoint.TpsCurrentIP
	params := client.getParams(ep, signType, nil)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return "", err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if curIP, ok := data["current_ip"].(string); ok {
			return curIP, nil
		}
	}
	return "", errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}

// ChangeTpsIP 更换隧道代理IP, 返回新的IP地址, 仅支持支持换IP周期>=1分钟的隧道代理订单
func (client Client) ChangeTpsIP(signType signtype.SignType) (string, error) {
	ep := endpoint.ChangeTpsIP
	params := client.getParams(ep, signType, nil)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return "", err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if curIP, ok := data["new_ip"].(string); ok {
			return curIP, nil
		}
	}
	return "", errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}
