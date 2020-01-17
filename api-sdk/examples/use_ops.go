package examples

import (
	"fmt"
	"log"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/auth"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/client"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/opslevel"
	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/signtype"
)

// 开放代理使用示例

// 接口鉴权说明：
// 接口鉴权方式为必填项, 目前支持的鉴权方式有"simple" 和 "hmacsha1"两种
// 可选值为signtype.SIMPLE和signtype.HmacSha1 或直接传"simple"或"hmacsha1"

// 返回值说明:
// 所有返回值都包括两个值，第一个为目标值，第二个为error类型, 值为nil说明成功，不为nil说明失败

func useOps() {
	auth := auth.Auth{OrderID: "test_order_id", APIKey: "test_api_key"}
	client := client.Client{Auth: auth}

	// 获取订单到期时间, 返回时间字符串
	expireTime, err := client.GetOrderExpireTime(signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("expire time: ", expireTime)

	// 提取开放代理, 参数有: 提取数量、开放代理套餐种类(normal, vip, svip, ent
	// 	分别对应opslevel.NORMAL, opslevel.VIP, opslevel.SVIP, opslevel.ENT)、
	// 鉴权方式及其他参数(放入map[string]interface{}中, 若无则传入nil)
	// (具体有哪些其他参数请参考帮助中心: "https://www.kuaidaili.com/doc/api/getops/")
	params := map[string]interface{}{"area": "北京,上海"}
	ips, err := client.GetProxy(2, opslevel.NORMAL, signtype.HmacSha1, params)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("ips: ", ips)

	// 检测开放代理有效性， 返回map[string]bool, ip:true/false
	valids, err := client.CheckOpsValid(ips, signtype.HmacSha1)
	if err != nil {
		log.Println(err)
	}
	fmt.Println("valids: ", valids)
}
