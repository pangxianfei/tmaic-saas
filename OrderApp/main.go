package main

import (
	"runtime"

	"gitee.com/pangxianfei/framework/facades"
	"gitee.com/pangxianfei/framework/kernel/debug"
	"golang.org/x/sync/errgroup"

	"tmaic/commonApp/bootstrap"
)

var app errgroup.Group
var Saas bootstrap.Saas

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Saas.Initialize()
	Saas.EnablingScheduledTask()
	app.Go(Saas.OrderApp)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
		facades.Log.Info(err.Error())
	}
}
