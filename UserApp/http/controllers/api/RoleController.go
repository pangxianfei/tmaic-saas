package api

import (
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type RoleController struct {
	Ctx iris.Context
}

// PostStore 添加角色
func (e *RoleController) PostStore(ctx iris.Context) *response.JsonResult {
	var createRole requests.CreateRole
	err := simple.ReadJSON(ctx, &createRole)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	CreateRoleModel, roleErr := paas.Gate.AddRole(ctx, createRole)
	if roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}

	return response.JsonData(CreateRoleModel)
}

// PostUpdate 更新角色
func (e *RoleController) PostUpdate() *response.JsonResult {
	return response.JsonSuccess()
}

// PostDelete 删除角色
func (e *RoleController) PostDelete() *response.JsonResult {
	var SelectCreate requests.SelectRole
	if err := e.Ctx.ReadJSON(&SelectCreate); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.RemoveRole(e.Ctx, SelectCreate.RoleId); roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonData("删除成功")
}

// PostEdit 编辑角色
func (e *RoleController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}

// PostAny 角色权限列表
func (e *RoleController) PostAny() *response.JsonResult {
	var selectRole requests.SelectRole
	err := e.Ctx.ReadJSON(&selectRole)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}

	return response.JsonData(paas.Gate.HasAnyRole(e.Ctx, selectRole))
}

// PostRemove 删除角色中某个权限
func (e *RoleController) PostRemove() *response.JsonResult {
	var RemoveRolePermission requests.GiveRolePermission

	if err := e.Ctx.ReadJSON(&RemoveRolePermission); err != nil {
		return response.JsonErrorMsg(err.Error())
	}

	roleErr := paas.Gate.RemovePermissionToRoles(e.Ctx, RemoveRolePermission)
	if roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}

	return response.JsonData(RemoveRolePermission)
}
