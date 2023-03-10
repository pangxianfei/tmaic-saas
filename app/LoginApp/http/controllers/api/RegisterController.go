package api

import (
	. "gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	"tmaic/app/LoginApp/http/requests"
	"tmaic/app/LoginApp/services"
)

type RegisterController struct {
	Ctx iris.Context
}

// PostRegister 注册
func (c *RegisterController) PostRegister() *JsonResult {

	var UserRegister requests.UserRegister

	if err := c.Ctx.ReadJSON(&UserRegister); err != nil {
		return JsonErrorMsg(err.Error())
	}

	user, err := services.RegisterService.Register(UserRegister)
	if err != nil {
		return JsonError(&CodeError{
			Code:    401,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return JsonData(user)
}
