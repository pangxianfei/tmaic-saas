package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/FmsApp/http/controllers/v1"
	FmsAppMiddleware "tmaic/FmsApp/http/middleware"
)

func FmsAppApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), FmsAppMiddleware.FmsAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/fms").Handle(new(v1.FmsController))
	})
}
