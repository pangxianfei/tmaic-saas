package CommonModel

import (
	"gitee.com/pangxianfei/framework/kernel/pb"
	"gitee.com/pangxianfei/framework/kernel/zone"
	"gitee.com/pangxianfei/framework/queue"
	message "gitee.com/pangxianfei/framework/queue/protocol_buffers"
	"gitee.com/pangxianfei/simple"
)

type FailedQueue struct {
	ID        *uint          `gorm:"column:failed_queue_id;primary_key;auto_increment"`
	Hash      *string        `gorm:"column:failed_queue_hash;type:varchar(100);unique_index;not null"`
	Topic     *string        `gorm:"column:failed_queue_topic_name;type:varchar(100);not null"`
	Channel   *string        `gorm:"column:failed_queue_channel_name;type:varchar(100);not null"`
	DataProto *[]byte        `gorm:"column:failed_queue_data;type:varbinary"`
	PushedAt  *zone.Time     `gorm:"column:failed_queue_pushed_at;not null"`
	Delay     *zone.Duration `gorm:"column:failed_queue_delay;type:bigint;not null"`
	Retries   *uint32        `gorm:"column:failed_queue_retries;type:integer unsigned;not null"`
	Tried     *uint32        `gorm:"column:failed_queue_tried;type:integer unsigned;not null"`
	Err       *string        `gorm:"column:failed_queue_err;size:65535"`

	CreatedAt *zone.Time `gorm:"column:failed_queue_created_at"`
	UpdatedAt zone.Time  `gorm:"column:failed_queue_updated_at"`
	DeletedAt *zone.Time `gorm:"column:failed_queue_deleted_at"`
}

func (fq *FailedQueue) TableName() string {
	return "mq_failed_queue"
}

func (fq *FailedQueue) RetryTopic() string {
	return *fq.Topic
}

func (fq *FailedQueue) RetryChannel() string {
	return *fq.Channel
}

func (fq *FailedQueue) RetryRetries() uint32 {
	return *fq.Retries
}

func (fq *FailedQueue) RetryDelay() zone.Duration {
	return *fq.Delay
}

func (fq *FailedQueue) RetryParamProtoBytes() []byte {
	return *fq.DataProto
}
func (fq *FailedQueue) RetryHash() string {
	return *fq.Hash
}

func (fq *FailedQueue) Default() interface{} {
	return FailedQueue{}
}

func (fq *FailedQueue) FailedToDatabase(topicName string, channelName string, msg *message.Message, errStr string) error {

	_fq := FailedQueue{
		Topic:     &topicName,
		Channel:   &channelName,
		Hash:      &msg.Hash,
		DataProto: &msg.Param,
		PushedAt:  pb.TimestampConvert(msg.PushedAt),
		Delay:     pb.DurationConvert(msg.Delay),
		Retries:   &msg.Retries,
		Tried:     &msg.Tried,
		Err:       &errStr,
	}
	return simple.DB().Create(&_fq).Error
}

func (fq *FailedQueue) FailedQueueById(id uint) (failedQueuerPtr queue.FailedQueuer, err error) {
	mq := FailedQueue{ID: &id}
	if err := simple.DB().First(&mq).Error; err != nil {
		return nil, err
	}
	return &mq, nil
}

func (fq *FailedQueue) DeleteQueueById(id uint) error {
	_fq := FailedQueue{ID: &id}
	if err := simple.DB().Delete(&_fq).Error; err != nil {
		return err
	}
	return nil
}
