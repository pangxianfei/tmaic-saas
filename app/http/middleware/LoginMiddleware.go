package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"tmaic/app/http/middleware/response"
	"tmaic/app/services"
	"tmaic/vendors/framework/helpers/debug"
)

// Login 获取JWT中的用户数据，封装成实体存放在ctx中供请求调用
var Login context.Handler

func initUserInfo() {
	Login = func(ctx iris.Context) {
		newUser := services.UserTokenService.GetUserInfo(ctx)
		//debug.Dump("获取JWT中的用户数据，封装成实体存放在ctx中供请求调用")
		if newUser == nil {
			_ = ctx.JSON(response.ErrorTokenInvalidation())
			return
		}
		debug.Dump(newUser)
		ctx.Values().Set("UserInfo", newUser)
		ctx.Next()
	}
}
