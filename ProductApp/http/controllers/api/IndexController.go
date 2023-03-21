package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/simple"
)

type PrcoductController struct {
	controller.BaseController
}

func (c *PrcoductController) PostInfo() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *PrcoductController) PostCreate() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *PrcoductController) PostEditBy() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostDeleteBy 删除文章
func (c *PrcoductController) PostDeleteBy() *simple.JsonResult {
	return simple.JsonSuccess()
}
