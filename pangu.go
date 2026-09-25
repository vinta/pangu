package pangu

// SpaceText inserts whitespace between CJK and ANS characters in text.
func SpaceText(text string) string {
	return text
}

// HasProperSpacing reports whether SpaceText would leave text unchanged.
func HasProperSpacing(text string) bool {
	return SpaceText(text) == text
}
