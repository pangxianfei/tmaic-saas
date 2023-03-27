package MQ

import (
	"gitee.com/pangxianfei/framework/work"

	"tmaic/FmsApp/jobs"
	"tmaic/FmsApp/jobs/proto3/protomodel"
)

//import (
//"gitee.com/pangxianfei/framework/work"

//"tmaic/FmsApp/jobs"
//"tmaic/FmsApp/jobs/proto3/protomodel"
//"tmaic/FmsApp/models"
//)

var FmsDemoMessage = new(FmsDemoDispatch)

type fms_demoMessage interface {
	Dispatch() bool
}
type FmsDemoDispatch struct{}

func (l *FmsDemoDispatch) Dispatch(admininfo *protomodel.FmsDemoJob) bool {

	FmsDemoJob := jobs.FmsDemoJob
	FmsDemoJob.SetParam(&protomodel.FmsDemoJob{
		UserName: admininfo.UserName,
		Mobile:   admininfo.Mobile,
	})
	if jobErr := work.Dispatch(FmsDemoJob); jobErr != nil {
		return false
	}
	return true
}
