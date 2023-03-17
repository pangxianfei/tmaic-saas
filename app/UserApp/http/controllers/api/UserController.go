package api

import (
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type UserController struct {
	Ctx iris.Context
}

// PostInfo 获取当前登录用户
func (c *UserController) PostInfo() *simple.JsonResult {

	return simple.JsonData(paas.Auth.User(c.Ctx))
}

// PostEditBy 修改用户资料
func (c *UserController) PostEditBy() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostUpdateAvatar 修改头像
func (c *UserController) PostUpdateAvatar() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostSetUsername 设置用户名
func (c *UserController) PostSetUsername() *simple.JsonResult {
	return simple.JsonSuccess()
}

// PostSetEmail 设置邮箱
func (c *UserController) PostSetEmail() *simple.JsonResult {

	return simple.JsonSuccess()
}

// PostSetPassword 设置密码
func (c *UserController) PostSetPassword() *simple.JsonResult {
	return nil
}

// PostUpdatePassword 修改密码
func (c *UserController) PostUpdatePassword() *simple.JsonResult {
	return nil
}
