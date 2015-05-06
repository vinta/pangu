package pangu_test

import (
	"testing"
)

func BenchmarkFileSpacing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExampleFileSpacing()
	}
}
