package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/framework/work"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"

	"tmaic/app/jobs"
	"tmaic/app/jobs/proto3/protomodel"
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

	//推送登陆消息队列
	LoginJob := jobs.LoginJob
	LoginJob.SetParam(&protomodel.LoginJob{
		UserName: adminInfo.UserName,
		Mobile:   adminInfo.Mobile,
		TenantId: adminInfo.TenantId,
		UserType: int32(adminInfo.UserType),
	})
	if jobErr := work.Dispatch(LoginJob); jobErr != nil {
		return response.JsonErrorMsg(jobErr.Error())
	}

	return response.JsonData(tmaic.V{"token": token, "adminInfo": adminInfo})
}

// GetSignout 退出登录
func (c *LoginController) GetSignout() *response.JsonResult {

	if paas.Auth.Logout(c.Context) {
		return response.JsonSuccess()
	}
	return response.JsonErrorMsg("登陆异常")
}
