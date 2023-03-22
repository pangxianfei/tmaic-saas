package bootstrap

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	LoginAppRoute "tmaic/LoginApp/routes"
	OrderAppRoute "tmaic/OrderApp/routes"
	ProductRoute "tmaic/ProductApp/routes"
	UserAppRoute "tmaic/UserApp/routes"
)

func (s *Saas) LoginApp() error {
	app := iris.New().SetName("LoginApp")
	LoginAppRoute.LoginAppRoute(app)
	s.UserRouteNameList(app, config.Instance.App.LoginApp, config.Instance.AppPort.LoginPort, config.Instance.AppNo.Login)
	loginErr := s.SetAppConfig(app, config.Instance.AppPort.LoginPort)
	_ = app.Build()
	return loginErr
}

func (s *Saas) OrderApp() error {
	app := iris.New().SetName("OrderApp")
	OrderAppRoute.OrderRoute(app)
	s.UserRouteNameList(app, config.Instance.App.OrderApp, config.Instance.AppPort.OrderPort, config.Instance.AppNo.Order)
	OrderErr := s.SetAppConfig(app, config.Instance.AppPort.OrderPort)
	_ = app.Build()
	return OrderErr
}

func (s *Saas) UserApp() error {
	app := iris.New().SetName("UserApp")
	UserAppRoute.UserAppRoute(app)
	s.UserRouteNameList(app, config.Instance.App.UserApp, config.Instance.AppPort.UserPort, config.Instance.AppNo.User)
	UserErr := s.SetAppConfig(app, config.Instance.AppPort.UserPort)
	_ = app.Build()
	return UserErr
}
func (s *Saas) ProductApp() error {
	app := iris.New().SetName("ProductApp")
	ProductRoute.ProductRoute(app)
	s.UserRouteNameList(app, config.Instance.App.ProductApp, config.Instance.AppPort.ProductPort, config.Instance.AppNo.Product)
	UserErr := s.SetAppConfig(app, config.Instance.AppPort.ProductPort)
	_ = app.Build()
	return UserErr
}
