package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"sort"
	"strings"

	"github.com/kuaidaili/golang-sdk/api-sdk/kdl/utils"
)

// Auth 用于保存用户SecretID、SecretKey，以及计算签名
type Auth struct {
	SecretID  string
	SecretKey string
}

// GetStringToSign 生成签名原文字符串
func (auth Auth) GetStringToSign(method string, endpoint string, params map[string]interface{}) string {
	s := method + strings.Split(endpoint, ".com")[1] + "?"
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var arr []string
	for _, k := range keys {
		arr = append(arr, k+"="+utils.TypeSwitcher(params[k]))
	}
	queryStr := strings.Join(arr, "&")
	return s + queryStr
}

// SignStr 计算签名串
func (auth Auth) SignStr(rawStr string) string {
	key := []byte(auth.SecretKey)
	hash := hmac.New(sha1.New, key)
	hash.Write([]byte(rawStr))
	sig := base64.StdEncoding.EncodeToString([]byte(string(hash.Sum(nil))))
	return sig
}
