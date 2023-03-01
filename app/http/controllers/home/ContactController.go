package home

import "github.com/kataras/iris/v12/mvc"

type ContactController struct{}

var contactView = mvc.View{
	Name: "contact.html",
	Data: map[string]interface{}{
		"Title":   "联系",
		"Message": "我的联系方式",
	},
}

func (c *ContactController) Get() mvc.View {
	return contactView
}
