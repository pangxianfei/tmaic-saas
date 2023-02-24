package bootstrap

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"tmaic/app/events"
	"tmaic/app/http/middleware"
	"tmaic/app/jobs"
	"tmaic/app/listeners"
	"tmaic/config"
	"tmaic/resources/lang"
	"tmaic/routes"
	"tmaic/vendors/framework/cache"
	"tmaic/vendors/framework/console"
	"tmaic/vendors/framework/helpers/zone"
	"tmaic/vendors/framework/queue"
	"tmaic/vendors/framework/simple"
)

// Initialize 出始化框架
func Initialize() {
	config.Initialize()
	//database.Initialize()
	//variable.InitBasePath()
	//sentry.Initialize()
	//logs.Initialize()
	zone.Initialize()
	lang.Initialize()
	cache.Initialize()
	//m.Initialize()
	queue.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()

}

func Run(parentCtx context.Context) {

	app := iris.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.Use(logger.New())
	app.AllowMethods(iris.MethodOptions)
	//跨域
	app.Use(middleware.CORS)
	//注册路由
	route.Route(app)
	RouteNameList(app)
	httpServer := &http.Server{Addr: ":" + config.Instance.Port}
	handleSignal(httpServer, parentCtx)
	err := app.Run(iris.Server(httpServer), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}), iris.WithoutInterruptHandler)

	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
	<-parentCtx.Done()
}

func handleSignal(server *http.Server, parentCtx context.Context) {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now", s)
		if err := server.Shutdown(parentCtx); nil != err {
			logrus.Errorf("server close failed: " + err.Error())
		}
		simple.CloseDB()
		logrus.Infof("Exited")
		os.Exit(0)
	}()
}

// RouteNameList 打印路由列表
func RouteNameList(app *iris.Application) {
	routeList := app.GetRoutes()
	var index int = 1
	for _, value := range routeList {
		if strings.Contains(value.MainHandlerName, "tmaic") || strings.Contains(value.MainHandlerName, "iris") {
			continue
		}

		if value.Method == "POST" || value.Method == "GET" {
			console.Println(console.CODE_SUCCESS, " "+console.Sprintf(console.CODE_SUCCESS, "%-6d", index)+console.Sprintf(console.CODE_SUCCESS, "%-50s", value.MainHandlerName)+console.Sprintf(console.CODE_SUCCESS, "%-6s", value.Method)+" "+console.Sprintf(console.CODE_SUCCESS, "%-35s", value.Path))
			index++
		}

	}
}
