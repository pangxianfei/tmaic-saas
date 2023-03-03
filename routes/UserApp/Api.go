package UserAppRoute

import (
	"gitee.com/pangxianfei/library/config"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/UserApp/http/controllers/api"
	UserAppMiddleware "tmaic/app/UserApp/http/middleware"
	UserAppModel "tmaic/app/UserApp/model"
	"tmaic/app/http/middleware"
	"tmaic/app/http/middleware/response"
)

func UserAppApi(app *iris.Application) {
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Party("/auth").Handle(new(api.LoginController))
	})

	auth := app.Party("/api")

	mvc.Configure(auth, func(m *mvc.Application) {
		m.Router.Use(middleware.LoginMiddleware(), UserAppMiddleware.UserAppMiddleware)
		m.Party("/user").Handle(new(api.UserController))
		m.Party("/saas").Handle(new(api.SaasController))
	})
}

func verifier() context.Handler {

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
