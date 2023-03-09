package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	route "tmaic/routes"
	SysApp "tmaic/routes/SysApp"
)

func SysRun() error {
	app := iris.New().SetName("SysApp")
	//注册资源及心跳路由
	route.Route(app)
	SysApp.SysAppRoute(app)
	RouteNameList(app, config.Instance.App.SysApp, config.Instance.AppPort.SysAppPort, config.Instance.AppNo.Sys)
	SysErr := SetAppConfig(app, config.Instance.AppPort.SysAppPort)
	_ = app.Build()
	return SysErr

}
