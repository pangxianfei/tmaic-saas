package UserAppRoute

import (
	"github.com/kataras/iris/v12"
)

// UserAppRoute 注册路由
func UserAppRoute(app *iris.Application) {
	UserAppApi(app)

}
