package job

import "tmaic/vendors/framework/queue"

func topicName(j jobber) string {
	return "job-" + j.Name()
}
func channelName(j jobber) string {
	return j.Name()
}
func RegisterQueue() {
	for _, j := range jobMap {
		if err := queue.Queue().Register(topicName(j), channelName(j)); err != nil {
			panic(err)
		}
	}
}
