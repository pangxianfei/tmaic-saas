package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"

	LoginAppRoute "tmaic/LoginApp/routes"
	OrderAppRoute "tmaic/OrderApp/routes"
	ProductRoute "tmaic/ProductApp/routes"
	UserAppRoute "tmaic/UserApp/routes"
)

// 登陆中心
func (s *Saas) LoginApp() error {
	app := iris.New().SetName("LoginApp")
	LoginAppRoute.Route(app)
	s.UserRouteNameList(app, config.Instance.App.LoginApp, config.Instance.AppPort.LoginPort, config.Instance.AppNo.Login)
	loginErr := s.SetAppConfig(app, config.Instance.AppPort.LoginPort)
	_ = app.Build()
	return loginErr
}

// 订单管理系统
func (s *Saas) OrderApp() error {
	app := iris.New().SetName("OrderApp")
	OrderAppRoute.Route(app)
	s.UserRouteNameList(app, config.Instance.App.OrderApp, config.Instance.AppPort.OrderPort, config.Instance.AppNo.Order)
	OrderErr := s.SetAppConfig(app, config.Instance.AppPort.OrderPort)
	_ = app.Build()
	return OrderErr
}

// 人事管理系统
func (s *Saas) UserApp() error {
	app := iris.New().SetName("UserApp")
	UserAppRoute.Route(app)
	s.UserRouteNameList(app, config.Instance.App.UserApp, config.Instance.AppPort.UserPort, config.Instance.AppNo.User)
	UserErr := s.SetAppConfig(app, config.Instance.AppPort.UserPort)
	_ = app.Build()
	return UserErr
}

// 产品管理系统
func (s *Saas) ProductApp() error {
	app := iris.New().SetName("ProductApp")
	ProductRoute.Route(app)
	s.UserRouteNameList(app, config.Instance.App.ProductApp, config.Instance.AppPort.ProductPort, config.Instance.AppNo.Product)
	UserErr := s.SetAppConfig(app, config.Instance.AppPort.ProductPort)
	_ = app.Build()
	return UserErr
}
