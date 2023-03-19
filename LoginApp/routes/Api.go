package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	v12 "tmaic/LoginApp/http/controllers/v1"
)

func LoginApi(app *iris.Application) {

	mvc.Configure(app.Party("/auth"), func(m *mvc.Application) {
		m.Party("/").Handle(new(v12.RegisterController))
		m.Party("/").Handle(new(v12.LoginController))
		m.Party("/createtable").Handle(new(v12.CreateTableController))
	})

}
