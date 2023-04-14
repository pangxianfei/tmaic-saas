package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"

	"tmaic/UserApp/http/request"
)

type PermissionController struct {
	controller.BaseController
}

func (this *PermissionController) PostMenu() *response.JsonResult {

	var MenuRequest request.MenuRequest

	if err, returnData := this.Validation.Json(this.Context, &MenuRequest, MenuRequest.Messages()); err != nil {
		return this.JsonFail(returnData)
	}
	return this.JsonQueryData(paas.Gate.GetAppMenu(this.Context, MenuRequest.AppId))
}

// PostRole 角色分配权限
func (this *PermissionController) PostRole() *response.JsonResult {
	var RolePermission requests.RolePermission
	if err := this.Context.ReadJSON(&RolePermission); err != nil {
		return this.JsonErrorMsg("参数不能为空")
	}
	if roleErr := paas.Gate.SyncPermissionToRoles(this.Context, RolePermission); roleErr != nil {
		return this.JsonErrorMsg(roleErr.Error())
	}
	return this.JsonData("授权成功")
}

// PostSync 撤销权限、并添加新的权限
func (this *PermissionController) PostSync() *response.JsonResult {
	var PermissionArr requests.SyncPermission
	if err := this.Context.ReadJSON(&PermissionArr); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.SyncPermissionsTo(this.Context, PermissionArr); roleErr != nil {
		return this.JsonErrorMsg(roleErr.Error())
	}
	return this.JsonData("权限添加成功")
}

// PostRevokeUser 撤销用户权限
func (this *PermissionController) PostRevokeUser() *response.JsonResult {
	var SynPermission requests.SyncPermission
	if err := this.Context.ReadJSON(&SynPermission); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.RevokePermissionTo(this.Context, SynPermission); roleErr != nil {
		return this.JsonErrorMsg(roleErr.Error())
	}
	return this.JsonData("撤销权限成功")
}

// PostGiveRole 给角色添加一个权限
func (this *PermissionController) PostGiveRole() *response.JsonResult {
	var GiveRolePermission requests.GiveRolePermission
	if err := this.Context.ReadJSON(&GiveRolePermission); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveRoleToPermission(this.Context, GiveRolePermission) {
		return this.JsonData("添加成功")
	}
	return this.JsonData("添加失败")
}

// PostGiveUserRole 给用户添加角色权限
func (this *PermissionController) PostGiveUserRole() *response.JsonResult {
	var selectRole requests.SelectRole
	if err := this.Context.ReadJSON(&selectRole); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	if gErr := paas.Gate.GiveUserRolePermission(this.Context, selectRole); gErr != nil {
		return this.JsonErrorMsg(gErr.Error())
	}
	return this.JsonData("添加授权成功")
}

// PostRevokeRole 撤销角色权限
func (this *PermissionController) PostRevokeRole() *response.JsonResult {
	var selectRole requests.SelectRole
	if err := this.Context.ReadJSON(&selectRole); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RevokeRoleToPermission(this.Context, selectRole) {
		return this.JsonData("撤销角色权限成功")
	}
	return this.JsonErrorMsg("撤销角色权限失败")
}

// PostHasRoleAuthority 确定角色是否具有某种权限
func (this *PermissionController) PostHasRoleAuthority() *response.JsonResult {
	var giveRolePermission requests.GiveRolePermission
	if err := this.Context.ReadJSON(&giveRolePermission); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	if paas.Gate.HasRoleToPermission(this.Context, giveRolePermission) {
		return this.JsonData("有权限")
	}
	return this.JsonErrorMsg("无权限")
}

// PostRemoveUserRole 撤消用角色下所有权限
func (this *PermissionController) PostRemoveUserRole() *response.JsonResult {
	var SelectRole requests.SelectRole
	if err := this.Context.ReadJSON(&SelectRole); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RemoveUserRole(this.Context, SelectRole) {
		return this.JsonData("撤销成功")
	}
	return this.JsonErrorMsg("撤销失败")
}
