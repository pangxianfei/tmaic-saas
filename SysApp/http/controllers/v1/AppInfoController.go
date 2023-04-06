package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"

	"tmaic/SysApp/http/requests"
	"tmaic/SysApp/services"
)

type AppInfoController struct {
	controller.BaseController
}

// PostList 获取应用列表
func (this *AppInfoController) PostList() *response.JsonResult {
	newAppInfo := services.AppInfoService.GetByList()
	return this.JsonQueryData(newAppInfo)
}

// PostStore 创建应用
func (this *AppInfoController) PostStore() *response.JsonResult {
	var AppInfo requests.AppInfo
	if err := this.Context.ReadJSON(&AppInfo); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	newAppInfo, err := services.AppInfoService.Create(AppInfo)
	if err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	return this.JsonData(newAppInfo)
}
