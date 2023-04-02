package console

import (
	"gitee.com/pangxianfei/framework/cmd"
)

// Schedule 执行定时任务
func Schedule(schedule *cmd.Schedule) {
	schedule.Command("queue:work login").EverySecond()
}
