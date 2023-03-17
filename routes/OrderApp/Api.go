package OrderApp

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/OrderApp/http/controllers/api"
	OrderAppMiddleware "tmaic/app/OrderApp/http/middleware"
)

func OrderApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), OrderAppMiddleware.OrderAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/order").Handle(new(api.OrderController))
	})
}
