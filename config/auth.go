package config

import (
	"tmaic/app/model"
	. "tmaic/vendors/framework/config"
)

func init() {
	auth := make(map[string]interface{})

	auth["sign_key"] = Env("AUTH_SIGN_KEY", "f6P8FPkpNbWmkeRr")
	auth["model_ptr"] = &model.User{}

	Add("auth", auth)
}
