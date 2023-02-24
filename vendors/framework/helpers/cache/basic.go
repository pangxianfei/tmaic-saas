package cache

import (
	c "tmaic/vendors/framework/cache"
	"tmaic/vendors/framework/config"
	//"tmaic/vendors/framework/helpers/zone"
	"time"
)

func Prefix() string {
	return c.Cache().Prefix()
}

func Has(key string) bool {
	return c.Cache().Has(key)
}

func setCacheTime() (Duration time.Duration) {
	var TIMELENGTH = time.Duration(config.GetInt("cache.cache_time"))
	Duration = TIMELENGTH * time.Second
	return
}
func setTokenTime() (TokenTime time.Duration) {
	var TokenLoogTime = time.Duration(config.GetInt("cache.token_time"))
	TokenTime = TokenLoogTime * time.Second
	return
}

func GetString(key string) string {
	return c.Cache().GetString(key)
}

func Get(key string, defaultValue ...interface{}) interface{} {
	return c.Cache().Get(key, defaultValue...)
}
func Pull(key string, defaultValue ...interface{}) interface{} {
	return c.Cache().Pull(key, defaultValue...)
}
func Put(key string, value interface{}) bool {
	return c.Cache().Put(key, value, setCacheTime())
}

func AddTokenCache(key string, value interface{}) bool {
	return c.Cache().Put(key, value, setTokenTime())
}

func SetNx(key string, value string) bool {
	return c.Cache().SetNx(key, value, setCacheTime())
}

func Add(key string, value interface{}) bool {
	return c.Cache().Add(key, value, setCacheTime())
}
func Increment(key string, value int64) (incremented int64, success bool) {
	return c.Cache().Increment(key, value)
}
func Decrement(key string, value int64) (decremented int64, success bool) {
	return c.Cache().Decrement(key, value)
}
func Forever(key string, value interface{}) bool {
	return c.Cache().Forever(key, value)
}
func Forget(key string) bool {
	return c.Cache().Forget(key)
}
