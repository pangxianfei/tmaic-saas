package bootstrap

import (
	"flag"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
	config "tmaic/config"
	c "tmaic/vendors/framework/config"
	"tmaic/vendors/framework/simple"
)

// ConfigInit 出始化配置
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

// DbInit 连接数据库
func DbInit(gormConf *gorm.Config) {
	if err := simple.OpenDB(config.Instance.DB.Url, c.GetString("database.db_prefix"), gormConf, config.Instance.DB.MaxIdleConns, config.Instance.DB.MaxOpenConns); err != nil {
		logrus.Error(err)
	}
}

// InitializationLog 初始化日志
func InitializationLog(gormConf *gorm.Config) bool {
	if file, err := os.OpenFile(c.GetString("app.LogFile"), os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666); err == nil {
		logrus.SetOutput(io.MultiWriter(os.Stdout, file))
		if c.GetBool("database.show-sql") {
			gormConf.Logger = logger.New(log.New(file, "\r\n", log.LstdFlags), logger.Config{
				SlowThreshold: time.Second,
				Colorful:      true,
				LogLevel:      logger.Info,
			})
		}
	} else {
		logrus.SetOutput(os.Stdout)
		logrus.Error(err)
		return false
	}
	return true
}
