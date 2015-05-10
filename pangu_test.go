package pangu_test

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/stretchr/testify/suite"
	"github.com/vinta/pangu"
	"io"
	"io/ioutil"
	"os"
	"testing"
)

type PanguTestSuite struct {
	suite.Suite
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func md5Of(filename string) string {
	var result []byte

	file, err := os.Open(filename)
	checkError(err)
	defer file.Close()

	hash := md5.New()
	_, err = io.Copy(hash, file)
	checkError(err)

	checksum := hex.EncodeToString(hash.Sum(result))

	return checksum
}

func (suite *PanguTestSuite) TestTextSpacing() {
	suite.Equal(`新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`, pangu.TextSpacing(`新八的構造成分有95%是眼鏡、3%是水、2%是垃圾`))
	suite.Equal(`新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`, pangu.TextSpacing(`新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`))

	suite.Equal(`所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`, pangu.TextSpacing(`所以,請問Jackey的鼻子有幾個?3.14個!`))
	suite.Equal(`所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`, pangu.TextSpacing(`所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`))
}

func (suite *PanguTestSuite) TestLatin1Supplement() {
	suite.Equal(`中文 Ø 漢字`, pangu.TextSpacing(`中文Ø漢字`))
	suite.Equal(`中文 Ø 漢字`, pangu.TextSpacing(`中文 Ø 漢字`))
}

func (suite *PanguTestSuite) TestGeneralPunctuation() {
	suite.Equal(`中文 • 漢字`, pangu.TextSpacing(`中文•漢字`))
	suite.Equal(`中文 • 漢字`, pangu.TextSpacing(`中文 • 漢字`))
}

func (suite *PanguTestSuite) TestNumberForms() {
	suite.Equal(`中文 Ⅶ 漢字`, pangu.TextSpacing(`中文Ⅶ漢字`))
	suite.Equal(`中文 Ⅶ 漢字`, pangu.TextSpacing(`中文 Ⅶ 漢字`))
}

func (suite *PanguTestSuite) TestCJKRadicalsSupplement() {
	suite.Equal(`abc ⻤ 123`, pangu.TextSpacing(`abc⻤123`))
	suite.Equal(`abc ⻤ 123`, pangu.TextSpacing(`abc ⻤ 123`))
}

func (suite *PanguTestSuite) TestKangxiRadicals() {
	suite.Equal(`abc ⾗ 123`, pangu.TextSpacing(`abc⾗123`))
	suite.Equal(`abc ⾗ 123`, pangu.TextSpacing(`abc ⾗ 123`))
}

func (suite *PanguTestSuite) TestHiragana() {
	suite.Equal(`abc あ 123`, pangu.TextSpacing(`abcあ123`))
	suite.Equal(`abc あ 123`, pangu.TextSpacing(`abc あ 123`))
}

func (suite *PanguTestSuite) TestKatakana() {
	suite.Equal(`abc ア 123`, pangu.TextSpacing(`abcア123`))
	suite.Equal(`abc ア 123`, pangu.TextSpacing(`abc ア 123`))
}

func (suite *PanguTestSuite) TestBopomofo() {
	suite.Equal(`abc ㄅ 123`, pangu.TextSpacing(`abcㄅ123`))
	suite.Equal(`abc ㄅ 123`, pangu.TextSpacing(`abc ㄅ 123`))
}

func (suite *PanguTestSuite) TestEnclosedCJKLettersAndMonths() {
	suite.Equal(`abc ㈱ 123`, pangu.TextSpacing(`abc㈱123`))
	suite.Equal(`abc ㈱ 123`, pangu.TextSpacing(`abc ㈱ 123`))
}

func (suite *PanguTestSuite) TestCJKUnifiedIdeographsExtensionA() {
	suite.Equal(`abc 㐂 123`, pangu.TextSpacing(`abc㐂123`))
	suite.Equal(`abc 㐂 123`, pangu.TextSpacing(`abc 㐂 123`))
}

func (suite *PanguTestSuite) TestCJKUnifiedIdeographs() {
	suite.Equal(`abc 丁 123`, pangu.TextSpacing(`abc丁123`))
	suite.Equal(`abc 丁 123`, pangu.TextSpacing(`abc 丁 123`))
}

func (suite *PanguTestSuite) TestCJKCompatibilityIdeographs() {
	suite.Equal(`abc 車 123`, pangu.TextSpacing(`abc車123`))
	suite.Equal(`abc 車 123`, pangu.TextSpacing(`abc 車 123`))
}

func (suite *PanguTestSuite) TestTilde() {
	suite.Equal(`前面~ 後面`, pangu.TextSpacing(`前面~後面`))
	suite.Equal(`前面 ~ 後面`, pangu.TextSpacing(`前面 ~ 後面`))
	suite.Equal(`前面~ 後面`, pangu.TextSpacing(`前面~ 後面`))
}

func (suite *PanguTestSuite) TestBackQuote() {
	suite.Equal("前面 ` 後面", pangu.TextSpacing("前面`後面"))
	suite.Equal("前面 ` 後面", pangu.TextSpacing("前面 ` 後面"))
	suite.Equal("前面 ` 後面", pangu.TextSpacing("前面` 後面"))
}

func (suite *PanguTestSuite) TestExclamationMark() {
	suite.Equal(`前面! 後面`, pangu.TextSpacing(`前面!後面`))
	suite.Equal(`前面 ! 後面`, pangu.TextSpacing(`前面 ! 後面`))
	suite.Equal(`前面! 後面`, pangu.TextSpacing(`前面! 後面`))
}

func (suite *PanguTestSuite) TestAt() {
	// https://twitter.com/vinta
	suite.Equal(`請 @vinta 吃大便`, pangu.TextSpacing(`請@vinta吃大便`))
	suite.Equal(`請 @vinta 吃大便`, pangu.TextSpacing(`請 @vinta 吃大便`))

	// http://weibo.com/vintalines
	suite.Equal(`請 @陳上進 吃大便`, pangu.TextSpacing(`請@陳上進 吃大便`))
	suite.Equal(`請 @陳上進 吃大便`, pangu.TextSpacing(`請 @陳上進 吃大便`))

	// TODO
	// suite.Equal(`陳上進@地球`, pangu.TextSpacing(`陳上進@地球`))
}

func (suite *PanguTestSuite) TestHash() {
	suite.Equal(`前面 #H2G2 後面`, pangu.TextSpacing(`前面#H2G2後面`))
	suite.Equal(`前面 #銀河便車指南 後面`, pangu.TextSpacing(`前面#銀河便車指南 後面`))
	suite.Equal(`前面 #銀河公車指南 #銀河拖吊車指南 後面`, pangu.TextSpacing(`前面#銀河公車指南 #銀河拖吊車指南 後面`))

	suite.Equal(`前面 #H2G2# 後面`, pangu.TextSpacing(`前面#H2G2#後面`))
	suite.Equal(`前面 #銀河閃電霹靂車指南# 後面`, pangu.TextSpacing(`前面#銀河閃電霹靂車指南#後面`))
}

func (suite *PanguTestSuite) TestDollar() {
	suite.Equal(`前面 $ 後面`, pangu.TextSpacing(`前面$後面`))
	suite.Equal(`前面 $ 後面`, pangu.TextSpacing(`前面 $ 後面`))

	suite.Equal(`前面 $100 後面`, pangu.TextSpacing(`前面$100後面`))

	// TODO
	// suite.Equal(`前面 $一百塊 後面`, pangu.TextSpacing(`前面$一百塊 後面`))
}

func (suite *PanguTestSuite) TestPercent() {
	suite.Equal(`前面 % 後面`, pangu.TextSpacing(`前面%後面`))
	suite.Equal(`前面 % 後面`, pangu.TextSpacing(`前面 % 後面`))

	suite.Equal(`前面 100% 後面`, pangu.TextSpacing(`前面100%後面`))
}

func (suite *PanguTestSuite) TestCarat() {
	suite.Equal(`前面 ^ 後面`, pangu.TextSpacing(`前面^後面`))
	suite.Equal(`前面 ^ 後面`, pangu.TextSpacing(`前面 ^ 後面`))
}

func (suite *PanguTestSuite) TestAmpersand() {
	suite.Equal(`前面 & 後面`, pangu.TextSpacing(`前面&後面`))
	suite.Equal(`前面 & 後面`, pangu.TextSpacing(`前面 & 後面`))

	suite.Equal(`Vinta&Mollie`, pangu.TextSpacing(`Vinta&Mollie`))
	suite.Equal(`Vinta & 陳上進`, pangu.TextSpacing(`Vinta&陳上進`))
	suite.Equal(`陳上進 & Vinta`, pangu.TextSpacing(`陳上進&Vinta`))

	suite.Equal(`得到一個 A&B 的結果`, pangu.TextSpacing(`得到一個A&B的結果`))
}

func (suite *PanguTestSuite) TestAsterisk() {
	suite.Equal(`前面 * 後面`, pangu.TextSpacing(`前面*後面`))
	suite.Equal(`前面 * 後面`, pangu.TextSpacing(`前面 * 後面`))

	suite.Equal(`Vinta*Mollie`, pangu.TextSpacing(`Vinta*Mollie`))
	suite.Equal(`Vinta * 陳上進`, pangu.TextSpacing(`Vinta*陳上進`))
	suite.Equal(`陳上進 * Vinta`, pangu.TextSpacing(`陳上進*Vinta`))

	suite.Equal(`得到一個 A*B 的結果`, pangu.TextSpacing(`得到一個A*B的結果`))
}

func (suite *PanguTestSuite) TestParenthesis() {
	suite.Equal(`前面 ( 後面`, pangu.TextSpacing(`前面(後面`))
	suite.Equal(`前面 ( 後面`, pangu.TextSpacing(`前面 ( 後面`))

	suite.Equal(`前面 ) 後面`, pangu.TextSpacing(`前面)後面`))
	suite.Equal(`前面 ) 後面`, pangu.TextSpacing(`前面 ) 後面`))

	suite.Equal(`前面 (中文 123 漢字) 後面`, pangu.TextSpacing(`前面(中文123漢字)後面`))
	suite.Equal(`前面 (中文 123) 後面`, pangu.TextSpacing(`前面(中文123)後面`))
	suite.Equal(`前面 (123 漢字) 後面`, pangu.TextSpacing(`前面(123漢字)後面`))
	suite.Equal(`前面 (中文 123 漢字) tail`, pangu.TextSpacing(`前面(中文123漢字) tail`))
	suite.Equal(`head (中文 123 漢字) 後面`, pangu.TextSpacing(`head (中文123漢字)後面`))
	suite.Equal(`head (中文 123 漢字) tail`, pangu.TextSpacing(`head (中文123漢字) tail`))
}

func (suite *PanguTestSuite) TestMinus() {
	suite.Equal(`前面 - 後面`, pangu.TextSpacing(`前面-後面`))
	suite.Equal(`前面 - 後面`, pangu.TextSpacing(`前面 - 後面`))

	suite.Equal(`Vinta-Mollie`, pangu.TextSpacing(`Vinta-Mollie`))
	suite.Equal(`Vinta - 陳上進`, pangu.TextSpacing(`Vinta-陳上進`))
	suite.Equal(`陳上進 - Vinta`, pangu.TextSpacing(`陳上進-Vinta`))

	suite.Equal(`得到一個 A-B 的結果`, pangu.TextSpacing(`得到一個A-B的結果`))
}

func (suite *PanguTestSuite) TestUnderscore() {
	suite.Equal(`前面_後面`, pangu.TextSpacing(`前面_後面`))
	suite.Equal(`前面 _ 後面`, pangu.TextSpacing(`前面 _ 後面`))
}

func (suite *PanguTestSuite) TestPlus() {
	suite.Equal(`前面 + 後面`, pangu.TextSpacing(`前面+後面`))
	suite.Equal(`前面 + 後面`, pangu.TextSpacing(`前面 + 後面`))

	suite.Equal(`Vinta+Mollie`, pangu.TextSpacing(`Vinta+Mollie`))
	suite.Equal(`Vinta + 陳上進`, pangu.TextSpacing(`Vinta+陳上進`))
	suite.Equal(`陳上進 + Vinta`, pangu.TextSpacing(`陳上進+Vinta`))

	suite.Equal(`得到一個 A+B 的結果`, pangu.TextSpacing(`得到一個A+B的結果`))

	suite.Equal(`得到一個 C++ 的結果`, pangu.TextSpacing(`得到一個C++的結果`))

	// TODO
	// suite.Equal(`得到一個 A+ 的結果`, pangu.TextSpacing(`得到一個A+的結果`))
}

func (suite *PanguTestSuite) TestEqual() {
	suite.Equal(`前面 = 後面`, pangu.TextSpacing(`前面=後面`))
	suite.Equal(`前面 = 後面`, pangu.TextSpacing(`前面 = 後面`))

	suite.Equal(`Vinta=Mollie`, pangu.TextSpacing(`Vinta=Mollie`))
	suite.Equal(`Vinta = 陳上進`, pangu.TextSpacing(`Vinta=陳上進`))
	suite.Equal(`陳上進 = Vinta`, pangu.TextSpacing(`陳上進=Vinta`))

	suite.Equal(`得到一個 A=B 的結果`, pangu.TextSpacing(`得到一個A=B的結果`))
}

func (suite *PanguTestSuite) TestBrace() {
	suite.Equal(`前面 { 後面`, pangu.TextSpacing(`前面{後面`))
	suite.Equal(`前面 { 後面`, pangu.TextSpacing(`前面 { 後面`))

	suite.Equal(`前面 } 後面`, pangu.TextSpacing(`前面}後面`))
	suite.Equal(`前面 } 後面`, pangu.TextSpacing(`前面 } 後面`))

	suite.Equal(`前面 {中文 123 漢字} 後面`, pangu.TextSpacing(`前面{中文123漢字}後面`))
	suite.Equal(`前面 {中文 123} 後面`, pangu.TextSpacing(`前面{中文123}後面`))
	suite.Equal(`前面 {123 漢字} 後面`, pangu.TextSpacing(`前面{123漢字}後面`))
	suite.Equal(`前面 {中文 123 漢字} tail`, pangu.TextSpacing(`前面{中文123漢字} tail`))
	suite.Equal(`head {中文 123 漢字} 後面`, pangu.TextSpacing(`head {中文123漢字}後面`))
	suite.Equal(`head {中文 123 漢字} tail`, pangu.TextSpacing(`head {中文123漢字} tail`))
}

func (suite *PanguTestSuite) TestBracket() {
	suite.Equal(`前面 [ 後面`, pangu.TextSpacing(`前面[後面`))
	suite.Equal(`前面 [ 後面`, pangu.TextSpacing(`前面 [ 後面`))

	suite.Equal(`前面 ] 後面`, pangu.TextSpacing(`前面]後面`))
	suite.Equal(`前面 ] 後面`, pangu.TextSpacing(`前面 ] 後面`))

	suite.Equal(`前面 [中文 123 漢字] 後面`, pangu.TextSpacing(`前面[中文123漢字]後面`))
	suite.Equal(`前面 [中文 123] 後面`, pangu.TextSpacing(`前面[中文123]後面`))
	suite.Equal(`前面 [123 漢字] 後面`, pangu.TextSpacing(`前面[123漢字]後面`))
	suite.Equal(`前面 [中文 123 漢字] tail`, pangu.TextSpacing(`前面[中文123漢字] tail`))
	suite.Equal(`head [中文 123 漢字] 後面`, pangu.TextSpacing(`head [中文123漢字]後面`))
	suite.Equal(`head [中文 123 漢字] tail`, pangu.TextSpacing(`head [中文123漢字] tail`))
}

func (suite *PanguTestSuite) TestPipe() {
	suite.Equal(`前面 | 後面`, pangu.TextSpacing(`前面|後面`))
	suite.Equal(`前面 | 後面`, pangu.TextSpacing(`前面 | 後面`))

	suite.Equal(`Vinta|Mollie`, pangu.TextSpacing(`Vinta|Mollie`))
	suite.Equal(`Vinta | 陳上進`, pangu.TextSpacing(`Vinta|陳上進`))
	suite.Equal(`陳上進 | Vinta`, pangu.TextSpacing(`陳上進|Vinta`))

	suite.Equal(`得到一個 A|B 的結果`, pangu.TextSpacing(`得到一個A|B的結果`))
}

func (suite *PanguTestSuite) TestBackslash() {
	suite.Equal(`前面 \ 後面`, pangu.TextSpacing(`前面\後面`))
}

func (suite *PanguTestSuite) TestColon() {
	suite.Equal(`前面: 後面`, pangu.TextSpacing(`前面:後面`))
	suite.Equal(`前面 : 後面`, pangu.TextSpacing(`前面 : 後面`))
	suite.Equal(`前面: 後面`, pangu.TextSpacing(`前面: 後面`))
}

func (suite *PanguTestSuite) TestSemicolon() {
	suite.Equal(`前面; 後面`, pangu.TextSpacing(`前面;後面`))
	suite.Equal(`前面 ; 後面`, pangu.TextSpacing(`前面 ; 後面`))
	suite.Equal(`前面; 後面`, pangu.TextSpacing(`前面; 後面`))
}

func (suite *PanguTestSuite) TestQuote() {
	// suite.Equal(`前面 " 後面`, pangu.TextSpacing(`前面"後面`))
	// suite.Equal(`前面 "" 後面`, pangu.TextSpacing(`前面""後面`))
	// suite.Equal(`前面 " " 後面`, pangu.TextSpacing(`前面" "後面`))

	suite.Equal(`前面 "中文 123 漢字" 後面`, pangu.TextSpacing(`前面"中文123漢字"後面`))
	suite.Equal(`前面 "中文 123" 後面`, pangu.TextSpacing(`前面"中文123"後面`))
	suite.Equal(`前面 "123 漢字" 後面`, pangu.TextSpacing(`前面"123漢字"後面`))
	suite.Equal(`前面 "中文 123 漢字" tail`, pangu.TextSpacing(`前面"中文123漢字" tail`))
	suite.Equal(`head "中文 123 漢字" 後面`, pangu.TextSpacing(`head "中文123漢字"後面`))
	suite.Equal(`head "中文 123 漢字" tail`, pangu.TextSpacing(`head "中文123漢字" tail`))

	// \u201c and \u201d
	suite.Equal(`前面 “中文 123 漢字” 後面`, pangu.TextSpacing(`前面“中文123漢字”後面`))
}

func (suite *PanguTestSuite) TestSingleQuote() {
	// suite.Equal(`前面 ' 後面`, pangu.TextSpacing(`前面'後面`))
	// suite.Equal(`前面 '' 後面`, pangu.TextSpacing(`前面''後面`))
	// suite.Equal(`前面 ' ' 後面`, pangu.TextSpacing(`前面' '後面`))

	suite.Equal(`前面 '中文 123 漢字' 後面`, pangu.TextSpacing(`前面'中文123漢字'後面`))
	suite.Equal(`前面 '中文 123' 後面`, pangu.TextSpacing(`前面'中文123'後面`))
	suite.Equal(`前面 '123 漢字' 後面`, pangu.TextSpacing(`前面'123漢字'後面`))
	suite.Equal(`前面 '中文 123 漢字' tail`, pangu.TextSpacing(`前面'中文123漢字' tail`))
	suite.Equal(`head '中文 123 漢字' 後面`, pangu.TextSpacing(`head '中文123漢字'後面`))
	suite.Equal(`head '中文 123 漢字' tail`, pangu.TextSpacing(`head '中文123漢字' tail`))

	suite.Equal(`陳上進 likes 林依諾's status.`, pangu.TextSpacing(`陳上進 likes 林依諾's status.`))
}

func (suite *PanguTestSuite) TestLessThan() {
	suite.Equal(`前面 < 後面`, pangu.TextSpacing(`前面<後面`))
	suite.Equal(`前面 < 後面`, pangu.TextSpacing(`前面 < 後面`))

	suite.Equal(`Vinta<Mollie`, pangu.TextSpacing(`Vinta<Mollie`))
	suite.Equal(`Vinta < 陳上進`, pangu.TextSpacing(`Vinta<陳上進`))
	suite.Equal(`陳上進 < Vinta`, pangu.TextSpacing(`陳上進<Vinta`))

	suite.Equal(`得到一個 A<B 的結果`, pangu.TextSpacing(`得到一個A<B的結果`))

	suite.Equal(`前面 <中文 123 漢字> 後面`, pangu.TextSpacing(`前面<中文123漢字>後面`))
	suite.Equal(`前面 <中文 123> 後面`, pangu.TextSpacing(`前面<中文123>後面`))
	suite.Equal(`前面 <123 漢字> 後面`, pangu.TextSpacing(`前面<123漢字>後面`))
	suite.Equal(`前面 <中文 123 漢字> tail`, pangu.TextSpacing(`前面<中文123漢字> tail`))
	suite.Equal(`head <中文 123 漢字> 後面`, pangu.TextSpacing(`head <中文123漢字>後面`))
	suite.Equal(`head <中文 123 漢字> tail`, pangu.TextSpacing(`head <中文123漢字> tail`))
}

func (suite *PanguTestSuite) TestComma() {
	suite.Equal(`前面, 後面`, pangu.TextSpacing(`前面,後面`))
	suite.Equal(`前面 , 後面`, pangu.TextSpacing(`前面 , 後面`))
	suite.Equal(`前面, 後面`, pangu.TextSpacing(`前面, 後面`))
}

func (suite *PanguTestSuite) TestGreaterThan() {
	suite.Equal(`前面 > 後面`, pangu.TextSpacing(`前面>後面`))
	suite.Equal(`前面 > 後面`, pangu.TextSpacing(`前面 > 後面`))

	suite.Equal(`Vinta>Mollie`, pangu.TextSpacing(`Vinta>Mollie`))
	suite.Equal(`Vinta > 陳上進`, pangu.TextSpacing(`Vinta>陳上進`))
	suite.Equal(`陳上進 > Vinta`, pangu.TextSpacing(`陳上進>Vinta`))

	suite.Equal(`得到一個 A>B 的結果`, pangu.TextSpacing(`得到一個A>B的結果`))
}

func (suite *PanguTestSuite) TestPeriod() {
	suite.Equal(`前面. 後面`, pangu.TextSpacing(`前面.後面`))
	suite.Equal(`前面 . 後面`, pangu.TextSpacing(`前面 . 後面`))
	suite.Equal(`前面. 後面`, pangu.TextSpacing(`前面. 後面`))
}

func (suite *PanguTestSuite) TestQuestionMark() {
	suite.Equal(`前面? 後面`, pangu.TextSpacing(`前面?後面`))
	suite.Equal(`前面 ? 後面`, pangu.TextSpacing(`前面 ? 後面`))
	suite.Equal(`前面? 後面`, pangu.TextSpacing(`前面? 後面`))
}

func (suite *PanguTestSuite) TestSlash() {
	suite.Equal(`前面 / 後面`, pangu.TextSpacing(`前面/後面`))
	suite.Equal(`前面 / 後面`, pangu.TextSpacing(`前面 / 後面`))

	suite.Equal(`Vinta/Mollie`, pangu.TextSpacing(`Vinta/Mollie`))
	suite.Equal(`Vinta / 陳上進`, pangu.TextSpacing(`Vinta/陳上進`))
	suite.Equal(`陳上進 / Vinta`, pangu.TextSpacing(`陳上進/Vinta`))

	suite.Equal(`得到一個 A/B 的結果`, pangu.TextSpacing(`得到一個A/B的結果`))

	// TODO
	// suite.Equal(`陳上進 / Vinta / Mollie`, pangu.TextSpacing(`陳上進/Vinta/Mollie`))
}

func (suite *PanguTestSuite) TestFileSpacing() {
	input := "_fixtures/test_file.txt"
	output := "_fixtures/test_file.pangu.txt"

	fw, err := os.Create(output)
	checkError(err)
	defer fw.Close()

	err = pangu.FileSpacing(input, fw)
	suite.Nil(err)
	suite.Equal(md5Of(output), md5Of("_fixtures/test_file.expected.txt"))
}

func (suite *PanguTestSuite) TestFileSpacingNoNewlineAtEOF() {
	input := "_fixtures/test_file_no_eof_newline.txt"
	output := "_fixtures/test_file_no_eof_newline.pangu.txt"

	fw, err := os.Create(output)
	checkError(err)
	defer fw.Close()

	err = pangu.FileSpacing(input, fw)
	suite.Nil(err)
	suite.Equal(md5Of(output), md5Of("_fixtures/test_file_no_eof_newline.expected.txt"))
}

func (suite *PanguTestSuite) TestFileSpacingNoSuchFile() {
	input := "_fixtures/none.exist"

	err := pangu.FileSpacing(input, ioutil.Discard)
	suite.EqualError(err, "open _fixtures/none.exist: no such file or directory")
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPanguTestSuite(t *testing.T) {
	suite.Run(t, new(PanguTestSuite))
}
