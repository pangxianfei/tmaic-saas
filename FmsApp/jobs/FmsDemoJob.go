package jobs

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	"gitee.com/pangxianfei/framework/work"
	"github.com/golang/protobuf/proto"

	"tmaic/FmsApp/jobs/proto3/protomodel"
)

func init() {
	work.Add(&fms_demoJob{})
}

var FmsDemoJob = new(fms_demoJob)

type fms_demoJob struct {
	work.Job
}

// Retries 失败重启次数
func (jb *fms_demoJob) Retries() uint32 {
	return 3
}

// Name 列队名称  Topics 名
func (jb *fms_demoJob) Name() string {
	return "fms_demo"
}

func (jb *fms_demoJob) SetParam(paramPtr proto.Message) {
	jb.Job.SetParam(paramPtr)
}

func (jb *fms_demoJob) ParamData() proto.Message {
	return jb.Job.ParamProto()
}

// ParamProto proto 类名参数 实例
func (jb *fms_demoJob) ParamProto() proto.Message {
	return &protomodel.FmsDemoJob{}
}

// Handle 执行
func (jb *fms_demoJob) Handle(paramPtr proto.Message) error {
	fms_demoObj := paramPtr.(*protomodel.FmsDemoJob)
	debug.Dd(fms_demoObj)
	return nil
}
