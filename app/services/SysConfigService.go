package services

import (
	"strconv"
	"tmaic/app/buffer"
	"tmaic/app/common/constants"
	"tmaic/app/model"

	"github.com/spf13/cast"
	"tmaic/vendors/framework/simple/jsons"
	"tmaic/vendors/framework/simple/strs"

	"github.com/sirupsen/logrus"
)

var SysConfigService = newSysConfigService()

func newSysConfigService() *sysConfigService {
	return &sysConfigService{}
}

type sysConfigService struct {
}

func (s *sysConfigService) Get(id int64) *model.SysConfig {
	//return repositories.SysConfigRepository.Get(sqls.DB(), id)
	return nil
}

func (s *sysConfigService) GetTokenExpireDays() int {
	tokenExpireDaysStr := buffer.SysConfigCache.GetValue(constants.SysConfigTokenExpireDays)
	tokenExpireDays, err := strconv.Atoi(tokenExpireDaysStr)
	if err != nil {
		tokenExpireDays = constants.DefaultTokenExpireDays
	}
	if tokenExpireDays <= 0 {
		tokenExpireDays = constants.DefaultTokenExpireDays
	}
	return tokenExpireDays
}

func (s *sysConfigService) GetLoginMethod() model.LoginMethod {
	loginMethodStr := buffer.SysConfigCache.GetValue(constants.SysConfigLoginMethod)

	useDefault := true
	var loginMethod model.LoginMethod
	if strs.IsNotBlank(loginMethodStr) {
		if err := jsons.Parse(loginMethodStr, &loginMethod); err != nil {
			logrus.Warn("登录方式数据错误", err)
		} else {
			useDefault = false
		}
	}
	if useDefault {
		loginMethod = model.LoginMethod{
			Password: true,
			QQ:       true,
			Github:   true,
		}
	}
	return loginMethod
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
