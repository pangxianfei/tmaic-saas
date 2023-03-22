package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
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

	roleErr := paas.Gate.SyncPermissionToRoles(c.Context, RolePermission)
	if roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonData("授权成功")
}

// PostSync 撤销权限、并添加新的权限
func (c *PermissionController) PostSync() *response.JsonResult {
	var PermissionArr requests.SyncPermission
	err := c.Context.ReadJSON(&PermissionArr)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	roleErr := paas.Gate.SyncPermissionsTo(c.Context, PermissionArr)
	if roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonData("权限添加成功")
}

// PostRevokeUser 撤销用户权限
func (c *PermissionController) PostRevokeUser() *simple.JsonResult {
	var SynPermission requests.SyncPermission
	err := c.Context.ReadJSON(&SynPermission)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.RevokePermissionTo(c.Context, SynPermission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}
	return simple.JsonData("撤销权限成功")
}

// PostGiveRole 给角色添加一个权限
func (c *PermissionController) PostGiveRole() *response.JsonResult {
	var GiveRolePermission requests.GiveRolePermission
	err := c.Context.ReadJSON(&GiveRolePermission)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.GiveRoleToPermission(c.Context, GiveRolePermission) {
		return response.JsonData("添加成功")
	}
	return response.JsonData("添加失败")
}

// PostGiveUserRole 给用户添加角色权限
func (c *PermissionController) PostGiveUserRole(ctx iris.Context) *response.JsonResult {
	var selectRole requests.SelectRole
	err := c.Context.ReadJSON(&selectRole)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if gErr := paas.Gate.GiveUserRolePermission(ctx, selectRole); gErr != nil {

		return response.JsonErrorMsg(gErr.Error())
	}
	return response.JsonData("添加授权成功")
}

// PostRevokeRole 撤销角色权限
func (c *PermissionController) PostRevokeRole() *response.JsonResult {
	var selectRole requests.SelectRole
	err := c.Context.ReadJSON(&selectRole)
	if err != nil {
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
	err := c.Context.ReadJSON(&giveRolePermission)
	if err != nil {
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
	err := c.Context.ReadJSON(&SelectRole)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if paas.Gate.RemoveUserRole(c.Context, SelectRole) {
		return response.JsonData("撤销成功")
	}
	return response.JsonErrorMsg("撤销失败")
}
