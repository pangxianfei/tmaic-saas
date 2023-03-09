package LoginApp

import (
	"github.com/kataras/iris/v12"
)

// LoginAppRoute 注册路由
func LoginAppRoute(app *iris.Application) {
	LoginApi(app)
}
