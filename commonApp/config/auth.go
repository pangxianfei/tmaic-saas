package config

import (
	. "gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/sysmodel"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "f6P8FPkpNbWmkeRr")
	auth["model_ptr"] = &sysmdel.TenantUser{}
	Add("auth", auth)
}
