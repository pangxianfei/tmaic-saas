package v1

import (
	"gitee.com/pangxianfei/framework/facades"
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/framework/kernel/debug"
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

	name := newFile.GetClientOriginalName()
	extension := newFile.GetClientOriginalExtension()
	debug.Dd(name)
	debug.Dd(extension)
	HashName := newFile.HashName()
	debug.Dd(HashName)
	AA, _ := facades.Storage.PutFile("pang", newFile)
	facades.Log.Info(HashName)
	return response.JsonCreateData(AA)
}
