package home

import "github.com/kataras/iris/v12/mvc"

type AboutController struct{}

var aboutView = mvc.View{
	Name: "about.html",
	Data: map[string]interface{}{
		"Title":   "关于",
		"Message": "关于介绍",
	},
}

func (c *AboutController) Get() mvc.View {
	return aboutView
}
