package bootstrap

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/recover"
	OrderAppRoute "tmaic/routes/OrderApp"
	UserAppRoute "tmaic/routes/UserApp"
)

func OrderApp() error {
	app := iris.New().SetName("OrderApp")
	app.Use(recover.New())
	OrderAppRoute.OrderRoute(app)
	RouteNameList(app, "订单应用路由列表:")
	OrderErr := app.Listen(":8081")
	return OrderErr
}

func UserApp() error {
	app := iris.New().SetName("UserApp")
	app.Use(recover.New())
	UserAppRoute.UserAppRoute(app)
	RouteNameList(app, "用户应用路由列表:")
	OrderErr := app.Listen(":8080")
	return OrderErr
}
