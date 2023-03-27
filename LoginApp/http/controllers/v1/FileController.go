package v1

import (
	"time"

	"gitee.com/pangxianfei/framework/facades"
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
)

type FileController struct {
	controller.BaseController
}

func (c *FileController) PostCreate() *response.JsonResult {

	newFile, err := c.File("avatar")

	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}

	file, _ := facades.Storage.PutFile("pang", newFile)

	url, loadErr := facades.Storage.TemporaryUrl(
		file, time.Now().Add(20*time.Second),
	)
	if loadErr != nil {
		return response.JsonError(loadErr)
	}
	return response.JsonCreateData(url)
}
