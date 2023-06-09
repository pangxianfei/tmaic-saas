package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type PangsController struct {
	controller.BaseController
}

func (pang *PangsController) PostIndex() *response.JsonResult {
	return response.JsonQueryData("查询成功")
}

func (pang *PangsController) PostShow() *response.JsonResult {
	return response.JsonQueryData("查询成功")
}

func (pang *PangsController) PostStore() *response.JsonResult {
	return response.JsonUpdateSuccess("创建成功")
}

func (pang *PangsController) PostUpdate() *response.JsonResult {
	return response.JsonDeleteFail("更新失败，请稍后尝试~")
}

func (pang *PangsController) PostDelete() *response.JsonResult {
	return response.JsonDeleteFail("删除失败，请稍后尝试~")
}
