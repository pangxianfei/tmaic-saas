package config

import (
	. "gitee.com/pangxianfei/library/config"
	"tmaic/app/UserApp/model"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "f6P8FPkpNbWmkeRr")
	auth["model_ptr"] = &UserAppModel.TenantUser{}
	Add("auth", auth)
}
