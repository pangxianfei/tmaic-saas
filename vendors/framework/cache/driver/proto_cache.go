package driver

import (
	"time"

	"github.com/golang/protobuf/proto"
)

type ProtoCacheGetter interface {
	Pget(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error
}
type ProtoCache interface {
	Ppull(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error
	Pput(key string, valuePtr proto.Message, future time.Duration) bool
	Padd(key string, valuePtr proto.Message, future time.Duration) bool
	Pforever(key string, valuePtr proto.Message) bool

	ProtoCacheGetter
}
