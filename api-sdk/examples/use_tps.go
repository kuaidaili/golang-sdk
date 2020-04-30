package examples

import (
	"fmt"
	"log"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/auth"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/client"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
)

// 隧道代理使用示例

// 接口鉴权说明：
// 接口鉴权方式为必填项, 目前支持的鉴权方式有"simple" 和 "hmacsha1"两种
// 可选值为signtype.SIMPLE和signtype.HmacSha1 或直接传"simple"或"hmacsha1"

// 返回值说明:
// 所有返回值都包括两个值，第一个为目标值，第二个为error类型, 值为nil说明成功，不为nil说明失败

func useTps() {
	auth := auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
	client := client.Client{Auth: auth}

	// 获取订单到期时间, 返回时间字符串
	expireTime, err := client.GetOrderExpireTime(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("expire time: ", expireTime)

	// 获取ip白名单, 返回ip切片, 类型为[]string
	ipWhitelist, err := client.GetIPWhitelist(signtype.SIMPLE)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ip whitelist: ", ipWhitelist)

	//设置ip白名单，参数类型为[]string
	_, err = client.SetIPWhitelist([]string{"test_ip1", "test_ip2"}, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}

	// 获取隧道代理当前的ip
	ip, err := client.TpsCurrentIP(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("current_ip: ", ip)

	// 改变当前隧道ip
	newIP, err := client.ChangeTpsIP(signtype.SIMPLE)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("new_ip: ", newIP)

	
	// 获取订单访问代理IP的鉴权信息
	proxyAuthorization, err:= client.GetProxyAuthorization(1,signtype.HmacSha1)
	fmt.Println("proxyAuthorization: ", proxyAuthorization)


	//提取隧道代理IP, 参数有: 提取数量、鉴权方式及其他参数(放入map[string]interface{}中, 若无则传入nil)
	//(具体有哪些其他参数请参考帮助中心: "https://www.kuaidaili.com/doc/api/gettps/")
	params := map[string]interface{}{"format": "json"}
	ips, err := client.GetTpsIp(2, signtype.HmacSha1, params)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ips: ", ips)
}
