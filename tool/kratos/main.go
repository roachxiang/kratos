package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "kratos"
	app.Usage = "kratos工具集"
	app.Version = Version
	app.Commands = []cli.Command{
		{
			Name:    "new",
			Aliases: []string{"n"},
			Usage:   "create new project",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:        "o",
					Value:       "",
					Usage:       "project owner for create project",
					Destination: &p.Owner,
				},
				cli.StringFlag{
					Name:        "d",
					Value:       "",
					Usage:       "project directory for create project",
					Destination: &p.Path,
				},
				cli.BoolFlag{
					Name:        "grpc",
					Usage:       "whether to use grpc for create project",
					Destination: &p.WithGRPC,
				},
			},
			Action: runNew,
		},
		{
			Name:    "build",
			Aliases: []string{"b"},
			Usage:   "kratos build",
			Action:  buildAction,
		},
		{
			Name:    "run",
			Aliases: []string{"r"},
			Usage:   "kratos run",
			Action:  runAction,
		},

		{
			Name:    "tool",
			Aliases: []string{"t"},
			Usage:   "kratos tool",
			Action:  toolAction,
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "kratos version",
			Action: func(c *cli.Context) error {
				fmt.Println(getVersion())
				return nil
			},
		},
		{
			Name:   "self-upgrade",
			Usage:  "kratos self-upgrade",
			Action: upgradeAction,
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
