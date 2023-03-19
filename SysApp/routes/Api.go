package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	api2 "tmaic/SysApp/http/controllers/api"
	SysAppiddleware "tmaic/SysApp/http/middleware"
)

func AppRouteApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), SysAppiddleware.SysAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/events").Handle(new(api2.EventsController))
		m.Party("/upload").Handle(new(api2.UploadController))
		m.Party("/app").Handle(new(api2.AppInfoController))
		m.Party("/authority").Handle(new(api2.AuthorityController))
	})
}
