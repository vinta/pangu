package pangu_test

import "testing"

func TestCJKAlphabetsNumbers(t *testing.T) {
	t.Run("handle short text", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中a", "中 a"},
			{"a中", "a 中"},
			{"1中", "1 中"},
			{"中1", "中 1"},
			{"中a1", "中 a1"},
			{"a1中", "a1 中"},
			{"a中1", "a 中 1"},
			{"1中a", "1 中 a"},
		})
	})
	t.Run("handle alphabets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中文abc", "中文 abc"},
			{"abc中文", "abc 中文"},
		})
	})
	t.Run("handle numbers", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中文123", "中文 123"},
			{"123中文", "123 中文"},
		})
	})
	// https://symbl.cc/en/unicode-table/#latin-1-supplement
	t.Run("handle Latin-1 Supplement", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中文Ø漢字", "中文 Ø 漢字"},
			{"中文 Ø 漢字", "中文 Ø 漢字"},
		})
	})
	// https://symbl.cc/en/unicode-table/#greek-coptic
	t.Run("handle Greek and Coptic", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中文β漢字", "中文 β 漢字"},
			{"中文 β 漢字", "中文 β 漢字"},
			{"我是α，我是Ω", "我是 α，我是 Ω"},
		})
	})
	// https://symbl.cc/en/unicode-table/#number-forms
	t.Run("handle Number Forms", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"中文Ⅶ漢字", "中文 Ⅶ 漢字"},
			{"中文 Ⅶ 漢字", "中文 Ⅶ 漢字"},
		})
	})
	// https://symbl.cc/en/unicode-table/#letterlike-symbols
	t.Run("handle Letterlike Symbols", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"今天123℃很熱", "今天 123℃ 很熱"},
			{"攝氏25℃到30℃之間", "攝氏 25℃ 到 30℃ 之間"},
			{"水溫98℉了", "水溫 98℉ 了"},
			{"5℃~10℃之間", "5℃~10℃ 之間"},
			{"溫度是℃單位", "溫度是 ℃ 單位"},
			{"第№5號", "第 №5 號"},
			{"電阻10Ω很小", "電阻 10Ω 很小"},
			{"中文ℝ漢字", "中文 ℝ 漢字"},
			{"中文 ℝ 漢字", "中文 ℝ 漢字"},
			{"符號ℓ表示長度", "符號 ℓ 表示長度"},
			{"資訊ℹ圖示", "資訊 ℹ 圖示"},
			{"估計℮500ml", "估計 ℮500ml"},
		})
	})
	// https://symbl.cc/en/unicode-table/#dingbats
	t.Run("handle Dingbats symbols", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"剪刀✂符號", "剪刀 ✂ 符號"},
			{"完成✅了", "完成 ✅ 了"},
			{"愛心❤符號", "愛心 ❤ 符號"},
		})
	})
	// https://symbl.cc/en/unicode-table/#cjk-radicals-supplement
	t.Run("handle CJK Radicals Supplement", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abc⻤123", "abc ⻤ 123"},
			{"abc ⻤ 123", "abc ⻤ 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#kangxi-radicals
	t.Run("handle Kangxi Radicals", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abc⾗123", "abc ⾗ 123"},
			{"abc ⾗ 123", "abc ⾗ 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#hiragana
	t.Run("handle Hiragana", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abcあ123", "abc あ 123"},
			{"abc あ 123", "abc あ 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#katakana
	t.Run("handle Katakana", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abcア123", "abc ア 123"},
			{"abc ア 123", "abc ア 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#bopomofo
	t.Run("handle Bopomofo", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abcㄅ123", "abc ㄅ 123"},
			{"abc ㄅ 123", "abc ㄅ 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#enclosed-cjk-letters-and-months
	t.Run("handle Enclosed CJK Letters And Months", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abc㈱123", "abc ㈱ 123"},
			{"abc ㈱ 123", "abc ㈱ 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#cjk-unified-ideographs-extension-a
	t.Run("handle CJK Unified Ideographs Extension-A", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abc㐂123", "abc 㐂 123"},
			{"abc 㐂 123", "abc 㐂 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#cjk-unified-ideographs
	t.Run("handle CJK Unified Ideographs", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abc丁123", "abc 丁 123"},
			{"abc 丁 123", "abc 丁 123"},
		})
	})
	// https://symbl.cc/en/unicode-table/#cjk-compatibility-ideographs
	t.Run("handle CJK Compatibility Ideographs", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"abc車123", "abc 車 123"},
			{"abc 車 123", "abc 車 123"},
		})
	})
}

func TestNameSuffix(t *testing.T) {
	t.Run("handle product names with suffixes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"Apple Fitness+推出新課程", "Apple Fitness+ 推出新課程"},
			{"Apple TV+上架了新片", "Apple TV+ 上架了新片"},
			{"Discovery+和discovery+都上架了", "Discovery+ 和 discovery+ 都上架了"},
			{"Disney+上架了新片", "Disney+ 上架了新片"},
			{"Disney+上架了C++課程", "Disney+ 上架了 C++ 課程"},
			{"mo店+免運無限次 天天超取290起-momo購物網", "mo 店+ 免運無限次 天天超取 290 起 - momo 購物網"},
			{"mo 店+ 免運無限次", "mo 店+ 免運無限次"},
			{"PS+會員", "PS+ 會員"},
			{"公視+上架了新片", "公視+ 上架了新片"},
			{"如何使用PTS+（公視+）註冊與觀看？", "如何使用 PTS+（公視+）註冊與觀看？"},
			{"公視+(免費平台)", "公視+ (免費平台)"},
			{"今天來看公視+", "今天來看公視+"},
			{"MOD影劇館+上架了新片", "MOD 影劇館+ 上架了新片"},
			{"Netflix、Disney+、Apple TV+、MOD影劇館+、公視+等串流平台", "Netflix、Disney+、Apple TV+、MOD 影劇館+、公視+ 等串流平台"},
		})
	})
	t.Run("handle product tiers with suffixes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"vivo X70 Pro+開賣", "vivo X70 Pro+ 開賣"},
		})
	})
	t.Run("handle credit ratings with suffixes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"評等介於AA-和AA+之間", "評等介於 AA- 和 AA+ 之間"},
			{"惠譽給予BBB+評等", "惠譽給予 BBB+ 評等"},
			{"惠譽給予BBB-評等", "惠譽給予 BBB- 評等"},
			{"中華信評給予twAA+評等", "中華信評給予 twAA+ 評等"},
		})
	})
	t.Run("handle blood types with suffixes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"血型是AB+的人", "血型是 AB+ 的人"},
			{"血型是AB-的人", "血型是 AB- 的人"},
			{"血型是Rh+的人", "血型是 Rh+ 的人"},
			{"血型是Rh-的人", "血型是 Rh- 的人"},
			{"型號AB-123的零件，血型是AB-的人", "型號 AB-123 的零件，血型是 AB- 的人"},
		})
	})
	t.Run("handle closing punctuation tight after a suffixes", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"公視+，今天有新片", "公視+，今天有新片"},
			{"公視+。", "公視+。"},
			{"(公視+）", "(公視+）"},
			{"[公視+]", "[公視+]"},
			{"「公視+」", "「公視+」"},
		})
	})
	t.Run("handle non-preserved names", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"私視+上線", "私視 + 上線"},
			{"Disney-上架了新片", "Disney - 上架了新片"},
		})
	})
}

func TestURL(t *testing.T) {
	t.Run("leave the inside of a URL untouched", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			// Issue https://github.com/vinta/pangu.js/issues/147
			{"第三條的內容為http://se.360.cn/", "第三條的內容為 http://se.360.cn/"},
			// Issue https://github.com/vinta/pangu.js/issues/149
			{"你https://%E5%A6%82", "你 https://%E5%A6%82"},
			// Issue https://github.com/vinta/pangu.js/issues/155
			{"https://xxxxx/自动加空格.html", "https://xxxxx/自动加空格.html"},
			{"打開此連結，https://www.google.com/search?q=%E5%9B%BD%E5%AF%86SM2%2F3%2F4%E7%AE%97%E6%B3%95+360", "打開此連結，https://www.google.com/search?q=%E5%9B%BD%E5%AF%86SM2%2F3%2F4%E7%AE%97%E6%B3%95+360"},
			{"https://www.google.com/search?q=中文&hl=zh-TW", "https://www.google.com/search?q=中文&hl=zh-TW"},
			{"https://zh.wikipedia.org/w/index.php?title=中文&action=history", "https://zh.wikipedia.org/w/index.php?title=中文&action=history"},
			{"網址是https://zh.wikipedia.org/wiki/%E4%B8%AD%E6%96%87", "網址是 https://zh.wikipedia.org/wiki/%E4%B8%AD%E6%96%87"},
			{"https://zh.wikipedia.org/wiki/中文#歷史", "https://zh.wikipedia.org/wiki/中文#歷史"},
			{"參考https://zh.wikipedia.org/wiki/中文#歷史的說明", "參考 https://zh.wikipedia.org/wiki/中文#歷史的說明"},
			{"https://zh.wikipedia.org/wiki/盤古", "https://zh.wikipedia.org/wiki/盤古"},
		})
	})
	// FIXME: CJK characters continue the URL, and no rule tells URL-internal CJK from prose written tight after the URL. See ADR 0026
	t.Run("space a URL from CJK on both sides", func(t *testing.T) {
		testSpaceTextFails(t, []spaceTextCase{
			{"搜尋https://www.google.com/search?q=pangu.js&hl=zh-TW看看", "搜尋 https://www.google.com/search?q=pangu.js&hl=zh-TW 看看"},
			{"看https://github.com/vinta/pangu.js/issues/155這個issue", "看 https://github.com/vinta/pangu.js/issues/155 這個 issue"},
			{"文件在https://developer.mozilla.org/zh-TW/docs/Web/API/URL/canParse_static這裡", "文件在 https://developer.mozilla.org/zh-TW/docs/Web/API/URL/canParse_static 這裡"},
		})
	})
	t.Run("stop a URL at CJK punctuation, quotes, and brackets", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"請看https://vinta.ws/code/。", "請看 https://vinta.ws/code/。"},
			{"請看https://vinta.ws/code/，謝謝", "請看 https://vinta.ws/code/，謝謝"},
			{"（https://vinta.ws/code/）", "（https://vinta.ws/code/）"},
			{"(https://vinta.ws/code/)", "(https://vinta.ws/code/)"},
			{"「https://vinta.ws/code/」", "「https://vinta.ws/code/」"},
		})
	})
	t.Run("leave trailing half-width punctuation outside the URL", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"詳見https://vinta.ws/code/.", "詳見 https://vinta.ws/code/."},
		})
	})
	t.Run("leave a URL inside an attribute value untouched", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"<a href=\"https://zh.wikipedia.org/wiki/中文#歷史\">中文</a>", "<a href=\"https://zh.wikipedia.org/wiki/中文#歷史\">中文</a>"},
			{"<a href=\"http://vinta.ws/中文網址with英文.html\">oh一個超連結with英文，網址包含中文</a>", "<a href=\"http://vinta.ws/中文網址with英文.html\">oh 一個超連結 with 英文，網址包含中文</a>"},
		})
	})
	t.Run("space a hashtag on a line that also holds a URL", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"看完這篇#pangu 的介紹 https://vinta.ws/code/", "看完這篇 #pangu 的介紹 https://vinta.ws/code/"},
		})
	})
	t.Run("leave a URL with no CJK contact untouched", func(t *testing.T) {
		testSpaceText(t, []spaceTextCase{
			{"see https://vinta.ws/code/ and 中文", "see https://vinta.ws/code/ and 中文"},
		})
	})
}
