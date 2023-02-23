package cache

import (
	"github.com/goburrow/cache"
)

// 转化int64
func key2Int64(key cache.Key) int64 {
	return key.(int64)
}
