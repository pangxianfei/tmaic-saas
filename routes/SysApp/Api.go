package SysAppRoute

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/SysApp/http/controllers/api"
	SysAppiddleware "tmaic/app/SysApp/http/middleware"
)

func AppRouteApi(app *iris.Application) {

	mvc.Configure(app.Party("/"), func(m *mvc.Application) {
		m.Router.Use(SysAppiddleware.SysAppMiddleware)
		m.Party("/events").Handle(new(api.EventsController))
		m.Party("/upload").Handle(new(api.UploadController))
		m.Party("/app").Handle(new(api.AppInfoController))
		m.Party("/authority").Handle(new(api.AuthorityController))
	})
}
