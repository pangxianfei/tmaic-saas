package jobs

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/pangxianfei/framework/work"
	"github.com/golang/protobuf/proto"

	"tmaic/app/jobs/proto3/protomodel"
)

func init() {
	work.Add(&demoJob{})
}

var DemoJob = new(demoJob)

type demoJob struct {
	work.Job
}

// Retries 失败重启次数
func (e *demoJob) Retries() uint32 {
	return 3
}

// Name 列队名称  Topics 名
func (e *demoJob) Name() string {
	return "demo"
}

func (e *demoJob) SetParam(paramPtr proto.Message) {
	e.Job.SetParam(paramPtr)
}

func (e *demoJob) ParamData() proto.Message {
	return e.Job.ParamProto()
}

// ParamProto proto 类名参数 实例
func (e *demoJob) ParamProto() proto.Message {
	return &protomodel.DemoJob{}
}

// Handle 执行
func (e *demoJob) Handle(paramPtr proto.Message) error {
	demoObj := paramPtr.(*protomodel.DemoJob)
	filePath := "G:/go/src/saas/commonApp/storage/logs/file.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	_, _ = write.WriteString(demoObj.String() + "\n")
	_ = write.Flush()
	return nil
	//return errors.New("异常")
}
