package route

import (
	"gitee.com/pangxianfei/framework/trans"
	"github.com/kataras/iris/v12"
)

// Route 注册公共路由
func Route(app *iris.Application) {
	//心跳路由
	app.Any("/ping", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": 200, "status": true, "lang": trans.Get("address.not.exist")})
	})

	app.Any("/", func(app iris.Context) {
		_, _ = app.HTML("<h1>Powered by Tmaic SAAS Framework</h1>")
	})

	app.HandleDir("/storage", "./commonApp/storage/app")
}

func NotFound(ctx iris.Context) {
	ctx.JSON(iris.Map{"code": 404, "msg": trans.Get("address.not.exist")})
}

func InternalServerError(ctx iris.Context) {
	ctx.JSON(iris.Map{"code": iris.StatusInternalServerError, "msg": "Oups something went wrong, try again"})
}
