package services

import (
	"gitee.com/pangxianfei/saas/paas"
	"github.com/kataras/iris/v12"
)

var AppPermissionService = new(AppPermission)

type AppPermission struct {
}

func (user *AppPermission) PostApplyFor(ctx iris.Context, AppId int64) error {
	err := paas.Gate.AddPermissionToApp(ctx, AppId)

	return err
}
