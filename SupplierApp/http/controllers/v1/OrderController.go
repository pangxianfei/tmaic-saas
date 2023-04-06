package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
)

type OrderController struct {
	controller.BaseController
}

// PostList 获取当前登录用户
func (this *OrderController) PostList() *response.JsonResult {
	return this.JsonQueryData(nil)
}

// PostUpdate 更新
func (this *OrderController) PostUpdate() *response.JsonResult {
	return this.JsonUpdateSuccess(nil)
}

// PostDelete 更新
func (this *OrderController) PostDelete() *response.JsonResult {
	return this.JsonDeleteSuccess(nil)
}
