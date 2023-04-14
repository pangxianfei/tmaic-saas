package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type WmController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *WmController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 添加
func (c *WmController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostEdit 修改
func (c *WmController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}
