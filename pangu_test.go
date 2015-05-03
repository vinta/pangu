package pangu

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type PanguTestSuite struct {
	suite.Suite
}

func (suite *PanguTestSuite) assertEqualTextSpacing(expected, input string) {
	actual := TextSpacing(input)
	suite.Equal(expected, actual)
}

func (suite *PanguTestSuite) TestTextSpacing() {
	suite.Equal(`新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`, TextSpacing(`新八的構造成分有95%是眼鏡、3%是水、2%是垃圾`))
	suite.Equal(`新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`, TextSpacing(`新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾`))

	suite.Equal(`所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`, TextSpacing(`所以,請問Jackey的鼻子有幾個?3.14個!`))
	suite.Equal(`所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`, TextSpacing(`所以, 請問 Jackey 的鼻子有幾個? 3.14 個!`))
}

func (suite *PanguTestSuite) TestFileSpacing() {
	inPath := "_fixtures/test_file.txt"
	outPath, err := FileSpacing(inPath)
	suite.NotNil(outPath)
	suite.Nil(err)
}

func (suite *PanguTestSuite) TestCJKRadicalsSupplement() {
	suite.Equal(`abc ⻤ 123`, TextSpacing(`abc⻤123`))
	suite.Equal(`abc ⻤ 123`, TextSpacing(`abc ⻤ 123`))
}

func (suite *PanguTestSuite) TestKangxiRadicals() {
	suite.Equal(`abc ⾗ 123`, TextSpacing(`abc⾗123`))
	suite.Equal(`abc ⾗ 123`, TextSpacing(`abc ⾗ 123`))
}

func (suite *PanguTestSuite) TestHiragana() {
	suite.Equal(`abc あ 123`, TextSpacing(`abcあ123`))
	suite.Equal(`abc あ 123`, TextSpacing(`abc あ 123`))
}

func (suite *PanguTestSuite) TestKatakana() {
	suite.Equal(`abc ア 123`, TextSpacing(`abcア123`))
	suite.Equal(`abc ア 123`, TextSpacing(`abc ア 123`))
}

func (suite *PanguTestSuite) TestBopomofo() {
	suite.Equal(`abc ㄅ 123`, TextSpacing(`abcㄅ123`))
	suite.Equal(`abc ㄅ 123`, TextSpacing(`abc ㄅ 123`))
}

func (suite *PanguTestSuite) TestEnclosedCJKLettersAndMonths() {
	suite.Equal(`abc ㈱ 123`, TextSpacing(`abc㈱123`))
	suite.Equal(`abc ㈱ 123`, TextSpacing(`abc ㈱ 123`))
}

func (suite *PanguTestSuite) TestCJKUnifiedIdeographsExtensionA() {
	suite.Equal(`abc 㐂 123`, TextSpacing(`abc㐂123`))
	suite.Equal(`abc 㐂 123`, TextSpacing(`abc 㐂 123`))
}

func (suite *PanguTestSuite) TestCJKUnifiedIdeographs() {
	suite.Equal(`abc 丁 123`, TextSpacing(`abc丁123`))
	suite.Equal(`abc 丁 123`, TextSpacing(`abc 丁 123`))
}

func (suite *PanguTestSuite) TestCJKCompatibilityIdeographs() {
	suite.Equal(`abc 車 123`, TextSpacing(`abc車123`))
	suite.Equal(`abc 車 123`, TextSpacing(`abc 車 123`))
}

func (suite *PanguTestSuite) TestTilde() {
	suite.Equal(`前面 ~ 後面`, TextSpacing(`前面~後面`))
	suite.Equal(`前面 ~ 後面`, TextSpacing(`前面 ~ 後面`))
	suite.Equal(`前面 ~ 後面`, TextSpacing(`前面~ 後面`))
}

func (suite *PanguTestSuite) TestBackQuote() {
	suite.Equal("前面 ` 後面", TextSpacing("前面`後面"))
	suite.Equal("前面 ` 後面", TextSpacing("前面 ` 後面"))
	suite.Equal("前面 ` 後面", TextSpacing("前面` 後面"))
}

func (suite *PanguTestSuite) TestExclamationMark() {
	suite.Equal(`前面! 後面`, TextSpacing(`前面!後面`))
	suite.Equal(`前面 ! 後面`, TextSpacing(`前面 ! 後面`))
	suite.Equal(`前面! 後面`, TextSpacing(`前面! 後面`))
}

func (suite *PanguTestSuite) TestAt() {
	suite.Equal(`請 @vinta 吃大便`, TextSpacing(`請@vinta吃大便`))
	suite.Equal(`請 @vinta 吃大便`, TextSpacing(`請 @vinta 吃大便`))

	suite.Equal(`請 @陳上進 吃大便`, TextSpacing(`請@陳上進 吃大便`))
	suite.Equal(`請 @陳上進 吃大便`, TextSpacing(`請 @陳上進 吃大便`))

	// TODO
	// suite.Equal(`陳上進@地球`, TextSpacing(`陳上進@地球`))
}

func (suite *PanguTestSuite) TestHash() {
	suite.Equal(`前面 #H2G2 後面`, TextSpacing(`前面#H2G2後面`))

	suite.Equal(`前面 #銀河便車指南 後面`, TextSpacing(`前面#銀河便車指南 後面`))

	suite.Equal(`前面 #銀河公車指南 #銀河大客車指南 後面`, TextSpacing(`前面#銀河公車指南 #銀河大客車指南 後面`))

	suite.Equal(`前面 #銀河閃電霹靂車指南# 後面`, TextSpacing(`前面#銀河閃電霹靂車指南#後面`))
}

func (suite *PanguTestSuite) TestDollar() {
	suite.Equal(`前面 $ 後面`, TextSpacing(`前面$後面`))
	suite.Equal(`前面 $ 後面`, TextSpacing(`前面 $ 後面`))

	suite.Equal(`前面 $100 後面`, TextSpacing(`前面$100後面`))

	// TODO
	// suite.Equal(`前面 $一百塊 後面`, TextSpacing(`前面$一百塊 後面`))
}

func (suite *PanguTestSuite) TestPercent() {
	suite.Equal(`前面 % 後面`, TextSpacing(`前面%後面`))
	suite.Equal(`前面 % 後面`, TextSpacing(`前面 % 後面`))

	suite.Equal(`前面 100% 後面`, TextSpacing(`前面100%後面`))
}

func (suite *PanguTestSuite) TestCarat() {
	suite.Equal(`前面 ^ 後面`, TextSpacing(`前面^後面`))
	suite.Equal(`前面 ^ 後面`, TextSpacing(`前面 ^ 後面`))
}

func (suite *PanguTestSuite) TestAmpersand() {
	suite.Equal(`前面 & 後面`, TextSpacing(`前面&後面`))
	suite.Equal(`前面 & 後面`, TextSpacing(`前面 & 後面`))

	suite.Equal(`Vinta & Mollie`, TextSpacing(`Vinta&Mollie`))
	suite.Equal(`Vinta & 陳上進`, TextSpacing(`Vinta&陳上進`))
	suite.Equal(`陳上進 & Vinta`, TextSpacing(`陳上進&Vinta`))

	suite.Equal(`得到一個 A & B 的結果`, TextSpacing(`得到一個A&B的結果`))
}

func (suite *PanguTestSuite) TestAsterisk() {
	suite.Equal(`前面 * 後面`, TextSpacing(`前面*後面`))
	suite.Equal(`前面 * 後面`, TextSpacing(`前面 * 後面`))

	suite.Equal(`Vinta * Mollie`, TextSpacing(`Vinta*Mollie`))
	suite.Equal(`Vinta * 陳上進`, TextSpacing(`Vinta*陳上進`))
	suite.Equal(`陳上進 * Vinta`, TextSpacing(`陳上進*Vinta`))

	suite.Equal(`得到一個 A * B 的結果`, TextSpacing(`得到一個A*B的結果`))
}

func (suite *PanguTestSuite) TestParenthesis() {
	suite.Equal(`前面 ( 後面`, TextSpacing(`前面(後面`))
	suite.Equal(`前面 ( 後面`, TextSpacing(`前面 ( 後面`))

	suite.Equal(`前面 ) 後面`, TextSpacing(`前面)後面`))
	suite.Equal(`前面 ) 後面`, TextSpacing(`前面 ) 後面`))

	suite.Equal(`前面 (中文 123 漢字) 後面`, TextSpacing(`前面(中文123漢字)後面`))
	suite.Equal(`前面 (中文 123) 後面`, TextSpacing(`前面(中文123)後面`))
	suite.Equal(`前面 (123 漢字) 後面`, TextSpacing(`前面(123漢字)後面`))
	suite.Equal(`前面 (中文 123 漢字) tail`, TextSpacing(`前面(中文123漢字) tail`))
	suite.Equal(`head (中文 123 漢字) 後面`, TextSpacing(`head (中文123漢字)後面`))
	suite.Equal(`head (中文 123 漢字) tail`, TextSpacing(`head (中文123漢字) tail`))
}

func (suite *PanguTestSuite) TestMinus() {
	suite.Equal(`前面 - 後面`, TextSpacing(`前面-後面`))
	suite.Equal(`前面 - 後面`, TextSpacing(`前面 - 後面`))

	suite.Equal(`Vinta - Mollie`, TextSpacing(`Vinta-Mollie`))
	suite.Equal(`Vinta - 陳上進`, TextSpacing(`Vinta-陳上進`))
	suite.Equal(`陳上進 - Vinta`, TextSpacing(`陳上進-Vinta`))

	suite.Equal(`得到一個 A - B 的結果`, TextSpacing(`得到一個A-B的結果`))
}

func (suite *PanguTestSuite) TestUnderscore() {
	suite.Equal(`前面_後面`, TextSpacing(`前面_後面`))
	suite.Equal(`前面 _ 後面`, TextSpacing(`前面 _ 後面`))
}

func (suite *PanguTestSuite) TestPlus() {
	suite.Equal(`前面 + 後面`, TextSpacing(`前面+後面`))
	suite.Equal(`前面 + 後面`, TextSpacing(`前面 + 後面`))

	suite.Equal(`Vinta + 陳上進`, TextSpacing(`Vinta+陳上進`))
	suite.Equal(`陳上進 + Vinta`, TextSpacing(`陳上進+Vinta`))

	suite.Equal(`得到一個 A + B 的結果`, TextSpacing(`得到一個A+B的結果`))

	suite.Equal(`得到一個 C++ 的結果`, TextSpacing(`得到一個C++的結果`))

	// TODO
	// suite.Equal(`得到一個 A+ 的結果`, TextSpacing(`得到一個A+的結果`))
}

func (suite *PanguTestSuite) TestEqual() {
	suite.Equal(`前面 = 後面`, TextSpacing(`前面=後面`))
	suite.Equal(`前面 = 後面`, TextSpacing(`前面 = 後面`))

	suite.Equal(`Vinta = Mollie`, TextSpacing(`Vinta=Mollie`))
	suite.Equal(`Vinta = 陳上進`, TextSpacing(`Vinta=陳上進`))
	suite.Equal(`陳上進 = Vinta`, TextSpacing(`陳上進=Vinta`))

	suite.Equal(`得到一個 A = B 的結果`, TextSpacing(`得到一個A=B的結果`))
}

func (suite *PanguTestSuite) TestBrace() {
	suite.Equal(`前面 { 後面`, TextSpacing(`前面{後面`))
	suite.Equal(`前面 { 後面`, TextSpacing(`前面 { 後面`))

	suite.Equal(`前面 } 後面`, TextSpacing(`前面}後面`))
	suite.Equal(`前面 } 後面`, TextSpacing(`前面 } 後面`))

	suite.Equal(`前面 {中文 123 漢字} 後面`, TextSpacing(`前面{中文123漢字}後面`))
	suite.Equal(`前面 {中文 123} 後面`, TextSpacing(`前面{中文123}後面`))
	suite.Equal(`前面 {123 漢字} 後面`, TextSpacing(`前面{123漢字}後面`))
	suite.Equal(`前面 {中文 123 漢字} tail`, TextSpacing(`前面{中文123漢字} tail`))
	suite.Equal(`head {中文 123 漢字} 後面`, TextSpacing(`head {中文123漢字}後面`))
	suite.Equal(`head {中文 123 漢字} tail`, TextSpacing(`head {中文123漢字} tail`))
}

func (suite *PanguTestSuite) TestBracket() {
	suite.Equal(`前面 [ 後面`, TextSpacing(`前面[後面`))
	suite.Equal(`前面 [ 後面`, TextSpacing(`前面 [ 後面`))

	suite.Equal(`前面 ] 後面`, TextSpacing(`前面]後面`))
	suite.Equal(`前面 ] 後面`, TextSpacing(`前面 ] 後面`))

	suite.Equal(`前面 [中文 123 漢字] 後面`, TextSpacing(`前面[中文123漢字]後面`))
	suite.Equal(`前面 [中文 123] 後面`, TextSpacing(`前面[中文123]後面`))
	suite.Equal(`前面 [123 漢字] 後面`, TextSpacing(`前面[123漢字]後面`))
	suite.Equal(`前面 [中文 123 漢字] tail`, TextSpacing(`前面[中文123漢字] tail`))
	suite.Equal(`head [中文 123 漢字] 後面`, TextSpacing(`head [中文123漢字]後面`))
	suite.Equal(`head [中文 123 漢字] tail`, TextSpacing(`head [中文123漢字] tail`))
}

func (suite *PanguTestSuite) TestPipe() {
	suite.Equal(`前面 | 後面`, TextSpacing(`前面|後面`))
	suite.Equal(`前面 | 後面`, TextSpacing(`前面 | 後面`))

	suite.Equal(`Vinta | Mollie`, TextSpacing(`Vinta|Mollie`))
	suite.Equal(`Vinta | 陳上進`, TextSpacing(`Vinta|陳上進`))
	suite.Equal(`陳上進 | Vinta`, TextSpacing(`陳上進|Vinta`))

	suite.Equal(`得到一個 A | B 的結果`, TextSpacing(`得到一個A|B的結果`))
}

func (suite *PanguTestSuite) TestBackslash() {
	suite.Equal(`前面 \ 後面`, TextSpacing(`前面\後面`))
}

func (suite *PanguTestSuite) TestColon() {
	suite.Equal(`前面: 後面`, TextSpacing(`前面:後面`))

	suite.Equal(`前面 : 後面`, TextSpacing(`前面 : 後面`))
}

func (suite *PanguTestSuite) TestSemicolon() {
	suite.Equal(`前面; 後面`, TextSpacing(`前面;後面`))

	suite.Equal(`前面 ; 後面`, TextSpacing(`前面 ; 後面`))
}

func (suite *PanguTestSuite) TestQuote() {
	suite.Equal(`前面 " 後面`, TextSpacing(`前面"後面`))

	suite.Equal(`前面 "中文 123 漢字" 後面`, TextSpacing(`前面"中文123漢字"後面`))

	suite.Equal(`前面 "" 後面`, TextSpacing(`前面""後面`))
}

func (suite *PanguTestSuite) TestSingleQuote() {
	suite.Equal(`前面 ' 後面`, TextSpacing(`前面'後面`))

	suite.Equal(`前面 '中文 123 漢字' 後面`, TextSpacing(`前面'中文123漢字'後面`))

	suite.Equal(`前面 '' 後面`, TextSpacing(`前面''後面`))
	suite.Equal(`前面 ' ' 後面`, TextSpacing(`前面' '後面`))

	suite.Equal(`陳上進's 丸子`, TextSpacing(`陳上進's 丸子`))
}

func (suite *PanguTestSuite) TestLessThan() {
	suite.Equal(`前面 < 後面`, TextSpacing(`前面<後面`))
	suite.Equal(`前面 < 後面`, TextSpacing(`前面 < 後面`))

	suite.Equal(`Vinta < Mollie`, TextSpacing(`Vinta<Mollie`))
	suite.Equal(`Vinta < 陳上進`, TextSpacing(`Vinta<陳上進`))
	suite.Equal(`陳上進 < Vinta`, TextSpacing(`陳上進<Vinta`))

	suite.Equal(`得到一個 A < B 的結果`, TextSpacing(`得到一個A<B的結果`))

	// TODO
	// suite.Equal(`前面 <中文 123 漢字> 後面`, TextSpacing(`前面<中文123漢字>後面`))
	// suite.Equal(`前面 <中文 123> 後面`, TextSpacing(`前面<中文123>後面`))
	// suite.Equal(`前面 <123 漢字> 後面`, TextSpacing(`前面<123漢字>後面`))
	// suite.Equal(`前面 <中文 123 漢字> tail`, TextSpacing(`前面<中文123漢字> tail`))
	// suite.Equal(`head <中文 123 漢字> 後面`, TextSpacing(`head <中文123漢字>後面`))
	// suite.Equal(`head <中文 123 漢字> tail`, TextSpacing(`head <中文123漢字> tail`))
}

func (suite *PanguTestSuite) TestComma() {
	suite.Equal(`前面, 後面`, TextSpacing(`前面,後面`))
	suite.Equal(`前面 , 後面`, TextSpacing(`前面 , 後面`))
}

func (suite *PanguTestSuite) TestGreaterThan() {
	suite.Equal(`前面 > 後面`, TextSpacing(`前面>後面`))
	suite.Equal(`前面 > 後面`, TextSpacing(`前面 > 後面`))

	suite.Equal(`Vinta > Mollie`, TextSpacing(`Vinta>Mollie`))
	suite.Equal(`Vinta > 陳上進`, TextSpacing(`Vinta>陳上進`))
	suite.Equal(`陳上進 > Vinta`, TextSpacing(`陳上進>Vinta`))

	suite.Equal(`得到一個 A > B 的結果`, TextSpacing(`得到一個A>B的結果`))
}

func (suite *PanguTestSuite) TestPeriod() {
	suite.Equal(`前面. 後面`, TextSpacing(`前面.後面`))
	suite.Equal(`前面 . 後面`, TextSpacing(`前面 . 後面`))
}

func (suite *PanguTestSuite) TestQuestionMark() {
	suite.Equal(`前面? 後面`, TextSpacing(`前面?後面`))
	suite.Equal(`前面 ? 後面`, TextSpacing(`前面 ? 後面`))
}

func (suite *PanguTestSuite) TestSlash() {
	suite.Equal(`前面 / 後面`, TextSpacing(`前面/後面`))
	suite.Equal(`前面 / 後面`, TextSpacing(`前面 / 後面`))

	suite.Equal(`Vinta / Mollie`, TextSpacing(`Vinta/Mollie`))
	suite.Equal(`Vinta / 陳上進`, TextSpacing(`Vinta/陳上進`))
	suite.Equal(`陳上進 / Vinta`, TextSpacing(`陳上進/Vinta`))

	suite.Equal(`得到一個 A / B 的結果`, TextSpacing(`得到一個A/B的結果`))
}

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestPanguTestSuite(t *testing.T) {
	suite.Run(t, new(PanguTestSuite))
}
