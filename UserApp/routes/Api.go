package routes

import (
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/UserApp/http/controllers/v1"
	UserAppMiddleware "tmaic/UserApp/http/middleware"
)

func UserAppApi(app *iris.Application) {
	auth := app.Party("/staff")
	auth.Use(middleware.LoginMiddleware(), UserAppMiddleware.UserAppMiddleware, middleware.Permissions)
	mvc.Configure(auth, func(m *mvc.Application) {
		m.Party("/user").Handle(new(v1.AdminController))
		m.Party("/personnel").Handle(new(v1.StaffController))
		m.Party("/permission").Handle(new(v1.PermissionController))
		m.Party("/role").Handle(new(v1.RoleController))
	})
}
