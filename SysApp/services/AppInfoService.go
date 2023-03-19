package services

import (
	"errors"
	"gitee.com/pangxianfei/simple"
	"time"
	"tmaic/SysApp/http/requests"
	"tmaic/SysApp/model"
	"tmaic/SysApp/repositories"
)

var AppInfoService = newAppInfoService()

func newAppInfoService() *appInfoService {
	return &appInfoService{}
}

type appInfoService struct {
}

// GetByName 根据用户名查找
func (s *appInfoService) GetByName(mobile string) *SysAppModel.AppInfo {
	return repositories.AppInfoRepository.GetByName(simple.DB(), mobile)
}

func (s *appInfoService) GetByList() []SysAppModel.AppInfo {

	return repositories.AppInfoRepository.GetByList(simple.DB())
}

func (s *appInfoService) Create(appInfo requests.AppInfo) (AppInfo *SysAppModel.AppInfo, err error) {

	if oldApp := s.GetByName(appInfo.AppName); oldApp != nil {
		return nil, errors.New("SAAS应用:" + appInfo.AppName + " 已被占用")
	}
	//保存至DB
	newAppInfo := &SysAppModel.AppInfo{
		Name:        appInfo.AppName,
		Key:         appInfo.Key,
		Description: appInfo.Description,
		Status:      1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	newAppInfo, err = repositories.AppInfoRepository.Create(simple.DB(), newAppInfo)
	if err != nil {
		return nil, errors.New("SAAS应用创建失败")
	}
	return newAppInfo, nil
}

func (s *appInfoService) GetByAppCreateList() []SysAppModel.AppInfo {
	return repositories.AppInfoRepository.GetByAppCreateList(simple.DB())
}
