package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/WmApp/http/controllers/v1"
	WmAppMiddleware "tmaic/WmApp/http/middleware"
)

func WmAppApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), WmAppMiddleware.WmAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/wm").Handle(new(v1.WmController))
	})
}
