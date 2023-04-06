package v1

import (
	"gitee.com/pangxianfei/framework/http/controller"
	"gitee.com/pangxianfei/library/response"

	"tmaic/FmsApp/mq"
)

type EventsController struct {
	controller.BaseController
}

// PostInfo 队列
func (c *EventsController) PostInfo() *response.JsonResult {
	MQ.FmsDemoMessage.Dispatch()
	return response.JsonSuccess()
}
