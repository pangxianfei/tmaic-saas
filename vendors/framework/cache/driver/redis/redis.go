package redis

import (
	"errors"
	r "github.com/go-redis/redis"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	"time"

	. "tmaic/vendors/framework/cache/utils"
)

func NewRedis(host string, port string, password string, dbIndex int, prefix string) *redis {
	client := r.NewClient(&r.Options{
		Addr:     host + ":" + port,
		Password: password,
		DB:       dbIndex,
	})

	return &redis{
		redisBasic{
			cache:  client,
			prefix: prefix,
		},
	}
}

type redis struct {
	redisBasic
}

func (re *redis) SetNx(key string, value string, future time.Duration) bool {
	k := NewKey(key, re.Prefix())

	_, err := re.cache.Set(k.Prefixed(), value, future).Result()

	if err != nil {
		return false
	}
	return true
}

func (re *redis) GetString(key string) string {
	k := NewKey(key, re.Prefix())

	valStr, err := re.cache.Get(k.Prefixed()).Result()
	if err != nil {
		return ""
	}
	return valStr
}

func (re *redis) Pget(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error {
	var valueBytes []byte

	k := NewKey(key, re.Prefix())

	if !re.Has(k.Raw()) {

		if len(defaultValuePtr) > 0 {

			return copier.Copy(valuePtr, defaultValuePtr[0])
		}
		return errors.New("key not exist")
	}
	err := re.cache.Get(k.Prefixed()).Scan(&valueBytes)
	if err != nil {

		return err
	}

	if err := proto.Unmarshal(valueBytes, valuePtr); err != nil {

		return err
	}

	return nil
}

// ------------------------------------------------------------------------------
// the same
// ------------------------------------------------------------------------------

func (re *redis) Ppull(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error {
	k := NewKey(key, re.Prefix())

	err := re.Pget(k.Raw(), valuePtr, defaultValuePtr...)
	if err != nil {
		return err
	}

	re.Forget(k.Raw())

	return nil
}

func (re *redis) Pput(key string, valuePtr proto.Message, future time.Duration) bool {
	valueBytes, err := proto.Marshal(valuePtr)
	if err != nil {
		return false
	}
	return re.Put(key, valueBytes, future)
}

func (re *redis) Padd(key string, valuePtr proto.Message, future time.Duration) bool {
	valueBytes, err := proto.Marshal(valuePtr)
	if err != nil {
		return false
	}
	return re.Add(key, valueBytes, future)
}
func (re *redis) Pforever(key string, valuePtr proto.Message) bool {
	valueBytes, err := proto.Marshal(valuePtr)
	if err != nil {
		return false
	}
	return re.Forever(key, valueBytes)
}
