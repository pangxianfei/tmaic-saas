package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type FinanceController struct {
	controller.BaseController
}

// PostList 列表
func (c *FinanceController) PostList() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostInfo 获取当前登录用户
func (c *FinanceController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 添加
func (c *FinanceController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostEdit 修改
func (c *FinanceController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}
