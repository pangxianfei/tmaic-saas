package config

import (
	"gitee.com/pangxianfei/library/config"
)

func init() {
	config.Add("filesystems", map[string]any{
		// Default Filesystem Disk
		//
		// Here you may specify the default filesystem disk that should be used
		// by the framework. The "local" disk, as well as a variety of cloud
		// based disks are available to your application. Just store away!
		"default": config.Env("FILESYSTEM_DISK", "local"),

		// Filesystem Disks
		//
		// Here you may configure as many filesystem "disks" as you wish, and you
		// may even configure multiple disks of the same driver. Defaults have
		// been set up for each driver as an example of the required values.
		//
		// Supported Drivers: "local", "s3", "oss", "cos", "custom"
		"disks": map[string]any{
			"local": map[string]any{
				"driver": "local",
				"root":   "./commonApp/storage/app",
				"url":    config.Env("APP_URL").(string) + "/storage",
			},
			"s3": map[string]any{
				"driver": "s3",
				"key":    config.Env("AWS_ACCESS_KEY_ID"),
				"secret": config.Env("AWS_ACCESS_KEY_SECRET"),
				"region": config.Env("AWS_DEFAULT_REGION"),
				"bucket": config.Env("AWS_BUCKET"),
				"url":    config.Env("AWS_URL"),
			},
			"oss": map[string]any{
				"driver":   "oss",
				"key":      config.Env("ALIYUN_ACCESS_KEY_ID"),
				"secret":   config.Env("ALIYUN_ACCESS_KEY_SECRET"),
				"bucket":   config.Env("ALIYUN_BUCKET"),
				"url":      config.Env("ALIYUN_URL"),
				"endpoint": config.Env("ALIYUN_ENDPOINT"),
			},
			"cos": map[string]any{
				"driver": "cos",
				"key":    config.Env("TENCENT_ACCESS_KEY_ID"),
				"secret": config.Env("TENCENT_ACCESS_KEY_SECRET"),
				"bucket": config.Env("TENCENT_BUCKET"),
				"url":    config.Env("TENCENT_URL"),
			},
		},
	})
}

/*
	filesystems := make(map[string]interface{})

	filesystems["filesystems"] = map[string]any{
		// Default Filesystem Disk
		//
		// Here you may specify the default filesystem disk that should be used
		// by the framework. The "local" disk, as well as a variety of cloud
		// based disks are available to your application. Just store away!
		"default": Env("FILESYSTEM_DISK", "local"),

		// Filesystem Disks
		//
		// Here you may configure as many filesystem "disks" as you wish, and you
		// may even configure multiple disks of the same driver. Defaults have
		// been set up for each driver as an example of the required values.
		//
		// Supported Drivers: "local", "s3", "oss", "cos", "custom"
		"disks": map[string]any{
			"local": map[string]any{
				"driver": "local",
				"root":   "storage/app",
				"url":    Env("APP_URL").(string) + "/storage",
			},
			"s3": map[string]any{
				"driver": "s3",
				"key":    Env("AWS_ACCESS_KEY_ID"),
				"secret": Env("AWS_ACCESS_KEY_SECRET"),
				"region": Env("AWS_DEFAULT_REGION"),
				"bucket": Env("AWS_BUCKET"),
				"url":    Env("AWS_URL"),
			},
			"oss": map[string]any{
				"driver":   "oss",
				"key":      Env("ALIYUN_ACCESS_KEY_ID"),
				"secret":   Env("ALIYUN_ACCESS_KEY_SECRET"),
				"bucket":   Env("ALIYUN_BUCKET"),
				"url":      Env("ALIYUN_URL"),
				"endpoint": Env("ALIYUN_ENDPOINT"),
			},
			"cos": map[string]any{
				"driver": "cos",
				"key":    Env("TENCENT_ACCESS_KEY_ID"),
				"secret": Env("TENCENT_ACCESS_KEY_SECRET"),
				"bucket": Env("TENCENT_BUCKET"),
				"url":    Env("TENCENT_URL"),
			},
		},
	}
	Add("filesystems", filesystems)
}
*/
