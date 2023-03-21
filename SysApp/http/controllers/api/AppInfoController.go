package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"tmaic/SysApp/http/requests"
	"tmaic/SysApp/services"
)

type AppInfoController struct {
	controller.BaseController
}

// PostStore 创建应用
func (c *AppInfoController) PostStore() *response.JsonResult {
	var AppInfo requests.AppInfo
	if err := c.Context.ReadJSON(&AppInfo); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	newAppInfo, err := services.AppInfoService.Create(AppInfo)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	return response.JsonData(newAppInfo)
}

// PostList 获取应用列表
func (c *AppInfoController) PostList() *response.JsonResult {
	newAppInfo := services.AppInfoService.GetByList()
	return response.JsonData(newAppInfo)
}
