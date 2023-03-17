package route

import (
	"gitee.com/pangxianfei/framework/trans"
	"github.com/kataras/iris/v12"
	"tmaic/routes/LoginApp"
)

// Route 注册公共路由
func Route(app *iris.Application) {
	LoginApp.RoutePing(app)
	LoginApp.RouteRedirect(app)
	LoginApp.RouteStatic(app)
}

func NotFound(ctx iris.Context) {
	ctx.JSON(iris.Map{"code": 404, "msg": trans.Get("address.not.exist")})
}

func InternalServerError(ctx iris.Context) {
	ctx.JSON(iris.Map{"code": iris.StatusInternalServerError, "msg": "Oups something went wrong, try again"})
}
