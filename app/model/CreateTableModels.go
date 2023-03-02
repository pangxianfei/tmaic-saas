package model

import (
	model3 "tmaic/app/OrderApp/model"
	model2 "tmaic/app/SysApp/model"
	"tmaic/app/UserApp/model"
)

var CreateTableModels = []interface{}{
	&model.User{},
	&model.UserToken{},
	&model3.Article{},
	&model2.SysConfig{},
}
