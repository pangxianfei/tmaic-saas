package api

import (
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type RoleController struct {
	Ctx iris.Context
}

// PostStore 添加角色
func (e *RoleController) PostStore(ctx iris.Context) *simple.JsonResult {
	var createRole requests.CreateRole
	err := simple.ReadJSON(ctx, &createRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	CreateRoleModel, roleErr := paas.Gate.AddRole(ctx, createRole)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}

	return simple.JsonData(CreateRoleModel)
}

// PostUpdate 更新角色
func (e *RoleController) PostUpdate() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostDelete 删除角色
func (e *RoleController) PostDelete() *simple.JsonResult {
	var SelectCreate requests.SelectRole
	if err := e.Ctx.ReadJSON(&SelectCreate); err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.RemoveRole(e.Ctx, SelectCreate.RoleId); roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}
	return simple.JsonData("删除成功")
}

// PostEdit 编辑角色
func (e *RoleController) PostEdit() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostAny 角色权限列表
func (e *RoleController) PostAny() *simple.JsonResult {
	var selectRole requests.SelectRole
	err := e.Ctx.ReadJSON(&selectRole)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData(paas.Gate.HasAnyRole(e.Ctx, selectRole))
}

// PostRemove 删除角色中某个权限
func (e *RoleController) PostRemove() *simple.JsonResult {
	var RemoveRolePermission requests.GiveRolePermission

	if err := e.Ctx.ReadJSON(&RemoveRolePermission); err != nil {
		return simple.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.RemovePermissionToRoles(e.Ctx, RemoveRolePermission)
	if roleErr != nil {
		return simple.JsonErrorMsg(roleErr.Error())
	}

	return simple.JsonData(RemoveRolePermission)
}
