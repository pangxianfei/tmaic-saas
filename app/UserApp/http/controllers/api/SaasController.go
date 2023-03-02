package api

import (
	"gitee.com/pangxianfei/saas"
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
	db := saas.DB.Initiation(c.Ctx)
	var user UserAppModel.User
	db.First(&user)
	return JsonData(user.UserName)
}
