package OrderApp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/OrderApp/http/controllers/api"
	"tmaic/app/http/middleware"
)

func OrderApi(app *iris.Application) {
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middleware.Login)
		m.Party("/order").Handle(new(api.OrderController))
	})
}
