package UserAppRoute

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/UserApp/http/controllers/api"
	UserAppMiddleware "tmaic/app/UserApp/http/middleware"
	"tmaic/app/http/middleware"
)

func UserAppApi(app *iris.Application) {
	mvc.Configure(app.Party("/auth"), func(m *mvc.Application) {
		m.Party("/").Handle(new(api.LoginController))
	})

	auth := app.Party("/")
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Router.Use(middleware.LoginMiddleware(), UserAppMiddleware.UserAppMiddleware)
		m.Party("/staff").Handle(new(api.UserController))
		m.Party("/saas").Handle(new(api.SaasController))
	})
}
