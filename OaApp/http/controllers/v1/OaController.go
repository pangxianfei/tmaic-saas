package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type OaController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *OaController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}
