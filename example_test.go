package pangu_test

import (
	"fmt"
	"github.com/vinta/pangu"
	"os"
)

func checkErr(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ExampleTextSpacing() {
	s := pangu.TextSpacing("所以,請問Jackey的鼻子有幾個?3.14個!")
	fmt.Println(s)
	// Output:
	// 所以, 請問 Jackey 的鼻子有幾個? 3.14 個!
}

func ExampleFileSpacing() {
	outPath, err := pangu.FileSpacing("_fixtures/test_file.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(outPath)
	// Output:
	// test_file.pangu.txt
}
