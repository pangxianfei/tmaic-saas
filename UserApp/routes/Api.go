package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	api2 "tmaic/UserApp/http/controllers/api"
	UserAppMiddleware "tmaic/UserApp/http/middleware"
)

func UserAppApi(app *iris.Application) {

	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), UserAppMiddleware.UserAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/staff").Handle(new(api2.UserController))
		m.Party("/permission").Handle(new(api2.PermissionController))
		m.Party("/role").Handle(new(api2.RoleController))
		m.Party("/saas").Handle(new(api2.SaasController))
	})
}
