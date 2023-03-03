package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"tmaic/app/http/middleware"
	OrderAppRoute "tmaic/routes/OrderApp"
	UserAppRoute "tmaic/routes/UserApp"
)

func OrderApp() error {
	app := iris.New().SetName("OrderApp")
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS)
	OrderAppRoute.OrderRoute(app)
	RouteNameList(app, config.Instance.App.OrderApp, config.Instance.AppPort.OrderPort)
	OrderErr := SetAppConfig(app, config.Instance.AppPort.OrderPort)
	_ = app.Build()
	return OrderErr
}

func UserApp() error {
	app := iris.New().SetName("UserApp")
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS)
	UserAppRoute.UserAppRoute(app)
	RouteNameList(app, config.Instance.App.UserApp, config.Instance.AppPort.UserPort)
	UserErr := SetAppConfig(app, config.Instance.AppPort.UserPort)
	_ = app.Build()
	return UserErr
}

func SetAppConfig(app *iris.Application, Port string) error {
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
