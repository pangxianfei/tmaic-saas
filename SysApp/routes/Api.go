package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/SysApp/http/controllers/api"
	adminMiddleware "tmaic/SysApp/http/middleware"
)

func AppRouteApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), adminMiddleware.Admin)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/events").Handle(new(api.EventsController))
		m.Party("/upload").Handle(new(api.UploadController))
		m.Party("/app").Handle(new(api.AppInfoController))
		m.Party("/app/permission").Handle(new(api.AppPermissionController))
		m.Party("/authority").Handle(new(api.AuthorityController))
		m.Party("/saas").Handle(new(api.CreateDatabaseController))
	})
}
