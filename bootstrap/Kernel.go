package bootstrap

import (
	"gitee.com/pangxianfei/framework/cache"
	"gitee.com/pangxianfei/framework/kernel/zone"
	"gitee.com/pangxianfei/framework/logs"
	"gitee.com/pangxianfei/framework/queue"
	"tmaic/app/events"
	"tmaic/app/jobs"
	"tmaic/app/listeners"
	"tmaic/config"
	"tmaic/lang"
)

// Initialize 出始化框架
func Initialize() {
	config.Initialize()
	//database.Initialize()
	//variable.InitBasePath()
	//sentry.Initialize()
	logs.Initialize()
	zone.Initialize()
	lang.Initialize()
	cache.Initialize()
	//m.Initialize()
	queue.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()

}
