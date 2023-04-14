package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type CreateTableController struct {
	controller.BaseController
}

func (c *CreateTableController) PostCreate() *response.JsonResult {
	return response.JsonSuccess()
}
