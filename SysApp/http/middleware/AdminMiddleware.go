package middleware

import (
	"gitee.com/pangxianfei/library/consts"
	"gitee.com/pangxianfei/saas/paas"
	"gitee.com/pangxianfei/saas/response"
	"gitee.com/pangxianfei/saas/services"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

var Admin context.Handler

func GetAdminPermissions() {
	Admin = func(ctx iris.Context) {
		AdminInfo := services.UserTokenService.GetUserInfo(ctx)
		if AdminInfo == nil || AdminInfo.Id <= 0 {
			ctx.StatusCode(iris.StatusForbidden)
			_ = ctx.JSON(response.ErrorUnauthorized())
			return
		}
		//debug.Dd(AdminInfo)
		//debug.Dd(AdminInfo.UserType)
		if AdminInfo.Id == consts.RoleAdministrator {
			//超级管理员放行
		} else if paas.Gate.HasPermission(ctx, ctx.GetCurrentRoute().Name()) == false {
			ctx.StatusCode(iris.StatusForbidden)
			_ = ctx.JSON(response.ErrorNoHaveAuthority())
			return
		}
		ctx.Next()
	}
}
