package client

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kdl/endpoint"
	"kdl/signtype"
	"kdl/utils"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// getParams 构造请求参数, 返回参数map
func (client Client) getParams(ep endpoint.EndPoint, signType signtype.SignType, kwargs map[string]interface{}) map[string]interface{} {
	params := make(map[string]interface{})
	if client.Auth.OrderID == "" {
		panic("order id is required for auth")
	}
	params["orderid"] = client.Auth.OrderID
	params["sign_type"] = signType
	for k, v := range kwargs {
		params[k] = v
	}
	if signType == "" {
		return params
	}
	if client.Auth.APIKey == "" {
		panic("api key is required for signature")
	}
	if signType == signtype.SIMPLE {
		params["signature"] = client.Auth.APIKey
	} else if signType == signtype.HmacSha1 {
		params["timestamp"] = time.Now().Unix()
		var rawStr string
		if ep == endpoint.SetIPWhitelist {
			rawStr = client.Auth.GetStringToSign("POST", string(ep), params)
		} else {
			rawStr = client.Auth.GetStringToSign("GET", string(ep), params)
		}
		params["signature"] = client.Auth.SignStr(rawStr)
	} else {
		panic("unknown sign type" + signType)
	}
	return params
}

// getBaseRes, 公共http请求方法, 处理一些针对返回结果的公共逻辑
// return: jsonRes struct
func (client Client) getBaseRes(method string, endpoint endpoint.EndPoint, params map[string]interface{}) (jsonRes, error) {
	var r *http.Response
	var err error
	if method == "GET" {
		r, err = httpGet("https://"+string(endpoint), params)
	} else if method == "POST" {
		r, err = httpPost("https://"+string(endpoint), params)
	}
	var res jsonRes
	if err != nil {
		return res, err
	}
	defer r.Body.Close()
	if r.StatusCode != http.StatusOK {
		return res, errors.New("KdlError: request status code error: " + strconv.Itoa(r.StatusCode))
	}
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic("fail to read body bytes")
	}
	bodyString := string(bodyBytes)
	err = json.Unmarshal([]byte(bodyString), &res)
	if err != nil {
		// json解析失败，返回原始字符串
		return jsonRes{Code: 1043, Msg: bodyString}, errors.New("KdlError: fail to parse result content: " + bodyString)
	}
	if res.Code != 0 {
		return res, errors.New("KdlError: response status code error: " + strconv.Itoa(res.Code) + ", err msg: " + res.Msg)
	}
	return res, nil
}

// httpGet 发起Get请求
func httpGet(url string, params map[string]interface{}) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, errors.New("KdlError: new request is fail")
	}
	// add params
	q := req.URL.Query()
	if params != nil {
		for k, v := range params {
			q.Add(k, utils.TypeSwitcher(v))
		}
		// req.URL.RawQuery = q.Encode()
		req.URL.RawQuery = q.Encode()
	}
	client := &http.Client{}
	return client.Do(req)
}

// httpPost 发起Post请求
func httpPost(_url string, body map[string]interface{}) (*http.Response, error) {
	var req *http.Request
	form := url.Values{}
	for k, v := range body {
		form.Add(k, utils.TypeSwitcher(v))
	}
	req, err := http.NewRequest(http.MethodPost, _url, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, errors.New("KdlError: new request is fail")
	}
	client := &http.Client{}
	return client.Do(req)
}
