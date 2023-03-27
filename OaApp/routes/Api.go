package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/OaApp/http/controllers/v1"
	OaAppMiddleware "tmaic/OaApp/http/middleware"
)

func OaAppApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), OaAppMiddleware.OaAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/fms").Handle(new(v1.OaController))
	})
}
