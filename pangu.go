package pangu

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"regexp"
	"text/template"
)

const VERSION = "1.0.0"

// CJK is short for Chinese, Japanese and Korean.
//
// The constant cjk contains following Unicode blocks:
// 	\u2e80-\u2eff CJK Radicals Supplement
// 	\u2f00-\u2fdf Kangxi Radicals
// 	\u3040-\u309f Hiragana
// 	\u30a0-\u30ff Katakana
// 	\u3100-\u312f Bopomofo
// 	\u3200-\u32ff Enclosed CJK Letters and Months
// 	\u3400-\u4dbf CJK Unified Ideographs Extension A
// 	\u4e00-\u9fff CJK Unified Ideographs
// 	\uf900-\ufaff CJK Compatibility Ideographs
//
// For more information about Unicode blocks, see
// 	http://unicode-table.com/en/
const cjk = "" +
	"\u2e80-\u2eff" +
	"\u2f00-\u2fdf" +
	"\u3040-\u309f" +
	"\u30a0-\u30ff" +
	"\u3100-\u312f" +
	"\u3200-\u32ff" +
	"\u3400-\u4dbf" +
	"\u4e00-\u9fff" +
	"\uf900-\ufaff"

// ANS is short for Alphabets, Numbers
// and Symbols (`~!@#$%^&*()-_=+[]{}\|;:'",<.>/?).
//
// The constant ans doesn't contain all symbols above.
const ans = "A-Za-z0-9`~\\$%\\^&\\*\\-=\\+\\\\|<>/"

var cjk_quote = regexp.MustCompile(re("([{{ .CJK }}])" + "([\"'])"))
var quote_cjk = regexp.MustCompile(re("([\"'])" + "([{{ .CJK }}])"))
var fix_quote = regexp.MustCompile(re("([\"'])" + "(\\s*)" + "(.+?)" + "(\\s*)" + "([\"'])"))
var fix_single_quote = regexp.MustCompile(re("([{{ .CJK }}])" + "( )" + "(')" + "([A-Za-z0-9])"))

var cjk_bracket_cjk = regexp.MustCompile(re("([{{ .CJK }}])" + "([\\({\\[]+(.*?)[\\)}\\]]+)" + "([{{ .CJK }}])"))
var cjk_bracket = regexp.MustCompile(re("([{{ .CJK }}])" + "([\\(\\){}\\[\\]])"))
var bracket_cjk = regexp.MustCompile(re("([\\(\\){}\\[\\]])" + "([{{ .CJK }}])"))
var fix_bracket = regexp.MustCompile(re("([(\\({\\[)]+)" + "(\\s*)" + "(.+?)" + "(\\s*)" + "([\\)}\\]]+)"))

var cjk_hash = regexp.MustCompile(re("([{{ .CJK }}])" + "(#(\\S+))"))
var hash_cjk = regexp.MustCompile(re("((\\S+)#)" + "([{{ .CJK }}])"))

var fix_operator = regexp.MustCompile(re("([A-Za-z0-9{{ .CJK }}])" + "([\\+\\-\\*/=&\\|<>])" + "([A-Za-z0-9{{ .CJK }}])"))

var fix_symbol = regexp.MustCompile(re("([{{ .CJK }}])" + "([!;:,\\.\\?])" + "([A-Za-z0-9])"))

var cjk_ans = regexp.MustCompile(re("([{{ .CJK }}])([{{ .ANS }}@])"))
var ans_cjk = regexp.MustCompile(re("([{{ .ANS }}!;:,\\.\\?])([{{ .CJK }}])"))

type pattern struct {
	CJK string
	ANS string
}

func re(exp string) string {
	var buf bytes.Buffer

	var tmpl = template.New("pangu")
	tmpl, _ = tmpl.Parse(exp)
	pat := pattern{
		CJK: cjk,
		ANS: ans,
	}
	tmpl.Execute(&buf, pat)
	expr := buf.String()

	return expr
}

// TextSpacing performs a paranoid text spacing on input.
// It returns the processed text, with love.
func TextSpacing(input string) string {
	if len(input) < 2 {
		return input
	}

	text := input

	text = cjk_quote.ReplaceAllString(text, "$1 $2")
	text = quote_cjk.ReplaceAllString(text, "$1 $2")
	text = fix_quote.ReplaceAllString(text, "$1$3$5")
	text = fix_single_quote.ReplaceAllString(text, "$1$3$4")

	oldText := text
	newText := cjk_bracket_cjk.ReplaceAllString(oldText, "$1 $2 $4")
	text = newText
	if oldText == newText {
		text = cjk_bracket.ReplaceAllString(text, "$1 $2")
		text = bracket_cjk.ReplaceAllString(text, "$1 $2")
	}
	text = fix_bracket.ReplaceAllString(text, "$1$3$5")

	text = cjk_hash.ReplaceAllString(text, "$1 $2")
	text = hash_cjk.ReplaceAllString(text, "$1 $3")

	text = fix_operator.ReplaceAllString(text, "$1 $2 $3")

	text = fix_symbol.ReplaceAllString(text, "$1$2 $3")

	text = cjk_ans.ReplaceAllString(text, "$1 $2")
	text = ans_cjk.ReplaceAllString(text, "$1 $2")

	return text
}

// FileSpacing performs a paranoid text spacing on path, and
// generates a new file that contains processed content.
// It returns the path to the new generated file.
func FileSpacing(path string) (outPath string, err error) {
	fin, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fin.Close()

	outPath = "test_file.pangu.txt"
	fout, err := os.Create(outPath)
	if err != nil {
		return "", err
	}
	defer fout.Close()

	scanner := bufio.NewScanner(fin)
	writer := bufio.NewWriter(fout)
	for scanner.Scan() {
		line := TextSpacing(scanner.Text())
		fmt.Fprintln(writer, line)
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	defer writer.Flush()

	return outPath, nil
}
