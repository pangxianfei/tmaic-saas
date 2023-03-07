package api

import (
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"tmaic/app/SysApp/http/requests"
	"tmaic/app/SysApp/services"
)

type AppInfoController struct {
	Ctx iris.Context
}

// PostStore 创建应用
func (c *AppInfoController) PostStore() *JsonResult {
	var AppInfo requests.AppInfo
	if err := c.Ctx.ReadJSON(&AppInfo); err != nil {
		return JsonErrorMsg(err.Error())
	}
	newAppInfo, err := services.AppInfoService.Create(AppInfo)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}
	return JsonData(newAppInfo)
}

// PostList 获取应用列表
func (c *AppInfoController) PostList() *JsonResult {
	newAppInfo := services.AppInfoService.GetByList()
	return JsonData(newAppInfo)
}
