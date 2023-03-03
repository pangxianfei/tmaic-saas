package middleware

import (
	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/saas"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	UserAppModel "tmaic/app/UserApp/model"
	"tmaic/app/UserApp/services"
	"tmaic/app/http/middleware/response"
)

var OrderAppMiddleware context.Handler

func OrderAppInfo() {
	OrderAppMiddleware = func(ctx iris.Context) {

		var UserAppHeader response.AppHeaderAuthorization
		if err := ctx.ReadHeaders(&UserAppHeader); err != nil {
			_ = ctx.JSON(response.ErrorTokenInvalidation())
			return
		}
		TokenModel := jwt.Get(ctx).(*UserAppModel.Token)
		standardClaims := jwt.GetVerifiedToken(ctx).StandardClaims
		timeLeft := standardClaims.Timeleft()
		newUser := services.UserTokenService.GetUserInfo(ctx)
		if TokenModel == nil || newUser == nil {
			_ = ctx.JSON(response.ErrorTokenInvalidation())
			return
		}

		var AppId int64 = 3
		if sassDb := saas.DB.SetTenantsDb(TokenModel.TenantId, AppId); sassDb == nil {
			_ = ctx.JSON(response.ErrorUnregisteredTenantAppDb())
			return
		}

		ctx.Values().Set("TenantId", TokenModel.TenantId)
		ctx.Values().Set("AppId", AppId)
		ctx.Values().Set("UserId", TokenModel.UserId)

		console.Println(console.CODE_WARNING, " "+console.Sprintf(console.CODE_WARNING, "%s  %s %s", ctx.Method(), ctx.Path(), timeLeft))
		ctx.Next()
	}
}
