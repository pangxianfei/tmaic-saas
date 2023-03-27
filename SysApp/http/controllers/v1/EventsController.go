package v1

import (
	"gitee.com/pangxianfei/framework/facades"
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"

	"gitee.com/pangxianfei/framework/hub"
	"gitee.com/pangxianfei/framework/kernel/tmaic"
	"gitee.com/pangxianfei/framework/work"

	"tmaic/LoginApp/events"
	"tmaic/LoginApp/events/protocol_model/listenmodel"
	"tmaic/LoginApp/jobs"
	"tmaic/LoginApp/jobs/proto3/protomodel"
)

type EventsController struct {
	controller.BaseController
}

// PostInfo 队列
func (c *EventsController) PostInfo() *response.JsonResult {

	test := new(events.TestEvents)
	test.SetParam(&listenmodel.TestEvents{
		Id:   1000,
		Name: "PangXianFei",
	})

	if errs := hub.Emit(test); errs != nil {
		facades.Log.Info("user registered event emit failed", tmaic.V{"event": test, "errors": errs})
	}

	jb := jobs.DemoJob
	jb.SetParam(&protomodel.DemoJob{
		Query:         "parmigiana",
		PageNumber:    1,
		ResultPerPage: 3,
	})

	if err := work.Dispatch(jb); err != nil {
		return response.JsonErrorMsg(err.Error())
	}

	return response.JsonSuccess()
}
