package api

import (
	"gitee.com/pangxianfei/saas/paas"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	UserAppModel "tmaic/app/UserApp/model"
)

type SaasController struct {
	Ctx iris.Context
}

// PostDb 测试连接Db
func (c *SaasController) PostDb() *JsonResult {
	//TenantId := c.Ctx.Values().Get("TenantId").(int64)
	//AppId := c.Ctx.Values().Get("AppId").(int64)
	//db := saas.DB.SetTenantsDb(TenantId, AppId)
	db := paas.DB.Initiation(c.Ctx)
	var Admin UserAppModel.TenantUser
	db.First(&Admin)
	return JsonData(Admin.UserName)
}
