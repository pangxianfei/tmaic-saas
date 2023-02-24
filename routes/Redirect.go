package route

import "github.com/kataras/iris/v12"

func routeRedirect(app *iris.Application) {
	app.Any("/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{"code": 200, "status": true})
	})
}
