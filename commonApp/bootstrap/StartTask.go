package bootstrap

import (
	"tmaic/commonApp/common"
)

func (s *Saas) EnablingScheduledTask() {
	if !common.IsProd() {
		return
	}
	return
}
