package route

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/http/controllers/api"
	"tmaic/app/http/middleware"
)

func routeUser(app *iris.Application) {
	//心跳路由
	app.Any("/ping", func(ctx iris.Context) {
		//ctx.JSON(iris.Map{"code": 200, "status": true, "lang": helpers.L("auth.register.failed_token_generate_error")})
		ctx.JSON(iris.Map{"code": 200, "status": true})
	})

	mvc.Configure(app.Party("/"), func(m *mvc.Application) {
		m.Party("auth").Handle(new(api.LoginController))
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middleware.Login)
		m.Party("/user").Handle(new(api.UserController))
	})
}
