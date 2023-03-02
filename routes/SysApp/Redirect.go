package SysAppRoute

import "github.com/kataras/iris/v12"

func RouteRedirect(app *iris.Application) {
	app.Any("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": 200, "status": true})
	})
}
