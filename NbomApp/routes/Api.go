package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/NbomApp/http/controllers/v1"
	NbomAppMiddleware "tmaic/NbomApp/http/middleware"
)

func NbomAppApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), NbomAppMiddleware.NbomAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/bnom").Handle(new(v1.NbomController))
	})
}
