package main

import (
	"os"

	"gitee.com/pangxianfei/framework/cmd"
	"gitee.com/pangxianfei/framework/cmd/commands/controllers"
	"gitee.com/pangxianfei/framework/cmd/commands/model"
	"gitee.com/pangxianfei/framework/cmd/commands/request"
	"gitee.com/pangxianfei/framework/cmd/commands/schedule"
	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/framework/kernel/log"
	"github.com/urfave/cli"

	"tmaic/app/console/commands"
	"tmaic/commonApp/bootstrap"
	"tmaic/commonApp/database/migrations"

	commandeer "gitee.com/pangxianfei/framework/cmd/commands/queue"

	appschedule "tmaic/app/console"
)

func init() {
	new(bootstrap.Saas).Initialize()
	controller.Initialize()
	model.Initialize()
	request.Initialize()
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
		log.Fatal(err.Error())
	}
}
