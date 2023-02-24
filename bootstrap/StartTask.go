package bootstrap

import (
	"tmaic/app/common"
	"tmaic/app/cron"
)

// EnablingScheduledTask 开启定时任务
func EnablingScheduledTask() {
	if !common.IsProd() {
		return
	}
	cron.StartSchedule()
}
