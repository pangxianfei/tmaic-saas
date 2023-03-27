package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/OrderApp/http/controllers/v1"
	OrderAppMiddleware "tmaic/OrderApp/http/middleware"
)

func Route(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), OrderAppMiddleware.OrderAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/order").Handle(new(v1.OrderController))
	})
}
