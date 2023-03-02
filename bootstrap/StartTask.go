package bootstrap

import (
	"tmaic/app/common"
	"tmaic/app/cron"
)

func EnablingScheduledTask() {
	if !common.IsProd() {
		return
	}
	cron.StartSchedule()
}
