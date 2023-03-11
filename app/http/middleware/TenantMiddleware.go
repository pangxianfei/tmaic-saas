package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	LoginAppModel "tmaic/app/LoginApp/model"
)

var Tenant context.Handler

func TenantInfo() {
	Tenant = func(ctx iris.Context) {

		claims := jwt.Get(ctx).(*LoginAppModel.Token)
		ctx.Values().Set("TenantId", claims.TenantId)
		ctx.Values().Set("AppId", 1)
		ctx.Values().Set("UserId", claims.UserId)
		ctx.Next()
	}
}
