package main

import (
	"os"

	"gitee.com/pangxianfei/framework/cmd"
	"gitee.com/pangxianfei/framework/cmd/commands/controllers"
	"gitee.com/pangxianfei/framework/cmd/commands/job"
	"gitee.com/pangxianfei/framework/cmd/commands/model"
	"gitee.com/pangxianfei/framework/cmd/commands/modelcache"
	"gitee.com/pangxianfei/framework/cmd/commands/mq"
	"gitee.com/pangxianfei/framework/cmd/commands/repository"
	"gitee.com/pangxianfei/framework/cmd/commands/request"
	"gitee.com/pangxianfei/framework/cmd/commands/schedule"
	"gitee.com/pangxianfei/framework/cmd/commands/services"
	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/framework/facades"
	"github.com/urfave/cli"

	appschedule "tmaic/app/console"
	"tmaic/app/console/commands"
	"tmaic/commonApp/bootstrap"
	"tmaic/commonApp/database/migrations"

	commandeer "gitee.com/pangxianfei/framework/cmd/commands/queue"
)

func init() {
	new(bootstrap.Saas).Initialize()
	controller.Initialize()
	model.Initialize()
	request.Initialize()
	services.Initialize()
	repository.Initialize()
	job.Initialize()
	mq.Initialize()
	modelcache.Initialize()
	migrations.Initialize()
	commandeer.Initialize()
	commands.Initialize()
	schedule.Initialize()
	appschedule.Schedule(cmd.NewSchedule())
}

func main() {
	cliServe()
}

func cliServe() {
	app := cli.NewApp()
	app.Name = "artisan"
	app.Usage = "让我们像工匠一样工作"
	app.Version = "1.1.5"
	app.Commands = cmd.List()
	app.Action = func(c *cli.Context) error {
		console.Println(console.CODE_SUCCESS, "commands:")
		for _, cate := range app.Categories() {
			categoryName := cate.Name
			if categoryName == "" {
				categoryName = "kernel"
			}
			console.Println(console.CODE_WARNING, ""+categoryName+":")
			for _, cmds := range cate.Commands {
				cmdsName := console.Sprintf(console.CODE_SUCCESS, "%-28s", cmds.Name)
				ArgsUsage := console.Sprintf(console.CODE_PARAM, "%-25s", cmds.ArgsUsage)
				cmdsUsage := console.Sprintf(console.CODE_WARNING, "%s", cmds.Usage)
				console.Println(console.CODE_SUCCESS, "  "+cmdsName+""+ArgsUsage+""+cmdsUsage)
			}
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		facades.Log.Fatal(err.Error())
		console.Println(console.CODE_WARNING, err.Error())
	}
}
