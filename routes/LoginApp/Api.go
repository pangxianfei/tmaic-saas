package LoginApp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/LoginApp/http/controllers/api"
)

func LoginApi(app *iris.Application) {

	mvc.Configure(app.Party("/auth"), func(m *mvc.Application) {
		m.Party("/").Handle(new(api.RegisterController))
		m.Party("/").Handle(new(api.LoginController))
	})

}
