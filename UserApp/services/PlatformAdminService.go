package services

import (
	"gitee.com/pangxianfei/library/consts"
	"gitee.com/pangxianfei/saas/services"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/date"
	"github.com/kataras/iris/v12"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"tmaic/UserApp/http/request"
)

var PlatformAdminService = new(PlatformAdmin)

type PlatformAdmin struct {
}

func (admin *PlatformAdmin) AdminInfo(ctx iris.Context) *sysmdel.PlatformAdmin {
	AdminInfo := services.UserService.GetAdminInfo(ctx)
	if AdminInfo == nil && AdminInfo.Id <= 0 {
		return nil
	}
	return AdminInfo
}

func (admin *PlatformAdmin) UpdateAdminPassword(ctx iris.Context, UpdatePassword request.UpdatePasswordRequest) error {
	AdminInfo := admin.AdminInfo(ctx)

	if AdminInfo == nil && AdminInfo.Id <= 0 {
		return errors.New("用户不存在")
	}

	//检验旧密码是否正确
	if !simple.ValidatePassword(AdminInfo.Password, UpdatePassword.OldPassword) {

		return errors.New("旧密码不正确")
	}

	updateAdminPassWork := &sysmdel.PlatformAdmin{}
	TenantsErr := simple.DB().Transaction(func(txDb *gorm.DB) error {
		updateAdminPassWork.Password = simple.EncodePassword(UpdatePassword.Password)
		updateAdminPassWork.UpdateTime = date.NowTimestamp()
		if err := txDb.Debug().Model(&sysmdel.PlatformAdmin{Id: AdminInfo.Id}).Updates(updateAdminPassWork).Error; err != nil {
			return err
		}
		return nil
	})
	return TenantsErr
}

func (admin *PlatformAdmin) PostUpdateStatus(ctx iris.Context, UserStatusRequset request.UserStatusRequset) error {
	AdminInfo := admin.AdminInfo(ctx)

	if AdminInfo == nil && AdminInfo.Id <= 0 {
		return errors.New("用户不存在")
	}

	updateAdminPassWork := &sysmdel.PlatformAdmin{}
	TenantsErr := simple.DB().Transaction(func(txDb *gorm.DB) error {

		if UserStatusRequset.Status == 1 {
			updateAdminPassWork.Status = UserStatusRequset.Status
		} else {
			updateAdminPassWork.Status = consts.StatusDisable
		}
		updateAdminPassWork.UpdateTime = date.NowTimestamp()
		if err := txDb.Debug().Model(&sysmdel.PlatformAdmin{Id: AdminInfo.Id}).Updates(updateAdminPassWork).Error; err != nil {
			return err
		}
		return nil
	})
	return TenantsErr
}
