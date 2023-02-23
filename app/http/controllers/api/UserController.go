package api

import (
	"gitee.com/pangxianfei/simple"
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
	//jwtInfo := c.Ctx.Values().Get("jwt")
	//return simple.JsonData(jwtInfo)
	//user := services.UserTokenService.GetUserInfo(c.Ctx)
	//return simple.JsonData(user)
	//logined := c.Ctx.Values().Get("UserInfo").(*model.User)
	//data, _ := json.Marshal(logined)
	//cache.Put("pangxianfei", data)
	//cache.SetNx("test", "pangxianfei")
	return simple.JsonData(services.UserTokenService.GetUserInfo(c.Ctx))
	//return simple.JsonData(logined.Email)

}

// GetCurrent 获取当前登录用户
func (c *UserController) GetCurrent() *simple.JsonResult {

	/////user := services.UserTokenService.GetCurrent(c.Ctx)

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

// 设置用户名
func (c *UserController) PostSetUsername() *simple.JsonResult {

	return simple.JsonSuccess()
}

// 设置邮箱
func (c *UserController) PostSetEmail() *simple.JsonResult {

	return simple.JsonSuccess()
}

// 设置密码
func (c *UserController) PostSetPassword() *simple.JsonResult {
	return nil
}

// 修改密码
func (c *UserController) PostUpdatePassword() *simple.JsonResult {
	return nil
}

// 设置背景图
func (c *UserController) PostSetBackgroundImage() *simple.JsonResult {
	return nil
}

// 用户收藏
func (c *UserController) GetFavorites() *simple.JsonResult {
	return nil
}

// 获取最近3条未读消息
func (c *UserController) GetMsgrecent() *simple.JsonResult {
	return nil
}

// 用户消息
func (c *UserController) GetMessages() *simple.JsonResult {
	return nil
}

// 最新用户
func (c *UserController) GetNewest() *simple.JsonResult {

	return simple.JsonData(nil)
}

// 用户积分记录
func (c *UserController) GetScorelogs() *simple.JsonResult {
	return nil
}

// 积分排行
func (c *UserController) GetScoreRank() *simple.JsonResult {
	return nil
}
