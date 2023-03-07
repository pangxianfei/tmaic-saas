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

// PostLogin 用户名密码登录
func (c *LoginController) PostLogin() *JsonResult {

	var UserLogin requests.UserLogin
	var Admin *UserAppModel.Admin

	if err := c.Ctx.ReadJSON(&UserLogin); err != nil {
		return JsonErrorMsg(err.Error())
	}

	Admin, token, err := services.UserService.SignIn(c.Ctx, UserLogin)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}
	return JsonData(tmaic.V{"Admin": Admin, "token": token})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {
	err := services.UserTokenService.Signout(c.Ctx)
	if err != nil {
		return JsonError(NewErrorMsg("登出失败"))
	}
	return JsonSuccess()
}