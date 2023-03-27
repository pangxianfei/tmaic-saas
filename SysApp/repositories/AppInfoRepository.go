package repositories

import (
	"gorm.io/gorm"

	"tmaic/SysApp/models"
)

var AppInfoRepository = new(appInfoRepository)

type appInfoRepository struct {
}

func (r *appInfoRepository) Take(db *gorm.DB, where ...interface{}) *models.AppInfo {
	ret := &models.AppInfo{}
	if err := db.Debug().Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}
func (r *appInfoRepository) GetByName(db *gorm.DB, name string) *models.AppInfo {
	return r.Take(db, "name = ?", name)
}

func (r *appInfoRepository) Create(db *gorm.DB, AppInfo *models.AppInfo) (appInfo *models.AppInfo, err error) {
	var createStatus bool = false
	createStatus = db.Debug().Migrator().HasTable(&models.AppInfo{})
	if createStatus == false {
		err := db.Migrator().CreateTable(&models.AppInfo{})
		if err != nil {
			return nil, err
		}
	}
	err = db.Create(AppInfo).Error
	return
}

func (r *appInfoRepository) GetByList(db *gorm.DB) []models.AppInfo {

	var AppList []models.AppInfo
	db.Find(&AppList)
	return AppList
}

func (r *appInfoRepository) GetByAppCreateList(db *gorm.DB) []models.AppInfo {

	var AppList []models.AppInfo
	db.Where("is_created = ?", 1).Find(&AppList)
	return AppList
}
