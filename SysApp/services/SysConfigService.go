package services

import (
	"gitee.com/pangxianfei/simple/strs"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cast"
	"tmaic/SysApp/buffer"
	UserAppModel "tmaic/SysApp/model"
)

var SysConfigService = newSysConfigService()

func newSysConfigService() *sysConfigService {
	return &sysConfigService{}
}

type sysConfigService struct {
}

func (s *sysConfigService) Get(id int64) *UserAppModel.SysConfig {
	//return repositories.SysConfigRepository.Get(sqls.DB(), id)
	return nil
}

func (s *sysConfigService) GetTokenExpireDays() int {
	return 0
}

func (s *sysConfigService) GetStr(key, def string) (value string) {
	value = buffer.SysConfigCache.GetValue(key)
	if strs.IsBlank(value) {
		value = def
	}
	return
}

func (s *sysConfigService) GetInt(key string, def int) (value int) {
	str := buffer.SysConfigCache.GetValue(key)
	if strs.IsBlank(str) {
		value = def
		return
	}
	var err error
	if value, err = cast.ToIntE(value); err != nil {
		value = def
		logrus.Warn("Get int config error, use default value:", def, " key: ", key, " value: ", str)
		return
	}
	return
}
