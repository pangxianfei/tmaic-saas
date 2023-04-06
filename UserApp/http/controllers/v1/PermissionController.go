package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
)

type PermissionController struct {
	controller.BaseController
}

// PostRole 角色分配权限
func (c *PermissionController) PostRole() *response.JsonResult {
	var RolePermission requests.RolePermission
	if err := c.Context.ReadJSON(&RolePermission); err != nil {
		return response.JsonErrorMsg("参数不能为空")
	}
	if roleErr := paas.Gate.SyncPermissionToRoles(c.Context, RolePermission); roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonData("授权成功")
}

// PostSync 撤销权限、并添加新的权限
func (c *PermissionController) PostSync() *response.JsonResult {
	var PermissionArr requests.SyncPermission
	if err := c.Context.ReadJSON(&PermissionArr); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.SyncPermissionsTo(c.Context, PermissionArr); roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonData("权限添加成功")
}

// PostRevokeUser 撤销用户权限
func (c *PermissionController) PostRevokeUser() *simple.JsonResult {
	var SynPermission requests.SyncPermission
	if err := c.Context.ReadJSON(&SynPermission); err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.RevokePermissionTo(c.Context, SynPermission); roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}
	return simple.JsonData("撤销权限成功")
}

// PostGiveRole 给角色添加一个权限
func (c *PermissionController) PostGiveRole() *response.JsonResult {
	var GiveRolePermission requests.GiveRolePermission
	if err := c.Context.ReadJSON(&GiveRolePermission); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveRoleToPermission(c.Context, GiveRolePermission) {
		return response.JsonData("添加成功")
	}
	return response.JsonData("添加失败")
}

// PostGiveUserRole 给用户添加角色权限
func (c *PermissionController) PostGiveUserRole() *response.JsonResult {
	var selectRole requests.SelectRole
	if err := c.Context.ReadJSON(&selectRole); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if gErr := paas.Gate.GiveUserRolePermission(c.Context, selectRole); gErr != nil {
		return response.JsonErrorMsg(gErr.Error())
	}
	return response.JsonData("添加授权成功")
}

// PostRevokeRole 撤销角色权限
func (c *PermissionController) PostRevokeRole() *response.JsonResult {
	var selectRole requests.SelectRole
	if err := c.Context.ReadJSON(&selectRole); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RevokeRoleToPermission(c.Context, selectRole) {
		return response.JsonData("撤销角色权限成功")
	}
	return response.JsonErrorMsg("撤销角色权限失败")
}

// PostHasRoleAuthority 确定角色是否具有某种权限
func (c *PermissionController) PostHasRoleAuthority() *response.JsonResult {
	var giveRolePermission requests.GiveRolePermission
	if err := c.Context.ReadJSON(&giveRolePermission); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.HasRoleToPermission(c.Context, giveRolePermission) {
		return response.JsonData("有权限")
	}
	return response.JsonErrorMsg("无权限")
}

// PostRemoveUserRole 撤消用角色下所有权限
func (c *PermissionController) PostRemoveUserRole() *response.JsonResult {
	var SelectRole requests.SelectRole
	if err := c.Context.ReadJSON(&SelectRole); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RemoveUserRole(c.Context, SelectRole) {
		return response.JsonData("撤销成功")
	}
	return response.JsonErrorMsg("撤销失败")
}
