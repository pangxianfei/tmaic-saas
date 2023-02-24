package queue

import (
	"errors"
	"strconv"

	"tmaic/vendors/framework/cmd"
	"tmaic/vendors/framework/queue"
)

func init() {
	cmd.Add(&Retry{})
}

type Retry struct {
}

func (r *Retry) Command() string {
	return "queue:retry {queue_id}"
}

func (r *Retry) Description() string {
	return "Retry a failed queue from database"
}

func (r *Retry) Handler(arg *cmd.Arg) error {
	queueIdPtr, err := arg.Get("queue_id")
	if err != nil {
		return err
	}

	if queueIdPtr == nil {
		return errors.New("queue_id is invalid")
	}

	queueId, err := strconv.ParseUint(*queueIdPtr, 10, 32)
	if err != nil {
		return err
	}

	if err := queue.Retry(uint(queueId)); err != nil {
		return err
	}

	return nil
}
