package middleware

import (
	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	"tmaic/app/http/middleware/response"
	"tmaic/vendors/framework/config"
)

var JWT *jwt.Middleware

func initJWT() {
	JWT = jwt.New(jwt.Config{
		ErrorHandler: func(ctx iris.Context, err error) {
			if err == nil {
				_ = ctx.JSON(response.ErrorUnauthorized(err))
				return
			}
			ctx.StopExecution()
			ctx.StatusCode(iris.StatusUnauthorized)
			_ = ctx.JSON(response.ErrorUnauthorized(err))
			return
		},

		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetString("auth.sign_key")), nil
		},

		SigningMethod: jwt.SigningMethodHS256,
	})
}
