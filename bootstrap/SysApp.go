package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	route "tmaic/routes"
	SysApp "tmaic/routes/SysApp"
)

func SysRun() error {

	app := iris.New().SetName("SysApp")
	app.Validator = validator.New()
	//注册路由
	route.Route(app)
	SysApp.SysAppRoute(app)
	RouteNameList(app, config.Instance.App.SysApp, config.Instance.AppPort.SysAppPort)

	SysErr := SetAppConfig(app, config.Instance.AppPort.SysAppPort)
	_ = app.Build()
	return SysErr

}
