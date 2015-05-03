package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/vinta/pangu"
	"os"
)

const (
	NAME    = "pangu-axe"
	USAGE   = "Paranoid text spacing for good readability."
	VERSION = "1.0.0"
	AUTHOR  = "Vinta Chen"
	EMAIL   = "vinta.chen@gmail.com"
)

func main() {
	app := cli.NewApp()
	app.Name = NAME
	app.Usage = USAGE
	app.Version = VERSION
	app.Author = AUTHOR
	app.Email = EMAIL
	app.Commands = []cli.Command{
		{
			Name:    "text",
			Usage:   "Spacing text",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					fmt.Println("You need to provide a text.")
					os.Exit(1)
				}

				input := c.Args().First()
				fmt.Println(pangu.TextSpacing(input))
			},
		},
		{
			Name:    "file",
			Usage:   "Spacing file",
			Aliases: []string{"f"},
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					fmt.Println("You need to provide a file path.")
					os.Exit(1)
				}

				path := c.Args().First()
				outPath, err := pangu.FileSpacing(path)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(outPath)
			},
		},
	}

	app.Run(os.Args)
}
