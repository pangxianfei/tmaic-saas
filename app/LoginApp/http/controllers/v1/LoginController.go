package v1

import (
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type LoginController struct {
	Ctx iris.Context
}

// PostLogin 用户名密码登录
func (c *LoginController) PostLogin() *JsonResult {
	var UserLogin requests.UserLogin
	if err := ReadJSON(c.Ctx, &UserLogin); err != nil {
		return JsonErrorMsg(err.Error())
	}

	adminInfo, token, err := paas.Auth.Login(c.Ctx, UserLogin)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}
	return JsonData(tmaic.V{"token": token, "adminInfo": adminInfo})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *JsonResult {

	if paas.Auth.Logout(c.Ctx) {
		return JsonSuccess()
	}
	return JsonError(NewErrorMsg("已登出"))

}
