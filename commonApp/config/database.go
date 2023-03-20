package config

import (
	. "gitee.com/pangxianfei/library/config"
)

func init() {
	database := make(map[string]interface{})

	database["default"] = Env("DB_CONNECTION", "mysql")
	database["db_prefix"] = Env("DB_PREFIX", "tmaic_")
	database["show-sql"] = Env("SHOW_SQL", false)
	database["dns"] = Env("DNS")

	database["connections"] = map[string]interface{}{
		"mysql": map[string]interface{}{
			"driver":    "mysql",
			"host":      Env("DB_HOST", "127.0.0.1"),
			"port":      Env("DB_PORT", "3306"),
			"database":  Env("DB_DATABASE", ""),
			"username":  Env("DB_USERNAME", ""),
			"password":  Env("DB_PASSWORD", ""),
			"charset":   "utf8mb4",
			"collation": "utf8mb4_unicode_ci",
			"prefix":    Env("DB_PREFIX", ""),
		},
		"logs": map[string]interface{}{
			"driver":    "mysql",
			"host":      Env("LOGS_DB_HOST", "127.0.0.1"),
			"port":      Env("LOGS_DB_PORT", "3306"),
			"database":  Env("LOGS_DB_DATABASE", ""),
			"username":  Env("LOGS_DB_USERNAME", ""),
			"password":  Env("LOGS_DB_PASSWORD", ""),
			"charset":   "utf8mb4",
			"collation": "utf8mb4_unicode_ci",
			"prefix":    Env("LOGS_DB_PREFIX", ""),
		},
	}
	database["migrations"] = "migrations"
	database["max_idle_connections"] = Env("DB_MAX_IDLE_CONNECTIONS", 2) // 2 is the cpu cores
	database["max_open_connections"] = Env("DB_MAX_OPEN_CONNECTIONS", 0) // 2 is the cpu cores
	database["max_life_seconds"] = Env("DB_MAX_LIFE_SECONDS", 0)         // 2 is the cpu cores

	database["redis"] = map[string]interface{}{

		"options": map[string]interface{}{
			"prefix": Env("CACHE_PREFIX", ""),
		},
		"default": map[string]interface{}{
			"host":     Env("REDIS_HOST", "127.0.0.1"),
			"port":     Env("REDIS_PORT", "6379"),
			"password": Env("REDIS_PASSWORD", ""),
			"database": Env("REDIS_DB", 0),
		},
		"cache": map[string]interface{}{
			"host":     Env("REDIS_HOST", "127.0.0.1"),
			"port":     Env("REDIS_PORT", "6379"),
			"password": Env("REDIS_PASSWORD", ""),
			"database": Env("REDIS_CACHE_DB", 1),
		},
	}

	Add("database", database)
}
