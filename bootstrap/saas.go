package bootstrap

import "github.com/kataras/iris/v12"

type bootstrap interface {
	Initialize()
	SysRun() error
	LoginApp() error
	OrderApp() error
	UserApp() error
	SetAppConfig(app *iris.Application, Port string) error
	EnablingScheduledTask()
	RouteNameList(app *iris.Application, AppName string, Port string, AppId int64)
}
