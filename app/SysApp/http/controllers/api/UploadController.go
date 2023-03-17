package api

import (
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

const uploadMaxBytes int64 = 1024 * 1024 * 3

type UploadController struct {
	Ctx iris.Context
}

func (c *UploadController) PostUpload() *simple.JsonResult {
	return simple.JsonSuccess()
}
