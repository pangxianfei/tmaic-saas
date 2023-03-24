package listeners

import (
	"errors"

	"gitee.com/pangxianfei/framework/hub"
	"gitee.com/pangxianfei/saas/sysmodel"
	"github.com/golang/protobuf/proto"

	"tmaic/app/events"
	"tmaic/app/events/protocol_model/listenmodel"
	//"tmaic/app/platform/models"
	"gitee.com/pangxianfei/framework/kernel/log"
)

func init() {
	hub.Register(&Test{})
}

type Test struct {
	user sysmdel.TenantUser
	hub.Listen
}

// Name topics主题名称
func (user *Test) Name() hub.ListenerName {
	return "add-test"
}

func (user *Test) Subscribe() (eventPtrList []hub.Eventer) {
	return []hub.Eventer{
		&events.TestEvents{},
	}
}

func (user *Test) Construct(paramPtr proto.Message) error {
	log.Info("第一执行这里业务代码 这里通常为 Handle 数据作为准备的工作")
	test, status := paramPtr.(*listenmodel.TestEvents)
	// 出始化 user 数据 handle 直接拿user 数据

	if status == false {
		return errors.New("数据有误")
	}

	user.user.Id = int64(test.Id)
	user.user.UserName = test.Name
	return nil
}

func (user *Test) Handle() error {
	log.Info("第二执行这里业务代码")
	log.Info(user.user.Id)
	log.Info(user.user.UserName)
	//errors.New("处理完成")
	return nil
}
