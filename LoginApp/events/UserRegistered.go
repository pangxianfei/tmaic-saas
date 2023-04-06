package events

import (
	"gitee.com/pangxianfei/framework/queue/hub"
	"github.com/golang/protobuf/proto"

	"tmaic/LoginApp/events/protocol_model/listenmodel"
)

func init() {
	hub.Make(&UserRegistered{})
}

type UserRegistered struct {
	hub.Event
	UserId uint32
}

func (ur *UserRegistered) ParamProto() proto.Message {
	return &listenmodel.UserRegistered{}
}
