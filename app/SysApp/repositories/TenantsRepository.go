package repositories

import (
	"gorm.io/gorm"
	SysAppModel "tmaic/app/SysApp/model"
)

var TenantsRepository = new(tenantsRepository)

type tenantsRepository struct {
}

func (r *tenantsRepository) Take(db *gorm.DB, where ...interface{}) *SysAppModel.TenantsInfo {
	ret := &SysAppModel.TenantsInfo{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *tenantsRepository) GetByMobile(db *gorm.DB, mobile string) *SysAppModel.TenantsInfo {
	return r.Take(db, "mobile = ?", mobile)
}

func (r *tenantsRepository) GetByTenantName(db *gorm.DB, TenantName string) *SysAppModel.TenantsInfo {
	return r.Take(db, "tenant_name = ?", TenantName)

}

func (r *tenantsRepository) Create(db *gorm.DB, TenantsInfo *SysAppModel.TenantsInfo) (Tenants *SysAppModel.TenantsInfo, err error) {
	var createStatus bool = false
	createStatus = db.Debug().Migrator().HasTable(&SysAppModel.TenantsInfo{})
	if createStatus == false {
		err := db.Migrator().CreateTable(&SysAppModel.TenantsInfo{})
		if err != nil {
			return nil, err
		}
	}
	err = db.Create(TenantsInfo).Error
	return TenantsInfo, err
}
