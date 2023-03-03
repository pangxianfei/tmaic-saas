package bootstrap

import (
	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"strings"
	"tmaic/app/http/middleware"
	OrderAppRoute "tmaic/routes/OrderApp"
	UserAppRoute "tmaic/routes/UserApp"
)

func OrderApp() error {
	app := iris.New().SetName("OrderApp")
	OrderAppRoute.OrderRoute(app)
	RouteNameList(app, config.Instance.App.OrderApp, config.Instance.AppPort.OrderPort)
	OrderErr := SetAppConfig(app, config.Instance.AppPort.OrderPort)
	_ = app.Build()
	return OrderErr
}

func UserApp() error {
	app := iris.New().SetName("UserApp")
	UserAppRoute.UserAppRoute(app)
	RouteNameList(app, config.Instance.App.UserApp, config.Instance.AppPort.UserPort)
	UserErr := SetAppConfig(app, config.Instance.AppPort.UserPort)
	_ = app.Build()
	return UserErr
}

func SetAppConfig(app *iris.Application, Port string) error {
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS)
	return app.Run(iris.Addr(Port), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 true,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		TimeFormat:                        "Mon, 02 Jan 2006 15:04:05 GMT",
		Charset:                           "UTF-8",
	}))
}

// RouteNameList 打印路由列表
func RouteNameList(app *iris.Application, AppName string, Port string) {
	routeList := app.GetRoutes()
	var index int = 1
	for _, value := range routeList {
		if strings.Contains(value.MainHandlerName, "tmaic") || strings.Contains(value.MainHandlerName, "iris") {
			continue
		}
		if value.Method == "POST" || value.Method == "GET" {
			console.Println(console.CODE_SUCCESS, " "+console.Sprintf(console.CODE_SUCCESS, "%-4d", index)+console.Sprintf(console.CODE_SUCCESS, "%-40s", value.MainHandlerName)+console.Sprintf(console.CODE_SUCCESS, "%-5s", value.Method)+" "+console.Sprintf(console.CODE_SUCCESS, "%-45s", value.Path)+console.Sprintf(console.CODE_SUCCESS, "%-30s", value.FormattedPath)+console.Sprintf(console.CODE_SUCCESS, "%-20s", value.RegisterFileName))
			index++
		}
	}
	console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "%s listening on: http://localhost%s", AppName, Port))
}
