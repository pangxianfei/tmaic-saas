package routes

import (
	"github.com/kataras/iris/v12"
)

// OrderRoute 注册路由
func OrderRoute(app *iris.Application) {
	OrderApi(app)
}
