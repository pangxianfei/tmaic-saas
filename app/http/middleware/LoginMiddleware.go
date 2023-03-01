package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"tmaic/app/http/middleware/response"
	"tmaic/app/services"
)

var Login context.Handler

func initUserInfo() {
	Login = func(ctx iris.Context) {
		newUser := services.UserTokenService.GetUserInfo(ctx)
		if newUser == nil {
			_ = ctx.JSON(response.ErrorTokenInvalidation())
			return
		}
		ctx.Values().Set("UserInfo", newUser)
		ctx.Next()
	}
}