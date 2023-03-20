package console

import (
	"gitee.com/pangxianfei/framework/cmd"
)

func Schedule(schedule *cmd.Schedule) {
	schedule.Command("demo:say words,9999999999").EverySecond()
}
