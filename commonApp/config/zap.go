package config

import (
	"gitee.com/pangxianfei/library/config"
)

func init() {

	config.Add("zap", map[string]any{
		"filename":   config.Env("ZAP_FILENAME", "./commonApp/storage/logs/zap.log"),
		"max_size":   config.Env("ZAP_MAX_SIZE", 128),
		"max_backup": config.Env("ZAP_MAX_BACKUP", 30),
		"max_age":    config.Env("ZAP_AGE", 7),
		"compress":   config.Env("ZAP_COMPRESS", true),
	})
}
