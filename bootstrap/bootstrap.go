package bootstrap

import (
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	"tmaic/app/http/middleware"
	WrongRoute "tmaic/routes"
)

type Saas struct{}

func (s *Saas) SetAppConfig(app *iris.Application, Port string) error {
	app.Validator = validator.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS)
	app.OnErrorCode(iris.StatusNotFound, WrongRoute.NotFound)
	app.OnErrorCode(iris.StatusInternalServerError, WrongRoute.InternalServerError)
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
