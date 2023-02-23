package home

import (
	"github.com/kataras/iris/v12/mvc"
)

type IndexController struct{}

func (c *IndexController) Get() mvc.View {
	return mvc.View{
		Name: "home.html",
		Data: map[string]interface{}{
			"Title":   "index",
			"Message": "Your application description page.",
		},
	}
}
