package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type result struct {
	code           int
	stdout, stderr string
}

// runCLI runs the CLI with stdin piped when piped is true, and with stdin on a terminal otherwise
func runCLI(t *testing.T, args []string, stdin string, piped bool) result {
	t.Helper()
	var stdout, stderr strings.Builder
	code := run(args, strings.NewReader(stdin), !piped, &stdout, &stderr)
	return result{code, stdout.String(), stderr.String()}
}

func writeTemp(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "temp_test.txt")
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func check(t *testing.T, got, want result) {
	t.Helper()
	if got != want {
		t.Errorf("got  %+v\nwant %+v", got, want)
	}
}

func TestHelp(t *testing.T) {
	got := runCLI(t, []string{"--help"}, "", false)
	if got.code != 0 || !strings.Contains(got.stdout, "usage: pangu") || !strings.Contains(got.stdout, "Paranoid text spacing") {
		t.Errorf("got %+v", got)
	}
}

func TestVersion(t *testing.T) {
	got := runCLI(t, []string{"-v"}, "", false)
	if got.code != 0 || !strings.HasPrefix(got.stdout, "pangu.go ") {
		t.Errorf("got %+v", got)
	}
}

func TestText(t *testing.T) {
	check(t, runCLI(t, []string{"-t", "你從什麼時候開始產生了我沒使用Monkey Patch的錯覺？"}, "", false), result{0, "你從什麼時候開始產生了我沒使用 Monkey Patch 的錯覺？\n", ""})
}

func TestTextByDefault(t *testing.T) {
	check(t, runCLI(t, []string{"與PM戰鬥的人"}, "", false), result{0, "與 PM 戰鬥的人\n", ""})
}

func TestFile(t *testing.T) {
	path := writeTemp(t, "老婆餅裡面沒有老婆，JavaScript裡面也沒有Java")
	check(t, runCLI(t, []string{"-f", path}, "", false), result{0, "老婆餅裡面沒有老婆，JavaScript 裡面也沒有 Java", ""})
}

func TestFileFixtures(t *testing.T) {
	for _, name := range []string{"text-file", "text-file-no-eof-newline"} {
		t.Run(name, func(t *testing.T) {
			want, err := os.ReadFile(filepath.Join("testdata", name+".expected.txt"))
			if err != nil {
				t.Fatal(err)
			}
			check(t, runCLI(t, []string{"-f", filepath.Join("testdata", name+".txt")}, "", false), result{0, string(want), ""})
		})
	}
}

// File mode only does spacing: no newline appended, the file's own EOF newlines and line endings pass through untouched
func TestFileKeepsNewlines(t *testing.T) {
	for _, tc := range []struct{ input, want string }{
		{"中文abc", "中文 abc"},
		{"中文abc\n", "中文 abc\n"},
		{"中文abc\n\n", "中文 abc\n\n"},
		{"中文abc\r\n第二行ABC\r\n", "中文 abc\r\n第二行 ABC\r\n"},
		{"中文abc\r尾行ABC\r", "中文 abc\r尾行 ABC\r"},
		{"中文abc\r\n尾行XYZ", "中文 abc\r\n尾行 XYZ"},
		{"字A\r\n空格 在行尾 \r\n", "字 A\r\n空格 在行尾 \r\n"},
	} {
		t.Run(tc.input, func(t *testing.T) {
			check(t, runCLI(t, []string{"-f", writeTemp(t, tc.input)}, "", false), result{0, tc.want, ""})
		})
	}
}

func TestFileMissing(t *testing.T) {
	got := runCLI(t, []string{"-f", filepath.Join(t.TempDir(), "missing.txt")}, "", false)
	if got.code != 1 || !strings.Contains(got.stderr, "pangu: error:") {
		t.Errorf("got %+v", got)
	}
}

func TestStdin(t *testing.T) {
	for _, tc := range []struct {
		name  string
		args  []string
		stdin string
		want  string
	}{
		{"no argument", nil, "當你凝視著bug，bug也凝視著你\n", "當你凝視著 bug，bug 也凝視著你\n"},
		{"-t", []string{"-t"}, "測試CLI參數\n", "測試 CLI 參數\n"},
		{"- argument", []string{"-"}, "老婆餅裡面沒有老婆\n", "老婆餅裡面沒有老婆\n"},
		{"-f -", []string{"-f", "-"}, "老婆餅裡面沒有老婆，JavaScript裡面也沒有Java\n", "老婆餅裡面沒有老婆，JavaScript 裡面也沒有 Java\n"},
		{"multi-line", nil, "第一行有bug\n第二行有Java\n", "第一行有 bug\n第二行有 Java\n"},
		{"argument wins over stdin", []string{"-t", "命令列的文字有PM"}, "標準輸入的文字有bug\n", "命令列的文字有 PM\n"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			check(t, runCLI(t, tc.args, tc.stdin, true), result{0, tc.want, ""})
		})
	}
}

func TestCheck(t *testing.T) {
	check(t, runCLI(t, []string{"-c"}, "當你凝視著 bug，bug 也凝視著你\n", true), result{0, "", ""})
	check(t, runCLI(t, []string{"-c"}, "當你凝視著bug\n", true), result{1, "", "Corrected: 當你凝視著 bug\n"})
	check(t, runCLI(t, []string{"-c", "中文abc"}, "", false), result{1, "", "Corrected: 中文 abc\n"})
}

func TestUsageErrors(t *testing.T) {
	for _, tc := range []struct {
		name  string
		args  []string
		piped bool
		want  string
	}{
		{"-f without a path even when piped", []string{"-f"}, true, "pangu: error: argument --file: expected a file path"},
		{"-f with an empty path", []string{"-f", ""}, true, "pangu: error: argument --file: expected a file path"},
		{"mutually exclusive modes", []string{"-t", "-f", "x.txt"}, false, "pangu: error: argument --file: not allowed with argument --text"},
		{"no input on a terminal", nil, false, "pangu: error: the following arguments are required"},
		{"unknown flag", []string{"-x"}, false, "flag provided but not defined: -x"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got := runCLI(t, tc.args, "當你凝視著bug\n", tc.piped)
			if got.code != 2 || got.stdout != "" || !strings.Contains(got.stderr, tc.want) {
				t.Errorf("got %+v\nwant code 2 and stderr containing %q", got, tc.want)
			}
		})
	}
}
