package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type FmsController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *FmsController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 添加
func (c *FmsController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostEdit 修改
func (c *FmsController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}
