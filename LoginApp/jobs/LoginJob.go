package jobs

import (
	"bufio"
	"fmt"
	"os"

	"gitee.com/pangxianfei/framework/work"
	"github.com/golang/protobuf/proto"

	"tmaic/LoginApp/jobs/proto3/protomodel"
)

func init() {
	work.Add(&loginJob{})
}

var LoginJob = new(loginJob)

type loginJob struct {
	work.Job
}

// Retries 失败重启次数
func (e *loginJob) Retries() uint32 {
	return 3
}

// Name 列队名称  Topics 名
func (e *loginJob) Name() string {
	return "login"
}

func (e *loginJob) SetParam(paramPtr proto.Message) {
	e.Job.SetParam(paramPtr)
}

func (e *loginJob) ParamData() proto.Message {
	return e.Job.ParamProto()
}

// ParamProto proto 类名参数 实例
func (e *loginJob) ParamProto() proto.Message {
	return &protomodel.LoginJob{}
}

// Handle 执行
func (e *loginJob) Handle(paramPtr proto.Message) error {
	LoginJobObj := paramPtr.(*protomodel.LoginJob)
	filePath := "G:/go/src/saas/commonApp/storage/logs/file.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	_, _ = write.WriteString(LoginJobObj.UserName + "-" + LoginJobObj.Mobile + "\n")
	_ = write.Flush()
	return nil
}
