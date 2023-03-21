package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	"gitee.com/pangxianfei/simple"
)

type RegisterController struct {
	controller.BaseController
}

// PostRegister 注册租户及用户帐号
func (c *RegisterController) PostRegister() *simple.JsonResult {
	var UserRegister requests.UserRegister
	if err := c.Context.ReadJSON(&UserRegister); err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	//保存租户信息
	tenantsInfo, err := paas.Tenant.Add(c.Context, UserRegister)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	//创建登陆管理员账号、密码
	createUser, CreateUserErr := paas.Instance.CreateUser(UserRegister, tenantsInfo)
	if CreateUserErr != nil {
		return simple.JsonErrorMsg(CreateUserErr.Error())
	}
	//同步登陆账号、密码
	if createUser.Id > 0 {
		_, err := paas.Tenant.SynTenantUser(createUser)
		if err != nil {
			return simple.JsonErrorMsg(err.Error())
		}
	}
	//创建租户数据库账号、密码
	err = paas.Instance.CreateDatabaseUserName(UserRegister, createUser)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	//创建舫数据库实例
	appInstance, InstanceErr := paas.Instance.CreateAppInstance(UserRegister, createUser)
	if InstanceErr != nil {
		return simple.JsonErrorMsg(InstanceErr.Error())
	}
	return simple.JsonData(appInstance)
}