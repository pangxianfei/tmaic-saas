package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type OverviewController struct {
	controller.BaseController
}

func (this *OverviewController) PostIndex() *response.JsonResult {
	return this.JsonSuccess()
}

// PostStore 添加
func (this *OverviewController) PostStore() *response.JsonResult {
	return this.JsonSuccess()
}

// PostUpdate 更新
func (this *OverviewController) PostUpdate() *response.JsonResult {
	return this.JsonUpdateSuccess(nil)
}

// PostDelete 更新
func (this *OverviewController) PostDelete() *response.JsonResult {
	return this.JsonDeleteSuccess(nil)
}
