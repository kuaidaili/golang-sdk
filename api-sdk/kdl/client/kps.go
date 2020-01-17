package client

import (
	"errors"
	"fmt"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/endpoint"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/utils"
)

// GetKps 获取独享代理
// return: 代理slice
func (client Client) GetKps(num int, signType signtype.SignType, kwargs map[string]interface{}) ([]string, error) {
	ep := endpoint.GetKpsProxy
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
