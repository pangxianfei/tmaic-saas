package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type PrcoductController struct {
	controller.BaseController
}

func (this *PrcoductController) PostInfo() *response.JsonResult {
	return this.JsonSuccess()
}

func (this *PrcoductController) PostStore() *response.JsonResult {
	return this.JsonCreateSucces(nil)
}

func (this *PrcoductController) PostEdit() *response.JsonResult {
	return this.JsonSuccess()
}

func (this *PrcoductController) PostUpdate() *response.JsonResult {
	return this.JsonSuccess()
}

// PostDelete 删除
func (this *PrcoductController) PostDelete() *response.JsonResult {
	return this.JsonSuccess()
}
