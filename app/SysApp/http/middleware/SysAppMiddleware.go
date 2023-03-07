package middleware

import (
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var SysAppMiddleware context.Handler

func SysAppInfo() {
	SysAppMiddleware = func(ctx iris.Context) {
		//console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "应用编号: %d  %s", config.Instance.AppNo.Sys, ctx.Method()))
		middleware.TenantTenantMiddleware(ctx, config.Instance.AppNo.Sys)

	}
}
