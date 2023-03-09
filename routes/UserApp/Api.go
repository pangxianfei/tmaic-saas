package UserAppRoute

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/UserApp/http/controllers/api"
	UserAppMiddleware "tmaic/app/UserApp/http/middleware"
)

func UserAppApi(app *iris.Application) {

	auth := app.Party("/")
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Router.Use(UserAppMiddleware.UserAppMiddleware)
		m.Party("/staff").Handle(new(api.UserController))
		m.Party("/saas").Handle(new(api.SaasController))
	})
}
