package middleware

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	UserAppModel "tmaic/app/UserApp/model"
	"tmaic/app/http/middleware/response"
)

// LoginMiddleware 检查登陆验证器
func LoginMiddleware() context.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, []byte(config.GetString("auth.sign_key")))
	verifier.WithDefaultBlocklist()
	verifier.ErrorHandler = func(ctx *context.Context, err error) {
		_ = ctx.JSON(response.ErrorTokenInvalidation())
		return
	}
	return verifier.Verify(func() interface{} {
		return new(UserAppModel.Token)
	})
}
