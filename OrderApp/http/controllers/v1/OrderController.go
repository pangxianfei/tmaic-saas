package v1

import (
	"gitee.com/pangxianfei/library/response"
	"github.com/kataras/iris/v12"
)

type OrderController struct {
	Ctx iris.Context
}

func (c *OrderController) PostInfo() *response.JsonResult {
	return response.JsonSuccess()
}

func (c *OrderController) PostCreate() *response.JsonResult {
	return response.JsonSuccess()
}

func (c *OrderController) PostEditBy() *response.JsonResult {
	return response.JsonSuccess()
}

// PostDeleteBy 删除文章
func (c *OrderController) PostDeleteBy() *response.JsonResult {
	return response.JsonSuccess()
}
