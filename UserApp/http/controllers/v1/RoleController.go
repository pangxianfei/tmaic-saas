package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
)

type RoleController struct {
	controller.BaseController
}

// PostAny 角色权限列表
func (e *RoleController) PostAny() *response.JsonResult {
	var QueryRole requests.QueryRole
	err := e.Context.ReadJSON(&QueryRole)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	return response.JsonQueryData(paas.Gate.HasAnyRole(e.Context, QueryRole))
}

// PostStore 添加角色
func (e *RoleController) PostStore() *response.JsonResult {
	var createRole requests.CreateRole

	if err := simple.ReadJSON(e.Context, &createRole); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	CreateRoleModel, roleErr := paas.Gate.AddRole(e.Context, createRole)
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
	if err := e.Context.ReadJSON(&SelectCreate); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	if roleErr := paas.Gate.RemoveRole(e.Context, SelectCreate.RoleId); roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonData("删除成功")
}

// PostEdit 编辑角色
func (e *RoleController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}

// PostRemove 删除角色中某个权限
func (e *RoleController) PostRemove() *response.JsonResult {
	var RemoveRolePermission requests.GiveRolePermission
	if err := e.Context.ReadJSON(&RemoveRolePermission); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	roleErr := paas.Gate.RemovePermissionToRoles(e.Context, RemoveRolePermission)
	if roleErr != nil {
		return response.JsonErrorMsg(roleErr.Error())
	}
	return response.JsonDeleteSuccess(RemoveRolePermission)
}
