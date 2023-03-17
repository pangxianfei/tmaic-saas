package v1

import (
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type CreateTableController struct {
	Ctx iris.Context
}

func (c *CreateTableController) PostCreate() *simple.JsonResult {

	return simple.JsonSuccess()
}
