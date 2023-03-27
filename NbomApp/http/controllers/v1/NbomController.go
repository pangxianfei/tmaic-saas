package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type NbomController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *NbomController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 添加
func (c *NbomController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostEdit 修改
func (c *NbomController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}
