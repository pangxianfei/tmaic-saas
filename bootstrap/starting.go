package bootstrap

import (
	"flag"
	"gitee.com/pangxianfei/simple"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"time"
	"tmaic/app/common"
	"tmaic/app/model"
	"tmaic/config"
	c "tmaic/vendors/framework/config"
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
	}
	logrus.Debug(c.GetInterface("database.connections.mysql"))
	// 连接数据库
	if err := simple.OpenDB(c.GetString("database.dns"), gormConf, 10, 20, model.Models...); err != nil {
		logrus.Error(err)
	}
}

// StartOn 开启定时任务
func StartOn() {
	if !common.IsProd() {
		return
	}
	startSchedule()
}
