package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"

	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"

	"tmaic/UserApp/http/request"
	"tmaic/UserApp/services"
)

type AdminController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (this *AdminController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(this.Context))
}

// PostStore 创建登陆帐号
func (this *AdminController) PostStore() *response.JsonResult {
	var UserStore request.UserRequset
	if newErr, returnData := this.Validation.Json(this.Context, &UserStore, UserStore.Messages()); newErr != nil {
		return response.JsonDataError(returnData)
	}
	if _, createErr := paas.Instance.CreateLoginAccount(this.Context, UserStore.UserName, UserStore.Mobile, UserStore.Password); createErr != nil {
		return response.JsonErrorMsg(createErr.Error())
	}
	return this.JsonCreateSucces("")
}

// PostEditBy 修改用户资料
func (this *AdminController) PostEditBy() *response.JsonResult {
	return this.JsonSuccess()
}

// PostSetUsername 设置用户名
func (this *AdminController) PostSetUsername() *response.JsonResult {
	return this.JsonSuccess()
}

// PostSetEmail 设置邮箱
func (this *AdminController) PostSetEmail() *response.JsonResult {
	return this.JsonSuccess()
}

// PostUpdatePassword 修改密码
func (this *AdminController) PostUpdatePassword() *response.JsonResult {
	var UpdatePassword request.UpdatePasswordRequest
	if err, returnData := this.Validation.Json(this.Context, &UpdatePassword, UpdatePassword.Messages()); err != nil {
		return this.JsonFail(returnData)
	}
	if updateErr := services.PlatformAdminService.UpdateAdminPassword(this.Context, UpdatePassword); updateErr != nil {
		return this.JsonFail(updateErr.Error())
	}
	return this.JsonUpdateSuccess("密码更新成功")
}

// PostUpdateStatus 启用或者禁用登陆帐号
func (this *AdminController) PostUpdateStatus() *response.JsonResult {
	var UserStatusRequset request.UserStatusRequset

	if err, returnData := this.Validation.Json(this.Context, &UserStatusRequset, UserStatusRequset.Messages()); err != nil {
		return this.JsonFail(returnData)
	}

	if updateErr := services.PlatformAdminService.PostUpdateStatus(this.Context, UserStatusRequset); updateErr != nil {
		return this.JsonFail(updateErr.Error())
	}

	return this.JsonUpdateSuccess("")
}
