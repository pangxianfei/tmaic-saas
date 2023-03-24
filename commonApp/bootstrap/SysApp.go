package bootstrap

import (
	"fmt"
	"time"

	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/services"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
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
	s.UserRouteNameList(app, config.Instance.App.SysApp, config.Instance.AppPort.SysAppPort, config.Instance.AppNo.Sys)
	SysErr := s.SetAppConfig(app, config.Instance.AppPort.SysAppPort)
	_ = app.Build()
	return SysErr
}

func (s *Saas) ApplicationExample(AppInfo AppTypeInfo) func() error {
	return func() error {
		app := iris.New().SetName(AppInfo.AppName)
		route.Route(app)
		s.UserRouteNameList(app, AppInfo.AppName, AppInfo.Port, AppInfo.No)
		err := s.SetAppConfig(app, AppInfo.Port)
		_ = app.Build()
		return err
	}
}

func (s *Saas) StartApplication() {
	var appList []sysmdel.AppInfo
	appList = services.AppInfoService.GetStartApplication()
	for _, AppInfo := range appList {
		var sys AppTypeInfo
		sys.AppName = AppInfo.AppName
		sys.Port = fmt.Sprintf(":%s", AppInfo.Port)
		sys.No = AppInfo.No
		AppCore.Go(s.ApplicationExample(sys))
		time.Sleep(time.Millisecond * 1000)
	}
}
