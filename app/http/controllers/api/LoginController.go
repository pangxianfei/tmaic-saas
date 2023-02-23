package api

import (
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"tmaic/app/services"
	"tmaic/vendors/framework/helpers/tmaic"
)

type LoginController struct {
	Ctx iris.Context
}

// PostRegister 注册
func (c *LoginController) PostRegister() *JsonResult {
	return JsonSuccess()
}

// PostLogin 用户名密码登录
func (c *LoginController) PostLogin() *JsonResult {
	var (
		username = c.Ctx.PostValueTrim("username")
		password = c.Ctx.PostValueTrim("password")
	)
	user, token, err := services.UserService.SignIn(username, password)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	return JsonData(tmaic.V{"user": user, "token": token})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {
	return JsonSuccess()
}
