package main

import (
	"runtime"
	"time"

	"gitee.com/pangxianfei/framework/kernel/debug"
	_ "gitee.com/pangxianfei/library/config"
	"golang.org/x/sync/errgroup"

	"tmaic/commonApp/bootstrap"
)

var app errgroup.Group
var Saas bootstrap.Saas
var startApplication map[string]func() error

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Saas.Initialize()
	Saas.EnablingScheduledTask()
	startApplication = make(map[string]func() error)
	startApplication["SysRun"] = Saas.SysRun
	startApplication["ProductApp"] = Saas.ProductApp
	startApplication["LoginApp"] = Saas.LoginApp
	startApplication["UserApp"] = Saas.UserApp
	startApplication["OrderApp"] = Saas.OrderApp
	for _, appName := range startApplication {
		app.Go(appName)
		delayTime()
	}
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}
}
func delayTime() {
	time.Sleep(time.Millisecond * 200)
}
