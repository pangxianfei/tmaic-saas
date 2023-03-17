package main

import (
	"os"
	"tmaic/app/console/commands"
	"tmaic/database/migrations"

	"gitee.com/pangxianfei/framework/cmd"
	"gitee.com/pangxianfei/framework/cmd/commands/schedule"
	"gitee.com/pangxianfei/framework/console"
	"gitee.com/pangxianfei/framework/kernel/log"

	"tmaic/bootstrap"

	"github.com/urfave/cli"

	commandeer "gitee.com/pangxianfei/framework/cmd/commands/queue"

	appschedule "tmaic/app/console"
)

func init() {
	bootstrap.Initialize()
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
		console.Println(console.CODE_INFO, "COMMANDS:")

		for _, cate := range app.Categories() {
			categoryName := cate.Name
			if categoryName == "" {
				categoryName = "kernel"
			}
			console.Println(console.CODE_WARNING, "    "+categoryName+":")

			for _, cmds := range cate.Commands {
				console.Println(console.CODE_SUCCESS, "        "+cmds.Name+" "+console.Sprintf(console.CODE_INFO, "%s", cmds.ArgsUsage)+"    "+console.Sprintf(console.CODE_WARNING, "%s", cmds.Usage))
			}
		}
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}

}
