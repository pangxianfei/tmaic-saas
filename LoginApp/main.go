package main

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	_ "gitee.com/pangxianfei/library/config"
	"golang.org/x/sync/errgroup"
	"runtime"
	"tmaic/bootstrap"
)

var app errgroup.Group
var Saas bootstrap.Saas

func init() {}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Saas.Initialize()
	Saas.EnablingScheduledTask()
	app.Go(Saas.LoginApp)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}
}
