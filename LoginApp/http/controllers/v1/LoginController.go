package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"

	MQ "tmaic/LoginApp/mq"
)

type LoginController struct {
	controller.BaseController
}

// PostLogin 用户名密码登录
func (c *LoginController) PostLogin() *response.JsonResult {
	var UserLogin requests.UserLogin
	if err := c.Context.ReadJSON(&UserLogin); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	adminInfo, token, err := paas.Auth.Login(c.Context, UserLogin)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}

	MQ.LoginMessage.Dispatch(adminInfo)
	return response.JsonData(tmaic.V{"token": token, "adminInfo": adminInfo})
}

// PostSignout 退出登录
func (c *LoginController) PostSignout() *response.JsonResult {

	if paas.Auth.Logout(c.Context) {
		return response.JsonSuccess()
	}
	return response.JsonErrorMsg("登陆异常")
}
