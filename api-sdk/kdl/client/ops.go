package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/endpoint"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/utils"
)

// GetProxy 获取开放代理
// return: 代理slice
func (client Client) GetProxy(num int, orderLevel string, signType signtype.SignType, kwargs map[string]interface{}) ([]string, error) {
	var ep endpoint.EndPoint
	if orderLevel == "normal" || orderLevel == "vip" {
		ep = endpoint.GetOpsProxyNormalOrVip
	} else if orderLevel == "svip" {
		ep = endpoint.GetOpsProxySvip
	} else if orderLevel == "ent" {
		ep = endpoint.GetOpsProxyEnt
	}
	if kwargs != nil {
		kwargs["num"] = num
	} else {
		kwargs = map[string]interface{}{"num": num}
	}
	params := client.getParams(ep, signType, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		if res.Code == 1043 { // format 不为json, 将原始字符串放入[]string并返回
			if strings.HasPrefix(res.Msg, "ERROR") {
				return []string{}, errors.New("KdlError: KdlError: " + res.Msg)
			}
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

// CheckOpsValid 检测私密代理有效性
// return: map[string]bool, key: proxy, value: true/false
func (client Client) CheckOpsValid(proxyList []string, signType signtype.SignType) (map[string]bool, error) {
	if len(proxyList) == 0 {
		return map[string]bool{}, errors.New("KdlError: proxyList can't be empty")
	}
	proxy := strings.Join(proxyList, ",")
	ep := endpoint.CheckOpsValid
	params := client.getParams(ep, signType, map[string]interface{}{"proxy": proxy})
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return map[string]bool{}, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		var rv = map[string]bool{}
		for proxy, valid := range data {
			if v, ok := valid.(bool); ok {
				rv[proxy] = v
			}
		}
		return rv, nil
	}
	return map[string]bool{}, errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}
