package CommonModel

import SysAppModel "tmaic/app/SysApp/model"
import UserAppModel "tmaic/app/UserApp/model"
import LoginAppModel "tmaic/app/LoginApp/model"
import "gitee.com/pangxianfei/saas/sysmodel"

var CreateTableModels = []interface{}{

	//sysApp
	&SysAppModel.Authority{},
	&SysAppModel.SysConfig{},
	&SysAppModel.AppInfo{},
	&SysAppModel.AppInfo{},

	//UserAPP
	&UserAppModel.Roles{},

	//loginApp
	&LoginAppModel.TenantsInfo{},
	&LoginAppModel.UserToken{},

	//sysmodel
	&sysmdel.InstanceDb{},
	&sysmdel.InstanceDb{},
}
