package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/ProductApp/http/controllers/api"
	ProductMiddleware "tmaic/ProductApp/http/middleware"
)

func Route(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), ProductMiddleware.ProductAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/bom").Handle(new(api.PrcoductController))
	})
}
