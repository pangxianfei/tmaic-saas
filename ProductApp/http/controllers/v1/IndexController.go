package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
)

type PrcoductController struct {
	controller.BaseController
}

func (c *PrcoductController) PostInfo() *response.JsonResult {
	return response.JsonSuccess()
}

func (c *PrcoductController) PostCreate() *response.JsonResult {
	return response.JsonSuccess()
}

func (c *PrcoductController) PostEditBy() *response.JsonResult {
	return response.JsonSuccess()
}

// PostDeleteBy 删除文章
func (c *PrcoductController) PostDeleteBy() *response.JsonResult {
	return response.JsonSuccess()
}
