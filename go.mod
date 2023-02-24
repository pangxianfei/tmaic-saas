module tmaic

go 1.20

replace framework => ./vendors/framework

replace simple => ./vendors/simple

require (
	github.com/PuerkitoBio/goquery v1.5.1
	github.com/davecgh/go-spew v1.1.1
	github.com/fatih/color v1.7.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/goburrow/cache v0.1.2
	github.com/golang/protobuf v1.5.2
	github.com/iancoleman/strcase v0.2.0
	github.com/iris-contrib/go.uuid v2.0.0+incompatible
	github.com/iris-contrib/middleware/cors v0.0.0-20201115103636-07e8bced147f
	github.com/iris-contrib/middleware/jwt v0.0.0-20221109225525-f806663b83a0
	github.com/iris-contrib/schema v0.0.6
	github.com/jinzhu/copier v0.3.5
	github.com/kataras/iris/v12 v12.2.0-beta6
	github.com/nsqio/go-nsq v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/robfig/cron v1.2.0
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.3.1
	github.com/spf13/viper v1.8.1
	github.com/urfave/cli v1.22.10
	github.com/ztrue/tracerr v0.3.0
	golang.org/x/crypto v0.0.0-20221005025214-4161e89ecf1b
	google.golang.org/protobuf v1.28.1
	gopkg.in/go-playground/validator.v9 v9.31.0
	gopkg.in/yaml.v2 v2.4.0
	gorm.io/driver/mysql v1.0.1
	gorm.io/driver/sqlserver v1.4.1
	gorm.io/gorm v1.24.2
)

require (
	github.com/BurntSushi/toml v1.2.0 // indirect
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/CloudyKit/jet/v6 v6.1.0 // indirect
	github.com/Joker/jade v1.1.3 // indirect
	github.com/Shopify/goreferrer v0.0.0-20220729165902-8cddb4f5de06 // indirect
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/aymerick/douceur v0.2.0 // indirect
	github.com/blang/semver/v4 v4.0.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.0-20190314233015-f79a8a8ca69d // indirect
	github.com/eknkc/amber v0.0.0-20171010120322-cdade1c07385 // indirect
	github.com/fatih/structs v1.1.0 // indirect
	github.com/flosch/pongo2/v4 v4.0.2 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/golang-jwt/jwt/v4 v4.4.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/css v1.0.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kataras/blocks v0.0.7 // indirect
	github.com/kataras/golog v0.1.7 // indirect
	github.com/kataras/neffos v0.0.20 // indirect
	github.com/kataras/pio v0.0.11 // indirect
	github.com/kataras/sitemap v0.0.6 // indirect
	github.com/kataras/tunnel v0.0.4 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/logrusorgru/aurora v0.0.0-20181002194514-a7b3b318ed4e // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailgun/raymond/v2 v2.0.47 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.0.9 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/mediocregopher/radix/v3 v3.8.0 // indirect
	github.com/microcosm-cc/bluemonday v1.0.21 // indirect
	github.com/microsoft/go-mssqldb v0.17.0 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/nats-io/nats.go v1.16.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nicksnyder/go-i18n/v2 v2.2.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/schollz/closestmatch v2.1.0+incompatible // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/tdewolff/minify/v2 v2.12.4 // indirect
	github.com/tdewolff/parse/v2 v2.6.4 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/vmihailenco/msgpack/v5 v5.3.5 // indirect
	github.com/vmihailenco/tagparser/v2 v2.0.0 // indirect
	github.com/yosssi/ace v0.0.5 // indirect
	golang.org/x/net v0.0.0-20221002022538-bcab6841153b // indirect
	golang.org/x/sys v0.0.0-20220919091848-fb04ddd9f9c8 // indirect
	golang.org/x/text v0.4.0 // indirect
	golang.org/x/time v0.0.0-20220922220347-f3bd1da661af // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
