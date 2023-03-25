package events

import (
	"gitee.com/pangxianfei/framework/hub"
	"github.com/golang/protobuf/proto"

	event "tmaic/LoginApp/events/protocol_model/listenmodel"
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
