package bootstrap

import (
	"tmaic/app/common"
	"tmaic/app/cron"
)

func (s *Saas) EnablingScheduledTask() {
	if !common.IsProd() {
		return
	}
	cron.StartSchedule()
	return
}
