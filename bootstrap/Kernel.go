package bootstrap

import (
	"gitee.com/pangxianfei/framework/cache"
	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/framework/helpers/zone"
	"gitee.com/pangxianfei/framework/logs"
	"gitee.com/pangxianfei/framework/queue"
	"github.com/kataras/iris/v12"
	"strings"
	"tmaic/app/events"
	"tmaic/app/jobs"
	"tmaic/app/listeners"
	"tmaic/config"
	"tmaic/resources/lang"
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

// RouteNameList 打印路由列表
func RouteNameList(app *iris.Application, msgOut string) {

	routeList := app.GetRoutes()
	var index int = 1
	//console.Println(console.CODE_INFO, msgOut)
	for _, value := range routeList {
		if strings.Contains(value.MainHandlerName, "tmaic") || strings.Contains(value.MainHandlerName, "iris") {
			continue
		}

		if value.Method == "POST" || value.Method == "GET" {
			console.Println(console.CODE_SUCCESS, " "+console.Sprintf(console.CODE_SUCCESS, "%-4d", index)+console.Sprintf(console.CODE_SUCCESS, "%-40s", value.MainHandlerName)+console.Sprintf(console.CODE_SUCCESS, "%-5s", value.Method)+" "+console.Sprintf(console.CODE_SUCCESS, "%-45s", value.Path)+console.Sprintf(console.CODE_SUCCESS, "%-30s", value.FormattedPath)+console.Sprintf(console.CODE_SUCCESS, "%-20s", value.RegisterFileName))
			index++
		}

	}
}
