package repositories

import (
	"gitee.com/pangxianfei/simple"
	"gorm.io/gorm"
	UserAppModel "tmaic/app/UserApp/model"
)

var UserRepository = new(userRepository)

type userRepository struct {
}

func (r *userRepository) Get(db *gorm.DB, id int64) *UserAppModel.Admin {
	ret := &UserAppModel.Admin{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) Take(db *gorm.DB, where ...interface{}) *UserAppModel.Admin {
	ret := &UserAppModel.Admin{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []UserAppModel.Admin) {
	cnd.Find(db, &list)
	return
}

func (r *userRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *UserAppModel.Admin {
	ret := &UserAppModel.Admin{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []UserAppModel.Admin, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []UserAppModel.Admin, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &UserAppModel.Admin{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userRepository) Create(db *gorm.DB, t *UserAppModel.Admin) (err error) {
	err = db.Create(t).Error
	return
}

func (r *userRepository) Update(db *gorm.DB, t *UserAppModel.Admin) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&UserAppModel.Admin{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&UserAppModel.Admin{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&UserAppModel.Admin{}, "id = ?", id)
}

func (r *userRepository) GetByEmail(db *gorm.DB, email string) *UserAppModel.Admin {
	return r.Take(db, "email = ?", email)
}

func (r *userRepository) GetByUsername(db *gorm.DB, username string) *UserAppModel.Admin {
	return r.Take(db, "username = ?", username)
}

func (r *userRepository) GetByMobile(db *gorm.DB, mobile string) *UserAppModel.Admin {
	return r.Take(db, "mobile = ?", mobile)
}