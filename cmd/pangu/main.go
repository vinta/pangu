// Command pangu inserts whitespace between CJK and half-width characters in text, a file, or stdin.
package main

import (
	"os"

	"github.com/vinta/pangu/internal/cli"
)

func main() {
	os.Exit(cli.Main())
}
