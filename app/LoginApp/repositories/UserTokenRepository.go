package repositories

import (
	"gitee.com/pangxianfei/simple"
	"gorm.io/gorm"
	UserAppModel "tmaic/app/UserApp/model"
)

var UserTokenRepository = new(userTokenRepository)

type userTokenRepository struct {
}

func (r *userTokenRepository) GetByToken(db *gorm.DB, token string) *UserAppModel.UserToken {
	if len(token) == 0 {
		return nil
	}
	return r.Take(db, "token = ?", token)
}

func (r *userTokenRepository) Get(db *gorm.DB, id int64) *UserAppModel.UserToken {
	ret := &UserAppModel.UserToken{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Take(db *gorm.DB, where ...interface{}) *UserAppModel.UserToken {
	ret := &UserAppModel.UserToken{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []UserAppModel.UserToken) {
	cnd.Find(db, &list)
	return
}

func (r *userTokenRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *UserAppModel.UserToken {
	ret := &UserAppModel.UserToken{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []UserAppModel.UserToken, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userTokenRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []UserAppModel.UserToken, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &UserAppModel.UserToken{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userTokenRepository) Create(db *gorm.DB, t *UserAppModel.UserToken) (err error) {
	err = db.Debug().Create(t).Error
	return
}

func (r *userTokenRepository) Update(db *gorm.DB, t *UserAppModel.UserToken) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userTokenRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&UserAppModel.UserToken{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userTokenRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&UserAppModel.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userTokenRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&UserAppModel.UserToken{}, "id = ?", id)
}
