package bootstrap

import (
	"context"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"tmaic/app/events"
	"tmaic/app/http/middleware"
	"tmaic/app/jobs"
	"tmaic/app/listeners"
	"tmaic/config"
	"tmaic/resources/lang"
	"tmaic/routes"
	"tmaic/vendors/framework/cache"
	"tmaic/vendors/framework/helpers/log"
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

func Run(parentCtx context.Context, wg *sync.WaitGroup) {

	go func() {
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
	}()
	_, cancel := context.WithTimeout(parentCtx, 5*zone.Second)
	defer cancel()
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
	for _, value := range routeList {
		if value.Method == "POST" || value.Method == "GET" {
			log.Info(fmt.Sprintf(" %-6s %-35s --> %-50s", value.Method, value.Path, value.Name))
		}
	}
}
