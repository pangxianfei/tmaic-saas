package bootstrap

import (
	"gitee.com/pangxianfei/framework/cache"
	"gitee.com/pangxianfei/framework/database"
	"gitee.com/pangxianfei/framework/kernel/zone"
	"gitee.com/pangxianfei/framework/logs"
	"gitee.com/pangxianfei/framework/queue"

	"tmaic/LoginApp/events"
	"tmaic/LoginApp/jobs"
	"tmaic/LoginApp/listeners"
	"tmaic/commonApp/config"
	"tmaic/commonApp/lang"
)

func (s *Saas) Initialize() {
	config.Initialize()
	database.Initialize()
	logs.Initialize()
	zone.Initialize()
	lang.Initialize()
	cache.Initialize()
	queue.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()
}
