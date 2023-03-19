package OrderApp

import (
	"gitee.com/pangxianfei/framework/kernel/debug"
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
	app.Go(Saas.OrderApp)
	if err := app.Wait(); err != nil {
		debug.Dd(err.Error())
	}
}
