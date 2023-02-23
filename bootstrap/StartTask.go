package bootstrap

import (
	"tmaic/app/common"
	"tmaic/app/cron"
)

// EnablingAscheduledTask 开启定时任务
func EnablingAscheduledTask() {
	if !common.IsProd() {
		return
	}
	cron.StartSchedule()
}
