module tmaic

go 1.15

replace framework => ./vendors/framework

require (
	gitee.com/pangxianfei/frame v1.1.9
	gitee.com/pangxianfei/simple v1.0.6
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.7.0
	github.com/getsentry/raven-go v0.2.0
	github.com/gin-contrib/sentry v0.0.0-20191119142041-ff0e9556d1b7
	github.com/gin-gonic/gin v1.8.1
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/goburrow/cache v0.1.2
	github.com/golang/protobuf v1.5.2
	github.com/iancoleman/strcase v0.2.0
	github.com/iris-contrib/middleware/cors v0.0.0-20201115103636-07e8bced147f
	github.com/iris-contrib/middleware/jwt v0.0.0-20221109225525-f806663b83a0
	github.com/jinzhu/copier v0.3.5
	github.com/kataras/iris/v12 v12.2.0-beta6
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/nsqio/go-nsq v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.8.1
	github.com/urfave/cli v1.22.10
	github.com/ztrue/tracerr v0.3.0
	google.golang.org/protobuf v1.28.1
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.0.1
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.2
)
