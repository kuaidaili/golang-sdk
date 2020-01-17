package signtype

// SignType 签名计算方式
type SignType = string

const (
	// SIMPLE 直接以APIKey为签名
	SIMPLE SignType = "simple"
	// HmacSha1 hmac算法计算签名
	HmacSha1 SignType = "hmacsha1"
)
