package bootstrap

import (
	"tmaic/app/cron"
	"tmaic/commonApp/common"
)

func (s *Saas) EnablingScheduledTask() {
	if !common.IsProd() {
		return
	}
	cron.StartSchedule()
	return
}
