package utils

import (
	"fmt"
	"strconv"
)

// TypeSwitcher 空接口转换为string
func TypeSwitcher(t interface{}) string {
	switch v := t.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	case int64:
		return strconv.Itoa(int(v))
	case float64:
		return fmt.Sprintf("%.0f", v)
	default:
		return ""
	}
}
