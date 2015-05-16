package main

import (
	"os"
	"testing"
)

func TestText(t *testing.T) {
	os.Args = []string{NAME, "text", "新八的構造成分有95%是眼鏡、3%是水、2%是垃圾"}
	main()
}
