package cache

import (
	"time"
	cache "tmaic/vendors/framework/cache"

	"github.com/golang/protobuf/proto"
)

func Pget(key string, valuePtr proto.Message, defaultValue ...proto.Message) error {
	return cache.Cache().Pget(key, valuePtr, defaultValue...)
}

// Ppull 取出值且删除(只取一次值并且删除，不复存在)
func Ppull(key string, valuePtr proto.Message, defaultValue ...proto.Message) error {
	return cache.Cache().Ppull(key, valuePtr, defaultValue...)
}

// Pput 取出并不删除原始数据
func Pput(key string, value proto.Message, future time.Duration) bool {
	return cache.Cache().Pput(key, value, future)
}
func Padd(key string, value proto.Message, future time.Duration) bool {
	return cache.Cache().Padd(key, value, future)
}
func Pforever(key string, value proto.Message) bool {
	return cache.Cache().Pforever(key, value)
}
