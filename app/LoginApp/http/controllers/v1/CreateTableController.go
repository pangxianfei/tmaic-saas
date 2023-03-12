package v1

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	"gitee.com/pangxianfei/simple"
	"github.com/kataras/iris/v12"
	UserAppModel "tmaic/app/UserApp/model"
)

type CreateTableController struct {
	Ctx iris.Context
}

func (c *CreateTableController) PostCreate() *simple.JsonResult {
	db := simple.DB()

	err := db.AutoMigrate(&UserAppModel.Roles{})
	debug.Dd("数据表创建失败:", err.Error())
	return simple.JsonErrorMsg(err.Error())
}
