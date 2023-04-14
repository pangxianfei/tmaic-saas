package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
)

type AuthorityController struct {
	controller.BaseController
}

// PostList 权限信息
func (e *AuthorityController) PostList() *response.JsonResult {
	return response.JsonSuccess()
}

// PostInfo 权限信息
func (e *AuthorityController) PostInfo() *response.JsonResult {
	return response.JsonSuccess()
}

// PostStore 添加权限
func (e *AuthorityController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostUpdate 更新权限
func (e *AuthorityController) PostUpdate() *response.JsonResult {
	return response.JsonSuccess()
}

// PostDelete 删除权限
func (e *AuthorityController) PostDelete() *response.JsonResult {
	return response.JsonSuccess()
}
