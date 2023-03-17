package api

import (
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"

	//"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type PermissionController struct {
	Ctx iris.Context
}

// PostApplyFor 申请应用权限
func (e *PermissionController) PostApplyFor() *simple.JsonResult {
	var ApplyFor requests.ApplyFor

	if err := e.Ctx.ReadJSON(&ApplyFor); err != nil {
		return simple.JsonErrorMsg("参数不能为空")
	}

	err := paas.Gate.AddPermissionToApp(e.Ctx, ApplyFor.AppId)
	if err != nil {
		return simple.JsonData(err.Error())
	}

	return simple.JsonData(ApplyFor)
}

// PostRole 角色分配权限
func (e *PermissionController) PostRole() *simple.JsonResult {
	var RolePermission requests.RolePermission

	if err := e.Ctx.ReadJSON(&RolePermission); err != nil {
		return simple.JsonErrorMsg("参数不能为空")
	}

	roleErr := paas.Gate.SyncPermissionToRoles(e.Ctx, RolePermission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}
	return simple.JsonData(RolePermission)
}

// PostSync 撤销权限、并添加新的权限
func (e *PermissionController) PostSync() *simple.JsonResult {
	var SyncPermission requests.SyncPermission
	err := e.Ctx.ReadJSON(&SyncPermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.SyncPermissionsTo(e.Ctx, SyncPermission.Permission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}

	return simple.JsonData("权限添加成功")
}

// PostRevokeUser 撤销用户权限
func (e *PermissionController) PostRevokeUser() *simple.JsonResult {
	var SyncPermission requests.SyncPermission
	err := e.Ctx.ReadJSON(&SyncPermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.RevokePermissionTo(e.Ctx, SyncPermission.Permission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}

	return simple.JsonData("撤销权限成功")
}

// PostGiveRole 给角色添加一个权限
func (e *PermissionController) PostGiveRole(ctx iris.Context) *simple.JsonResult {
	var GiveRolePermission requests.GiveRolePermission
	err := e.Ctx.ReadJSON(&GiveRolePermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveRoleToPermission(ctx, GiveRolePermission) {
		return simple.JsonData("添加成功")
	}
	return simple.JsonErrorMsg("添加失败")
}

// PostGiveUserRole 给用户添加角色权限
func (e *PermissionController) PostGiveUserRole(ctx iris.Context) *simple.JsonResult {
	var selectRole requests.SelectRole
	err := e.Ctx.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveUserRolePermission(ctx, selectRole) {
		return simple.JsonData("添加成功")
	}

	return simple.JsonErrorMsg("添加失败")
}

// PostRevokeRole 撤销角色权限
func (e *PermissionController) PostRevokeRole(ctx iris.Context) *simple.JsonResult {
	var selectRole requests.SelectRole
	err := e.Ctx.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RevokeRoleToPermission(ctx, selectRole) {
		return simple.JsonData("撤销角色权限成功")
	}
	return simple.JsonErrorMsg("撤销角色权限失败")
}

// PostHasRoleAuthority 确定角色是否具有某种权限
func (e *PermissionController) PostHasRoleAuthority(ctx iris.Context) *simple.JsonResult {
	var giveRolePermission requests.GiveRolePermission
	err := e.Ctx.ReadJSON(&giveRolePermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.HasRoleToPermission(ctx, giveRolePermission) {
		return simple.JsonData("有权限")
	}
	return simple.JsonErrorMsg("无权限")
}

// PostRemoveUserRole 确定角色是否具有某种权限
func (e *PermissionController) PostRemoveUserRole(ctx iris.Context) *simple.JsonResult {
	var selectRole requests.SelectRole
	err := e.Ctx.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RemoveUserRole(ctx, selectRole) {
		return simple.JsonData("撤销成功")
	}
	return simple.JsonErrorMsg("撤销失败")
}
