// 请求socks代理服务器(用户名密码认证)
// http和https网页均适用

package main

import (
	"compress/gzip"
	"fmt"
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// 用户名密码(私密代理/独享代理)
	username := "myusername"
	password := "mypassword"

	auth := proxy.Auth{
		User:     username,
		Password: password,
	}

	proxy_str := "59.38.241.25:23916"

	// 目标网页
	page_url := "http://dev.kdlapi.com/testproxy"

	// 设置代理
	dialer, err := proxy.SOCKS5("tcp", proxy_str, &auth, proxy.Direct)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// 请求目标网页
	client := &http.Client{Transport: &http.Transport{Dial: dialer.Dial}}
	req, _ := http.NewRequest("GET", page_url, nil)
	req.Header.Add("Accept-Encoding", "gzip") //使用gzip压缩传输数据让访问更快
	res, err := client.Do(req)

	if err != nil {
		// 请求发生异常
		fmt.Println(err.Error())
	} else {
		defer res.Body.Close() //保证最后关闭Body

		fmt.Println("status code:", res.StatusCode) // 获取状态码

		// 有gzip压缩时,需要解压缩读取返回内容
		if res.Header.Get("Content-Encoding") == "gzip" {
			reader, _ := gzip.NewReader(res.Body) // gzip解压缩
			defer reader.Close()
			io.Copy(os.Stdout, reader)
			os.Exit(0) // 正常退出
		}

		// 无gzip压缩, 读取返回内容
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Println(string(body))
	}
}
