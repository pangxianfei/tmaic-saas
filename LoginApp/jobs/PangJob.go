package jobs

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	"gitee.com/pangxianfei/framework/work"
	"github.com/golang/protobuf/proto"

	"tmaic/LoginApp/jobs/proto3/protomodel"
)

func init() {
	work.Add(&pangJob{})
}

var PangJob = new(pangJob)

type pangJob struct {
	work.Job
}

// Retries 失败重启次数
func (jb *pangJob) Retries() uint32 {
	return 3
}

// Name 列队名称  Topics 名
func (jb *pangJob) Name() string {
	return "pang"
}

func (jb *pangJob) SetParam(paramPtr proto.Message) {
	jb.Job.SetParam(paramPtr)
}

func (jb *pangJob) ParamData() proto.Message {
	return jb.Job.ParamProto()
}

// ParamProto proto 类名参数 实例
func (jb *pangJob) ParamProto() proto.Message {
	return &protomodel.PangJob{}
}

// Handle 执行
func (jb *pangJob) Handle(paramPtr proto.Message) error {
	pangObj := paramPtr.(*protomodel.PangJob)
	debug.Dd(pangObj)
	return nil
}
