package main

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
	_ "gitee.com/pangxianfei/library/config"
	"time"

	"golang.org/x/sync/errgroup"
	"tmaic/bootstrap"
)

var app errgroup.Group

func init() {}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//bootstrap.Run(ctx)
	//bootstrap.SysRun(ctx)
	//defer cancel()
	bootstrap.ConfigInit()
	bootstrap.Initialize()
	bootstrap.EnablingScheduledTask()
	app.Go(bootstrap.SysRun)
	time.Sleep(time.Second * 1)
	app.Go(bootstrap.LoginApp)
	time.Sleep(time.Second * 1)
	app.Go(bootstrap.UserApp)
	time.Sleep(time.Second * 1)
	app.Go(bootstrap.OrderApp)
	time.Sleep(time.Second * 1)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}
}
