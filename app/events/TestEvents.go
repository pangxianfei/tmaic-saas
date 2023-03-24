package events

import (
	event "tmaic/app/events/protocol_model/listenmodel"

	"gitee.com/pangxianfei/framework/hub"
	"github.com/golang/protobuf/proto"
)

func init() {
	hub.Make(&TestEvents{})
}

type TestEvents struct {
	hub.Event
	Id uint32
}

func (ur *TestEvents) ParamProto() proto.Message {
	return &event.TestEvents{}
}
