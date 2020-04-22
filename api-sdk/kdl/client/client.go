package client

import (
	"errors"
	"fmt"
	"strings"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/endpoint"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/utils"
)

// GetOrderExpireTime 获取订单过期时间
// return: 过期时间字符串
func (client Client) GetOrderExpireTime(signType signtype.SignType) (string, error) {
	ep := endpoint.GetOrderExpireTime
	params := client.getParams(ep, signType, nil)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return "", err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if expireTime, ok := data["expire_time"].(string); ok {
			return expireTime, nil
		}
	}
	return "", errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}

//GetProxyAuthorization 获取代理鉴权信息
// return: 鉴权信息字典
func (client Client) GetProxyAuthorization(plaintext int, signType signtype.SignType) (map[string]string, error) {
	ret := make(map[string]string)
	ep := endpoint.GetProxyAuthorization
	params := client.getParams(ep, signType, map[string]interface{}{"plaintext": plaintext})
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return ret, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if value, ok := data["type"].(string); ok {
			ret["type"] = value
		}
		if value, ok := data["credentials"].(string); ok {
			ret["credentials"] = value
		}
		if plaintext == 1 {
			if value, ok := data["username"].(string); ok {
				ret["username"] = value
			}
			if value, ok := data["password"].(string); ok {
				ret["password"] = value
			}
		}
	}
	return ret, err
}


// GetIPWhitelist 获取订单的IP白名单
// return: ip白名单slice
func (client Client) GetIPWhitelist(signType signtype.SignType) ([]string, error) {
	ep := endpoint.GetIPWhitelist
	params := client.getParams(ep, signType, nil)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return []string{}, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if count, ok := data["count"].(float64); ok && count == 0 {
			return []string{}, nil
		} else if count > 0 {
			if whitelist, ok := data["ipwhitelist"].([]interface{}); ok {
				var ips []string
				for _, v := range whitelist {
					ips = append(ips, utils.TypeSwitcher(v))
				}
				return ips, nil
			}
			return []string{}, nil
		}
	}
	return []string{}, errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}

// SetIPWhitelist 设置IP白名单
// return: true/false
func (client Client) SetIPWhitelist(ipList []string, signType signtype.SignType) (bool, error) {
	ips := strings.Join(ipList, ",")
	ep := endpoint.SetIPWhitelist
	params := client.getParams(ep, signType, map[string]interface{}{"iplist": ips})
	_, err := client.getBaseRes("POST", ep, params)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetIPBalance 获取计数版订单ip余额
// return: ip余额
func (client Client) GetIPBalance(signType signtype.SignType) (int, error) {
	ep := endpoint.GetIPBalance
	params := client.getParams(ep, signType, nil)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return -1, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if balance, ok := data["balance"].(float64); ok {
			return int(balance), nil
		}
	}
	return -1, errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}
