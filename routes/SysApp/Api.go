package SysAppRoute

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/SysApp/http/controllers/api"
	SysAppiddleware "tmaic/app/SysApp/http/middleware"
)

func AppRouteApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), SysAppiddleware.SysAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/events").Handle(new(api.EventsController))
		m.Party("/upload").Handle(new(api.UploadController))
		m.Party("/app").Handle(new(api.AppInfoController))
		m.Party("/authority").Handle(new(api.AuthorityController))
	})
}
