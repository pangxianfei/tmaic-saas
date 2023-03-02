package bootstrap

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"

	"tmaic/app/http/middleware"
	route "tmaic/routes"
	SysApp "tmaic/routes/SysApp"
)

func SysRun() error { //parentCtx context.Context

	app := iris.New().SetName("SysApp")
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.Use(logger.New())
	app.AllowMethods(iris.MethodOptions)
	//跨域
	app.Use(middleware.CORS)
	//注册路由
	route.Route(app)
	SysApp.SysAppRoute(app)

	RouteNameList(app, "平台应用路由列表:")
	_ = app.Build()
	return app.Listen(":9999")
}
