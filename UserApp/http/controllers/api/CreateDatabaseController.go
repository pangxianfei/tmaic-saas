package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/services"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"gitee.com/pangxianfei/simple"
)

type CreateDatabaseController struct {
	controller.BaseController
}

// PostCreateDbUser 创建租户数据库帐号
func (c *CreateDatabaseController) PostCreateDbUser() *simple.JsonResult {
	AdminId, _ := c.Context.User().GetRaw()
	AdminToken := AdminId.(*sysmdel.Token)
	if AdminToken == nil || AdminToken.UserId <= 0 {
		return simple.JsonErrorMsg("未授权")
	}
	err := paas.Instance.CreateDBuser(AdminToken.UserId)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData("创建租户数据库账号、密码成功")
}

func (c *CreateDatabaseController) PostCreateTenantInstance() *simple.JsonResult {
	AdminId, _ := c.Context.User().GetRaw()
	AdminToken := AdminId.(*sysmdel.Token)
	if AdminToken == nil || AdminToken.UserId <= 0 {
		return simple.JsonErrorMsg("未授权")
	}
	_, err := paas.Instance.CreateTenantsDatabase(AdminToken.UserId)
	if err != nil {
		return simple.JsonErrorMsg(err.Error())
	}
	return simple.JsonData("创建实例成功")
}

func (c *CreateDatabaseController) PostSynTenantUser() *simple.JsonResult {
	admin := paas.Tenant.GetTenantAdminInfo(services.UserService.GetUserId(c.Context))
	user, err := paas.Tenant.SynTenantUser(admin)
	return simple.JsonData(tmaic.V{"user": user, "msg": err.Error()})
}
