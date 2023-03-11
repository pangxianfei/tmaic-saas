package LoginApp

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"tmaic/app/LoginApp/http/controllers/v1"
)

func LoginApi(app *iris.Application) {

	mvc.Configure(app.Party("/auth"), func(m *mvc.Application) {
		m.Party("/").Handle(new(v1.RegisterController))
		m.Party("/").Handle(new(v1.LoginController))
		m.Party("/createtable").Handle(new(v1.CreateTableController))
	})

}
