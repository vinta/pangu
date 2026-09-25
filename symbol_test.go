package pangu_test

import "testing"

func TestSymbolAmpersand(t *testing.T) {
	t.Run("handle & symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面&後面", "前面 & 後面"},
			{"Vinta&陳上進", "Vinta & 陳上進"},
			{"陳上進&Vinta", "陳上進 & Vinta"},
			// DO NOT change if already spacing
			{"前面 & 後面", "前面 & 後面"},
			{"Vinta & Abc123", "Vinta & Abc123"},
			{"Vinta & 陳上進", "Vinta & 陳上進"},
			{"陳上進 & Vinta", "陳上進 & Vinta"},
			{"得到一個 A & B 的結果", "得到一個 A & B 的結果"},
		})
	})
	t.Run("handle & symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta&Abc123", "Vinta&Abc123"},
			{"得到一個A&B的結果", "得到一個 A&B 的結果"},
			{"本週S&P 500及Nasdaq同時下跌", "本週 S&P 500 及 Nasdaq 同時下跌"},
			{"接下來是Q&A時間", "接下來是 Q&A 時間"},
		})
	})
}

func TestSymbolAngleBrackets(t *testing.T) {
	t.Run("handle < > symbols as angle brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面<中文123漢字>後面", "前面 <中文 123 漢字> 後面"},
			{"前面<中文123>後面", "前面 <中文 123> 後面"},
			{"前面<123漢字>後面", "前面 <123 漢字> 後面"},
			{"前面<中文123> tail", "前面 <中文 123> tail"},
			{"head <中文123漢字>後面", "head <中文 123 漢字> 後面"},
			{"head <中文123漢字> tail", "head <中文 123 漢字> tail"},
		})
	})
	t.Run("handle < > symbols as HTML tags", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"<p>一行文本</p>", "<p>一行文本</p>"},
			{"<p>文字<strong>加粗</strong></p>", "<p>文字<strong>加粗</strong></p>"},
			{"<div>測試<span>內容</span>結束</div>", "<div>測試<span>內容</span>結束</div>"},
			{"<a href=\"#\">連結</a>", "<a href=\"#\">連結</a>"},
			{"<input value=\"測試123\">", "<input value=\"測試 123\">"},
			{"<img src=\"test.jpg\" alt=\"測試圖片\">", "<img src=\"test.jpg\" alt=\"測試圖片\">"},
			// Multiple tags
			{"<p>第一段</p><p>第二段</p>", "<p>第一段</p><p>第二段</p>"},
			{"<h1>標題</h1><p>內容</p>", "<h1>標題</h1><p>內容</p>"},
			// Nested tags
			{"<div><p>嵌套<strong>測試</strong></p></div>", "<div><p>嵌套<strong>測試</strong></p></div>"},
			// <br> or <hr> must stay untouched
			{"文字<br>換行", "文字<br>換行"},
			{"文字<br />換行", "文字<br />換行"},
			{"第一段<hr>第二段", "第一段<hr>第二段"},
			{"第一段<hr />第二段", "第一段<hr />第二段"},
			// Real-world markup stays untouched
			{"<ul><li>第一項</li><li>第二項</li></ul>", "<ul><li>第一項</li><li>第二項</li></ul>"},
			{"<button disabled>送出表單</button>", "<button disabled>送出表單</button>"},
			{"<img src=\"photo.jpg\">上面是圖片", "<img src=\"photo.jpg\">上面是圖片"},
			{"<attackOnJava>那一天，人類終於回想起了，曾經一度被XML所支配的恐懼</attackOnJava> <!-- 進擊的Java -->", "<attackOnJava>那一天，人類終於回想起了，曾經一度被 XML 所支配的恐懼</attackOnJava> <!-- 進擊的 Java -->"},
		})
	})
	t.Run("handle < > symbols as tag mention", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"在這裡插入一個<div>標籤", "在這裡插入一個 <div> 標籤"},
			{"型別是List<String>的容器", "型別是 List<String> 的容器"},
			{"把文字包在<span>裡面", "把文字包在 <span> 裡面"},
			{"每個<li>代表一個列表項目", "每個 <li> 代表一個列表項目"},
			{"用<table>排版是過時的做法", "用 <table> 排版是過時的做法"},
			{"HTML的<head>放的是metadata", "HTML 的 <head> 放的是 metadata"},
			{"<html>是整份文件的根元素", "<html> 是整份文件的根元素"},
			{"這裡放<Spinner />元件", "這裡放 <Spinner /> 元件"},
			// Generic type parameters read as tag mentions too
			{"回傳Promise<string>就好", "回傳 Promise<string> 就好"},
			{"用Vec<u8>儲存位元組", "用 Vec<u8> 儲存位元組"},
			{"先引入<iostream>標頭檔", "先引入 <iostream> 標頭檔"},
			// A mention next to real markup: only the mention is spaced
			{"<p>用<code>標記程式碼</p>", "<p>用 <code> 標記程式碼</p>"},
		})
	})
}

func TestSymbolAsterisk(t *testing.T) {
	t.Run("handle * symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面*後面", "前面 * 後面"},
			{"Vinta*陳上進", "Vinta * 陳上進"},
			{"陳上進*Vinta", "陳上進 * Vinta"},
			{"標示*的欄位代表必填", "標示 * 的欄位代表必填"},
			{"時薪*(平日時數+假日時數)", "時薪 * (平日時數 + 假日時數)"},
			// Rare cases, ignore
			// {"時薪*[平日時數+假日時數]", "時薪 * [平日時數 + 假日時數]"},
			// {"係數*[2+3]", "係數 * [2+3]"},
			// {"係數*[abc]", "係數 * [abc]"},
			// {"係數*[[0-9].log]", "係數 * [[0-9].log]"},
			// DO NOT change if already spacing
			{"前面 * 後面", "前面 * 後面"},
			{"Vinta * Abc123", "Vinta * Abc123"},
			{"Vinta * 陳上進", "Vinta * 陳上進"},
			{"陳上進 * Vinta", "陳上進 * Vinta"},
			{"得到一個 A * B 的結果", "得到一個 A * B 的結果"},
		})
	})
	t.Run("handle * symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta*Abc123", "Vinta*Abc123"},
			{"得到一個A*B的結果", "得到一個 A*B 的結果"},
			{"算式是2*3的積", "算式是 2*3 的積"},
		})
	})
	t.Run("handle * symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"刪掉*.log的檔案", "刪掉 *.log 的檔案"},
		})
	})
	t.Run("preserve bracket globs", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"刪掉*[0-9].log的檔案", "刪掉 *[0-9].log 的檔案"},
			{"刪掉*[a-z].log的檔案", "刪掉 *[a-z].log 的檔案"},
			{"刪掉*[!0-9].log的檔案", "刪掉 *[!0-9].log 的檔案"},
			{"刪掉*[0-9].tar.gz的檔案", "刪掉 *[0-9].tar.gz 的檔案"},
			{"刪掉*[0-9][0-9].log的檔案", "刪掉 *[0-9][0-9].log 的檔案"},
			{"刪掉*[0-9]*.log的檔案", "刪掉 *[0-9]*.log 的檔案"},
			// DO NOT change if already spacing
			{"刪掉 *[0-9].log 的檔案", "刪掉 *[0-9].log 的檔案"},
			{"刪掉 *[a-z].log 的檔案", "刪掉 *[a-z].log 的檔案"},
			{"刪掉 *[!0-9].log 的檔案", "刪掉 *[!0-9].log 的檔案"},
			{"刪掉 *[0-9].tar.gz 的檔案", "刪掉 *[0-9].tar.gz 的檔案"},
			{"刪掉 *[0-9][0-9].log 的檔案", "刪掉 *[0-9][0-9].log 的檔案"},
			{"刪掉 *[0-9]*.log 的檔案", "刪掉 *[0-9]*.log 的檔案"},
		})
	})
}

func TestSymbolAt(t *testing.T) {
	t.Run("handle @ symbol as at", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面@vinta後面", "前面 @vinta 後面"},
			{"前面@vinta_chen後面", "前面 @vinta_chen 後面"},
			{"前面@VintaChen後面", "前面 @VintaChen 後面"},
			{"前面@陳上進 後面", "前面 @陳上進 後面"},
		})
	})
}

func TestSymbolBackslash(t *testing.T) {
	t.Run("handle \\ symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面\\後面", "前面 \\ 後面"},
			{"前面 \\ 後面", "前面 \\ 後面"},
		})
	})
	t.Run("handle \\ symbol as escape character", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"\\n", "\\n"},
			{"\\t", "\\t"},
		})
	})
	t.Run("handle \\ symbol as Windows file path", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"檔案在C:\\Users\\name\\", "檔案在 C:\\Users\\name\\"},
			{"程式在D:\\Program Files\\", "程式在 D:\\Program Files\\"},
			{"在C:\\Windows\\System32", "在 C:\\Windows\\System32"},
		})
	})
}

func TestSymbolBacktick(t *testing.T) {
	t.Run("handle ` ` symbols as quotes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面`中間`後面", "前面 `中間` 後面"},
			{"`! git commit -a -m \"蛤\"`", "`! git commit -a -m \"蛤\"`"},
			{"从结果来看，当a.b销毁后，`a.getB()`返回值为null", "从结果来看，当 a.b 销毁后，`a.getB()` 返回值为 null"},
			{"雖然知道可以在Claude Code直接執行shell指令，例如`! git commit -a -m \"蛤\"`，但是看了文件才知道原來在 http://command.md 裡面也可以用`!`啊#TIL", "雖然知道可以在 Claude Code 直接執行 shell 指令，例如 `! git commit -a -m \"蛤\"`，但是看了文件才知道原來在 http://command.md 裡面也可以用 `!` 啊 #TIL"},
			{"雖然知道可以在 Claude Code 直接執行 shell 指令，例如 `! git commit -a -m \"蛤\"`，但是看了文件才知道原來在 http://command.md 裡面也可以用 `!` 啊 #TIL", "雖然知道可以在 Claude Code 直接執行 shell 指令，例如 `! git commit -a -m \"蛤\"`，但是看了文件才知道原來在 http://command.md 裡面也可以用 `!` 啊 #TIL"},
		})
	})
}

func TestSymbolCaret(t *testing.T) {
	t.Run("handle ^ symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面^後面", "前面 ^ 後面"},
			{"前面 ^ 後面", "前面 ^ 後面"},
		})
	})
}

func TestSymbolColon(t *testing.T) {
	t.Run("handle : symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面:後面", "前面: 後面"},
			{"電話:123456789", "電話: 123456789"},
			{"前面:I have no idea後面", "前面: I have no idea 後面"},
			// DO NOT change if already spacing
			{"前面 : 後面", "前面 : 後面"},
			{"前面: 後面", "前面: 後面"},
			{"前面 :後面", "前面 :後面"},
			{"前面: I have no idea後面", "前面: I have no idea 後面"},
		})
	})
	// FIXME: See https://github.com/vinta/pangu.js/issues/316
	t.Run("handle : symbol as emoticon", func(t *testing.T) {
		testSpaceTextFails(t, []spaceTextCase{
			{"前面:)後面", "前面 :) 後面"},
		})
	})
	// FIXME: But rare cases, I suppose?
	t.Run("handle : symbol as separator", func(t *testing.T) {
		testSpaceTextFails(t, []spaceTextCase{
			{"前面:後面:再後面", "前面:後面:再後面"},
			{"前面:後面:再後面:更後面", "前面:後面:再後面:更後面"},
			{"前面:後面:再後面:更後面:超後面", "前面:後面:再後面:更後面:超後面"},
		})
	})
}

func TestSymbolComma(t *testing.T) {
	t.Run("handle , symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面,後面", "前面, 後面"},
			{"\"你好\",她說", "\"你好\", 她說"},
			{"每月只要1,000元", "每月只要 1,000 元"},
			{"精采5G購機方案(30個月),月繳599元購機優惠(30個月)", "精采 5G 購機方案 (30 個月), 月繳 599 元購機優惠 (30 個月)"},
			// DO NOT change if already spacing
			{"前面 , 後面", "前面 , 後面"},
			{"前面, 後面", "前面, 後面"},
			// Rare cases, ignore
			// {"前面 ,後面", "前面 ,後面"},
		})
	})
}

func TestSymbolCurlyBrackets(t *testing.T) {
	t.Run("handle { } symbols as curly brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面{中文123漢字}後面", "前面 {中文 123 漢字} 後面"},
			{"前面{中文123}後面", "前面 {中文 123} 後面"},
			{"前面{123漢字}後面", "前面 {123 漢字} 後面"},
			{"前面{中文123} tail", "前面 {中文 123} tail"},
			{"head {中文123漢字}後面", "head {中文 123 漢字} 後面"},
			{"head {中文123漢字} tail", "head {中文 123 漢字} tail"},
		})
	})
	t.Run("handle multiline content in curly brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// A space before a newline is mid-content, not a bracket-edge space: only the literal string edges get stripped
			{"{x \n}中", "{x \n} 中"},
		})
	})
}

func TestSymbolDashes(t *testing.T) {
	// \u2014
	// FIXME
	t.Run("Symbol —", func(t *testing.T) {
		t.Run("handle — symbol", func(t *testing.T) {
			testSpaceTextFails(t, []spaceTextCase{
				{"他說——不對", "他說 —— 不對"},
				{"台灣——美麗之島", "台灣 —— 美麗之島"},
				{"他說———不對", "他說 ——— 不對"},
				{"他說 —— 不對", "他說 —— 不對"},
			})
		})
		// No CJK contact, no change
		t.Run("keep — between ANS untouched", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"A—B", "A—B"},
				{"2020—2024年", "2020—2024 年"},
			})
		})
	})
	// \u2500
	// FIXME
	t.Run("Symbol ─", func(t *testing.T) {
		t.Run("handle ─ symbol", func(t *testing.T) {
			testSpaceTextFails(t, []spaceTextCase{
				{"他說──不對", "他說 ── 不對"},
				{"於是──各位觀眾", "於是 ── 各位觀眾"},
				{"於是 ── 各位觀眾", "於是 ── 各位觀眾"},
			})
		})
		// No CJK contact, no change
		t.Run("keep ─ between ANS untouched", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"A─B", "A─B"},
				{"2020──2024", "2020──2024"},
			})
		})
	})
}

func TestSymbolDollarSign(t *testing.T) {
	t.Run("handle $ symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面$後面", "前面 $ 後面"},
			{"前面 $ 後面", "前面 $ 後面"},
			{"前面$100後面", "前面 $100 後面"},
		})
	})
}

func TestSymbolDoubleQuotes(t *testing.T) {
	t.Run("Symbol \" \"", func(t *testing.T) {
		t.Run("handle \" \" symbols as quotes around CJK", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面\"中文123漢字\"後面", "前面 \"中文 123 漢字\" 後面"},
				{"前面\"中文123\"後面", "前面 \"中文 123\" 後面"},
				{"前面\"中文abc\"後面", "前面 \"中文 abc\" 後面"},
				{"前面\"123漢字\"後面", "前面 \"123 漢字\" 後面"},
				{"前面\"中文123\" tail", "前面 \"中文 123\" tail"},
				{"head \"中文123漢字\"後面", "head \"中文 123 漢字\" 後面"},
				{"head \"中文123漢字\" tail", "head \"中文 123 漢字\" tail"},
			})
		})
		t.Run("handle separator spacing inside quotes", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"\"字+\"", "\"字 +\""},
				{"\"字|\"", "\"字 |\""},
				{"你好\"字+\"世界", "你好 \"字 +\" 世界"},
				{"前面\"字|\"後面", "前面 \"字 |\" 後面"},
				{"多行\"字+\"\n下行\"字|\"", "多行 \"字 +\"\n下行 \"字 |\""},
			})
		})
		t.Run("handle \" \" symbols as quotes around English", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"我們也不可以說\"We invited the reverend to dinner.\"", "我們也不可以說 \"We invited the reverend to dinner.\""},
				{"\"We invited the Rev. Darling.\"我們也不可以說", "\"We invited the Rev. Darling.\" 我們也不可以說"},
				{"它應該這樣使用：\"We invited\"", "它應該這樣使用：\"We invited\""},
				{"\"! git commit -a -m '蛤'\"", "\"! git commit -a -m '蛤'\""},
				{"Rev. (Reverend；牧師的尊稱)這個縮寫嚴格來說並不是一項頭銜，而是形容詞。所以，它應該這樣使用：\"We invited the Rev. Alan Darling.\" 或\u00a0 \"We\u00a0invited the Rev. Mr. Darling.\"，而非\"We invited the Rev. Darling.\"我們也不可以說\"We invited the reverend to dinner.\" -- Only a cad would invite the rev. (只有下流的人才會招致批評：句中的 rev. 是 review 的縮寫，算是雙關語)", "Rev. (Reverend；牧師的尊稱) 這個縮寫嚴格來說並不是一項頭銜，而是形容詞。所以，它應該這樣使用：\"We invited the Rev. Alan Darling.\" 或\u00a0 \"We\u00a0invited the Rev. Mr. Darling.\"，而非 \"We invited the Rev. Darling.\" 我們也不可以說 \"We invited the reverend to dinner.\" -- Only a cad would invite the rev. (只有下流的人才會招致批評：句中的 rev. 是 review 的縮寫，算是雙關語)"},
			})
		})
		t.Run("handle \" \" across the line breaks of a wrapped HTML source", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"使用：\"We\ninvited Darling.\" 或 \"We invited.\"", "使用：\"We\ninvited Darling.\" 或 \"We invited.\""},
				{"Rev. (Reverend；牧師的尊稱) \n    這個縮寫嚴格來說並不是一項頭銜，而是形容詞。所以，它應該這樣使用：\"We \n    invited the Rev. Alan Darling.\" 或\u00a0 \"We\u00a0 invited the Rev. Mr. \n    Darling.\" ，而非 \"We invited the Rev. Darling.\" 我們也不可以說\u00a0 \n    \"We invited the reverend to dinner.\" -- Only a cad would invite the rev. (只有下流的人才會招致批評：句中的 \n    rev. 是 review 的縮寫，算是雙關語) ", "Rev. (Reverend；牧師的尊稱) \n    這個縮寫嚴格來說並不是一項頭銜，而是形容詞。所以，它應該這樣使用：\"We \n    invited the Rev. Alan Darling.\" 或\u00a0 \"We\u00a0 invited the Rev. Mr. \n    Darling.\" ，而非 \"We invited the Rev. Darling.\" 我們也不可以說\u00a0 \n    \"We invited the reverend to dinner.\" -- Only a cad would invite the rev. (只有下流的人才會招致批評：句中的 \n    rev. 是 review 的縮寫，算是雙關語) "},
			})
		})
		// Rare cases, ignore
		// See https://github.com/vinta/pangu.js/issues/287
		// t.Run("handle \" \" mis-pairing (known limitation)", func(t *testing.T) {
		// 	testSpaceText(t, []spaceTextCase{
		// 		{"Darling.\" 或 \"We", "Darling.\" 或 \"We"},
		// 	})
		// })
	})
	// \u201c
	// \u201d
	t.Run("Symbol “ ”", func(t *testing.T) {
		t.Run("handle “ ” symbols as quotes", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"獲標準普爾長期信用評等“AA”全球電信業之首", "獲標準普爾長期信用評等 “AA” 全球電信業之首"},
				{"阿里云开源“计算王牌”Blink，实时计算时代已来", "阿里云开源 “计算王牌” Blink，实时计算时代已来"},
				{"苹果撤销Facebook“企业证书”后者股价一度短线走低", "苹果撤销 Facebook “企业证书” 后者股价一度短线走低"},
				{"【UCG中字】“數毛社”DF的《戰神4》全新演示解析", "【UCG 中字】“數毛社” DF 的《戰神 4》全新演示解析"},
			})
		})
		t.Run("handle misused ” ” quote pairs", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"他说”你好”啊", "他说 ”你好” 啊"},
				{"《战斧骨》里还有个镜头挺有意思，就是男主”见路不走”，不从峡谷入口走，而选择了从侧面翻越，还顺便借着口哨吸引出来一个食人族给杀了。", "《战斧骨》里还有个镜头挺有意思，就是男主 ”见路不走”，不从峡谷入口走，而选择了从侧面翻越，还顺便借着口哨吸引出来一个食人族给杀了。"},
			})
		})
	})
}

func TestSymbolEllipsis(t *testing.T) {
	// \u2026
	t.Run("handle … symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面…後面", "前面… 後面"},
			{"前面……後面", "前面…… 後面"},
		})
	})
}

func TestSymbolEqualsSign(t *testing.T) {
	t.Run("handle = symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面=後面", "前面 = 後面"},
			{"Vinta=陳上進", "Vinta = 陳上進"},
			{"陳上進=Vinta", "陳上進 = Vinta"},
			{"年增率=(今年-去年)", "年增率 = (今年 - 去年)"},
			{"總價=(單價*數量)", "總價 = (單價 * 數量)"},
			// DO NOT change if already spacing
			{"前面 = 後面", "前面 = 後面"},
			{"Vinta = Abc123", "Vinta = Abc123"},
			{"Vinta = 陳上進", "Vinta = 陳上進"},
			{"陳上進 = Vinta", "陳上進 = Vinta"},
			{"得到一個 A = B 的結果", "得到一個 A = B 的結果"},
		})
	})
	t.Run("handle = symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta=Abc123", "Vinta=Abc123"},
			{"得到一個A=B的結果", "得到一個 A=B 的結果"},
			{"設定a=1之後執行", "設定 a=1 之後執行"},
			{"網址是example.com?foo=bar&baz=1的頁面", "網址是 example.com?foo=bar&baz=1 的頁面"},
		})
	})
	t.Run("handle = symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"用=>寫箭頭函式", "用 => 寫箭頭函式"},
		})
	})
}

func TestSymbolExclamationMark(t *testing.T) {
	t.Run("handle ! symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面!", "前面!"},
			{"前面!!", "前面!!"},
			{"前面!!!", "前面!!!"},
			{"前面!後面", "前面! 後面"},
			{"前面!!後面", "前面!! 後面"},
			{"前面!!!後面", "前面!!! 後面"},
			{"前面!abc", "前面! abc"},
			{"前面!123", "前面! 123"},
			{"前面2!的階乘", "前面 2! 的階乘"},
			{"你還在用Yahoo!奇摩？", "你還在用 Yahoo! 奇摩？"},
			{"! git commit -a -m \"蛤\"", "! git commit -a -m \"蛤\""},
			// DO NOT change if already spacing
			{"前面 ! 後面", "前面 ! 後面"},
			{"前面! 後面", "前面! 後面"},
			// Rare cases, ignore
			// {"前面 !後面", "前面 !後面"},
		})
	})
}

func TestSymbolGreaterThanSign(t *testing.T) {
	t.Run("handle > symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面>後面", "前面 > 後面"},
			{"Vinta>陳上進", "Vinta > 陳上進"},
			{"陳上進>Vinta", "陳上進 > Vinta"},
			{"溫度>30就開冷氣", "溫度 > 30 就開冷氣"},
			// DO NOT change if already spacing
			{"前面 > 後面", "前面 > 後面"},
			{"Vinta > Abc123", "Vinta > Abc123"},
			{"Vinta > 陳上進", "Vinta > 陳上進"},
			{"陳上進 > Vinta", "陳上進 > Vinta"},
			{"得到一個 A > B 的結果", "得到一個 A > B 的結果"},
		})
	})
	t.Run("handle > symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta>Abc123", "Vinta>Abc123"},
			{"得到一個A>B的結果", "得到一個 A>B 的結果"},
		})
	})
	t.Run("handle > symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"流程是A->B的方向", "流程是 A->B 的方向"},
		})
	})
}

func TestSymbolHashtag(t *testing.T) {
	t.Run("handle # symbol as hashtag", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面#後面", "前面 #後面"},
			{"前面#H2G2後面", "前面 #H2G2 後面"},
			{"前面 #銀河便車指南 後面", "前面 #銀河便車指南 後面"},
			{"前面#銀河便車指南 後面", "前面 #銀河便車指南 後面"},
			{"前面#銀河公車指南 #銀河拖吊車指南 後面", "前面 #銀河公車指南 #銀河拖吊車指南 後面"},
		})
	})
	t.Run("handle # symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面C#後面", "前面 C# 後面"},
			{"前面F#後面", "前面 F# 後面"},
			{"前端/後端/資料庫：C#和Python", "前端/後端/資料庫：C# 和 Python"},
		})
	})
	t.Run("handle # symbol as hashtag in a slash list", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"dae-dae-o/#絕地家庭小會議/#今天大掃除了沒有/", "dae-dae-o/#絕地家庭小會議/#今天大掃除了沒有/"},
		})
	})
}

func TestSymbolLessThanSign(t *testing.T) {
	t.Run("handle < symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面<後面", "前面 < 後面"},
			{"Vinta<陳上進", "Vinta < 陳上進"},
			{"陳上進<Vinta", "陳上進 < Vinta"},
			// DO NOT change if already spacing
			{"前面 < 後面", "前面 < 後面"},
			{"Vinta < Abc123", "Vinta < Abc123"},
			{"Vinta < 陳上進", "Vinta < 陳上進"},
			{"陳上進 < Vinta", "陳上進 < Vinta"},
			{"得到一個 A < B 的結果", "得到一個 A < B 的結果"},
		})
	})
	t.Run("handle < symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta<Abc123", "Vinta<Abc123"},
			{"得到一個A<B的結果", "得到一個 A<B 的結果"},
			{"如果A<B就繼續", "如果 A<B 就繼續"},
			{"條件是1<2的情況", "條件是 1<2 的情況"},
		})
	})
}

func TestSymbolMiddleDot(t *testing.T) {
	// \u00b7
	t.Run("Symbol ·", func(t *testing.T) {
		t.Run("handle · symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面·後面", "前面・後面"},
				{"喬治·R·R·馬丁", "喬治・R・R・馬丁"},
				{"M·奈特·沙马兰", "M・奈特・沙马兰"},
			})
		})
		t.Run("should not convert if already spaced", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"哥爾 · D · 羅傑", "哥爾 · D · 羅傑"},
				{"看过 · · ·", "看过 · · ·"},
				{"看过 · · · (2026部)", "看过 · · · (2026 部)"},
			})
		})
	})
	// \u2022
	t.Run("Symbol •", func(t *testing.T) {
		t.Run("handle • symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面•後面", "前面・後面"},
				{"喬治•R•R•馬丁", "喬治・R・R・馬丁"},
				{"M•奈特•沙马兰", "M・奈特・沙马兰"},
			})
		})
		t.Run("should not convert if already spaced", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面 • 後面", "前面 • 後面"},
				{"喬治 • R • R • 馬丁", "喬治 • R • R • 馬丁"},
				{"M • 奈特 • 沙马兰", "M • 奈特 • 沙马兰"},
			})
		})
		t.Run("should not convert consecutive • symbols", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"國泰CUBE卡 •••• 1234", "國泰 CUBE 卡 •••• 1234"},
				{"ether.fi Cash Card •••• 5678", "ether.fi Cash Card •••• 5678"},
			})
		})
	})
	// \u2027
	t.Run("Symbol ‧", func(t *testing.T) {
		t.Run("handle ‧ symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面‧後面", "前面・後面"},
				{"喬治‧R‧R‧馬丁", "喬治・R・R・馬丁"},
				{"M‧奈特‧沙马兰", "M・奈特・沙马兰"},
			})
		})
		t.Run("should not convert if already spaced", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面 ‧ 後面", "前面 ‧ 後面"},
				{"喬治 ‧ R ‧ R ‧ 馬丁", "喬治 ‧ R ‧ R ‧ 馬丁"},
				{"M ‧ 奈特 ‧ 沙马兰", "M ‧ 奈特 ‧ 沙马兰"},
			})
		})
	})
}

func TestSymbolMinusSign(t *testing.T) {
	t.Run("handle - symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面-後面", "前面 - 後面"},
			{"Vinta-陳上進", "Vinta - 陳上進"},
			{"陳上進-Vinta", "陳上進 - Vinta"},
			{"博客來-Rewire-神經可塑性：用神經科學突破行為模式迴圈，終結焦慮、恐慌和憂鬱，實現最佳的心理健康", "博客來 - Rewire - 神經可塑性：用神經科學突破行為模式迴圈，終結焦慮、恐慌和憂鬱，實現最佳的心理健康"},
			{"博客來-經濟學原理 10/e Mankiw (授權經銷版)", "博客來 - 經濟學原理 10/e Mankiw (授權經銷版)"},
			{"財政部電子發票整合服務平台[自然人憑證]-歸戶設定通知", "財政部電子發票整合服務平台 [自然人憑證] - 歸戶設定通知"},
			{"博客來-4%法則：讓錢活得比你久的提領金律(電子書)", "博客來 - 4% 法則：讓錢活得比你久的提領金律 (電子書)"},
			{"长者的智慧和复杂的维斯特洛- 文章", "长者的智慧和复杂的维斯特洛 - 文章"},
			{"1976年-2018年", "1976 年 - 2018 年"},
			// Regex-boundary cases, no real text found
			{"年增率-(GDP)", "年增率 - (GDP)"},
			{"年增率-[GDP]", "年增率 - [GDP]"},
			{"年增率-(季調)", "年增率 - (季調)"},
			// Hyphen reading: a hyphen in direct contact with CJK flips the hyphens between brackets on its line
			{"全球-實質國內生產毛額[GDP]-(年增率, IMF 預估)", "全球 - 實質國內生產毛額 [GDP] - (年增率, IMF 預估)"},
			{"全球-名目國內生產毛額[GDP]-(NSA,美元,IMF 預估)", "全球 - 名目國內生產毛額 [GDP] - (NSA, 美元, IMF 預估)"},
			{"台灣-消費者物價指數[CPI]-(年增率)", "台灣 - 消費者物價指數 [CPI] - (年增率)"},
			{"美國-核心消費者物價指數[Core CPI]-(SA,年增率)", "美國 - 核心消費者物價指數 [Core CPI] - (SA, 年增率)"},
			{"美國-個人消費支出物價指數[PCE]-(年增率)-第10百分位數", "美國 - 個人消費支出物價指數 [PCE] - (年增率) - 第 10 百分位數"},
			// DO NOT change if already spacing
			{"前面 - 後面", "前面 - 後面"},
			{"Vinta - Abc123", "Vinta - Abc123"},
			{"Vinta - 陳上進", "Vinta - 陳上進"},
			{"陳上進 - Vinta", "陳上進 - Vinta"},
			{"得到一個 A - B 的結果", "得到一個 A - B 的結果"},
		})
	})
	t.Run("handle - symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta-Abc123", "Vinta-Abc123"},
			{"得到一個A-B的結果", "得到一個 A-B 的結果"},
			{"去5-A教室上課", "去 5-A 教室上課"},
			{"搭2-A的公車", "搭 2-A 的公車"},
			{"範圍是1-10的整數", "範圍是 1-10 的整數"},
			{"用USB-C充電", "用 USB-C 充電"},
			{"照X-RAY檢查", "照 X-RAY 檢查"},
			// No hyphen on the line is in direct contact with CJK
			{"毛額[GDP]-(NSA)", "毛額 [GDP]-(NSA)"},
			// Hyphenated English names
			{"英文姓名須與護照上相同，包含標點符號；範例：王小明，英文名為WANG, HSIAO-MING，請於英文姓(Surname)欄位填入WANG,、英文名(Given Names)欄位填入HSIAO-MING。", "英文姓名須與護照上相同，包含標點符號；範例：王小明，英文名為 WANG, HSIAO-MING，請於英文姓 (Surname) 欄位填入 WANG,、英文名 (Given Names) 欄位填入 HSIAO-MING。"},
		})
	})
	t.Run("handle - symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// Compound words
			{"Sci-Fi", "Sci-Fi"},
			{"X-RAY", "X-RAY"},
			{"USB Type-C", "USB Type-C"},
			{"The company offered a state-of-the-art machine-learning-powered real-time fraud-detection system with end-to-end encryption and cutting-edge performance.", "The company offered a state-of-the-art machine-learning-powered real-time fraud-detection system with end-to-end encryption and cutting-edge performance."},
			{"這間公司提供了一套state-of-the-art、machine-learning-powered的real-time fraud-detection系統，具備end-to-end加密功能以及cutting-edge的效能。", "這間公司提供了一套 state-of-the-art、machine-learning-powered 的 real-time fraud-detection 系統，具備 end-to-end 加密功能以及 cutting-edge 的效能。"},
			{"Anthropic的claude-4-opus模型", "Anthropic 的 claude-4-opus 模型"},
			{"OpenAI的o3-pro模型", "OpenAI 的 o3-pro 模型"},
			{"OpenAI的gpt-4o模型", "OpenAI 的 gpt-4o 模型"},
			{"OpenAI的GPT-5模型", "OpenAI 的 GPT-5 模型"},
			{"Google的gemini-2.5-pro模型", "Google 的 gemini-2.5-pro 模型"},
		})
	})
	t.Run("handle - symbol as affix", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// CLI flags
			{"你可以使用uname -m指令來檢查你的Linux作業系統是32位元或是[敏感词已被屏蔽]位元", "你可以使用 uname -m 指令來檢查你的 Linux 作業系統是 32 位元或是 [敏感词已被屏蔽] 位元"},
			{"參數要加-m的旗標", "參數要加 -m 的旗標"},
			// Grades
			{"得到一個D-的結果", "得到一個 D- 的結果"},
			{"得到一個D--的結果", "得到一個 D-- 的結果"},
			// NOTE: fixed by AI spacing, see browser-extensions/chrome/src/ai-spacing/shapes/hyphen-digit.ts
			// The hyphen sign reading was dropped, CJK-N reads as an operator, see ADR 0015
			// {"氣溫是-5度左右", "氣溫是 -5 度左右"},
			// {"Nasdaq-100本週下跌-13.44%", "Nasdaq-100 本週下跌 -13.44%"},
		})
	})
}

func TestSymbolNBSP(t *testing.T) {
	t.Run("handle solitary &nbsp;", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// The &nbsp; already separates the runs it sits between, so only the genuinely missing 說|We junction gets a space
			{"我們說We\u00a0invited", "我們說 We\u00a0invited"},
			{"第\u00a05\u00a0章", "第\u00a05\u00a0章"},
		})
	})
	t.Run("handle solitary &nbsp; adjacent to a half-width space", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// A doubled gap the author wrote. CSS collapses two half-width spaces but never collapses &nbsp; + space, so this paints wider than one space.
			// Dropping either character would be a rewrite, so both stay
			{"或\u00a0 \"We invited\"", "或\u00a0 \"We invited\""},
		})
	})
	t.Run("handle consecutive &nbsp;", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// Runs of 2+ &nbsp;s are deliberate formatting (e.g. paragraph indentation)
			{"中文\u00a0\u00a0\u00a0\u00a0中文", "中文\u00a0\u00a0\u00a0\u00a0中文"},
		})
	})
	t.Run("handle &nbsp; adjacent to other whitespace", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中文\u00a0\n中文", "中文\u00a0\n中文"},
		})
	})
	t.Run("handle &nbsp; at string boundaries", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"\u00a0中文abc", "\u00a0中文 abc"},
			{"中文abc\u00a0", "中文 abc\u00a0"},
		})
	})
	t.Run("handle &nbsp; separating a hashtag from CJK", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// The hashtag guard has to read an &nbsp; as the gap it is, otherwise the # reads as glued to 台北 and gets split off
			{"台北\u00a0#中文", "台北\u00a0#中文"},
			{"中文#\u00a0abc", "中文#\u00a0abc"},
		})
	})
}

func TestSymbolPercentSign(t *testing.T) {
	t.Run("handle % symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面%後面", "前面 % 後面"},
			{"前面 % 後面", "前面 % 後面"},
			{"前面100%後面", "前面 100% 後面"},
			{"新八的構造成分有95%是眼鏡、3%是水、2%是垃圾", "新八的構造成分有 95% 是眼鏡、3% 是水、2% 是垃圾"},
			{"丹寧控注意Levi's全館任2件25%OFF滿額再享85折！", "丹寧控注意 Levi's 全館任 2 件 25% OFF 滿額再享 85 折！"},
		})
	})
}

func TestSymbolPeriod(t *testing.T) {
	t.Run("handle . symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面.", "前面."},
			{"前面..", "前面.."},
			{"前面...", "前面..."},
			{"前面.後面", "前面. 後面"},
			{"前面..後面", "前面.. 後面"},
			{"前面...後面", "前面... 後面"},
			// DO NOT change if already spacing
			{"前面 . 後面", "前面 . 後面"},
			{"前面. 後面", "前面. 後面"},
			{"前面 .後面", "前面 .後面"},
			// Abbreviations
			{"前面vs.後面", "前面 vs. 後面"},
			{"前面U.S.A.後面", "前面 U.S.A. 後面"},
			{"Mr.龍島主道：「Let's Party!各位高明博雅君子！", "Mr. 龍島主道：「Let's Party! 各位高明博雅君子！"},
			{"Mr.龍島主道:「Let's Party!各位高明博雅君子!", "Mr. 龍島主道:「Let's Party! 各位高明博雅君子!"},
			{"世.界.，草.班.与千.早.爱.音.", "世. 界.，草. 班. 与千. 早. 爱. 音."},
		})
	})
	t.Run("handle . symbol as file extension", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// File extensions should keep spacing
			{"使用Python.py檔案", "使用 Python.py 檔案"},
			{"設定檔.env很重要", "設定檔.env 很重要"},
			{"編輯器.vscode目錄", "編輯器.vscode 目錄"},
			{"黑人問號.jpg後面", "黑人問號.jpg 後面"},
			{"黑人問號.jpg 後面", "黑人問號.jpg 後面"},
			// Multiple dots
			{"檔案package.lock.json存在", "檔案 package.lock.json 存在"},
			// CJK before dot patterns
			{"環境.env", "環境.env"},
			{"測試.test.js", "測試.test.js"},
			{"專案.gitignore", "專案.gitignore"},
			// Mixed patterns
			{"使用環境.env配置", "使用環境.env 配置"},
			{"專案.prettierrc和.eslintrc", "專案.prettierrc 和.eslintrc"},
		})
	})
	t.Run("handle . symbol as version number", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"版本v1.2.3發布了", "版本 v1.2.3 發布了"},
			{"pangu.js v1.2.3橫空出世", "pangu.js v1.2.3 橫空出世"},
			{"pangu.js 1.2.3橫空出世", "pangu.js 1.2.3 橫空出世"},
		})
	})
}

func TestSymbolPipe(t *testing.T) {
	t.Run("handle | symbol as separator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面|後面", "前面 | 後面"},
			{"Vinta|貓咪", "Vinta | 貓咪"},
			{"貓咪|Vinta", "貓咪 | Vinta"},
			{"陳上進|貓咪|Abc123", "陳上進 | 貓咪 | Abc123"},
			{"陳上進|Abc123|貓咪", "陳上進 | Abc123 | 貓咪"},
			{"Abc123|Vinta|貓咪", "Abc123 | Vinta | 貓咪"},
			{"Abc123|陳上進|貓咪", "Abc123 | 陳上進 | 貓咪"},
			{"作詞|林夕", "作詞 | 林夕"},
			{"文|張三 圖|李四", "文 | 張三 圖 | 李四"},
			{"支援的 Apple TV 型號|Disney+ 幫助中心|TW", "支援的 Apple TV 型號 | Disney+ 幫助中心 | TW"},
			// DO NOT change if already spacing
			{"前面 | 後面", "前面 | 後面"},
			{"Vinta | Abc123", "Vinta | Abc123"},
			{"Vinta | Abc123 | Kitten", "Vinta | Abc123 | Kitten"},
			{"陳上進 | 貓咪 | Abc123", "陳上進 | 貓咪 | Abc123"},
			{"陳上進 | Abc123 | 貓咪", "陳上進 | Abc123 | 貓咪"},
			{"Abc123 | Vinta | 貓咪", "Abc123 | Vinta | 貓咪"},
			{"Abc123 | 陳上進 | 貓咪", "Abc123 | 陳上進 | 貓咪"},
		})
	})
	t.Run("handle | symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta|Abc123", "Vinta|Abc123"},
			{"Vinta|Abc123|Kitten", "Vinta|Abc123|Kitten"},
			{"ps aux|grep node", "ps aux|grep node"},
			{"條件是x|y的情況", "條件是 x|y 的情況"},
			{"得到一個A|B的結果", "得到一個 A|B 的結果"},
			{"得到一個A||B的結果", "得到一個 A||B 的結果"},
		})
	})
}

func TestSymbolPlusSign(t *testing.T) {
	t.Run("handle + symbol as operator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面+後面", "前面 + 後面"},
			{"陳上進+Vinta", "陳上進 + Vinta"},
			{"Vinta+陳上進", "Vinta + 陳上進"},
			{"你+我=我們", "你 + 我 = 我們"},
			{"Switch+健身環套組", "Switch + 健身環套組"},
			{"MacBook Air M2+滑鼠組合", "MacBook Air M2 + 滑鼠組合"},
			// DO NOT change if already spacing
			{"前面 + 後面", "前面 + 後面"},
			{"Vinta + Abc123", "Vinta + Abc123"},
			{"Vinta + 陳上進", "Vinta + 陳上進"},
			{"陳上進 + Vinta", "陳上進 + Vinta"},
			{"得到一個 A + B 的結果", "得到一個 A + B 的結果"},
		})
	})
	t.Run("handle + symbol as separator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// A plus after a word in CJK contact is undecided by any affix, so plus reading spaces it as a separator, in a bundle plan and on a brand line alike
			{"Switch OLED+健身環+保護貼", "Switch OLED + 健身環 + 保護貼"},
			// Plus reading runs before the operator rules, so a CJK+A contact flips the line's later joiners too; a line with no contact keeps them
			{"陳上進+Vinta+Abc123", "陳上進 + Vinta + Abc123"},
			{"HiNet光世代+MOD+Wi-Fi全屋通", "HiNet 光世代 + MOD + Wi-Fi 全屋通"},
			{"套餐含MOD+Netflix+Disney", "套餐含 MOD+Netflix+Disney"},
			{"HiNet光世代+MOD+影劇館+/全選/自選20/特選餐/豪華餐(5選1)+Wi-Fi全屋通(1台)", "HiNet 光世代 + MOD + 影劇館+/全選/自選 20/特選餐/豪華餐 (5 選 1) + Wi-Fi 全屋通 (1 台)"},
			{"【速在必行方案】HiNet光世代+Wi-Fi全屋通1台+MOD影劇館+(300M/300M)", "【速在必行方案】HiNet 光世代 + Wi-Fi 全屋通 1 台 + MOD 影劇館+ (300M/300M)"},
			{"HiNet光世代+MOD+自選餐(全選)+「影劇館+」", "HiNet 光世代 + MOD + 自選餐 (全選) +「影劇館+」"},
			{"自選餐(全選)+「影劇館」", "自選餐 (全選) +「影劇館」"},
		})
	})
	// FIXME
	t.Run("handle + symbol as separator after a product name ending in a digit", func(t *testing.T) {
		testSpaceTextFails(t, []spaceTextCase{
			{"Switch 2+瑪利歐賽車世界同捆組", "Switch 2 + 瑪利歐賽車世界同捆組"},
		})
	})
	t.Run("handle + symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta+Abc123", "Vinta+Abc123"},
			{"前面A+B後面", "前面 A+B 後面"},
			{"得到一個A+B的結果", "得到一個 A+B 的結果"},
			{"答案是5+5的和", "答案是 5+5 的和"},
		})
	})
	t.Run("handle + symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"得到一個C++的結果", "得到一個 C++ 的結果"},
			{"得到一個 C++的結果", "得到一個 C++ 的結果"},
			{"得到一個i++的結果", "得到一個 i++ 的結果"},
			{"我會寫C++的程式", "我會寫 C++ 的程式"},
		})
	})
	t.Run("handle + symbol as affix", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// Grades
			{"得到一個A+的結果", "得到一個 A+ 的結果"},
			{"得到一個 A+ 的結果", "得到一個 A+ 的結果"},
			{"成績是A+的等級", "成績是 A+ 的等級"},
			// Sign before digits
			{"打+886這個號碼", "打 +886 這個號碼"},
			{"氣溫是+5度左右", "氣溫是 +5 度左右"},
			// Suffix after digits
			{"有100+的選擇", "有 100+ 的選擇"},
			{"這裡有18+的內容", "這裡有 18+ 的內容"},
			{"評分3.5+的餐廳", "評分 3.5+ 的餐廳"},
			{"Python 3+的版本", "Python 3+ 的版本"},
		})
	})
}

func TestSymbolQuestionMark(t *testing.T) {
	t.Run("handle ? symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面?", "前面?"},
			{"前面??", "前面??"},
			{"前面???", "前面???"},
			{"前面?後面", "前面? 後面"},
			{"前面??後面", "前面?? 後面"},
			{"前面???後面", "前面??? 後面"},
			{"前面?abc", "前面? abc"},
			{"前面?123", "前面? 123"},
			{"所以,請問Jackey的鼻子有幾個?3.14個", "所以, 請問 Jackey 的鼻子有幾個? 3.14 個"},
			// DO NOT change if already spacing
			{"前面 ? 後面", "前面 ? 後面"},
			{"前面? 後面", "前面? 後面"},
			// Rare cases, ignore
			// {"前面 ?後面", "前面 ?後面"},
		})
	})
}

func TestSymbolRoundBrackets(t *testing.T) {
	t.Run("handle ( ) symbols as round brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面(中文123漢字)後面", "前面 (中文 123 漢字) 後面"},
			{"前面(中文123)後面", "前面 (中文 123) 後面"},
			{"前面(123漢字)後面", "前面 (123 漢字) 後面"},
			{"前面(中文123) tail", "前面 (中文 123) tail"},
			{"head (中文123漢字)後面", "head (中文 123 漢字) 後面"},
			{"head (中文123漢字) tail", "head (中文 123 漢字) tail"},
			{"(or simply \"React\")", "(or simply \"React\")"},
			{"function(123)", "function(123)"},
			{"我看过的电影(1404)", "我看过的电影 (1404)"},
			{"預定於繳款截止日114/07/02(遇假日順延)之次一營業日進行扣款", "預定於繳款截止日 114/07/02 (遇假日順延) 之次一營業日進行扣款"},
			{"OperationalError: (2006, 'MySQL server has gone away')", "OperationalError: (2006, 'MySQL server has gone away')"},
			{"Chang Stream(变更记录流)是指collection(数据库集合)的变更事件流", "Chang Stream (变更记录流) 是指 collection (数据库集合) 的变更事件流"},
		})
	})
	t.Run("handle multiline content in round brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// A space before a newline is mid-content, not a bracket-edge space: only the literal string edges get stripped
			{"(x \n)中", "(x \n) 中"},
			{"(參數 \n)後面", "(參數 \n) 後面"},
		})
	})
}

func TestSymbolSemicolon(t *testing.T) {
	t.Run("handle ; symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面;後面", "前面; 後面"},
			// DO NOT change if already spacing
			{"前面 ; 後面", "前面 ; 後面"},
			{"前面; 後面", "前面; 後面"},
			// Rare cases, ignore
			// {"前面 ;後面", "前面 ;後面"},
		})
	})
}

func TestSymbolSingleQuotes(t *testing.T) {
	t.Run("handle ' ' symbols as quotes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Why are Python's 'private' methods not actually private?", "Why are Python's 'private' methods not actually private?"},
			{"举个栗子，如果一道题只包含'A' ~ 'Z'意味着字符集大小是", "举个栗子，如果一道题只包含 'A' ~ 'Z' 意味着字符集大小是"},
			{"后续会直接用iframe window.addEventListener('message')", "后续会直接用 iframe window.addEventListener('message')"},
			{"'! git commit -a -m \"蛤\"'", "'! git commit -a -m \"蛤\"'"},
			// Single quotes around Chinese text should not have spaces added
			{"Remove '铁蕾' from 1 Folder?", "Remove '铁蕾' from 1 Folder?"},
		})
	})
	t.Run("handle ' symbols as apostrophe", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"陳上進 likes 林依諾's status.", "陳上進 likes 林依諾's status."},
		})
	})
}

func TestSymbolSlash(t *testing.T) {
	t.Run("handle / symbol as separator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面/後面", "前面/後面"},
			{"Vinta/貓咪", "Vinta/貓咪"},
			{"貓咪/Vinta", "貓咪/Vinta"},
			{"速度是60公里/小時", "速度是 60 公里/小時"},
			{"價格是$100/每小時", "價格是 $100/每小時"},
			{"我/你\n他/她", "我/你\n他/她"},
			{"歡迎光臨/再見\n參考 https://example.com/docs", "歡迎光臨/再見\n參考 https://example.com/docs"},
			// DO NOT change if already spacing
			{"前面 / 後面", "前面 / 後面"},
			{"Vinta / Abc123", "Vinta / Abc123"},
			{"Abc123 / 陳上進", "Abc123 / 陳上進"},
			{"陳上進 / Abc123", "陳上進 / Abc123"},
			{"得到一個 A / B 的結果", "得到一個 A / B 的結果"},
			{"好人 / bad guy", "好人 / bad guy"},
			{"吃apple / banana", "吃 apple / banana"},
		})
	})
	t.Run("handle / symbol as joiner token", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Vinta/Abc123", "Vinta/Abc123"},
			{"得到一個A/B的結果", "得到一個 A/B 的結果"},
			{"他要做A/B測試", "他要做 A/B 測試"},
			{"打東東26/30", "打東東 26/30"},
			{"打東東1/denominator", "打東東 1/denominator"},
			{"吃apple/banana", "吃 apple/banana"},
			{"選A/B其中一個", "選 A/B 其中一個"},
			{"答案是6/2的商數", "答案是 6/2 的商數"},
			{"安装指令：npx skills add vinta/hal-9000", "安装指令：npx skills add vinta/hal-9000"},
		})
	})
	t.Run("handle / symbol as list", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"陳上進/貓咪/Abc123", "陳上進/貓咪/Abc123"},
			{"陳上進/Abc123/貓咪", "陳上進/Abc123/貓咪"},
			{"Abc123/Vinta/貓咪", "Abc123/Vinta/貓咪"},
			{"Abc123/陳上進/貓咪", "Abc123/陳上進/貓咪"},
			{"日期是2024/01/22的早上", "日期是 2024/01/22 的早上"},
			{"8964/3★集會所接待員/克隆·麻煩大師/手卷師傅（已退休）/主程式毀滅者/dae-dae-o/#絕地家庭小會議/#今天大掃除了沒有/NS編號在banner裡/discord:史單力#3230", "8964/3★集會所接待員/克隆・麻煩大師/手卷師傅（已退休）/主程式毀滅者/dae-dae-o/#絕地家庭小會議/#今天大掃除了沒有/NS 編號在 banner 裡/discord: 史單力 #3230"},
			{"8964/3★集會所接待員/克隆·麻煩大師/手卷師傅(已退休)/主程式毀滅者/dae-dae-o/#絕地家庭小會議/#今天大掃除了沒有/NS編號在banner裡/discord:史單力#3230", "8964/3★集會所接待員/克隆・麻煩大師/手卷師傅 (已退休)/主程式毀滅者/dae-dae-o/#絕地家庭小會議/#今天大掃除了沒有/NS 編號在 banner 裡/discord: 史單力 #3230"},
			{"after 80'/气象工作者/不苟同/关注abc天气变化/向往123自由/热爱科学、互联网、编程Node.js Web C++ Julia Python", "after 80'/气象工作者/不苟同/关注 abc 天气变化/向往 123 自由/热爱科学、互联网、编程 Node.js Web C++ Julia Python"},
			// DO NOT change if already spacing
			{"陳上進 / 貓咪 / Abc123", "陳上進 / 貓咪 / Abc123"},
			{"陳上進 / Abc123 / 貓咪", "陳上進 / Abc123 / 貓咪"},
			{"Abc123 / Vinta / 貓咪", "Abc123 / Vinta / 貓咪"},
			{"Abc123 / 陳上進 / 貓咪", "Abc123 / 陳上進 / 貓咪"},
			{"2016-12-26(奇幻电影节) / 2017-01-20(美国) / 詹姆斯麦卡沃伊", "2016-12-26 (奇幻电影节) / 2017-01-20 (美国) / 詹姆斯麦卡沃伊"},
		})
	})
	t.Run("handle / symbol as Unix absolute file path", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"/home和/root是Linux中的頂級目錄", "/home 和 /root 是 Linux 中的頂級目錄"},
			{"/home/與/root是Linux中的頂級目錄", "/home/ 與 /root 是 Linux 中的頂級目錄"},
			{"\"/home/\"和\"/root\"是Linux中的頂級目錄", "\"/home/\" 和 \"/root\" 是 Linux 中的頂級目錄"},
			{"當你用cat和od指令查看/dev/random和/dev/urandom的內容時", "當你用 cat 和 od 指令查看 /dev/random 和 /dev/urandom 的內容時"},
			{"當你用cat和od指令查看\"/dev/random\"和\"/dev/urandom\"的內容時", "當你用 cat 和 od 指令查看 \"/dev/random\" 和 \"/dev/urandom\" 的內容時"},
			// Basic Unix paths
			{"在/home目錄", "在 /home 目錄"},
			{"查看/etc/passwd文件", "查看 /etc/passwd 文件"},
			{"進入/usr/local/bin目錄", "進入 /usr/local/bin 目錄"},
			// Paths with dots
			{"配置檔在/etc/nginx/nginx.conf", "配置檔在 /etc/nginx/nginx.conf"},
			{"隱藏檔案/.bashrc很重要", "隱藏檔案 /.bashrc 很重要"},
			{"查看/home/.config/settings", "查看 /home/.config/settings"},
			// Paths with version numbers
			{"安裝到/usr/lib/python3.9/", "安裝到 /usr/lib/python3.9/"},
			{"位於/opt/node-v16.14.0/bin", "位於 /opt/node-v16.14.0/bin"},
			// Paths with special characters
			{"備份到/mnt/backup.2024-01-01/", "備份到 /mnt/backup.2024-01-01/"},
			{"日誌在/var/log/app-name.log", "日誌在 /var/log/app-name.log"},
			// Paths with @ symbols (npm packages)
			{"模組在/node_modules/@babel/core", "模組在 /node_modules/@babel/core"},
			{"套件在/node_modules/@types/node", "套件在 /node_modules/@types/node"},
			// Paths with + symbols
			{"編譯器在/usr/bin/g++", "編譯器在 /usr/bin/g++"},
			{"套件在/usr/lib/gcc/x86_64-linux-gnu/11++", "套件在 /usr/lib/gcc/x86_64-linux-gnu/11++"},
			// Paths ending with slash before CJK
			{"目錄/usr/bin/包含執行檔", "目錄 /usr/bin/ 包含執行檔"},
			{"資料夾/etc/nginx/存放設定", "資料夾 /etc/nginx/ 存放設定"},
		})
	})
	t.Run("handle / symbol as Unix relative file path", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// Basic relative paths
			{"檢查src/main.py文件", "檢查 src/main.py 文件"},
			{"構建dist/index.js完成", "構建 dist/index.js 完成"},
			{"運行test/spec.js測試", "運行 test/spec.js 測試"},
			{"編輯docs/README.md文檔", "編輯 docs/README.md 文檔"},
			// Project directories
			{"查看templates/base.html模板", "查看 templates/base.html 模板"},
			{"複製assets/images/logo.png圖片", "複製 assets/images/logo.png 圖片"},
			{"配置config/database.yml設定", "配置 config/database.yml 設定"},
			{"執行scripts/deploy.sh腳本", "執行 scripts/deploy.sh 腳本"},
			// Build/output directories
			{"清理build/temp/目錄", "清理 build/temp/ 目錄"},
			{"輸出到target/release/資料夾", "輸出到 target/release/ 資料夾"},
			{"發布到public/static/路徑", "發布到 public/static/ 路徑"},
			// Development directories
			{"安裝node_modules/@babel/core套件", "安裝 node_modules/@babel/core 套件"},
			{"設定.git/hooks/pre-commit鉤子", "設定 .git/hooks/pre-commit 鉤子"},
			{"編輯.vscode/settings.json配置", "編輯 .vscode/settings.json 配置"},
			// With leading ./
			{"參考./docs/API.md文件", "參考 ./docs/API.md 文件"},
			{"執行./scripts/test.sh腳本", "執行 ./scripts/test.sh 腳本"},
			{"查看./.claude/CLAUDE.md說明", "查看 ./.claude/CLAUDE.md 說明"},
			// Nested paths
			{"位於src/components/Button/index.tsx", "位於 src/components/Button/index.tsx"},
			{"存放在assets/fonts/Inter/Regular.woff2", "存放在 assets/fonts/Inter/Regular.woff2"},
			// Multiple file paths in one sentence
			{"從src/utils.js複製到dist/utils.js", "從 src/utils.js 複製到 dist/utils.js"},
			{"比較test/fixtures/input.txt和test/fixtures/output.txt", "比較 test/fixtures/input.txt 和 test/fixtures/output.txt"},
		})
	})
	t.Run("handle / symbol as glob pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"聽說桐島rm -rf /*了", "聽說桐島 rm -rf /* 了"},
			{"模板在templates/*.html裡", "模板在 templates/*.html 裡"},
			{"測試所有test/**/*.js檔案", "測試所有 test/**/*.js 檔案"},
		})
	})
}

func TestSymbolSquareBrackets(t *testing.T) {
	t.Run("handle [ ] symbols as square brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面[中文123漢字]後面", "前面 [中文 123 漢字] 後面"},
			{"前面[中文123]後面", "前面 [中文 123] 後面"},
			{"前面[123漢字]後面", "前面 [123 漢字] 後面"},
			{"前面[中文123] tail", "前面 [中文 123] tail"},
			{"head [中文123漢字]後面", "head [中文 123 漢字] 後面"},
			{"head [中文123漢字] tail", "head [中文 123 漢字] tail"},
		})
	})
	t.Run("handle multiline content in square brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// A space before a newline is mid-content, not a bracket-edge space: only the literal string edges get stripped
			{"[x \n]中", "[x \n] 中"},
			{"中[ 多行\n內容 ]", "中 [多行\n內容]"},
		})
	})
}

func TestSymbolSuperscript(t *testing.T) {
	t.Run("Symbol Superscripts", func(t *testing.T) {
		t.Run("handle superscript as suffix", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"前面E=mc²後面", "前面 E=mc² 後面"},
				{"115年賽事加碼：7/21-10/8 新申請MOD+自選餐(全選)/影劇館⁺加碼", "115 年賽事加碼：7/21-10/8 新申請 MOD + 自選餐 (全選)/影劇館⁺ 加碼"},
				{"甲⁰乙、甲¹乙、甲²乙、甲³乙、甲⁴乙、甲⁵乙、甲⁶乙、甲⁷乙、甲⁸乙、甲⁹乙", "甲⁰ 乙、甲¹ 乙、甲² 乙、甲³ 乙、甲⁴ 乙、甲⁵ 乙、甲⁶ 乙、甲⁷ 乙、甲⁸ 乙、甲⁹ 乙"},
				{"甲ⁱ乙、甲ⁿ乙、甲⁺乙、甲⁻乙、甲⁼乙", "甲ⁱ 乙、甲ⁿ 乙、甲⁺ 乙、甲⁻ 乙、甲⁼ 乙"},
				{"甲⁽註⁾乙", "甲⁽註⁾ 乙"},
			})
		})
		// \u2122
		t.Run("handle ™ symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"Trademark™後面", "Trademark™ 後面"},
				{"商標™後面", "商標™ 後面"},
			})
		})
		// \u2120
		t.Run("handle ℠ symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"Service Mark℠後面", "Service Mark℠ 後面"},
				{"服務商標℠後面", "服務商標℠ 後面"},
			})
		})
	})
	// \u00ae
	t.Run("Symbol ®", func(t *testing.T) {
		t.Run("handle ® symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"Registered Trademark®後面", "Registered Trademark® 後面"},
				{"註冊商標®公司", "註冊商標® 公司"},
				{"註冊商標®與Trademark™", "註冊商標® 與 Trademark™"},
			})
		})
	})
	// \u00a9
	t.Run("Symbol ©", func(t *testing.T) {
		t.Run("handle © symbol", func(t *testing.T) {
			testSpaceText(t, []spaceTextCase{
				{"版權所有©2026東亞重工", "版權所有 © 2026 東亞重工"},
				{"版權所有©2012-2026東亞重工", "版權所有 © 2012-2026 東亞重工"},
				{"Copyright © 2026東亞重工", "Copyright © 2026 東亞重工"},
				{"Copyright © 2012-2026東亞重工", "Copyright © 2012-2026 東亞重工"},
			})
		})
	})
}

func TestSymbolTilde(t *testing.T) {
	t.Run("handle ~ symbol", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面~", "前面~"},
			{"前面~~", "前面~~"},
			{"前面~~~", "前面~~~"},
			{"前面~後面", "前面~ 後面"},
			{"前面~~後面", "前面~~ 後面"},
			{"前面~~~後面", "前面~~~ 後面"},
			{"前面~abc", "前面~ abc"},
			{"前面~123", "前面~ 123"},
			// DO NOT change if already spacing
			{"前面 ~ 後面", "前面 ~ 後面"},
			{"前面~ 後面", "前面~ 後面"},
			{"前面 ~後面", "前面 ~後面"},
		})
	})
	t.Run("handle ~ symbol as preserved pattern", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面~=後面", "前面 ~= 後面"},
			{"前面 ~= 後面", "前面 ~= 後面"},
		})
	})
}

func TestSymbolUnderscore(t *testing.T) {
	t.Run("handle _ symbol as separator", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"前面_後面", "前面_後面"},
			{"Vinta_Abc123", "Vinta_Abc123"},
			{"Vinta_Abc123_Kitten", "Vinta_Abc123_Kitten"},
			{"Vinta_貓咪", "Vinta_貓咪"},
			{"貓咪_Vinta", "貓咪_Vinta"},
			{"陳上進_貓咪_Abc123", "陳上進_貓咪_Abc123"},
			{"陳上進_Abc123_貓咪", "陳上進_Abc123_貓咪"},
			{"Abc123_Vinta_貓咪", "Abc123_Vinta_貓咪"},
			{"Abc123_陳上進_貓咪", "Abc123_陳上進_貓咪"},
			{"得到一個A_B的結果", "得到一個 A_B 的結果"},
			{"為什麼你們就是不能加個空格呢？_20771210_最終版_v365.7.24.zip", "為什麼你們就是不能加個空格呢？_20771210_最終版_v365.7.24.zip"},
			// Rare cases, ignore
			// {"前面 _ 後面", "前面 _ 後面"},
			// {"Vinta _ Abc123", "Vinta _ Abc123"},
			// {"Vinta _ Abc123 _ Kitten", "Vinta _ Abc123 _ Kitten"},
			// {"陳上進 _ 貓咪 _ Abc123", "陳上進 _ 貓咪 _ Abc123"},
			// {"陳上進 _ Abc123 _ 貓咪", "陳上進 _ Abc123 _ 貓咪"},
			// {"Abc123 _ Vinta _ 貓咪", "Abc123 _ Vinta _ 貓咪"},
			// {"Abc123 _ 陳上進 _ 貓咪", "Abc123 _ 陳上進 _ 貓咪"},
		})
	})
}
