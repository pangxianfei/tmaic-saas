package common

import (
	"gitee.com/pangxianfei/library/config"
)

// IsProd 是否是正式环境
func IsProd() bool {
	return config.Instance.Env == "prod"
}
