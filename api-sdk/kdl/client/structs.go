package client

import "github.com/kuaidaili/golang-sdk/api-sdk/kdl/auth"

// SecretToken信息存放路径
const SecretPath = ".secret"

// Client 请求客户端
type Client struct {
	Auth auth.Auth
}

// jsonRes 返回结果的json解析结构体
type jsonRes struct {
	Msg  string
	Code int
	Data interface{}
}
