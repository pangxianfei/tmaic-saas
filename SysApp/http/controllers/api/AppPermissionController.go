package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
	"tmaic/SysApp/services"
)

type AppPermissionController struct {
	controller.BaseController
}

// PostApplyFor 申请应用权限
func (e *AppPermissionController) PostApplyFor() *simple.JsonResult {
	var ApplyFor requests.ApplyFor
	if err := e.Context.ReadJSON(&ApplyFor); err != nil {
		return simple.JsonErrorMsg("参数不能为空")
	}
	err := services.AppPermissionService.PostApplyFor(e.Context, ApplyFor.AppId)
	if err != nil {
		return simple.JsonData(err.Error())
	}

	return simple.JsonData(ApplyFor)
}
