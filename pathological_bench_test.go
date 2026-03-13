package pathologize

import "testing"

func BenchmarkClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Clean("CON.")
		Clean("some_normal_file_name.txt")
		Clean("LPT1.txt")
	}
}

func BenchmarkRemoveTrailing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeTrailing("foo.bar.baz.  ")
	}
}

func BenchmarkRemoveReservedNames(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeReservedNames("LPT1")
	}
}
