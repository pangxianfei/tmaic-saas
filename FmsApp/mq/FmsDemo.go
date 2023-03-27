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

type fmsDemoMessage interface {
	Dispatch() bool
}
type FmsDemoDispatch struct{}

func (l *FmsDemoDispatch) Dispatch() bool {
	FmsDemoJob := jobs.FmsDemoJob
	FmsDemoJob.SetParam(&protomodel.FmsDemoJob{UserName: "pangxianfei", Mobile: "18688678181"})
	if jobErr := work.Dispatch(FmsDemoJob); jobErr != nil {
		return false
	}
	return true
}
