package config

import (
	. "gitee.com/pangxianfei/library/config"
	"tmaic/app/model"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "f6P8FPkpNbWmkeRr")
	auth["model_ptr"] = &model.User{}

	Add("auth", auth)
}
