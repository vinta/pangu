// Command pangu inserts whitespace between CJK and half-width characters in text, a file, or stdin.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime/debug"
	"strings"

	"github.com/vinta/pangu"
)

const usage = `usage: pangu [-h] [-v] [-t | -f | -c] [text_or_path]

pangu.go -- Paranoid text spacing for good readability, to automatically insert whitespace between CJK and half-width characters (alphabetical letters, numerical digits and symbols).

positional arguments:
  text_or_path   the text or file path to apply spacing; omit it to read stdin when input is piped

options:
  -h, --help     show this help message and exit
  -v, --version  show program's version number and exit
  -t, --text     treat the input as text (default)
  -f, --file     treat the input as a file path (pass - to read from stdin)
  -c, --check    check whether the input already has proper spacing (exit 0 if yes, 1 if no)

notes:
  - an explicit argument wins over piped stdin; stdin is read only when no argument is given
  - an explicit - argument always means stdin
  - put -- before text that starts with -
`

func main() {
	stat, err := os.Stdin.Stat()
	stdinIsTerminal := err == nil && stat.Mode()&os.ModeCharDevice != 0
	os.Exit(run(os.Args[1:], os.Stdin, stdinIsTerminal, os.Stdout, os.Stderr))
}

// run executes the CLI and returns its exit code: 0 on success, 1 when -c finds improper spacing or a file cannot be read, 2 on a usage error
func run(args []string, stdin io.Reader, stdinIsTerminal bool, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("pangu", flag.ContinueOnError)
	fs.SetOutput(stderr)
	fs.Usage = func() {}
	var isText, isFile, isCheck, showVersion bool
	for _, f := range []struct {
		value       *bool
		short, long string
	}{
		{&isText, "t", "text"},
		{&isFile, "f", "file"},
		{&isCheck, "c", "check"},
		{&showVersion, "v", "version"},
	} {
		fs.BoolVar(f.value, f.short, false, "")
		fs.BoolVar(f.value, f.long, false, "")
	}

	usageError := func(msg string) int {
		fmt.Fprintf(stderr, "%spangu: error: %s\n", usage, msg)
		return 2
	}

	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			fmt.Fprint(stdout, usage)
			return 0
		}
		// flag has already printed what went wrong
		fmt.Fprint(stderr, usage)
		return 2
	}

	if showVersion {
		fmt.Fprintln(stdout, "pangu.go", version())
		return 0
	}

	var modes []string
	for _, m := range []struct {
		set  bool
		name string
	}{{isText, "--text"}, {isFile, "--file"}, {isCheck, "--check"}} {
		if m.set {
			modes = append(modes, m.name)
		}
	}
	if len(modes) > 1 {
		return usageError(fmt.Sprintf("argument %s: not allowed with argument %s", modes[1], modes[0]))
	}
	if fs.NArg() > 1 {
		return usageError("unrecognized arguments: " + strings.Join(fs.Args()[1:], " "))
	}

	arg, hasArg := fs.Arg(0), fs.NArg() == 1
	// An empty string is what -f "$EMPTY_VAR" expands to, so it counts as a missing path rather than a file to open
	if isFile && arg == "" {
		return usageError("argument --file: expected a file path")
	}

	var source string
	switch {
	case arg == "-" || !hasArg && !stdinIsTerminal:
		// An explicit - always means stdin, under -f too (cf. tar -f -): the text itself arrives on stdin, so there is no file to open
		data, err := io.ReadAll(stdin)
		if err != nil {
			fmt.Fprintf(stderr, "pangu: error: %v\n", err)
			return 1
		}
		// Println puts the trailing newline back, so dropping one here passes piped input through unchanged
		source = strings.TrimSuffix(string(data), "\n")
		isFile = false
	case hasArg && isFile:
		data, err := os.ReadFile(arg)
		if err != nil {
			fmt.Fprintf(stderr, "pangu: error: %v\n", err)
			return 1
		}
		source = string(data)
	case hasArg:
		source = arg
	default:
		return usageError("the following arguments are required: text_or_path (or pipe text via stdin)")
	}

	spaced := pangu.SpaceText(source)

	if isCheck {
		if spaced == source {
			return 0
		}
		// stdout stays empty so -c composes in a pipeline; the correction goes to stderr so a failing check is debuggable
		fmt.Fprintf(stderr, "Corrected: %s\n", spaced)
		return 1
	}

	// File mode only does spacing: the file's own EOF newlines (or lack of one) pass through untouched
	if isFile {
		fmt.Fprint(stdout, spaced)
	} else {
		fmt.Fprintln(stdout, spaced)
	}
	return 0
}

// version is the module version go install recorded in the binary
func version() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		return info.Main.Version
	}
	return "(unknown)"
}
