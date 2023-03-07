package bootstrap

import (
	c "gitee.com/pangxianfei/library/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func ConfigInit() {
	// gorm配置
	gormConf := &gorm.Config{}
	// 初始化日志
	InitializationLog(gormConf)
	DbInit(gormConf)
}

func InitializationLog(gormConf *gorm.Config) bool {
	if file, err := os.OpenFile(c.GetString("app.Log_file"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		if c.GetBool("database.show-sql") {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		return false
	}
	return true
}
