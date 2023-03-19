package bootstrap

import (
	"gitee.com/pangxianfei/framework/cache"
	"gitee.com/pangxianfei/framework/database"
	"gitee.com/pangxianfei/framework/kernel/zone"
	"gitee.com/pangxianfei/framework/logs"
	"gitee.com/pangxianfei/framework/queue"
	"tmaic/app/events"
	"tmaic/app/jobs"
	"tmaic/app/listeners"
	"tmaic/config"
	"tmaic/lang"
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
