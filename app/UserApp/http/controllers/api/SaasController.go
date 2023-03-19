package api

import (
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"tmaic/app/UserApp/model"
)

type SaasController struct {
	Ctx iris.Context
}

// PostDb 测试连接Db
func (c *SaasController) PostDb() *simple.JsonResult {
	db := paas.DB.Initiation(c.Ctx)
	var Admin UserAppModel.TenantUser
	db.First(&Admin)
	return simple.JsonData(Admin.UserName)
}
