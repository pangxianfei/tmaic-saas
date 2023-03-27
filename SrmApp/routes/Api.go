package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/SrmApp/http/controllers/v1"
	SrmAppMiddleware "tmaic/SrmApp/http/middleware"
)

func SrmAppApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), SrmAppMiddleware.SrmAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/srm").Handle(new(v1.SrmController))
	})
}
