package middleware

import (
	"gitee.com/pangxianfei/framework/helpers/debug"
	"github.com/kataras/iris/v12/middleware/jwt"
	UserAppModel "tmaic/app/UserApp/model"

	//"gitee.com/pangxianfei/saas"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	//"tmaic/app/UserApp/services"
	//"tmaic/app/http/middleware/response"
)

var Tenant context.Handler

func TenantInfo() {
	Tenant = func(ctx iris.Context) {

		claims := jwt.Get(ctx).(*UserAppModel.Token)
		standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
		expiresAtString := standardClaims.ExpiresAt().Format(ctx.Application().ConfigurationReadOnly().GetTimeFormat())
		timeLeft := standardClaims.Timeleft()
		debug.Dd(claims.Mobile)
		debug.Dd(expiresAtString)
		debug.Dd(timeLeft)

		/*
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
		*/
		ctx.Values().Set("TenantId", claims.TenantId)
		ctx.Values().Set("AppId", 1)
		ctx.Values().Set("UserId", claims.UserId)
		ctx.Next()
	}
}
