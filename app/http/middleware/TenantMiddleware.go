package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var Tenant context.Handler

func TenantInfo() {
	Tenant = func(ctx iris.Context) {
		/*
			newUser := services.UserTokenService.GetUserInfo(ctx)
			if newUser == nil {
				_ = ctx.JSON(response.ErrorTokenInvalidation())
				return
			}*/
		ctx.Values().Set("TenantId", 1)
		ctx.Values().Set("AppId", 2)
		//debug.Dump("租户中间件")
		ctx.Next()
	}
}
