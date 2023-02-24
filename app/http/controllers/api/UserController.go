package api

import (
	//"gitee.com/pangxianfei/simple"
	"tmaic/vendors/framework/simple"

	"github.com/kataras/iris/v12"
	"tmaic/app/services"
	//"tmaic/vendors/framework/helpers/cache"
	//"tmaic/vendors/framework/helpers/tmaic"
)

type UserController struct {
	Ctx iris.Context
}

// PostInfo 获取当前登录用户
func (c *UserController) PostInfo() *simple.JsonResult {
	//return simple.JsonData(services.UserTokenService.GetUserInfo(c.Ctx))
	return simple.JsonData(services.UserTokenService.GetUserInfo(c.Ctx))
}

// GetCurrent 获取当前登录用户
func (c *UserController) GetCurrent() *simple.JsonResult {
	return simple.JsonData(c.Ctx.User().GetUsername)
}

// GetBy 用户详情
func (c *UserController) GetBy(userId int64) *simple.JsonResult {

	return simple.JsonErrorMsg("用户不存在")
}

// PostEditBy 修改用户资料
func (c *UserController) PostEditBy(userId int64) *simple.JsonResult {
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

// GetMsgrecent 获取最近3条未读消息
func (c *UserController) GetMsgrecent() *simple.JsonResult {
	return nil
}

// GetMessages 用户消息
func (c *UserController) GetMessages() *simple.JsonResult {
	return nil
}

// GetNewest 最新用户
func (c *UserController) GetNewest() *simple.JsonResult {
	return simple.JsonData(nil)
}
