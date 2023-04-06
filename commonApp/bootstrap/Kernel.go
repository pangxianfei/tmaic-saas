package bootstrap

import (
	"gitee.com/pangxianfei/framework/cache"
	"gitee.com/pangxianfei/framework/database"
	"gitee.com/pangxianfei/framework/filesystem"
	"gitee.com/pangxianfei/framework/kernel/zone"
	"gitee.com/pangxianfei/framework/log"
	"gitee.com/pangxianfei/framework/queue/producerconsumer"

	"tmaic/LoginApp/events"
	"tmaic/LoginApp/jobs"
	"tmaic/LoginApp/listeners"
	"tmaic/commonApp/config"
	"tmaic/commonApp/lang"
)

func (s *Saas) Initialize() {
	config.Initialize()
	log.Initialize()
	database.Initialize()
	zone.Initialize()
	lang.Initialize()
	cache.Initialize()
	producerconsumer.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()
	filesystem.Initialize()

}
