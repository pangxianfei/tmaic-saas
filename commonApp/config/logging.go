package config

import (
	"gitee.com/pangxianfei/library/config"
)

func init() {

	config.Add("logging", map[string]any{
		// Default Log Channel
		//
		// This option defines the default log channel that gets used when writing
		// messages to the logs. The name specified in this option should match
		// one of the channels defined in the "channels" configuration array.
		"default": config.Env("LOG_CHANNEL", "stack"),

		// Log Channels
		//
		// Here you may configure the log channels for your application.
		// Available Drivers: "single", "daily", "custom", "stack"
		// Available Level: "debug", "info", "warning", "error", "fatal", "panic"
		"channels": map[string]any{
			"stack": map[string]any{
				"driver":   "stack",
				"channels": []string{"daily"},
			},
			"single": map[string]any{
				"driver": "single",
				"path":   "./commonApp/storage/logs/saas.log",
				"level":  config.Env("LOG_LEVEL", "debug"),
			},
			"daily": map[string]any{
				"driver": "daily",
				"path":   "./commonApp/storage/logs/saas.log",
				"level":  config.Env("LOG_LEVEL", "debug"),
				"days":   7,
			},
		},
	})
}
