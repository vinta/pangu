package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestText(t *testing.T) {
	realStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	os.Args = []string{NAME, "text", "新八的構造成分有95%是眼鏡、3%是水、2%是垃圾"}
	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = realStdout
	fmt.Printf("%s", out)
}
