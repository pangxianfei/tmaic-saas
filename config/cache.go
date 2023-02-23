package config

import (
	. "tmaic/vendors/framework/config"
)

func init() {
	cache := make(map[string]interface{})

	cache["default"] = Env("CACHE_DRIVER", "memory")
	cache["cache_time"] = Env("CACHE_TIME", 60)
	cache["stores"] = map[string]interface{}{
		"memory": map[string]interface{}{
			"driver":                    "memory",
			"default_expiration_minute": 5,
			"cleanup_interval_minute":   5,
		},
		"redis": map[string]interface{}{
			"driver":     "redis",
			"connection": "cache",
		},
	}

	Add("cache", cache)
}
