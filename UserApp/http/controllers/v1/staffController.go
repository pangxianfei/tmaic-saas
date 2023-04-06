package v1

/**
人事人员信息
*/
import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"
	"gitee.com/pangxianfei/saas/paas"
)

type StaffController struct {
	controller.BaseController
}

// PostList 列表
func (c *StaffController) PostList() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostInfo 获取当前登录用户
func (c *StaffController) PostInfo() *response.JsonResult {
	return response.JsonData(paas.Auth.User(c.Context))
}

// PostStore 添加员工资料
func (c *StaffController) PostStore() *response.JsonResult {
	return response.JsonSuccess()
}

// PostEdit 修改员工资料
func (c *StaffController) PostEdit() *response.JsonResult {
	return response.JsonSuccess()
}
