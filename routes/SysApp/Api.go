package SysAppRoute

import (
	"gitee.com/pangxianfei/framework/helpers"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/http/controllers/api"
	"tmaic/app/http/middleware"
)

func SysAppRouteApi(app *iris.Application) {
	//心跳路由
	app.Any("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": 200, "status": true, "lang": helpers.L("auth.register.failed_token_generate_error")})
	})

	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middleware.Login)
		m.Party("/ev").Handle(new(api.EventsController))
		m.Party("/upload").Handle(new(api.UploadController))
	})
}
