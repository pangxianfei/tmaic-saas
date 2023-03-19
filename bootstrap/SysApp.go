package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	"tmaic/routes"
	"tmaic/routes/SysApp"
)

func (s *Saas) SysRun() error {
	app := iris.New().SetName("SysApp")
	//注册资源及心跳路由
	route.Route(app)
	SysAppRoute.SysAppRoute(app)
	s.RouteNameList(app, config.Instance.App.SysApp, config.Instance.AppPort.SysAppPort, config.Instance.AppNo.Sys)
	SysErr := s.SetAppConfig(app, config.Instance.AppPort.SysAppPort)
	_ = app.Build()
	return SysErr
}
