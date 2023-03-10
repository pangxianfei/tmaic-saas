package repositories

import (
	"gorm.io/gorm"
	LoginAppModel "tmaic/app/LoginApp/model"
)

var AdminRepository = new(adminRepository)

type adminRepository struct {
}

func (r *adminRepository) Take(db *gorm.DB, where ...interface{}) *LoginAppModel.Admin {
	ret := &LoginAppModel.Admin{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
func (r *adminRepository) GetByMobile(db *gorm.DB, mobile string) *LoginAppModel.Admin {
	return r.Take(db, "mobile = ?", mobile)
}

func (r *adminRepository) Create(db *gorm.DB, admin *LoginAppModel.Admin) (err error) {
	err = db.Create(admin).Error
	return
}
