package bootstrap

import "github.com/kataras/iris/v12"

type bootstrap interface {
	Initialize()
	SysRun() error
	LoginApp() error
	OrderApp() error
	ProductApp() error
	UserApp() error
	SrmApp() error
	SetAppConfig(app *iris.Application, Port string) error
	EnablingScheduledTask()
	UserRouteNameList(app *iris.Application, AppName string, Port string, AppId int64)
	SysRouteNameList(app *iris.Application, AppName string, Port string, AppId int64)
}
