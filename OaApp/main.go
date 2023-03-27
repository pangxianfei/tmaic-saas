package main

import (
	"runtime"
	"time"

	"gitee.com/pangxianfei/framework/facades"
	"gitee.com/pangxianfei/framework/kernel/debug"
	_ "gitee.com/pangxianfei/library/config"
	"golang.org/x/sync/errgroup"

	"tmaic/commonApp/bootstrap"
)

var app errgroup.Group
var Saas bootstrap.Saas

func init() {}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Saas.Initialize()
	Saas.EnablingScheduledTask()
	app.Go(Saas.OaApp)
	time.Sleep(time.Millisecond * 200)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
		facades.Log.Info(err.Error())
	}
}
