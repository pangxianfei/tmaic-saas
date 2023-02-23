package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/http/controllers/api"
	"tmaic/app/http/middleware"
)

func Route(app *iris.Application) {
	routeStatic(app)
	routeRedirect(app)

	mvc.Configure(app.Party("/user"), func(m *mvc.Application) {
		m.Party("/login").Handle(new(api.LoginController))
	})
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middleware.JWT.Serve, middleware.Login)
		m.Party("/article").Handle(new(api.ArticleController))
		m.Party("/ev").Handle(new(api.EventsController))
		m.Party("/user").Handle(new(api.UserController))
		m.Party("/upload").Handle(new(api.UploadController))

	})

}
