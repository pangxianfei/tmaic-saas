package routes

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	"tmaic/LoginApp/http/controllers/v1"
)

func Route(app *iris.Application) {

	mvc.Configure(app.Party("/auth"), func(m *mvc.Application) {
		m.Party("/").Handle(new(v1.RegisterController))
		m.Party("/").Handle(new(v1.LoginController))
		m.Party("/createtable").Handle(new(v1.CreateTableController))
		m.Party("/file").Handle(new(v1.FileController))
	})

}
