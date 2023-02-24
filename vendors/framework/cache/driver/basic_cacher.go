package driver

import "time"

type BasicCache interface {
	Prefix() string

	Has(key string) bool
	Get(key string, defaultValue ...interface{}) interface{}
	Pull(key string, defaultValue ...interface{}) interface{}
	Put(key string, value interface{}, future time.Duration) bool
	Add(key string, value interface{}, future time.Duration) bool
	Increment(key string, value int64) (incremented int64, success bool)
	Decrement(key string, value int64) (decremented int64, success bool)
	Forever(key string, value interface{}) bool
	Forget(key string) bool
	SetNx(key string, value string, future time.Duration) bool
	GetString(key string) string
	Close() error
}
