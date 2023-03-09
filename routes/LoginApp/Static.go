package LoginApp

import "github.com/kataras/iris/v12"

func RouteStatic(app *iris.Application) {
	app.HandleDir("/", "./assets")
}
