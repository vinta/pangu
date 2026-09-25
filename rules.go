package pangu

import "regexp"

// Character classes, written as bracket-expression bodies. A is A-Za-z, N is 0-9, S varies per rule.
const (
	cjk = `\x{2e80}-\x{2eff}` + // CJK Radicals Supplement
		`\x{2f00}-\x{2fdf}` + // Kangxi Radicals
		`\x{3040}-\x{309f}` + // Hiragana
		`\x{30a0}-\x{30fa}\x{30fc}-\x{30ff}` + // Katakana without \x{30fb}, the character middleDot converts to
		`\x{3100}-\x{312f}` + // Bopomofo
		`\x{3200}-\x{32ff}` + // Enclosed CJK Letters and Months
		`\x{3400}-\x{4dbf}` + // CJK Unified Ideographs Extension A
		`\x{4e00}-\x{9fff}` + // CJK Unified Ideographs
		`\x{f900}-\x{faff}` // CJK Compatibility Ideographs

	alpha = `A-Za-z`
	alnum = `A-Za-z0-9`

	// Latin-1 starts one past NBSP (\x{a0}) so an NBSP is in no class and no rule rewrites it. See ADR 0009
	ansExtended = `\x{0370}-\x{03ff}` + // Greek and Coptic
		`0-9`
	ansExtendedTail = `\x{00a1}-\x{00ff}` + // Latin-1 Supplement after NBSP
		`\x{2150}-\x{218f}` + // Number Forms
		`\x{2700}-\x{27bf}` + // Dingbats
		`\x{2100}-\x{214f}` // Letterlike Symbols

	// Superscript suffixes stay attached on the left
	superscriptSuffixes = `\x{00ae}\x{00b2}\x{00b3}\x{00b9}\x{2070}\x{2071}\x{2074}-\x{207c}\x{207e}\x{207f}\x{2120}\x{2122}`

	ansCJKAfter  = alpha + ansExtended + `@\$%\^&\*\-\+\\=` + ansExtendedTail
	ansBeforeCJK = alpha + ansExtended + `\$%\^&\*\-\+\\=` + ansExtendedTail + superscriptSuffixes

	// No + because affixes and plus reading decide every plus before the operator rules run
	operators = `\*=&\-`

	quotes = "`" + `"\x{05f4}`

	leftBracketsBasic     = `\(\[\{`
	rightBracketsBasic    = `\)\]\}`
	leftBracketsExtended  = `\(\[\{<>\x{201c}`
	rightBracketsExtended = `\)\]\}<>\x{201d}`

	// JavaScript's \s. Go's \s is ASCII only
	jsSpace = `\t\n\v\f\r \x{00a0}\x{1680}\x{2000}-\x{200a}\x{2028}\x{2029}\x{202f}\x{205f}\x{3000}\x{feff}`

	// Full-width punctuation that keeps an adjacent plus tight
	fullWidthPunctuation = `\x{ff0c}\x{3002}\x{ff1b}\x{ff1a}\x{ff01}\x{ff1f}\x{3001}\x{ff08}\x{ff09}\x{300c}\x{300d}\x{300e}\x{300f}\x{3010}\x{3011}\x{300a}\x{300b}`
)

const (
	filePathDirs     = `home|root|usr|etc|var|opt|tmp|dev|mnt|proc|sys|bin|boot|lib|media|run|sbin|srv|node_modules|path|project|src|dist|test|tests|docs|templates|assets|public|static|config|scripts|tools|build|out|target|your|\.claude|\.git|\.vscode`
	filePathChars    = `[A-Za-z0-9_\-\.@\+\*]+`
	unixAbsolutePath = `/(?:\.?(?:` + filePathDirs + `)|\.(?:[A-Za-z0-9_\-]+))(?:/` + filePathChars + `)*`
	unixRelativePath = `(?:\./)?(?:` + filePathDirs + `)(?:/` + filePathChars + `)+`
	windowsPath      = `[A-Z]:\\(?:[A-Za-z0-9_\-\. ]+\\?)+`
)

// Name suffixes: product names and tiers take + only, credit ratings and blood types take + or -
const (
	productName      = `Apple TV|CATCHPLAY|[Dd]iscovery|Disney|ESPN|Fitness|iCloud|mo ?店|Paramount|PS`
	productNameInCJK = `公視|影劇館`
	productTier      = `Pro`
	creditRating     = `(?:tw)?(?:AA|BBB|BB|CCC)|tw[AB]`
	bloodType        = `AB|RhD|Rh`
)

// Placeholder delimiters come from the Private Use Area so they do not collide with ordinary text
const (
	htmlTagMentionStart = "\ue004"
	htmlTagMentionEnd   = "\ue005"
	httpURLStart        = "\ue00a"
)

var (
	anyCJK = regexp.MustCompile(`[` + cjk + `]`)

	isCJK                = charClass(cjk)
	isCJKOrAlnum         = charClass(cjk + alnum)
	isRightBracket       = charClass(rightBracketsBasic)
	isLeftBracket        = charClass(leftBracketsBasic)
	isSuperscript        = charClass(superscriptSuffixes)
	isSpaceOrPipe        = charClass(jsSpace + `|`)
	isPlusNeighbor       = charClass(jsSpace + `+` + fullWidthPunctuation)
	isFullWidthLeft      = charClass(`\x{ff08}\x{300c}\x{300e}\x{3010}\x{300a}`)
	isMiddleDotGap       = charClass(` \x{00a0}\x{00b7}\x{2022}\x{2027}`)
	isClosingAfterSuffix = charClass(`/)\]}\x{ff09}\x{3011}\x{3015}\x{3009}\x{300b}\x{300d}\x{300f}\x{ff0c}\x{3002}\x{3001}\x{ff1b}\x{ff1a}\x{ff01}\x{ff1f}`)

	// A punctuation run after CJK gets a trailing space when CJK, a letter, or a digit follows
	cjkPunctuation = regexp.MustCompile(`([` + cjk + `])([!;,\?:]+)`)
	// A punctuation run directly before CJK gets a space after it, whatever sits on its left. See ADR 0007
	punctuationCJK = regexp.MustCompile(`[!;,\?]+`)

	// Tilde has its own rule so ~= stays intact
	cjkTilde       = regexp.MustCompile(`([` + cjk + `])(~+)`)
	cjkTildeEquals = regexp.MustCompile(`([` + cjk + `])(~=)`)

	// A letter or digit after the period reads as a file extension, so only CJK after it gets a space
	cjkPeriod   = regexp.MustCompile(`([` + cjk + `])(\.)`)
	anPeriodCJK = regexp.MustCompile(`([` + alnum + `])(\.)([` + cjk + `])`)

	anColonCJK = regexp.MustCompile(`([` + alnum + `])(:)([` + cjk + `])`)
	// The only full-width conversion: a colon after CJK, directly before a parenthesis
	fixCJKColonANS = regexp.MustCompile(`([` + cjk + `]):([A-Z0-9\(\)])`)

	dotsCJK = regexp.MustCompile(`(\.{2,}|\x{2026})([` + cjk + `])`)

	// The quote class excludes ' because single quotes have their own rules
	cjkQuote = regexp.MustCompile(`([` + cjk + `])([` + quotes + `])`)
	quoteCJK = regexp.MustCompile(`([` + quotes + `])([` + cjk + `])`)
	// [\s\S] so a quoted segment that spans a line break still pairs with its own closing quote
	fixQuoteAnyQuote = regexp.MustCompile(`([` + quotes + `]+)[ ]*([\s\S]+?)[ ]*([` + quotes + `]+)`)
	quoteAN          = regexp.MustCompile(`(\x{201d})([` + alnum + `])`)
	cjkQuoteAN       = regexp.MustCompile(`([` + cjk + `])(")([` + alnum + `])`)

	cjkSingleQuoteButPossessive = regexp.MustCompile(`([` + cjk + `])('[^s])`)
	singleQuoteCJK              = regexp.MustCompile(`(')([` + cjk + `])`)
	fixPossessiveSingleQuote    = regexp.MustCompile(`([` + alnum + cjk + `])( )('s)`)
	singleQuotePureCJK          = regexp.MustCompile(`'[` + cjk + `]+'`)

	hashCJKHash = regexp.MustCompile(`([` + cjk + `])(#)([` + cjk + `]+)(#)([` + cjk + `])`)
	// The hashtag guard rejects an NBSP the same way it rejects a space
	cjkHash = regexp.MustCompile(`([` + cjk + `])(#[^ \x{00a0}])`)
	// A hashtag right after a slash in a list (/#tag) is a hashtag, not a C# shape
	hashCJK = regexp.MustCompile(`([^ \x{00a0}/]#)([` + cjk + `])`)

	// Matches when the text ends with a name suffix. (?:^|[^A-Za-z0-9]) stands in for upstream's (?<![A-Za-z0-9])
	nameSuffixAtEnd = regexp.MustCompile(`(?:^|[^A-Za-z0-9])(?:(?:` + productName + `|` + productTier + `)\+|(?:` + creditRating + `|` + bloodType + `)[+-])$|(?:` + productNameInCJK + `)\+$`)

	// Only direct CJK contact makes - * = & an operator; between two half-width characters it is a joiner token
	cjkOperatorANS = regexp.MustCompile(`([` + cjk + `])([` + operators + `])([` + alnum + leftBracketsBasic + `])`)
	ansOperatorCJK = regexp.MustCompile(`([` + alnum + rightBracketsBasic + `])([` + operators + `])([` + cjk + `])`)

	// Hyphen, pipe, and plus readings are decided per line by direct CJK contact
	hyphenCJKContact = regexp.MustCompile(`[` + cjk + `]-|-[` + cjk + `]`)
	pipeCJKContact   = regexp.MustCompile(`[` + cjk + `]\||\|[` + cjk + `]`)
	pipeSeparator    = regexp.MustCompile(`([^` + jsSpace + `|])[ ]*(\|+)[ ]*`)
	plusCJKContact   = regexp.MustCompile(`[` + cjk + `]\+|\+[` + cjk + `]`)
	hyphen           = regexp.MustCompile(`-`)
	plus             = regexp.MustCompile(`\+`)

	// Single-letter grades (A+, B-, C*) before CJK get the space after the symbol
	singleLetterGradeCJK = regexp.MustCompile(`\b([` + alpha + `])([\+\-\*])([` + cjk + `])`)

	// Affixes attach a symbol to its half-width side: sign +886, flag -m, suffix 100+. See ADR 0015 and 0024
	cjkSignDigit  = regexp.MustCompile(`([` + cjk + `])(\+)([0-9])`)
	cjkHyphenFlag = regexp.MustCompile(`([` + cjk + `])(-)([a-z])\b`)
	digitPlusCJK  = regexp.MustCompile(`\b([0-9]+)(\+)([` + cjk + `])`)

	cjkLessThan    = regexp.MustCompile(`([` + cjk + `])(<)([` + alnum + `])`)
	lessThanCJK    = regexp.MustCompile(`([` + alnum + `])(<)([` + cjk + `])`)
	cjkGreaterThan = regexp.MustCompile(`([` + cjk + `])(>)([` + alnum + `])`)
	greaterThanCJK = regexp.MustCompile(`([` + alnum + `])(>)([` + cjk + `])`)

	cjkLeftBracket  = regexp.MustCompile(`([` + cjk + `])([` + leftBracketsExtended + `])`)
	rightBracketCJK = regexp.MustCompile(`([` + rightBracketsExtended + `])([` + cjk + `])`)

	ansCJKLeftQuoteAnyRightQuote = regexp.MustCompile(`([` + alnum + cjk + `])[ ]*(\x{201c})([` + alnum + cjk + `\-_ ]+)(\x{201d})`)
	leftQuoteAnyRightQuoteANSCJK = regexp.MustCompile(`(\x{201c})([` + alnum + cjk + `\-_ ]+)(\x{201d})[ ]*([` + alnum + cjk + `])`)
	// Some input habits type both quotes of a pair as closing curly quotes
	ansCJKRightQuoteAnyRightQuote = regexp.MustCompile(`([` + alnum + cjk + `])[ ]*(\x{201d})[ ]*([` + alnum + cjk + `\-_ ]+?)[ ]*(\x{201d})`)

	anLeftBracket  = regexp.MustCompile(`([` + alnum + `])([` + leftBracketsBasic + `])`)
	rightBracketAN = regexp.MustCompile(`([` + rightBracketsBasic + `])([` + alnum + `])`)

	cjkUnixAbsolutePath      = regexp.MustCompile(`([` + cjk + `])(` + unixAbsolutePath + `)`)
	cjkUnixRelativePath      = regexp.MustCompile(`([` + cjk + `])(` + unixRelativePath + `)`)
	cjkWindowsPath           = regexp.MustCompile(`([` + cjk + `])(` + windowsPath + `)`)
	unixAbsolutePathSlashCJK = regexp.MustCompile(`(` + unixAbsolutePath + `/)([` + cjk + `])`)
	unixRelativePathSlashCJK = regexp.MustCompile(`(` + unixRelativePath + `/)([` + cjk + `])`)

	cjkANS = regexp.MustCompile(`([` + cjk + `])([` + ansCJKAfter + `])`)
	ansCJK = regexp.MustCompile(`([` + ansBeforeCJK + `])([` + cjk + `])`)

	percentAlpha = regexp.MustCompile(`(%)([` + alpha + `])`)
	// The copyright sign reads with a space before a year
	copyrightDigit = regexp.MustCompile(`(\x{00a9})([0-9])`)

	// Only a lone, tight middle dot converts: a run is a mask, a spaced one is the author's separator
	middleDot = regexp.MustCompile(`[\x{00b7}\x{2022}\x{2027}]`)

	// Hyphen-joined runs that read as one name: state-of-the-art, GPT-4o, claude-4-opus. An all-uppercase pair like ABC-DEF does not qualify
	compoundWord = regexp.MustCompile(`\b(?:[A-Za-z0-9]*[a-z][A-Za-z0-9]*-[A-Za-z0-9]+|[A-Za-z0-9]+-[A-Za-z0-9]*[a-z][A-Za-z0-9]*|[A-Za-z]+-[0-9]+|[A-Za-z]+[0-9]+-[A-Za-z0-9]+)(?:-[A-Za-z0-9]+)*\b`)

	backtickContent = regexp.MustCompile("`([^`]+)`")

	// Only opening, closing, and self-closing tags with a real tag name, so stray < > is not read as a tag
	htmlTag        = regexp.MustCompile(`</?[a-zA-Z][a-zA-Z0-9]*(?:[` + jsSpace + `]+[^>]*)?>`)
	htmlAttr       = regexp.MustCompile(`(\w+)="([^"]*)"`)
	closingHTMLTag = regexp.MustCompile(`</([a-zA-Z][a-zA-Z0-9]*)`)
	// A bare unpaired non-void tag amid prose is a tag mention, not markup. See ADR 0005
	bareHTMLTag       = regexp.MustCompile(`^<([a-zA-Z][a-zA-Z0-9]*)[` + jsSpace + `]*/?>$`)
	cjkHTMLTagMention = regexp.MustCompile(`([` + cjk + `])(` + htmlTagMentionStart + `)`)
	htmlTagMentionCJK = regexp.MustCompile(`(` + htmlTagMentionEnd + `)([` + cjk + `])`)

	// Scheme-anchored, and CJK continues the URL. The body stops at the Private Use Area so a URL never swallows a placeholder. See ADR 0026
	httpURL                    = regexp.MustCompile(`https?://[^` + jsSpace + `<>"` + "`" + `\x{3000}-\x{303f}\x{ff00}-\x{ffef}\x{2018}\x{2019}\x{201c}\x{201d}\x{2026}\x{e000}-\x{f8ff}]+`)
	httpURLTrailingPunctuation = regexp.MustCompile(`[.,;:!?'"]+$`)
	cjkHTTPURL                 = regexp.MustCompile(`([` + cjk + `])(` + httpURLStart + `)`)

	// Brackets whose inner edge spaces get stripped
	bracketPairs = []*regexp.Regexp{
		regexp.MustCompile(`<[^<>]*>`),
		regexp.MustCompile(`\([^()]*\)`),
		regexp.MustCompile(`\[[^\[\]]*\]`),
		regexp.MustCompile(`\{[^{}]*\}`),
	}
)

var voidHTMLTags = map[string]bool{
	"area": true, "base": true, "br": true, "col": true, "embed": true, "hr": true, "img": true,
	"input": true, "link": true, "meta": true, "param": true, "source": true, "track": true, "wbr": true,
}

// charClass returns a predicate for one rune against a bracket-expression body
func charClass(class string) func(rune) bool {
	re := regexp.MustCompile(`^[` + class + `]$`)
	return func(r rune) bool { return re.MatchString(string(r)) }
}
