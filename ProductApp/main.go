package main

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
	"tmaic/commonApp/bootstrap"
)

var app errgroup.Group
var Saas bootstrap.Saas

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Saas.Initialize()
	Saas.EnablingScheduledTask()
	app.Go(Saas.ProductApp)
	time.Sleep(time.Millisecond * 200)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}
}
