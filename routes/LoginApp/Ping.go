package LoginApp

import (
	"gitee.com/pangxianfei/framework/trans"
	"github.com/kataras/iris/v12"
)

func RoutePing(app *iris.Application) {
	//心跳路由
	app.Any("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": 200, "status": true, "lang": trans.Get("address.not.exist")})
	})

	app.Any("/", func(app iris.Context) {
		_, _ = app.HTML("<h1>Powered by Tmaic SAAS Framework</h1>")
	})
}
