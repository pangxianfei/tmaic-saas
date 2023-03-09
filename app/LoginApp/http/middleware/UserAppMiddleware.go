package middleware

import (
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var UserAppMiddleware context.Handler

func UserAppInfo() {
	UserAppMiddleware = func(ctx iris.Context) {
		//console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "应用编号: %d  %s", config.Instance.AppNo.User, ctx.Method()))
		middleware.TenantTenantMiddleware(ctx, config.Instance.AppNo.User)
	}
}
