package services

import (
	"errors"
	"gitee.com/pangxianfei/simple"
	"time"
	"tmaic/app/LoginApp/http/requests"
	SysAppModel "tmaic/app/SysApp/model"
	"tmaic/app/SysApp/repositories"
)

var TenantsInfoService = newTenantsInfoService()

func newTenantsInfoService() *tenantsInfoService {
	return &tenantsInfoService{}
}

type tenantsInfoService struct {
}

// GetByTenantName 根据租户名查找
func (s *tenantsInfoService) GetByTenantName(TenantName string) *SysAppModel.TenantsInfo {
	return repositories.TenantsRepository.GetByTenantName(simple.DB(), TenantName)
}

// Create 存入DB
func (s *tenantsInfoService) Create(RegisterInfo requests.UserRegister) (Tenants *SysAppModel.TenantsInfo, err error) {

	//保存至DB
	TenantsInfo := &SysAppModel.TenantsInfo{
		TenantName: RegisterInfo.TenantName,
		Mobile:     RegisterInfo.Mobile,
		Status:     1,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	Tenants, err = repositories.TenantsRepository.Create(simple.DB(), TenantsInfo)
	if err != nil {
		return nil, errors.New("租户信息创建失败")
	}

	return Tenants, nil
}
