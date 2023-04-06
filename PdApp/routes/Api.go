package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/PdApp/http/controllers/v1"
	PdAppMiddleware "tmaic/PdApp/http/middleware"
)

func PdAppApi(app *iris.Application) {
	auth := app.Party("/pd")
	auth.Use(middleware.LoginMiddleware(), PdAppMiddleware.PdAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/overview").Handle(new(v1.OverviewController))

	})
}
