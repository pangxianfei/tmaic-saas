package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type OverviewController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *OverviewController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 创建
func (c *OverviewController) PostStore() *response.JsonResult {

	return response.JsonCreateSucces("")
}

// PostEdit 修改用户资料
func (c *OverviewController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}

// PostUpdate 更新
func (c *OverviewController) PostUpdate() *response.JsonResult {
	return response.JsonSuccess()
}
