// Command pangu-go is pangu under an unambiguous name, for when pangu.js or pangu.py shadows pangu on PATH.
package main

import (
	"os"

	"github.com/vinta/pangu/internal/cli"
)

func main() {
	os.Exit(cli.Main())
}
