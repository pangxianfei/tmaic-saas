package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	requests2 "tmaic/UserApp/http/requests"
)

type PermissionController struct {
	controller.BaseController
}

// PostApplyFor 申请应用权限
func (c *PermissionController) PostApplyFor() *simple.JsonResult {
	var GiveRolePermission requests2.GiveRolePermission
	newErr, returnData := c.ValidateJSON(c.Context, &GiveRolePermission, GiveRolePermission.Messages())
	if newErr != nil {
		return simple.JsonError(returnData)
	}
	return simple.JsonData(GiveRolePermission)
}

/*

// PostApplyFor 申请应用权限
func (e *PermissionController) PostApplyFor() *simple.JsonResult {
	var ApplyFor requests.ApplyFor

	if err := e.Ctx.ReadJSON(&ApplyFor); err != nil {
		return simple.JsonErrorMsg("参数不能为空")
	}

	err := services.PermissionService.PostApplyFor(e.Ctx, ApplyFor.AppId)
	if err != nil {
		return simple.JsonData(err.Error())
	}

	return simple.JsonData(ApplyFor)
}
*/

// PostRole 角色分配权限
func (c *PermissionController) PostRole() *simple.JsonResult {
	var RolePermission requests.RolePermission

	if err := c.Context.ReadJSON(&RolePermission); err != nil {
		return simple.JsonErrorMsg("参数不能为空")
	}

	roleErr := paas.Gate.SyncPermissionToRoles(c.Context, RolePermission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}
	return simple.JsonData(RolePermission)
}

// PostSync 撤销权限、并添加新的权限
func (c *PermissionController) PostSync() *simple.JsonResult {
	var SyncPermission requests.SyncPermission
	err := c.Context.ReadJSON(&SyncPermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.SyncPermissionsTo(c.Context, SyncPermission.Permission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}

	return simple.JsonData("权限添加成功")
}

// PostRevokeUser 撤销用户权限
func (c *PermissionController) PostRevokeUser() *simple.JsonResult {
	var SyncPermission requests.SyncPermission
	err := c.Context.ReadJSON(&SyncPermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.RevokePermissionTo(c.Context, SyncPermission.Permission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}

	return simple.JsonData("撤销权限成功")
}

// PostGiveRole 给角色添加一个权限
func (c *PermissionController) PostGiveRole() *simple.JsonResult {
	var GiveRolePermission requests.GiveRolePermission
	err := c.Context.ReadJSON(&GiveRolePermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveRoleToPermission(c.Context, GiveRolePermission) {
		return simple.JsonData("添加成功")
	}
	return simple.JsonData("添加失败")
}

// PostGiveUserRole 给用户添加角色权限
func (c *PermissionController) PostGiveUserRole(ctx iris.Context) *simple.JsonResult {
	var selectRole requests.SelectRole
	err := c.Context.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveUserRolePermission(ctx, selectRole) {
		return simple.JsonData("添加成功")
	}

	return simple.JsonErrorMsg("添加失败")
}

// PostRevokeRole 撤销角色权限
func (c *PermissionController) PostRevokeRole() *simple.JsonResult {
	var selectRole requests.SelectRole
	err := c.Context.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RevokeRoleToPermission(c.Context, selectRole) {
		return simple.JsonData("撤销角色权限成功")
	}
	return simple.JsonErrorMsg("撤销角色权限失败")
}

// PostHasRoleAuthority 确定角色是否具有某种权限
func (c *PermissionController) PostHasRoleAuthority() *simple.JsonResult {
	var giveRolePermission requests.GiveRolePermission
	err := c.Context.ReadJSON(&giveRolePermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.HasRoleToPermission(c.Context, giveRolePermission) {
		return simple.JsonData("有权限")
	}
	return simple.JsonErrorMsg("无权限")
}

// PostRemoveUserRole 确定角色是否具有某种权限
func (c *PermissionController) PostRemoveUserRole() *simple.JsonResult {
	var selectRole requests.SelectRole
	err := c.Context.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RemoveUserRole(c.Context, selectRole) {
		return simple.JsonData("撤销成功")
	}
	return simple.JsonErrorMsg("撤销失败")
}
