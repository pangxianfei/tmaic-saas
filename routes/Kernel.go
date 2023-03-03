package route

import (
	"github.com/kataras/iris/v12"
	SysAppRoute "tmaic/routes/SysApp"
)

// Route 注册路由
func Route(app *iris.Application) {
	SysAppRoute.RouteRedirect(app)
	SysAppRoute.RouteStatic(app)
	//注册应用路由
}
