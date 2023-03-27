package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"

	"tmaic/UserApp/model"
)

type SaasController struct {
	controller.BaseController
}

// PostDb 测试连接Db
func (c *SaasController) PostDb() *response.JsonResult {
	db := paas.DB.Initiation(c.Context)
	var Admin UserAppModel.TenantUser
	db.First(&Admin)
	return response.JsonData(Admin.UserName)
}
