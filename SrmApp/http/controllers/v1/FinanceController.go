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

// PostStore 添加
func (c *FinanceController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostUpdate 更新
func (this *FinanceController) PostUpdate() *response.JsonResult {
	return this.JsonUpdateSuccess(nil)
}

// PostDelete 更新
func (this *FinanceController) PostDelete() *response.JsonResult {
	return this.JsonDeleteSuccess(nil)
}