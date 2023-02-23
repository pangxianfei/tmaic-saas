package config

import (
	. "tmaic/vendors/framework/config"
)

func init() {
	app := make(map[string]interface{})

	app["name"] = Env("APP_NAME", "tmaic")
	app["env"] = Env("APP_ENV", "production")
	app["debug"] = Env("APP_DEBUG", false)
	app["log_out"] = Env("LOG_OUT", false)
	app["log_level"] = Env("APP_LOG_LEVEL", "trace")
	app["port"] = Env("APP_PORT", "8080")
	app["orderPort"] = Env("APP_ORDER_PORT", "8081")
	app["productPort"] = Env("APP_PRODUCT_PORT", "8082")
	app["timezone"] = "Asia/Shanghai"
	app["locale"] = Env("APP_LOCALE", "zh")
	app["fallback_locale"] = "zh"
	app["key"] = Env("APP_KEY")
	app["cipher"] = "AES-256-CBC"
	app["read_timeout_seconds"] = 10
	app["write_timeout_seconds"] = 10
	app["storage"] = Env("STORAGE", "storage")
	app["LogFile"] = Env("LOGFILE")

	Add("app", app)
}
