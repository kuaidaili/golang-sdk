# 简介
快代理api SDK - Golang

# 依赖环境
1. Golang 1.10及以上
2. 从[快代理](https://www.kuaidaili.com)购买相应产品
3. [获取订单的`secret_id`和`secret_key`](https://www.kuaidaili.com/usercenter/api/settings/)

# 获取安装
安装 Golang SDK 前，请先获取订单对应的secret_id和secret_key。 secret_id是用于标识订单，secret_key必须严格保管，避免泄露。

## 通过go get安装(推荐)
您可以通过`go get`将SDK安装到您的`GOPATH`中：
```
go get github.com/kuaidaili/golang-sdk
```

## 通过源码包安装
前往 [Github 代码托管地址](https://github.com/kuaidaili/golang-sdk/tree/master/api-sdk) 下载最新代码，解压后将其放在您的`GOPATH`中

## 示例
以私密代理订单使用为例
``` golang
package examples

import (
	"fmt"
	"log"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/auth"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/client"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
)

// 私密代理使用示例

// 接口鉴权说明：
// 接口鉴权方式为必填项, 目前支持的鉴权方式有"token" 和 "hmacsha1"两种
// 可选值为signtype.TOKEN和signtype.HmacSha1 或直接传"token"或"hmacsha1"

// 返回值说明:
// 所有返回值都包括两个值，第一个为目标值，第二个为error类型, 值为nil说明成功，不为nil说明失败

func useDps() {
	auth := auth.Auth{SecretID: "secret_id", SecretKey: "secret_key"}
	client := client.Client{Auth: auth}

	// 获取订单到期时间, 返回时间字符串
	expireTime, err := client.GetOrderExpireTime(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("expire time: ", expireTime)

	// 获取User-Agent, 返回ua列表
	ua, err := client.GetUA(10 , signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("User-Agent: ", ua)

	// 获取ip白名单, 返回ip切片, 类型为[]string
	ipWhitelist, err := client.GetIPWhitelist(signtype.Token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ip whitelist: ", ipWhitelist)

	//设置ip白名单，参数类型为[]string
	_, err = client.SetIPWhitelist([]string{"test_ip1", "test_ip2"}, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}

	// 提取私密代理, 参数有: 提取数量、鉴权方式及其他参数(放入map[string]interface{}中, 若无则传入nil)
	// (具体有哪些其他参数请参考帮助中心: "https://www.kuaidaili.com/doc/api/getdps/")
	params := map[string]interface{}{"format": "json", "area": "北京,上海"}
	ips, err := client.GetDps(2, signtype.HmacSha1, params)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ips: ", ips)

	// 检测私密代理有效性， 返回map[string]bool, ip:true/false
	valids, err := client.CheckDpsValid(ips, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("valids: ", valids)

	// 获取私密代理剩余时间(单位为秒), 返回map[string]string, ip:seconds
	seconds, err := client.GetDpsValidTime(ips, signtype.Token)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("seconds: ", seconds)

	// 获取计数版ip余额（仅私密代理计数版）
	balance, err := client.GetIPBalance(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("balance: ", balance)

	
	// 获取订单访问代理IP的鉴权信息
	proxyAuthorization, err:= client.GetProxyAuthorization(1,signtype.HmacSha1)
	fmt.Println("proxyAuthorization: ", proxyAuthorization)
}

```
您可以在examples目录下找到更详细的示例

## 参考资料

* [查看API列表](https://www.kuaidaili.com/doc/api/)
* [了解API鉴权](https://www.kuaidaili.com/doc/api/auth/)
