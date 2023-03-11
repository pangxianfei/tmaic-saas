package bootstrap

import (
	"fmt"
	"gitee.com/pangxianfei/library/config"
	CreateTable "tmaic/app/CommonModel"

	"gitee.com/pangxianfei/simple"
	"gorm.io/gorm"
)

func DbInit(gormConf *gorm.Config) {

	var UseDbType = config.Instance.DatabaseType.UseDbType
	switch UseDbType {

	case "mysql":
		_ = simple.OpenDB(UseDbType, config.Instance.DB.Dns, config.Instance.DB.Prefix, gormConf, config.Instance.DB.MaxIdleConns, config.Instance.DB.MaxOpenConns, CreateTable.CreateTableModels...)
	case "mssql":
		_ = simple.OpenDB(UseDbType, config.Instance.MSSQLDB.Sqlserver, config.Instance.MSSQLDB.Prefix, gormConf, config.Instance.MSSQLDB.SetMaxIdleConns, config.Instance.MSSQLDB.SetMaxOpenConns, CreateTable.CreateTableModels...)
	default:
		fmt.Printf("未知型")
	}
}
