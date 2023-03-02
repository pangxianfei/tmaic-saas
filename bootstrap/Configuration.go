package bootstrap

import (
	"flag"
	"log"
	"os"
	"time"
	"tmaic/config"

	//config "tmaic/config"

	c "gitee.com/pangxianfei/library/config"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConfigInit() {
	var configFile = flag.String("config", "./config/platform.yaml", "配置文件路径")
	flag.Parse()
	// 初始化配置
	_ = config.Init(*configFile)
	// gorm配置
	gormConf := &gorm.Config{}
	// 初始化日志
	InitializationLog(gormConf)
	DbInit(gormConf)
}

func InitializationLog(gormConf *gorm.Config) bool {
	//debug.Dd(c.GetString("app.Log_file"))
	if file, err := os.OpenFile(c.GetString("app.Log_file"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		if c.GetBool("database.show-sql") {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		//debug.Dd(err.Error())
		return false
	}
	return true
}
