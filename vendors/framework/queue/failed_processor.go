package queue

import (
	"tmaic/vendors/framework/config"
	"tmaic/vendors/framework/helpers/zone"
	message "tmaic/vendors/framework/queue/protocol_buffers"
)

var failedProcessor FailedProcessor

func initializeFailedProcessor() {
	failedProcessor = config.Get("queue.failed_db_processor_model").(FailedProcessor)
}

type FailedProcessor interface {
	FailedToDatabase(topicName string, channelName string, msg *message.Message, errStr string) error
	FailedQueuer
}

type FailedQueuer interface {
	RetryTopic() string
	RetryHash() string
	RetryChannel() string
	RetryRetries() uint32
	RetryDelay() zone.Duration
	RetryParamProtoBytes() []byte
	FailedQueueById(id uint) (failedQueuerPtr FailedQueuer, err error)
	DeleteQueueById(id uint) error
}
