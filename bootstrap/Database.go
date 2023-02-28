package bootstrap

import (
	"fmt"
	"gorm.io/gorm"
	"tmaic/app/model"
	"tmaic/config"
	"tmaic/vendors/framework/simple"
)

// DbInit 连接数据库
func DbInit(gormConf *gorm.Config) {

	var UseDbType = config.Instance.DatabaseType.UseDbType
	switch UseDbType {

	case "mysql":
		_ = simple.OpenDB("mysql", config.Instance.DB.Dns, config.Instance.DB.Prefix, gormConf, config.Instance.DB.MaxIdleConns, config.Instance.DB.MaxOpenConns, model.CreateTableModels...)
	case "mssql":
		_ = simple.OpenDB(config.Instance.DatabaseType.UseDbType, config.Instance.MSSQLDB.Sqlserver, config.Instance.MSSQLDB.Prefix, gormConf, config.Instance.MSSQLDB.SetMaxIdleConns, config.Instance.MSSQLDB.SetMaxOpenConns, model.CreateTableModels...)
	default:
		fmt.Printf("未知型")
	}
}
