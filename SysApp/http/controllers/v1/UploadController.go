package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

const uploadMaxBytes int64 = 1024 * 1024 * 3

type UploadController struct {
	controller.BaseController
}

func (c *UploadController) PostUpload() *response.JsonResult {
	return response.JsonSuccess()
}
