package routes

import (
	"github.com/kataras/iris/v12"
)

// Route 注册路由
func Route(app *iris.Application) {
	WmAppApi(app)
}
