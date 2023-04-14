package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/services"

	"tmaic/SysApp/http/requests"
)

type TenantController struct {
	controller.BaseController
}

// PostList 列表
func (this *TenantController) PostList() *response.JsonResult {

	return this.JsonQueryData(nil)
}

// PostCreateDbUser 创建租户数据库帐号
func (this *TenantController) PostCreateDbUser() *response.JsonResult {
	var DbUser requests.DbUserRequest
	if errReq := this.Context.ReadJSON(&DbUser); errReq != nil {
		return this.JsonErrorMsg("参数不能为空")
	}
	if err := paas.Instance.CreateDBuser(DbUser.UserId); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	return this.JsonData("创建租户数据库账号、密码成功")
}

func (this *TenantController) PostCreateTenantInstance() *response.JsonResult {
	var DbUser requests.DbUserRequest
	if errReq := this.Context.ReadJSON(&DbUser); errReq != nil {
		return this.JsonErrorMsg("参数不能为空")
	}

	if _, err := paas.Instance.CreateTenantsDatabase(DbUser.UserId); err != nil {
		return this.JsonErrorMsg(err.Error())
	}
	return this.JsonData("创建实例成功")
}

func (this *TenantController) PostSynTenantUser() *response.JsonResult {
	admin := paas.Tenant.GetTenantAdminInfo(services.UserService.GetUserId(this.Context))
	user, err := paas.Tenant.SynTenantUser(admin)
	return response.JsonData(tmaic.V{"user": user, "msg": err.Error()})
}
