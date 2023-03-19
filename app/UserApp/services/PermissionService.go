package services

import (
	"gitee.com/pangxianfei/saas/paas"
	"github.com/kataras/iris/v12"
)

var PermissionService = new(UserPermission)

type UserPermission struct {
}

func (user *UserPermission) PostApplyFor(ctx iris.Context, AppId int64) error {
	err := paas.Gate.AddPermissionToApp(ctx, AppId)

	return err
}

func (user *UserPermission) PostApplyForV2(ctx iris.Context, AppId int64) error {
	err := paas.Gate.AddPermissionToApp(ctx, AppId)

	return err
}
