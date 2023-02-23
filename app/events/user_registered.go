package events

import (
	"github.com/golang/protobuf/proto"
	listenmodel "tmaic/app/events/protocol_model/listenmodel"
	"tmaic/vendors/framework/hub"
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
