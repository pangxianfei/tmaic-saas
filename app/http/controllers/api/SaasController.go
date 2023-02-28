package api

import (
	"github.com/kataras/iris/v12"
	"tmaic/vendors/framework/helpers/tmaic"

	. "tmaic/vendors/framework/simple"
)

type SaasController struct {
	Ctx iris.Context
}

// PostDb 测试连接Db
func (c *SaasController) PostDb() *JsonResult {
	TenantId := c.Ctx.Values().Get("TenantId")
	AppId := c.Ctx.Values().Get("AppId")
	return JsonData(tmaic.V{"TenantId": TenantId, "AppId": AppId})
}
