package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"

	"tmaic/LoginApp/mq"
)

type LoginController struct {
	controller.BaseController
}

// PostLogin 用户名密码登录
func (this *LoginController) PostLogin() *response.JsonResult {
	var UserLogin requests.UserLogin
	if err := this.Context.ReadJSON(&UserLogin); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	adminInfo, token, err := paas.Auth.Login(this.Context, UserLogin)
	if err != nil {
		return this.JsonErrorMsg(err.Error())
	}

	MQ.LoginMessage.Dispatch(adminInfo)
	return this.JsonData(tmaic.V{"token": token, "adminInfo": adminInfo})
}

// PostSignout 退出登录
func (this *LoginController) PostSignout() *response.JsonResult {

	if paas.Auth.Logout(this.Context) {
		return response.JsonSuccess()
	}
	return this.JsonErrorMsg("登陆异常")
}

// PostJosn 退出登录
func (this *LoginController) PostJosn() *response.JsonResult {
	return this.JsonErrorMsg("请求有误")
}
