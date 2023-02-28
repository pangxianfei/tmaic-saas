package api

import (
	"github.com/kataras/iris/v12"
	"tmaic/vendors/framework/simple"
)

const uploadMaxBytes int64 = 1024 * 1024 * 3

type UploadController struct {
	Ctx iris.Context
}

func (c *UploadController) Post() *simple.JsonResult {
	return simple.JsonSuccess()
}
