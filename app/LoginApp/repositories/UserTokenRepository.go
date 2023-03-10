package repositories

import (
	"gitee.com/pangxianfei/simple"
	"gorm.io/gorm"
	LoginAppModel "tmaic/app/LoginApp/model"
)

var UserTokenRepository = new(userTokenRepository)

type userTokenRepository struct {
}

func (r *userTokenRepository) GetByToken(db *gorm.DB, token string) *LoginAppModel.UserToken {
	if len(token) == 0 {
		return nil
	}
	return r.Take(db, "token = ?", token)
}

func (r *userTokenRepository) Get(db *gorm.DB, id int64) *LoginAppModel.UserToken {
	ret := &LoginAppModel.UserToken{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Take(db *gorm.DB, where ...interface{}) *LoginAppModel.UserToken {
	ret := &LoginAppModel.UserToken{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) Find(db *gorm.DB, cnd *simple.SqlCnd) (list []LoginAppModel.UserToken) {
	cnd.Find(db, &list)
	return
}

func (r *userTokenRepository) FindOne(db *gorm.DB, cnd *simple.SqlCnd) *LoginAppModel.UserToken {
	ret := &LoginAppModel.UserToken{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *userTokenRepository) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []LoginAppModel.UserToken, paging *simple.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *userTokenRepository) FindPageByCnd(db *gorm.DB, cnd *simple.SqlCnd) (list []LoginAppModel.UserToken, paging *simple.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &LoginAppModel.UserToken{})

	paging = &simple.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *userTokenRepository) Create(db *gorm.DB, t *LoginAppModel.UserToken) (err error) {
	err = db.Debug().Create(t).Error
	return
}

func (r *userTokenRepository) Update(db *gorm.DB, t *LoginAppModel.UserToken) (err error) {
	err = db.Save(t).Error
	return
}

func (r *userTokenRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&LoginAppModel.UserToken{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *userTokenRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&LoginAppModel.UserToken{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *userTokenRepository) Delete(db *gorm.DB, id int64) {
	db.Delete(&LoginAppModel.UserToken{}, "id = ?", id)
}
