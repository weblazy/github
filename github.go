package main

import (
	"github/action"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.EnableBashCompletion = true
	app.Name = "github"
	app.Usage = "github命令行管理工具"
	app.HideHelp = true
	app.HideHelpCommand = true
	token := &cli.StringFlag{
		Name:    "token",
		Usage:   "设置github开发者的accesstoken",
		Aliases: []string{"a"},
		EnvVars: []string{"GITHUB_TOKEN"},
	}
	// app.Authors = []*cli.Author{&cli.Author{Name: "lazy", Email: "2276282419@qq.com"}}
	// app.Version = "v1.0.0"
	app.Commands = []*cli.Command{
		{
			Name:    "release",
			Aliases: []string{"r"},
			Usage:   "管理github的release",
			Subcommands: []*cli.Command{
				{
					Name:    "list",
					Aliases: []string{"l"},
					Usage:   "获得release列表",
					Flags: []cli.Flag{
						token,
					},
					Action: action.List,
				},
				{
					Name:    "add",
					Aliases: []string{"a"},
					Usage:   "添加一个release",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:     "tag",
							Usage:    "设置对应版本tag",
							Aliases:  []string{"t"},
							Required: true,
						},
						&cli.StringFlag{
							Name:    "desc",
							Usage:   "设置描述信息",
							Aliases: []string{"d"},
						},
						&cli.StringFlag{
							Name:    "branch",
							Usage:   "设置默认分支",
							Aliases: []string{"b"},
							Value:   "master",
						},
						token,
					},
					Action: action.Add,
				},
				{
					Name:    "edit",
					Aliases: []string{"e"},
					Usage:   "编辑一个release",
					Action:  action.Edit,
				},
				{
					Name:    "delete",
					Aliases: []string{"d"},
					Usage:   "删除一个release",
					Action:  action.Delete,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
