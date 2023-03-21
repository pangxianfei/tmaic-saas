package api

import (
	"gitee.com/pangxianfei/library/response"
	"github.com/kataras/iris/v12"
)

type AuthorityController struct {
	Ctx iris.Context
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

// PostShow 权限信息
func (e *AuthorityController) PostShow() *response.JsonResult {
	return response.JsonSuccess()
}
