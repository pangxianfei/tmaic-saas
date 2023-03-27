package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	"golang.org/x/sync/errgroup"

	"tmaic/SysApp/routes"
	"tmaic/commonApp/routes"
)

var AppCore errgroup.Group

type StartApplication struct {
	Application map[string]func() error
}
type AppTypeInfo struct {
	AppName string
	Port    string
	No      int64
}

func (s *Saas) SysRun() error {
	app := iris.New().SetName("SysApp")
	//注册资源及心跳路由
	route.Route(app)
	routes.SysAppRoute(app)
	s.UserRouteNameList(app, config.Instance.App.SysApp, config.Instance.AppPort.SysPort, config.Instance.AppNo.Sys)
	SysErr := s.SetAppConfig(app, config.Instance.AppPort.SysPort)
	_ = app.Build()
	return SysErr
}
