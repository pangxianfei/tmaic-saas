package v1

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
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
	debug.Dd(tenantsInfo)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}

	//创建登陆帐号
	createUser, err := paas.Instance.CreateUser(UserRegister, tenantsInfo)
	debug.Dd(createUser)
	if err != nil {
		return JsonErrorMsg(err.Error())
	}
	//创建租户数据库帐号
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
