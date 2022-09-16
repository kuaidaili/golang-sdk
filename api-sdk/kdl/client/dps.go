package client

import (
	"errors"
	"fmt"
	"strings"

	"kdl/endpoint"
	"kdl/signtype"
	"kdl/utils"
)

// GetDps 获取私密代理
// return: 代理slice
func (client Client) GetDps(num int, signType signtype.SignType, kwargs map[string]interface{}) ([]string, error) {
	ep := endpoint.GetDpsProxy
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

// CheckDpsValid 检测私密代理有效性
// return: map[string]bool, key: proxy, value:true/false
func (client Client) CheckDpsValid(proxyList []string, signType signtype.SignType) (map[string]bool, error) {
	if len(proxyList) == 0 {
		return map[string]bool{}, errors.New("KdlError: proxyList can't be empty")
	}
	proxy := strings.Join(proxyList, ",")
	ep := endpoint.CheckDpsValid
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

// GetDpsValidTime 检测私密代理的有效时间(剩余秒数)
// return: map[string]bool, key: proxy, value:true/false
func (client Client) GetDpsValidTime(proxyList []string, signType signtype.SignType) (map[string]string, error) {
	if len(proxyList) == 0 {
		return map[string]string{}, errors.New("KdlError: proxyList can't be empty")
	}
	proxy := strings.Join(proxyList, ",")
	ep := endpoint.GetDpsValidTime
	params := client.getParams(ep, signType, map[string]interface{}{"proxy": proxy})
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return map[string]string{}, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		var rv = map[string]string{}
		for proxy, seconds := range data {
			if v, ok := seconds.(float64); ok {
				rv[proxy] = utils.TypeSwitcher(v)
			}
		}
		return rv, nil
	}
	return map[string]string{}, errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}
