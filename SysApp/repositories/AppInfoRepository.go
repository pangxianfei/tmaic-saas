package repositories

import (
	"gorm.io/gorm"
	"tmaic/SysApp/model"
)

var AppInfoRepository = new(appInfoRepository)

type appInfoRepository struct {
}

func (r *appInfoRepository) Take(db *gorm.DB, where ...interface{}) *SysAppModel.AppInfo {
	ret := &SysAppModel.AppInfo{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
func (r *appInfoRepository) GetByName(db *gorm.DB, name string) *SysAppModel.AppInfo {
	return r.Take(db, "name = ?", name)
}

func (r *appInfoRepository) Create(db *gorm.DB, AppInfo *SysAppModel.AppInfo) (appInfo *SysAppModel.AppInfo, err error) {
	var createStatus bool = false
	createStatus = db.Debug().Migrator().HasTable(&SysAppModel.AppInfo{})
	if createStatus == false {
		err := db.Migrator().CreateTable(&SysAppModel.AppInfo{})
		if err != nil {
			return nil, err
		}
	}
	err = db.Create(AppInfo).Error
	return
}

func (r *appInfoRepository) GetByList(db *gorm.DB) []SysAppModel.AppInfo {

	var AppList []SysAppModel.AppInfo
	db.Find(&AppList)
	return AppList
}

func (r *appInfoRepository) GetByAppCreateList(db *gorm.DB) []SysAppModel.AppInfo {

	var AppList []SysAppModel.AppInfo
	db.Where("is_created = ?", 1).Find(&AppList)
	return AppList
}
