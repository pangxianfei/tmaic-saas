package middleware

import (
	"gitee.com/pangxianfei/library/config"
	"gitee.com/pangxianfei/saas/middleware"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var NbomAppMiddleware context.Handler

func NbomAppInfo() {
	NbomAppMiddleware = func(ctx iris.Context) {
		middleware.TenantTenantMiddleware(ctx, config.Instance.AppNo.Nbom)
	}
}
