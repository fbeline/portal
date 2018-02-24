package main

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func main() {
	s := NewStorage()
	app := cli.NewApp()
	app.Name = "portal"
	app.Usage = "With it, you can create your own portals."
	app.Version = "0.0.1"

	app.Commands = []cli.Command{
		{
			Name:    "save",
			Aliases: []string{"s"},
			Usage:   "store directory",
			Action: func(c *cli.Context) error {
				s.Add(c.Args().First())
				return nil
			},
		},
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "list stored directories",
			Action: func(c *cli.Context) error {
				fmt.Println(s.List())
				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		path, err := Match(s.directories, c.Args().First())
		if err == nil {
			fmt.Print(path)
		} else {
			fmt.Println(err)
		}
		return err
	}

	app.Run(os.Args)
}
