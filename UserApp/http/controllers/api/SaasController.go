package api

import (
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
	"github.com/kataras/iris/v12"
	"tmaic/UserApp/model"
)

type SaasController struct {
	Ctx iris.Context
}

// PostDb 测试连接Db
func (c *SaasController) PostDb() *response.JsonResult {
	db := paas.DB.Initiation(c.Ctx)
	var Admin UserAppModel.TenantUser
	db.First(&Admin)
	return response.JsonData(Admin.UserName)
}
