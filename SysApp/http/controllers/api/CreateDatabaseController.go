package api

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/services"
	"tmaic/SysApp/http/requests"
)

type CreateDatabaseController struct {
	controller.BaseController
}

// PostCreateDbUser 创建租户数据库帐号
func (c *CreateDatabaseController) PostCreateDbUser() *response.JsonResult {
	var DbUser requests.DbUserRequest
	if errReq := c.Context.ReadJSON(&DbUser); errReq != nil {
		return response.JsonErrorMsg("参数不能为空")
	}
	err := paas.Instance.CreateDBuser(DbUser.UserId)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	return response.JsonData("创建租户数据库账号、密码成功")
}

func (c *CreateDatabaseController) PostCreateTenantInstance() *response.JsonResult {
	var DbUser requests.DbUserRequest
	if errReq := c.Context.ReadJSON(&DbUser); errReq != nil {
		return response.JsonErrorMsg("参数不能为空")
	}
	_, err := paas.Instance.CreateTenantsDatabase(DbUser.UserId)
	if err != nil {
		return response.JsonErrorMsg(err.Error())
	}
	return response.JsonData("创建实例成功")
}

func (c *CreateDatabaseController) PostSynTenantUser() *response.JsonResult {
	admin := paas.Tenant.GetTenantAdminInfo(services.UserService.GetUserId(c.Context))
	user, err := paas.Tenant.SynTenantUser(admin)
	return response.JsonData(tmaic.V{"user": user, "msg": err.Error()})
}
