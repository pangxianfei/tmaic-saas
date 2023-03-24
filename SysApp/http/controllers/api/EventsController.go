package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"

	"tmaic/app/events"
	"tmaic/app/events/protocol_model/listenmodel"
	"tmaic/app/jobs/proto3/protomodel"

	"gitee.com/pangxianfei/framework/hub"
	"gitee.com/pangxianfei/framework/kernel/log"
	"gitee.com/pangxianfei/framework/kernel/tmaic"
	"gitee.com/pangxianfei/framework/work"

	"tmaic/app/jobs"
)

type EventsController struct {
	controller.BaseController
}

// PostInfo 队列
func (e *EventsController) PostInfo() *response.JsonResult {

	test := new(events.TestEvents)
	test.SetParam(&listenmodel.TestEvents{
		Id:   1000,
		Name: "PangXianFei",
	})

	if errs := hub.Emit(test); errs != nil {
		log.Info("user registered event emit failed", tmaic.V{"event": test, "errors": errs})
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
