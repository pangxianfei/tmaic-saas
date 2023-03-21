package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/UserApp/http/controllers/api"
	UserAppMiddleware "tmaic/UserApp/http/middleware"
)

func UserAppApi(app *iris.Application) {
	auth := app.Party("/")
	auth.Use(middleware.LoginMiddleware(), UserAppMiddleware.UserAppMiddleware)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/staff").Handle(new(api.UserController))
		m.Party("/permission").Handle(new(api.PermissionController))
		m.Party("/role").Handle(new(api.RoleController))
		m.Party("/saas").Handle(new(api.SaasController))
		m.Party("/saas").Handle(new(api.CreateDatabaseController))
	})
}
