package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/ProductApp/http/controllers/v1"
	ProductMiddleware "tmaic/ProductApp/http/middleware"
)

func Route(app *iris.Application) {
	auth := app.Party("/bom")
	auth.Use(middleware.LoginMiddleware(), ProductMiddleware.ProductAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/overview").Handle(new(v1.OverviewController))
		m.Party("/prcoduct").Handle(new(v1.PrcoductController))
	})
}
