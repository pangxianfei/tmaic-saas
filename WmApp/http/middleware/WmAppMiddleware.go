package middleware

import (
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var WmAppMiddleware context.Handler

func WmAppInfo() {
	WmAppMiddleware = func(ctx iris.Context) {
		middleware.TenantTenantMiddleware(ctx, config.Instance.AppNo.Wm)
	}
}
