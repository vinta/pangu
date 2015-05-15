package pangu_test

import (
	"github.com/vinta/pangu"
	"testing"
)

func BenchmarkTextSpacing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pangu.TextSpacing("所以,請問Jackey的鼻子有幾個?3.14個!")
	}
}

func BenchmarkFileSpacing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleFileSpacing()
	}
}
