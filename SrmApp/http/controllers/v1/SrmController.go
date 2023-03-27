package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type SrmController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *SrmController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 添加
func (c *SrmController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostEdit 修改
func (c *SrmController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}
