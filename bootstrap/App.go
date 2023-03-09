package bootstrap

import (
	"fmt"
	"strings"
	"tmaic/app/http/middleware"

	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/framework/helpers/tmaic"
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/simple"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"

	SysAppModel "tmaic/app/SysApp/model"
	LoginAppRoute "tmaic/routes/LoginApp"
	OrderAppRoute "tmaic/routes/OrderApp"
	UserAppRoute "tmaic/routes/UserApp"
)

func LoginApp() error {
	app := iris.New().SetName("LoginApp")
	LoginAppRoute.LoginAppRoute(app)
	RouteNameList(app, config.Instance.App.LoginApp, config.Instance.AppPort.LoginPort, config.Instance.AppNo.Login)
	loginErr := SetAppConfig(app, config.Instance.AppPort.LoginPort)
	_ = app.Build()
	return loginErr
}

func OrderApp() error {
	app := iris.New().SetName("OrderApp")
	OrderAppRoute.OrderRoute(app)
	RouteNameList(app, config.Instance.App.OrderApp, config.Instance.AppPort.OrderPort, config.Instance.AppNo.Order)
	OrderErr := SetAppConfig(app, config.Instance.AppPort.OrderPort)
	_ = app.Build()
	return OrderErr
}

func UserApp() error {
	app := iris.New().SetName("UserApp")
	UserAppRoute.UserAppRoute(app)
	RouteNameList(app, config.Instance.App.UserApp, config.Instance.AppPort.UserPort, config.Instance.AppNo.User)
	UserErr := SetAppConfig(app, config.Instance.AppPort.UserPort)
	_ = app.Build()
	return UserErr
}

func SetAppConfig(app *iris.Application, Port string) error {
	app.Validator = validator.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.AllowMethods(iris.MethodOptions)
	app.Use(middleware.CORS, middleware.LoginMiddleware())

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
func RouteNameList(app *iris.Application, AppName string, Port string, AppId int64) {
	routeList := app.GetRoutes()
	var index int = 1
	db := simple.DB()

	for _, value := range routeList {

		if strings.Contains(value.MainHandlerName, "tmaic") || strings.Contains(value.MainHandlerName, "iris") {
			continue
		}

		if value.Method == "POST" || value.Method == "GET" {

			AuthorityMain := SysAppModel.Authority{
				Description:    value.Description,
				Pid:            0,
				AppId:          AppId,
				SourceFileName: value.SourceFileName,
				Name:           value.MainHandlerName,
			}
			db.Model(&SysAppModel.Authority{}).FirstOrCreate(&AuthorityMain, SysAppModel.Authority{Description: value.Description, Pid: 0, SourceFileName: value.SourceFileName})

			//debug.Dd(value.StaticPath())
			//debug.Dd(value.Title)
			//debug.Dd(value.Description)
			//debug.Dd(value.Name)
			//debug.Dd(value.SourceFileName)
			console.Println(console.CODE_SUCCESS, " "+console.Sprintf(console.CODE_SUCCESS, "%-4d", index)+console.Sprintf(console.CODE_SUCCESS, "%-40s", value.MainHandlerName)+console.Sprintf(console.CODE_SUCCESS, "%-5s", value.Method)+" "+console.Sprintf(console.CODE_SUCCESS, "%-45s", value.Path)+console.Sprintf(console.CODE_SUCCESS, "%-30s", value.FormattedPath)+console.Sprintf(console.CODE_SUCCESS, "%-20s", value.RegisterFileName))

			/*
				checkAuthority := &SysAppModel.Authority{
					Description:    value.Description,
					Pid:            0,
					SourceFileName: value.SourceFileName,
				}
			*/

			var Authority SysAppModel.Authority
			var isAuthority SysAppModel.Authority

			db.Where(&SysAppModel.Authority{Description: value.Description}).First(&Authority)

			if Authority.Id > 0 {
				Md5ValueFile := fmt.Sprintf("%s_%s", value.MainHandlerName, value.RegisterFileName)
				Md5Value := tmaic.MD5(Md5ValueFile)
				db.Where("pid > ? and md5_value = ?", 0, tmaic.MD5(Md5Value)).Find(&isAuthority)
				//debug.Dd(Authority)
				if isAuthority.Id <= 0 {
					createAuthority := &SysAppModel.Authority{
						Pid:              Authority.Id,
						AppId:            AppId,
						Name:             str(value.MainHandlerName, "."),
						Description:      value.Description,
						RegisterFileName: value.RegisterFileName,
						MainHandlerName:  value.MainHandlerName,
						Method:           value.Method,
						FormattedPath:    value.FormattedPath,
						StaticPath:       value.StaticPath(),
						Path:             value.Path,
						SourceFileName:   value.SourceFileName,
						RouteName:        value.Name,
						Status:           1,
						Md5Value:         tmaic.MD5(Md5Value),
					}
					db.Create(createAuthority)
				}

			}

			index++
		}
	}
	console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "%s listening on: http://localhost%s", AppName, Port))
}

func str(str string, point string) string {

	words := strings.SplitAfter(str, point)
	if len(words) > 0 {
		return words[len(words)-1]
	}
	return ""
}
