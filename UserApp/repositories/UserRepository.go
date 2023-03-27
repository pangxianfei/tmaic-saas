package repositories

import (
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/sqlcmd"
	"gorm.io/gorm"

	"tmaic/UserApp/models"
)

var UserRepository = new(userRepository)

type userRepository struct {
}

func (r *userRepository) Get(db *gorm.DB, id int64) *models.TenantUser {
	ret := &models.TenantUser{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) Take(db *gorm.DB, where ...interface{}) *models.TenantUser {
	ret := &models.TenantUser{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) Find(db *gorm.DB, cnd *sqlcmd.Cnd) (list []models.TenantUser) {
	cnd.Find(db, &list)
	return
}

func (r *userRepository) FindOne(db *gorm.DB, cnd *sqlcmd.Cnd) *models.TenantUser {
	ret := &models.TenantUser{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []models.TenantUser, paging *sqlcmd.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userRepository) FindPageByCnd(db *gorm.DB, cnd *sqlcmd.Cnd) (list []models.TenantUser, paging *sqlcmd.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &models.TenantUser{})

	paging = &sqlcmd.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userRepository) Create(db *gorm.DB, t *models.TenantUser) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userRepository) Update(db *gorm.DB, t *models.TenantUser) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&models.TenantUser{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&models.TenantUser{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&models.TenantUser{}, "id = ?", id)
}

func (r *userRepository) GetByEmail(db *gorm.DB, email string) *models.TenantUser {
	return r.Take(db, "email = ?", email)
}

func (r *userRepository) GetByUsername(db *gorm.DB, username string) *models.TenantUser {
	return r.Take(db, "username = ?", username)
}

func (r *userRepository) GetByMobile(db *gorm.DB, mobile string) *models.TenantUser {
	return r.Take(db, "mobile = ?", mobile)
}
