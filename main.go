package main

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	_ "gitee.com/pangxianfei/library/config"
	"golang.org/x/sync/errgroup"
	"runtime"
	"time"
	"tmaic/commonApp/bootstrap"
)

var app errgroup.Group
var Saas bootstrap.Saas

func init() {}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Saas.Initialize()
	Saas.EnablingScheduledTask()
	app.Go(Saas.SysRun)
	time.Sleep(time.Millisecond * 200)
	app.Go(Saas.LoginApp)
	time.Sleep(time.Millisecond * 200)
	app.Go(Saas.UserApp)
	time.Sleep(time.Millisecond * 200)
	app.Go(Saas.OrderApp)
	time.Sleep(time.Millisecond * 200)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}
}
