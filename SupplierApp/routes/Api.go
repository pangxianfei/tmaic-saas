package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/SupplierApp/http/controllers/v1"
	SupplierAppMiddleware "tmaic/SupplierApp/http/middleware"
)

func SupplierAppApi(app *iris.Application) {
	auth := app.Party("/super")
	auth.Use(middleware.LoginMiddleware(), SupplierAppMiddleware.SupplierAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/finance").Handle(new(v1.FinanceController))
		m.Party("/order").Handle(new(v1.OrderController))
	})
}
