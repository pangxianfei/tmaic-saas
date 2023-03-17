package v1

import (
	"gitee.com/pangxianfei/framework/kernel/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/requests"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
)

type RegisterController struct {
	Ctx iris.Context
}

// PostRegister 注册租户及用户帐号
func (c *RegisterController) PostRegister() *JsonResult {

	var UserRegister requests.UserRegister

	if err := c.Ctx.ReadJSON(&UserRegister); err != nil {
		return JsonErrorMsg(err.Error())
	}

	//保存租户信息
	tenantsInfo, err := paas.Tenant.Add(c.Ctx, UserRegister)

	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	//创建登陆管理员账号、密码
	createUser, err := paas.Instance.CreateUser(UserRegister, tenantsInfo)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	//同步登陆账号、密码
	if createUser.Id > 0 {
		_, err := paas.Tenant.SynTenantUser(createUser)
		if err != nil {
			return JsonErrorMsg(err.Error())
		}
	}
	//创建租户数据库账号、密码
	err = paas.Instance.CreateDatabaseUserName(UserRegister, createUser)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}
	//创建舫数据库实例
	appInstance, err := paas.Instance.CreateAppInstance(UserRegister, createUser)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	return JsonData(appInstance)

}

func (c *RegisterController) PostSynTenantInfo() *JsonResult {
	return JsonData(paas.Tenant.GetTenantAdminInfo(1))
}

func (c *RegisterController) PostSynTenantUser() *JsonResult {

	admin := paas.Tenant.GetTenantAdminInfo(1)

	user, err := paas.Tenant.SynTenantUser(admin)

	return JsonData(tmaic.V{"user": user, "msg": err.Error()})
}
