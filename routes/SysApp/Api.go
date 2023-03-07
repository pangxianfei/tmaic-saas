package SysAppRoute

import (
	"gitee.com/pangxianfei/framework/trans"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/SysApp/http/controllers/api"
	SysAppiddleware "tmaic/app/SysApp/http/middleware"
	"tmaic/app/http/middleware"
)

func AppRouteApi(app *iris.Application) {
	//心跳路由
	app.Any("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": 200, "status": true, "lang": trans.Get("auth.register.failed_token_generate_error")})
	})

	mvc.Configure(app.Party("/auth"), func(m *mvc.Application) {
		m.Party("/").Handle(new(api.RegisterController))
	})

	mvc.Configure(app.Party("/"), func(m *mvc.Application) {
		m.Router.Use(middleware.LoginMiddleware(), SysAppiddleware.SysAppMiddleware)
		m.Party("/events").Handle(new(api.EventsController))
		m.Party("/upload").Handle(new(api.UploadController))
		m.Party("/app").Handle(new(api.AppInfoController))
	})
}
