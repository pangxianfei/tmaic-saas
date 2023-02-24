package route

import (
	"github.com/kataras/iris/v12"
)

// Route 注册路由
func Route(app *iris.Application) {
	routeRedirect(app)
	routeStatic(app)
	routeRedirect(app)
	routeUser(app)
	routeApi(app)
}
