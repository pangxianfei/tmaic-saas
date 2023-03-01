package events

import (
	event "tmaic/app/events/protocol_model/listenmodel"

	"gitee.com/pangxianfei/framework/hub"
	"github.com/golang/protobuf/proto"
)

func init() {
	hub.Make(&Test{})
}

type Test struct {
	hub.Event
	Id uint32
}

func (ur *Test) ParamProto() proto.Message {
	return &event.Test{}
}
