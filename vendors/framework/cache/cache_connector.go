package cache

import (
	"tmaic/vendors/framework/cache/driver/memory"
	"tmaic/vendors/framework/cache/driver/redis"
	"tmaic/vendors/framework/config"
)

var cer Cacher

func Initialize() {
	cer = setStore("default")
}
func setStore(store string) (cer Cacher) {

	_conn := store
	if store == "default" {
		_conn = config.GetString("cache." + store)
		if _conn == "" {
			panic("cache connection parse error")
		}
	}
	//debug.Dump(_conn)
	// get driver instance and connect cache store
	switch _conn {
	case "memory":
		cer = memory.NewMemory(
			config.GetString("cache.stores.memory.prefix"),
			config.GetUint("cache.stores.memory.default_expiration_minute"),
			config.GetUint("cache.stores.memory.cleanup_interval_minute"),
		)
	case "redis":
		connection := config.GetString("cache.stores.redis.connection") // cache
		cer = redis.NewRedis(
			config.GetString("database.redis."+connection+".host"),
			config.GetString("database.redis."+connection+".port"),
			config.GetString("database.redis."+connection+".password"),
			config.GetInt("database.redis."+connection+".database"),
			config.GetString("database.redis.options.prefix"),
		)
	default:
		panic("incorrect cache connection provided")
	}

	return cer
}

func Store(store string) Cacher {
	return setStore(store)
}

func Cache() Cacher {
	return cer
}
