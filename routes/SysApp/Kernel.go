package SysAppRoute

import (
	"github.com/kataras/iris/v12"
)

// SysAppRoute 注册路由
func SysAppRoute(app *iris.Application) {
	AppRouteApi(app)
}
