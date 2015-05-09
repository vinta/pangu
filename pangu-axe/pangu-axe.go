package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/vinta/pangu"
	"os"
	"path/filepath"
	"strings"
)

const (
	NAME    = "pangu-axe"
	USAGE   = "Paranoid text spacing"
	VERSION = "1.0.0"
	AUTHOR  = "Vinta Chen"
	EMAIL   = "vinta.chen@gmail.com"
)

func checkErrorExit(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func outFilename(path, specified string) string {
	if len(specified) > 0 {
		return specified
	}

	absPath, err := filepath.Abs(path)
	checkErrorExit(err)

	dir := filepath.Dir(absPath)
	err = os.MkdirAll(dir, 0755)
	checkErrorExit(err)

	ext := filepath.Ext(absPath)
	suffix := ".pangu"
	output := strings.Replace(absPath, ext, suffix+ext, 1)

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
			Usage:   "Spacing a text",
			Aliases: []string{"t"},
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					fmt.Println("USAGE:")
					fmt.Println(`   pangu-axe text "your text"`)

					return
				}

				text := c.Args().First()
				fmt.Println(pangu.TextSpacing(text))
			},
		},
		{
			Name:    "file",
			Usage:   "Spacing a file or files",
			Aliases: []string{"f"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "output, o",
					Value: "",
					Usage: "output filename",
				},
			},
			Action: func(c *cli.Context) {
				if len(c.Args()) == 0 {
					fmt.Println("USAGE:")
					fmt.Println("   pangu-axe file your_file.txt")
					fmt.Println("   pangu-axe file your_file.txt -o your_custom_output_filename")
					fmt.Println("   pangu-axe file your_file.txt -o stdout")

					return
				}

				filename := c.Args().First()

				o := c.String("output")
				output := outFilename(filename, o)

				fw, err := os.Create(output)
				checkErrorExit(err)
				defer fw.Close()

				err = pangu.FileSpacing(filename, fw)
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(output)
			},
		},
	}

	app.Run(os.Args)
}
