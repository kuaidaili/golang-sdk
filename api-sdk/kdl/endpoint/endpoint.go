package endpoint

// EndPoint 接口地址
type EndPoint string

const (
	// GetOrderExpireTime 获取订单过期时间
	GetOrderExpireTime EndPoint = "dev.kdlapi.com/api/getorderexpiretime"
	// GetIPWhitelist 获取IP白名单
	GetIPWhitelist EndPoint = "dev.kdlapi.com/api/getipwhitelist"
	// SetIPWhitelist 设置IP白名单
	SetIPWhitelist EndPoint = "dev.kdlapi.com/api/setipwhitelist"
	// GetKpsProxy 获取独享代理
	GetKpsProxy EndPoint = "kps.kdlapi.com/api/getkps"
	// GetDpsProxy 获取私密代理
	GetDpsProxy EndPoint = "dps.kdlapi.com/api/getdps"
	// GetOpsProxyNormalOrVip 获取开放代理普通版和vip版代理
	GetOpsProxyNormalOrVip EndPoint = "dev.kdlapi.com/api/getproxy"
	// GetOpsProxySvip 获取开放代理Svip版
	GetOpsProxySvip EndPoint = "svip.kdlapi.com/api/getproxy"
	// GetOpsProxyEnt 获取开放代理企业版
	GetOpsProxyEnt EndPoint = "ent.kdlapi.com/api/getproxy"
	// CheckDpsValid 验证私密代理是否有效
	CheckDpsValid EndPoint = "dps.kdlapi.com/api/checkdpsvalid"
	// CheckOpsValid 验证开放代理是否有效
	CheckOpsValid EndPoint = "dev.kdlapi.com/api/checkopsvalid"
	// GetIPBalance 获取IP可用余额
	GetIPBalance EndPoint = "dps.kdlapi.com/api/getipbalance"
	// GetDpsValidTime 获取私密代理可用时间
	GetDpsValidTime EndPoint = "dps.kdlapi.com/api/getdpsvalidtime"
	// TpsCurrentIP 获取当前隧道代理IP
	TpsCurrentIP EndPoint = "tps.kdlapi.com/api/tpscurrentip"
	// ChangeTpsIP 更改当前隧道代理IP
	ChangeTpsIP EndPoint = "tps.kdlapi.com/api/changetpsip"
	//GetProxyAuthorization 获取代理鉴权信息
	GetProxyAuthorization EndPoint = "dev.kdlapi.com/api/getproxyauthorization"
	//GetTpsIP 获取隧道代理IP
	GetTpsIp EndPoint = "tps.kdlapi.com/api/gettps"

	//工具接口
	GetUA             EndPoint = "dev.kdlapi.com/api/getua"             //获取User Agent
	GetAreaCode       EndPoint = "dev.kdlapi.com/api/getareacode"       //获取指定地区编码
	GetAccountBalance EndPoint = "dev.kdlapi.com/api/getaccountbalance" //获取账户余额

	// 订单相关
	CreateOrder  EndPoint = "dev.kdlapi.com/api/createorder"  // 创建订单
	GetOrderInfo EndPoint = "dev.kdlapi.com/api/getorderinfo" // 获取订单信息
	SetAutoRenew EndPoint = "dev.kdlapi.com/api/setautorenew" // 开启/关闭自动续费
	CloseOrder   EndPoint = "dev.kdlapi.com/api/closeorder"   // 关闭订单
	QueryKpsCity EndPoint = "dev.kdlapi.com/api/querykpscity" // 查询独享代理城市信息

	GetSecretToken EndPoint = "auth.kdlapi.com/api/get_secret_token" // 获取密钥令牌
)
