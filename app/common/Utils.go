package common

import (
	"github.com/goburrow/cache"
	"tmaic/config"
)

// IsProd 是否是正式环境
func IsProd() bool {
	return config.Instance.Env == "prod"
}

// 转化int64
func key2Int64(key cache.Key) int64 {
	return key.(int64)
}
