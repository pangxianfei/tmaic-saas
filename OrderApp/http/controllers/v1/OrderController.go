package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type OverviewController struct {
	controller.BaseController
}

func (this *OverviewController) PostList() *response.JsonResult {
	return this.JsonSuccess()
}

func (this *OverviewController) PostInfo() *response.JsonResult {
	return this.JsonSuccess()
}

func (this *OverviewController) PostCreate() *response.JsonResult {
	return this.JsonSuccess()
}

func (this *OverviewController) PostEdit() *response.JsonResult {
	return this.JsonSuccess()
}

// PostDelete 删除
func (this *OverviewController) PostDelete() *response.JsonResult {
	return this.JsonSuccess()
}
