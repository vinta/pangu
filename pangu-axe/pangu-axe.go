package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/vinta/pangu"
	"os"
	"path/filepath"
)

const (
	NAME    = "pangu-axe"
	USAGE   = "Paranoid text spacing"
	VERSION = "1.1.0"
	AUTHOR  = "Vinta Chen"
	EMAIL   = "vinta.chen@gmail.com"
)

const PREFIX = "readable."

func checkErrorExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func outputFilename(path, specified string) string {
	if len(specified) > 0 {
		return specified
	}

	filename := filepath.Base(path)
	output := PREFIX + filename

	return output
}

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
			Usage:   "Performs paranoid text spacing on text",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					cli.ShowSubcommandHelp(c)
					return
				}

				text := c.Args().First()
				fmt.Println(pangu.TextSpacing(text))
			},
		},
		{
			Name:    "file",
			Usage:   "Performs paranoid text spacing on files",
			Aliases: []string{"f"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: fmt.Sprintf(`Specifies the output file name. If not specified, the output file name will be "%sfilename.ext"`, PREFIX),
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					cli.ShowSubcommandHelp(c)
					return
				}

				o := c.String("output")

				if len(c.Args()) > 1 && len(o) > 0 {
					fmt.Println(`can't use the "-output" flag with multiple files`)
					os.Exit(1)
				}

				for _, filename := range c.Args() {
					if _, err := os.Stat(filename); os.IsNotExist(err) {
						fmt.Printf("no such file or directory: %s\n", filename)
						continue
					}

					var fw *os.File
					var err error

					switch o {
					case "stdout", "STDOUT":
						fw = os.Stdout
					case "stderr", "STDERR":
						fw = os.Stderr
					default:
						output := outputFilename(filename, o)
						fw, err = os.Create(output)
						checkErrorExit(err)
						defer fw.Close()
					}

					err = pangu.FileSpacing(filename, fw)
					checkErrorExit(err)
				}
			},
		},
	}

	app.Run(os.Args)
}
