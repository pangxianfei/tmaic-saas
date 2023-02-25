package api

import (
	"github.com/kataras/iris/v12"
	"tmaic/app/services"
	"tmaic/vendors/framework/helpers/tmaic"
	. "tmaic/vendors/framework/simple"
)

type LoginController struct {
	Ctx iris.Context
}

// PostRegister 注册
func (c *LoginController) PostRegister() *JsonResult {
	var (
		//email      = c.Ctx.PostValueTrim("email")
		//username   = c.Ctx.PostValueTrim("username")
		mobile     = c.Ctx.PostValueTrim("mobile")
		password   = c.Ctx.PostValueTrim("password")
		rePassword = c.Ctx.PostValueTrim("rePassword")
		//nickname   = c.Ctx.PostValueTrim("nickname")
	)
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
	var (
		mobile   = c.Ctx.PostValueTrim("mobile")
		password = c.Ctx.PostValueTrim("password")
	)
	user, token, err := services.UserService.SignIn(mobile, password)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	return JsonData(tmaic.V{"user": user, "token": token})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {
	return JsonSuccess()
}
