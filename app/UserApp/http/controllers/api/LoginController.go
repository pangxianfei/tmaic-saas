package api

import (
	"gitee.com/pangxianfei/library/tmaic"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"tmaic/app/UserApp/http/requests"
	UserAppModel "tmaic/app/UserApp/model"
	services "tmaic/app/UserApp/services"
)

type LoginController struct {
	Ctx iris.Context
}

// PostRegister 注册
func (c *LoginController) PostRegister() *JsonResult {

	var mobile = c.Ctx.PostValueTrim("mobile")
	var password = c.Ctx.PostValueTrim("password")
	var rePassword = c.Ctx.PostValueTrim("rePassword")

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

	var UserLogin requests.UserLogin
	var User *UserAppModel.User

	if err := c.Ctx.ReadJSON(&UserLogin); err != nil {
		return JsonErrorMsg(err.Error())
	}

	User, token, err := services.UserService.SignIn(UserLogin.Mobile, UserLogin.Password)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	return JsonData(tmaic.V{"user": User, "token": token})

}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {
	err := services.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return JsonError(NewErrorMsg("登出失败"))
	}
	return JsonSuccess()
}
