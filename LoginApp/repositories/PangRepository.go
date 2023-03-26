package repositories

import (
	"gitee.com/pangxianfei/simple"
	"gitee.com/pangxianfei/simple/sqlcmd"
	"gorm.io/gorm"

	"tmaic/LoginApp/models/pang"
)

var PangRepository = new(PangDao)

type PangDao struct{}

func (r *PangDao) Get(db *gorm.DB, id int64) *pang.Pang {
	ret := &pang.Pang{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *PangDao) Take(db *gorm.DB, where ...interface{}) *pang.Pang {
	ret := &pang.Pang{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (r *PangDao) Find(db *gorm.DB, cnd *sqlcmd.Cnd) (list []pang.Pang) {
	cnd.Find(db, &list)
	return
}

func (r *PangDao) FindOne(db *gorm.DB, cnd *sqlcmd.Cnd) *pang.Pang {
	ret := &pang.Pang{}
	if err := cnd.FindOne(db, &ret); err != nil {
		return nil
	}
	return ret
}

func (r *PangDao) FindPageByParams(db *gorm.DB, params *simple.QueryParams) (list []pang.Pang, paging *sqlcmd.Paging) {
	return r.FindPageByCnd(db, &params.SqlCnd)
}

func (r *PangDao) FindPageByCnd(db *gorm.DB, cnd *sqlcmd.Cnd) (list []pang.Pang, paging *sqlcmd.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &pang.Pang{})

	paging = &sqlcmd.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *PangDao) Count(db *gorm.DB, cnd *sqlcmd.Cnd) int64 {
	return cnd.Count(db, &pang.Pang{})
}

func (r *PangDao) Create(db *gorm.DB, t *pang.Pang) (err error) {
	err = db.Create(t).Error
	return
}

func (r *PangDao) Update(db *gorm.DB, t *pang.Pang) (err error) {
	err = db.Save(t).Error
	return
}

func (r *PangDao) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&pang.Pang{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *PangDao) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&pang.Pang{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (r *PangDao) Delete(db *gorm.DB, id int64) error{
	return db.Delete(&pang.Pang{}, "id = ?", id).Error
}
