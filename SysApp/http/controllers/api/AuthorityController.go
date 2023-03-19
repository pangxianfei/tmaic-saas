package api

import (
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type AuthorityController struct {
	Ctx iris.Context
}

// PostStore 添加权限
func (e *AuthorityController) PostStore() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostUpdate 更新权限
func (e *AuthorityController) PostUpdate() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostDelete 删除权限
func (e *AuthorityController) PostDelete() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostShow 权限信息
func (e *AuthorityController) PostShow() *simple.JsonResult {
	return simple.JsonSuccess()
}
