package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type PdController struct {
	controller.BaseController
}

// PostInfo 获取当前登录用户
func (c *PdController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 创建
func (c *PdController) PostStore() *response.JsonResult {

	return response.JsonCreateSucces("")
}

// PostEdit 修改用户资料
func (c *PdController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}

// PostUpdate 更新
func (c *PdController) PostUpdate() *response.JsonResult {
	return response.JsonSuccess()
}
