package api

import (
	"tmaic/app/services"

	"gitee.com/pangxianfei/framework/helpers/tmaic"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

// PostRegister 注册
func (c *LoginController) PostRegister() *JsonResult {

	var mobile = c.Ctx.PostValueTrim("mobile")
	var password = c.Ctx.PostValueTrim("password")
	var rePassword = c.Ctx.PostValueTrim("rePassword")

	loginMethod := services.SysConfigService.GetLoginMethod()
	if !loginMethod.Password {
		return JsonErrorMsg("账号密码登录/注册已禁用")
	}

	user, err := services.UserService.SignUp(mobile, password, rePassword)
	if err != nil {
		return JsonError(&CodeError{
			Code:    401,
			Message: err.Error(),
			Data:    nil,
		})
	}
	return JsonData(user)
}

// PostLogin 用户名密码登录
func (c *LoginController) PostLogin() *JsonResult {

	var mobile = c.Ctx.PostValueTrim("mobile")
	var password = c.Ctx.PostValueTrim("password")
	user, token, err := services.UserService.SignIn(mobile, password)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	return JsonData(tmaic.V{"user": user, "token": token})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {
	err := services.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return JsonError(NewErrorMsg("登出失败"))
	}
	return JsonSuccess()
}
