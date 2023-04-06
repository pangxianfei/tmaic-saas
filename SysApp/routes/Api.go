package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/SysApp/http/controllers/v1"
	adminMiddleware "tmaic/SysApp/http/middleware"
)

func AppRouteApi(app *iris.Application) {
	auth := app.Party("/sys")
	auth.Use(middleware.LoginMiddleware(), adminMiddleware.Admin)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/tenant").Handle(new(v1.TenantController))
		m.Party("/events").Handle(new(v1.EventsController))
		m.Party("/upload").Handle(new(v1.UploadController))
		m.Party("/app").Handle(new(v1.AppInfoController))
		m.Party("/app/permission").Handle(new(v1.AppPermissionController))
		m.Party("/authority").Handle(new(v1.AuthorityController))
	})
}
