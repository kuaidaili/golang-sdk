# 快代理SDK - Golang

## 快代理API-SDK - Go

[快代理GO-API-SDK](https://github.com/kuaidaili/golang-sdk/tree/master/go-api-sdk)

## 调用API
### api.go
使用标准库调用api链接并解析返回内容示例

## Http代理
### http.go
使用标准库请求Http代理服务器, 请求http和https网页均适用

## Socks代理
### socks_auth.go
以`用户名密码`认证形式使用`golang.org/x/net`包请求Socks代理服务器, 请求http和https网页均适用

### socks_whitelist.go
以`白名单`认证形式使用`golang.org/x/net`包请求Socks代理服务器, 请求http和https网页均适用
```
使用提示：
请先获取官方net包: go get golang.org/x/net
```

## 隧道代理
### tps.go
以`用户名密码`使用标准库请求Http代理服务器, 请求http和https网页均适用


## 技术支持

如果您发现代码有任何问题, 请提交`Issue`。

欢迎提交`Pull request`以使代码样例更加完善。

获取更多关于调用API和代理服务器使用的资料，请参考[开发者指南](https://help.kuaidaili.com/dev/api/)。

* 技术支持微信：<a href="https://img.kuaidaili.com/img/service_wx.jpg">kuaidaili</a>
* 技术支持QQ：<a href="http://q.url.cn/CDksXo?_type=wpa&qidian=true">800849628</a>
