package v1

import (
	"time"

	"gitee.com/pangxianfei/framework/facades"
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type FileController struct {
	controller.BaseController
}

func (this *FileController) PostCreate() *response.JsonResult {

	newFile, err := this.File("avatar")

	if err != nil {
		return this.JsonErrorMsg(err.Error())
	}

	file, _ := facades.Storage.PutFile("pang", newFile)

	url, loadErr := facades.Storage.TemporaryUrl(
		file, time.Now().Add(20*time.Second),
	)
	if loadErr != nil {
		return this.JsonError(loadErr)
	}
	return this.JsonCreateSucces(url)
}
