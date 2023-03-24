package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"tmaic/UserApp/http/requests"
)

type UserController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *UserController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 创建登陆帐号
func (c *UserController) PostStore() *response.JsonResult {
	var UserStore requests.UserRequset
	if newErr, returnData := c.Validation.Json(c.Context, &UserStore, UserStore.Messages()); newErr != nil {
		return response.JsonError(returnData)
	}
	if _, createErr := paas.Instance.CreateLoginAccount(c.Context, UserStore.UserName, UserStore.Mobile, UserStore.Password); createErr != nil {
		return response.JsonErrorMsg(createErr.Error())
	}
	return response.JsonCreateData("")
}

// PostEditBy 修改用户资料
func (c *UserController) PostEditBy() *response.JsonResult {
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

// PostUpdatePassword 修改密码
func (c *UserController) PostUpdatePassword() *response.JsonResult {
	return nil
}
