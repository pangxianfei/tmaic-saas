package services

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"

	"tmaic/app/LoginApp/http/requests"

	"gitee.com/pangxianfei/framework/helpers/debug"
	"gitee.com/pangxianfei/library/tmaic"
	"gitee.com/pangxianfei/saas"
	sysmdel "gitee.com/pangxianfei/saas/sysmodel"
	"gitee.com/pangxianfei/simple"
	LoginAppModel "tmaic/app/LoginApp/model"
)

var TenantInstanceService = newTenantInstanceService()

func newTenantInstanceService() *tenantInstanceService {
	return &tenantInstanceService{}
}

type tenantInstanceService struct {
}

func (tenant *tenantInstanceService) CreateDb() {

}

func (tenant *tenantInstanceService) CreateDbUser() {

}

func (tenant *tenantInstanceService) CreateDbGrantUser() {

}

func (tenant *tenantInstanceService) InitializeTenantInstance(UserRegister requests.UserRegister, Admin *LoginAppModel.Admin) error {

	AppList := AppInfoService.GetByAppCreateList()
	if len(AppList) <= 0 {
		return errors.New("租户：" + Admin.UserName + " 数据库出始化失败")
	}

	tenantDbWhere := &sysmdel.RetrieveDB{
		Status: 1,
		AppId:  1000000,
	}
	var InstanceDB sysmdel.InstanceDb

	Result := simple.DB().Debug().Model(&sysmdel.InstanceDb{}).Where(tenantDbWhere).First(&InstanceDB)

	if Result.RowsAffected <= 0 {
		return errors.New("租户： 数据库出始化失败")
	}

	TenantsErr := simple.DB().Transaction(func(tx *gorm.DB) error {
		for _, appinfo := range AppList {
			//租户数据库
			TenantDbName := fmt.Sprintf("db_%s_%d", strings.ToLower(appinfo.Key), Admin.TenantId)
			dbPassword := tmaic.MD5(UserRegister.Password)
			saasDb := saas.DB.SetTenantsDb(InstanceDB.TenantsId, InstanceDB.AppId)
			CreateUser := fmt.Sprintf("CREATE USER '%s'@'%s' IDENTIFIED BY '%s'", TenantDbName, "%", dbPassword)
			GRANT := fmt.Sprintf("GRANT ALL PRIVILEGES ON `%s`.* TO '%s'@'%s' WITH GRANT OPTION", TenantDbName, TenantDbName, "%")
			if err := saasDb.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_general_ci", TenantDbName)).Error; err != nil {
				return err
			}

			if saasDb.Exec(CreateUser).Error != nil && saasDb.Exec(GRANT).Error != nil {
				debug.Dd(CreateUser)
				debug.Dd(GRANT)
			}

			//创建应用数据库连接信息
			createInstanceDb := &sysmdel.InstanceDb{
				TenantsId:          Admin.TenantId,
				Code:               appinfo.Key,
				AppName:            appinfo.Name,
				UseType:            "save",
				AppId:              appinfo.Id,
				Host:               InstanceDB.Host,
				DriverName:         InstanceDB.DriverName,
				Port:               InstanceDB.Port,
				Prefix:             InstanceDB.Prefix,
				DbName:             TenantDbName,
				Dbuser:             TenantDbName,
				Charset:            InstanceDB.Charset,
				Collation:          InstanceDB.Collation,
				SetmaxIdleconns:    3,
				Setmaxopenconns:    10,
				Setconnmaxlifetime: 60,
				Password:           dbPassword,
			}
			checkInstanceDB := &sysmdel.InstanceDb{
				TenantsId: Admin.TenantId,
				AppId:     appinfo.Id,
			}
			result := simple.DB().Where(checkInstanceDB).First(checkInstanceDB)
			if result.RowsAffected <= 0 {

				if createInstanceDbErr := simple.DB().Create(createInstanceDb).Error; createInstanceDbErr != nil {
					return createInstanceDbErr
				}
			}
		}
		return nil
	})

	return TenantsErr
}
