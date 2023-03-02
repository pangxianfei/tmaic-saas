package main

import (
	"gitee.com/pangxianfei/framework/helpers/debug"
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

	bootstrap.ConfigInit()
	bootstrap.Initialize()
	bootstrap.EnablingScheduledTask()
	//bootstrap.SysRun(ctx)
	//defer cancel()
	app.Go(bootstrap.SysRun)
	time.Sleep(time.Second * 1)
	app.Go(bootstrap.UserApp)
	time.Sleep(time.Second * 1)
	app.Go(bootstrap.OrderApp)
	time.Sleep(time.Second * 1)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}

}
