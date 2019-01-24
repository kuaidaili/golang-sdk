package main

// 调用API接口

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Result & Data 为根据api的返回结果构造的结构体, 请注意大小写
type Result struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
	Data *Data  `json:"data"`
}

type Data struct {
	Count      int      `json:"count"`
	Proxy_list []string `json:"proxy_list"`
}

func get_proxies() (ips []string, err error) {
	// 生成的api链接
	api_url := "http://dev.kdlapi.com/api/getproxy/?orderid=965102959536478&num=100&protocol=1&method=2&an_ha=1&sep=1"
	r, err := http.Get(api_url)
	if err != nil { // 请求异常
		fmt.Printf("%s", err)
	}
	defer r.Body.Close()                    // 保证最后关闭Body
	contents, err := ioutil.ReadAll(r.Body) // 读取请求的返回体
	if err != nil {
		fmt.Printf("%s", err) // 读取出错
		os.Exit(1)
	}

	result := &Result{
		Msg:  "",
		Code: -1,
		Data: &Data{},
	}

	err = json.Unmarshal(contents, result)
	if err != nil { // 转化为Result对象出错
		return nil, err
	}

	if result.Code != 0 { // api接口返回异常
		return nil, errors.New(result.Msg)
	}

	ips = result.Data.Proxy_list
	return ips, nil
}

func main() {
	ips, err := get_proxies()
	if err != nil {
		fmt.Println("fail to get proxies", err) // 获取代理ip失败
		os.Exit(-1)
	}
	// 成功获取代理ip
	fmt.Println("proxies list:", ips)
}
