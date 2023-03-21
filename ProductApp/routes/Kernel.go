package routes

import (
	"github.com/kataras/iris/v12"
)

// ProductRoute 注册路由
func ProductRoute(app *iris.Application) {
	ProductApi(app)
}
