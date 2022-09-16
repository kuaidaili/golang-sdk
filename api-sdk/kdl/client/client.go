package client

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"kdl/endpoint"
	"kdl/signtype"
	"kdl/utils"
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

//GetUA 获取User Agent
// return: user agent数组
func (client Client) GetUA(num int, signType signtype.SignType) ([]string, error) {
	ep := endpoint.GetUA
	kwargs := make(map[string]interface{})
	kwargs["num"] = num
	params := client.getParams(ep, signType, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return []string{}, err
	}
	if data, ok := res.Data.(map[string]interface{}); ok {
		if count, ok := data["count"].(float64); ok && count == 0 {
			return []string{}, nil
		} else if count > 0 {
			if ua_list, ok := data["ua_list"].([]interface{}); ok {
				var uaList []string
				for _, v := range ua_list {
					uaList = append(uaList, utils.TypeSwitcher(v))
				}
				return uaList, nil
			}
			return []string{}, nil
		}
	}
	return []string{}, errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}

//GetAreaCode 获取指定地区编码
// return: 地区编码信息字典
func (client Client) GetAreaCode(area string, signType signtype.SignType) (map[string]string, error) {
	ret := make(map[string]string)
	ep := endpoint.GetAreaCode
	kwargs := make(map[string]interface{})
	kwargs["area"] = area
	params := client.getParams(ep, signType, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return ret, err
	}

	if data, ok := res.Data.(map[string]interface{}); ok {
		if area_name, ok := data["area_name"].(string); ok {
			ret["area_name"] = area_name
		}
		if area_code, ok := data["area_code"].(string); ok {
			ret["area_code"] = area_code
		}
	}
	return ret, nil
}

// GetAreaCode 获取账户余额
// return: 账户余额字符串
func (client Client) GetAccountBalance(signType signtype.SignType) (string, error) {
	ep := endpoint.GetAccountBalance
	params := client.getParams(ep, signType, nil)
	res, err := client.getBaseRes("GET", ep, params)
	if err != nil {
		return "", err
	}

	if data, ok := res.Data.(map[string]interface{}); ok {
		if balance, ok := data["balance"].(string); ok {
			return balance, nil
		}
	}
	return "", errors.New("KdlError: fail to parse response data: " + fmt.Sprint(res.Data))
}

// CreateOrder 创建订单，自动从账户余额里结算费用
// return: jsonRes struct
func (client Client) CreateOrder(product string, pay_type string) (jsonRes, error) {
	ep := endpoint.CreateOrder
	kwargs := make(map[string]interface{})
	kwargs["product"] = product
	kwargs["pay_type"] = pay_type
	params := client.getParams(ep, signtype.HmacSha1, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	return res, err
}

// GetOrderInfo 获取订单的详细信息
// return: jsonRes struct
func (client Client) GetOrderInfo() (jsonRes, error) {
	ep := endpoint.GetOrderInfo
	kwargs := make(map[string]interface{})
	params := client.getParams(ep, signtype.HmacSha1, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	return res, err
}

// SetAutoRenew 开启/关闭自动续费
// return: jsonRes struct
func (client Client) SetAutoRenew(autorenew string) (jsonRes, error) {
	ep := endpoint.SetAutoRenew
	kwargs := make(map[string]interface{})
	kwargs["autorenew"] = autorenew
	params := client.getParams(ep, signtype.HmacSha1, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	return res, err
}

// CloseOrder 关闭指定订单, 此接口只对按量付费(后付费)订单有效
// return: jsonRes struct
func (client Client) CloseOrder() (jsonRes, error) {
	ep := endpoint.CloseOrder
	kwargs := make(map[string]interface{})
	params := client.getParams(ep, signtype.HmacSha1, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	return res, err
}

// QueryKpsCity 查询独享代理有哪些城市可供开通。对于IP共享型还可查询到每个城市可开通的IP数量。
// return: jsonRes struct
func (client Client) QueryKpsCity(serie string) (jsonRes, error) {
	ep := endpoint.QueryKpsCity
	kwargs := make(map[string]interface{})
	kwargs["serie"] = serie
	params := client.getParams(ep, signtype.HmacSha1, kwargs)
	res, err := client.getBaseRes("GET", ep, params)
	return res, err
}

func (client Client) _ReadSecretToken() string {
	var secret_token string
	f, err := os.Open(SecretPath)
	if err != nil {
		f.Close()
		panic("fail to read .secret")
	}
	_byte, _ := ioutil.ReadAll(f)
	secretSlice := strings.Split(string(_byte), "|")
	secret_token = secretSlice[0]
	expire, _ := strconv.ParseInt(secretSlice[1], 10, 64)
	_time, _ := strconv.ParseInt(secretSlice[2], 10, 64)
	if (_time + expire - 3*60) < time.Now().Unix() { // 3分钟过期时重新获取
		secret_token, expire, _time = client._GetSecretToken()
		f, err := os.OpenFile(SecretPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			f.Close()
			panic("fail to write .secret")
		}
		f.WriteString(fmt.Sprintf("%v|%v|%v", secret_token, expire, _time))
	}
	return secret_token
}

func (client Client) _GetSecretToken() (string, int64, int64) {
	ep := endpoint.GetSecretToken
	kwargs := make(map[string]interface{})
	kwargs["secret_id"] = client.Auth.SecretID
	kwargs["secret_key"] = client.Auth.SecretKey
	params := client.getParams(ep, signtype.HmacSha1, kwargs)
	res, _ := client.getBaseRes("POST", ep, params)
	secret_token := res.Data.(map[string]interface{})["secret_token"].(string)
	expire_float64, _ := res.Data.(map[string]interface{})["expire"].(float64)
	expire := int64(expire_float64)
	_time := time.Now().Unix() // 秒级时间戳
	return secret_token, expire, _time
}

// GetSecretToken 获取密钥令牌
// return: jsonRes struct
func (client Client) GetSecretToken() string {
	var secret_token string
	if _, err := os.Stat(SecretPath); err == nil {
		// 存在secret文件
		secret_token = client._ReadSecretToken()
	} else {
		tmp_secret_token, expire, _time := client._GetSecretToken()
		secret_token = tmp_secret_token
		f, err := os.OpenFile(SecretPath, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			f.Close()
			panic("fail to write .secret")
		}
		f.WriteString(fmt.Sprintf("%v|%v|%v", secret_token, expire, _time))
	}
	return secret_token
}
