package jobs

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	"gitee.com/pangxianfei/framework/queue/work"
	"github.com/golang/protobuf/proto"

	"tmaic/FmsApp/jobs/proto3/protomodel"
)

func init() {
	work.Add(&fmsDemoJob{})
}

var FmsDemoJob = new(fmsDemoJob)

type fmsDemoJob struct {
	work.Job
}

// Retries 失败重启次数
func (jb *fmsDemoJob) Retries() uint32 {
	return 3
}

// Name 列队名称  Topics 名
func (jb *fmsDemoJob) Name() string {
	return "fmsdemo"
}

func (jb *fmsDemoJob) SetParam(paramPtr proto.Message) {
	jb.Job.SetParam(paramPtr)
}

func (jb *fmsDemoJob) ParamData() proto.Message {
	return jb.Job.ParamProto()
}

// ParamProto proto 类名参数 实例
func (jb *fmsDemoJob) ParamProto() proto.Message {
	return &protomodel.FmsDemoJob{}
}

// Handle 执行
func (jb *fmsDemoJob) Handle(paramPtr proto.Message) error {
	fmsDemoJobObj := paramPtr.(*protomodel.FmsDemoJob)
	debug.Dd(fmsDemoJobObj)
	return nil
}
