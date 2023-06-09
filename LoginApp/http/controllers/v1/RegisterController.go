package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
)

type RegisterController struct {
	controller.BaseController
}

// PostRegister 注册租户及用户帐号
func (this *RegisterController) PostRegister() *response.JsonResult {
	var UserRegister requests.UserRegister
	if err := this.Context.ReadJSON(&UserRegister); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	//保存租户信息
	tenantsInfo, err := paas.Tenant.Add(this.Context, UserRegister)
	if err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	//创建登陆管理员账号、密码
	createUser, CreateUserErr := paas.Instance.CreateUser(UserRegister, tenantsInfo)
	if CreateUserErr != nil {
		return response.JsonErrorMsg(CreateUserErr.Error())
	}
	//创建舫数据库实例
	appInstance, InstanceErr := paas.Instance.CreateAppInstance(UserRegister, createUser)
	if InstanceErr != nil {
		return response.JsonErrorMsg(InstanceErr.Error())
	}
	//同步登陆账号、密码
	if createUser.Id > 0 && createUser.TenantId > 0 {
		_, err := paas.Tenant.SynTenantUser(createUser)
		if err != nil {
			return response.JsonErrorMsg(err.Error())
		}
	}
	return response.JsonData(appInstance)
}
func (this *RegisterController) PostRegisterTest() *response.JsonResult {
	var UserRegister requests.UserRegister
	if err := this.Context.ReadJSON(&UserRegister); err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	//保存租户信息
	tenantsInfo, err := paas.Tenant.Add(this.Context, UserRegister)
	if err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	//创建登陆管理员账号、密码
	createUser, CreateUserErr := paas.Instance.CreateUser(UserRegister, tenantsInfo)
	if CreateUserErr != nil {
		return this.JsonErrorMsg(CreateUserErr.Error())
	}
	//创建租户数据库账号、密码
	if cErr := paas.Instance.CreateDatabaseUserName(UserRegister, createUser); cErr != nil {
		return this.JsonErrorMsg(cErr.Error())
	}
	return this.JsonCreateSucces(createUser)
}
