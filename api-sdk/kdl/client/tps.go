package client

import (
	"errors"
	"fmt"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/endpoint"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/utils"
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

//GetTpsIp 获取隧道代理IP, 获取订单对应的隧道代理IP。
func (client Client) GetTpsIp(num int, signType signtype.SignType, kwargs map[string]interface{}) ([]string, error) {
	ep := endpoint.GetTpsIp
	if kwargs != nil {
		kwargs["num"] = num
	} else {
		kwargs = map[string]interface{}{"num": num}
	}
	params := client.getParams(ep, signType, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		if res.Code == 1043 { // format 不为json, 将原始字符串放入[]string并返回
			return []string{res.Msg}, nil
		}
		return []string{}, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if count, ok := data["count"].(float64); ok && count == 0 {
			return []string{}, nil
		} else if count > 0 {
			if proxies, ok := data["proxy_list"].([]interface{}); ok {
				var ips []string
				for _, v := range proxies {
					ips = append(ips, utils.TypeSwitcher(v))
				}
				return ips, nil
			}
			return []string{}, nil
		}
	}
	return []string{}, errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}
