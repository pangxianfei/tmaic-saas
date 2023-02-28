package simple

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
	"tmaic/vendors/framework/helpers/debug"
)

var (
	db    *gorm.DB
	sqlDB *sql.DB
)

func OpenDB(DbType string, dsn string, TablePrefix string, config *gorm.Config, maxIdleConns, maxOpenConns int, models ...interface{}) (err error) {
	if config == nil {
		config = &gorm.Config{}
	}

	if config.NamingStrategy == nil {
		config.NamingStrategy = schema.NamingStrategy{
			TablePrefix:   TablePrefix,
			SingularTable: true,
		}
	}
	var dberr error
	if DbType == "mssql" {
		db, dberr = gorm.Open(sqlserver.Open(dsn), config)
	} else if DbType == "mysql" {
		db, err = gorm.Open(mysql.Open(dsn), config)
	}

	if dberr != nil {
		//log.Fatal("Error creating connection pool: ", dberr.Error())
		debug.Dd("Error creating connection pool: ", dberr.Error())
		return dberr
	}

	if sqlDB, err = db.DB(); err == nil {
		sqlDB.SetMaxIdleConns(maxIdleConns)
		sqlDB.SetMaxOpenConns(maxOpenConns)
		sqlDB.SetConnMaxLifetime(time.Hour)
	} else {
		//log.Error(err)
		debug.Dd(err.Error())
		return err
	}

	if err = db.AutoMigrate(models...); nil != err {
		//log.Errorf("auto migrate tables failed: %s", err.Error())
		debug.Dd("数据表创建失败:", err.Error())
	}

	return
}

// DB 获取数据库链接
func DB() *gorm.DB {
	return db
}

// CloseDB 关闭连接
func CloseDB() {
	if sqlDB == nil {
		return
	}
	if err := sqlDB.Close(); nil != err {
		debug.Dd("Disconnect from database failed: %s", err.Error())
	}
}
