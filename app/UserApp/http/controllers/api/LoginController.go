package api

import (
	"gitee.com/pangxianfei/framework/helpers/tmaic"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	services3 "tmaic/app/UserApp/services"
)

type LoginController struct {
	Ctx iris.Context
}

// PostRegister 注册
func (c *LoginController) PostRegister() *JsonResult {

	var mobile = c.Ctx.PostValueTrim("mobile")
	var password = c.Ctx.PostValueTrim("password")
	var rePassword = c.Ctx.PostValueTrim("rePassword")

	user, err := services3.UserService.SignUp(mobile, password, rePassword)
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
	user, token, err := services3.UserService.SignIn(mobile, password)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	return JsonData(tmaic.V{"user": user, "token": token})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {
	err := services3.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return JsonError(NewErrorMsg("登出失败"))
	}
	return JsonSuccess()
}
