package main

import (
	"context"
	"sync"
	"tmaic/bootstrap"
	_ "tmaic/vendors/framework/config"
)

func init() {

}

func main() {
	ctx, _ := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	bootstrap.ConfigInit()
	bootstrap.Initialize()
	bootstrap.EnablingAscheduledTask()
	bootstrap.Run(ctx, wg)

	wg.Done()
}
