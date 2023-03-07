package repositories

import (
	"gorm.io/gorm"
	SysAppModel "tmaic/app/SysApp/model"
	UserAppModel "tmaic/app/UserApp/model"
)

var AdminRepository = new(adminRepository)

type adminRepository struct {
}

func (r *adminRepository) Take(db *gorm.DB, where ...interface{}) *UserAppModel.Admin {
	ret := &UserAppModel.Admin{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
func (r *adminRepository) GetByMobile(db *gorm.DB, mobile string) *UserAppModel.Admin {
	return r.Take(db, "mobile = ?", mobile)
}

func (r *adminRepository) Create(db *gorm.DB, admin *SysAppModel.Admin) (err error) {
	err = db.Create(admin).Error
	return
}
