package pathologize

import (
	"testing"
)

func BenchmarkCleanOriginal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clean("valid-filename.txt")
		Clean("invalid:file*name?.txt")
		Clean("CON.txt")
	}
}
