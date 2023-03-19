package api

import (
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type OrderController struct {
	Ctx iris.Context
}

func (c *OrderController) PostInfo() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) PostCreate() *simple.JsonResult {
	return simple.JsonSuccess()
}

func (c *OrderController) PostEditBy() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostDeleteBy 删除文章
func (c *OrderController) PostDeleteBy() *simple.JsonResult {
	return simple.JsonSuccess()
}
