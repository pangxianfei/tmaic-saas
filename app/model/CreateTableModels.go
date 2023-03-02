package model

import (
	OrderAppModel "tmaic/app/OrderApp/model"
	SysAppModel "tmaic/app/SysApp/model"
	UserAppModel "tmaic/app/UserApp/model"
)

var CreateTableModels = []interface{}{
	&UserAppModel.User{},
	&UserAppModel.UserToken{},
	&OrderAppModel.Article{},
	&SysAppModel.SysConfig{},
}
