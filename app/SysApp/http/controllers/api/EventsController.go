package api

import (
	"tmaic/app/events"
	pbs "tmaic/app/events/protocol_model/listenmodel"
	jobsPbs "tmaic/app/jobs/protocol_jobs"

	"gitee.com/pangxianfei/framework/hub"
	jobber "gitee.com/pangxianfei/framework/job"
	"gitee.com/pangxianfei/framework/kernel/log"
	"gitee.com/pangxianfei/framework/kernel/tmaic"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"

	ab "tmaic/app/jobs"
)

type EventsController struct {
	Ctx iris.Context
}

// GetByInfo 队列
func (e *EventsController) PostInfo(articleId int64) *simple.JsonResult {
	test := new(events.Test)
	test.SetParam(&pbs.Test{
		Id:   1000,
		Name: "PangXianFei",
	})

	if errs := hub.Emit(test); errs != nil {
		log.Info("user registered event emit failed", tmaic.V{"event": test, "errors": errs})
	}

	jb := new(ab.DemoJob)
	jb.SetParam(&jobsPbs.DemoJob{
		Query:         string("parmigiana"),
		PageNumber:    1,
		ResultPerPage: 3,
	})

	err := jobber.Dispatch(jb)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	return simple.JsonSuccess()
}
