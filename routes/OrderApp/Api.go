package OrderApp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/OrderApp/http/controllers/api"
	OrderAppMiddleware "tmaic/app/OrderApp/http/middleware"
	"tmaic/app/http/middleware"
)

func OrderApi(app *iris.Application) {
	mvc.Configure(app.Party("/"), func(m *mvc.Application) {
		m.Router.Use(middleware.LoginMiddleware(), OrderAppMiddleware.OrderAppMiddleware)
		m.Party("/order").Handle(new(api.OrderController))
	})
}
