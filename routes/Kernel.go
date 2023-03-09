package route

import (
	"github.com/kataras/iris/v12"
	"tmaic/routes/LoginApp"
)

// Route 注册公共路由
func Route(app *iris.Application) {
	LoginApp.RoutePing(app)
	LoginApp.RouteRedirect(app)
	LoginApp.RouteStatic(app)
}
