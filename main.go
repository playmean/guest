package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/playmean/guest/commands"
	"github.com/playmean/guest/storage"

	"github.com/urfave/cli/v2"
)

//go:embed assets/*
var embedded embed.FS

//go:generate go run tools/typesgen.go

func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	fmt.Println(scanner.Text())
	// }

	app := &cli.App{
		Name:  "guest",
		Usage: "be a nice guest to protocols",
		Commands: []*cli.Command{
			{
				Name:      "grab",
				Usage:     "download workspace from resource",
				UsageText: "guest grab [options] PATH",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "storage",
						Aliases: []string{"s"},
						Usage:   "storage for grabbing workspace",
					},
				},
				Action: commands.Grab,
			},
			{
				Name:      "init",
				Usage:     "init local workspace",
				UsageText: "guest init [options] [PATH]",
				Action:    commands.Init,
			},
			{
				Name:      "knock",
				Usage:     "knock something",
				UsageText: "guest knock [options] [PATH]",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "storage",
						Aliases: []string{"s"},
						Usage:   "storage for using workspace from",
					},
					&cli.StringSliceFlag{
						Name:  "var",
						Usage: "variable to pass (key=value)",
					},
				},
				Action: commands.Knock,
			},
			{
				Name:      "ui",
				Usage:     "start web ui",
				UsageText: "guest ui options [PORT]",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Usage:   "port for web server",
					},
				},
				Action: commands.StartUI,
			},
		},
	}

	storage.DefaultManager.RegisterStorage("local", storage.NewLocal())
	storage.DefaultManager.RegisterStorage("git", storage.NewGit())

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)

		os.Exit(1)
	}
}
