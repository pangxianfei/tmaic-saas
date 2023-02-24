package memory

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/jinzhu/copier"
	c "github.com/patrickmn/go-cache"
	"time"

	. "tmaic/vendors/framework/cache/utils"
	"tmaic/vendors/framework/helpers/zone"
)

func NewMemory(prefix string, defaultExpirationMinute uint, cleanUpIntervalMinute uint) *memory {
	return &memory{
		memoryBasic{
			cache:  c.New(zone.Duration(defaultExpirationMinute)*zone.Minute, zone.Duration(cleanUpIntervalMinute)*zone.Minute),
			prefix: prefix,
		},
	}
}

type memory struct {
	memoryBasic
}

func (m *memory) SetNx(key string, value string, future time.Duration) bool {
	//TODO implement me
	panic("implement me")
}

func (m *memory) GetString(key string) string {
	//TODO implement me
	panic("implement me")
}

/*
func (m *memory) Set(key string, value string, future time.Duration) bool {
	return true
}
func (m *memory) GetString(key string) string {
	return ""
}
*/

func (m *memory) Pget(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error {
	k := NewKey(key, m.Prefix())

	valueInterface, found := m.cache.Get(k.Prefixed())
	if !found {

		if len(defaultValuePtr) > 0 {
			return copier.Copy(valuePtr, defaultValuePtr[0])
		}
		return errors.New("key not exist")
	}

	valueBytes, ok := valueInterface.([]byte)
	if !ok {
		return errors.New("key's value is not a valid proto buffer")
	}
	if err := proto.Unmarshal(valueBytes, valuePtr); err != nil {
		return err
	}
	return nil
}

// ------------------------------------------------------------------------------
// the same
// ------------------------------------------------------------------------------

func (m *memory) Ppull(key string, valuePtr proto.Message, defaultValuePtr ...proto.Message) error {
	k := NewKey(key, m.Prefix())

	err := m.Pget(k.Raw(), valuePtr, defaultValuePtr...)
	if err != nil {
		return err
	}

	m.Forget(k.Raw())

	return nil
}
func (m *memory) Pput(key string, valuePtr proto.Message, future time.Duration) bool {
	valueBytes, err := proto.Marshal(valuePtr)
	if err != nil {
		return false
	}
	return m.Put(key, valueBytes, future)
}
func (m *memory) Padd(key string, valuePtr proto.Message, future time.Duration) bool {
	valueBytes, err := proto.Marshal(valuePtr)
	if err != nil {
		return false
	}
	return m.Add(key, valueBytes, future)
}
func (m *memory) Pforever(key string, valuePtr proto.Message) bool {
	valueBytes, err := proto.Marshal(valuePtr)
	if err != nil {
		return false
	}
	return m.Forever(key, valueBytes)
}
