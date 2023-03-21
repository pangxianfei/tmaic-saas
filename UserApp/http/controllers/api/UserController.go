package api

import (
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

// PostInfo 获取当前登录用户
func (c *UserController) PostInfo() *response.JsonResult {

	return response.JsonData(paas.Auth.User(c.Ctx))
}

// PostEditBy 修改用户资料
func (c *UserController) PostEditBy() *response.JsonResult {
	return response.JsonSuccess()
}

// PostUpdateAvatar 修改头像
func (c *UserController) PostUpdateAvatar() *response.JsonResult {
	return response.JsonSuccess()
}

// PostSetUsername 设置用户名
func (c *UserController) PostSetUsername() *response.JsonResult {
	return response.JsonSuccess()
}

// PostSetEmail 设置邮箱
func (c *UserController) PostSetEmail() *response.JsonResult {

	return response.JsonSuccess()
}

// PostSetPassword 设置密码
func (c *UserController) PostSetPassword() *response.JsonResult {
	return nil
}

// PostUpdatePassword 修改密码
func (c *UserController) PostUpdatePassword() *response.JsonResult {
	return nil
}
