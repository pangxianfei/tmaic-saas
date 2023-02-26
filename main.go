package main

import (
	"context"
	"tmaic/bootstrap"
	_ "tmaic/vendors/framework/config"
)

func init() {}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	bootstrap.ConfigInit()
	bootstrap.Initialize()
	bootstrap.EnablingScheduledTask()
	bootstrap.Run(ctx)
	defer cancel()

}
