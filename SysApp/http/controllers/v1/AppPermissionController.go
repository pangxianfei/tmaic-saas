package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/requests"

	"tmaic/SysApp/services"
)

type AppPermissionController struct {
	controller.BaseController
}

// PostApplyFor 申请应用权限
func (e *AppPermissionController) PostApplyFor() *response.JsonResult {
	var ApplyFor requests.ApplyFor
	if err, retrueData := e.Validation.Json(e.Context, &ApplyFor, ApplyFor.Messages()); err != nil {
		return response.JsonDataError(retrueData)
	}
	if appErr := services.AppPermissionService.PostApplyFor(e.Context, ApplyFor.AppId); appErr != nil {
		return response.JsonFailData(appErr.Error())
	}
	return response.JsonQueryData("应用授权成功")
}
