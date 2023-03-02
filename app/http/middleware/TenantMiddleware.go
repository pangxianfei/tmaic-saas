package middleware

import (
	"gitee.com/pangxianfei/saas"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"tmaic/app/OrderApp/services"
	"tmaic/app/http/middleware/response"
)

var Tenant context.Handler

func TenantInfo() {
	Tenant = func(ctx iris.Context) {

		newUser := services.UserTokenService.GetUserInfo(ctx)
		if newUser == nil {
			_ = ctx.JSON(response.ErrorTokenInvalidation())
			return
		}
		var AppId int64 = 1
		sassDb := saas.DB.SetTenantsDb(newUser.TenantId, AppId)

		if sassDb == nil {
			_ = ctx.JSON(response.ErrorUnregisteredTenantAppDb())
			return
		}

		ctx.Values().Set("TenantId", newUser.TenantId)
		ctx.Values().Set("AppId", AppId)
		ctx.Values().Set("UserId", newUser.Id)
		ctx.Next()
	}
}
