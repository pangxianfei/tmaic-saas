package api

import (
	"gitee.com/pangxianfei/saas"
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"math/rand"
	"time"
)

type SaasController struct {
	Ctx iris.Context
}

// PostDb 测试连接Db
func (c *SaasController) PostDb() *JsonResult {
	TenantId := c.Ctx.Values().Get("TenantId").(int64)
	//AppId := c.Ctx.Values().Get("AppId").(int64)
	var AppDb saas.AppDb

	rand.Seed(time.Now().Unix())
	var randNum int64
	randNum = GenerateRangeNum(1, 5)
	//db := AppDb.Initiation(c.Ctx)
	if randNum > 0 {
		AppDb.SetTenantsDb(TenantId, randNum)
	}

	//var user model.User
	//db.First(&user)
	return JsonData(randNum)

}
func GenerateRangeNum(min int64, max int64) int64 {

	randNum := rand.Int63n(max - min)
	randNum = randNum + min

	return randNum
}
