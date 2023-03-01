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
		var TenantId int64 = 1
		var AppId int64 = 1
		var UserId int64 = 1
		ctx.Values().Set("TenantId", TenantId)
		ctx.Values().Set("AppId", AppId)
		ctx.Values().Set("UserId", UserId)
		//debug.Dump("租户中间件")
		ctx.Next()
	}
}
