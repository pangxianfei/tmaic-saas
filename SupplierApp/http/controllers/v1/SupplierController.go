package v1

import (
	"gitee.com/pangxianfei/framework/http/controller/iriscontroller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type SupplierController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *SupplierController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}
