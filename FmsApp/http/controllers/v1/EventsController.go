package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/framework/work"
	"gitee.com/pangxianfei/library/response"

	"tmaic/FmsApp/jobs"
	"tmaic/FmsApp/jobs/proto3/protomodel"
)

type EventsController struct {
	controller.BaseController
}

// PostInfo 队列
func (c *EventsController) PostInfo() *response.JsonResult {

	jb := jobs.FmsDemoJob
	jb.SetParam(&protomodel.FmsDemoJob{
		UserName: "pangxianfei",
		Mobile:   "18688678181",
	})

	if err := work.Dispatch(jb); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	return response.JsonSuccess()
}
