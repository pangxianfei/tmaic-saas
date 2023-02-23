package events

import (
	"github.com/golang/protobuf/proto"
	event "tmaic/app/events/protocol_model/listenmodel"
	"tmaic/vendors/framework/hub"
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
