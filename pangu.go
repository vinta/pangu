package pangu

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

// SpaceText inserts whitespace between CJK and ANS characters in text.
func SpaceText(text string) string {
	if !anyCJK.MatchString(text) {
		return text
	}

	// Hide backtick content from the quote rules; the backticks themselves still get spacing
	backticks := &placeholders{placeholderKind: backtickPlaceholder}
	text = replaceAll(backtickContent, text, func(m []int) string {
		return "`" + backticks.store(text[m[2]:m[3]]) + "`"
	})

	// Hide every URL from the rules
	urls := &placeholders{placeholderKind: httpURLPlaceholder}
	text = replaceLookaround(httpURL, text, func(s string, m []int) bool {
		r, ok := runeBefore(s, m[0])
		return !ok || !isASCIIAlnum(r)
	}, func(s string, m []int) string {
		match := s[m[0]:m[1]]
		url := trimHTTPURL(match)
		return urls.store(url) + match[len(url):]
	})

	htmlTags := &placeholders{placeholderKind: htmlTagPlaceholder}
	mentionedTags := &placeholders{placeholderKind: htmlTagMentionPlaceholder}
	hasHTMLTags := strings.Contains(text, "<")
	if hasHTMLTags {
		// Tag names whose closing tag appears anywhere in the text: their opening tags are paired markup
		closedTagNames := map[string]bool{}
		for _, m := range closingHTMLTag.FindAllStringSubmatch(text, -1) {
			closedTagNames[strings.ToLower(m[1])] = true
		}
		// Hide every real tag behind a placeholder; attribute values get spacing first
		text = htmlTag.ReplaceAllStringFunc(text, func(tag string) string {
			if m := bareHTMLTag.FindStringSubmatch(tag); m != nil {
				name := strings.ToLower(m[1])
				if !voidHTMLTags[name] && !closedTagNames[name] {
					return mentionedTags.store(tag)
				}
			}
			return htmlTags.store(replaceAll(htmlAttr, tag, func(m []int) string {
				return tag[m[2]:m[3]] + `="` + SpaceText(tag[m[4]:m[5]]) + `"`
			}))
		})
	}

	// Middle dots go before the spacing rules, which would space a tight one as ANS
	text = replaceLookaround(middleDot, text, func(s string, m []int) bool {
		return !runeBeforeIs(s, m[0], isMiddleDotGap) && !runeAfterIs(s, m[1], isMiddleDotGap)
	}, literal("\u30fb"))

	// Dot runs go first, before the single-period rule
	text = dotsCJK.ReplaceAllString(text, "${1} ${2}")

	text = replaceLookaround(cjkPunctuation, text, nextIs(isCJKOrAlnum), expand(cjkPunctuation, "${1}${2} "))
	text = replaceLookaround(punctuationCJK, text, nextIs(isCJK), expand(punctuationCJK, "${0} "))
	text = replaceLookaround(cjkTilde, text, nextIs(isCJKOrAlnum), expand(cjkTilde, "${1}${2} "))
	text = cjkTildeEquals.ReplaceAllString(text, "${1} ${2} ")
	text = replaceLookaround(cjkPeriod, text, nextIs(isCJK), expand(cjkPeriod, "${1}${2} "))
	text = anPeriodCJK.ReplaceAllString(text, "${1}${2} ${3}")
	text = anColonCJK.ReplaceAllString(text, "${1}${2} ${3}")
	text = fixCJKColonANS.ReplaceAllString(text, "${1}\uff1a${2}")

	text = cjkQuote.ReplaceAllString(text, "${1} ${2}")
	text = quoteCJK.ReplaceAllString(text, "${1} ${2}")
	text = fixQuoteAnyQuote.ReplaceAllString(text, "${1}${2}${3}")

	text = quoteAN.ReplaceAllString(text, "${1} ${2}")
	text = cjkQuoteAN.ReplaceAllString(text, "${1}${2} ${3}")

	text = fixPossessiveSingleQuote.ReplaceAllString(text, "${1}'s")

	// Quoted pure-CJK content keeps its quotes tight, so hide it before the single-quote rules run
	singleQuotes := &placeholders{placeholderKind: singleQuotePlaceholder}
	text = singleQuotePureCJK.ReplaceAllStringFunc(text, singleQuotes.store)
	text = cjkSingleQuoteButPossessive.ReplaceAllString(text, "${1} ${2}")
	text = singleQuoteCJK.ReplaceAllString(text, "${1} ${2}")
	text = singleQuotes.restore(text)

	text = hashCJKHash.ReplaceAllString(text, "${1} ${2}${3}${4} ${5}")
	text = cjkHash.ReplaceAllString(text, "${1} ${2}")
	text = hashCJK.ReplaceAllString(text, "${1} ${2}")

	// Protect compound words from operator spacing
	compoundWords := &placeholders{placeholderKind: compoundWordPlaceholder}
	text = compoundWord.ReplaceAllStringFunc(text, compoundWords.store)

	// Single-letter grades run before the operator rules so A+CJK becomes A+ CJK, not A + CJK
	text = singleLetterGradeCJK.ReplaceAllString(text, "${1}${2} ${3}")

	// Affixes run before the operator rules so the symbol stays attached to its half-width side
	text = cjkSignDigit.ReplaceAllString(text, "${1} ${2}${3}")
	text = cjkHyphenFlag.ReplaceAllString(text, "${1} ${2}${3}")
	text = digitPlusCJK.ReplaceAllString(text, "${1}${2} ${3}")

	// Plus reading is per line: a plus in direct contact with CJK makes every undecided plus on the line a separator. A decided plus keeps its reading: space-adjacent,
	// affix-attached (100+, +886), or in a ++ run (C++). It runs before the operator rules, so a CJK+A contact flips the line's joiners like CJK+CJK does
	text = mapLines(text, func(line string) string {
		if plusCJKContact.MatchString(line) {
			line = replaceLookaround(plus, line, func(s string, m []int) bool {
				return runeBeforeIs(s, m[0], not(isPlusNeighbor)) && runeAfterIs(s, m[1], not(isPlusNeighbor))
			}, func(s string, m []int) string {
				// Read through compound placeholders to recognize names such as non-Disney+
				if !endsWithNameSuffix(compoundWords.restore(s[:m[1]])) {
					return " + "
				}
				if runeAfterIs(s, m[1], isClosingAfterSuffix) {
					return "+"
				}
				return "+ "
			})
		}
		// A closing bracket cannot carry a name suffix, so before a full-width opener the separator space goes on the closing-bracket side only
		return replaceLookaround(plus, line, func(s string, m []int) bool {
			return runeBeforeIs(s, m[0], isRightBracket) && runeAfterIs(s, m[1], isFullWidthLeft)
		}, literal(" +"))
	})

	// Hyphen reading is per line and runs before the operator rules space the CJK contact away. Only a hyphen between a closing and an opening bracket flips
	text = mapLines(text, func(line string) string {
		if !hyphenCJKContact.MatchString(line) {
			return line
		}
		return replaceLookaround(hyphen, line, func(s string, m []int) bool {
			return runeBeforeIs(s, m[0], isRightBracket) && runeAfterIs(s, m[1], isLeftBracket)
		}, literal(" - "))
	})

	// An asterisk before a square bracket opens a bracket glob (*[0-9].log), so it stays tight. See ADR 0033
	text = replaceLookaround(cjkOperatorANS, text, func(s string, m []int) bool {
		return s[m[4]:m[5]] != "*" || s[m[6]:m[7]] != "["
	}, expand(cjkOperatorANS, "${1} ${2} ${3}"))
	// Listed name suffixes keep their signs attached
	text = replaceLookaround(ansOperatorCJK, text, func(s string, m []int) bool {
		return !endsWithNameSuffix(s[:m[5]])
	}, expand(ansOperatorCJK, "${1} ${2} ${3}"))

	text = cjkLessThan.ReplaceAllString(text, "${1} ${2} ${3}")
	text = lessThanCJK.ReplaceAllString(text, "${1} ${2} ${3}")
	text = cjkGreaterThan.ReplaceAllString(text, "${1} ${2} ${3}")
	text = greaterThanCJK.ReplaceAllString(text, "${1} ${2} ${3}")

	text = cjkUnixAbsolutePath.ReplaceAllString(text, "${1} ${2}")
	text = cjkUnixRelativePath.ReplaceAllString(text, "${1} ${2}")
	text = cjkWindowsPath.ReplaceAllString(text, "${1} ${2}")
	text = unixAbsolutePathSlashCJK.ReplaceAllString(text, "${1} ${2}")
	text = unixRelativePathSlashCJK.ReplaceAllString(text, "${1} ${2}")

	// Pipe reading is per line: a pipe in direct CJK contact makes every pipe on the line a separator (CJK | CJK, as in concatenated page titles)
	text = mapLines(text, func(line string) string {
		if !pipeCJKContact.MatchString(line) {
			return line
		}
		return replaceLookaround(pipeSeparator, line, nextIs(not(isSpaceOrPipe)), expand(pipeSeparator, "${1} ${2} "))
	})

	// A pipe or plus separator space can land just inside a closing quote; strip it again so a second pass changes nothing
	text = fixQuoteAnyQuote.ReplaceAllString(text, "${1}${2}${3}")

	text = compoundWords.restore(text)

	text = cjkLeftBracket.ReplaceAllString(text, "${1} ${2}")
	text = rightBracketCJK.ReplaceAllString(text, "${1} ${2}")
	text = ansCJKLeftQuoteAnyRightQuote.ReplaceAllString(text, "${1} ${2}${3}${4}")
	text = leftQuoteAnyRightQuoteANSCJK.ReplaceAllString(text, "${1}${2}${3} ${4}")
	// A right quote opens a pair only when no unclosed left quote precedes it on the line
	text = replaceLookaround(ansCJKRightQuoteAnyRightQuote, text, func(s string, m []int) bool {
		return !hasUnclosedLeftQuote(s[:m[4]])
	}, expand(ansCJKRightQuoteAnyRightQuote, "${1} ${2}${3}${4}"))

	// A dotted name keeps its call parenthesis tight (Math.floor(x)), a bare name does not (foo (x))
	text = replaceLookaround(anLeftBracket, text, func(s string, m []int) bool {
		return !isDottedName(s[:m[3]])
	}, expand(anLeftBracket, "${1} ${2}"))
	text = rightBracketAN.ReplaceAllString(text, "${1} ${2}")

	text = replaceLookaround(cjkANS, text, func(s string, m []int) bool {
		r, _ := utf8.DecodeRuneInString(s[m[4]:])
		return !isSuperscript(r) && !endsWithNameSuffix(s[:m[5]])
	}, expand(cjkANS, "${1} ${2}"))
	text = ansCJK.ReplaceAllString(text, "${1} ${2}")

	text = percentAlpha.ReplaceAllString(text, "${1} ${2}")
	text = copyrightDigit.ReplaceAllString(text, "${1} ${2}")

	text = fixBracketSpacing(text)

	if hasHTMLTags {
		text = cjkHTMLTagMention.ReplaceAllString(text, "${1} ${2}")
		text = htmlTagMentionCJK.ReplaceAllString(text, "${1} ${2}")
		text = mentionedTags.restore(text)
		text = htmlTags.restore(text)
	}

	text = cjkHTTPURL.ReplaceAllString(text, "${1} ${2}")
	text = urls.restore(text)
	return backticks.restore(text)
}

// HasProperSpacing reports whether SpaceText would leave text unchanged.
func HasProperSpacing(text string) bool {
	return SpaceText(text) == text
}

// trimHTTPURL drops trailing punctuation and unbalanced closing parentheses, which belong to the prose, not the URL
func trimHTTPURL(url string) string {
	unbalanced := strings.Count(url, ")") - strings.Count(url, "(")
	for {
		trimmed := httpURLTrailingPunctuation.ReplaceAllString(url, "")
		if strings.HasSuffix(trimmed, ")") && unbalanced > 0 {
			unbalanced--
			url = trimmed[:len(trimmed)-1]
			continue
		}
		if trimmed == url {
			return url
		}
		url = trimmed
	}
}

// fixBracketSpacing strips the spaces just inside a bracket pair
func fixBracketSpacing(text string) string {
	for _, pair := range bracketPairs {
		text = pair.ReplaceAllStringFunc(text, func(m string) string {
			return m[:1] + strings.Trim(m[1:len(m)-1], " ") + m[len(m)-1:]
		})
	}
	return text
}

// hasUnclosedLeftQuote reports whether a left curly quote opens on the current line of text and is not closed yet
func hasUnclosedLeftQuote(text string) bool {
	for i := strings.LastIndexAny(text, "\u201c\u201d\n"); i >= 0; {
		return strings.HasPrefix(text[i:], "\u201c")
	}
	return false
}

// isDottedName reports whether text ends with a dot followed by letters and digits
func isDottedName(text string) bool {
	trimmed := strings.TrimRightFunc(text, isASCIIAlnum)
	return strings.HasSuffix(trimmed, ".")
}

func isASCIIAlnum(r rune) bool {
	return r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <= '9'
}

// placeholderKind is one family of Private Use Area markers that hide matched text so later rules cannot touch it
type placeholderKind struct {
	start, name, end string
	pattern          *regexp.Regexp
}

func newPlaceholderKind(start, name, end string) *placeholderKind {
	return &placeholderKind{start, name, end, regexp.MustCompile(start + name + `(\d+)` + end)}
}

var (
	backtickPlaceholder       = newPlaceholderKind("\ue000", "BACKTICK_CONTENT_", "\ue001")
	htmlTagPlaceholder        = newPlaceholderKind("\ue002", "HTML_TAG_PLACEHOLDER_", "\ue003")
	htmlTagMentionPlaceholder = newPlaceholderKind(htmlTagMentionStart, "HTML_TAG_MENTION_", htmlTagMentionEnd)
	singleQuotePlaceholder    = newPlaceholderKind("\ue006", "SINGLE_QUOTE_CJK_PLACEHOLDER_", "\ue007")
	compoundWordPlaceholder   = newPlaceholderKind("\ue008", "COMPOUND_WORD_PLACEHOLDER_", "\ue009")
	httpURLPlaceholder        = newPlaceholderKind(httpURLStart, "HTTP_URL_PLACEHOLDER_", "\ue00b")
)

// placeholders holds the text that one SpaceText call hid behind one kind of placeholder
type placeholders struct {
	*placeholderKind
	items []string
}

func (p *placeholders) store(item string) string {
	p.items = append(p.items, item)
	return p.start + p.name + strconv.Itoa(len(p.items)-1) + p.end
}

func (p *placeholders) restore(text string) string {
	if len(p.items) == 0 {
		return text
	}
	return replaceAll(p.pattern, text, func(m []int) string {
		i, _ := strconv.Atoi(text[m[2]:m[3]])
		if i < len(p.items) {
			return p.items[i]
		}
		return ""
	})
}

// replaceAll is ReplaceAllStringFunc with submatch indices
func replaceAll(re *regexp.Regexp, text string, repl func(m []int) string) string {
	var b strings.Builder
	last := 0
	for _, m := range re.FindAllStringSubmatchIndex(text, -1) {
		b.WriteString(text[last:m[0]])
		b.WriteString(repl(m))
		last = m[1]
	}
	b.WriteString(text[last:])
	return b.String()
}

// replaceLookaround stands in for a JavaScript regex whose lookarounds became accept, which checks a match against the whole text. A rejected match is retried one
// character later, the way a JavaScript regex scans, instead of being skipped whole. The regex must not start with \b or ^, which lose their left context here
func replaceLookaround(re *regexp.Regexp, text string, accept func(s string, m []int) bool, repl func(s string, m []int) string) string {
	var b strings.Builder
	last, pos := 0, 0
	for pos <= len(text) {
		m := re.FindStringSubmatchIndex(text[pos:])
		if m == nil {
			break
		}
		for i := range m {
			if m[i] >= 0 {
				m[i] += pos
			}
		}
		if !accept(text, m) {
			_, size := utf8.DecodeRuneInString(text[m[0]:])
			pos = m[0] + max(size, 1)
			continue
		}
		b.WriteString(text[last:m[0]])
		b.WriteString(repl(text, m))
		last, pos = m[1], m[1]
	}
	b.WriteString(text[last:])
	return b.String()
}

func expand(re *regexp.Regexp, template string) func(s string, m []int) string {
	return func(s string, m []int) string {
		return string(re.ExpandString(nil, template, s, m))
	}
}

func literal(repl string) func(s string, m []int) string {
	return func(string, []int) string { return repl }
}

// nextIs accepts a match that the given class follows directly
func nextIs(class func(rune) bool) func(s string, m []int) bool {
	return func(s string, m []int) bool { return runeAfterIs(s, m[1], class) }
}

func not(class func(rune) bool) func(rune) bool {
	return func(r rune) bool { return !class(r) }
}

// runeAfterIs reports whether the rune starting at i exists and is in class
func runeAfterIs(s string, i int, class func(rune) bool) bool {
	if i >= len(s) {
		return false
	}
	r, _ := utf8.DecodeRuneInString(s[i:])
	return class(r)
}

// runeBeforeIs reports whether the rune ending at i exists and is in class
func runeBeforeIs(s string, i int, class func(rune) bool) bool {
	r, ok := runeBefore(s, i)
	return ok && class(r)
}

func runeBefore(s string, i int) (rune, bool) {
	if i <= 0 {
		return 0, false
	}
	r, _ := utf8.DecodeLastRuneInString(s[:i])
	return r, true
}

// mapLines applies f to each line of text
func mapLines(text string, f func(string) string) string {
	var b strings.Builder
	first := true
	for line := range strings.SplitSeq(text, "\n") {
		if !first {
			b.WriteByte('\n')
		}
		first = false
		b.WriteString(f(line))
	}
	return b.String()
}

// endsWithNameSuffix reports whether text ends with a name suffix. It reads only the last 32 bytes: the longest suffix plus the rune before it is 14, so the cut never splits a
// match, and it keeps a full-text regex scan out of per-match checks, which would make SpaceText quadratic
func endsWithNameSuffix(text string) bool {
	if i := len(text) - 32; i > 0 {
		for !utf8.RuneStart(text[i]) {
			i++
		}
		text = text[i:]
	}
	return nameSuffixAtEnd.MatchString(text)
}
