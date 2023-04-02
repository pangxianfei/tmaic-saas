package MQ

import (
	"gitee.com/pangxianfei/framework/queue/work"

	"tmaic/FmsApp/jobs"
	"tmaic/FmsApp/jobs/proto3/protomodel"
)

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
