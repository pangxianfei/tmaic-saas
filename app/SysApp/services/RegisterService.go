package services

import (
	"errors"
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/date"
	"gorm.io/gorm"
	"strings"
	"tmaic/app/SysApp/http/requests"
	"tmaic/app/SysApp/model"
	"tmaic/app/SysApp/repositories"
	UserAppModel "tmaic/app/UserApp/model"
	"tmaic/app/common/constants"

	"tmaic/app/common/validate"
)

var RegisterService = newRegisterService()

func newRegisterService() *registerService {
	return &registerService{}
}

type registerService struct {
}

// Register 创建租户
func (s *registerService) Register(UserRegister requests.UserRegister) (*SysAppModel.Admin, error) {
	mobile := strings.TrimSpace(UserRegister.Mobile)
	if len(mobile) == 0 {
		return nil, errors.New("手机号不能为空")
	}
	err := validate.IsPassword(UserRegister.Password, UserRegister.RePassword)
	if err != nil {
		return nil, err
	}

	//查询租户名称是否占用
	TenantsInfo := TenantsInfoService.GetByTenantName(UserRegister.TenantName)

	if TenantsInfo != nil {
		return nil, errors.New("租户名称：" + UserRegister.TenantName + " 已被占用")
	}

	// 验证手机号
	if validate.IsMobile(mobile) {
		if s.GetByMobile(mobile) != nil {
			return nil, errors.New("手机号：" + mobile + " 已被占用")
		}
	}

	//保存租户信息
	var Tenants *SysAppModel.TenantsInfo
	TenantsErr := simple.DB().Transaction(func(tx *gorm.DB) error {
		if Tenants, err = TenantsInfoService.Create(UserRegister); err != nil {
			return err
		}
		return nil
	})
	if TenantsErr != nil {
		return nil, errors.New("租户名称：" + UserRegister.TenantName + " 创建失败")
	}

	//开始注册租户
	Admin := &SysAppModel.Admin{}
	Admin.Nickname = mobile
	Admin.TenantId = Tenants.ID
	Admin.Mobile = mobile
	Admin.UserName = mobile
	Admin.Password = simple.EncodePassword(UserRegister.Password)
	Admin.Status = constants.StatusOk
	Admin.CreateTime = date.NowTimestamp()
	Admin.UpdateTime = date.NowTimestamp()

	AdminErr := simple.DB().Transaction(func(tx *gorm.DB) error {
		if err := repositories.AdminRepository.Create(tx, Admin); err != nil {
			return err
		}
		return nil
	})

	if AdminErr != nil {
		return nil, AdminErr
	}

	return Admin, nil

}

// GetByMobile 根据用户名查找
func (s *registerService) GetByMobile(mobile string) *UserAppModel.Admin {
	return repositories.AdminRepository.GetByMobile(simple.DB(), mobile)
}
