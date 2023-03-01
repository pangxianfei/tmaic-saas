package events

import (
	listenmodel "tmaic/app/events/protocol_model/listenmodel"

	"gitee.com/pangxianfei/framework/hub"
	"github.com/golang/protobuf/proto"
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
