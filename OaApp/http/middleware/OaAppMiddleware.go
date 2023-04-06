package middleware

import (
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var OaAppMiddleware context.Handler

func OaAppInfo() {
	OaAppMiddleware = func(ctx iris.Context) {
		middleware.TenantTenantMiddleware(ctx, config.Instance.AppNo.Oa)
	}
}
