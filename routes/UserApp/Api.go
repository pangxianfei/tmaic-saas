package UserAppRoute

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/UserApp/http/controllers/api"
	"tmaic/app/http/middleware"
)

func UserAppApi(app *iris.Application) {
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/auth").Handle(new(api.LoginController))
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middleware.Login, middleware.Tenant)
		m.Party("/user").Handle(new(api.UserController))
		m.Party("/saas").Handle(new(api.SaasController))
	})
}
