package main

import (
	"context"
	_ "gitee.com/pangxianfei/library/config"
	"tmaic/bootstrap"
)

// gitee.com:pangxianfei/library v1.0.1
func init() {}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	bootstrap.ConfigInit()
	bootstrap.Initialize()
	bootstrap.EnablingScheduledTask()
	bootstrap.Run(ctx)
	defer cancel()

}
