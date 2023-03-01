package home

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type HomeController struct {
	//上下文对象
	Ctx iris.Context

	//session对象
	Session *sessions.Session
}

func (c *HomeController) Get() mvc.View {

	c.Ctx.ViewData("datas", "comments")
	c.Ctx.ViewData("Title", "GoLang API Framework 开发框架")
	c.Ctx.ViewData("keywords", "Golang开发框架,Golang web mvc框架,GoLang API Framework 开发框架,gin框架,mvc设计模式")
	c.Ctx.ViewData("description", "tmaic 是一套简洁、优雅的Golang API Web开发框架(GoLang Web Framework)。支持mysql,mssql等多类型数据库，它可以让你从面条一样杂乱的代码中解脱出来；它可以帮你构建一个完美的网络应用，而且每行代码都可以简洁、富于表达力。")

	return mvc.View{Name: "home"}

}

func (c *HomeController) GetAbout() mvc.View {
	return mvc.View{
		Name: "about.html",
		Data: map[string]interface{}{
			"Title":   "About Page",
			"Message": "Your application description page.",
		},
	}
}

func (c *HomeController) GetContact() mvc.View {
	return mvc.View{
		Name: "contact.html",
		Data: map[string]interface{}{
			"Title":   "Contact Page",
			"Message": "Your application description page.",
		},
	}
}
