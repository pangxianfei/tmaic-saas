package lang

import (
	"gitee.com/pangxianfei/framework/kernel/locale"

	"gitee.com/pangxianfei/framework/resources/lang"
)

func init() {
	validationTranslation := lang.ValidationTranslation{}

	customTranslation := lang.CustomTranslation{
		"auth.register.failed_existed":              "user register failed, user has been registered before",
		"auth.register.failed_system_error":         "user register failed, system error",
		"auth.register.failed_token_generate_error": "用户注册失败，令牌生成失败",

		"auth.login.failed_not_exist":            "user login failed, user doesn't exist",
		"auth.login.failed_wrong_password":       "user login failed, user password incorrect",
		"auth.login.failed_token_generate_error": "用户登录失败，令牌生成失败",

		//URL地址不存在
		"address.not.exist": "URL地址不存在",
	}
	locale.AddLocale("zh", &customTranslation, &validationTranslation)
}
