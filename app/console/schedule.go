package console

import (
	"gitee.com/pangxianfei/framework/cmd"
)

func Schedule(schedule *cmd.Schedule) {
	schedule.Command("queue:work login").EverySecond()
}
